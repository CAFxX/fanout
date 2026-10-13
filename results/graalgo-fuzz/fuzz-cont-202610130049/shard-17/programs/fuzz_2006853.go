package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint32; F2 int8 }
func genFn0(p0 bool, p1 uint64, p2 int) int {
    v5 := []int{p2, p2}
    if (0 <= 1) {
        if ((0.0 * (((-2.25 - ((float64((1.5 * 0.1)) - -2.25) * 0.0)) + -2.25) / -2.25)) != 0.1) {
            fmt.Println(p1)
            p2 += (p2 / (((-(p2))) | 1))
            fmt.Println(p0, p1, p2)
        }
        var v6 int = p2
        _ = v6
    }
    fmt.Println(p2)
    v7 := GenS1{F0: 265, F1: -2147483648}
    fmt.Println(p1, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return p2
}
func genFn1() int64 {
    var v8 uint16 = 0
    if false {
        v9 := []int{genFn0((9223372036854775807 >= -370), 1, -805), 128, int(v8)}
        _ = v9
    } else {
        v8 = 65535
        v10 := GenS1{F0: 42, F1: 1}
        _ = v10
    }
    _ = v8
    return 42
}

func main() {
    var v1 int32 = 0
    var v2 string = ("foo bar" + ("a" + "hello"))
    var v3 int64 = -112
    var v4 int64 = (v3 & (v3 % ((v3) | 1)))
    v11 := GenS2{F0: 1}
    v12 := []int{-551, genFn0(false, 42, 0), 42, 467}
    if (true || (!false)) {
        if len(v12) > 0 {
            v12[0] = 32768
        }
        v2 = v2
    }
    if (((((float64(float32(float64((1e10 + -2.25)))) >= float64(1.5)) || true) || (1.5 >= 1.5)) || true) && true) {
        var v13 int16 = -275
        _ = v13
    } else {
        v14 := GenS2{F0: 280}
        _ = v14
    }
    v4 += (v3 / (((v3 / (((v4 + v4)) | 1))) | 1))
    v15 := GenS3{F0: "\n\t", F1: 74, F2: 100}
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 int16 = 1
        var v18 uint16 = 42
        for fi19 := 0; fi19 < 9; fi19++ {
            v3 += genFn1()
            v2 = "\n\t"
            _ = fi19
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    v12 = append(v12, 0)
    fmt.Println(v4)
    if true {
        var v20 bool = (-1 >= -958)
        _ = v20
    }
    v1 = v1
    for fi21 := 0; fi21 < 7; fi21++ {
        if true {
            var v22 uint16 = 42
            _ = v22
        }
        var v23 uint16 = 32767
        _ = fi21
        _ = v23
    }
    var v24 int = genFn0(((1e10 * -2.25) != 0.1), 444, 32768)
    switch v24 {
    case -497:
        var v25 uint32 = 2147483648
        _ = v25
    case 128:
        _ = len(v12)
    default:
        var v26 int16 = -32768
        _ = v26
    }
    var v27 int = (v24 + 128)
    fmt.Println(v27)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
}
