package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }

func main() {
    var v1 bool = false
    var v2 int32 = -434
    var v3 bool = true
    v4 := []int{4294967295, 32768, 42}
    var v5 int8 = 85
    var v6 uint32 = 503
    if false {
        if len(v4) > 0 {
            v4[0] = 42
        }
        var v7 int64 = 0
        _ = v7
    } else {
        _ = len(v4)
    }
    v2 -= int32(v5)
    if v3 {
        if (false && v1) {
            var v8 uint16 = 0
            v8 = (v8 ^ v8)
            _ = v8
        }
        v1 = v3
        v9 := GenS1{F0: "\n\t", F1: (v2 / ((v2) | 1))}
        _ = v9
    }
    var v10 uint8 = 127
    var v11 uint64 = 844
    fmt.Println(v1, v3)
    var v12 uint8 = v10
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 string = "hello"
        for fi15 := 0; fi15 < 14; fi15++ {
            if len(v4) > 0 {
                v4[0] = (fi15 / ((fi15) | 1))
            }
            var v16 uint8 = v12
            if (v1 && true) {
                var v17 bool = true
                fmt.Println(v2)
                v2 -= 320
                _ = v17
            }
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    switch (("" + ("" + "hello")) + "") {
    case "a":
        var v18 int32 = v2
        _ = v18
    default:
        var v19 uint8 = v10
        _ = v19
    }
    v12 = v10
    if v3 {
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
