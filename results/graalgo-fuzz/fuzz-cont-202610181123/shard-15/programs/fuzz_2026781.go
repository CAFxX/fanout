package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint32; F1 int; F2 int8 }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int32; F1 string; F2 uint8 }
func genFn0(p0 uint64) uint16 {
    if (false || true) {
        v4 := []int{-1, -162, -545}
        var v5 bool = (true && (1 != 65535))
        _ = v4
        _ = v5
    }
    p0 = p0
    v6 := []int{}
    var v7 int = -593
    var v8 string = ("foo bar" + "hello")
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 50
}
func genFn1(p0 uint8, p1 uint32, p2 bool) uint32 {
    for fi9 := 0; fi9 < 6; fi9++ {
        if (!p2) {
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 string = "hello"
                var v12 int = fi9
                _ = fi10
                _ = v11
                _ = v12
            }
            p2 = p2
            if (!true) {
                for fi13 := 0; fi13 < 3; fi13++ {
                    switch fi9 {
                    case 1:
                        fmt.Println(fi9)
                    case 4294967295:
                        v14 := []int{726, (^(fi9)), 42}
                        _ = v14
                    }
                    v15 := []int{(fi13 ^ -2147483648), (fi13 ^ 42), (fi9 / ((fi9) | 1)), 32767}
                    _ = fi13
                    _ = v15
                }
            }
        } else {
            if (p1 != p1) {
                if p2 {
                    if false {
                        if (-128 > 127) {
                            v16 := []int{}
                            p1 = 510
                            for fi17 := 0; fi17 < 8; fi17++ {
                                v16 = append(v16, (fi9 >> ((1) & 63)))
                                fmt.Println(p1, fi17)
                                var v18 uint32 = p1
                                p0 += (p0 << ((18446744073709551615) & 63))
                                _ = fi17
                                _ = v18
                            }
                            _ = v16
                        } else {
                            if (1 <= 0) {
                                if p2 {
                                    v19 := []int{fi9}
                                    fmt.Println(p1, p2)
                                    p2 = (!false)
                                    _ = v19
                                }
                                p0 = p0
                                v20 := GenS1{F0: (fi9 & -73)}
                                _ = v20
                            }
                            fmt.Println(fi9, p0, fi9)
                            if p2 {
                                var v21 int8 = 42
                                v22 := GenS1{F0: (fi9 >> ((684) & 63))}
                                _ = v21
                                _ = v22
                            }
                        }
                        p2 = p2
                    } else {
                        var v23 int64 = 9223372036854775807
                        fmt.Println(p2, fi9, p2)
                        var v24 uint16 = genFn0(18446744073709551615)
                        _ = v23
                        _ = v24
                    }
                }
            } else {
                if p2 {
                    v25 := GenS1{F0: (-(fi9))}
                    var v26 int = fi9
                    _ = v25
                    _ = v26
                } else {
                    p1 -= 0
                    p1 += (^(p1))
                }
            }
            v27 := []int{9223372036854775807, fi9}
            _ = v27
        }
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 string = "foo bar"
            _ = fi28
            _ = v29
        }
        var v30 int8 = 1
        _ = fi9
        _ = v30
    }
    var v31 uint8 = (p0 / ((p0) | 1))
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    return (p1 * (p1 << ((4294967295) & 63)))
}

func main() {
    var v1 int16 = 1
    var v2 uint32 = 2147483647
    v3 := make(map[string]uint32)
    var v32 uint = 42
    if true {
        v33 := []int{-732, -207}
        v34 := GenS1{F0: -9223372036854775808}
        if true {
            switch int(v32) {
            case 1:
                _ = len(v33)
            default:
                v35 := GenS1{F0: 589}
                _ = v35
            }
            if false {
                var v36 string = "hello"
                v36 = v36
                _ = v36
            }
            v37 := GenS1{F0: 42}
            _ = v37
        } else {
            switch "a" {
            case "hello":
                for fi38 := 0; fi38 < 2; fi38++ {
                    v39 := GenS1{F0: (fi38 / (((fi38 | (fi38 & -1))) | 1))}
                    var v40 float32 = ((1e10 - 0.1) / 1.5)
                    if (-806 == -270) {
                        if false {
                            fmt.Println(v32, v1)
                            fmt.Println(v40, v32, v40)
                        }
                    } else {
                        var v41 float32 = v40
                        _ = v41
                    }
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            case "foo bar":
                if true {
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                    v42 := GenS1{F0: 42}
                    _ = v42
                } else {
                    fmt.Println(v32, v32)
                    if len(v33) > 0 {
                        v33[0] = 32767
                    }
                    for fi43 := 0; fi43 < 15; fi43++ {
                        var v44 int8 = 4
                        switch -2147483648 {
                        case 4294967295:
                            for fi45 := 0; fi45 < 1; fi45++ {
                                _ = len(v33)
                                if true {
                                    v1 += -770
                                    var v46 uint16 = 65535
                                    fmt.Println(v32, v46, v44)
                                    _ = v46
                                } else {
                                    if (v44 != v44) {
                                        var v47 uint16 = genFn0(18446744073709551615)
                                        if (1 != 219) {
                                            v47 = v47
                                            v47 = (v47 & 65535)
                                        } else {
                                            v48 := GenS1{F0: -552}
                                            _ = v48
                                        }
                                        v49 := GenS2{F0: genFn1(2, uint32(v32), true), F1: ""}
                                        if (255 < (v32 % ((65535) | 1))) {
                                            var v50 uint16 = (v47 << ((18446744073709551615) & 63))
                                            _ = len(v3)
                                            _ = v50
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                v51 := v3[""]
                                                _ = v51
                                            case "foo bar":
                                                var v52 uint32 = (v2 / ((2) | 1))
                                                _ = v52
                                            default:
                                                fmt.Println(v44)
                                            }
                                        }
                                        _ = v47
                                        _ = v49
                                    } else {
                                        if (!(127 < 0)) {
                                            if false {
                                                v53 := GenS3{F0: v1, F1: "foo bar"}
                                                _ = v53
                                            }
                                            v54 := GenS4{F0: 1, F1: (fi45 / ((fi43) | 1)), F2: v44}
                                            v2 = v2
                                            _ = v54
                                        }
                                        v55 := v3[("hello" + "a")]
                                        var v56 uint16 = uint16(v32)
                                        _ = v55
                                        _ = v56
                                    }
                                }
                                v1 = int16(fi43)
                                v57 := GenS3{F0: 32767, F1: "hello"}
                                v58 := GenS2{F0: v2, F1: "hello"}
                                _ = fi45
                                _ = v57
                                _ = v58
                            }
                        case -171:
                            _ = len(v3)
                        default:
                            v59 := GenS3{F0: v1, F1: "foo bar"}
                            _ = v59
                        }
                        switch fi43 {
                        case 4294967295:
                            var v60 uint = (v32 - v32)
                            _ = v60
                        case 217:
                            if true {
                                for fi61 := 0; fi61 < 14; fi61++ {
                                    v44 = 0
                                    var v62 int = (fi43 - (fi43 >> ((v32) & 63)))
                                    var v63 string = "foo bar"
                                    var v64 uint16 = 2
                                    var v65 uint8 = 1
                                    _ = fi61
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                                v1 = (v1 / ((-288) | 1))
                                var v66 bool = true
                                _ = v66
                            } else {
                                if true {
                                    fmt.Println(v1, v1)
                                    fmt.Println(v1, v2)
                                    var v67 int16 = (-(v1))
                                    _ = v67
                                }
                                v32 -= (v32 / ((v32) | 1))
                                v1 -= v1
                            }
                        default:
                            switch "foo bar" {
                            case "a":
                                var v68 int = -9223372036854775808
                                _ = v68
                            case "\n\t":
                                var v69 uint32 = v2
                                _ = v69
                            case "":
                                fmt.Println(v2, v2)
                            }
                        }
                        _ = fi43
                        _ = v44
                    }
                }
            default:
                fmt.Println(v1)
            }
        }
        _ = v33
        _ = v34
    } else {
        v2 -= (v2 << (((^(v32))) & 63))
    }
    for fi70 := 0; fi70 < 17; fi70++ {
        v71 := GenS2{F0: 1, F1: ""}
        _ = fi70
        _ = v71
    }
    var v72 bool = false
    for fi73 := 0; fi73 < 4; fi73++ {
        if (!v72) {
            if false {
                v2 -= v2
                var v74 float64 = 0.1
                _ = v74
            }
            if v72 {
                v32 -= 0
                var v75 uint32 = v2
                for fi76 := 0; fi76 < 15; fi76++ {
                    v77 := v3[("foo bar" + "foo bar")]
                    _ = fi76
                    _ = v77
                }
                _ = v75
            }
        }
        var v78 uint64 = 1
        _ = fi73
        _ = v78
    }
    v1 -= 0
    fmt.Println(v1, v32)
    if true {
        for fi79 := 0; fi79 < 1; fi79++ {
            switch "\n\t" {
            case "":
                _ = len(v3)
            default:
                if (1 < 983) {
                    for fi80 := 0; fi80 < 4; fi80++ {
                        fmt.Println(v2)
                        var v81 uint64 = 770
                        _ = fi80
                        _ = v81
                    }
                    switch (fi79 / (((-(fi79))) | 1)) {
                    case 127:
                        var v82 uint = v32
                        _ = v82
                    case 32767:
                        var v83 int8 = 9
                        _ = v83
                    case -121:
                        for fi84 := 0; fi84 < 7; fi84++ {
                            v32 = v32
                            v32 += (^(v32))
                            v85 := GenS5{F0: v1, F1: 0}
                            v86 := GenS6{F0: (v1 % ((0) | 1)), F1: -100}
                            var v87 string = "foo bar"
                            if (601 > 18446744073709551615) {
                                v72 = ((^(v32)) > 415)
                                v88 := []int{fi79}
                                v2 -= genFn1(1, (v2 ^ (^(v2))), v72)
                                var v89 int16 = v1
                                var v90 uint32 = 2147483648
                                v91 := GenS2{F0: 607, F1: "a"}
                                _ = v88
                                _ = v89
                                _ = v90
                                _ = v91
                            } else {
                                v92 := []int{fi84, fi84}
                                var v93 int16 = v1
                                _ = len(v92)
                                _ = v92
                                _ = v93
                            }
                            _ = fi84
                            _ = v85
                            _ = v86
                            _ = v87
                        }
                    default:
                        switch 2147483647 {
                        case -9223372036854775808:
                            v94 := GenS7{F0: 42}
                            _ = v94
                        default:
                            for fi95 := 0; fi95 < 2; fi95++ {
                                var v96 float64 = 1.5
                                v97 := []int{fi95, (fi79 << ((674) & 63)), fi79, fi95}
                                fmt.Println(fi95, fi95)
                                v96 += (1e10 / v96)
                                switch (fi95 - 127) {
                                case -442:
                                    switch ("foo bar" + ("" + "\n\t")) {
                                    case "hello":
                                        _ = len(v3)
                                    case "foo bar":
                                        v98 := GenS8{F0: 108}
                                        _ = v98
                                    case "":
                                        var v99 uint32 = v2
                                        _ = v99
                                    }
                                case 508:
                                    v32 = 255
                                }
                                if true {
                                    var v100 uint = v32
                                    v100 -= v32
                                    v101 := GenS9{F0: 1, F1: ("foo bar" + "foo bar"), F2: uint8(v32)}
                                    v102 := GenS1{F0: (fi79 | 32768)}
                                    var v103 int8 = -108
                                    _ = v100
                                    _ = v101
                                    _ = v102
                                    _ = v103
                                }
                                if (v72 && ((^(fi95)) == (fi95 - 0))) {
                                    if (!false) {
                                        if (!(18446744073709551615 == 256)) {
                                            fmt.Println(v1, fi79, v96)
                                        }
                                    }
                                    var v104 int = 127
                                    v105 := GenS6{F0: 1, F1: -116}
                                    _ = v104
                                    _ = v105
                                } else {
                                    _ = len(v3)
                                    fmt.Println(fi95, fi95, v1)
                                    if len(v97) >= 2 {
                                        v97 = v97[1:len(v97)]
                                    }
                                    if len(v97) > 0 {
                                        v97[0] = (fi95 ^ 1)
                                    }
                                }
                                var v106 float32 = 0.0
                                _ = fi95
                                _ = v106
                                _ = v96
                                _ = v97
                            }
                        }
                    }
                    var v107 int32 = 715
                    _ = v107
                }
            }
            v1 = (v1 / (((v1 - v1)) | 1))
            v108 := []int{fi79, (fi79 ^ (fi79 + fi79))}
            _ = fi79
            _ = v108
        }
        var v109 uint = 145
        _ = v109
    }
    var v110 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v110
    _ = v2
    _ = v3
    _ = v32
    _ = v72
}
