package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint, p1 int8, p2 int64) bool {
    if false {
        if true {
            var v6 bool = false
            _ = v6
        } else {
            if (false || (0 > 560)) {
                switch 1 {
                case 127:
                    p2 = (p2 << (((p0 & (^(p0)))) & 63))
                case 32767:
                    var v7 int32 = 1
                    _ = v7
                default:
                    v8 := []int{-880, 4294967295, 148}
                    _ = v8
                }
                fmt.Println(p0)
                fmt.Println(p1)
            }
            var v9 int = 2147483647
            _ = v9
        }
        v10 := GenS1{F0: ("hello" + "\n\t"), F1: 1, F2: 468}
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(p0)
        var v12 uint64 = 305
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return (!false)
}
func genFn1() uint64 {
    v13 := []int{128, -9223372036854775808}
    var v14 string = "a"
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = -1
    var v3 float32 = 0.0
    v4 := make(map[string]string)
    v5 := make(map[string]bool)
    v15 := []int{-2147483648, 2147483647}
    var v16 int = 619
    if (128 < 103) {
        var v17 int8 = 100
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = v17
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := GenS2{F0: 2}
            var v20 uint16 = 2
            _ = fi18
            _ = v19
            _ = v20
        }
        if ((true || (-9223372036854775808 != v2)) && true) {
            var v21 int16 = 0
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            var v22 uint = 0
            _ = v21
            _ = v22
        } else {
            v4[("" + ("hello" + "\n\t"))] = "\n\t"
        }
        var v23 bool = false
        _ = v23
    }
    v24 := GenS2{F0: 690}
    v25 := GenS2{F0: 0}
    if ((!false) || false) {
        var v26 float64 = 0.1
        v27 := GenS1{F0: "a", F1: (v2 * 9223372036854775807), F2: 91}
        _ = v26
        _ = v27
    } else {
        for fi28 := 0; fi28 < 20; fi28++ {
            v29 := GenS1{F0: "\n\t", F1: (^(v2)), F2: -32768}
            _ = fi28
            _ = v29
        }
    }
    if (((3.14159 - -2.25) / (1.5 / 0.1)) >= (0.1 - 0.1)) {
        var v30 int64 = (v1 - (v1 - v1))
        if ("a" >= "") {
            for fi31 := 0; fi31 < 1; fi31++ {
                v32 := GenS2{F0: 667}
                _ = fi31
                _ = v32
            }
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            v33 := GenS1{F0: "a", F1: v30, F2: -974}
            _ = v33
        }
        _ = v30
    } else {
        if ((true && true) || true) {
            v34 := GenS3{F0: 128}
            _ = v34
        }
    }
    switch v16 {
    case 302:
        v16 = 991
    case 9223372036854775807:
        if true {
            var v35 int32 = 2147483647
            fmt.Println(v16)
            if (false && ((v16 | v16) == (v16 % ((v16) | 1)))) {
                _ = len(v15)
            }
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
