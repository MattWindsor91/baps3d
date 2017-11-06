package list_test

import (
	"fmt"

	"github.com/UniversityRadioYork/baps3d/list"
)

func ExampleNew() {
	l := list.New()

	fmt.Println(l.Count())

	idx, _ := l.Selection()
	fmt.Println(idx)

	fmt.Println(l.AutoMode())

	// Output:
	// 0
	// -1
	// off
}

// ExampleList_SetAutoMode tests List.SetAutoMode in an example style.
func ExampleList_SetAutoMode() {
	l := list.New()

	l.SetAutoMode(list.AutoOff)
	fmt.Println(l.AutoMode())

	l.SetAutoMode(list.AutoDrop)
	fmt.Println(l.AutoMode())

	l.SetAutoMode(list.AutoNext)
	fmt.Println(l.AutoMode())

	changed := l.SetAutoMode(list.AutoShuffle)
	fmt.Println(l.AutoMode(), changed)

	changedAgain := l.SetAutoMode(list.AutoShuffle)
	fmt.Println(l.AutoMode(), changedAgain)

	// Output:
	// off
	// drop
	// next
	// shuffle true
	// shuffle false
}

// ExampleList_Selection tests List.Selection in an example style.
func ExampleList_Selection() {
	// New lists have no selection.
	l := list.New()

	idx, _ := l.Selection()
	fmt.Println(idx)

	// If we change the selection, Selection updates.
	if err := l.Add(list.NewTrack("xyz", "foo.mp3"), 0); err != nil {
		panic(err)
	}
	if _, err := l.Select(0, "xyz"); err != nil {
		panic(err)
	}

	idx, _ = l.Selection()
	fmt.Println(idx)

	// Output:
	// -1
	// 0
}

// ExampleList_Freeze tests List.Freeze in an example style.
func ExampleList_Freeze() {
	l := list.New()

	if err := l.Add(list.NewTrack("abc", "foo.mp3"), 0); err != nil {
		panic(err)
	}
	if err := l.Add(list.NewTrack("xyz", "bar.mp3"), 1); err != nil {
		panic(err)
	}

	items := l.Freeze()

	// 'items' will remain the same even if we add a new item.
	if err := l.Add(list.NewText("def", "baz"), 1); err != nil {
		panic(err)
	}

	for _, item := range items {
		fmt.Println(item.Hash())
	}

	// Output:
	// abc
	// xyz
}
