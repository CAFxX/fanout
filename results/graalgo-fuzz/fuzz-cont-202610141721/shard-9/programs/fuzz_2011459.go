package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 bool }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8, p1 string, p2 uint) int32 {
    p0 -= p0
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := []string{("hello" + (("" + "") + "hello"))}
    var v2 int16 = 1
    var v3 uint64 = 0
    v4 := []uint32{4294967295}
    var v5 uint64 = (v3 % (((v3 & 950)) | 1))
    v6 := make(map[uint32]bool)
    var v7 int64 = -9223372036854775808
    var v8 float32 = 3.14159
    fmt.Println(v7)
    v3 -= v3
    v2 = v2
    var v9 bool = (127 != 129)
    v10 := GenS1{F0: 2147483648, F1: 128, F2: (!(!(genFn0(143, "foo bar", 65535) < int32(v2))))}
    for fi11 := 0; fi11 < 17; fi11++ {
        if (false || false) {
            for fi12 := 0; fi12 < 7; fi12++ {
                fmt.Println(v3, v8)
                switch fi11 {
                case -516:
                    _ = len(v6)
                }
                _ = fi12
            }
            var v13 int32 = -6
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS1{F0: 65535, F1: 127, F2: v9}
                if (100 >= 114) {
                    var v16 uint32 = 0
                    _ = v16
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        v5 += (^(v3))
                    case "":
                        var v17 int16 = (-(v2))
                        _ = v17
                    default:
                        v18 := GenS1{F0: 867, F1: 127, F2: false}
                        _ = v18
                    }
                    fmt.Println(v7, fi11)
                    v13 -= (v13 << ((2) & 63))
                }
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        v19 := GenS1{F0: 65535, F1: 70, F2: v9}
        switch ("" + "foo bar") {
        case "\n\t":
            for fi20 := 0; fi20 < 14; fi20++ {
                v3 = 1
                switch (fi20 & 1) {
                case 4294967295:
                    if len(v4) > 0 {
                        v4[0] = 2
                    }
                case -14:
                    if (-1 <= -656) {
                        v21 := GenS1{F0: 2147483647, F1: 1, F2: (109 != uint8(v3))}
                        if false {
                            v5 += v5
                        }
                        var v22 int64 = (v7 >> ((249) & 63))
                        _ = v21
                        _ = v22
                    }
                case 760:
                    var v23 uint16 = 2
                    _ = v23
                default:
                    v24 := GenS1{F0: 422, F1: 144, F2: ("a" <= (("hello" + ("foo bar" + ("a" + ("foo bar" + ("a" + ""))))) + "a"))}
                    _ = v24
                }
                v3 -= (^(v5))
                _ = fi20
            }
        case "":
            fmt.Println(fi11, fi11)
        default:
            switch (("" + ("foo bar" + "a")) + "hello") {
            case "a":
                switch 9223372036854775807 {
                case -9223372036854775808:
                    if false {
                        var v25 uint8 = 255
                        _ = v25
                    }
                case 290:
                    var v26 bool = true
                    _ = v26
                case 916:
                    for fi27 := 0; fi27 < 13; fi27++ {
                        fmt.Println(fi11, v5)
                        _ = fi27
                    }
                default:
                    fmt.Println(v3, v8, v7)
                }
            }
        }
        _ = fi11
        _ = v19
    }
    switch 4294967295 {
    case -434:
        var v28 int16 = int16(v7)
        _ = v28
    case -108:
        v6[0] = v9
    default:
        var v29 int64 = 42
        _ = v29
    }
    if v9 {
        var v30 int32 = 0
        _ = v30
    } else {
        v31 := GenS2{F0: true}
        _ = v31
    }
    v3 -= v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
