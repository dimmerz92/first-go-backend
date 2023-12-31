# My First Go Project

## What is it?

It's nothing in particular, really. I was interesting in trying out Go and seeing how it worked.

## Why?

I like trying new things. I also thing using TS/JS on the backend introduces unnecessary complexities in addition to reliance on dependencies for nearly everything. Having used C, I enjoy writing my own functions, and prefer to use dependencies only really when required or it makes the dev experience far more friendly.

## First impressions

I quite enjoyed using Go, and I think I will try to use it more in the future where I have the option. I am yet to combine it with a front end framework - I will probably try out HTMX next.

The learning curve for Go was relatively gentle (although I am still learning). To me it feels like a safer version of C with far less lines required to do the same thing. I also like the fact that I have far more control over memory, as well as it being more appropriate for performance (not that I have ever had to develop for performance.. yet..).

## How to run?

1. `docker build -t psql`
2. `docker run -p 5432:5432 --name psql psql`
   - To stop: `docker stop psql`
   - To remove container: `docker rm psql`
   - To remove image: `docker rmi psql`
3. `go mod tidy`
4. `go mod vendor`
5. `go build && ./go_rss_app`

After this, it should be up and running. There is no front end, you can only currently interact with something like postman.
