package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int8; F2 string }
type GenS5 struct { F0 uint16; F1 uint32 }
func genFn0() int {
    switch ("" + "\n\t") {
    case "\n\t":
        var v7 uint8 = 2
        _ = v7
    }
    var v8 bool = false
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := []int{fi9}
        v11 := GenS1{F0: "foo bar"}
        v12 := GenS1{F0: ""}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 int = -685
    _ = v13
    _ = v8
    return 257
}
func genFn1(p0 bool) uint32 {
    if ((!p0) || (1 >= -9223372036854775808)) {
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := []int{}
            _ = fi14
            _ = v15
        }
        v16 := GenS2{F0: 65535, F1: ""}
        if (!(65535 <= 699)) {
            fmt.Println(p0)
        }
        _ = v16
    } else {
        v17 := GenS1{F0: "a"}
        _ = v17
    }
    fmt.Println(p0)
    if false {
        v18 := []int{genFn0(), 1, 42, 0}
        p0 = p0
        p0 = (float64((-2.25 + float32(3.14159))) != 1e10)
        _ = v18
    } else {
        var v19 uint32 = 2147483647
        p0 = p0
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := GenS2{F0: 2, F1: ""}
            v19 = (v19 / ((v19) | 1))
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    if (4294967295 <= 9223372036854775807) {
        var v22 uint32 = 895
        if (65535 <= 2) {
            var v23 float32 = (0.0 - (0.1 + 1e10))
            var v24 uint16 = 1
            _ = v23
            _ = v24
        }
        p0 = p0
        _ = v22
    }
    if (-1 >= -1) {
        for fi25 := 0; fi25 < 19; fi25++ {
            var v26 uint16 = 539
            _ = fi25
            _ = v26
        }
        p0 = (215 > 255)
        var v27 uint64 = 105
        _ = v27
    } else {
        fmt.Println(p0)
        fmt.Println(p0)
        var v28 bool = (9223372036854775807 <= 296)
        _ = v28
    }
    _ = p0
    return 340
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]uint32)
    var v3 int64 = -282
    var v4 int8 = 42
    var v5 uint32 = 42
    var v6 uint32 = v5
    fmt.Println(v4, v5)
    if ((-128 == -1) || true) {
        v29 := GenS3{F0: (!false)}
        v30 := GenS2{F0: (v1 + uint16(v6)), F1: "\n\t"}
        var v31 int32 = 724
        _ = v29
        _ = v30
        _ = v31
    } else {
        for fi32 := 0; fi32 < 19; fi32++ {
            fmt.Println(v3, v3)
            for fi33 := 0; fi33 < 17; fi33++ {
                if ((!true) || ("\n\t" < ("" + "foo bar"))) {
                    if ((v4 | 0) <= 6) {
                        fmt.Println(v1)
                        switch "foo bar" {
                        case "a":
                            var v34 uint64 = 18446744073709551615
                            _ = v34
                        default:
                            v6 = genFn1((0 > 651))
                        }
                    } else {
                        v35 := []int{}
                        fmt.Println(v1, fi33, v1)
                        v35 = append(v35, fi32)
                        _ = v35
                    }
                    switch "foo bar" {
                    case "\n\t":
                        var v36 uint16 = v1
                        _ = v36
                    case "a":
                        v6 = v6
                    }
                }
                var v37 bool = true
                _ = fi33
                _ = v37
            }
            if ((v1 / ((697) | 1)) == v1) {
                var v38 int64 = v3
                v39 := []int{128, (fi32 % ((32767) | 1))}
                _ = v38
                _ = v39
            } else {
                v6 += v5
                if (!(false && false)) {
                    v40 := GenS2{F0: (v1 >> ((18446744073709551615) & 63)), F1: (("\n\t" + "hello") + "a")}
                    _ = v40
                }
                switch fi32 {
                case 128:
                    var v41 uint64 = 0
                    _ = v41
                case 127:
                    var v42 uint32 = (v5 + 42)
                    _ = v42
                default:
                    var v43 int8 = (v4 / ((v4) | 1))
                    _ = v43
                }
            }
            _ = fi32
        }
        fmt.Println(v4, v6)
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        if (!false) {
            v45 := v2[genFn1(((!true) || (0 == 142)))]
            for fi46 := 0; fi46 < 7; fi46++ {
                if false {
                    v47 := []int{genFn0()}
                    var v48 int64 = v3
                    _ = v47
                    _ = v48
                } else {
                    if true {
                        var v49 int64 = (v3 << ((2) & 63))
                        _ = v49
                    }
                    v50 := v2[(^(v5))]
                    var v51 int16 = -1
                    _ = v50
                    _ = v51
                }
                _ = fi46
            }
            v1 = (v1 + v1)
            _ = v45
        }
        v3 = 1
        _ = fi44
    }
    var v52 int64 = 423
    v6 += 2147483647
    switch 986 {
    case 128:
        v52 -= (-(v52))
    case -523:
        var v53 int = 4294967295
        _ = v53
    }
    if true {
        for fi54 := 0; fi54 < 14; fi54++ {
            v55 := GenS4{F0: false, F1: (v4 ^ (v4 >> ((483) & 63))), F2: ("" + "foo bar")}
            v5 = (v5 / (((v5 ^ v5)) | 1))
            _ = fi54
            _ = v55
        }
        switch "\n\t" {
        case "":
            fmt.Println(v3, v1)
        case "hello":
            var v56 int8 = (^(v4))
            _ = v56
        }
    } else {
        _ = len(v2)
        v57 := GenS5{F0: v1, F1: (v5 % (((^(v6))) | 1))}
        v6 += genFn1(false)
        _ = v57
    }
    v58 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v58
    _ = v6
}
