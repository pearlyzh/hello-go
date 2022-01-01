# Go
Messing around with Go concepts.

## Characteristics
- Naturally non-blocking
- Compiled directly to machine code
- M:N Scheduling Model
- Is a divergence language

## Goroutine
- https://medium.com/@riteeksrivastava/a-complete-journey-with-goroutines-8472630c7f5c
- https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html
- https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html
- https://rakyll.org/scheduler/

### Compare to Green Thread
- https://softwareengineering.stackexchange.com/questions/222642/are-go-langs-goroutine-pools-just-green-threads


## Pass by Value
- Pass by reference: https://www.ibm.com/docs/en/zos/2.4.0?topic=calls-pass-by-reference-c-only
- Snippets:
```
type Person struct {
	firstName string
	lastName  string
}

// Pass by value of a pointer, like Java
func changeName(p *Person) {
	p.firstName = "Bob"
}

func main() {
	person := Person {
		firstName: "Alice",
		lastName:  "Dow",
	}

	changeName(&person)
	fmt.Println(person) // Output:Bob Dow
}

```

## References
 - https://nickkou.me/2020/06/golang-value-type-and-reference-type/
 - https://go.dev/blog/slices-intro
 - https://medium.com/rungo/the-anatomy-of-slices-in-go-6450e3bb2b94#:~:text=%E2%98%9B%20slice%20is%20a%20struct&text=len%20and%20cap%20is%20the,referenced)%20array%20is%20composed%20of.&text=But%20when%20a%20slice%20references,pointer%20will%20not%20be%20nil%20.
 - https://www.youtube.com/watch?v=a4HcEsJ1hIE
 - https://stackoverflow.com/a/27084475/6085492