package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }

func main() {
    v1 := []int{-258, 765, -347}
    var v2 uint32 = 42
    var v3 uint = 42
    v4 := GenS1{F0: "a"}
    if true {
        v2 += (^(v2))
    }
    var v5 bool = true
    v5 = ((42 != 42) || (v5 || v5))
    v6 := GenS2{F0: (v3 + 2)}
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 bool = (1 > 496)
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 368
    var v10 string = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
