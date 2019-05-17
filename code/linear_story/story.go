// Linked List Data Structure
// Recursion

//Functions - has a return value, may also execute commands
//Procedures - has no return value, just executes commands
//Methods - functions attached to a struct/object/etc

package main

import "fmt"

type storyPage struct {
	text     string
	nextPage *storyPage
}

// Recursion
// func playStory(page *storyPage) {
// 	if page == nil {
// 		return
// 	}
// 	fmt.Println(page.text)
// 	playStory(page.nextPage)
// }

// (page *storyPage) == receiver type
// func (page *storyPage) playStory() {
// 	if page == nil {
// 		return
// 	}
// 	fmt.Println(page.text)
// 	page.nextPage.playStory()
// }

// For Loop to run the game
func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}

// Add page to middle
// Receive current page
// Create a new page with the given text and set its next page to the current page's next page
// Set the current page's next page to the new page
func (page *storyPage) addAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

// Add page to end
// Traverse the pages until you hit one with a NIL nextpage
func (page *storyPage) addToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = &storyPage{text, nil}

}

func main() {
	page1 := storyPage{"It was a dark and stormy night.", nil}
	page1.addToEnd("You are alone and you need to find the sacred helmet before the bad guys do.")
	page1.addToEnd("You see a troll ahead")
	page1.addAfter("Testing Page After")
	page1.playStory()

}

