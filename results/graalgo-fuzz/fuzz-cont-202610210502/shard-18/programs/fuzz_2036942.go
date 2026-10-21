package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int8 }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0(p0 int64, p1 uint32) int16 {
    v4 := []int{-1, -1, -463, 0}
    v5 := GenS1{F0: (p1 / (((^(p1))) | 1)), F1: p0, F2: -1}
    var v6 uint = 911
    var v7 uint16 = 32767
    var v8 int32 = 42
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 float32 = (0.1 * 1e10)
    v2 := make(map[uint32]int)
    var v3 int8 = 0
    v9 := GenS1{F0: 97, F1: 9223372036854775807, F2: 55}
    var v10 uint32 = 2147483648
    for fi11 := 0; fi11 < 12; fi11++ {
        v1 = (v1 * v1)
        _ = fi11
    }
    switch "a" {
    case "foo bar":
        if (!(false || false)) {
            v2[v10] = 128
        }
    case "":
        fmt.Println(v10)
    default:
        v10 -= (v10 ^ (^(v10)))
    }
    v3 = v3
    var v12 uint64 = 0
    if ("hello" < "a") {
        v10 += (v10 % (((v10 - 216)) | 1))
    } else {
        var v13 uint32 = v10
        if (true && false) {
            for fi14 := 0; fi14 < 2; fi14++ {
                for fi15 := 0; fi15 < 10; fi15++ {
                    fmt.Println(v12, v1)
                    fmt.Println(fi14)
                    _ = fi15
                }
                var v16 float64 = 1.5
                _ = fi14
                _ = v16
            }
            var v17 int = 1
            var v18 uint64 = (v12 / (((^(v12))) | 1))
            _ = v17
            _ = v18
        } else {
            var v19 bool = true
            _ = v19
        }
        _ = v13
    }
    var v20 int8 = (v3 * v3)
    if (424 != 9223372036854775807) {
        fmt.Println(v3)
        if (!true) {
            v21 := GenS2{F0: 0, F1: 0}
            _ = v21
        }
    }
    v22 := GenS1{F0: v10, F1: -9223372036854775808, F2: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v9
}
