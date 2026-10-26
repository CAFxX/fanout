package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0() string {
    switch "hello" {
    case "foo bar":
        v5 := GenS1{F0: -128, F1: "foo bar", F2: 55}
        _ = v5
    case "":
        if (1 > -1) {
            var v6 string = "hello"
            fmt.Println(v6)
            v7 := []int{}
            _ = v6
            _ = v7
        }
    case "hello":
        var v8 int8 = 96
        _ = v8
    }
    var v9 int64 = 0
    _ = v9
    return ""
}

func main() {
    var v1 string = ("" + "\n\t")
    var v2 uint64 = 0
    var v3 uint8 = 0
    var v4 int = 4294967295
    v4 += v4
    fmt.Println(v3)
    switch 881 {
    case 895:
        switch (v1 + "\n\t") {
        case "\n\t":
            fmt.Println(v4)
        case "foo bar":
            switch v1 {
            case "a":
                v4 -= 2147483647
            }
        }
    case 1:
        v10 := []int{(^(v4))}
        _ = v10
    default:
        v3 -= v3
    }
    v2 += (v2 * (v2 - (v2 | 1)))
    v2 = uint64(v3)
    switch (v4 >> ((384) & 63)) {
    case 2147483647:
        v11 := []int{v4, 612, -2147483648}
        _ = v11
    case 922:
        var v12 int64 = -1
        _ = v12
    default:
        if (42 == 513) {
            v13 := GenS1{F0: -34, F1: "foo bar", F2: v3}
            v14 := []int{v4, v4}
            fmt.Println(v3, v2)
            _ = v13
            _ = v14
        } else {
            var v15 int8 = 42
            _ = v15
        }
    }
    v16 := GenS2{F0: (v2 | v2)}
    var v17 int64 = 1
    var v18 uint64 = 317
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
