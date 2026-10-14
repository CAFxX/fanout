package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int64; F2 int64 }
type GenS4 struct { F0 string }
func genFn0() int16 {
    var v7 float32 = (float32(0.1) * 1e10)
    if (32767 != 2147483647) {
        v8 := []int{}
        switch 14 {
        case 9223372036854775807:
            switch 128 {
            case 1:
                var v9 uint32 = 42
                _ = v9
            default:
                v7 -= (v7 + 1.5)
            }
        case -468:
            fmt.Println(v7)
        case 1:
            for fi10 := 0; fi10 < 9; fi10++ {
                if true {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v7 = 1.5
                        var v12 int64 = 905
                        _ = fi11
                        _ = v12
                    }
                    if (false && true) {
                        fmt.Println(fi10, v7)
                        var v13 int8 = 60
                        _ = v13
                    }
                }
                _ = fi10
            }
        default:
            var v14 int8 = 100
            _ = v14
        }
        _ = v8
    } else {
        v7 = float32(0.1)
        v15 := GenS1{F0: 4294967295, F1: 1, F2: 18446744073709551615}
        _ = v15
    }
    _ = v7
    return 42
}
func genFn1(p0 bool, p1 uint16) int64 {
    var v16 int64 = 678
    if p0 {
        if (93 <= -128) {
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 uint = 1
                if true {
                    var v19 int8 = -128
                    if true {
                        v16 = (^(v16))
                        v20 := GenS1{F0: 42, F1: v18, F2: v18}
                        _ = v20
                    }
                    v21 := []int{(-(fi17))}
                    _ = v19
                    _ = v21
                }
                _ = fi17
                _ = v18
            }
        } else {
            var v22 uint16 = p1
            _ = v22
        }
    } else {
        v23 := []int{32768}
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v16
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 int8 = 1
    var v4 uint32 = 4294967295
    var v5 float64 = ((3.14159 / 3.14159) + -2.25)
    var v6 uint32 = 42
    v4 = (v4 ^ v6)
    var v24 int8 = 76
    var v25 uint = uint(v4)
    fmt.Println(v3)
    switch "hello" {
    case "hello":
        if true {
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 string = (("a" + "\n\t") + (("" + "foo bar") + (("foo bar" + "foo bar") + ("\n\t" + "hello"))))
                switch fi26 {
                case -9223372036854775808:
                    var v28 int8 = v24
                    _ = v28
                case 4294967295:
                    v29 := []int{}
                    _ = v29
                case 162:
                    switch 127 {
                    case 399:
                        var v30 int64 = 173
                        _ = v30
                    case 362:
                        var v31 uint64 = 1
                        _ = v31
                    case 32767:
                        fmt.Println(v6)
                    default:
                        v32 := v2[(^(fi26))]
                        _ = v32
                    }
                default:
                    v33 := GenS2{F0: int32(fi26)}
                    _ = v33
                }
                v5 += 1e10
                _ = fi26
                _ = v27
            }
            var v34 uint8 = 1
            _ = v34
        }
    default:
        var v35 int = -9223372036854775808
        _ = v35
    }
    v36 := []int{42}
    v37 := GenS3{F0: true, F1: -1, F2: -9223372036854775808}
    switch 32768 {
    case 127:
        var v38 int8 = (v24 ^ v3)
        _ = v38
    }
    var v39 uint8 = 2
    _ = len(v36)
    v5 = float64(1.5)
    for fi40 := 0; fi40 < 6; fi40++ {
        var v41 uint16 = 255
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        fmt.Println(v39, v3, v3)
        v24 = (v24 >> ((42) & 63))
        var v43 uint8 = (v39 | (v39 & (v39 - (^(v39)))))
        _ = fi42
        _ = v43
    }
    v36 = append(v36, 42)
    v5 += float64(-2.25)
    v44 := GenS4{F0: "\n\t"}
    var v45 int8 = 100
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
