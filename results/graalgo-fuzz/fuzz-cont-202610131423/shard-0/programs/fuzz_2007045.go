package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int16 }
func genFn0() bool {
    var v5 uint8 = 2
    v6 := []int{}
    var v7 uint8 = (^(v5))
    if (false || true) {
        fmt.Println(v5, v7)
        v8 := GenS1{F0: 619, F1: 2147483647}
        _ = v8
    } else {
        v9 := GenS2{F0: ("" + ("hello" + ("\n\t" + "a"))), F1: 1}
        _ = v9
    }
    v7 -= (v7 ^ (^(v7)))
    _ = v5
    _ = v6
    _ = v7
    return (100 <= 0)
}
func genFn1(p0 uint64, p1 uint8) int {
    fmt.Println(p1)
    v10 := []int{4294967295, -1, -958, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v10
    return -709
}

func main() {
    v1 := []int{-1, 127, -1, 2147483647}
    v2 := make(map[uint32]bool)
    var v3 bool = true
    var v4 string = ("a" + "hello")
    v4 = "hello"
    switch v4 {
    case "foo bar":
        v3 = true
    }
    switch 127 {
    case 1:
        v4 = v4
    case 4294967295:
        var v11 uint = 65535
        _ = v11
    }
    fmt.Println(v4, v4)
    var v12 float32 = -2.25
    v3 = v3
    var v13 float32 = v12
    switch -600 {
    case 0:
        v14 := GenS3{F0: 32767}
        _ = v14
    default:
        v15 := GenS3{F0: 0}
        _ = v15
    }
    switch -9223372036854775808 {
    case 4294967295:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 667:
        switch v4 {
        case "\n\t":
            fmt.Println(v3, v12, v3)
        case "a":
            v12 = v13
        default:
            var v16 uint16 = 564
            _ = v16
        }
    case -1:
        fmt.Println(v3, v12, v4)
    }
    v17 := GenS3{F0: 1}
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            if (1 >= 0) {
                v13 -= ((v12 / ((v12 - v13) - 3.14159)) - v13)
                var v20 uint64 = 2
                _ = v20
            }
            v3 = (false && false)
            if v3 {
                delete(v2, 422)
                v21 := GenS2{F0: (v4 + ("foo bar" + v4)), F1: 384}
                switch fi19 {
                case 592:
                    v13 += ((v13 - ((v12 + float32(-2.25)) - v12)) + v12)
                case 32767:
                    fmt.Println(v4, v13, v12)
                default:
                    var v22 float64 = float64((1e10 + v13))
                    _ = v22
                }
                _ = v21
            } else {
                v13 += v13
                var v23 float64 = ((((0.1 - 1e10) + float64(((1e10 / v13) / -2.25))) / 1e10) - (-2.25 / -2.25))
                _ = v23
            }
            _ = fi19
        }
        v3 = (1 < -113)
        for fi24 := 0; fi24 < 19; fi24++ {
            fmt.Println(fi18, v3)
            _ = fi24
        }
        _ = fi18
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 bool = ((fi25 - (fi25 << ((4294967295) & 63))) >= fi25)
        fmt.Println(v4, v13, v12)
        _ = fi25
        _ = v26
    }
    v27 := GenS3{F0: 1000}
    var v28 uint = 65535
    fmt.Println(v13, v3, v4)
    if v3 {
        fmt.Println(v28, v3)
        v28 += (v28 / (((v28 >> ((0) & 63))) | 1))
        var v29 uint = (^(v28))
        _ = v29
    } else {
        _ = len(v1)
        v3 = (false && v3)
        for fi30 := 0; fi30 < 12; fi30++ {
            v31 := GenS2{F0: v4, F1: 404}
            _ = fi30
            _ = v31
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
