package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type keyMapping struct {
	key rune
	m   model
}

var keyMappings = map[model]map[string]keyMapping{
	collectionSelectorModel: {
		"J": {key: 'j', m: todoListModel},
		"K": {key: 'k', m: todoListModel},
		" ": {key: ' ', m: todoListModel},
	},
}

func mapMessage(m model, key tea.KeyMsg) (model, tea.KeyMsg) {
	if mappedCollection, ok := keyMappings[m]; ok {
		if mappedRune, ok := mappedCollection[key.String()]; ok {
			key.Runes = []rune{mappedRune.key}
			return mappedRune.m, key
		}
	}
	return m, key
}
