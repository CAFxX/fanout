package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 string; F2 uint16 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string; F1 bool }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint }
func genFn0(p0 uint32, p1 int64, p2 string) uint8 {
    p0 -= p0
    if true {
        var v7 uint64 = 0
        _ = v7
    } else {
        if true {
            var v8 float32 = 0.0
            for fi9 := 0; fi9 < 2; fi9++ {
                switch (fi9 * 2147483647) {
                case -762:
                    var v10 int = -276
                    _ = v10
                default:
                    for fi11 := 0; fi11 < 13; fi11++ {
                        switch fi9 {
                        case 127:
                            v12 := GenS1{F0: 1}
                            _ = v12
                        case 42:
                            v13 := GenS1{F0: -32768}
                            var v14 bool = ((p1 << ((623) & 63)) >= (p1 / (((p1 ^ 529)) | 1)))
                            _ = v13
                            _ = v14
                        }
                        _ = fi11
                    }
                }
                _ = fi9
            }
            p2 = p2
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 uint64, p1 uint, p2 uint64) int32 {
    for fi15 := 0; fi15 < 9; fi15++ {
        p2 += p0
        v16 := []int{fi15, 42, 1, (fi15 / ((1) | 1))}
        _ = fi15
        _ = v16
    }
    v17 := []int{4294967295}
    for fi18 := 0; fi18 < 13; fi18++ {
        p1 -= (p1 ^ p1)
        if true {
            switch "hello" {
            case "":
                if (-2.25 > (-2.25 / 1.0)) {
                    p2 = (p2 >> ((4294967295) & 63))
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v20 := GenS1{F0: 916}
                        if (!false) {
                            var v21 int32 = 42
                            v22 := GenS2{F0: (p0 / ((p0) | 1)), F1: 42}
                            p2 += (p0 << (((p1 | 2)) & 63))
                            fmt.Println(v21)
                            var v23 uint16 = 884
                            switch (("a" + "hello") + "foo bar") {
                            case "foo bar":
                                if true {
                                    var v24 uint8 = genFn0(2, 736, "hello")
                                    var v25 uint64 = 1
                                    if false {
                                        v17 = append(v17, fi18)
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                        v17 = append(v17, 4294967295)
                                    } else {
                                        p0 -= p2
                                        var v26 int32 = (-(v21))
                                        var v27 uint8 = (v24 << ((p1) & 63))
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = v24
                                    _ = v25
                                } else {
                                    if true {
                                        if false {
                                            var v28 float64 = (0.1 - float64(((((((-2.25 / (float32(((0.1 - 3.14159) / 3.14159)) + (1.5 * ((-2.25 * (float32(((3.14159 * 3.14159) + 3.14159)) * (-2.25 - 0.0))) - 0.0)))) - 3.14159) - 0.0) + (0.0 * 1e10)) / (0.1 / float32(1.5))) * (1.5 / ((((-2.25 / 1e10) - 0.1) * 1e10) - 1e10)))))
                                            v29 := GenS3{F0: "\n\t"}
                                            _ = v28
                                            _ = v29
                                        } else {
                                            v21 += -209
                                            v30 := GenS1{F0: 1000}
                                            if len(v17) >= 2 {
                                                v17 = v17[1:len(v17)]
                                            }
                                            var v31 int16 = -949
                                            _ = v30
                                            _ = v31
                                        }
                                    }
                                    var v32 uint32 = 2147483648
                                    _ = v32
                                }
                            case "":
                                v33 := GenS4{F0: uint32(p2)}
                                _ = v33
                            case "a":
                                var v34 int8 = 45
                                if (!((p2 + (p2 * uint64(v23))) >= (p2 / ((92) | 1)))) {
                                    v35 := GenS4{F0: 1}
                                    v36 := GenS5{F0: 1}
                                    v21 = (-(v21))
                                    _ = v35
                                    _ = v36
                                } else {
                                    var v37 uint64 = (^(p2))
                                    v38 := GenS5{F0: 1}
                                    if (true || true) {
                                        var v39 string = ("foo bar" + "a")
                                        _ = v39
                                    } else {
                                        v37 = (^(v37))
                                    }
                                    _ = v37
                                    _ = v38
                                }
                                _ = v34
                            default:
                                if (!true) {
                                    if (!true) {
                                        v40 := GenS6{F0: p1, F1: "", F2: (v23 << ((p1) & 63))}
                                        v21 = v21
                                        p1 = (p1 - 18446744073709551615)
                                        _ = v40
                                    }
                                    var v41 bool = true
                                    _ = v41
                                } else {
                                    if len(v17) > 0 {
                                        v17[0] = (fi18 | 9223372036854775807)
                                    }
                                    v42 := GenS6{F0: p1, F1: (("a" + "\n\t") + "a"), F2: 0}
                                    v43 := GenS7{F0: (127 < 42), F1: (^(v21))}
                                    _ = v42
                                    _ = v43
                                }
                            }
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        _ = fi19
                        _ = v20
                    }
                    fmt.Println(p0, fi18, p2)
                }
            default:
                p0 = (p0 & p2)
            }
            var v44 float64 = 1.5
            v44 = v44
            _ = v44
        } else {
            if false {
                p1 = (^(p1))
                p0 += 42
            }
            v45 := GenS8{F0: (42 > -155)}
            var v46 uint = p1
            _ = v45
            _ = v46
        }
        switch "" {
        case "\n\t":
            if (true || false) {
                var v47 float32 = 3.14159
                switch 9223372036854775807 {
                case 26:
                    for fi48 := 0; fi48 < 11; fi48++ {
                        if true {
                            fmt.Println(v47, v47)
                            if false {
                                p1 -= 4294967295
                                var v49 int16 = 890
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                _ = v49
                            }
                            if false {
                                v50 := GenS9{F0: "hello", F1: ((0 < 726) && true)}
                                _ = v50
                            }
                        } else {
                            var v51 int32 = -1
                            p2 -= 0
                            _ = v51
                        }
                        _ = fi48
                    }
                default:
                    p2 = (p0 >> ((p1) & 63))
                }
                _ = v47
            } else {
                if false {
                    switch "foo bar" {
                    case "foo bar":
                        switch "\n\t" {
                        case "":
                            v17 = append(v17, fi18)
                        default:
                            p2 = (^(p0))
                        }
                    case "a":
                        var v52 float64 = 3.14159
                        _ = v52
                    }
                    v53 := GenS7{F0: (!false), F1: 215}
                    v54 := GenS10{F0: (^(fi18))}
                    _ = v53
                    _ = v54
                }
                fmt.Println(fi18, p0)
            }
        }
        _ = fi18
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return -415
}

func main() {
    var v1 int16 = 42
    v2 := make(map[int]bool)
    v3 := []uint32{1}
    v4 := []uint32{294, 302, 392, 42}
    var v5 string = ""
    v6 := []int{-140}
    _ = len(v2)
    var v55 uint = 702
    v56 := v2[-9223372036854775808]
    v55 = (v55 * v55)
    v56 = (v5 <= (v5 + v5))
    v55 += v55
    v57 := GenS11{F0: v55}
    var v58 int16 = (v1 | (-(v1)))
    var v59 float64 = ((((1.5 - (-2.25 / ((1.5 - ((float64((1.5 - (1e10 + (float32(float64((0.1 - (float32((((0.1 + 1.5) + (0.1 / 1.0)) * (0.0 * 0.0))) + -2.25)))) + 0.1)))) / 1.5) * 0.1)) - 3.14159))) * 0.1) * (1e10 + (((((3.14159 + (1e10 - (1e10 + 0.1))) + 1e10) / 1.0) * float64((1e10 - 3.14159))) * 0.1))) + (0.0 / 3.14159))
    fmt.Println(v59, v56, v1)
    v60 := GenS9{F0: "\n\t", F1: (true && (0 >= 148))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
