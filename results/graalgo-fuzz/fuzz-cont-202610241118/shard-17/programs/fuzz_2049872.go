package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int16; F1 int32; F2 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64; F1 uint8; F2 int }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int64 }
func genFn0(p0 uint8, p1 bool, p2 bool) int64 {
    var v4 bool = false
    v5 := GenS1{F0: 4294967295}
    for fi6 := 0; fi6 < 1; fi6++ {
        if p2 {
            v7 := GenS1{F0: 0}
            _ = v7
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -266
}
func genFn1(p0 int16, p1 int) int8 {
    var v8 int64 = -9223372036854775808
    v9 := []int{}
    var v10 uint8 = 42
    p1 = p1
    var v11 bool = (false && false)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return -107
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    switch 8 {
    case 128:
        switch ("foo bar" + "") {
        case "":
            v12 := GenS2{F0: -2147483648, F1: genFn1(-32768, 2147483647)}
            _ = v12
        }
    default:
        if ((((2 != 1) && true) && (false && false)) || false) {
            var v13 uint16 = 42
            _ = v13
        } else {
            var v14 uint32 = 978
            for fi15 := 0; fi15 < 9; fi15++ {
                if (!true) {
                    delete(v2, 2147483648)
                    var v16 uint8 = 128
                    if true {
                        var v17 uint = 4294967295
                        v18 := GenS3{F0: -128, F1: (v14 % (((^(v14))) | 1))}
                        _ = v17
                        _ = v18
                    }
                    _ = v16
                } else {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        v20 := GenS4{F0: 655, F1: 674, F2: 32767}
                        if (32767 <= 65535) {
                            fmt.Println(fi19, fi15)
                            fmt.Println(fi15)
                        } else {
                            var v21 string = ""
                            v14 += (^(v14))
                            delete(v3, v14)
                            _ = v21
                        }
                        if (float32(1e10) > (((-2.25 - float32(float64(1e10))) - 1e10) + 3.14159)) {
                            var v22 uint16 = 32767
                            fmt.Println(fi19, fi15, v22)
                            _ = v22
                        }
                        var v23 bool = false
                        var v24 uint = 18446744073709551615
                        _ = fi19
                        _ = v20
                        _ = v23
                        _ = v24
                    }
                    var v25 int64 = -1
                    _ = v25
                }
                var v26 float32 = 1.5
                v27 := []int{}
                _ = fi15
                _ = v26
                _ = v27
            }
            var v28 uint32 = v14
            _ = v14
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 uint64 = 18446744073709551615
        if false {
            if true {
                v31 := []int{fi29, (^(fi29)), fi29}
                _ = v31
            }
            var v32 uint16 = 2
            if (255 >= 65535) {
                v30 = v30
                var v33 float32 = 3.14159
                _ = v33
            } else {
                if false {
                    for fi34 := 0; fi34 < 14; fi34++ {
                        v32 += (^(v32))
                        var v35 int16 = 654
                        v36 := GenS3{F0: 1, F1: 181}
                        delete(v3, 2147483647)
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    v37 := GenS2{F0: 2147483647, F1: 0}
                    var v38 uint64 = 824
                    _ = v37
                    _ = v38
                } else {
                    if (1 <= 0) {
                        switch "\n\t" {
                        case "hello":
                            var v39 uint16 = (v32 >> ((42) & 63))
                            _ = v39
                        case "":
                            for fi40 := 0; fi40 < 9; fi40++ {
                                var v41 float64 = -2.25
                                if false {
                                    if true {
                                        var v42 int = fi40
                                        _ = v42
                                    } else {
                                        var v43 int16 = 184
                                        _ = v43
                                    }
                                    var v44 string = ("" + ("\n\t" + "a"))
                                    _ = v44
                                } else {
                                    switch ("\n\t" + "\n\t") {
                                    case "hello":
                                        v45 := []int{(fi40 >> ((1) & 63)), fi40, fi40}
                                        _ = v45
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "hello":
                                            v46 := []int{-749, -2147483648, (^(fi29)), fi40}
                                            _ = v46
                                        case "\n\t":
                                            v1[fi40] = 65535
                                        case "foo bar":
                                            _ = len(v2)
                                        default:
                                            delete(v2, uint32(v30))
                                        }
                                    }
                                    if ((-2.25 - (float64(-2.25) - (v41 * float64(3.14159)))) <= 3.14159) {
                                        switch fi29 {
                                        case 2147483647:
                                            v32 += v32
                                        case 164:
                                            var v47 uint8 = 2
                                            _ = v47
                                        default:
                                            var v48 int8 = -79
                                            _ = v48
                                        }
                                    }
                                }
                                _ = fi40
                                _ = v41
                            }
                        case "\n\t":
                            switch ("foo bar" + ("\n\t" + "a")) {
                            case "hello":
                                v30 += v30
                            case "foo bar":
                                var v49 bool = (!false)
                                _ = v49
                            }
                        default:
                            var v50 float64 = float64(-2.25)
                            _ = v50
                        }
                        fmt.Println(v30, fi29, v30)
                    } else {
                        if (false || true) {
                            var v51 uint64 = v30
                            for fi52 := 0; fi52 < 15; fi52++ {
                                v51 = v51
                                var v53 int8 = 0
                                fmt.Println(v51)
                                v54 := GenS1{F0: 1}
                                fmt.Println(v51, fi29, fi29)
                                delete(v3, 4294967295)
                                switch "\n\t" {
                                case "a":
                                    var v55 uint16 = v32
                                    _ = v55
                                case "":
                                    fmt.Println(v53, v30, v51)
                                }
                                v30 = v51
                                _ = fi52
                                _ = v53
                                _ = v54
                            }
                            _ = v51
                        }
                        v32 += 1
                    }
                }
            }
            _ = v32
        }
        v56 := []int{(fi29 + 158), -2147483648, 42}
        _ = fi29
        _ = v30
        _ = v56
    }
    _ = len(v2)
    switch "\n\t" {
    case "":
        for fi57 := 0; fi57 < 4; fi57++ {
            v58 := GenS1{F0: 776}
            _ = fi57
            _ = v58
        }
    default:
        v59 := []int{}
        _ = v59
    }
    v60 := GenS4{F0: -32768, F1: -2147483648, F2: 127}
    var v61 uint = 0
    for fi62 := 0; fi62 < 19; fi62++ {
        fmt.Println(v61, v61)
        if false {
            if (!true) {
                switch "\n\t" {
                case "a":
                    v61 = v61
                case "hello":
                    v61 += v61
                }
                v61 += v61
                v61 += v61
            } else {
                v63 := GenS2{F0: 42, F1: 79}
                v64 := GenS5{F0: -2147483648}
                v61 = 65535
                _ = v63
                _ = v64
            }
            v65 := GenS6{F0: (v61 % ((v61) | 1)), F1: 2147483647, F2: 255}
            v66 := GenS2{F0: 0, F1: 1}
            _ = v65
            _ = v66
        } else {
            var v67 string = ""
            var v68 uint16 = 65535
            v61 = v61
            _ = v67
            _ = v68
        }
        if true {
            switch 32767 {
            case 2147483647:
                for fi69 := 0; fi69 < 16; fi69++ {
                    v70 := []int{}
                    if true {
                        switch fi62 {
                        case 128:
                            var v71 int8 = genFn1(1, fi62)
                            v72 := GenS2{F0: 1, F1: 42}
                            _ = v71
                            _ = v72
                        }
                        delete(v1, (fi69 / ((128) | 1)))
                    }
                    if ("" == "hello") {
                        v61 += v61
                        v73 := v3[4294967295]
                        v74 := GenS1{F0: (v61 << ((259) & 63))}
                        var v75 int16 = 770
                        _ = v73
                        _ = v74
                        _ = v75
                    } else {
                        switch "hello" {
                        case "\n\t":
                            v61 -= 379
                            if (1 >= 2147483647) {
                                v76 := GenS1{F0: (^(v61))}
                                _ = v76
                            } else {
                                var v77 int = fi69
                                v78 := GenS7{F0: v77}
                                v79 := GenS8{F0: genFn0(134, ((true && false) && false), ((false || (42 >= 488)) && (0 > 4294967295))), F1: 1, F2: (v77 * 325)}
                                if len(v70) > 0 {
                                    v70[0] = (-(v77))
                                }
                                v61 -= v61
                                v61 -= (^(v61))
                                v77 = fi62
                                _ = v77
                                _ = v78
                                _ = v79
                            }
                            v61 += (^(v61))
                        case "foo bar":
                            v61 = 4294967295
                        default:
                            delete(v3, 0)
                        }
                    }
                    _ = fi69
                    _ = v70
                }
            default:
                var v80 float64 = 0.1
                _ = v80
            }
            v61 = 255
            v3[2147483648] = 42
        } else {
            var v81 float32 = (3.14159 + (3.14159 * 0.0))
            var v82 float32 = v81
            for fi83 := 0; fi83 < 14; fi83++ {
                switch (fi83 & fi62) {
                case 1:
                    v84 := GenS9{F0: fi62}
                    _ = v84
                }
                _ = fi83
            }
            _ = v81
            _ = v82
        }
        _ = fi62
    }
    var v85 float32 = 1e10
    v86 := GenS10{F0: 42}
    var v87 int = -236
    v88 := GenS4{F0: -605, F1: -886, F2: (v87 * v87)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v61:", v61)
    fmt.Println("v85:", v85)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v60
    _ = v61
    _ = v85
    _ = v86
    _ = v87
    _ = v88
}
