package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int }
func genFn0(p0 uint8) uint16 {
    v6 := []int{-2147483648}
    fmt.Println(p0)
    var v7 float32 = (0.1 / 1e10)
    _ = p0
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 int8, p1 uint16, p2 int64) bool {
    var v8 int16 = 213
    var v9 uint16 = 502
    switch "hello" {
    case "a":
        var v10 uint16 = v9
        _ = v10
    case "":
        if (!true) {
            switch ("a" + "foo bar") {
            case "a":
                for fi11 := 0; fi11 < 17; fi11++ {
                    v12 := GenS1{F0: v8, F1: ("" + "\n\t")}
                    _ = fi11
                    _ = v12
                }
            case "":
                v13 := GenS1{F0: v8, F1: "foo bar"}
                _ = v13
            }
            var v14 int16 = v8
            fmt.Println(p2)
            _ = v14
        }
    case "\n\t":
        switch int(v8) {
        case 0:
            v15 := GenS1{F0: 805, F1: "\n\t"}
            _ = v15
        case -1:
            switch 128 {
            case 570:
                for fi16 := 0; fi16 < 10; fi16++ {
                    var v17 uint64 = 42
                    _ = fi16
                    _ = v17
                }
            case 1:
                v18 := GenS1{F0: v8, F1: "a"}
                _ = v18
            default:
                v9 = (p1 >> ((634) & 63))
            }
        case 182:
            var v19 uint16 = (^(v9))
            _ = v19
        }
    default:
        switch ("hello" + "hello") {
        case "foo bar":
            v20 := GenS1{F0: -744, F1: ("" + ("hello" + "a"))}
            _ = v20
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint32 = (v1 - (v1 % ((v1) | 1)))
    v3 := []uint32{(v2 % (((v2 * (v1 ^ v2))) | 1)), (v1 >> ((18446744073709551615) & 63)), v1, v1}
    v4 := []uint32{(v2 >> ((2) & 63)), 65535, (v1 | (v1 | v2)), v1}
    v5 := make(map[int]bool)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v21 string = "a"
    v21 = ("" + v21)
    if (((-2.25 + float64(1e10)) + ((1e10 + (0.0 / 1.0)) + (3.14159 - 1e10))) >= (0.1 * 1.5)) {
        v5[1] = true
        fmt.Println(v1, v1)
    }
    if ((711 != 658) || (!false)) {
        for fi22 := 0; fi22 < 1; fi22++ {
            v1 -= v1
            _ = fi22
        }
        if (!(true || (!true))) {
            v23 := GenS1{F0: -738, F1: (v21 + ("hello" + (v21 + "\n\t")))}
            v1 = (v1 | v2)
            if true {
                v2 -= (v2 ^ (v2 / (((v2 * 336)) | 1)))
                v2 = v1
            }
            _ = v23
        } else {
            v24 := v5[1]
            _ = v24
        }
    } else {
        for fi25 := 0; fi25 < 17; fi25++ {
            for fi26 := 0; fi26 < 1; fi26++ {
                v27 := GenS1{F0: 1, F1: (("foo bar" + "a") + (v21 + (v21 + v21)))}
                v28 := GenS1{F0: 362, F1: "a"}
                _ = fi26
                _ = v27
                _ = v28
            }
            v21 = (v21 + v21)
            if (65535 != 348) {
                v21 = v21
            } else {
                for fi29 := 0; fi29 < 7; fi29++ {
                    v21 = (v21 + v21)
                    var v30 float32 = (1e10 - 1.5)
                    _ = fi29
                    _ = v30
                }
                fmt.Println(v2, v2, v2)
                v3 = append(v3, v2)
            }
            _ = fi25
        }
        _ = len(v3)
        if false {
            v21 = v21
            v31 := GenS1{F0: 0, F1: ""}
            _ = v31
        }
    }
    v32 := GenS2{F0: 4294967295}
    v33 := GenS1{F0: 1, F1: (v21 + ((v21 + ("foo bar" + v21)) + v21))}
    for fi34 := 0; fi34 < 12; fi34++ {
        var v35 uint = 249
        v21 = v21
        var v36 float32 = -2.25
        _ = fi34
        _ = v35
        _ = v36
    }
    v21 = v21
    var v37 float64 = (1.5 - float64(-2.25))
    v21 = ("hello" + (v21 + "a"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v4
    _ = v5
}
