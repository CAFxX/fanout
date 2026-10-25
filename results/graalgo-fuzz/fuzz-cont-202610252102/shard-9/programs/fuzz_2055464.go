package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 bool; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 int16 }
func genFn0(p0 int64, p1 bool) int32 {
    fmt.Println(p0, p0)
    var v6 int64 = (p0 << ((255) & 63))
    _ = p0
    _ = p1
    _ = v6
    return 1
}
func genFn1(p0 int) int32 {
    v7 := GenS1{F0: 1, F1: (("hello" + "foo bar") + "a")}
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int = p0
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(p0)
    for fi11 := 0; fi11 < 12; fi11++ {
        switch "a" {
        case "":
            fmt.Println(fi11)
        case "\n\t":
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 uint = 1
                if (!(2 >= 18446744073709551615)) {
                    var v14 float64 = ((-2.25 / float64(3.14159)) + 0.1)
                    var v15 uint64 = 71
                    _ = v14
                    _ = v15
                } else {
                    if false {
                        v16 := []int{}
                        if len(v16) > 0 {
                            v16[0] = 32767
                        }
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        var v17 int8 = 13
                        if len(v16) > 0 {
                            v16[0] = fi11
                        }
                        fmt.Println(v13, v13, fi11)
                        _ = v16
                        _ = v17
                    } else {
                        v18 := GenS2{F0: false, F1: v13, F2: 128}
                        var v19 int64 = -276
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi12
                _ = v13
            }
        case "foo bar":
            p0 += fi11
        default:
            for fi20 := 0; fi20 < 18; fi20++ {
                fmt.Println(fi11)
                switch "\n\t" {
                case "hello":
                    if (!false) {
                        if false {
                            var v21 float32 = (-2.25 + 0.0)
                            p0 += fi20
                            fmt.Println(v21)
                            var v22 int32 = -2147483648
                            _ = v21
                            _ = v22
                        }
                    }
                case "\n\t":
                    if (!(true || true)) {
                        var v23 float32 = 0.1
                        fmt.Println(fi11, v23, p0)
                        if true {
                            v24 := GenS1{F0: -1, F1: ("\n\t" + (("a" + "foo bar") + ""))}
                            _ = v24
                        }
                        var v25 int16 = -88
                        if true {
                            v25 = v25
                            p0 += (fi20 - (p0 * -2147483648))
                            v26 := []int{fi11, (p0 + fi11), (fi11 >> ((2) & 63)), (fi11 * fi11)}
                            _ = v26
                        } else {
                            v27 := []int{(fi11 - (p0 % ((42) | 1))), 1, (fi11 << ((153) & 63))}
                            p0 = (p0 / ((fi20) | 1))
                            _ = v27
                        }
                        _ = v23
                        _ = v25
                    } else {
                        var v28 float32 = float32((-2.25 - ((float64(1e10) - (3.14159 + 3.14159)) - 1e10)))
                        _ = v28
                    }
                default:
                    var v29 bool = false
                    _ = v29
                }
                p0 -= (p0 >> ((18446744073709551615) & 63))
                _ = fi20
            }
        }
        if (!false) {
            v30 := []int{fi11, (fi11 ^ (p0 + p0))}
            fmt.Println(fi11, p0)
            _ = v30
        }
        for fi31 := 0; fi31 < 10; fi31++ {
            var v32 uint = 0
            _ = fi31
            _ = v32
        }
        _ = fi11
    }
    _ = p0
    _ = v7
    return -265
}

func main() {
    var v1 uint32 = 757
    var v2 uint = 42
    v3 := []string{"a"}
    var v4 int32 = 42
    v5 := make(map[int]bool)
    fmt.Println(v1, v4)
    for fi33 := 0; fi33 < 17; fi33++ {
        v34 := GenS1{F0: -1, F1: "foo bar"}
        v35 := GenS1{F0: -32768, F1: "foo bar"}
        _ = fi33
        _ = v34
        _ = v35
    }
    switch 32767 {
    case 32768:
        for fi36 := 0; fi36 < 17; fi36++ {
            v37 := GenS3{F0: (v2 | (^(v2)))}
            switch fi36 {
            case -9223372036854775808:
                fmt.Println(v1, v2)
            default:
                if false {
                    v1 = uint32(v2)
                }
            }
            _ = fi36
            _ = v37
        }
    case 9223372036854775807:
        var v38 uint32 = v1
        _ = v38
    default:
        v39 := GenS4{F0: 957, F1: false}
        _ = v39
    }
    var v40 int8 = -128
    var v41 uint = v2
    v42 := GenS5{F0: 1}
    for fi43 := 0; fi43 < 3; fi43++ {
        fmt.Println(v1, v1, v1)
        _ = fi43
    }
    v4 = v4
    v2 = (v41 / ((0) | 1))
    v41 = v2
    switch "foo bar" {
    case "a":
        if false {
            var v44 int64 = 0
            _ = v44
        }
    default:
        switch "\n\t" {
        case "\n\t":
            v2 += 255
        case "":
            v45 := GenS2{F0: false, F1: v41, F2: 176}
            _ = v45
        case "a":
            if ((1.5 * 3.14159) != 1.5) {
                _ = len(v5)
                v1 = v1
                v46 := GenS4{F0: 972, F1: false}
                _ = v46
            }
        }
    }
    var v47 int8 = 42
    v1 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v47
    _ = v5
}
