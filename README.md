# GOregon

REST api with gin

test post route in postman, curl, console or wherever

Example:
Start api with - `make dev` //runs the makefile containing instructions for running main.go

console test: 
await fetch('/newsfeed', {
  method: 'POST',
  headers: {
    "content-type": "application/json"
  },
  body: JSON.stringify({
    title: "some title",
    post: "the body of the post"
  })
})
