package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS8 struct { F0 uint8; F1 int8; F2 int32 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint32; F1 int32; F2 int16 }
func genFn0(p0 uint16) uint32 {
    if (((("a" + "") + "hello") + "a") != "") {
        fmt.Println(p0)
        fmt.Println(p0)
    } else {
        var v6 int = -2147483648
        _ = v6
    }
    p0 = (p0 - (p0 & p0))
    v7 := []int{9223372036854775807, -9223372036854775808, 4294967295, 128}
    if true {
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                if true {
                    if ((2147483647 >= 0) && (2 != 127)) {
                        v10 := GenS1{F0: fi9, F1: -9223372036854775808}
                        v11 := GenS2{F0: 128}
                        p0 -= p0
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        v12 := GenS3{F0: 357, F1: 100}
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        fmt.Println(fi9, fi9)
                        fmt.Println(p0, fi9, p0)
                        v13 := GenS4{F0: 32767}
                        _ = v13
                    }
                }
                _ = fi9
            }
            v14 := GenS5{F0: 18446744073709551615}
            _ = fi8
            _ = v14
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        v7 = append(v7, 59)
    } else {
        for fi15 := 0; fi15 < 16; fi15++ {
            if true {
                v16 := GenS6{F0: p0}
                switch fi15 {
                case 0:
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v18 := GenS7{F0: (p0 & (p0 % ((p0) | 1))), F1: 1, F2: 255}
                        switch ("a" + "\n\t") {
                        case "a":
                            p0 = p0
                        default:
                            v19 := GenS3{F0: 649, F1: 17}
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                case 9223372036854775807:
                    _ = len(v7)
                }
                fmt.Println(p0, p0)
                _ = v16
            }
            fmt.Println(fi15, p0)
            _ = fi15
        }
    }
    _ = p0
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{477}
    v3 := make(map[int]bool)
    var v4 int64 = 42
    var v5 int32 = 42
    v20 := GenS8{F0: 255, F1: 113, F2: (-(v5))}
    v5 += v5
    var v21 int32 = 42
    v22 := GenS9{F0: 0}
    v5 = 2147483647
    fmt.Println(v5, v4)
    for fi23 := 0; fi23 < 10; fi23++ {
        v24 := GenS10{F0: 0, F1: (v5 % ((int32(fi23)) | 1)), F2: 1}
        _ = fi23
        _ = v24
    }
    var v25 float64 = float64(((3.14159 - 0.0) - 1e10))
    _ = len(v2)
    var v26 string = ("foo bar" + "hello")
    var v27 int16 = -1
    var v28 uint16 = 901
    fmt.Println(v27, v26, v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
