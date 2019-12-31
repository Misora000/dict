`dict` is a simple and light dictionary implemented by trie tree.  
  
```go
// Create a new dictionary instance.
d := dict.NewDictionary()

// Add your words.
str := []string{
	"app",
	"apple",
	"application",
	"apply",
	"apps",
	"appstore",
}

for _, s := range str {
	d.Add(s)
}

// Find string
node, found := d.Find("apple")

// Auto complete string
ac := d.Predict("app")

// Dump the whole dictionary
d.Dump()
```
