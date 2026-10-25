package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }
func genFn0(p0 int16, p1 uint16) int {
    p1 = (p1 * p1)
    var v5 uint8 = 2
    fmt.Println(p0)
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint64 = 0
        v8 := []int{(fi6 >> ((0) & 63)), (fi6 + 128)}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint64 = 25
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 0
}

func main() {
    var v1 uint8 = 0
    v2 := []string{"", "hello", "", ""}
    var v3 int = -781
    v4 := []int{}
    if len(v2) > 0 {
        v2[0] = ("" + "hello")
    }
    _ = len(v2)
    var v10 uint = 1
    var v11 int16 = 1
    fmt.Println(v11, v11)
    v1 = v1
    for fi12 := 0; fi12 < 8; fi12++ {
        fmt.Println(v10, v1)
        v13 := GenS1{F0: true, F1: (v1 >> (((^(v10))) & 63)), F2: 18446744073709551615}
        _ = fi12
        _ = v13
    }
    v11 = (v11 + (v11 << ((uint(v1)) & 63)))
    var v14 int32 = -2147483648
    v10 -= v10
    fmt.Println(v11)
    fmt.Println(v14, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
