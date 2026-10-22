package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint8, p1 string) int64 {
    var v5 int8 = -19
    v6 := GenS1{F0: (!((true || false) || true)), F1: uint64(p0)}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 702
}

func main() {
    var v1 int32 = 1
    var v2 string = ""
    var v3 int16 = 42
    v4 := make(map[string]uint32)
    v4[v2] = 2
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int16 = (v3 / ((-781) | 1))
        fmt.Println(fi7, v3)
        v9 := GenS2{F0: 2147483648, F1: 429, F2: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = 9223372036854775807
    var v11 uint64 = 864
    fmt.Println(v10)
    v11 += 42
    v10 += -1
    var v12 int8 = 1
    v13 := []int{-2147483648, 32768}
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 uint32 = 2147483648
            v3 += v3
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v17 := GenS3{F0: 1}
    _ = len(v13)
    var v18 int16 = 32767
    var v19 bool = (!(0.0 > 0.1))
    for fi20 := 0; fi20 < 4; fi20++ {
        _ = len(v4)
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
