package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint; F2 string }
func genFn0(p0 uint32) int32 {
    fmt.Println(p0)
    fmt.Println(p0)
    v7 := GenS1{F0: -147}
    p0 = (p0 + p0)
    if false {
        if (!true) {
            var v8 uint8 = 127
            _ = v8
        }
    }
    _ = p0
    _ = v7
    return -2147483648
}
func genFn1(p0 uint, p1 int) uint8 {
    var v9 int32 = genFn0(331)
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := []int{(^(p1)), 1, (fi11 - p1)}
            _ = len(v12)
            _ = fi11
            _ = v12
        }
        var v13 uint = 484
        if (32767 != 208) {
            switch p1 {
            case -1:
                v14 := GenS2{F0: true, F1: 751}
                _ = v14
            case 9223372036854775807:
                v15 := GenS3{F0: (v9 >> (((p0 | p0)) & 63))}
                _ = v15
            case 892:
                v16 := []int{int(v9), 32768, 390, p1}
                _ = v16
            }
        } else {
            fmt.Println(fi10, v13)
        }
        _ = fi10
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v9
    return 64
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 0
    var v3 int = 4294967295
    var v4 int32 = int32(v2)
    var v5 int32 = v4
    var v6 float64 = float64((float32(0.0) - -2.25))
    v4 -= (v5 >> ((1) & 63))
    var v17 int64 = 42
    for fi18 := 0; fi18 < 18; fi18++ {
        fmt.Println(v5, v17, v5)
        _ = fi18
    }
    v19 := v1[651]
    v4 -= (v5 & v5)
    v20 := GenS4{F0: 0, F1: 188, F2: ""}
    if false {
        v17 -= int64(v4)
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint = 65535
            v23 := v1[439]
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    v4 -= (-(v4))
    var v24 int32 = genFn0(65535)
    v3 = 4294967295
    var v25 uint8 = 127
    v26 := []int{v19, v19, (v3 << ((65535) & 63)), (-(v19))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
