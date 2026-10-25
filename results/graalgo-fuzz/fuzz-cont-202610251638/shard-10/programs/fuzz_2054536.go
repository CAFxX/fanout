package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int; F1 string; F2 uint16 }
func genFn0(p0 int32) uint8 {
    if (-545 <= 1) {
        v7 := []int{32768}
        _ = v7
    }
    fmt.Println(p0)
    var v8 int16 = 32767
    switch -837 {
    case 32768:
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v8, v8)
        case "hello":
            v9 := GenS1{F0: (36 <= 1), F1: 65535, F2: 704}
            _ = v9
        }
    case -873:
        p0 = 122
    case 32767:
        v8 -= v8
    default:
        var v10 float32 = -2.25
        _ = v10
    }
    v11 := GenS2{F0: 18446744073709551615, F1: 1}
    _ = p0
    _ = v11
    _ = v8
    return 2
}

func main() {
    var v1 string = ("hello" + "a")
    var v2 uint16 = 0
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    var v5 int32 = -699
    v6 := make(map[int]uint32)
    if (3.14159 >= (-2.25 * (0.0 * (0.0 + float32(0.1))))) {
        v12 := GenS2{F0: 18446744073709551615, F1: -800}
        var v13 float64 = (1e10 + 0.1)
        _ = v12
        _ = v13
    } else {
        v14 := []int{-9223372036854775808, 441}
        _ = v14
    }
    v2 = (^(v2))
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 float64 = 3.14159
        _ = fi15
        _ = v16
    }
    if false {
        var v17 uint16 = (v2 + v2)
        for fi18 := 0; fi18 < 4; fi18++ {
            switch fi18 {
            case 9223372036854775807:
                v1 = "foo bar"
            default:
                var v19 int = (fi18 + 42)
                _ = v19
            }
            if (255 < 255) {
                if false {
                    v17 = 65535
                } else {
                    v20 := GenS3{F0: -977, F1: int8(fi18)}
                    if (1000 < 955) {
                        var v21 bool = (!((-1 <= 42) || (false || ((fi18 / ((fi18) | 1)) <= -107))))
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 18; fi22++ {
                            v17 += v2
                            _ = fi22
                        }
                        v23 := v3[2]
                        _ = v23
                    }
                    v1 = ("foo bar" + (v1 + ("a" + (("foo bar" + "hello") + ""))))
                    _ = v20
                }
                var v24 uint = 0
                _ = v24
            }
            for fi25 := 0; fi25 < 15; fi25++ {
                if true {
                    if (v17 == (^(v17))) {
                        v26 := []int{(-(fi18)), int(v17), (fi25 >> ((42) & 63)), fi25}
                        _ = v26
                    }
                    v1 = "\n\t"
                    if (32767 > 274) {
                        v2 += v17
                    } else {
                        v1 = v1
                        v1 = "foo bar"
                    }
                }
                v2 += v2
                var v27 uint = 40
                var v28 uint64 = 372
                _ = fi25
                _ = v27
                _ = v28
            }
            _ = fi18
        }
        _ = v17
    }
    var v29 uint = 2
    var v30 float64 = (1.5 + (-2.25 + (1e10 / (0.1 / -2.25))))
    v31 := GenS4{F0: int(v5), F1: v1, F2: (v2 / (((v2 | v2)) | 1))}
    switch 98 {
    case 127:
        fmt.Println(v5, v29, v1)
    }
    switch "foo bar" {
    case "":
        v30 = v30
    case "\n\t":
        v1 = v1
    default:
        v1 = v1
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
