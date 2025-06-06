<h1 align="middle"> <img src="https://skillicons.dev/icons?i=go" width="35"/> GO LANG FUNDAMENTALS <img src="https://skillicons.dev/icons?i=go" width="35" /></h1>

<div align="center">

This repository contains notes and code examples from my GoLang Fundamentals course, structured as a series of blog-style entries.

[![bhaumikmaan - GoLang-Fundamentals](https://img.shields.io/static/v1?label=bhaumikmaan&message=GoLang-Fundamentals&color=00ADD8&logo=github)](https://github.com/bhaumikmaan/Understanding-GoLang "Go to GitHub repo")
&nbsp; [![stars - GoLang-Fundamentals](https://img.shields.io/github/stars/bhaumikmaan/Understanding-GoLang?style=social)](https://github.com/bhaumikmaan/Understanding-GoLang)
&nbsp; [![forks - GoLang-Fundamentals](https://img.shields.io/github/forks/bhaumikmaan/Understanding-GoLang?style=social)](https://github.com/bhaumikmaan/Understanding-GoLang)
&nbsp; [![License](https://img.shields.io/badge/License-Apache--2.0_license-00ADD8)](#license)
<br/>

<a href="https://medium.com/@bhaumikmaan/golang-fundamentals-101-1cc420f7747a" target="_blank">
<img src="https://img.shields.io/badge/Read%20Blog%201%20on-Medium-black?logo=medium&style=for-the-badge" alt="Read Blog 1 on Medium"/>
</a>
&nbsp;
<a href="https://medium.com/@bhaumikmaan/golang-fundamentals-102-beyond-the-basics-7f4161f75daf" target="_blank">
<img src="https://img.shields.io/badge/Read%20Blog%202%20on-Medium-black?logo=medium&style=for-the-badge" alt="Read Blog 2 on Medium"/>
</a>
&nbsp;
<a href="https://medium.com/@bhaumikmaan/golang-advanced-103-deep-dive-into-gos-power-features-2e75788e9e68" target="_blank">
 <img src="https://img.shields.io/badge/Read%20Blog%203%20on-Medium-black?logo=medium&style=for-the-badge" alt="Read Blog 3 on Medium"/>
</a>
&nbsp;
<a href="https://github.com/bhaumikmaan/Understanding-GoLang/stargazers" target="_blank">
 <img src="https://img.shields.io/badge/Star-This%20Repo-yellow?style=for-the-badge&logo=github" alt="Star this repo"/>
</a>
</div>

---

## Building with Go: My Journey Through GoLang Fundamentals 🚀

Go has rapidly become a go-to language for building scalable, high-performance, and concurrent applications. After completing a comprehensive course, I've consolidated my learnings into this series of blog posts. If you're looking to understand why Go is so powerful and how to get started, this is for you. Dive in to explore its elegant simplicity and robust features!
<br>

![img.png](img.png)

---

## Your Path to Go Proficiency 🛠️

This series breaks down Go's core concepts step-by-step, providing a clear roadmap to understanding its capabilities.

All the diagrams from there series are available [here](https://github.com/bhaumikmaan/Understanding-GoLang/tree/main/diagrams). To view, copy and paste into [draw.io](http://drawio.com)


---

### [GoLang Fundamentals 101: Getting Started with Go](https://medium.com/@bhaumikmaan/golang-fundamentals-101-1cc420f7747a)
This post covers the absolute essentials to kickstart your Go journey.
* **Why Go?** Understand Go's philosophy of simplicity, performance, and concurrency.
* **Environment Setup:** Get started with Go installation and essential CLI commands like `go run`, `go build`, and `go fmt`.
* **Packages & Modules:** Learn how Go organizes code into reusable **packages** and manages dependencies with **Go Modules**.
* **Basic Declarations:** Master defining **variables** (with type inference and zero values) and **functions** (with multiple returns).
* **Arrays & Slices:** Explore fixed-length **arrays** and Go's dynamic, flexible **slices**, including how to append and iterate.

Checkout the Implementation here: [Card - Slices & Modules](https://github.com/bhaumikmaan/Understanding-GoLang/tree/main/cards)

---

### [GoLang Fundamentals 102: Beyond the Basics – Types, Structs, and More](https://medium.com/@bhaumikmaan/golang-fundamentals-102-beyond-the-basics-7f4161f75daf)
This part expands your Go toolkit, enabling more sophisticated data handling and development practices.
* **Custom Types & Receivers:** Define your own custom types and attach methods using **receivers**, bringing behavior to your data.
* **Structs:** Learn to create **structs** for grouping related properties and how to embed them for powerful composition.
* **Maps:** Understand **maps** as Go's versatile key-value data structure, perfect for dynamic collections.
* **File Management:** Interact with the file system using the `os` package to **read from and write to files**.
* **Testing in Go:** Discover Go's built-in, lightweight **testing framework** for ensuring code quality and reliability.
  
Checkout the Implementation here: 
* [Structs](https://github.com/bhaumikmaan/Understanding-GoLang/tree/main/structs)
* [Maps](https://github.com/bhaumikmaan/Understanding-GoLang/tree/main/maps)
* [Tests](https://github.com/bhaumikmaan/Understanding-GoLang/blob/main/cards/deck_test.go)


---

### [GoLang Fundamentals 103: Unlocking Go's True Potential](https://medium.com/@bhaumikmaan/golang-advanced-103-deep-dive-into-gos-power-features-2e75788e9e68)
The final installment dives into Go's most powerful and distinctive features, essential for building high-performance, scalable applications.
* **Pointers:** Grasp the concept of **pointers** for direct memory access and understand their role in Go's pass-by-value paradigm, especially with structs.
* **Interfaces:** Explore Go's elegant and **implicit interfaces** for achieving polymorphism and writing flexible, decoupled code.
* **Concurrency (Goroutines & Channels):** Dive into **Goroutines** (lightweight execution units) and **Channels** (for safe communication between them), which form the backbone of Go's powerful concurrency model.
* **HTTP Operations:** Learn how to perform basic HTTP requests and understand the fundamental role of `io.Reader` and `io.Writer` interfaces in handling web data.

Checkout the Implementation here:
* [HTTP Operations & File Management](https://github.com/bhaumikmaan/Understanding-GoLang/tree/main/http)
* [Interfaces](https://github.com/bhaumikmaan/Understanding-GoLang/tree/main/interfaces)
* [Channels & Routines](https://github.com/bhaumikmaan/Understanding-GoLang/tree/main/interfaces)
---

## Resources & Further Study 📚

Ready to dive deeper? Here's a curated list of resources to enhance your Go knowledge:

- **The Go Programming Language Specification**: [Link](https://go.dev/ref/spec)
- **Go Standard Library Documentation**: [Link](https://pkg.go.dev/)
- **Go Build Options**: [Link](https://dev.to/cduggn/go-build-options-570e)
- **Building and Consuming Custom Packages in Go**: [Link](https://medium.com/@ansujain/building-and-consuming-custom-packages-in-go-a-complete-guide-ee6566bf6782)
- **Package Management in Go**: [Link](https://dev.to/cduggn/package-management-in-go-2hjl)
- **Golang Project Structuring — Ben Johnson Way**: [Link](https://medium.com/sellerapp/golang-project-structuring-ben-johnson-way-2a11035f94bc)
- **Golang Deep and Shallow Copy a Slice**: [Link](https://yuminlee2.medium.com/golang-deep-and-shallow-copy-a-slice-7dbcd9eb044a)
- **Internals of Go**: [Link](https://medium.com/breaking-tech/internals-of-go-317e60fae1dc)
- **A Look at Iterators in Go**: [Link](https://medium.com/eureka-engineering/a-look-at-iterators-in-go-f8e86062937c)
- **Leveraging Go's Iterator Pattern**: [Link](https://engineering.teknasyon.com/leveraging-gos-iterator-pattern-30b7d3be783f)
- **Interfaces in Golang: A Deeper Look**: [Link](https://blog.stackademic.com/interfaces-in-golang-a-deeper-look-4a0931481c00)
- **A Straightforward Guide for Go Interface**: [Link](https://dev.to/eyo000000/a-straightforward-guide-for-go-interface-16b2)
- **Go Interfaces for New Gophers**: [Link](https://medium.com/codex/go-interfaces-for-new-gophers-3d608db99d46)
- **Understanding the Power of Go Interfaces: A Comprehensive Guide**: [Link](https://medium.com/@jamal.kaksouri/understanding-the-power-of-go-interfaces-a-comprehensive-guide-835955101b7e)
- **Interfaces in Golang with Examples**: [Link](https://amit-kumar-manjhi.medium.com/interfaces-in-golang-with-examples-ab5c5b7d34b7)