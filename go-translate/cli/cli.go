package cli

import "net/http"

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

const translateURL = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(body *RequestBody) {
	client := &http.Client{}
	req
	client.Do(req)
}
