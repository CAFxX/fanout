package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 int16; F1 uint16; F2 int }
func genFn0(p0 uint, p1 uint64) uint16 {
    var v4 int8 = 110
    v5 := GenS1{F0: 0}
    var v6 int64 = 806
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 int32, p1 uint64) bool {
    if ((1.5 - float64((0.1 - 0.1))) < (-2.25 + (1e10 / float64(3.14159)))) {
        fmt.Println(p0, p1)
    } else {
        fmt.Println(p1)
    }
    switch 2147483647 {
    case 42:
        v7 := GenS1{F0: 811}
        _ = v7
    case -184:
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := GenS1{F0: -352}
            var v10 float32 = ((0.0 * 0.0) / (0.0 + float32(float64(((3.14159 / 1.5) / 3.14159)))))
            v11 := GenS2{F0: 0, F1: p1}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    case -933:
        switch 128 {
        case 42:
            p1 = (p1 + p1)
        case 32767:
            var v12 uint16 = 468
            _ = v12
        case 448:
            var v13 uint16 = 65535
            _ = v13
        }
    default:
        var v14 uint32 = 4294967295
        _ = v14
    }
    switch ("" + ((("foo bar" + "hello") + "") + "a")) {
    case "\n\t":
        if (false && (65535 <= 1)) {
            v15 := GenS2{F0: 2147483647, F1: (^(p1))}
            if (!(42 < int64(p0))) {
                var v16 uint8 = 138
                _ = v16
            } else {
                p0 += (^(p0))
                var v17 bool = (-31 >= -1)
                _ = v17
            }
            var v18 int32 = (p0 + (p0 % (((p0 << ((42) & 63))) | 1)))
            _ = v15
            _ = v18
        }
    case "foo bar":
        p1 += 987
    case "a":
        p1 = 18446744073709551615
    }
    v19 := []int{42, 779, 975}
    _ = len(v19)
    _ = p0
    _ = p1
    _ = v19
    return true
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint16 = 65535
    v3 := make(map[string]int)
    v2 = (v2 / ((genFn0(0, uint64(v2))) | 1))
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 uint = 2
        v2 += genFn0((v21 | (v21 << ((v21) & 63))), (v1 ^ v1))
        _ = fi20
        _ = v21
    }
    switch "hello" {
    case "":
        if false {
            var v22 uint8 = 2
            _ = v22
        }
    case "a":
        fmt.Println(v2)
    }
    if (-128 != 0) {
        switch -2147483648 {
        case 32768:
            for fi23 := 0; fi23 < 20; fi23++ {
                fmt.Println(v2, v2, v1)
                _ = fi23
            }
        case 9223372036854775807:
            if true {
                for fi24 := 0; fi24 < 6; fi24++ {
                    v25 := v3[("a" + "")]
                    for fi26 := 0; fi26 < 4; fi26++ {
                        switch 127 {
                        case 2147483647:
                            v2 += (^(v2))
                        case 32768:
                            fmt.Println(v2, v2, v25)
                        case -9223372036854775808:
                            var v27 uint16 = 1
                            _ = v27
                        }
                        if ((!true) && true) {
                            v28 := GenS3{F0: 65535, F1: v1}
                            _ = v28
                        }
                        _ = fi26
                    }
                    _ = fi24
                    _ = v25
                }
            }
        case 32767:
            var v29 int16 = 353
            _ = v29
        }
        if (-1 >= 1) {
            for fi30 := 0; fi30 < 16; fi30++ {
                v31 := GenS2{F0: (fi30 + (fi30 << ((255) & 63))), F1: v1}
                _ = fi30
                _ = v31
            }
        } else {
            if ("hello" >= "a") {
                v32 := GenS4{F0: 709, F1: 255}
                for fi33 := 0; fi33 < 9; fi33++ {
                    for fi34 := 0; fi34 < 11; fi34++ {
                        fmt.Println(fi33, v2)
                        v2 -= v2
                        _ = fi34
                    }
                    fmt.Println(v2)
                    _ = fi33
                }
                _ = v32
            }
            v2 += v2
            v35 := GenS3{F0: 0, F1: (v1 * 2)}
            _ = v35
        }
    } else {
        v2 += 32767
    }
    _ = len(v3)
    v2 -= genFn0(uint(v1), 18446744073709551615)
    fmt.Println(v2, v1)
    fmt.Println(v2)
    var v36 int8 = 1
    v37 := GenS5{F0: -121, F1: 42}
    v38 := GenS6{F0: 1, F1: (v2 ^ 0), F2: -162}
    var v39 int8 = v36
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
