package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int64 }
type GenS2 struct { F0 bool; F1 int16 }

func main() {
    var v1 int32 = 2147483647
    var v2 float32 = -2.25
    var v3 int32 = (v1 / (((v1 * -829)) | 1))
    var v4 int64 = -1
    var v5 int64 = v4
    v6 := make(map[int]uint32)
    if (65535 != 0) {
        v5 = (v4 >> ((0) & 63))
        v2 -= v2
    } else {
        var v7 float64 = float64((v2 * v2))
        var v8 int64 = (v5 & (v5 << ((65535) & 63)))
        _ = v7
        _ = v8
    }
    var v9 int64 = (v4 / (((-(v5))) | 1))
    v10 := GenS1{F0: 40, F1: 386, F2: (^(v5))}
    v11 := GenS1{F0: 127, F1: (^(v1)), F2: (-(v5))}
    var v12 uint8 = 202
    for fi13 := 0; fi13 < 2; fi13++ {
        if (!false) {
            if (true && true) {
                v5 = 1
            } else {
                _ = len(v6)
                v2 = 0.1
                v9 += -9223372036854775808
            }
            if (!false) {
                fmt.Println(v2, fi13)
            }
        }
        _ = fi13
    }
    v14 := GenS1{F0: (v12 << ((890) & 63)), F1: v1, F2: (v5 % (((v5 & v9)) | 1))}
    fmt.Println(v4)
    if ("hello" <= "a") {
        v2 = (v2 / 1e10)
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 string = ""
            v2 = v2
            _ = fi15
            _ = v16
        }
        switch 32767 {
        case 56:
            fmt.Println(v5, v2)
        case 1:
            var v17 uint = 42
            _ = v17
        }
    }
    v18 := GenS1{F0: 153, F1: -674, F2: v4}
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 bool = false
        fmt.Println(v1)
        v21 := GenS2{F0: false, F1: -1}
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
