package validators

import "golang.org/x/exp/slices"

func OptionValidator(options []string,value string) bool{
	return slices.Contains(options,value)
}

func MinimumValueValidator(value int ,limit int) bool {
	return value <= limit
}

