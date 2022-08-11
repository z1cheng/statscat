package stats

import (
	"math/rand"

	"github.com/pterm/pterm"
)

func NewInfoSpinner(msg string) *pterm.SpinnerPrinter {
	spinner, _ := pterm.DefaultSpinner.Start(msg)
	spinner.InfoPrinter = &pterm.PrefixPrinter{
		MessageStyle: &pterm.ThemeDefault.DefaultText,
		Prefix: pterm.Prefix{
			Style: &pterm.ThemeDefault.InfoPrefixStyle,
			Text:  randomCatEmoji(),
		},
	}
	return spinner
}

func init() {
	pterm.Success.MessageStyle = &pterm.ThemeDefault.DefaultText
	pterm.Success.Prefix = pterm.Prefix{
		Style: &pterm.ThemeDefault.SuccessPrefixStyle,
		Text:  "STATS CAT🐈",
	}

	pterm.Error.Prefix = pterm.Prefix{
		Style: &pterm.ThemeDefault.ErrorPrefixStyle,
		Text:  "ERROR🙀",
	}
}

// randomCatEmoji returns a random cat emoji.
func randomCatEmoji() string {
	emojis := []rune("😺😸😹😻😼😽🙀😿😾🐱")
	return string(emojis[rand.Intn(len(emojis))])
}
