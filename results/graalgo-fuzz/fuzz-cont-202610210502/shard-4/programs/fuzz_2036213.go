package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint {
    switch 127 {
    case 4294967295:
        var v7 uint16 = 2
        _ = v7
    case 1:
        var v8 float64 = 1.5
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int = (fi9 & fi9)
        _ = fi9
        _ = v10
    }
    v11 := []int{-700, 193, -402, -2147483648}
    _ = v11
    return 0
}

func main() {
    var v1 int16 = 637
    v2 := make(map[int]int)
    var v3 int = 0
    var v4 uint = 1
    var v5 uint = 2
    var v6 uint = (v4 << ((v4) & 63))
    v12 := []int{127, v3}
    switch "" {
    case "hello":
        var v13 int64 = -776
        _ = v13
    case "foo bar":
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    fmt.Println(v5, v1)
    v6 = 301
    if true {
        v1 += -17
    } else {
        v3 += v3
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := GenS1{F0: 0}
            v2[-2147483648] = fi14
            _ = len(v12)
            _ = fi14
            _ = v15
        }
    }
    v4 -= (^(v4))
    fmt.Println(v4)
    if (369 <= 9223372036854775807) {
        v5 = v4
    } else {
        delete(v2, v3)
        if (!true) {
            v3 = 1
            fmt.Println(v1, v4, v6)
            v16 := GenS1{F0: 42}
            _ = v16
        }
        v1 -= (-(v1))
    }
    var v17 uint16 = 1
    fmt.Println(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
