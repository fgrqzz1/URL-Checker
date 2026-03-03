package models

import "errors"

func (i *InputData) Validate() error {
	if i.ShowHelp {
		return nil
	}

	if len(i.URL) == 0 && i.FilePath == "" {
		return errors.New("Не указаны URL для проверки")
	}

	if i.Timeout <= 0 {
		return errors.New("Не валидный таймаут | > 0")
	}

	return nil
}

