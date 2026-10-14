package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 uint64; F1 string; F2 uint64 }
func genFn0() int8 {
    var v6 uint64 = 42
    v6 = 466
    switch 128 {
    case 127:
        v7 := []int{}
        _ = v7
    default:
        fmt.Println(v6)
    }
    v8 := GenS1{F0: 1}
    _ = v6
    _ = v8
    return 12
}

func main() {
    var v1 uint16 = 32767
    var v2 float32 = ((0.1 * (0.1 / 1.5)) * 1e10)
    var v3 uint16 = (v1 ^ 65535)
    var v4 float64 = 0.1
    var v5 uint16 = v1
    var v9 int = 42
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint = 2
        v4 = 0.0
        v12 := []int{0, (^(fi10)), (fi10 ^ (v9 * (v9 | fi10)))}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 bool = true
    v13 = ((int(v5) <= (v9 % ((32767) | 1))) && (v4 <= (((-2.25 - v4) * float64(v2)) / 1.0)))
    var v14 bool = (-2147483648 >= 32768)
    v1 -= (^(v3))
    v15 := GenS1{F0: 857}
    fmt.Println(v3)
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 int8 = -64
        _ = fi16
        _ = v17
    }
    if (float32((-2.25 + 0.1)) >= v2) {
        v18 := []int{(v9 & v9)}
        fmt.Println(v4, v5)
        _ = v18
    } else {
        if v14 {
            v19 := []int{(v9 / (((-(v9))) | 1)), 2147483647, v9, v9}
            _ = v19
        } else {
            v20 := []int{(v9 + v9), (v9 + 9223372036854775807), (v9 << ((2) & 63)), 127}
            for fi21 := 0; fi21 < 11; fi21++ {
                if v13 {
                    v2 = (v2 + 1e10)
                    v22 := GenS1{F0: (v1 >> ((42) & 63))}
                    _ = v22
                } else {
                    v1 -= v1
                }
                _ = fi21
            }
            v3 = 382
            _ = v20
        }
        var v23 uint8 = 128
        if (1 != 0) {
            var v24 int8 = -109
            var v25 int32 = -2147483648
            v25 -= v25
            _ = v24
            _ = v25
        } else {
            v26 := []int{v9, v9, (v9 / (((v9 - -2147483648)) | 1))}
            var v27 int = v9
            _ = v26
            _ = v27
        }
        _ = v23
    }
    var v28 uint16 = (v5 * (v3 - 2))
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 int16 = 0
        var v31 int8 = 127
        v32 := []int{v9}
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    if v13 {
        v33 := GenS2{F0: 4294967295, F1: 222, F2: 2}
        _ = v33
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        v5 += v3
        _ = fi34
    }
    switch "foo bar" {
    case "\n\t":
        var v35 int8 = -128
        _ = v35
    case "":
        v3 = (v1 % ((v1) | 1))
    case "foo bar":
        v36 := GenS3{F0: 42, F1: "\n\t", F2: 979}
        _ = v36
    default:
        v5 = (^(v5))
    }
    var v37 uint16 = (v28 ^ 255)
    v38 := GenS3{F0: 2, F1: "foo bar", F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v9
}
