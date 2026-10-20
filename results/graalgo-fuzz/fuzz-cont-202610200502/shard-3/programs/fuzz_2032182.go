package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint; F1 int64; F2 uint64 }
func genFn0(p0 uint) bool {
    var v7 float64 = 1e10
    var v8 uint64 = 4294967295
    _ = p0
    _ = v7
    _ = v8
    return (!false)
}

func main() {
    v1 := []int{-487, 2147483647, 42, -1}
    var v2 int16 = 1000
    var v3 int16 = (v2 & (v2 | -581))
    var v4 int8 = 1
    v5 := []int{275, 127}
    v6 := make(map[string]uint32)
    var v9 bool = false
    v1 = append(v1, int(v2))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v10 := GenS1{F0: -9223372036854775808, F1: 255}
    var v11 int32 = -2147483648
    var v12 float64 = float64(((0.0 + ((-2.25 * (0.0 / ((0.1 + 1e10) - (1e10 * 1e10)))) / 1.0)) - -2.25))
    v13 := GenS2{F0: 653, F1: 0, F2: 170}
    v12 = 1e10
    if (v9 && v9) {
        v14 := GenS2{F0: 0, F1: -1, F2: 1}
        for fi15 := 0; fi15 < 12; fi15++ {
            v11 += (v11 + 42)
            v11 = (^(v11))
            _ = fi15
        }
        v4 = (v4 % (((v4 * v4)) | 1))
        _ = v14
    } else {
        v4 = v4
    }
    var v16 int64 = int64(v2)
    v17 := v6[""]
    var v18 uint32 = (^(v17))
    if (v9 && v9) {
        v4 = (v4 - v4)
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 string = ""
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 bool = v9
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 2; fi23++ {
            v9 = false
            _ = fi23
        }
        _ = fi19
        _ = v20
    }
    var v24 uint8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
