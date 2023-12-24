package snakegame

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	width  = 20
	height = 20
)

type point struct {
	x, y int
}

type snake struct {
	body  []point
	dir   point
	grow  bool
	score int
}

var food point

func SnakeGame() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	snake := newSnake()
	placeFood()

	gameOver := false

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	ticker := time.NewTicker(200 * time.Millisecond)

	for !gameOver {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyArrowUp:
					snake.setDirection(0, -1)
				case termbox.KeyArrowDown:
					snake.setDirection(0, 1)
				case termbox.KeyArrowLeft:
					snake.setDirection(-1, 0)
				case termbox.KeyArrowRight:
					snake.setDirection(1, 0)
				case termbox.KeyEsc:
					gameOver = true
				}
			}
		case <-ticker.C:
			gameOver = !snake.move()
			if snake.head() == food {
				snake.grow = true
				placeFood()
			}
		}

		render(snake)
	}

	fmt.Println("Game Over! Your score:", snake.score)
}

func newSnake() snake {
	head := point{width / 2, height / 2}
	return snake{
		body:  []point{head},
		dir:   point{1, 0},
		grow:  false,
		score: 0,
	}
}

func (s *snake) setDirection(dx, dy int) {
	if dx != -s.dir.x || dy != -s.dir.y {
		s.dir = point{dx, dy}
	}
}

func (s *snake) move() bool {
	head := s.head()
	newHead := point{head.x + s.dir.x, head.y + s.dir.y}

	if newHead.x < 0 || newHead.x >= width || newHead.y < 0 || newHead.y >= height {
		return false // Hit the wall
	}

	for _, bodyPart := range s.body {
		if bodyPart == newHead {
			return false // Hit itself
		}
	}

	s.body = append([]point{newHead}, s.body...)
	if !s.grow {
		s.body = s.body[:len(s.body)-1]
	} else {
		s.grow = false
		s.score++
	}

	return true
}

func (s *snake) head() point {
	return s.body[0]
}

func placeFood() {
	rand.Seed(time.Now().UnixNano())
	food = point{rand.Intn(width), rand.Intn(height)}
}

func render(s snake) {
	termbox.Clear(termbox.ColorWhite, termbox.ColorWhite)

	for _, p := range s.body {
		termbox.SetCell(p.x, p.y, '■', termbox.ColorDefault, termbox.ColorBlack)
	}

	termbox.SetCell(food.x, food.y, '★', termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}
