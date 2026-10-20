package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
func genFn0() int64 {
    var v7 uint8 = 215
    v8 := GenS1{F0: 0}
    var v9 int64 = 100
    if true {
        var v10 int64 = 934
        var v11 bool = (2147483647 < -2147483648)
        _ = v10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]string)
    var v3 int = 990
    var v4 int16 = -1
    var v5 uint16 = 157
    var v6 float64 = (3.14159 - (3.14159 - 3.14159))
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint = 0
        v1 -= uint64(v5)
        _ = fi12
        _ = v13
    }
    v1 -= (v1 ^ v1)
    _ = len(v2)
    var v14 uint32 = 0
    v15 := GenS1{F0: 0}
    for fi16 := 0; fi16 < 17; fi16++ {
        if (1 >= (fi16 >> ((4294967295) & 63))) {
            var v17 uint = 212
            v18 := GenS2{F0: genFn0()}
            _ = v17
            _ = v18
        } else {
            v4 += (v4 >> ((4294967295) & 63))
            switch v3 {
            case 4294967295:
                var v19 bool = false
                _ = v19
            case 188:
                var v20 int = (-(v3))
                _ = v20
            }
        }
        _ = fi16
    }
    var v21 float64 = (0.0 / 3.14159)
    var v22 int64 = 37
    var v23 uint64 = 2
    switch v3 {
    case 1:
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 uint64 = (v1 - (v1 | v1))
            var v26 int = (fi24 % (((fi24 - 9223372036854775807)) | 1))
            _ = fi24
            _ = v25
            _ = v26
        }
    default:
        v14 = uint32(v1)
    }
    var v27 int = v3
    v28 := GenS1{F0: -2147483648}
    v14 = v14
    for fi29 := 0; fi29 < 7; fi29++ {
        v30 := GenS2{F0: (^(v22))}
        _ = fi29
        _ = v30
    }
    v31 := []int{v3, v3, v3}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
