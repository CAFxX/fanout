package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int64; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 bool; F2 int }
func genFn0(p0 uint16, p1 uint16, p2 int16) string {
    p0 -= (p1 & 2)
    v6 := []int{300, 128, 0, 4294967295}
    if true {
        v7 := GenS1{F0: 65535, F1: 42, F2: -1}
        if true {
            p2 = p2
            fmt.Println(p0, p2, p0)
        } else {
            p2 = (-(p2))
            p1 -= 255
            v8 := GenS2{F0: -32768}
            _ = v8
        }
        _ = v7
    }
    if false {
        v9 := GenS3{F0: 1, F1: -9223372036854775808, F2: p2}
        var v10 uint8 = 42
        p2 = (p2 % (((p2 << ((255) & 63))) | 1))
        _ = v10
        _ = v9
    }
    var v11 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    return ""
}

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]uint32)
    var v3 float32 = 0.0
    v4 := make(map[int]string)
    var v5 uint = 727
    v3 -= v3
    v12 := GenS2{F0: v1}
    var v13 int16 = (v1 % ((1) | 1))
    var v14 int32 = -154
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS1{F0: 24, F1: 127, F2: -832}
        _ = fi15
        _ = v16
    }
    v17 := []int{2147483647}
    var v18 int8 = 127
    v1 = v13
    fmt.Println(v13, v18)
    v19 := GenS4{F0: 0}
    var v20 uint64 = 2
    var v21 int16 = (v1 * v13)
    fmt.Println(v3, v14)
    if (0.1 != v3) {
        v3 = 1.5
    } else {
        switch "a" {
        case "hello":
            switch 42 {
            case 830:
                var v22 uint64 = v20
                _ = v22
            default:
                v18 = 42
            }
        case "":
            if (!(v1 >= v21)) {
                fmt.Println(v18, v3, v5)
                if (true || (-1 == 9223372036854775807)) {
                    v3 -= float32(3.14159)
                    v3 -= ((v3 + -2.25) / float32(float64(float32((3.14159 - 0.0)))))
                }
            } else {
                if (true || true) {
                    for fi23 := 0; fi23 < 8; fi23++ {
                        for fi24 := 0; fi24 < 13; fi24++ {
                            var v25 uint = (v5 << (((v5 >> (((v5 << ((1) & 63))) & 63))) & 63))
                            _ = fi24
                            _ = v25
                        }
                        _ = fi23
                    }
                }
                if (1 > (^(v18))) {
                    for fi26 := 0; fi26 < 19; fi26++ {
                        fmt.Println(v3)
                        _ = fi26
                    }
                    v5 = v5
                }
                switch "" {
                case "hello":
                    v27 := GenS5{F0: v18, F1: (0 >= 2147483647), F2: -2147483648}
                    _ = v27
                case "":
                    fmt.Println(v18, v21)
                case "foo bar":
                    if false {
                        v28 := GenS3{F0: v5, F1: -9223372036854775808, F2: 42}
                        _ = v28
                    } else {
                        v29 := v2[uint32(v5)]
                        _ = v29
                    }
                default:
                    v3 = v3
                }
            }
        case "a":
            fmt.Println(v5, v18, v3)
        default:
            switch -922 {
            case 9223372036854775807:
                for fi30 := 0; fi30 < 15; fi30++ {
                    var v31 uint16 = 32767
                    _ = fi30
                    _ = v31
                }
            }
        }
    }
    if true {
        v32 := GenS1{F0: 65535, F1: 42, F2: 250}
        _ = v32
    }
    if true {
        if ((v5 << ((18446744073709551615) & 63)) >= 718) {
            var v33 uint = 2
            var v34 uint16 = 42
            _ = v33
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
