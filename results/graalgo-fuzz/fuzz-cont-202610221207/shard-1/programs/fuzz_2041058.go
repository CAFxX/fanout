package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
func genFn0() string {
    var v5 string = "\n\t"
    v5 = v5
    v6 := GenS1{F0: 42, F1: 69}
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{(fi7 / ((fi7) | 1)), 239}
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return (("a" + (((v5 + (v5 + "foo bar")) + v5) + v5)) + "a")
}

func main() {
    v1 := []string{"a"}
    var v2 uint = 848
    var v3 float64 = 1e10
    var v4 int16 = 1
    var v9 float64 = (v3 - (v3 / (float64(1.5) + 0.1)))
    fmt.Println(v2)
    fmt.Println(v9, v9, v4)
    v9 = (v3 + v3)
    var v10 uint16 = 596
    var v11 int = 506
    switch 32767 {
    case 1:
        var v12 uint8 = 57
        _ = v12
    case 127:
        fmt.Println(v2)
    case 4294967295:
        v1 = append(v1, "a")
    }
    v2 += 540
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
