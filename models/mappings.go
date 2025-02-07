package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type keyMapping struct {
	key rune
	m   view
}

var keyMappings = map[view]map[string]keyMapping{
	viewCollectionSelector: {
		"J": {key: 'j', m: viewTodoList},
		"K": {key: 'k', m: viewTodoList},
		" ": {key: ' ', m: viewTodoList},
		"a": {key: 'a', m: viewTodoList},
	},
}

func mapMessage(m view, key tea.KeyMsg) (view, tea.KeyMsg) {
	if mappedCollection, ok := keyMappings[m]; ok {
		if mappedRune, ok := mappedCollection[key.String()]; ok {
			key.Runes = []rune{mappedRune.key}
			return mappedRune.m, key
		}
	}
	return m, key
}
