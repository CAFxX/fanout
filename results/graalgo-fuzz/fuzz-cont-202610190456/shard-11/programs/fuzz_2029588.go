package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int32; F2 uint64 }

func main() {
    var v1 int16 = 42
    v2 := []string{(("a" + "\n\t") + "a"), "", "a"}
    v3 := []int{32767}
    fmt.Println(v1)
    v4 := GenS1{F0: (-2147483648 != 0)}
    _ = len(v3)
    var v5 uint8 = 0
    fmt.Println(v5, v5)
    switch -975 {
    case -16:
        var v6 string = ("foo bar" + "hello")
        _ = v6
    case -9223372036854775808:
        v3 = append(v3, 32767)
    }
    if (false && (372 == 4294967295)) {
        v5 += 2
        v7 := GenS1{F0: (42 == 604)}
        _ = v7
    } else {
        v8 := GenS2{F0: 0, F1: 42, F2: 474}
        v9 := GenS2{F0: 1, F1: 1, F2: 42}
        _ = v8
        _ = v9
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, ("foo bar" + "\n\t"))
    v5 -= v5
    v5 += 2
    var v10 int = 968
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
