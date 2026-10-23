package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint8; F2 int64 }
func genFn0(p0 int8) int32 {
    if (((float64(float32(1.5)) * (-2.25 * 3.14159)) - (0.1 * float64(3.14159))) >= (1.5 - 0.0)) {
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := []int{(fi6 * 128), (fi6 & (fi6 >> ((18446744073709551615) & 63)))}
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(fi8, fi8)
            v9 := []int{(fi8 ^ (-(fi8)))}
            _ = fi8
            _ = v9
        }
    }
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 42
}

func main() {
    v1 := []int{804, 733, 0}
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    var v4 uint8 = 127
    var v5 float32 = ((1e10 - 1e10) * 1e10)
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 bool = true
            var v13 uint64 = 2
            var v14 float64 = -2.25
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    var v15 string = "\n\t"
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS1{F0: 42, F1: 777}
        v18 := GenS1{F0: 2147483648, F1: 42}
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 int32 = 42
    switch -2147483648 {
    case 240:
        var v20 float64 = (1.5 - (-2.25 - 1.5))
        _ = v20
    case 4294967295:
        v15 = v15
    default:
        switch ((v15 + "") + ((v15 + v15) + v15)) {
        case "hello":
            for fi21 := 0; fi21 < 2; fi21++ {
                _ = len(v1)
                v5 += ((3.14159 - v5) + (v5 * ((1e10 - v5) + 1e10)))
                v19 -= v19
                _ = fi21
            }
        case "foo bar":
            for fi22 := 0; fi22 < 11; fi22++ {
                v5 = 0.1
                _ = fi22
            }
        case "":
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        default:
            var v23 uint8 = (v4 & 255)
            _ = v23
        }
    }
    var v24 int8 = -1
    v25 := GenS2{F0: ("hello" + (v15 + "\n\t")), F1: 0}
    fmt.Println(v24, v15, v19)
    var v26 bool = true
    var v27 uint32 = 4294967295
    fmt.Println(v24)
    if false {
        v2[675] = "a"
        var v28 uint8 = (v4 << ((15) & 63))
        _ = v28
    }
    var v29 int8 = (v24 ^ (v24 & 0))
    var v30 uint8 = v4
    var v31 uint32 = 2147483648
    if (v31 != v31) {
        v4 += v4
    }
    v32 := GenS3{F0: v26, F1: v30, F2: -80}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
