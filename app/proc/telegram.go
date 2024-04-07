package proc

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/denisbrodbeck/striphtmltags"
	log "github.com/go-pkgz/lgr"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/umputun/feed-master/app/feed"
)

// TelegramClient client
type TelegramClient struct {
	Bot             *tb.Bot
	Timeout         time.Duration
	DurationService DurationService
	TelegramSender  TelegramSender
}

// TelegramSender is the interface for sending messages to telegram
type TelegramSender interface {
	Send(tb.Audio, *tb.Bot, tb.Recipient, *tb.SendOptions) (*tb.Message, error)
}

// DurationService is the interface for reading duration from files
type DurationService interface {
	File(fname string) int
}

// NewTelegramClient init telegram client
func NewTelegramClient(token, apiURL string, timeout time.Duration, ds DurationService, tgs TelegramSender) (*TelegramClient, error) {
	log.Printf("[INFO] create telegram client for %s, timeout: %s", apiURL, timeout)
	if timeout == 0 {
		timeout = time.Second * 60
	}

	if token == "" {
		return &TelegramClient{
			Bot:     nil,
			Timeout: timeout,
		}, nil
	}

	bot, err := tb.NewBot(tb.Settings{
		URL:    apiURL,
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		Client: &http.Client{Timeout: timeout},
	})
	if err != nil {
		return nil, err
	}

	result := TelegramClient{
		Bot:             bot,
		Timeout:         timeout,
		DurationService: ds,
		TelegramSender:  tgs,
	}
	return &result, err
}

// Send message, skip if telegram token empty
func (client TelegramClient) Send(channelID string, feed feed.Rss2, item feed.Item) (err error) {
	if client.Bot == nil || channelID == "" {
		return nil
	}

	message, err := client.sendText(channelID, feed, item)
	if err != nil {
		return errors.Wrapf(err, "can't send to telegram for %+v", item.Enclosure)
	}

	log.Printf("[DEBUG] telegram message sent: \n%s", message.Text)
	return nil
}

func (client TelegramClient) sendText(channelID string, feed feed.Rss2, item feed.Item) (*tb.Message, error) {
	message, err := client.Bot.Send(
		recipient{chatID: channelID},
		client.getMessageHTML(feed, item, htmlMessageParams{WithMp3Link: true}),
		tb.ModeHTML,
		tb.NoPreview,
	)

	return message, err
}

// https://core.telegram.org/bots/api#html-style
func (client TelegramClient) tagLinkOnlySupport(htmlText string) string {
	p := bluemonday.NewPolicy()
	p.AllowAttrs("href").OnElements("a")
	return html.UnescapeString(p.Sanitize(htmlText))
}

type htmlMessageParams struct{ WithMp3Link, TrimCaption bool }

// getMessageHTML generates HTML message from provided feed.Item
func (client TelegramClient) getMessageHTML(feed feed.Rss2, item feed.Item, params htmlMessageParams) string {
	var header, footer string
	title := strings.TrimSpace(item.Title)
	if title != "" && item.Link == "" {
		header = fmt.Sprintf("%s\n\n", title)
	} else if title != "" && item.Link != "" {
		header = fmt.Sprintf("<a href=%q>%s</a>\n\n", item.Link, title)
	}

	feed.Title = strings.TrimSpace(feed.Title)
	feedTitle := fmt.Sprintf("<b>%s</b>\n\n", feed.Title)

	return feedTitle + header + footer
}

type recipient struct {
	chatID string
}

func (r recipient) Recipient() string {
	if _, err := strconv.ParseInt(r.chatID, 10, 64); err != nil && !strings.HasPrefix(r.chatID, "@") {
		return "@" + r.chatID
	}

	return r.chatID
}

// CropText shrinks the provided string, removing HTML tags in case it's exceeding the limit
func CropText(inp string, max int) string {
	if len([]rune(inp)) > max {
		return CleanText(inp, max)
	}
	return inp
}

// CleanText removes html tags and shrinks result
func CleanText(inp string, max int) string {
	res := striphtmltags.StripTags(inp)
	if len([]rune(res)) > max {
		// 4 symbols reserved for space and three dots on the end
		snippet := []rune(res)[:max-4]
		// go back in snippet and found first space
		for i := len(snippet) - 1; i >= 0; i-- {
			if snippet[i] == ' ' {
				snippet = snippet[:i]
				break
			}
		}
		res = string(snippet) + " ..."
	}
	return res
}

// TelegramSenderImpl is a TelegramSender implementation that sends messages to Telegram for real
type TelegramSenderImpl struct{}

// Send sends a message to Telegram
func (tg *TelegramSenderImpl) Send(audio tb.Audio, bot *tb.Bot, rcp tb.Recipient, opts *tb.SendOptions) (*tb.Message, error) {
	return audio.Send(bot, rcp, opts)
}
