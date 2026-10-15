package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint; F2 int32 }
func genFn0(p0 uint8) uint32 {
    var v7 int16 = 32767
    v8 := GenS1{F0: 1, F1: 1, F2: p0}
    var v9 int64 = -441
    v10 := GenS2{F0: (42 == 18446744073709551615), F1: (p0 ^ (p0 & p0))}
    v11 := []int{10, -2147483648, 32767, -1}
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (-32768 < 1)
    var v3 uint32 = 65535
    var v4 uint16 = 0
    v5 := []int{int(v4), -9223372036854775808, 32767}
    var v6 int64 = 0
    var v12 bool = (2 >= 1)
    var v13 uint32 = 970
    var v14 float64 = 0.0
    var v15 int8 = 0
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 bool = v2
        _ = len(v1)
        var v18 bool = (v6 > v6)
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := v1[v13]
    v4 -= (v4 * (v4 - v4))
    fmt.Println(v19, v2)
    v20 := GenS1{F0: v19, F1: -338, F2: 134}
    var v21 string = ""
    var v22 uint16 = 65535
    fmt.Println(v4, v6)
    fmt.Println(v14, v3)
    v23 := GenS1{F0: (v13 - v19), F1: -148, F2: 255}
    for fi24 := 0; fi24 < 18; fi24++ {
        if v12 {
            v25 := v1[(v13 & (v3 % (((^(v19))) | 1)))]
            _ = v25
        }
        v26 := GenS2{F0: true, F1: 127}
        v27 := GenS3{F0: ((((v14 / (v14 - v14)) - 1e10) * (v14 * float64(0.0))) != v14), F1: 432, F2: -458}
        _ = fi24
        _ = v26
        _ = v27
    }
    v28 := GenS3{F0: v12, F1: 14, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
