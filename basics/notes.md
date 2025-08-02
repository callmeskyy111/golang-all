---

### ✅ What is **Tree Shaking**?

**Tree shaking** is a term mostly used in **JavaScript bundlers** (like Webpack, Rollup, or ESBuild). It refers to the **removal of unused code** during the build process — like shaking a tree to get rid of dead leaves. If some imports or functions are never used, the final output bundle **does not include** them.

---

### 🔧 Example in JavaScript:

```js
// math.js
export function add(a, b) { return a + b; }
export function multiply(a, b) { return a * b; }

// app.js
import { add } from './math.js';

console.log(add(2, 3));
```

If `multiply()` is never used, tree shaking removes it from the final bundle.

---

### 🚫 GoLang **does not have traditional tree shaking** like JavaScript bundlers.

However, Go achieves **similar results** at compile time, but in a simpler, more direct way.

---

### ✅ How Go handles unused code:

* If **you import a package but don’t use it**, the Go compiler throws an **error**.
* If **a function, variable, or struct is not used**, it will also be **excluded** from the compiled binary (and possibly raise a warning/error if it's not allowed).

---

### 🔍 Example in Go:

```go
import "fmt" // Error if fmt is unused

func main() {
    // do nothing
}
```

You’ll get a compile-time error:
`imported and not used: "fmt"`

So Go is **strict** about unused imports — it *forces you* to only keep what’s used, effectively achieving tree shaking by **compiler enforcement**, not through bundlers.

---

### ✳️ Use of `_ "some/package"` (Blank Identifier)

Sometimes, Go allows unused imports **intentionally** using `_`:

```go
import _ "github.com/mattn/go-sqlite3"
```

This tells the Go compiler:
“I’m importing this only for its side-effects (e.g., `init()` functions), not direct usage.”

In this case, Go won’t complain, and it won’t tree-shake this out — because **you’re explicitly asking to run side-effects**.

---

### 🧠 Summary:

| Feature              | JavaScript       | GoLang                        |
| -------------------- | ---------------- | ----------------------------- |
| Tree shaking         | Done by bundlers | Done at compile-time by rules |
| Keeps only used code | Yes              | Yes                           |
| Allows unused code   | Sometimes        | No (except with `_`)          |

---
