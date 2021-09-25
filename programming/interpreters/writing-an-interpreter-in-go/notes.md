---
title: "Writing an Interpreter in Go"
author:
  - Thorsten Ball
abstract: |
  These are my notes for ["Writing an Interpreter in Go"](https://interpreterbook.com/). This document consists of a mix of Markdown and \LaTeX; to convert to PDF, use [pandoc](https://pandoc.org/). 
---

## 1

### 1.1

To interpret source code, it needs to be turned into a more accessible form. We'll do it in two steps:
\begin{equation*}
  \text{Source Code} \longrightarrow \text{Tokens} \longrightarrow \text{Abstract Syntax Tree (AST)}
\end{equation*}

Transformation from source code to tokens is called **lexical analysis**. It's done by a **lexer**. Tokens are easily categoralizable data structures that are fed into parser to produce an AST.

### 1.3

To support Unicode and UTF-8, one would need to represent characters using `rune` instead of `byte` in Go.

Function for eliminating whitespace appears in a lot of parsers.

## 2

### 2.6

Operator types:

* Prefix operator, e.g.\ `--` in `--5`
* Postfix operator, e.g.\ `++` in `myVar++`
* Infix operator, e.g.\ `*` in `4 * 3`
