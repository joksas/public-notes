---
title: "Information Theory"
---

These are my notes for information theory based mostly on the textbook "Elements of Information Theory" by Thomas M. Cover and Joy A. Thomas.

# Chapter 2

## 2.1

**Entropy** of a discrete random variable $X$ is
\begin{equation}
  H(x) \equiv - \sum_{x \in \mathcal{X}} p(x) \log p(x)
\end{equation}

**Expectation** of a random variable $X$ is
\begin{equation}
  E_p g(X) \equiv \sum_{x \in \mathcal{X}} g(x) p(x)
\end{equation}

## 2.2

**Joint entropy** of a pair of discrete random variables $(X, Y)$ with a joint distribution $p(x, y)$ is
\begin{equation}
  H(X, Y) \equiv - \sum_{x \in \mathcal{X}} \sum_{y \in \mathcal{Y}} p(x, y) \log p(x, y) = -E \log p(X, Y)
\end{equation}

**Conditional entropy** is
\begin{align}
  H(Y | X) &\equiv \sum_{x \in \mathcal{X}} p(x) H(Y | X = x) \\
	   &= -\sum_{x \in \mathcal{X}} p(x) \sum_{y \in \mathcal{Y}} p(y|x) \log p(y|x) \\
	   &= -\sum_{x \in \mathcal{X}} \sum_{y \in \mathcal{Y}} p(x, y) \log p(y|x)
\end{align}

**Chain rule** involving conditional entropy is
\begin{equation}
  H(X, Y) = H(X) + H(Y|X)
\end{equation}

## 2.3

**Relative entropy** or **Kullback-Leibler distance** is
\begin{align}
  D(p || q) &\equiv \sum_{x \in \mathcal{X}} p(x) \log \frac{p(x)}{q(x)} \\
	    &= E_p \log \frac{p(X)}{q(X)}
\end{align}

* Relative entropy is a measure of the distance between two distributions. If we knew distribution $p$, we could construct code with average description length $H(p)$.
* If we used code for distribution $q$, the average description length for $p$ would be $H(p) + D(p || q)$.

Relative entropy is not a true distance because it is not symmetric and does not satisfy triangle inequality. Nevertheless, it is often useful to think of relative entropy as "distance".

**Mutual information** is the relative entropy between the joint distribution and the product distribution:
\begin{equation}
  I(X; Y) \equiv D \left( p(x, y) || p(x) p(y) \right)
\end{equation}

## 2.4

We can show that
\begin{equation}
  I(X; Y) = H(X) - H(X|Y) = H(Y) - H(Y|X)
\end{equation}

Mutual information is the reduction in uncertainty of one variable due to the knowledge of the other variable.

Also
\begin{equation}
  I(X; Y) = H(X) + H(Y) - H(X, Y)
\end{equation}
and
\begin{equation}
  I(X; X) = H(X)
\end{equation}
