package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/somatom98/todoist/domain"
)

type keyMapping struct {
	key rune
	m   domain.View
}

var keyMappings = map[domain.View]map[string]keyMapping{
	domain.ViewCollectionSelector: {
		"J": {key: 'j', m: domain.ViewTaskList},
		"K": {key: 'k', m: domain.ViewTaskList},
		" ": {key: ' ', m: domain.ViewTaskList},
		"a": {key: 'a', m: domain.ViewTaskList},
		"c": {key: 'c', m: domain.ViewTaskList},
	},
}

func mapMessage(m domain.View, key tea.KeyMsg) (domain.View, tea.KeyMsg) {
	if mappedCollection, ok := keyMappings[m]; ok {
		if mappedRune, ok := mappedCollection[key.String()]; ok {
			key.Runes = []rune{mappedRune.key}
			return mappedRune.m, key
		}
	}
	return m, key
}
