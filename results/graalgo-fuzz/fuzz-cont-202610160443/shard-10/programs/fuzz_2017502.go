package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 int64; F1 int16 }
func genFn0(p0 uint, p1 int64, p2 uint) uint16 {
    fmt.Println(p0, p0, p1)
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := GenS1{F0: 728, F1: 862}
        fmt.Println(fi4, p2, p1)
        v6 := GenS1{F0: 2, F1: 2}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int64 = p1
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint = 524
        if true {
            if (p1 >= (v7 / (((p1 << ((p2) & 63))) | 1))) {
                var v10 float64 = 3.14159
                p1 -= (v7 % ((v7) | 1))
                if true {
                    var v11 int64 = (-(p1))
                    var v12 uint64 = 795
                    _ = v11
                    _ = v12
                }
                _ = v10
            }
            if (212 <= 255) {
                var v13 uint8 = 128
                v7 -= v7
                if true {
                    var v14 uint8 = v13
                    v15 := GenS2{F0: 42, F1: "foo bar"}
                    _ = v14
                    _ = v15
                }
                _ = v13
            }
        }
        v16 := []int{fi8, fi8, fi8}
        _ = fi8
        _ = v16
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}
func genFn1(p0 int8, p1 uint16) uint64 {
    switch ("a" + "") {
    case "":
        if (4294967295 >= 433) {
            if (!(772 == 0)) {
                v17 := GenS3{F0: 1, F1: 128}
                _ = v17
            } else {
                var v18 uint64 = 1
                _ = v18
            }
            if false {
                var v19 int8 = (p0 + -128)
                _ = v19
            } else {
                fmt.Println(p0, p0)
                if (1.5 >= float32(((1e10 * 0.0) * (3.14159 + -2.25)))) {
                    p0 += p0
                }
            }
            if (!true) {
                p0 = p0
            } else {
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 uint = 1
                    v21 = v21
                    _ = fi20
                    _ = v21
                }
                fmt.Println(p0)
            }
        }
    case "hello":
        p1 = p1
    default:
        if (((float64((1e10 + 3.14159)) * (-2.25 - 1.5)) + 1.5) >= 3.14159) {
            v22 := []int{136, -471}
            _ = v22
        } else {
            switch 972 {
            case 128:
                v23 := []int{0, 386}
                _ = v23
            case -1:
                var v24 uint64 = 4294967295
                _ = v24
            case 32767:
                v25 := []int{9223372036854775807, 50}
                _ = v25
            }
            var v26 uint64 = 42
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 uint64 = 959
        _ = fi27
        _ = v28
    }
    var v29 uint = 42
    _ = p0
    _ = p1
    _ = v29
    return 773
}

func main() {
    v1 := []string{(((("\n\t" + "\n\t") + ("\n\t" + "hello")) + "a") + "hello")}
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v30 int8 = -55
    if (!((756 >= 10) || ((9223372036854775807 > 291) && false))) {
        v30 += -107
        fmt.Println(v30)
    }
    var v31 uint64 = 570
    if (111 == 32767) {
        v32 := GenS1{F0: (^(v31)), F1: 778}
        _ = v32
    }
    var v33 float32 = float32((0.1 + (3.14159 * (1e10 * 1e10))))
    fmt.Println(v33, v30, v30)
    switch -1 {
    case 128:
        switch 42 {
        case 4294967295:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -1:
            v34 := GenS4{F0: 53}
            _ = v34
        default:
            var v35 float64 = 1e10
            _ = v35
        }
    case 127:
        v31 -= (v31 * v31)
    case -793:
        var v36 string = "foo bar"
        _ = v36
    default:
        var v37 uint16 = 65535
        _ = v37
    }
    v31 += 644
    for fi38 := 0; fi38 < 5; fi38++ {
        var v39 uint8 = 0
        _ = fi38
        _ = v39
    }
    if true {
        v30 += int8(v31)
        v40 := GenS5{F0: -1, F1: v30}
        _ = v40
    }
    v41 := GenS1{F0: 0, F1: 2}
    v42 := GenS6{F0: 0, F1: -882}
    var v43 string = ("foo bar" + (((("" + ((("\n\t" + "") + "foo bar") + "a")) + "") + "\n\t") + (("a" + "") + "foo bar")))
    for fi44 := 0; fi44 < 4; fi44++ {
        if len(v1) > 0 {
            v1[0] = v43
        }
        _ = fi44
    }
    if true {
        v45 := GenS2{F0: -1, F1: "a"}
        _ = v45
    } else {
        v43 = v43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v41
    _ = v42
    _ = v43
}
