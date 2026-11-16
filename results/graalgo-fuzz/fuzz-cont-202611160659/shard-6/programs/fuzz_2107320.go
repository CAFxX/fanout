package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int16; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 bool; F2 uint64 }
type GenS5 struct { F0 int32; F1 uint8; F2 uint16 }
func genFn0(p0 int16) uint32 {
    v6 := GenS1{F0: -704}
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS2{F0: 1, F1: (-(p0)), F2: fi7}
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: 386, F1: p0, F2: -1}
    v10 := GenS1{F0: 9223372036854775807}
    if ((p0 % (((-(p0))) | 1)) == -172) {
        v11 := []int{}
        for fi12 := 0; fi12 < 11; fi12++ {
            for fi13 := 0; fi13 < 6; fi13++ {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                p0 = (p0 & (p0 * 42))
                _ = fi13
            }
            _ = fi12
        }
        var v14 uint16 = 109
        _ = v11
        _ = v14
    }
    _ = p0
    _ = v10
    _ = v6
    _ = v9
    return 2147483648
}

func main() {
    var v1 bool = ((-9223372036854775808 == 0) || true)
    var v2 float32 = (0.0 + float32((0.1 - float64((3.14159 / 1e10)))))
    var v3 int64 = -1
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]int)
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int = fi15
        _ = fi15
        _ = v16
    }
    var v17 int64 = (v3 << ((183) & 63))
    v18 := GenS3{F0: -833}
    v19 := []int{4294967295, -185, 1}
    if ((v3 ^ (^(v3))) == 42) {
        fmt.Println(v3, v2, v2)
        v1 = (!false)
        fmt.Println(v3, v3)
    }
    v3 = (v17 ^ 1)
    v20 := GenS4{F0: v1, F1: false, F2: 1}
    fmt.Println(v3, v1, v3)
    v17 += v17
    var v21 string = "hello"
    if false {
        fmt.Println(v17, v3)
    }
    var v22 int64 = (v17 ^ (-(v17)))
    var v23 int16 = 32767
    v24 := GenS2{F0: 2, F1: (v23 - 42), F2: 32768}
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 uint64 = 192
        _ = fi25
        _ = v26
    }
    var v27 int8 = 127
    v28 := GenS5{F0: 42, F1: 151, F2: 42}
    var v29 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
