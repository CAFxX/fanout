package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint16 }
func genFn0(p0 bool, p1 uint, p2 uint16) int64 {
    p1 = 12
    if p0 {
        switch 42 {
        case 1:
            for fi6 := 0; fi6 < 15; fi6++ {
                p2 = (p2 ^ (p2 ^ 0))
                p1 -= (p1 ^ 0)
                _ = fi6
            }
        case 216:
            fmt.Println(p0, p1)
        case 766:
            p0 = (!(((((((float64(0.0) + 0.0) + (0.1 * (-2.25 - float64((1e10 - 1e10))))) * 0.1) * (-2.25 / (float64((float32(-2.25) * (0.1 * 0.1))) + 1.5))) + (((((1.5 * 1.5) * -2.25) + float64((0.1 * 0.0))) + (1.5 - float64(1e10))) / 1e10)) / 0.1) == (1e10 + 1.5)))
        default:
            p1 -= 18446744073709551615
        }
    }
    v7 := []int{617}
    switch "a" {
    case "foo bar":
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 422}
            var v10 uint8 = 255
            if ((v10 != v10) || p0) {
                if len(v7) > 0 {
                    v7[0] = (fi8 >> ((65535) & 63))
                }
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    case "":
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint16 = 772
            if false {
                p0 = p0
                p1 -= 579
            } else {
                var v13 int16 = -752
                v14 := GenS2{F0: 42, F1: 584, F2: 243}
                _ = v13
                _ = v14
            }
            p0 = (-2.25 != 3.14159)
            _ = fi11
            _ = v12
        }
    default:
        var v15 uint64 = 145
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 305
}
func genFn1() uint64 {
    var v16 int8 = 0
    v17 := GenS3{F0: true, F1: 1}
    switch ("a" + (("" + "\n\t") + (("a" + "hello") + "hello"))) {
    case "\n\t":
        v18 := GenS1{F0: 666}
        _ = v18
    default:
        var v19 uint = 505
        _ = v19
    }
    var v20 int64 = -441
    _ = v16
    _ = v17
    _ = v20
    return 326
}

func main() {
    var v1 int = 821
    var v2 int32 = -425
    var v3 uint32 = 509
    v4 := []uint32{}
    var v5 int64 = -1
    var v21 uint8 = 94
    switch -2147483648 {
    case -1:
        v21 += (v21 >> ((836) & 63))
    case 128:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        if true {
            fmt.Println(v21, v5)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if true {
                if (695 == 18446744073709551615) {
                    v3 = v3
                } else {
                    _ = len(v4)
                    for fi22 := 0; fi22 < 15; fi22++ {
                        var v23 string = "foo bar"
                        _ = fi22
                        _ = v23
                    }
                    var v24 string = (("" + "foo bar") + "hello")
                    _ = v24
                }
                var v25 int8 = 1
                _ = v25
            }
        }
    }
    fmt.Println(v5, v2)
    var v26 int32 = (v2 >> ((0) & 63))
    var v27 uint8 = v21
    v21 = 128
    fmt.Println(v21, v5)
    v28 := GenS3{F0: false, F1: 255}
    v29 := GenS4{F0: 814, F1: 2}
    v2 += 42
    var v30 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
