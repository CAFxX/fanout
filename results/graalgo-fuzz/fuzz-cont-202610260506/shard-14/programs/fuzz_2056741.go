package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }

func main() {
    var v1 string = ""
    var v2 string = v1
    var v3 uint8 = 8
    var v4 string = "\n\t"
    var v5 int64 = 1
    var v6 bool = true
    switch ((v2 + "") + (((v4 + v4) + "hello") + "a")) {
    case "":
        v7 := []int{32768, -2147483648}
        _ = v7
    case "\n\t":
        v8 := []int{2147483647, -9223372036854775808, 9223372036854775807}
        _ = v8
    case "a":
        if (100 < 100) {
            v3 += v3
            v1 = (((("foo bar" + v2) + "foo bar") + v4) + (((v4 + ("foo bar" + v4)) + v4) + v1))
        } else {
            switch 1 {
            case 120:
                fmt.Println(v3, v1)
            case -486:
                var v9 uint16 = 32767
                _ = v9
            }
            v2 = v2
        }
    default:
        if (65535 <= 954) {
            fmt.Println(v2)
            v1 = "hello"
        }
    }
    v2 = v2
    v6 = (!(1 != 32767))
    v10 := []int{}
    v4 = (((("" + v1) + ((v2 + v4) + v2)) + (v4 + v4)) + v1)
    v6 = v6
    var v11 int8 = 0
    var v12 int = int(v5)
    v13 := GenS1{F0: v4, F1: v11}
    v14 := GenS1{F0: "\n\t", F1: (^(v11))}
    var v15 uint = 257
    for fi16 := 0; fi16 < 5; fi16++ {
        fmt.Println(fi16, v15, v3)
        _ = fi16
    }
    fmt.Println(v11, v5, v15)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
