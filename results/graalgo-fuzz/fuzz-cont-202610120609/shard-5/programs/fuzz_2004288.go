package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]int)
    var v3 int = -138
    var v4 uint16 = 426
    var v5 bool = (v1 && v1)
    var v6 int8 = -1
    var v7 uint64 = 1
    v8 := v2[42]
    v3 += v8
    switch (("\n\t" + "\n\t") + ("foo bar" + ("" + "a"))) {
    case "\n\t":
        switch int(v6) {
        case 4294967295:
            var v9 string = ("foo bar" + "foo bar")
            _ = v9
        case -151:
            v7 += v7
        case -2147483648:
            delete(v2, 2)
        }
    case "a":
        switch v3 {
        case -1:
            v10 := []int{}
            _ = v10
        default:
            fmt.Println(v7, v8, v5)
        }
    }
    v11 := GenS1{F0: v3, F1: v5, F2: 15}
    var v12 float64 = 3.14159
    var v13 float64 = -2.25
    v14 := GenS1{F0: (v3 - -9223372036854775808), F1: v5, F2: 13}
    v5 = (!false)
    v8 += -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
