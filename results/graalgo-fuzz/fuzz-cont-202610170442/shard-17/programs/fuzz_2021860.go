package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 string }
type GenS2 struct { F0 int; F1 int16; F2 int }
type GenS3 struct { F0 int64; F1 int32; F2 int64 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint, p1 bool, p2 uint64) int16 {
    var v4 int = 32767
    switch v4 {
    case -306:
        for fi5 := 0; fi5 < 16; fi5++ {
            v4 = v4
            fmt.Println(fi5, p0)
            _ = fi5
        }
    default:
        if p1 {
            if p1 {
                for fi6 := 0; fi6 < 14; fi6++ {
                    v7 := GenS1{F0: (v4 | fi6), F1: 255, F2: ""}
                    switch "\n\t" {
                    case "":
                        v8 := []int{}
                        _ = v8
                    default:
                        var v9 uint = (p0 + p0)
                        _ = v9
                    }
                    fmt.Println(v4, fi6)
                    _ = fi6
                    _ = v7
                }
                var v10 int8 = 114
                v11 := GenS2{F0: v4, F1: 32767, F2: v4}
                _ = v10
                _ = v11
            }
        }
    }
    v4 += 42
    if (!true) {
        fmt.Println(p0, v4, p0)
        v12 := []int{(v4 % (((v4 ^ v4)) | 1)), (v4 / ((v4) | 1)), v4, (v4 % ((v4) | 1))}
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := GenS2{F0: v4, F1: -530, F2: 32767}
            for fi15 := 0; fi15 < 13; fi15++ {
                p1 = ((-2.25 / (((3.14159 * 1e10) - 0.0) + (0.1 + 0.1))) != 1.5)
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    if p1 {
        p0 -= 255
        v16 := []int{9223372036854775807, (v4 | (v4 * v4))}
        if true {
            p0 += (p0 / (((p0 - (p0 % ((0) | 1)))) | 1))
            if (0 == 32767) {
                var v17 float32 = (1.5 / 0.1)
                var v18 bool = (100 > -87)
                _ = v17
                _ = v18
            } else {
                p2 -= p2
                var v19 int = (-(v4))
                _ = v19
            }
        }
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -704
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 2
    var v3 float64 = 0.0
    switch 4294967295 {
    case -561:
        var v20 float64 = v3
        _ = v20
    default:
        v21 := GenS3{F0: 9223372036854775807, F1: 2147483647, F2: -1}
        _ = v21
    }
    v22 := GenS4{F0: true, F1: (128 != uint8(v2))}
    v3 -= 0.0
    switch 2147483647 {
    case 2147483647:
        fmt.Println(v2, v3)
    case 358:
        v3 = 1.5
    default:
        switch "foo bar" {
        case "":
            v23 := []int{-9223372036854775808, 761, -1, 0}
            _ = v23
        case "a":
            var v24 float64 = v3
            _ = v24
        default:
            v2 = 1
        }
    }
    v2 = 2
    v25 := []int{}
    _ = len(v25)
    switch "hello" {
    case "foo bar":
        for fi26 := 0; fi26 < 20; fi26++ {
            if len(v25) > 0 {
                v25[0] = (fi26 << ((v2) & 63))
            }
            var v27 float32 = 3.14159
            _ = fi26
            _ = v27
        }
    case "hello":
        switch 995 {
        case 9223372036854775807:
            for fi28 := 0; fi28 < 13; fi28++ {
                delete(v1, "")
                switch "a" {
                case "hello":
                    var v29 int32 = -2147483648
                    _ = v29
                case "":
                    var v30 string = ("foo bar" + "\n\t")
                    _ = v30
                case "\n\t":
                    var v31 float64 = -2.25
                    _ = v31
                default:
                    var v32 string = ""
                    _ = v32
                }
                _ = fi28
            }
        case 42:
            v33 := GenS5{F0: 0, F1: -49}
            _ = v33
        }
    case "":
        var v34 int = 767
        _ = v34
    }
    v2 -= (v2 | (v2 + (v2 + v2)))
    if ((3.14159 / 1.0) == (float32(float64(1.5)) / float32(1e10))) {
        v35 := GenS4{F0: false, F1: (true || ((1 != -9223372036854775808) && false))}
        var v36 uint16 = 0
        v1[(("a" + "a") + "\n\t")] = 0
        _ = v35
        _ = v36
    }
    if (-2.25 == (v3 + -2.25)) {
        v3 = 0.0
        var v37 uint32 = 0
        _ = v37
    } else {
        v38 := GenS6{F0: 226, F1: 2}
        _ = v38
    }
    if false {
        fmt.Println(v2, v3)
    } else {
        _ = len(v25)
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
}
