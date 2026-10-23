package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int32; F2 int8 }
type GenS4 struct { F0 int32; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint }
func genFn0() int16 {
    var v4 uint = 0
    for fi5 := 0; fi5 < 2; fi5++ {
        if true {
            v4 = v4
            if (830 == 457) {
                var v6 float32 = 0.0
                fmt.Println(fi5, v6)
                v6 = v6
                _ = v6
            }
            v7 := []int{(fi5 & (fi5 + (fi5 % ((fi5) | 1)))), (fi5 >> ((v4) & 63)), -2147483648, fi5}
            _ = v7
        } else {
            var v8 string = ""
            if false {
                var v9 int = (fi5 / (((fi5 ^ fi5)) | 1))
                v4 = v4
                _ = v9
            } else {
                v10 := []int{fi5}
                _ = v10
            }
            v11 := []int{(fi5 ^ (fi5 | (fi5 & (fi5 & fi5)))), (fi5 - fi5), fi5, 127}
            _ = v11
            _ = v8
        }
        v4 = (v4 / ((65535) | 1))
        _ = fi5
    }
    if (9223372036854775807 <= -9223372036854775808) {
        var v12 uint = v4
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 int32 = -152
            v14 += v14
            v15 := GenS1{F0: v14}
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 uint32 = 73
        _ = v12
        _ = v16
    } else {
        v4 -= (v4 + v4)
    }
    v17 := []int{42, 32767, 128, 9223372036854775807}
    _ = v17
    _ = v4
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 16
    var v3 int64 = 42
    v18 := GenS1{F0: -2147483648}
    v19 := GenS2{F0: 65535}
    switch 0 {
    case 78:
        for fi20 := 0; fi20 < 3; fi20++ {
            v3 += v3
            if false {
                v21 := GenS2{F0: 42}
                _ = v21
            } else {
                fmt.Println(fi20, v2)
            }
            _ = fi20
        }
    case 128:
        if false {
            var v22 uint32 = 4294967295
            _ = v22
        } else {
            v23 := GenS1{F0: 42}
            v1[2147483647] = "hello"
            for fi24 := 0; fi24 < 15; fi24++ {
                v2 -= (v2 & (v2 << ((224) & 63)))
                v25 := []int{(fi24 | (fi24 / ((fi24) | 1))), fi24, fi24, 42}
                _ = fi24
                _ = v25
            }
            _ = v23
        }
    }
    if true {
        v26 := GenS3{F0: 2147483647, F1: 337, F2: 0}
        var v27 float32 = 3.14159
        v2 = 42
        _ = v26
        _ = v27
    } else {
        v3 = (-(v3))
    }
    var v28 uint8 = 128
    for fi29 := 0; fi29 < 6; fi29++ {
        v30 := GenS4{F0: 0, F1: v28, F2: 1}
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 uint64 = 806
            var v33 uint32 = 65535
            fmt.Println(v33, v3)
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi29
        _ = v30
    }
    v34 := GenS5{F0: 65535, F1: 2147483647, F2: 319}
    v35 := GenS1{F0: 1}
    v36 := GenS3{F0: 1, F1: 0, F2: v2}
    v37 := GenS5{F0: 383, F1: 299, F2: 0}
    if (!false) {
        v38 := v1[440]
        _ = v38
    }
    v2 -= v2
    if (true && false) {
        fmt.Println(v28)
    } else {
        v28 = (v28 / ((v28) | 1))
        var v39 int32 = 2147483647
        _ = v39
    }
    v40 := []int{-598}
    switch 784 {
    case 9223372036854775807:
        v3 += int64(v28)
    }
    var v41 float64 = ((-2.25 * 1e10) - (3.14159 - float64(0.1)))
    v41 -= float64(((0.0 - ((1.5 / float32(float64((((1.5 - (-2.25 + 0.0)) - (((((1e10 + 1e10) + (1.5 / 0.1)) - 0.1) + (1e10 + (0.0 * -2.25))) / 0.1)) * 0.1)))) * 3.14159)) * (float32(float64((float32(3.14159) * (0.1 * 3.14159)))) + (float32(((v41 / -2.25) - v41)) * 1e10))))
    v28 = v28
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v40
    _ = v41
}
