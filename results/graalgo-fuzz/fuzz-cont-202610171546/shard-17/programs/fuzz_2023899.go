package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int64; F1 uint8; F2 int8 }
func genFn0(p0 uint32, p1 uint) int {
    p0 = p0
    v5 := []int{int(p0), -1, 1}
    if true {
        switch "foo bar" {
        case "foo bar":
            p0 += (p0 / ((663) | 1))
        default:
            switch -329 {
            case 128:
                p1 = (^(p1))
            case 127:
                v5 = append(v5, -891)
            case 285:
                fmt.Println(p0, p1)
            default:
                p1 -= p1
            }
        }
        fmt.Println(p0, p0)
    } else {
        var v6 uint = (p1 - 1)
        if false {
            v6 += 42
            var v7 float64 = 3.14159
            var v8 uint32 = p0
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: p0, F1: -1, F2: 18446744073709551615}
            _ = v9
        }
        v10 := GenS1{F0: (p0 >> (((v6 + 375)) & 63)), F1: 42, F2: 72}
        _ = v10
        _ = v6
    }
    p1 = (p1 % ((p1) | 1))
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]string)
    var v3 string = (("" + "hello") + ("\n\t" + "hello"))
    v4 := make(map[int]int)
    v11 := GenS2{F0: (1 <= 1), F1: 32768}
    v12 := v4[0]
    v12 = genFn0(2, 320)
    if true {
        v3 = "a"
        var v13 int8 = 1
        _ = v13
    }
    if (0 <= 85) {
        var v14 int16 = 1
        v3 = "hello"
        _ = v14
    }
    v3 = v3
    if (705 >= 1) {
        v15 := GenS2{F0: ((809 >= 42) || (!(!false))), F1: -1}
        fmt.Println(v1, v1, v3)
        fmt.Println(v1, v12)
        _ = v15
    } else {
        if true {
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 int8 = 55
                v17 = v17
                _ = fi16
                _ = v17
            }
            v18 := GenS3{F0: -906, F1: 127, F2: 127}
            _ = v18
        } else {
            v2[v12] = ((v3 + "hello") + (v3 + "\n\t"))
            var v19 int = 0
            _ = v19
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            v12 = 4294967295
            switch "a" {
            case "hello":
                if true {
                    var v21 uint8 = 153
                    _ = v21
                }
            case "":
                fmt.Println(v1)
            default:
                for fi22 := 0; fi22 < 2; fi22++ {
                    fmt.Println(fi20)
                    v3 = "a"
                    _ = fi22
                }
            }
            v3 = "\n\t"
            _ = fi20
        }
        if (v1 < v1) {
            var v23 int8 = -128
            _ = v23
        }
    }
    if false {
        switch v3 {
        case "":
            for fi24 := 0; fi24 < 17; fi24++ {
                for fi25 := 0; fi25 < 2; fi25++ {
                    v26 := []int{128, fi25}
                    var v27 float32 = v1
                    switch 128 {
                    case 128:
                        v27 = (v27 * -2.25)
                    case 4294967295:
                        var v28 int = 128
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                var v29 int64 = -9223372036854775808
                _ = fi24
                _ = v29
            }
        }
        v30 := []int{(v12 & (v12 / ((1) | 1))), v12, (v12 % ((genFn0(2147483647, 597)) | 1))}
        _ = v30
    } else {
        var v31 float64 = ((0.0 / (1.5 - float64((-2.25 - v1)))) / 3.14159)
        v31 -= v31
        var v32 float32 = float32(v31)
        _ = v31
        _ = v32
    }
    var v33 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
