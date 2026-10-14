package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int16; F1 string; F2 int32 }
type GenS4 struct { F0 int64; F1 uint; F2 int8 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 uint }
func genFn0(p0 int32, p1 uint32) uint16 {
    for fi6 := 0; fi6 < 3; fi6++ {
        p1 = 2
        _ = fi6
    }
    v7 := []int{32768, 127}
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 127, F1: fi8}
        var v10 float64 = ((1e10 * 1e10) - (3.14159 - 3.14159))
        var v11 uint32 = p1
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        switch ("\n\t" + "hello") {
        case "\n\t":
            if len(v7) > 0 {
                v7[0] = (fi12 - (fi12 ^ (fi12 & (fi12 & -826))))
            }
        }
        fmt.Println(p1, fi12, fi12)
        if (-2.25 != (((0.0 - 1e10) + 0.1) * (-2.25 * 0.0))) {
            p1 = p1
        } else {
            if true {
                v13 := GenS1{F0: 171, F1: fi12}
                _ = v13
            } else {
                var v14 bool = false
                _ = v14
            }
        }
        _ = fi12
    }
    if false {
        p1 += (^(p1))
        fmt.Println(p1, p1)
        v15 := GenS1{F0: 42, F1: 129}
        _ = v15
    } else {
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = v7
    return 255
}
func genFn1(p0 int32, p1 int16, p2 int) int8 {
    switch (p2 * (p2 * (p2 ^ (p2 | p2)))) {
    case 42:
        if false {
            if (63 < 105) {
                var v16 uint16 = 265
                fmt.Println(p1, p1)
                var v17 string = ("foo bar" + ("" + ""))
                _ = v16
                _ = v17
            } else {
                fmt.Println(p2, p0, p2)
                v18 := GenS1{F0: 255, F1: p2}
                var v19 bool = (-432 == p0)
                _ = v18
                _ = v19
            }
            p1 = (p1 & (p1 % ((32767) | 1)))
            v20 := []int{(-(p2))}
            _ = v20
        }
    case 1:
        if true {
            if (true && false) {
                var v21 uint8 = 128
                _ = v21
            }
            var v22 int64 = 9223372036854775807
            _ = v22
        } else {
            fmt.Println(p2, p1, p2)
        }
    default:
        v23 := GenS2{F0: 153, F1: 42}
        _ = v23
    }
    p2 -= 574
    for fi24 := 0; fi24 < 20; fi24++ {
        v25 := GenS3{F0: 42, F1: ("a" + ("" + ((("hello" + (("\n\t" + ("hello" + "")) + "\n\t")) + "\n\t") + ("\n\t" + ("\n\t" + "foo bar"))))), F2: p0}
        v26 := GenS4{F0: int64(p0), F1: 42, F2: -115}
        v27 := []int{(fi24 | int(p0)), int(p0), (fi24 / ((2147483647) | 1)), (p2 | fi24)}
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    v28 := GenS5{F0: "foo bar", F1: 32767}
    var v29 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v29
    return 100
}

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 894
    var v4 uint = 411
    var v5 int16 = v1
    v30 := GenS5{F0: "foo bar", F1: v1}
    fmt.Println(v5, v3)
    if false {
        var v31 float32 = (3.14159 - ((1.5 + 3.14159) - 1e10))
        _ = v31
    } else {
        if (42 == 892) {
            var v32 int = -1
            var v33 int8 = int8(v5)
            v32 -= 351
            _ = v32
            _ = v33
        } else {
            if (!false) {
                fmt.Println(v4, v4, v5)
                v34 := v2[v3]
                _ = v34
            }
        }
    }
    if (!true) {
        v4 = uint(v3)
        v4 = (^(v4))
        var v35 uint32 = (v3 + v3)
        _ = v35
    } else {
        fmt.Println(v5)
        v4 = (^(v4))
        v3 -= (v3 << ((v4) & 63))
    }
    v5 = (v5 + -974)
    for fi36 := 0; fi36 < 17; fi36++ {
        for fi37 := 0; fi37 < 19; fi37++ {
            v38 := []int{(fi37 - (-(fi37))), (fi36 / ((fi36) | 1)), fi37, (fi37 << ((v4) & 63))}
            v39 := GenS6{F0: 0, F1: (((("foo bar" + "\n\t") + "\n\t") + "\n\t") + "")}
            v5 -= 32767
            _ = fi37
            _ = v38
            _ = v39
        }
        v40 := []int{fi36, (fi36 | (fi36 % ((9223372036854775807) | 1))), (fi36 - fi36)}
        if true {
            if true {
                var v41 string = ""
                var v42 int16 = (^(v1))
                _ = v41
                _ = v42
            }
            var v43 string = ("\n\t" + "a")
            _ = v43
        } else {
            v5 = v5
            v40 = append(v40, fi36)
            if (!true) {
                v5 = v5
            } else {
                for fi44 := 0; fi44 < 8; fi44++ {
                    fmt.Println(v4, v4, v4)
                    v5 = -32768
                    if (-1 >= 14) {
                        var v45 bool = false
                        if (((1.5 / 0.1) == (0.1 - (-2.25 + 0.1))) && ((v4 / (((v4 + (v4 ^ v4))) | 1)) <= v4)) {
                            var v46 int = int(v5)
                            if true {
                                var v47 int64 = 547
                                var v48 int16 = (v5 / ((-1) | 1))
                                v40 = append(v40, 9223372036854775807)
                                _ = v47
                                _ = v48
                            } else {
                                var v49 uint64 = 18446744073709551615
                                v1 += v1
                                v49 = 0
                                v3 = v3
                                _ = v49
                            }
                            if v45 {
                                var v50 int16 = v1
                                if len(v40) >= 2 {
                                    v40 = v40[1:len(v40)]
                                }
                                _ = v50
                            } else {
                                v4 = v4
                                v51 := GenS4{F0: 0, F1: (v4 | 65535), F2: 100}
                                _ = len(v2)
                                _ = v51
                            }
                            if true {
                                v52 := GenS7{F0: 1}
                                switch (fi36 << ((v4) & 63)) {
                                case 286:
                                    _ = len(v2)
                                case 32768:
                                    var v53 int32 = 1
                                    _ = v53
                                case 856:
                                    switch 128 {
                                    case 32768:
                                        switch fi36 {
                                        case -9223372036854775808:
                                            v54 := GenS8{F0: (false || (v45 || (39 > 42))), F1: 4294967295}
                                            _ = v54
                                        }
                                    default:
                                        var v55 string = "a"
                                        _ = v55
                                    }
                                }
                                if (!true) {
                                    fmt.Println(fi36, v46)
                                } else {
                                    v56 := GenS7{F0: v3}
                                    v4 = v4
                                    _ = v56
                                }
                                _ = v52
                            } else {
                                if v45 {
                                    v40 = append(v40, (-(v46)))
                                }
                                _ = len(v40)
                                v40 = append(v40, fi44)
                            }
                            if v45 {
                                fmt.Println(v1, fi36)
                                var v57 int64 = -691
                                var v58 float64 = (0.1 / 3.14159)
                                if len(v40) > 0 {
                                    v40[0] = fi44
                                }
                                var v59 uint32 = 4294967295
                                _ = v57
                                _ = v58
                                _ = v59
                            } else {
                                v2[(v3 >> ((151) & 63))] = (v3 | 2)
                            }
                            fmt.Println(v4)
                            var v60 bool = false
                            v61 := v2[2]
                            var v62 uint8 = 132
                            _ = v46
                            _ = v60
                            _ = v61
                            _ = v62
                        }
                        _ = v45
                    }
                    _ = fi44
                }
                var v63 uint = (v4 + (v4 | v4))
                _ = v63
            }
        }
        _ = fi36
        _ = v40
    }
    var v64 int64 = -943
    switch (((("foo bar" + "a") + ("\n\t" + "hello")) + "") + ("" + ("hello" + "\n\t"))) {
    case "hello":
        switch -9223372036854775808 {
        case 32767:
            for fi65 := 0; fi65 < 18; fi65++ {
                fmt.Println(v5, v4)
                fmt.Println(v4, v1)
                var v66 uint64 = 221
                _ = fi65
                _ = v66
            }
        case 32768:
            switch -38 {
            case 815:
                if (true && (2147483647 != -577)) {
                    v1 = (v1 >> ((v4) & 63))
                    var v67 int16 = (v5 / ((v1) | 1))
                    _ = v67
                }
            }
        default:
            var v68 uint32 = v3
            _ = v68
        }
    case "\n\t":
        if true {
            for fi69 := 0; fi69 < 20; fi69++ {
                v64 -= v64
                v70 := GenS3{F0: 1000, F1: "a", F2: -2147483648}
                var v71 int64 = (v64 - (v64 >> (((v4 | v4)) & 63)))
                _ = fi69
                _ = v70
                _ = v71
            }
        }
    case "":
        var v72 uint16 = 37
        _ = v72
    default:
        fmt.Println(v4, v3, v1)
    }
    for fi73 := 0; fi73 < 9; fi73++ {
        fmt.Println(fi73, v64, v4)
        _ = fi73
    }
    if (v3 >= 2) {
        switch "foo bar" {
        case "\n\t":
            if true {
                fmt.Println(v1, v64)
            } else {
                v1 = (v5 ^ v5)
                var v74 int32 = -364
                _ = v74
            }
        case "foo bar":
            var v75 uint = 951
            _ = v75
        }
        fmt.Println(v3)
        var v76 float64 = (0.0 * (0.1 * (1.5 * 0.1)))
        _ = v76
    }
    var v77 uint = 927
    if false {
        v64 -= v64
    }
    v64 = v64
    v78 := GenS4{F0: v64, F1: 178, F2: genFn1(42, v1, 2147483647)}
    var v79 uint8 = 4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v77:", v77)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v64
    _ = v77
    _ = v78
    _ = v79
}
