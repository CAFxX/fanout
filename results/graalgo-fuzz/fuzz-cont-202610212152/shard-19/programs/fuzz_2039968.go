package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint32; F1 string }
func genFn0(p0 int64) uint {
    var v5 uint = 65535
    if false {
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := []int{(fi6 % ((-2147483648) | 1)), 127, (fi6 ^ (fi6 / (((fi6 | fi6)) | 1)))}
            _ = fi6
            _ = v7
        }
        fmt.Println(p0, v5)
        v8 := []int{946}
        _ = v8
    } else {
        if true {
            v9 := GenS1{F0: 42, F1: 127}
            _ = v9
        } else {
            var v10 int16 = 406
            v10 -= -1
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 int = (fi11 / ((fi11) | 1))
                fmt.Println(v12, p0, fi11)
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        var v13 float64 = float64((3.14159 + 1e10))
        _ = v13
    }
    _ = p0
    _ = v5
    return (v5 - (v5 | 4294967295))
}

func main() {
    var v1 int = 4294967295
    var v2 int8 = -94
    var v3 float64 = 1.5
    v4 := make(map[uint32]string)
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 18; fi15++ {
            v1 += fi15
            fmt.Println(fi15, fi14, v1)
            fmt.Println(fi14)
            _ = fi15
        }
        _ = fi14
    }
    fmt.Println(v3)
    var v16 uint32 = 42
    v3 = v3
    var v17 float64 = 1e10
    v3 += v3
    var v18 int = v1
    v19 := []int{(v18 >> ((1) & 63)), (^(v1))}
    var v20 int32 = -1
    v2 += (-(v2))
    if (v17 == v17) {
        if true {
            v20 += -2147483648
        }
    }
    v21 := GenS2{F0: (^(v16)), F1: "a"}
    v22 := GenS1{F0: v2, F1: 2}
    var v23 float64 = 3.14159
    if true {
        var v24 bool = ((0 > 0) && false)
        v23 = v23
        for fi25 := 0; fi25 < 16; fi25++ {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            v3 -= (v3 / 0.1)
            _ = fi25
        }
        _ = v24
    }
    var v26 int = 1
    var v27 uint32 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
