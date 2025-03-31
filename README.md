# 🚀 Gin Middleware – Smarter Rate Limiting for Your API!

Welcome to Gin Middleware – a powerful and easy-to-use set of middleware functions for your Gin-based applications! 🎉 Right now, we’re focusing on rate limiting, making sure your API stays fast, fair, and secure.

Built on top of [go-rate-limit](https://github.com/sirius1b/go-rate-limit), this library brings flexible and robust rate limiting strategies to your Gin routes with just a few lines of code.

![Go Version](https://img.shields.io/github/go-mod/go-version/sirius1b/gin-middleware)
![License](https://img.shields.io/github/license/sirius1b/gin-middleware)
![Issues](https://img.shields.io/github/issues/sirius1b/gin-middleware)
![Stars](https://img.shields.io/github/stars/sirius1b/gin-middleware?style=social)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)

## ✨ Why Use This?

With API traffic increasing, rate limiting isn’t just a good idea—it’s a must-have! This middleware makes it super easy to add rate limiting without all the hassle.

<ul>
<li>🔹 Protects your API from abuse 🚫</li>
<li>🔹 Ensures fair usage across clients ⚖️</li>
<li>🔹 Multiple rate-limiting algorithms (Fixed Window, Sliding Window, etc.) ⚙️</li>
<li>🔹 Quick and effortless integration ⏳</li>
</ul>

## Features

<ul>
<li>✅ IP-Based Rate Limiting – Control requests per client IP. </li>
<li>✅ Header Base Rate Limiting - Control requests as per tokens. </li>
<li>✅ Customizable Settings – Set your own limits & windows. </li>
<li>✅ Default Handler – Automatically handles denied requests. </li>
<li>✅ Plug & Play – Integrate seamlessly into any Gin app. </li>
</ul>

## 📦 Installation

Get started with a simple command:

```bash
go get github.com/sirius1b/gin-middleware/pkg
```

## 🔥 Quick Example

Check out the [example file](./examples/main.go) for a full setup.

Here’s how easy it is to add rate limiting to your Gin app:

```go
options := rl.Options{
	Limit:  1, // 1 request
	Window: time.Second, // Per second
}

limiter, err := rl.Require(rl.FixedWindow, options)

router.Use(pkg.ClientIPLimiter(
	&limiter,
	pkg.DefaultRateLimitHandler,
))
```

Boom! 🚀 Now your API is protected & optimized.

## 🤝 Contributions

Want to make this even better? PRs and issues are always welcome! 🎉
Before submitting, make sure your code follows best practices and includes tests.

Ready to take your Gin app to the next level? Give Gin Middleware a try and keep your API running smoothly & securely! 🚀🔥
