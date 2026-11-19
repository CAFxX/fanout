package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int16 }

func main() {
    v1 := []uint32{42, 2147483647, 42, 4294967295}
    var v2 uint8 = 157
    var v3 uint = 65535
    v4 := []int{-2147483648, 127}
    var v5 float32 = (float32(-2.25) - 1e10)
    var v6 uint = v3
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v3 -= 42
            v5 -= 1.5
            _ = fi8
        }
        v9 := GenS1{F0: -350, F1: 781, F2: -32768}
        _ = fi7
        _ = v9
    }
    var v10 bool = ((v3 + v3) < (v6 ^ 255))
    switch 2147483647 {
    case 128:
        switch "" {
        case "\n\t":
            var v11 float32 = (3.14159 - v5)
            _ = v11
        default:
            var v12 string = ""
            _ = v12
        }
    case 42:
        var v13 uint16 = uint16(v2)
        _ = v13
    default:
        var v14 int32 = 2147483647
        _ = v14
    }
    var v15 bool = v10
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        v16 := GenS1{F0: -1, F1: 0, F2: -865}
        _ = v16
    } else {
        v17 := GenS1{F0: -1, F1: 4294967295, F2: 1000}
        _ = v17
    }
    var v18 int8 = 100
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
