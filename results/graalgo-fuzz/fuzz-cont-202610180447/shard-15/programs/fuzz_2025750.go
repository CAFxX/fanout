package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int; F1 uint; F2 uint }

func main() {
    var v1 float32 = 0.0
    var v2 int32 = -688
    var v3 float64 = 1e10
    var v4 int = 42
    v5 := []int{}
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int16 = 652
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch ("hello" + ("foo bar" + (("foo bar" + "hello") + ""))) {
    case "a":
        v9 := GenS1{F0: -94, F1: (v4 + v4)}
        _ = v9
    case "\n\t":
        v10 := GenS2{F0: (v4 % (((-(v4))) | 1)), F1: 18446744073709551615, F2: 18446744073709551615}
        _ = v10
    }
    switch "" {
    case "\n\t":
        fmt.Println(v4, v3, v4)
    }
    if len(v5) > 0 {
        v5[0] = (v4 >> ((1) & 63))
    }
    fmt.Println(v4, v4)
    var v11 int16 = 1000
    var v12 string = "foo bar"
    v13 := GenS2{F0: v4, F1: 255, F2: 2}
    switch ("foo bar" + (((v12 + v12) + v12) + v12)) {
    case "\n\t":
        var v14 bool = true
        _ = v14
    case "foo bar":
        fmt.Println(v11)
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 int = v4
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
