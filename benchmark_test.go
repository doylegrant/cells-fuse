package main

import (
	"sync"
	"testing"
)

type DummyNode struct {
	Path      string
	MetaStore map[string]string
}

type DummyStat struct{}

func BenchmarkReaddirUnoptimized(b *testing.B) {
	nodes := make([]*DummyNode, 1000)
	for i := range nodes {
		nodes[i] = &DummyNode{Path: "path"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		children := make(map[string]*DummyStat)
		var mu sync.Mutex
		var wg sync.WaitGroup
		semaphore := make(chan struct{}, 10)

		for _, node := range nodes {
			wg.Add(1)
			semaphore <- struct{}{}
			go func(n *DummyNode) {
				defer wg.Done()
				defer func() { <-semaphore }()

				name := n.Path

				stat := &DummyStat{}

				mu.Lock()
				children[name] = stat
				mu.Unlock()
			}(node)
		}
		wg.Wait()
	}
}

func BenchmarkReaddirOptimized(b *testing.B) {
	nodes := make([]*DummyNode, 1000)
	for i := range nodes {
		nodes[i] = &DummyNode{Path: "path"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		children := make(map[string]*DummyStat, len(nodes))
		var mu sync.Mutex
		var wg sync.WaitGroup
		semaphore := make(chan struct{}, 10)

		for _, node := range nodes {
			wg.Add(1)
			semaphore <- struct{}{}
			go func(n *DummyNode) {
				defer wg.Done()
				defer func() { <-semaphore }()

				name := n.Path

				stat := &DummyStat{}

				mu.Lock()
				children[name] = stat
				mu.Unlock()
			}(node)
		}
		wg.Wait()
	}
}
