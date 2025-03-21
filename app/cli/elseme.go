package cli

import (
	"encoding/json"
	. "wu"
)

type ElseMeCli struct {
	*App
	ExtraResult any `json:"m,omitempty"`
}

var elsemeCommand = Command{
	Name:      "elseme",
	ShortName: "",
	Usage:     "Elseme mode",
	Action:    doElseMe,
}

// Prints help for the App
func doElseMe(c *Context) {
	ElseMe(c.App)
}

func ElseMe(app *App) {
	elseme := &ElseMeCli{App: app}
	result, err := json.Marshal(elseme)
	ErrFatal(err)
	print(string(result))
}

func ElseMeExtra(app *App, extraResult any) {
	elseme := &ElseMeCli{App: app, ExtraResult: extraResult}
	result, err := json.Marshal(elseme)
	ErrFatal(err)
	print(string(result))
}

// func ElseMe(app *App, extraResult any) {
// 	elseme := &ElseMeCli{App: app, ExtraResult: extraResult}
// 	// json, err := gojson.MarshalIndent(elseme, "", "\t")
// 	json, err := gojson.Marshal(elseme)
// 	if err != nil {
// 		fmt.Println(fmt.Sprintf("json Marshal err: ", err))
// 	}
// 	print(string(json))
// }
