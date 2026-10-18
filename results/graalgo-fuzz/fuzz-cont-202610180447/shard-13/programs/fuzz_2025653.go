package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint }

func main() {
    var v1 float32 = 3.14159
    var v2 uint8 = 127
    var v3 string = "hello"
    v4 := make(map[uint32]uint32)
    var v5 float64 = 1.5
    v6 := GenS1{F0: 634}
    v4[2] = 169
    switch 9223372036854775807 {
    case -9223372036854775808:
        v7 := GenS2{F0: 1000, F1: 1, F2: 2147483647}
        _ = v7
    case 32767:
        v1 -= v1
    default:
        var v8 int8 = 127
        _ = v8
    }
    switch (v3 + ("" + ("\n\t" + v3))) {
    case "foo bar":
        fmt.Println(v1, v1, v2)
    case "\n\t":
        switch int(v2) {
        case -75:
            v1 = v1
        case -824:
            var v9 uint64 = 1
            _ = v9
        case -2147483648:
            var v10 uint = 42
            _ = v10
        default:
            delete(v4, 530)
        }
    case "a":
        var v11 float64 = 3.14159
        _ = v11
    default:
        v5 = v5
    }
    var v12 float32 = 1.5
    switch "a" {
    case "hello":
        for fi13 := 0; fi13 < 6; fi13++ {
            if false {
                var v14 uint32 = 790
                _ = v14
            } else {
                var v15 uint32 = 562
                var v16 int32 = 42
                _ = v15
                _ = v16
            }
            v17 := v4[42]
            var v18 int16 = 1
            _ = fi13
            _ = v17
            _ = v18
        }
    default:
        v19 := GenS1{F0: 1}
        _ = v19
    }
    var v20 uint16 = 65535
    var v21 int32 = 1
    v1 += v1
    v22 := v4[42]
    var v23 uint8 = 255
    v24 := GenS1{F0: (v22 << ((42) & 63))}
    for fi25 := 0; fi25 < 12; fi25++ {
        v26 := GenS3{F0: 1}
        v27 := v4[(v22 ^ 4294967295)]
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 bool = false
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
