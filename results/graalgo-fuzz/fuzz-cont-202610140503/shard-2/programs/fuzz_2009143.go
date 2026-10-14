package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int = -1
    var v2 int64 = -9223372036854775808
    var v3 int8 = -9
    var v4 uint8 = 42
    var v5 float32 = (0.0 - (3.14159 * ((1e10 + float32(((float64(1e10) + float64(-2.25)) + ((0.1 - 0.1) - (3.14159 + (1e10 * 1.5)))))) * (((0.0 + -2.25) / 1.5) / (1e10 * -2.25)))))
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 float64 = 0.1
        var v8 uint = 0
        var v9 float64 = (v7 - v7)
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(v5, v1)
    v10 := GenS1{F0: 255, F1: 65535, F2: 0}
    var v11 bool = false
    if (!(float64(1e10) == 3.14159)) {
        var v12 int64 = v2
        _ = v12
    }
    v11 = (("a" <= "") || (!(-1 <= (v1 * (v1 | int(v2))))))
    for fi13 := 0; fi13 < 8; fi13++ {
        v1 -= v1
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := GenS2{F0: 0}
            v16 := []int{}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        fmt.Println(v1, v4)
        if v11 {
            if v11 {
                v5 -= (v5 + ((v5 + float32(((-2.25 / 3.14159) / -2.25))) - (((v5 / (0.1 + v5)) / 1.5) / 1.0)))
            }
            v1 += (fi17 >> ((2) & 63))
            v4 += (v4 << ((0) & 63))
        } else {
            v18 := []int{(v1 * (fi17 & (v1 | -9223372036854775808)))}
            var v19 int16 = -495
            _ = v18
            _ = v19
        }
        v20 := GenS3{F0: -1}
        _ = fi17
        _ = v20
    }
    var v21 float32 = v5
    fmt.Println(v5)
    v4 += (v4 >> ((65535) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
