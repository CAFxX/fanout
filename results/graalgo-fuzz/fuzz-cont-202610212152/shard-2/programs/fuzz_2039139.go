package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int32 }
type GenS2 struct { F0 int64; F1 bool; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 uint; F2 uint16 }
func genFn0() int {
    var v5 int32 = 42
    for fi6 := 0; fi6 < 10; fi6++ {
        if (2147483647 >= 4294967295) {
            switch "foo bar" {
            case "foo bar":
                for fi7 := 0; fi7 < 4; fi7++ {
                    if true {
                        switch "\n\t" {
                        case "":
                            v8 := []int{(fi6 & 127), (fi6 ^ fi6), fi7, (fi6 % (((fi6 - 32768)) | 1))}
                            _ = v8
                        default:
                            var v9 uint64 = 183
                            _ = v9
                        }
                        v10 := GenS1{F0: 42, F1: 377, F2: v5}
                        _ = v10
                    } else {
                        v11 := []int{}
                        _ = v11
                    }
                    v5 -= 42
                    _ = fi7
                }
            }
            v5 += 458
        }
        var v12 uint32 = 32
        _ = fi6
        _ = v12
    }
    v13 := GenS2{F0: -319, F1: true, F2: 42}
    v14 := []int{32768, -2147483648}
    _ = v13
    _ = v14
    _ = v5
    return -412
}
func genFn1() string {
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int16 = 367
        for fi17 := 0; fi17 < 15; fi17++ {
            v16 += v16
            v18 := GenS3{F0: 28}
            v19 := []int{(fi15 / ((fi17) | 1)), 593, 121, 1}
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 16; fi20++ {
            v16 += v16
            switch ("a" + "a") {
            case "":
                v16 = v16
            default:
                switch (-(fi20)) {
                case 127:
                    fmt.Println(fi15)
                }
            }
            v21 := GenS4{F0: 1, F1: "\n\t"}
            _ = fi20
            _ = v21
        }
        _ = fi15
        _ = v16
    }
    var v22 string = (("a" + ("hello" + "foo bar")) + "hello")
    _ = v22
    return ""
}

func main() {
    var v1 int32 = 1
    v2 := []uint32{2147483647, 1, 668, 65535}
    var v3 float64 = 0.0
    var v4 uint8 = 0
    v3 = 0.0
    var v23 uint16 = 42
    for fi24 := 0; fi24 < 19; fi24++ {
        switch "hello" {
        case "":
            switch (fi24 + (fi24 ^ (^(fi24)))) {
            case 4294967295:
                var v25 int = (fi24 % ((int(v1)) | 1))
                _ = v25
            case 393:
                fmt.Println(v23, v3)
            default:
                fmt.Println(fi24, v3, v4)
            }
        case "hello":
            var v26 string = "foo bar"
            _ = v26
        }
        if (int64(v23) < 0) {
            fmt.Println(v23)
            v23 -= v23
        } else {
            v1 -= 2147483647
            for fi27 := 0; fi27 < 12; fi27++ {
                v28 := GenS1{F0: (v1 / ((42) | 1)), F1: v23, F2: -2147483648}
                if (!true) {
                    _ = len(v2)
                    switch int(v23) {
                    case -478:
                        v29 := GenS3{F0: (^(v4))}
                        v4 += 255
                        v30 := GenS2{F0: -1, F1: false, F2: 18446744073709551615}
                        var v31 uint16 = v23
                        v32 := GenS1{F0: (v1 + -1), F1: v23, F2: (v1 % (((v1 >> ((436) & 63))) | 1))}
                        _ = len(v2)
                        _ = v29
                        _ = v30
                        _ = v31
                        _ = v32
                    case -2147483648:
                        v4 = (v4 / ((v4) | 1))
                    case 32767:
                        var v33 int8 = -96
                        _ = v33
                    default:
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                    }
                } else {
                    v3 -= (float64((((3.14159 - (float32((v3 * 0.0)) - float32(1e10))) + 1.5) / 1.0)) / (v3 * (v3 / v3)))
                    if len(v2) > 0 {
                        v2[0] = 2147483647
                    }
                    v23 = (v23 << ((255) & 63))
                }
                _ = fi27
                _ = v28
            }
        }
        v4 -= 168
        _ = fi24
    }
    v3 = (1e10 - 0.0)
    if true {
        var v34 float32 = 0.1
        v35 := GenS2{F0: -852, F1: true, F2: 541}
        v36 := GenS5{F0: -923}
        _ = v34
        _ = v35
        _ = v36
    }
    fmt.Println(v23, v4)
    v23 -= 0
    v1 += v1
    v37 := GenS1{F0: v1, F1: (^(v23)), F2: (v1 | (v1 / ((42) | 1)))}
    fmt.Println(v3)
    if (1 > 0) {
        var v38 uint32 = 65535
        _ = v38
    }
    var v39 uint = uint(v4)
    fmt.Println(v39, v4)
    v40 := GenS6{F0: 416, F1: (v39 + 18446744073709551615), F2: (v23 % ((v23) | 1))}
    if true {
        _ = len(v2)
        fmt.Println(v23, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v40
}
