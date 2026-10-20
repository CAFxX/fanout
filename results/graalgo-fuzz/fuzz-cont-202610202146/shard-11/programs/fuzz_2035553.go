package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint }
type GenS4 struct { F0 uint; F1 uint64; F2 int }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 bool; F1 int32; F2 int }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint16; F1 string }
func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    v5 := []int{1}
    var v6 int64 = 42
    _ = p0
    _ = v5
    _ = v6
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 1
    v3 := []uint32{567, 607, 4294967295, 295}
    var v4 uint = 136
    if false {
        v7 := GenS1{F0: 0, F1: 2}
        for fi8 := 0; fi8 < 8; fi8++ {
            if false {
                fmt.Println(fi8, v4, v2)
                fmt.Println(v4, fi8)
            }
            switch fi8 {
            case 4294967295:
                switch 865 {
                case -2147483648:
                    var v9 uint64 = 18446744073709551615
                    _ = v9
                case 1:
                    var v10 float64 = 0.1
                    _ = v10
                default:
                    v2 += (v2 ^ (v2 | v2))
                }
            default:
                switch 1 {
                case -2147483648:
                    if true {
                        v4 = (^(v4))
                        v2 += v2
                        if (-698 > 352) {
                            for fi11 := 0; fi11 < 6; fi11++ {
                                v1[2] = (42 >= 1)
                                fmt.Println(v2, fi11)
                                _ = fi11
                            }
                        }
                    } else {
                        v12 := GenS2{F0: (fi8 ^ 32768), F1: 32767}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (!false) {
                            for fi13 := 0; fi13 < 14; fi13++ {
                                var v14 int = (fi8 << ((v4) & 63))
                                var v15 uint32 = 2147483647
                                v16 := GenS1{F0: (v15 << ((693) & 63)), F1: 2}
                                _ = fi13
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                        } else {
                            delete(v1, 2147483648)
                        }
                        _ = v12
                    }
                default:
                    v4 = v4
                }
            }
            v1[2] = false
            _ = fi8
        }
        _ = v7
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 int16 = 1
        for fi19 := 0; fi19 < 6; fi19++ {
            switch fi19 {
            case 128:
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if true {
                        v1[4294967295] = (420 > 4294967295)
                    } else {
                        var v20 int8 = 1
                        _ = v20
                    }
                    switch "foo bar" {
                    case "":
                        var v21 string = (("\n\t" + "\n\t") + "foo bar")
                        _ = v21
                    case "hello":
                        var v22 uint = 255
                        _ = v22
                    case "a":
                        switch "" {
                        case "\n\t":
                            v23 := v1[222]
                            _ = v23
                        default:
                            if len(v3) > 0 {
                                v3[0] = 42
                            }
                        }
                    }
                    fmt.Println(v4, v18)
                    if false {
                        v24 := GenS3{F0: 2147483647, F1: 140, F2: v4}
                        fmt.Println(v4, fi17)
                        if false {
                            _ = len(v1)
                        } else {
                            fmt.Println(v2, v18)
                        }
                        _ = v24
                    } else {
                        if (!true) {
                            fmt.Println(fi19, fi19, v18)
                        } else {
                            fmt.Println(v4)
                        }
                        var v25 uint = v4
                        var v26 uint64 = 42
                        var v27 uint32 = 2147483648
                        v3 = append(v3, 4294967295)
                        if (990 < 0) {
                            switch ("foo bar" + ("foo bar" + (("foo bar" + "") + "hello"))) {
                            case "\n\t":
                                delete(v1, v27)
                            case "":
                                v25 = v25
                            case "foo bar":
                                v28 := v1[(^(v27))]
                                _ = v28
                            default:
                                v18 -= -853
                            }
                            v29 := v1[v27]
                            _ = v29
                        } else {
                            fmt.Println(fi17)
                        }
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                }
            case 127:
                var v30 uint = 2
                _ = v30
            case 32767:
                switch "" {
                case "hello":
                    v2 += -1
                case "\n\t":
                    switch (fi17 * (fi17 - (fi19 - (fi17 - 859)))) {
                    case -115:
                        v4 += 577
                        if true {
                            fmt.Println(v2, v2, v2)
                        } else {
                            switch (("" + "hello") + "foo bar") {
                            case "\n\t":
                                var v31 int32 = int32(v18)
                                _ = v31
                            default:
                                var v32 int8 = 35
                                _ = v32
                            }
                            var v33 string = "hello"
                            _ = v33
                        }
                    }
                    v34 := GenS4{F0: v4, F1: 1, F2: (fi17 % (((fi17 >> ((v4) & 63))) | 1))}
                    _ = v34
                default:
                    v35 := GenS5{F0: v18, F1: 0}
                    _ = v35
                }
            }
            switch -214 {
            case 1:
                if true {
                    if (127 >= 1) {
                        if (true || true) {
                            switch (fi17 >> ((v4) & 63)) {
                            case 173:
                                v4 -= (v4 ^ (v4 * v4))
                            }
                            var v36 uint8 = 1
                            _ = v36
                        }
                        if (int16(v2) != (v18 * v18)) {
                            fmt.Println(fi19, fi19)
                            var v37 uint = (v4 + v4)
                            var v38 uint64 = 0
                            _ = v37
                            _ = v38
                        }
                        if ("\n\t" < "") {
                            v18 = 1
                            switch ("\n\t" + "hello") {
                            case "":
                                v2 -= 1
                            case "a":
                                if (false && false) {
                                    v39 := GenS6{F0: (-9223372036854775808 < v2), F1: 1, F2: (fi17 + (fi17 + -1))}
                                    v4 -= 65535
                                    v3 = append(v3, 2147483648)
                                    v40 := GenS2{F0: 127, F1: 0}
                                    _ = v39
                                    _ = v40
                                } else {
                                    if (1 != 2147483647) {
                                        var v41 int16 = 873
                                        _ = v41
                                    } else {
                                        v4 = v4
                                        var v42 int64 = v2
                                        v43 := GenS3{F0: 2147483647, F1: 101, F2: v4}
                                        _ = v42
                                        _ = v43
                                    }
                                }
                            case "foo bar":
                                v4 += v4
                            }
                        }
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(fi19)
                        case "foo bar":
                            switch (fi17 << (((v4 * v4)) & 63)) {
                            case 718:
                                if (fi19 < fi17) {
                                    switch (fi19 << ((2) & 63)) {
                                    case 2147483647:
                                        v1[745] = false
                                        v44 := GenS7{F0: 1, F1: 128}
                                        if true {
                                            switch (fi19 - -1) {
                                            case 551:
                                                v45 := GenS8{F0: 4294967295, F1: (0 < 0)}
                                                _ = v45
                                            default:
                                                v46 := GenS9{F0: "hello"}
                                                _ = v46
                                            }
                                            var v47 uint = (v4 ^ (v4 % ((881) | 1)))
                                            var v48 int32 = -867
                                            _ = v47
                                            _ = v48
                                        }
                                        _ = v44
                                    case 127:
                                        v49 := v1[0]
                                        _ = v49
                                    }
                                    v50 := GenS5{F0: 32767, F1: 42}
                                    var v51 string = ("foo bar" + "\n\t")
                                    _ = v50
                                    _ = v51
                                } else {
                                    var v52 int32 = -2147483648
                                    _ = v52
                                }
                            case -537:
                                v1[42] = (42 <= 131)
                                if true {
                                    v53 := GenS10{F0: 405}
                                    var v54 uint32 = 4294967295
                                    if false {
                                        if true {
                                            v2 -= v2
                                            v55 := GenS4{F0: 318, F1: uint64(v4), F2: fi17}
                                            _ = v55
                                        } else {
                                            if true {
                                                v18 += (-(v18))
                                                var v56 uint = v4
                                                _ = v56
                                            } else {
                                                v3 = append(v3, (v54 & (v54 << (((v4 - v4)) & 63))))
                                                var v57 int64 = v2
                                                var v58 bool = ("\n\t" == "hello")
                                                _ = v57
                                                _ = v58
                                            }
                                            switch "foo bar" {
                                            case "\n\t":
                                                if false {
                                                    var v59 uint64 = 0
                                                    _ = v59
                                                }
                                            }
                                        }
                                        v18 += v18
                                        var v60 int32 = 0
                                        v54 += 65535
                                        v61 := GenS5{F0: (-(v18)), F1: 2}
                                        fmt.Println(fi17)
                                        delete(v1, (v54 & 0))
                                        _ = v60
                                        _ = v61
                                    } else {
                                        v62 := v1[v54]
                                        v63 := v1[(v54 / ((2) | 1))]
                                        _ = v62
                                        _ = v63
                                    }
                                    _ = v53
                                    _ = v54
                                }
                                fmt.Println(fi17, v18, fi17)
                                var v64 bool = true
                                _ = v64
                            case -9223372036854775808:
                                if (!false) {
                                    if false {
                                        fmt.Println(fi19, v4)
                                    } else {
                                        var v65 float32 = float32(1.5)
                                        _ = v65
                                    }
                                    if false {
                                        var v66 int32 = 10
                                        var v67 uint64 = 42
                                        var v68 uint32 = 4294967295
                                        _ = v66
                                        _ = v67
                                        _ = v68
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 164
                                    }
                                    var v69 string = "\n\t"
                                    _ = v69
                                }
                            default:
                                v70 := GenS5{F0: (v18 / ((v18) | 1)), F1: 988}
                                _ = v70
                            }
                        default:
                            var v71 float64 = 1e10
                            _ = v71
                        }
                    }
                    switch 128 {
                    case -1:
                        v72 := GenS5{F0: v18, F1: 0}
                        _ = v72
                    }
                    if (false || false) {
                        if true {
                            v73 := GenS11{F0: fi17}
                            _ = v73
                        }
                        var v74 uint16 = 65535
                        fmt.Println(v74, v74, v18)
                        _ = v74
                    }
                    var v75 float32 = (1e10 + (0.1 / 0.1))
                    v2 -= (v2 & v2)
                    _ = v75
                } else {
                    v76 := GenS1{F0: 2147483648, F1: 65535}
                    v18 = (v18 >> ((42) & 63))
                    _ = v76
                }
            }
            _ = fi19
        }
        var v77 int16 = (^(v18))
        _ = fi17
        _ = v18
        _ = v77
    }
    if true {
        fmt.Println(v2)
        var v78 string = ("foo bar" + (("hello" + ("hello" + ("foo bar" + ("hello" + (("foo bar" + "foo bar") + "hello"))))) + ""))
        delete(v1, uint32(v4))
        _ = v78
    } else {
        if false {
            if (127 == -1) {
                var v79 int16 = 0
                _ = v79
            }
        }
        var v80 uint8 = 127
        _ = v80
    }
    fmt.Println(v2, v2)
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    if true {
        v2 = -194
    }
    v81 := GenS5{F0: 1000, F1: 65535}
    if (false || false) {
        v3 = append(v3, 2147483648)
        v4 = (v4 << ((v4) & 63))
    }
    for fi82 := 0; fi82 < 4; fi82++ {
        delete(v1, 1)
        var v83 uint16 = 35
        _ = fi82
        _ = v83
    }
    for fi84 := 0; fi84 < 16; fi84++ {
        fmt.Println(v4)
        var v85 uint64 = 18446744073709551615
        _ = fi84
        _ = v85
    }
    var v86 int8 = 1
    v87 := GenS12{F0: 84, F1: (("foo bar" + "\n\t") + "foo bar")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v81
    _ = v86
    _ = v87
}
