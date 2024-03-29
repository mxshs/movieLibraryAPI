package mock_db

import (
	"mxshs/movieLibrary/src/domain"
	"sync/atomic"
)

type MockDB struct {
    aid atomic.Int64
    mid atomic.Int64
    uid atomic.Int64

    actors map[int]*domain.Actor
    movies map[int]*domain.Movie
    users map[int]*domain.UserDetail

    movieActors map[int]*LL
    actorMovies map[int]*LL
}

func NewDB() *MockDB {
    db := &MockDB{}
    db.actors = make(map[int]*domain.Actor)
    db.movies = make(map[int]*domain.Movie)
    db.users = make(map[int]*domain.UserDetail)
    db.movieActors = make(map[int]*LL)
    db.actorMovies = make(map[int]*LL)

    return db
}

type Node struct {
    id int
    next *Node
}

type LL struct {
    head *Node
    tail *Node
}

func (l *LL) Add(value int) {
    if (l.head == nil) {
        l.head = &Node{value, nil}
        l.tail = l.head
    } else {
        l.tail.next = &Node{value, nil}
        l.tail = l.tail.next
    }
}

func (l *LL) Remove(value int) {
    if l.head == nil {
        return
    } else if l.head.id == value {
        l.head = l.head.next
        return
    }

    ptr := l.head
    for (ptr.next != nil && ptr.next.id != value) {
        ptr = ptr.next
    }

    if ptr.next != nil {
        ptr.next = ptr.next.next
    }
}
