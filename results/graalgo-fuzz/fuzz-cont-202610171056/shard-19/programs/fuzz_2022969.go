package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint64 }
func genFn0(p0 string, p1 uint8, p2 string) uint32 {
    v7 := []int{-2147483648, -9223372036854775808, -193, 42}
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: p2, F1: fi8, F2: 340}
        var v10 int16 = 1
        _ = fi8
        _ = v10
        _ = v9
    }
    switch "a" {
    case "foo bar":
        var v11 float64 = 1.5
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 int8) string {
    fmt.Println(p0)
    v12 := []int{69, 1, 576, -2147483648}
    _ = p0
    _ = v12
    return "foo bar"
}

func main() {
    var v1 int = -2147483648
    var v2 bool = (-1 == 0)
    var v3 bool = (("" + "") < "\n\t")
    var v4 uint8 = 42
    var v5 uint = 774
    var v6 int = (v1 & (v1 * v1))
    fmt.Println(v2, v1, v5)
    v13 := GenS1{F0: (("foo bar" + "") + ("\n\t" + "")), F1: 568, F2: 2}
    v2 = v2
    v1 += (v1 / ((v1) | 1))
    v5 -= (^(v5))
    v14 := []int{-659}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    var v15 int16 = -531
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
