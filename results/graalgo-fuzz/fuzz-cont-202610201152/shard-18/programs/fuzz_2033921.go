package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint32 = 0
    var v2 uint = 0
    var v3 int8 = 1
    var v4 uint16 = 0
    v5 := GenS1{F0: false}
    if (false && (true && (true && true))) {
        v6 := GenS2{F0: 128}
        _ = v6
    }
    var v7 float64 = -2.25
    var v8 int64 = 42
    var v9 int32 = -817
    if (!true) {
        var v10 int16 = 42
        if (true || (uint(v1) >= (v2 % ((v2) | 1)))) {
            for fi11 := 0; fi11 < 19; fi11++ {
                v9 = 0
                _ = fi11
            }
        }
        _ = v10
    } else {
        v12 := GenS1{F0: (!false)}
        v7 -= (v7 - (1.5 - 1.5))
        v2 -= v2
        _ = v12
    }
    var v13 uint = (^(v2))
    v1 = uint32(v4)
    for fi14 := 0; fi14 < 15; fi14++ {
        if false {
            v15 := GenS3{F0: v3}
            _ = v15
        } else {
            v4 -= (v4 << ((v13) & 63))
            switch "" {
            case "hello":
                if (v2 == 255) {
                    var v16 int16 = -5
                    _ = v16
                }
            default:
                v1 = v1
            }
            fmt.Println(fi14)
        }
        _ = fi14
    }
    v1 += v1
    var v17 float64 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
