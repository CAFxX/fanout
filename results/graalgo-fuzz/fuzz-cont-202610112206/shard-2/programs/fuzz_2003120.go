package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 string }

func main() {
    var v1 uint32 = 414
    var v2 uint = 2
    var v3 int8 = 42
    var v4 uint8 = 69
    v5 := GenS1{F0: -32768}
    v6 := GenS1{F0: -32768}
    v3 += 11
    fmt.Println(v2, v2, v4)
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 string = "\n\t"
        var v9 string = v8
        _ = fi7
        _ = v8
        _ = v9
    }
    v4 -= (^(v4))
    v1 -= v1
    var v10 float64 = 1e10
    var v11 string = "foo bar"
    v12 := []int{}
    var v13 uint32 = (v1 / (((v1 - (^(v1)))) | 1))
    if false {
        v14 := GenS2{F0: 0, F1: v11}
        _ = v14
    }
    switch 153 {
    case -59:
        var v15 string = "foo bar"
        _ = v15
    default:
        var v16 string = (v11 + v11)
        _ = v16
    }
    var v17 float32 = (0.0 + 3.14159)
    for fi18 := 0; fi18 < 4; fi18++ {
        if (true || ((!(v4 != v4)) && (!(v17 < v17)))) {
            v3 = -128
        } else {
            var v19 uint = v2
            if ((false || true) || ((("a" + ("foo bar" + "a")) + ("foo bar" + v11)) == v11)) {
                var v20 int8 = v3
                fmt.Println(v19, v4, v20)
                _ = v20
            }
            var v21 int8 = (v3 >> ((v2) & 63))
            _ = v19
            _ = v21
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
