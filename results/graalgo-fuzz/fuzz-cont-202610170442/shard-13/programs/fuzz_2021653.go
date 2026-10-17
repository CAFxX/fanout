package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 bool; F1 int32 }

func main() {
    var v1 float32 = 0.0
    var v2 float64 = float64(v1)
    v3 := []uint32{639, 2147483647, 908}
    v4 := []uint32{329}
    var v5 uint16 = 747
    v6 := GenS1{F0: 489, F1: "foo bar"}
    var v7 int64 = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint = 0
        v10 := GenS2{F0: (true && false), F1: 1}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 bool = true
        _ = fi11
        _ = v12
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    var v13 uint = 65535
    var v14 float32 = v1
    var v15 uint32 = 718
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 uint8 = 127
        if len(v4) > 0 {
            v4[0] = (v15 << ((v13) & 63))
        }
        _ = fi16
        _ = v17
    }
    if (false || (-2.25 < v2)) {
        v13 = 2
        var v18 float32 = 0.1
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v3 = append(v3, (^(v15)))
        switch (fi19 >> ((v13) & 63)) {
        case -403:
            _ = len(v3)
        case -929:
            if ((fi19 << ((uint(v15)) & 63)) <= int(v5)) {
                var v20 int32 = 0
                v1 = v1
                v13 -= uint(v5)
                _ = v20
            } else {
                var v21 uint32 = v15
                v7 = -215
                v1 = float32((float64(1.5) - (v2 + (3.14159 - (v2 / v2)))))
                _ = v21
            }
        case -9223372036854775808:
            switch 42 {
            case 32767:
                var v22 string = "\n\t"
                _ = v22
            case 452:
                v5 = v5
            }
        default:
            v5 += (v5 ^ 117)
        }
        var v23 int64 = int64(v13)
        _ = fi19
        _ = v23
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
