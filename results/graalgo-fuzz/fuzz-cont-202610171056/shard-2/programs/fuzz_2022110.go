package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 int32; F1 uint64; F2 string }
type GenS9 struct { F0 uint }
func genFn0(p0 int16) uint16 {
    v5 := []int{902, 2147483647}
    v6 := GenS1{F0: 17, F1: "\n\t", F2: 9223372036854775807}
    p0 = (p0 | p0)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    p0 += (p0 + p0)
    _ = p0
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 int16, p1 uint, p2 int) string {
    v7 := []int{2147483647, (-(p2))}
    var v8 float64 = 0.1
    switch "hello" {
    case "\n\t":
        var v9 int8 = 100
        _ = v9
    case "foo bar":
        v7 = append(v7, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ("foo bar" + "\n\t")
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{}
    var v3 int64 = 457
    v4 := []int{32767, 9223372036854775807, -568}
    switch "hello" {
    case "a":
        v3 = (v3 - 9223372036854775807)
    case "hello":
        if false {
            v10 := GenS1{F0: 127, F1: "a", F2: (v3 >> ((871) & 63))}
            _ = v10
        }
    default:
        if false {
            v3 = v3
            v3 = v3
        } else {
            fmt.Println(v3)
            v3 = -1
        }
    }
    v11 := GenS2{F0: 213, F1: 87}
    var v12 int8 = 42
    v3 += -819
    var v13 string = "foo bar"
    if (!(!(3.14159 == (((0.0 + float64(float32(float64(((0.0 * (float32(0.1) / (3.14159 * 1.5))) * 1.5))))) - (3.14159 * (0.1 - 1e10))) - -2.25)))) {
        v14 := GenS3{F0: -1, F1: 2}
        _ = v14
    } else {
        if (!((1 != 512) || true)) {
            v15 := GenS4{F0: 1}
            _ = v15
        }
        var v16 int8 = v12
        _ = len(v1)
        _ = v16
    }
    fmt.Println(v12)
    for fi17 := 0; fi17 < 20; fi17++ {
        if false {
            v13 = ((((v13 + v13) + "hello") + "foo bar") + v13)
            var v18 int32 = 2147483647
            _ = v18
        } else {
            var v19 float64 = (1e10 * (0.0 + 0.0))
            v12 = v12
            for fi20 := 0; fi20 < 9; fi20++ {
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                v19 -= (float64((((((((float32(float64(-2.25)) * (0.0 - -2.25)) + 0.1) / 1e10) * 1.5) + 0.1) - (3.14159 + (0.0 * 1.5))) * 1e10)) / (0.1 - v19))
                _ = fi20
            }
            _ = v19
        }
        _ = fi17
    }
    var v21 int8 = int8(v3)
    fmt.Println(v21, v13)
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 2; fi23++ {
            v24 := GenS5{F0: ((v13 + ((v13 + v13) + v13)) + (v13 + ("hello" + v13))), F1: 1}
            fmt.Println(v21)
            if (!false) {
                if false {
                    if (!false) {
                        v13 = "\n\t"
                        var v25 uint32 = 65535
                        _ = len(v1)
                        _ = v25
                    }
                    fmt.Println(fi23)
                } else {
                    v13 = ""
                    delete(v1, 400)
                }
                switch 618 {
                case 567:
                    fmt.Println(fi22, v3)
                case 32767:
                    var v26 string = v13
                    _ = v26
                default:
                    var v27 float32 = 1.5
                    if (!false) {
                        v3 -= 9223372036854775807
                        var v28 bool = false
                        var v29 float32 = 0.0
                        _ = v28
                        _ = v29
                    } else {
                        v27 -= -2.25
                        var v30 uint64 = 18446744073709551615
                        switch (v13 + (v13 + "foo bar")) {
                        case "a":
                            fmt.Println(fi23, v21, v21)
                        case "":
                            var v31 bool = false
                            _ = v31
                        default:
                            fmt.Println(v27)
                            v32 := GenS6{F0: 99, F1: "hello"}
                            _ = v32
                        }
                        _ = v30
                    }
                    _ = v27
                }
                if true {
                    var v33 uint16 = 2
                    v21 -= v12
                    _ = v33
                }
            } else {
                if false {
                    v12 = (v21 << ((18446744073709551615) & 63))
                } else {
                    var v34 bool = (!(!(0 != 1)))
                    _ = v34
                }
                v35 := GenS7{F0: (^(v3)), F1: 0}
                if (3.14159 >= 1.5) {
                    v3 -= (v3 | v3)
                    var v36 int16 = 42
                    v36 -= v36
                    var v37 uint64 = 18446744073709551615
                    _ = len(v4)
                    _ = v36
                    _ = v37
                } else {
                    if false {
                        v12 = -128
                        if true {
                            fmt.Println(v12)
                            var v38 bool = (42 != 151)
                            v13 = v13
                            var v39 int = 128
                            _ = v38
                            _ = v39
                        } else {
                            _ = len(v1)
                            if (!false) {
                                if (!false) {
                                    v21 = 100
                                    if true {
                                        fmt.Println(fi22, v13)
                                    } else {
                                        var v40 string = "foo bar"
                                        _ = v40
                                    }
                                    fmt.Println(fi22)
                                } else {
                                    var v41 uint8 = 255
                                    var v42 int16 = 0
                                    _ = v41
                                    _ = v42
                                }
                                var v43 uint = 802
                                switch fi22 {
                                case 127:
                                    v44 := GenS7{F0: (v3 % ((v3) | 1)), F1: (v12 & v21)}
                                    _ = v44
                                }
                                _ = v43
                            } else {
                                if true {
                                    var v45 float32 = (((1e10 * ((1.5 + (3.14159 * 0.1)) + (0.0 + 3.14159))) / 1.0) + -2.25)
                                    fmt.Println(fi23, fi22)
                                    _ = v45
                                }
                                v12 = v12
                                if true {
                                    v46 := GenS8{F0: 2147483647, F1: 635, F2: "hello"}
                                    _ = v46
                                } else {
                                    if (((1.5 - ((0.0 + (((0.1 * float32(-2.25)) / 1.0) / -2.25)) * (float32(float64(-2.25)) * float32(3.14159)))) - 3.14159) != 1.5) {
                                        fmt.Println(v13)
                                        switch 4294967295 {
                                        case 32768:
                                            v47 := GenS2{F0: 65535, F1: 128}
                                            _ = v47
                                        case -9223372036854775808:
                                            if true {
                                                var v48 uint8 = 128
                                                var v49 int = (fi23 - fi22)
                                                _ = v48
                                                _ = v49
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        }
                                        v50 := GenS9{F0: 42}
                                        _ = v50
                                    } else {
                                        switch fi23 {
                                        case -115:
                                            if ((false && true) && (0.0 >= 0.0)) {
                                                var v51 int32 = 905
                                                if true {
                                                    v21 += v21
                                                    v21 = v12
                                                    delete(v1, 899)
                                                } else {
                                                    if ((-645 == v3) || false) {
                                                        _ = len(v1)
                                                        var v52 bool = (!(128 == 128))
                                                        _ = v52
                                                    } else {
                                                        v12 = 1
                                                        switch (fi23 & fi22) {
                                                        case 1:
                                                            v21 = v21
                                                        case -1:
                                                            var v53 uint32 = 0
                                                            _ = v53
                                                        }
                                                        var v54 string = (v13 + "a")
                                                        _ = v54
                                                    }
                                                    v13 = v13
                                                    fmt.Println(fi22, fi22)
                                                }
                                                _ = v51
                                            } else {
                                                var v55 int16 = -1
                                                _ = v55
                                            }
                                        default:
                                            v56 := GenS3{F0: 0, F1: 373}
                                            _ = v56
                                        }
                                    }
                                }
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v12 -= v21
                        }
                    }
                    if (v12 >= v12) {
                        fmt.Println(v13)
                        var v57 int = (fi23 | 127)
                        var v58 int16 = 1000
                        _ = v57
                        _ = v58
                    }
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                    var v59 uint8 = 127
                    v60 := GenS1{F0: (v21 << ((18446744073709551615) & 63)), F1: "hello", F2: v3}
                    _ = v59
                    _ = v60
                }
                _ = v35
            }
            _ = fi23
            _ = v24
        }
        v61 := GenS7{F0: v3, F1: -45}
        v1[42] = "foo bar"
        _ = fi22
        _ = v61
    }
    var v62 uint8 = 39
    if false {
        v63 := GenS7{F0: (v3 << ((871) & 63)), F1: v21}
        for fi64 := 0; fi64 < 1; fi64++ {
            if false {
                fmt.Println(v21)
                if ((float64(0.0) * (-2.25 * 3.14159)) != 1e10) {
                    if ((v3 + v3) <= v3) {
                        for fi65 := 0; fi65 < 9; fi65++ {
                            if false {
                                var v66 int16 = 0
                                _ = v66
                            }
                            v21 += v21
                            _ = fi65
                        }
                        fmt.Println(v13)
                    }
                } else {
                    for fi67 := 0; fi67 < 3; fi67++ {
                        v68 := GenS9{F0: 255}
                        fmt.Println(v62)
                        fmt.Println(v62)
                        _ = fi67
                        _ = v68
                    }
                    v13 = v13
                }
                v1[1] = (((v13 + v13) + v13) + ("foo bar" + "a"))
            } else {
                v62 -= v62
                fmt.Println(v21, v62)
            }
            var v69 uint8 = (v62 << ((261) & 63))
            _ = fi64
            _ = v69
        }
        _ = v63
    } else {
        var v70 int32 = -2147483648
        _ = v70
    }
    if (!false) {
        var v71 int = -9223372036854775808
        _ = v71
    } else {
        var v72 float64 = 1.5
        _ = v72
    }
    switch 0 {
    case -1:
        var v73 int16 = int16(v3)
        _ = v73
    case 128:
        var v74 int64 = (v3 % ((v3) | 1))
        _ = v74
    case 127:
        v13 = (((v13 + v13) + (("hello" + ("\n\t" + "a")) + (v13 + v13))) + ("\n\t" + v13))
    }
    if (v62 < 127) {
        fmt.Println(v21)
        v13 = ("hello" + (v13 + ("\n\t" + "foo bar")))
    } else {
        v4 = append(v4, -310)
        v13 = (v13 + (v13 + ((v13 + (v13 + (v13 + v13))) + v13)))
    }
    v12 -= (v12 + (v12 % ((42) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v62
}
