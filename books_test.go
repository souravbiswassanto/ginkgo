package ginkgo_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/souravbiswassanto/ginkgo"
)

/*var _ = Describe("Books", func() {
	var foxInSocks, lesMis *ginkgo.Book

	BeforeEach(func() {
		lesMis = &ginkgo.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		foxInSocks = &ginkgo.Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing books", func() {
		Context("with more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(lesMis.Category()).To(Equal(ginkgo.CategoryNovel))
			})
		})

		Context("with fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(foxInSocks.Category()).To(Equal(ginkgo.CategoryShortStory))
			})
		})
	})
})*/

var _ = Describe("Books", func() {
	var book1, book2 *ginkgo.Book

	BeforeEach(func() {
		book1 = &ginkgo.Book{
			Title:  "Hello World Book",
			Author: "Saurov Chandra Biswas",
			Pages:  350,
		}
		book2 = &ginkgo.Book{
			Title:  "New Book",
			Author: "Sample Writter",
			Pages:  250,
		}
	})

	Describe("Categorizing Books", func() {
		Context("with more than 300 pages", func() {
			It("Should be a novel", func() {
				Expect(book1.Category()).To(Equal(ginkgo.CategoryNovel))
			})
		})
		Context("with less than 300 pages", func() {
			It("should be a short story book", func() {
				Expect(book2.Category()).To(Equal(ginkgo.CategoryShortStory))
			})
		})
	})
})

/*

=== INVALID ===
var _ = It("has a color", func() {
	Context("when blue", func() { // NO! Nodes can only be nested in containers
		It("is blue", func() { // NO! Nodes can only be nested in containers

		})
	})
})

=== INVALID ===
var _ = Describe("book", func() {
  var book *Book
  Expect(book.Title()).To(BeFalse()) // NO!  Place in a setup node instead.

  It("tests something", func() {...})
})
=== INVALID ===
var _ = Describe("book", func() {
  var book *Book
  Expect(book.Title()).To(BeFalse()) // NO!  Place in a setup node instead.

  It("tests something", func() {...})
})

*/
