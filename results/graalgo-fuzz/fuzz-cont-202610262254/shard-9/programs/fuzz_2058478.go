package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int32; F1 uint32; F2 int8 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint16) uint32 {
    var v5 float32 = 3.14159
    var v6 uint32 = 2
    if (v6 > 4294967295) {
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := GenS1{F0: uint8(p1), F1: -577}
                v5 = 1.5
                p1 = p1
                _ = fi8
                _ = v9
            }
            if false {
                v10 := GenS1{F0: 42, F1: 9223372036854775807}
                _ = v10
            }
            _ = fi7
        }
    }
    v6 = (v6 % ((p0) | 1))
    var v11 int64 = 0
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    _ = v6
    return (v6 ^ v6)
}
func genFn1(p0 int32) int32 {
    if false {
        if (!true) {
            v12 := []int{-140, 42, 565}
            if len(v12) > 0 {
                v12[0] = 9223372036854775807
            }
            _ = v12
        }
    } else {
        if true {
            v13 := []int{4294967295}
            if (-1 > 9223372036854775807) {
                p0 += p0
                var v14 bool = false
                if v14 {
                    var v15 int64 = 0
                    for fi16 := 0; fi16 < 11; fi16++ {
                        v13 = append(v13, 0)
                        _ = fi16
                    }
                    _ = v15
                } else {
                    v14 = (255 != 923)
                }
                _ = v14
            }
            _ = v13
        } else {
            if true {
                if false {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        for fi18 := 0; fi18 < 13; fi18++ {
                            fmt.Println(p0)
                            _ = fi18
                        }
                        v19 := GenS2{F0: (p0 >> ((255) & 63)), F1: 1, F2: -1}
                        for fi20 := 0; fi20 < 4; fi20++ {
                            var v21 uint8 = 128
                            switch ("foo bar" + "a") {
                            case "a":
                                v22 := GenS1{F0: v21, F1: -9223372036854775808}
                                _ = v22
                            }
                            _ = fi20
                            _ = v21
                        }
                        _ = fi17
                        _ = v19
                    }
                    if (0 == 0) {
                        fmt.Println(p0)
                    } else {
                        fmt.Println(p0)
                        v23 := []int{}
                        _ = v23
                    }
                    v24 := GenS2{F0: p0, F1: 1, F2: 127}
                    _ = v24
                }
                p0 += -85
            }
            if false {
                v25 := []int{2147483647}
                if len(v25) > 0 {
                    v25[0] = 32768
                }
                _ = v25
            } else {
                var v26 int64 = 95
                if (!true) {
                    if true {
                        var v27 string = "hello"
                        for fi28 := 0; fi28 < 6; fi28++ {
                            var v29 int8 = 1
                            v30 := GenS1{F0: 0, F1: v26}
                            v26 = -9223372036854775808
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        var v31 uint64 = 159
                        _ = v27
                        _ = v31
                    } else {
                        var v32 int8 = 100
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 4; fi33++ {
                        if (247 >= 255) {
                            p0 += p0
                            if (900 <= 729) {
                                p0 = 42
                            }
                            var v34 uint16 = 255
                            _ = v34
                        }
                        p0 = p0
                        v26 = int64(fi33)
                        _ = fi33
                    }
                    var v35 int64 = 0
                    _ = v35
                } else {
                    fmt.Println(v26)
                    if true {
                        var v36 string = "\n\t"
                        _ = v36
                    }
                }
                _ = v26
            }
            if false {
                for fi37 := 0; fi37 < 11; fi37++ {
                    switch (("hello" + "") + "a") {
                    case "a":
                        var v38 int8 = -66
                        _ = v38
                    case "hello":
                        p0 = p0
                    case "foo bar":
                        p0 = 353
                    default:
                        p0 += p0
                    }
                    _ = fi37
                }
            }
        }
    }
    v39 := []int{}
    if true {
        var v40 uint16 = 0
        for fi41 := 0; fi41 < 9; fi41++ {
            if (252 <= 0) {
                if len(v39) > 0 {
                    v39[0] = 127
                }
                v42 := GenS3{F0: "a", F1: 2147483647}
                v43 := GenS2{F0: p0, F1: 91, F2: 100}
                _ = v42
                _ = v43
            } else {
                var v44 string = "a"
                switch "foo bar" {
                case "a":
                    if (601 < 2) {
                        var v45 uint = 4294967295
                        _ = v45
                    }
                }
                var v46 uint32 = 42
                _ = v44
                _ = v46
            }
            _ = fi41
        }
        _ = v40
    } else {
        fmt.Println(p0)
        if ((float64((0.1 + 1.5)) / (0.0 + 3.14159)) <= 1e10) {
            fmt.Println(p0)
            var v47 int = 4294967295
            _ = v47
        } else {
            switch "\n\t" {
            case "a":
                switch ("a" + "foo bar") {
                case "\n\t":
                    if false {
                        fmt.Println(p0)
                    }
                case "foo bar":
                    fmt.Println(p0)
                }
            case "foo bar":
                switch -868 {
                case 0:
                    if len(v39) >= 2 {
                        v39 = v39[1:len(v39)]
                    }
                }
            case "\n\t":
                var v48 int64 = 9223372036854775807
                _ = v48
            }
            p0 = (p0 % (((p0 % (((p0 + (p0 + p0))) | 1))) | 1))
            switch "a" {
            case "a":
                if len(v39) >= 2 {
                    v39 = v39[1:len(v39)]
                }
            default:
                var v49 bool = (!(-1 > 82))
                _ = v49
            }
        }
    }
    if (32767 == -614) {
        p0 = p0
        if (42 > 2147483647) {
            var v50 uint8 = 128
            fmt.Println(v50)
            var v51 uint32 = 244
            _ = v50
            _ = v51
        } else {
            var v52 float64 = 3.14159
            _ = v52
        }
    } else {
        for fi53 := 0; fi53 < 8; fi53++ {
            p0 = (p0 - p0)
            if (255 < 0) {
                var v54 int32 = p0
                p0 = 42
                if false {
                    var v55 int64 = -9223372036854775808
                    v55 = (-(v55))
                    p0 -= p0
                    _ = v55
                } else {
                    var v56 uint16 = 468
                    _ = v56
                }
                _ = v54
            } else {
                p0 += p0
                var v57 uint8 = 205
                _ = len(v39)
                _ = v57
            }
            p0 += (p0 - -404)
            _ = fi53
        }
        if len(v39) > 0 {
            v39[0] = 127
        }
        p0 += (p0 | (p0 ^ p0))
    }
    _ = p0
    _ = v39
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = ((1 >= -551) && false)
    var v3 float32 = 1.5
    v4 := []string{"", "a"}
    fmt.Println(v3, v3)
    var v58 int64 = -812
    v59 := GenS4{F0: 0, F1: 50}
    if (false || (298 <= -509)) {
        var v60 int32 = genFn1(42)
        _ = v60
    } else {
        if (-59 == 11) {
            v2 = true
            switch -932 {
            case 167:
                _ = len(v1)
            case 600:
                v61 := GenS5{F0: 83}
                _ = v61
            case 42:
                v3 += ((0.0 / 0.1) * ((v3 / v3) / (v3 - v3)))
            }
            if (("\n\t" + ("a" + "a")) >= "hello") {
                v62 := GenS2{F0: 2147483647, F1: genFn0(1, 32767), F2: 1}
                _ = v62
            } else {
                if len(v4) > 0 {
                    v4[0] = ("" + "")
                }
                v2 = v2
                switch "\n\t" {
                case "\n\t":
                    if (2 != 139) {
                        var v63 float64 = -2.25
                        v63 = (v63 * (float64(float32(v63)) * (float64(-2.25) - (float64(((float32(0.0) + ((v3 * 1e10) / v3)) * 0.0)) + (-2.25 * v63)))))
                        _ = v63
                    }
                case "":
                    v64 := GenS3{F0: "hello", F1: 2147483648}
                    _ = v64
                default:
                    v2 = (true || (!(0 >= 65535)))
                }
            }
        } else {
            var v65 float64 = float64((v3 + v3))
            for fi66 := 0; fi66 < 19; fi66++ {
                var v67 string = ("\n\t" + "foo bar")
                v2 = (-6 != 42)
                for fi68 := 0; fi68 < 9; fi68++ {
                    if (true && false) {
                        v67 = v67
                    }
                    _ = fi68
                }
                _ = fi66
                _ = v67
            }
            _ = v65
        }
    }
    switch "a" {
    case "a":
        var v69 uint64 = 14
        _ = v69
    case "foo bar":
        var v70 int16 = 0
        _ = v70
    case "hello":
        var v71 uint32 = 1
        _ = v71
    }
    v58 = (v58 - (v58 ^ -678))
    v72 := GenS3{F0: "a", F1: 2}
    v1[1] = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v59
    _ = v72
}
