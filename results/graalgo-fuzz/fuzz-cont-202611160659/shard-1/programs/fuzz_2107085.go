package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int64; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0(p0 int64, p1 uint16, p2 uint64) bool {
    switch (("a" + "") + "hello") {
    case "":
        v7 := GenS1{F0: 255, F1: p2}
        _ = v7
    case "hello":
        v8 := GenS1{F0: 2, F1: (p2 - 4294967295)}
        _ = v8
    default:
        v9 := []int{}
        _ = v9
    }
    if (p2 <= (p2 + (p2 / ((p2) | 1)))) {
        v10 := GenS1{F0: 20, F1: (p2 & (p2 & p2))}
        _ = v10
    }
    var v11 uint = 1
    var v12 int8 = 1
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 float64 = 0.0
        if (false || (0 <= -1)) {
            switch "hello" {
            case "\n\t":
                p0 += int64(p2)
            case "foo bar":
                if (!(0 <= 806)) {
                    switch (fi13 % (((fi13 + (fi13 | fi13))) | 1)) {
                    case 0:
                        v12 = (v12 + v12)
                    default:
                        v15 := []int{(-(fi13))}
                        _ = v15
                    }
                    v12 = v12
                } else {
                    switch (^(fi13)) {
                    case -358:
                        for fi16 := 0; fi16 < 15; fi16++ {
                            p0 = (p0 << ((v11) & 63))
                            _ = fi16
                        }
                    case 108:
                        var v17 uint32 = 312
                        _ = v17
                    case -1:
                        var v18 float64 = 1.5
                        _ = v18
                    }
                    var v19 int32 = -898
                    _ = v19
                }
            case "a":
                var v20 float32 = (-2.25 - (((1e10 * 1e10) - (1e10 * ((-2.25 - 1.5) + -2.25))) - 1e10))
                _ = v20
            }
            for fi21 := 0; fi21 < 11; fi21++ {
                var v22 int16 = int16(p0)
                v23 := GenS2{F0: ("a" + "foo bar"), F1: v22}
                _ = fi21
                _ = v22
                _ = v23
            }
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return true
}
func genFn1(p0 int32, p1 uint64) int {
    switch "foo bar" {
    case "\n\t":
        v24 := []int{-1, 2147483647, 757, 421}
        _ = v24
    case "hello":
        if (42 == -1) {
            var v25 int8 = -118
            var v26 uint16 = 32767
            _ = v25
            _ = v26
        }
    }
    v27 := GenS1{F0: 127, F1: (p1 + 487)}
    _ = p0
    _ = p1
    _ = v27
    return -1
}

func main() {
    var v1 uint32 = 1
    var v2 bool = (((true || (!(((1.5 + 1e10) > 1e10) || ((float32(-2.25) + 1e10) > 0.1)))) && (!true)) || false)
    v3 := make(map[string]uint32)
    var v4 bool = ("a" < ("foo bar" + "a"))
    var v5 int32 = -275
    var v6 int8 = 100
    if (int32(v1) >= -73) {
        if (94 != 514) {
            if (1 >= 1000) {
                switch 2147483647 {
                case 32768:
                    v6 = (v6 << ((622) & 63))
                case -9223372036854775808:
                    v3["foo bar"] = v1
                case 0:
                    v2 = (genFn1(v5, 18446744073709551615) <= 449)
                }
                var v28 uint32 = v1
                var v29 int = -1
                _ = v28
                _ = v29
            }
            var v30 int = 42
            _ = v30
        } else {
            for fi31 := 0; fi31 < 12; fi31++ {
                for fi32 := 0; fi32 < 15; fi32++ {
                    if ((!(!false)) && (v4 || ((float32(-2.25) >= 3.14159) || v4))) {
                        var v33 uint64 = 829
                        var v34 uint = 18446744073709551615
                        delete(v3, "a")
                        _ = v33
                        _ = v34
                    } else {
                        var v35 int64 = -1
                        _ = v35
                    }
                    v36 := GenS2{F0: "", F1: 1}
                    _ = fi32
                    _ = v36
                }
                v37 := []int{genFn1((^(v5)), 4294967295)}
                _ = fi31
                _ = v37
            }
            fmt.Println(v5, v6)
            for fi38 := 0; fi38 < 19; fi38++ {
                if v2 {
                    v2 = true
                    if (65535 >= 1) {
                        var v39 int32 = 2147483647
                        if (3.14159 < (float32(float64(3.14159)) + 1e10)) {
                            v40 := GenS3{F0: 1, F1: v1, F2: 32767}
                            var v41 uint8 = 0
                            if (1e10 > (((0.1 * 3.14159) * ((1.5 - 3.14159) / 1.0)) + 1.5)) {
                                v2 = (1 < 703)
                                v6 -= (v6 / ((1) | 1))
                            }
                            _ = v40
                            _ = v41
                        } else {
                            for fi42 := 0; fi42 < 16; fi42++ {
                                if v4 {
                                    var v43 uint64 = 508
                                    _ = v43
                                }
                                v44 := GenS4{F0: 843}
                                var v45 int = (fi38 / ((fi42) | 1))
                                _ = fi42
                                _ = v44
                                _ = v45
                            }
                        }
                        _ = v39
                    } else {
                        for fi46 := 0; fi46 < 17; fi46++ {
                            var v47 uint64 = 694
                            v5 = v5
                            if v4 {
                                v2 = ((((genFn1(0, (^(v47))) == int(v5)) && v2) && v2) || v2)
                            } else {
                                v48 := GenS3{F0: -337, F1: v1, F2: -32768}
                                _ = v48
                            }
                            _ = fi46
                            _ = v47
                        }
                        v4 = v4
                        var v49 int8 = v6
                        _ = v49
                    }
                    for fi50 := 0; fi50 < 14; fi50++ {
                        v5 = v5
                        v2 = (1 > 32767)
                        _ = fi50
                    }
                }
                var v51 float64 = 1e10
                _ = fi38
                _ = v51
            }
        }
        var v52 uint32 = (v1 - v1)
        var v53 string = "foo bar"
        _ = v52
        _ = v53
    }
    v2 = (true && false)
    switch "foo bar" {
    case "hello":
        if (4294967295 <= 863) {
            if v2 {
                v54 := GenS4{F0: 4294967295}
                v55 := []int{}
                _ = v54
                _ = v55
            }
            v56 := []int{}
            var v57 uint8 = 90
            _ = v56
            _ = v57
        } else {
            v58 := GenS5{F0: 6, F1: v5}
            if (!false) {
                v4 = v2
                var v59 bool = ((!(!(128 != 1))) && false)
                v2 = ((42 >= -128) || true)
                _ = v59
            }
            _ = v58
        }
    default:
        fmt.Println(v4, v1)
    }
    var v60 float64 = 1.5
    fmt.Println(v60, v1)
    var v61 bool = (-281 <= -1)
    switch -9223372036854775808 {
    case -9223372036854775808:
        if false {
            var v62 uint16 = 1
            v63 := []int{}
            _ = v62
            _ = v63
        }
    case -1:
        var v64 int32 = 1
        _ = v64
    default:
        v5 = (v5 % (((v5 / (((v5 | (v5 / (((v5 % ((v5) | 1))) | 1)))) | 1))) | 1))
    }
    var v65 float64 = v60
    v66 := []int{1, 127}
    _ = len(v66)
    v67 := GenS4{F0: 979}
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v66:", len(v66), v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
    _ = v65
    _ = v66
    _ = v67
}
