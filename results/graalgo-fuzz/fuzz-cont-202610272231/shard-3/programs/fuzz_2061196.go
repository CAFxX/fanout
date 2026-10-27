package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]string)
    v3 := []string{}
    var v4 float64 = float64(3.14159)
    var v5 int64 = -598
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "")
    }
    var v6 uint64 = 562
    if (!false) {
        if true {
            var v7 int8 = -128
            var v8 float64 = 3.14159
            _ = v7
            _ = v8
        } else {
            fmt.Println(v1)
            var v9 int64 = (-(v5))
            var v10 uint64 = (v6 - (v6 + v6))
            _ = v10
            _ = v9
        }
    } else {
        v1 = (v1 + v1)
    }
    switch -2147483648 {
    case 2147483647:
        var v11 uint8 = 0
        _ = v11
    case -1:
        for fi12 := 0; fi12 < 20; fi12++ {
            v6 -= v6
            var v13 int = (fi12 & (fi12 << ((65535) & 63)))
            _ = fi12
            _ = v13
        }
    default:
        v6 -= v6
    }
    var v14 uint32 = 454
    v15 := GenS1{F0: (("foo bar" + "\n\t") < "hello"), F1: 476}
    var v16 uint = 0
    v6 = v6
    var v17 uint = v16
    v16 += 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
