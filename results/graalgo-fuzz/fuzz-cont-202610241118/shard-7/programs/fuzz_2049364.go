package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 int8 }
func genFn0(p0 int8) int32 {
    var v5 int32 = 14
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 128, F1: 255, F2: v5}
        p0 = 127
        var v8 int16 = 32767
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint8 = 0
        var v11 uint16 = 1
        switch 42 {
        case 2147483647:
            v12 := []int{(fi9 ^ int(v5))}
            _ = v12
        case -9223372036854775808:
            if (!true) {
                switch ("foo bar" + ("foo bar" + "\n\t")) {
                case "":
                    v11 -= 32767
                case "foo bar":
                    var v13 uint16 = (^(v11))
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 8; fi14++ {
                        var v15 int = (^(fi14))
                        p0 = 0
                        _ = fi14
                        _ = v15
                    }
                }
                v5 -= 0
                p0 -= p0
            } else {
                v11 = (v11 + v11)
                v16 := GenS1{F0: (v10 * v10), F1: v11, F2: v5}
                _ = v16
            }
        case 581:
            v17 := []int{(fi9 << ((255) & 63)), fi9, fi9, 42}
            _ = v17
        default:
            if false {
                v18 := []int{}
                _ = v18
            } else {
                var v19 uint = 0
                v10 = v10
                var v20 int = 613
                _ = v19
                _ = v20
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    p0 = (^(p0))
    _ = p0
    _ = v5
    return (v5 - (-(v5)))
}
func genFn1(p0 int, p1 int64, p2 int8) uint64 {
    for fi21 := 0; fi21 < 20; fi21++ {
        if (2 != 473) {
            if (!true) {
                p2 -= (^(p2))
            } else {
                p1 += p1
            }
            v22 := []int{}
            _ = v22
        } else {
            if true {
                p0 -= (p0 / ((-555) | 1))
                switch "\n\t" {
                case "a":
                    if (!(113 > 4294967295)) {
                        var v23 uint32 = 535
                        var v24 int = (-(fi21))
                        var v25 int64 = (^(p1))
                        _ = v23
                        _ = v24
                        _ = v25
                    } else {
                        if false {
                            p2 = p2
                            for fi26 := 0; fi26 < 14; fi26++ {
                                if (false || false) {
                                    var v27 int8 = p2
                                    _ = v27
                                } else {
                                    if false {
                                        v28 := GenS1{F0: 146, F1: 32767, F2: -1}
                                        var v29 uint32 = 842
                                        v30 := []int{}
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v31 := GenS1{F0: 1, F1: 32767, F2: -255}
                                        _ = v31
                                    }
                                    var v32 int64 = (p1 << ((4294967295) & 63))
                                    v32 += v32
                                    _ = v32
                                }
                                p1 = -9223372036854775808
                                p0 += 1
                                _ = fi26
                            }
                            v33 := []int{p0, 2147483647, p0}
                            _ = v33
                        } else {
                            for fi34 := 0; fi34 < 17; fi34++ {
                                var v35 int64 = p1
                                v36 := GenS2{F0: 255, F1: 18446744073709551615}
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                            v37 := []int{}
                            if true {
                                for fi38 := 0; fi38 < 12; fi38++ {
                                    switch (fi21 / (((fi21 * 644)) | 1)) {
                                    case -979:
                                        if (false && (!((128 > 2) || (((float64(1e10) - 3.14159) / 1.5) <= -2.25)))) {
                                            p1 += (p1 * p1)
                                        }
                                    case 1:
                                        if false {
                                            v39 := GenS3{F0: 1}
                                            p0 = p0
                                            if true {
                                                p2 -= p2
                                                p2 -= (p2 << ((2) & 63))
                                                fmt.Println(p0, p1)
                                                p0 -= (-(fi21))
                                                var v40 uint32 = 65535
                                                _ = v40
                                            } else {
                                                var v41 int32 = genFn0((p2 | p2))
                                                _ = v41
                                            }
                                            if true {
                                                _ = len(v37)
                                            }
                                            _ = v39
                                        }
                                    case 9223372036854775807:
                                        var v42 string = (("hello" + "foo bar") + ("\n\t" + "hello"))
                                        _ = v42
                                    }
                                    if ((p0 < (fi38 - fi21)) && true) {
                                        switch p0 {
                                        case 145:
                                            fmt.Println(p0)
                                        case 32767:
                                            p0 = fi21
                                        case 817:
                                            fmt.Println(p2)
                                        default:
                                            if ((!true) || false) {
                                                if true {
                                                    if len(v37) > 0 {
                                                        v37[0] = p0
                                                    }
                                                    if (1000 >= 1) {
                                                        _ = len(v37)
                                                    }
                                                }
                                                if false {
                                                    var v43 uint64 = 1
                                                    _ = v43
                                                } else {
                                                    var v44 uint16 = 683
                                                    p2 -= (p2 | (p2 / ((p2) | 1)))
                                                    _ = v44
                                                }
                                                p0 -= (fi38 - fi21)
                                                fmt.Println(fi38, fi21)
                                                fmt.Println(p0, fi21)
                                                if (128 > 2) {
                                                    var v45 float64 = (-2.25 - 1e10)
                                                    v46 := GenS4{F0: -1}
                                                    v37 = append(v37, (fi38 + (fi38 >> ((976) & 63))))
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    var v47 int64 = (p1 + 1)
                                                    v37 = append(v37, int(p2))
                                                    _ = v47
                                                }
                                                v48 := GenS3{F0: 1}
                                                if (!true) {
                                                    p0 = fi38
                                                    if ((3.14159 > 0.0) && true) {
                                                        v49 := GenS1{F0: 255, F1: 42, F2: 0}
                                                        switch "hello" {
                                                        case "hello":
                                                            fmt.Println(p1, fi38, p0)
                                                        }
                                                        _ = v49
                                                    } else {
                                                        p1 -= (p1 + p1)
                                                        v50 := GenS1{F0: 214, F1: 651, F2: 1}
                                                        if false {
                                                            p2 = 127
                                                            p1 = p1
                                                            var v51 uint32 = 5
                                                            _ = v51
                                                        }
                                                        _ = v50
                                                    }
                                                    var v52 string = "foo bar"
                                                    _ = v52
                                                } else {
                                                    if true {
                                                        var v53 uint16 = 98
                                                        var v54 int8 = 42
                                                        if false {
                                                            v53 -= 65535
                                                        }
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        var v55 int16 = 674
                                                        v56 := GenS1{F0: 42, F1: 42, F2: 2147483647}
                                                        p1 -= p1
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                }
                                                _ = v48
                                            } else {
                                                p2 = (^(p2))
                                                var v57 string = "foo bar"
                                                v57 = "foo bar"
                                                _ = v57
                                            }
                                        }
                                    }
                                    _ = fi38
                                }
                            }
                            _ = v37
                        }
                    }
                case "":
                    for fi58 := 0; fi58 < 7; fi58++ {
                        if (1e10 != (((1.5 / float64(float32(1e10))) + 1e10) * (0.0 * float64(1.5)))) {
                            var v59 uint16 = 137
                            p1 = (p1 % ((-9223372036854775808) | 1))
                            _ = v59
                        }
                        var v60 bool = true
                        v61 := GenS4{F0: 1000}
                        _ = fi58
                        _ = v60
                        _ = v61
                    }
                }
            }
            v62 := GenS5{F0: 2147483648, F1: 772}
            if true {
                switch (p0 * fi21) {
                case 127:
                    var v63 float32 = -2.25
                    _ = v63
                case -441:
                    v64 := GenS6{F0: (p0 / ((fi21) | 1)), F1: ("\n\t" + "")}
                    _ = v64
                }
            } else {
                switch "\n\t" {
                case "hello":
                    var v65 float32 = 0.0
                    _ = v65
                case "":
                    switch "" {
                    case "foo bar":
                        if true {
                            var v66 int16 = 0
                            _ = v66
                        }
                    }
                case "foo bar":
                    var v67 int64 = -962
                    _ = v67
                default:
                    for fi68 := 0; fi68 < 20; fi68++ {
                        var v69 bool = (!(((false || (3.14159 >= (-2.25 * ((1e10 * float32(1e10)) - (-2.25 * 1e10))))) || true) && (-2.25 != (0.1 + ((0.1 - (0.1 + float64(1.5))) - 3.14159)))))
                        _ = fi68
                        _ = v69
                    }
                }
            }
            _ = v62
        }
        p1 -= p1
        _ = fi21
    }
    p2 -= p2
    p1 -= 1
    var v70 int = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v70
    return 657
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -2147483648
    var v3 uint16 = 494
    v4 := make(map[uint32]int)
    fmt.Println(v3, v3)
    v1[197] = 4294967295
    v2 = (v2 >> ((33) & 63))
    var v71 bool = true
    v3 = v3
    var v72 uint16 = 160
    v73 := GenS5{F0: 42, F1: 4294967295}
    v74 := GenS7{F0: 42}
    v75 := GenS6{F0: -1, F1: ""}
    var v76 bool = v71
    fmt.Println(v71, v72)
    var v77 uint8 = 42
    var v78 uint64 = 0
    v2 += (v2 | (v2 / ((v2) | 1)))
    fmt.Println(v3, v78)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v71
    _ = v72
    _ = v73
    _ = v74
    _ = v75
    _ = v76
    _ = v77
    _ = v78
}
