package main

import (
	"github.com/Wheeeel/todobot/hitoko"
	"strings"
	"math/rand"
)

func formatHitokotoRespMsg(args []string) (msg string) {
	if len(args) == 0 {
		msg = one()
		return
	}
	switch args[0] {
	case "-a":
		msg, _ = hitokoto(hitoko.TYPE_ANIME)
	case "-c":
		msg, _ = hitokoto(hitoko.TYPE_COMIC)
	case "-g":
		msg, _ = hitokoto(hitoko.TYPE_GAME)
	case "-n":
		msg, _ = hitokoto(hitoko.TYPE_NOVEL)
	case "-i":
		msg, _ = hitokoto(hitoko.TYPE_INTERNET)
	case "-o":
		msg, _ = hitokoto(hitoko.TYPE_OTHER)
	case "-m":
		msg = nya()
	default:
		msg = one()
	}
	return
}

func one() (koto string) {
	koto, _ = hitokoto(hitoko.TYPE_RANDOM)
	if len(koto) == 0 {
		koto = nya()
		return
	}
	return
}

func hitokoto(hitokotoType string) (koto string, by string) {
	r, err := hitoko.Fortune(hitokotoType)
	r.Hitokoto = strings.TrimSpace(r.Hitokoto)
	r.From = strings.TrimSpace(r.From)
	if err != nil || len(r.Hitokoto) == 0 {
		return
	}
	koto = r.Hitokoto
	by = r.From
	return
}

var nyaSuffix = []string{
	"",
	"~",
	"～",
	"？",
	"！",
	"……",
}

func nya() (koto string) {
	d := rand.Intn(6)
	if d < 4 { // 0 1 2 3
		o := d*2 + 1 // 1 3 5 7
		for i := 0; i < o; i++ {
			koto += "喵"
		}
	} else if d == 4 {
		koto += "喵呜"
	} else if d == 5 {
		koto += "咪呜"
	}
	koto += nyaSuffix[rand.Intn(len(nyaSuffix))]
	return
}
