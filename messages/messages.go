package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"text/template"

	"github.com/bwmarrin/discordgo"

	"github.com/SteMak/vanilla/out"
)

type embed struct {
	Color       int64
	Title       *template.Template
	Description *template.Template
	Footer      *template.Template
}

type message struct {
	Content *template.Template
	Embed   *embed
}

var (
	msgs = make(map[string]*message)
)

func loadMessage(name string, data json.RawMessage) error {
	type config struct {
		Content *string `json:"content,omitempty"`
		Embed   *struct {
			Color       *string `json:"color,omitempty"`
			Title       *string `json:"title,omitempty"`
			Description *string `json:"description,omitempty"`
			Footer      *string `json:"footer,omitempty"`
		} `json:"embed,omitempty"`
	}

	var c config
	err := json.Unmarshal(data, &c)
	if err != nil {
		return err
	}

	var m message

	if c.Content != nil {
		data, err := ioutil.ReadFile(*c.Content)
		if err != nil {
			return err
		}
		m.Content, err = template.New(*c.Content).
			Funcs(funcs).
			Parse(string(data))

		if err != nil {
			return err
		}
	}

	if c.Embed != nil {
		m.Embed = new(embed)
		if c.Embed.Color != nil {
			m.Embed.Color, err = strconv.ParseInt(string([]rune(*c.Embed.Color)[1:]), 16, 32)
			if err != nil {
				return err
			}
		}

		if c.Embed.Title != nil {
			data, err := ioutil.ReadFile(*c.Embed.Title)
			if err != nil {
				return err
			}

			m.Embed.Title, err = template.New(*c.Embed.Title).
				Funcs(funcs).
				Parse(string(data))

			if err != nil {
				return err
			}

		}

		if c.Embed.Description != nil {
			data, err := ioutil.ReadFile(*c.Embed.Description)
			if err != nil {
				return err
			}

			m.Embed.Description, err = template.New(*c.Embed.Description).
				Funcs(funcs).
				Parse(string(data))

			if err != nil {
				return err
			}
		}

		if c.Embed.Footer != nil {
			m.Embed.Footer, err = template.New(*c.Embed.Footer).
				Funcs(funcs).
				ParseFiles(*c.Embed.Footer)

			if err != nil {
				return err
			}
		}
	}

	msgs[name] = &m
	return nil
}

func LoadMessage(path *string) error {
	if path == nil {
		return nil
	}

	var messages map[string]json.RawMessage
	data, err := ioutil.ReadFile(*path)
	if err != nil {
		out.Fatal(err)
	}

	err = json.Unmarshal(data, &messages)
	if err != nil {
		return err
	}

	for name, raw := range messages {
		if err := loadMessage(name, raw); err != nil {
			return err
		}
	}

	return nil
}

func Get(name string, data interface{}) (*discordgo.MessageSend, error) {
	m, ok := msgs[name]
	if !ok {
		return nil, fmt.Errorf("message '%s' no found", name)
	}

	str := func(t *template.Template) (string, error) {
		buf := bytes.NewBufferString("")
		err := t.Execute(buf, data)
		if err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	result := new(discordgo.MessageSend)

	if m.Content != nil {
		value, err := str(m.Content)
		if err != nil {
			return nil, err
		}
		result.Content = value
	}

	if m.Embed != nil {
		embed := new(discordgo.MessageEmbed)
		embed.Color = int(m.Embed.Color)

		if m.Embed.Title != nil {
			value, err := str(m.Embed.Title)
			if err != nil {
				return nil, err
			}
			embed.Title = value
		}

		if m.Embed.Description != nil {
			value, err := str(m.Embed.Description)
			if err != nil {
				return nil, err
			}
			embed.Description = value
		}

		if m.Embed.Footer != nil {
			value, err := str(m.Embed.Footer)
			if err != nil {
				return nil, err
			}

			embed.Footer = &discordgo.MessageEmbedFooter{
				Text: value,
			}
		}

		result.Embed = embed
	}

	return result, nil
}
