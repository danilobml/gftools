# gftools

**gftools** is a lightweight Go package providing generic, functional-style utilities like `Filter`, `Map`, `Reduce`, `Find`, and predicate factories such as `GreaterThan`, `HasValue`, and `MatchRegex`.

---

## Features

- Functional utilities: `Filter`, `Map`, `Reduce`, `Find`, `Some`, `Every`, `IndexOf`, `FindIndex`, `Reverse`
- Predicate factories: `GreaterThan`, `LessThan`, `Between`, `MatchRegex`, `HasKey`, `HasValue`, `HasKeyValue`, etc.
- Type-safe and generic (Go 1.18+)

---

## Installation

Go 1.18 or later is required.

```sh
go get github.com/danilobml/gftools
```

## Usage Examples:

### Example slice of int:

```go
nums := []int{1, 2, 3, 4, 5}
```

### Filter:

```go
// Using a custom function:
evens := gftools.Filter(nums, func(n int) bool {
	return n % 2 == 0
})

// Or using predicate factory:
evens := gftools.Filter(nums, gftools.IsEven())

// Output: [2 4]
```
### Find

```go
found := gftools.Find(nums, func(n int) bool {
	return n > 3
})
// Output: 4

found2 := gftools.Find(nums, gftools.LessThan(3))
// Output: 1
```

### Map

```go
labels := gftools.Map(nums, func(n int) any {
	return fmt.Sprintf("Item-%d", n)
})
// Output: ["Item-1", "Item-2", "Item-3", "Item-4", "Item-5"]
```

### Reduce

```go
sum := gftools.Reduce(nums, 0, func(acc, curr int) int {
	return acc + curr
})
// Output: 15
```

### Some

```go
hasNegative := gftools.Some([]int{-1, 0, 1}, func(n int) bool {
	return n < 0
})
// Output: true

hasZero := gftools.Some([]int{1, 2, 3}, gftools.Equals(0))
// Output: false
```

### Every

```go
allPositive := gftools.Every(nums, func(n int) bool {
	return n > 0
})
// Output: true

noneGreaterThanTen := gftools.Every(nums, gftools.LessThan(10))
// Output: true
```

### IndexOf

```go
index := gftools.IndexOf([]string{"a", "b", "c"}, "b")
// Output: 1
```

### FindIndex

```go
index := gftools.FindIndex(nums, func(n int) bool {
	return n%2 == 0
})
// Output: 1 (for value 2)
```

### Reverse

```go
reversed := gftools.Reverse([]int{1, 2, 3})
// Output: [3 2 1]
```

### HasKey

```go
maps := []map[string]string{
	{"type": "fruit"},
	{"color": "red"},
}

hasType := gftools.HasKey[string, string]("type")
result := gftools.Find(maps, hasType)
// Output: map[type:fruit]
```

### HasValue

```go
maps := []map[string]string{
	{"animal": "prego"},
	{"animal": "aranha"},
}

hasAranha := gftools.HasValue[string, string]("aranha")
result := gftools.Find(maps, hasAranha)
// Output: map[animal:aranha]
```

### HasKeyValue

```go
maps := []map[string]string{
	{"animal": "prego"},
	{"animal": "aranha"},
}

match := gftools.HasKeyValue[string, string]("animal", "aranha")
result := gftools.Find(maps, match)
// Output: map[animal:aranha]
```

---

## 📖 License

MIT — see [LICENSE](LICENSE)

---

## 🙌 Contributions

Feel free to open issues or PRs to suggest new utilities or improvements!
