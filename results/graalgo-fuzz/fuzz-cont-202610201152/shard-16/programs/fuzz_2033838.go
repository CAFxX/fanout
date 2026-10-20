package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 0
    var v3 int8 = 127
    var v4 uint16 = 243
    if false {
        var v5 int16 = 148
        _ = v5
    } else {
        v6 := []int{42, -1}
        _ = v6
    }
    v3 -= 1
    for fi7 := 0; fi7 < 10; fi7++ {
        v2 += (v2 & (^(v2)))
        _ = fi7
    }
    var v8 int16 = -275
    v3 -= -1
    switch 9223372036854775807 {
    case 634:
        v9 := GenS1{F0: (-(v8))}
        _ = v9
    case 767:
        v8 += v8
    }
    if false {
        v2 += v2
    }
    var v10 int = -91
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
