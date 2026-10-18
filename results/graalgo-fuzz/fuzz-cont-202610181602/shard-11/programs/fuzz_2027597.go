package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int16 }
type GenS2 struct { F0 int16; F1 int }

func main() {
    var v1 int32 = -761
    var v2 int32 = (-(v1))
    var v3 int64 = 0
    var v4 uint = 0
    var v5 uint64 = 1
    v6 := []uint32{42, 566, 821, 2147483647}
    _ = len(v6)
    v2 = v2
    v2 -= 42
    v7 := GenS1{F0: (v1 / ((v2) | 1)), F1: (!true), F2: 0}
    v2 -= (v1 * (v1 - 0))
    if true {
        v8 := GenS1{F0: -351, F1: (-1 < (-(v1))), F2: 0}
        _ = len(v6)
        var v9 float64 = (1.5 / 0.1)
        _ = v8
        _ = v9
    } else {
        fmt.Println(v2)
        v5 += v5
        v4 = v4
    }
    if true {
        fmt.Println(v3, v2, v2)
    }
    if (false || true) {
        v5 = uint64(v4)
        var v10 string = ((("hello" + "\n\t") + "hello") + ("" + "foo bar"))
        v11 := GenS2{F0: -1, F1: -540}
        _ = v10
        _ = v11
    } else {
        fmt.Println(v2, v3, v5)
        var v12 int64 = (v3 + (v3 >> ((v4) & 63)))
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 int32 = 2147483647
            fmt.Println(v3, v4)
            v12 = 0
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    var v15 uint16 = 120
    if false {
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 float64 = 0.1
            if (fi16 != (fi16 - 2147483647)) {
                v15 = (v15 << ((v4) & 63))
                var v18 float64 = (0.0 / v17)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v18
            }
            _ = fi16
            _ = v17
        }
        v19 := GenS1{F0: v2, F1: true, F2: 0}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v19
    }
    v2 = 2147483647
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
