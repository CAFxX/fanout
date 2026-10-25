package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint }
func genFn0(p0 uint, p1 int16, p2 uint64) string {
    var v7 uint = p0
    var v8 int32 = -2147483648
    v9 := GenS1{F0: p2, F1: 127, F2: (v7 | 0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return ("foo bar" + "foo bar")
}

func main() {
    var v1 bool = false
    var v2 int32 = 0
    var v3 int8 = 1
    v4 := make(map[string]string)
    v5 := []uint32{936}
    v6 := []int{-1}
    v1 = (!(32767 == 32767))
    v3 -= (v3 & v3)
    var v10 bool = v1
    var v11 int64 = -9223372036854775808
    var v12 uint64 = 42
    v2 = (v2 + v2)
    v10 = false
    switch 32768 {
    case -9223372036854775808:
        if v1 {
            v12 += (v12 - v12)
            v12 = 1
        } else {
            v2 = (v2 / ((1) | 1))
            v3 -= v3
        }
    case 784:
        var v13 uint32 = uint32(v12)
        _ = v13
    case 388:
        v12 = 42
    default:
        if v10 {
            var v14 int64 = 9223372036854775807
            for fi15 := 0; fi15 < 14; fi15++ {
                _ = len(v4)
                fmt.Println(v2, v14, fi15)
                for fi16 := 0; fi16 < 12; fi16++ {
                    v17 := GenS1{F0: (v12 << ((2) & 63)), F1: 0, F2: 0}
                    var v18 bool = (3.14159 <= ((((float32(1e10) - (0.0 / ((float32(float64(0.1)) / 1.0) / 1e10))) / 1.0) + 1.5) - 3.14159))
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi15
            }
            _ = v14
        }
    }
    var v19 uint16 = 648
    v20 := GenS1{F0: (v12 ^ 587), F1: 1, F2: 0}
    var v21 string = ""
    var v22 float32 = (0.0 / 0.1)
    v10 = ((-2.25 + 0.0) >= (-2.25 - 1e10))
    var v23 uint32 = 840
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
