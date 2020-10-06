package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	// feed.Add(Item{"Item added?", "The body"})
	if len(feed.Items) == 0 {
		t.Errorf("No item added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	results := feed.GetAll()
	if len(results) != 1 {
		t.Errorf("No items")
	}
}