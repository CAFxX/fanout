package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int }

func main() {
    var v1 int64 = 0
    v2 := []uint32{2147483647}
    var v3 uint8 = 135
    v4 := []string{"\n\t", ("a" + "\n\t")}
    var v5 uint8 = (v3 | (v3 & 0))
    v6 := GenS1{F0: true, F1: 195}
    for fi7 := 0; fi7 < 12; fi7++ {
        v3 += v3
        var v8 string = ("\n\t" + "")
        _ = fi7
        _ = v8
    }
    if (v5 > v3) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v9 int16 = 0
        _ = v9
    } else {
        var v10 int64 = -716
        _ = v10
    }
    fmt.Println(v3)
    if (511 <= 0) {
        switch 4294967295 {
        case 32768:
            fmt.Println(v5, v1, v1)
        default:
            fmt.Println(v1, v3, v1)
        }
    } else {
        if (127 != -50) {
            var v11 int16 = -32768
            _ = v11
        } else {
            if (v3 < 255) {
                for fi12 := 0; fi12 < 3; fi12++ {
                    v1 += (v1 | v1)
                    _ = fi12
                }
                v3 += 32
                for fi13 := 0; fi13 < 14; fi13++ {
                    v1 -= (v1 << ((42) & 63))
                    _ = fi13
                }
            }
        }
    }
    if false {
        v14 := GenS1{F0: (128 < -2147483648), F1: 42}
        _ = v14
    }
    fmt.Println(v1, v5)
    var v15 int32 = 2147483647
    var v16 string = ""
    v17 := GenS2{F0: 1}
    var v18 int64 = -9223372036854775808
    var v19 float32 = (float32(1.5) / 1e10)
    _ = len(v2)
    var v20 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
