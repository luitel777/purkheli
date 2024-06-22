package utils

import "errors"

// cases where title is wrong
// string is empty
// string contains only whitespace
// string length is more than 255

var TITLE_ERROR = "Title cannot be empty"
var TEXTAREA_ERROR = "Textarea cannot be empty"
var WHITESPACE_ERROR = "Whitespace only sentence is not allowed"
var MAXLENGTH_ERROR = "Length exceeds 255 characters"
var MAXLENGTH_ERROR_TEXTBOX = "Length exceeds 5000 characters"

func ValidateTitle(s string) error {
	if s == "" {
		return errors.New(TITLE_ERROR)
	}

	whitespaceCounter := 0
	for _, c := range s {
		if c == ' ' {
			whitespaceCounter++
		}
	}
	if whitespaceCounter == len(s) {
		return errors.New(WHITESPACE_ERROR)
	}

	if len(s) > 255 {
		return errors.New(MAXLENGTH_ERROR)
	}
	return nil
}

func ValidateTextArea(s string) error {
	if s == "" {
		return errors.New(TEXTAREA_ERROR)
	}

	whitespaceCounter := 0
	for _, c := range s {
		if c == ' ' {
			whitespaceCounter++
		}
	}
	if whitespaceCounter == len(s) {
		return errors.New(WHITESPACE_ERROR)
	}

	if len(s) > 5000 {
		return errors.New(MAXLENGTH_ERROR_TEXTBOX)
	}
	return nil
}
