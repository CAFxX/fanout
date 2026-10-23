package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int64 }

func main() {
    var v1 float64 = ((1.5 - 3.14159) / 1.0)
    var v2 uint = 255
    var v3 int8 = -10
    var v4 uint16 = 32767
    var v5 int32 = 42
    var v6 int64 = -1
    v7 := []int{98, 4294967295, 4294967295, -702}
    switch 572 {
    case 217:
        var v8 int16 = 171
        _ = v8
    case 32767:
        var v9 int = 35
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: true, F2: v6}
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 bool = (0.0 != 0.0)
        v13 := GenS1{F0: -440, F1: ((v6 | v6) != v6), F2: 149}
        _ = fi11
        _ = v12
        _ = v13
    }
    v3 -= -1
    switch ("hello" + "") {
    case "a":
        var v14 int64 = (^(v6))
        _ = v14
    case "hello":
        if false {
            v6 += v6
        } else {
            fmt.Println(v3, v5)
        }
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v5 = (v5 + (v5 ^ (^(v5))))
        var v16 int64 = v6
        _ = fi15
        _ = v16
    }
    var v17 uint = (^(v2))
    if len(v7) > 0 {
        v7[0] = -1
    }
    v2 = 42
    var v18 int32 = int32(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
