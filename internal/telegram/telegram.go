package telegram

type Bot struct {
	api string
}

func New(token string) *Bot {
	return &Bot{
		api: token,
	}
}

func (b *Bot) Run() {
	return
}
