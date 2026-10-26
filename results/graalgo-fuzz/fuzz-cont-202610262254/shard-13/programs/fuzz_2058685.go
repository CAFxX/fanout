package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint; F1 string; F2 bool }
func genFn0(p0 uint64, p1 uint8, p2 uint32) uint8 {
    var v5 int8 = 0
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint16 = 42
        var v8 uint8 = p1
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 int8 = v5
            var v11 uint32 = (p2 << ((491) & 63))
            if ((v8 * p1) != v8) {
                v12 := []int{0, fi9}
                if len(v12) > 0 {
                    v12[0] = 4294967295
                }
                v10 += (^(v5))
                _ = v12
            } else {
                if (!(!false)) {
                    switch "\n\t" {
                    case "":
                        fmt.Println(fi9, p0, v11)
                    case "\n\t":
                        v13 := []int{fi9}
                        _ = v13
                    case "hello":
                        var v14 int = int(p2)
                        _ = v14
                    }
                    if (p1 >= (v8 | v8)) {
                        switch "hello" {
                        case "a":
                            v15 := []int{(fi9 | int(v8))}
                            _ = v15
                        case "\n\t":
                            var v16 int8 = (v5 << ((0) & 63))
                            _ = v16
                        case "hello":
                            p1 -= v8
                        default:
                            if (!((0.1 + (0.0 * 1.5)) != float64(1e10))) {
                                v17 := []int{32767, -1, (-(fi6))}
                                if (-496 >= 574) {
                                    var v18 float32 = 3.14159
                                    if false {
                                        p0 += 0
                                        _ = len(v17)
                                    } else {
                                        p1 += 255
                                    }
                                    p1 = v8
                                    _ = v18
                                } else {
                                    var v19 int = fi6
                                    v20 := GenS1{F0: uint64(p1), F1: 2, F2: 0}
                                    var v21 int8 = (v10 / ((v10) | 1))
                                    v10 -= 16
                                    var v22 uint = 258
                                    var v23 int64 = 42
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                                _ = v17
                            }
                            if (65535 < 374) {
                                var v24 int16 = int16(fi9)
                                _ = v24
                            } else {
                                switch ("" + "foo bar") {
                                case "a":
                                    v11 -= 4294967295
                                }
                                switch "a" {
                                case "a":
                                    switch ((("a" + "foo bar") + "foo bar") + ("a" + "")) {
                                    case "foo bar":
                                        fmt.Println(v7)
                                    case "":
                                        if false {
                                            v25 := []int{}
                                            if (p2 >= (v11 >> ((889) & 63))) {
                                                fmt.Println(fi6, p1, v7)
                                                switch "foo bar" {
                                                case "":
                                                    if len(v25) >= 2 {
                                                        v25 = v25[1:len(v25)]
                                                    }
                                                default:
                                                    if true {
                                                        var v26 uint32 = p2
                                                        p1 += v8
                                                        _ = v26
                                                    }
                                                }
                                                var v27 uint = 4294967295
                                                _ = v27
                                            }
                                            _ = v25
                                        } else {
                                            v28 := GenS1{F0: 993, F1: p2, F2: (v5 * -128)}
                                            if ((true && (!false)) && (false && (!(3.14159 == (1.5 + ((1.5 + (float32(0.0) - ((1e10 / (float32(3.14159) / (3.14159 + 0.0))) * 0.0))) * -2.25)))))) {
                                                v29 := GenS1{F0: p0, F1: p2, F2: v5}
                                                v30 := GenS2{F0: -99, F1: 1000}
                                                v10 = v5
                                                _ = v29
                                                _ = v30
                                            }
                                            _ = v28
                                        }
                                    case "\n\t":
                                        if (true && (((1.5 * 0.0) / ((0.1 * float32((((3.14159 * 0.1) + 1e10) - (0.1 * 1.5)))) + (0.0 / 1.5))) < float32(1.5))) {
                                            v31 := []int{4294967295}
                                            _ = v31
                                        }
                                    default:
                                        v8 += 235
                                    }
                                case "hello":
                                    if false {
                                        var v32 float32 = 1.5
                                        _ = v32
                                    }
                                case "\n\t":
                                    var v33 float32 = 0.0
                                    _ = v33
                                default:
                                    v34 := []int{}
                                    _ = v34
                                }
                                fmt.Println(fi9, v5, p0)
                                v35 := []int{963, (fi9 | fi9), 4294967295}
                                _ = v35
                            }
                            fmt.Println(v10)
                        }
                    } else {
                        if false {
                            fmt.Println(v10)
                        } else {
                            v36 := []int{2147483647}
                            if len(v36) >= 2 {
                                v36 = v36[1:len(v36)]
                            }
                            _ = v36
                        }
                    }
                }
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}

func main() {
    var v1 int = 127
    var v2 int8 = -77
    v3 := make(map[int]int)
    var v4 uint16 = 281
    if true {
        var v37 int = -2147483648
        if true {
            v4 += (v4 ^ (v4 >> ((1) & 63)))
            v4 += v4
            fmt.Println(v4, v1)
        }
        v1 -= 81
        _ = v37
    }
    var v38 int16 = 42
    v4 += (^(v4))
    if true {
        if false {
            delete(v3, v1)
        } else {
            var v39 int8 = (v2 / ((v2) | 1))
            switch "a" {
            case "foo bar":
                v3[(-(v1))] = (-(v1))
            case "hello":
                fmt.Println(v1, v1)
            case "a":
                v4 = 447
            }
            _ = v39
        }
        if false {
            v40 := v3[1]
            _ = v40
        } else {
            var v41 uint = 1
            if false {
                var v42 uint64 = 696
                _ = len(v3)
                v43 := []int{(v1 | v1), v1, v1}
                _ = v42
                _ = v43
            }
            _ = v41
        }
    }
    if (-2.25 < (float32(-2.25) / float32(1e10))) {
        var v44 int = v1
        _ = v44
    } else {
        if (v4 == (v4 >> ((300) & 63))) {
            v45 := []int{746, v1, v1, (v1 * -457)}
            _ = v45
        } else {
            var v46 float32 = 0.0
            _ = v46
        }
        v38 = (v38 % (((v38 & -1)) | 1))
        if (!true) {
            if (false && (29 == 42)) {
                if true {
                    fmt.Println(v4)
                } else {
                    switch v1 {
                    case 903:
                        var v47 float32 = -2.25
                        _ = v47
                    case 1:
                        v4 = v4
                    case 127:
                        v2 = (^(v2))
                    default:
                        v2 += v2
                    }
                }
            } else {
                for fi48 := 0; fi48 < 18; fi48++ {
                    var v49 string = "\n\t"
                    if ((v38 % (((v38 + v38)) | 1)) == v38) {
                        var v50 uint = 18446744073709551615
                        if ((true && (true && (true && false))) || true) {
                            v50 -= v50
                            v51 := GenS2{F0: 1, F1: (v38 % ((int16(v50)) | 1))}
                            v52 := GenS1{F0: 1, F1: 2147483648, F2: v2}
                            _ = v51
                            _ = v52
                        } else {
                            v53 := []int{42, v1}
                            if true {
                                switch v1 {
                                case 42:
                                    for fi54 := 0; fi54 < 17; fi54++ {
                                        fmt.Println(v1, v1)
                                        _ = fi54
                                    }
                                default:
                                    v55 := GenS1{F0: 2, F1: 529, F2: 127}
                                    _ = v55
                                }
                            } else {
                                v56 := GenS3{F0: v50, F1: (v49 + v49), F2: false}
                                for fi57 := 0; fi57 < 4; fi57++ {
                                    v58 := GenS3{F0: uint(v4), F1: "", F2: ((v50 / (((v50 & v50)) | 1)) < (v50 ^ v50))}
                                    v53 = append(v53, (fi57 << ((186) & 63)))
                                    fmt.Println(v1, fi57, fi48)
                                    _ = fi57
                                    _ = v58
                                }
                                _ = len(v53)
                                _ = v56
                            }
                            if (!(42 >= 18446744073709551615)) {
                                fmt.Println(v50)
                                _ = len(v53)
                                var v59 int64 = 561
                                _ = v59
                            }
                            _ = v53
                        }
                        _ = v50
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            if false {
                                var v60 int32 = -2147483648
                                v61 := GenS2{F0: v60, F1: v38}
                                delete(v3, (v1 | 32767))
                                _ = v60
                                _ = v61
                            }
                        case "\n\t":
                            for fi62 := 0; fi62 < 7; fi62++ {
                                var v63 string = "a"
                                _ = fi62
                                _ = v63
                            }
                        default:
                            for fi64 := 0; fi64 < 3; fi64++ {
                                var v65 uint64 = 823
                                _ = fi64
                                _ = v65
                            }
                        }
                        var v66 bool = (-2147483648 <= -1)
                        _ = v66
                    }
                    _ = fi48
                    _ = v49
                }
                var v67 int64 = -1
                _ = v67
            }
            v68 := GenS2{F0: 2147483647, F1: v38}
            _ = v68
        }
    }
    v2 -= (v2 % ((v2) | 1))
    v69 := []int{(^(v1)), (^(v1)), (-(v1)), 32768}
    v69 = append(v69, v1)
    var v70 uint32 = 1
    for fi71 := 0; fi71 < 16; fi71++ {
        var v72 bool = true
        fmt.Println(v2, v72, v1)
        _ = fi71
        _ = v72
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v69:", len(v69), v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v69
    _ = v70
}
