package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int; F1 int }
func genFn0(p0 int8) uint {
    p0 -= p0
    p0 += (p0 + p0)
    _ = p0
    return 4294967295
}
func genFn1(p0 uint) uint16 {
    if (1 >= 0) {
        var v6 string = "a"
        _ = v6
    }
    switch 1 {
    case 32767:
        p0 -= (p0 * (p0 - p0))
    default:
        switch 32767 {
        case 2147483647:
            p0 += (^(p0))
        case -1:
            var v7 uint16 = 0
            _ = v7
        default:
            var v8 bool = (158 <= 4294967295)
            _ = v8
        }
    }
    var v9 int = 0
    var v10 int32 = 232
    v9 = 42
    _ = p0
    _ = v10
    _ = v9
    return 488
}

func main() {
    var v1 int32 = -1
    v2 := make(map[string]bool)
    var v3 uint16 = 1
    var v4 int8 = -1
    var v5 int16 = 42
    v3 -= v3
    if false {
        if true {
            switch 481 {
            case 0:
                for fi11 := 0; fi11 < 14; fi11++ {
                    if (!(!(4294967295 >= 680))) {
                        v12 := GenS1{F0: 1, F1: "foo bar"}
                        v13 := GenS1{F0: -1, F1: ("a" + "a")}
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 42, F1: ("\n\t" + "foo bar")}
                    _ = fi11
                    _ = v14
                }
            default:
                v4 = v4
            }
            v15 := GenS2{F0: (v1 - v1)}
            _ = v15
        }
        if false {
            v1 = 0
            fmt.Println(v1, v5, v5)
        }
        if (4294967295 < 2147483647) {
            v16 := []int{2147483647, 32767}
            v17 := GenS3{F0: v4, F1: 82}
            if ("" > "") {
                var v18 bool = ("a" > ("hello" + "hello"))
                _ = v18
            }
            _ = v16
            _ = v17
        } else {
            if (!true) {
                v19 := GenS4{F0: (false || false), F1: 273}
                v5 = (v5 << ((255) & 63))
                _ = v19
            }
        }
    }
    fmt.Println(v3, v4)
    if true {
        _ = len(v2)
    }
    fmt.Println(v3, v1, v3)
    var v20 int = -1
    if (-1 >= -1) {
        var v21 int32 = -1
        var v22 int8 = v4
        v2[""] = (42 != -869)
        _ = v21
        _ = v22
    }
    v23 := GenS3{F0: v4, F1: 1}
    var v24 uint64 = 42
    v25 := []int{v20, v20}
    for fi26 := 0; fi26 < 17; fi26++ {
        v3 = v3
        fmt.Println(v24, v5, v24)
        var v27 uint16 = (v3 >> ((1) & 63))
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 float32 = 0.1
        v30 := GenS5{F0: fi28, F1: -2147483648}
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        v5 += v5
        _ = fi31
    }
    v32 := GenS2{F0: (v1 >> ((478) & 63))}
    for fi33 := 0; fi33 < 4; fi33++ {
        v5 = (v5 % (((v5 | v5)) | 1))
        v4 -= (-(v4))
        _ = fi33
    }
    v2["\n\t"] = (!(0 < int64(v24)))
    for fi34 := 0; fi34 < 7; fi34++ {
        v35 := GenS1{F0: -1, F1: "\n\t"}
        var v36 int16 = v5
        fmt.Println(v3, fi34, v3)
        _ = fi34
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
