package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int32 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 bool; F1 int16; F2 bool }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint64; F1 uint16; F2 string }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool; F1 int }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS12 struct { F0 int16; F1 uint16 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 bool; F1 bool; F2 bool }
func genFn0(p0 uint16, p1 int16) int8 {
    fmt.Println(p0, p1)
    v4 := []int{0, 32767, 2147483647}
    v5 := GenS1{F0: 723, F1: p0, F2: -529}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 uint = 479
    var v2 uint32 = 618
    v3 := make(map[uint32]uint32)
    v1 -= (v1 << (((v1 ^ (v1 - uint(v2)))) & 63))
    switch 4294967295 {
    case 1:
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v2, v1)
        case "a":
            for fi6 := 0; fi6 < 20; fi6++ {
                var v7 int8 = 42
                switch (fi6 - 809) {
                case 42:
                    if false {
                        if true {
                            v8 := []int{42, fi6}
                            v2 = 42
                            fmt.Println(fi6, v2, v7)
                            _ = v8
                        }
                        if (0.1 == 0.1) {
                            for fi9 := 0; fi9 < 2; fi9++ {
                                v7 += genFn0(902, 1000)
                                _ = fi9
                            }
                            switch (fi6 / ((1) | 1)) {
                            case -1:
                                v10 := GenS2{F0: (fi6 * (fi6 % ((256) | 1))), F1: 42}
                                _ = v10
                            default:
                                var v11 int = -966
                                _ = v11
                            }
                        }
                    } else {
                        for fi12 := 0; fi12 < 15; fi12++ {
                            if false {
                                switch "a" {
                                case "":
                                    v2 -= v2
                                case "hello":
                                    v13 := GenS3{F0: (!(!true)), F1: 0, F2: (false || false)}
                                    var v14 uint8 = 0
                                    _ = v13
                                    _ = v14
                                default:
                                    v7 -= (v7 | 0)
                                }
                                var v15 float32 = 1e10
                                if (!true) {
                                    if (((!(247 != 2)) || ((v2 * (v2 & (v2 & v2))) <= v2)) || true) {
                                        v15 = 3.14159
                                        if false {
                                            fmt.Println(v1)
                                        }
                                        v15 = v15
                                    }
                                    switch fi6 {
                                    case -821:
                                        v16 := GenS1{F0: (v1 | v1), F1: 113, F2: -574}
                                        _ = v16
                                    case 128:
                                        _ = len(v3)
                                        if false {
                                            var v17 uint16 = 387
                                            var v18 bool = ((v7 & v7) == v7)
                                            fmt.Println(v2, v18)
                                            if v18 {
                                                v7 += 127
                                            } else {
                                                switch (fi6 * 9223372036854775807) {
                                                case -9223372036854775808:
                                                    var v19 int32 = 1
                                                    _ = v19
                                                case 2147483647:
                                                    v20 := GenS4{F0: fi12, F1: 1000}
                                                    _ = v20
                                                case 32767:
                                                    var v21 int8 = v7
                                                    _ = v21
                                                default:
                                                    if (((127 <= 125) && false) && (v18 && (452 != -2147483648))) {
                                                        if v18 {
                                                            if (!v18) {
                                                                v22 := GenS5{F0: uint(v2), F1: 691}
                                                                _ = v22
                                                            }
                                                            v7 -= (v7 - genFn0(uint16(v1), 32767))
                                                            v17 = 336
                                                        } else {
                                                            v23 := []int{127}
                                                            _ = v23
                                                        }
                                                        delete(v3, v2)
                                                        var v24 float32 = float32((((0.0 / 1.0) + 3.14159) - float64(1.5)))
                                                        _ = v24
                                                    } else {
                                                        switch fi12 {
                                                        case 42:
                                                            var v25 int = fi12
                                                            _ = v25
                                                        case -1:
                                                            if ((((3.14159 - 0.0) - 3.14159) + (float64(-2.25) + 0.0)) == 0.1) {
                                                                v26 := GenS6{F0: 2, F1: v17, F2: ("hello" + "foo bar")}
                                                                var v27 int64 = -1
                                                                _ = v26
                                                                _ = v27
                                                            }
                                                        default:
                                                            if true {
                                                                var v28 string = ""
                                                                _ = v28
                                                            } else {
                                                                var v29 uint16 = (v17 / (((v17 << ((v1) & 63))) | 1))
                                                                _ = v29
                                                            }
                                                        }
                                                        var v30 uint16 = 0
                                                        _ = v30
                                                    }
                                                }
                                                v31 := GenS6{F0: 18446744073709551615, F1: (v17 >> ((v1) & 63)), F2: "foo bar"}
                                                v15 = (v15 - -2.25)
                                                _ = v31
                                            }
                                            _ = v17
                                            _ = v18
                                        }
                                        switch fi12 {
                                        case 9223372036854775807:
                                            v2 += 1
                                        default:
                                            if ((3.14159 > 1e10) && true) {
                                                fmt.Println(fi12, v2)
                                            } else {
                                                if false {
                                                    _ = len(v3)
                                                    v3[v2] = 2147483648
                                                }
                                            }
                                        }
                                        v32 := GenS7{F0: "foo bar", F1: 32767}
                                        _ = v32
                                    case -1:
                                        fmt.Println(v2)
                                    default:
                                        v33 := []int{-1}
                                        _ = v33
                                    }
                                } else {
                                    v15 = 3.14159
                                }
                                _ = v15
                            }
                            var v34 uint16 = 822
                            _ = fi12
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 8; fi35++ {
                            if true {
                                v1 = v1
                            }
                            v36 := []int{}
                            v7 = v7
                            _ = fi35
                            _ = v36
                        }
                        fmt.Println(fi6)
                    }
                }
                _ = fi6
                _ = v7
            }
        case "hello":
            var v37 uint16 = 1
            _ = v37
        default:
            if (false && false) {
                v38 := GenS8{F0: (^(v1))}
                var v39 string = "\n\t"
                _ = v38
                _ = v39
            }
        }
    case -9223372036854775808:
        v40 := GenS4{F0: -296, F1: 526}
        _ = v40
    default:
        switch ("\n\t" + "\n\t") {
        case "hello":
            var v41 uint32 = (v2 & v2)
            _ = v41
        }
    }
    v2 -= v2
    v2 += 743
    var v42 int32 = 0
    for fi43 := 0; fi43 < 12; fi43++ {
        var v44 uint = 65535
        if false {
            v42 -= (v42 >> ((18446744073709551615) & 63))
            v2 = 42
            if false {
                v1 -= v44
                for fi45 := 0; fi45 < 11; fi45++ {
                    if (v2 != v2) {
                        var v46 int8 = 45
                        _ = v46
                    }
                    v47 := GenS3{F0: false, F1: 960, F2: false}
                    if (v2 < (v2 >> (((^(v44))) & 63))) {
                        if ((!(false && true)) || true) {
                            var v48 int16 = -1
                            v2 = (^(v2))
                            if true {
                                v44 -= v44
                                var v49 uint8 = 255
                                if (true && (896 != 42)) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        if (0 >= 0) {
                                            v50 := GenS9{F0: (0 >= 42), F1: -905}
                                            fmt.Println(fi43)
                                            v42 += v42
                                            _ = v50
                                        }
                                    default:
                                        v51 := []int{}
                                        _ = v51
                                    }
                                    if (float64((1e10 - 1.5)) > 1.5) {
                                        v48 = 943
                                        switch fi45 {
                                        case 497:
                                            if ("" >= "a") {
                                                v44 += (v44 ^ (v44 >> ((v1) & 63)))
                                                if true {
                                                    var v52 float32 = -2.25
                                                    if (!(!true)) {
                                                        v2 += v2
                                                    } else {
                                                        var v53 uint = 4294967295
                                                        v54 := []int{}
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                    _ = v52
                                                }
                                                v42 -= (v42 | v42)
                                            }
                                        case -9223372036854775808:
                                            switch "hello" {
                                            case "\n\t":
                                                v55 := GenS5{F0: (v1 * (^(v44))), F1: 1}
                                                _ = v55
                                            case "hello":
                                                v49 += 0
                                            case "":
                                                switch fi43 {
                                                case -60:
                                                    v3[(v2 & (v2 + v2))] = 42
                                                case 4294967295:
                                                    v56 := []int{}
                                                    _ = v56
                                                case 127:
                                                    v57 := GenS7{F0: "", F1: 2}
                                                    _ = v57
                                                }
                                            default:
                                                var v58 bool = true
                                                _ = v58
                                            }
                                        case -949:
                                            var v59 uint8 = (v49 / (((v49 * v49)) | 1))
                                            _ = v59
                                        default:
                                            if false {
                                                switch "a" {
                                                case "foo bar":
                                                    var v60 int64 = 0
                                                    _ = v60
                                                case "\n\t":
                                                    v44 = v1
                                                case "a":
                                                    fmt.Println(v49, v44)
                                                default:
                                                    v61 := GenS1{F0: v44, F1: 909, F2: v42}
                                                    _ = len(v3)
                                                    _ = v61
                                                }
                                                if true {
                                                    var v62 string = "hello"
                                                    fmt.Println(v62)
                                                    var v63 uint8 = (v49 & v49)
                                                    _ = v62
                                                    _ = v63
                                                } else {
                                                    fmt.Println(v1, v1, v44)
                                                    if false {
                                                        v44 = v44
                                                    } else {
                                                        v49 -= v49
                                                        fmt.Println(v1, v44, v1)
                                                    }
                                                    v3[(v2 & (v2 & (v2 >> (((^(v1))) & 63))))] = v2
                                                }
                                            }
                                        }
                                        var v64 int = fi43
                                        v1 -= uint(v2)
                                        if (!true) {
                                            fmt.Println(fi45)
                                            fmt.Println(v48)
                                            if (false || true) {
                                                if (!false) {
                                                    if (440 < 42) {
                                                        switch fi43 {
                                                        case 9223372036854775807:
                                                            v65 := []int{(fi45 << ((v44) & 63)), fi45, (^(fi43))}
                                                            _ = v65
                                                        default:
                                                            if ((!((!(48 > 1)) && true)) && ((1.5 * -2.25) <= 0.1)) {
                                                                if true {
                                                                    var v66 int = (fi43 & v64)
                                                                    v66 = (-(v66))
                                                                    v67 := []int{int(v49), v66, v64}
                                                                    _ = v66
                                                                    _ = v67
                                                                } else {
                                                                    var v68 uint8 = (^(v49))
                                                                    var v69 uint = 255
                                                                    if false {
                                                                        v69 -= 4294967295
                                                                        var v70 int16 = v48
                                                                        _ = v70
                                                                    }
                                                                    _ = v68
                                                                    _ = v69
                                                                }
                                                                v71 := GenS10{F0: v42}
                                                                switch int(v48) {
                                                                case -405:
                                                                    v44 -= (v44 / ((v44) | 1))
                                                                case 4294967295:
                                                                    v44 = 18446744073709551615
                                                                case -209:
                                                                    v72 := GenS11{F0: v49, F1: 42, F2: (^(v42))}
                                                                    _ = v72
                                                                default:
                                                                    _ = len(v3)
                                                                }
                                                                _ = v71
                                                            }
                                                        }
                                                        v73 := GenS12{F0: (v48 % ((839) | 1)), F1: 1}
                                                        _ = v73
                                                    }
                                                    v44 = v44
                                                    v74 := GenS13{F0: 65535}
                                                    if true {
                                                        var v75 uint8 = v49
                                                        _ = v75
                                                    }
                                                    _ = v74
                                                } else {
                                                    var v76 float64 = 1e10
                                                    if true {
                                                        if false {
                                                            v77 := GenS4{F0: (fi43 / ((32768) | 1)), F1: int16(v64)}
                                                            var v78 uint64 = 89
                                                            _ = v77
                                                            _ = v78
                                                        } else {
                                                            fmt.Println(v49)
                                                        }
                                                        if false {
                                                            v79 := []int{0}
                                                            v3[v2] = 65535
                                                            _ = v79
                                                        }
                                                    }
                                                    _ = v76
                                                }
                                                v80 := GenS6{F0: 1, F1: 883, F2: "hello"}
                                                _ = v80
                                            } else {
                                                v1 += 42
                                                if false {
                                                    var v81 int = -571
                                                    switch ("\n\t" + ("a" + "")) {
                                                    case "foo bar":
                                                        var v82 uint16 = 32767
                                                        _ = v82
                                                    }
                                                    _ = v81
                                                } else {
                                                    v2 = v2
                                                    v1 -= 255
                                                    delete(v3, 42)
                                                }
                                                var v83 bool = true
                                                _ = v83
                                            }
                                        } else {
                                            v84 := v3[uint32(v1)]
                                            v42 = (v42 % (((v42 - (v42 >> ((v44) & 63)))) | 1))
                                            _ = v84
                                        }
                                        _ = v64
                                    }
                                    var v85 float64 = float64((1e10 - (0.1 - float32(1.5))))
                                    _ = len(v3)
                                    v86 := GenS2{F0: fi43, F1: 732}
                                    _ = v85
                                    _ = v86
                                } else {
                                    v48 = (v48 & v48)
                                    v1 = (^(v1))
                                    var v87 uint = v1
                                    _ = v87
                                }
                                v44 = v44
                                _ = v49
                            } else {
                                switch "hello" {
                                case "foo bar":
                                    var v88 uint8 = 255
                                    _ = v88
                                }
                            }
                            _ = v48
                        } else {
                            var v89 int32 = (v42 & (v42 ^ v42))
                            if false {
                                fmt.Println(fi45)
                                v42 += v42
                            } else {
                                v90 := GenS4{F0: (-(fi43)), F1: 604}
                                v2 = v2
                                var v91 int = fi43
                                _ = v90
                                _ = v91
                            }
                            v92 := GenS9{F0: (true && (true || ((v42 & 2147483647) < v42))), F1: fi45}
                            var v93 int64 = -111
                            _ = len(v3)
                            var v94 int8 = -128
                            _ = v89
                            _ = v92
                            _ = v93
                            _ = v94
                        }
                        switch (fi43 >> (((v1 ^ 65535)) & 63)) {
                        case 127:
                            fmt.Println(v44, v44, fi43)
                        default:
                            v1 += v1
                        }
                    } else {
                        v44 += v1
                    }
                    _ = fi45
                    _ = v47
                }
            }
        } else {
            switch "foo bar" {
            case "foo bar":
                if true {
                    v95 := v3[v2]
                    var v96 int32 = (v42 | (v42 | (v42 / ((-2147483648) | 1))))
                    _ = v95
                    _ = v96
                } else {
                    fmt.Println(fi43, v2, v2)
                    var v97 uint8 = 64
                    var v98 int32 = (v42 + 2147483647)
                    _ = v97
                    _ = v98
                }
            case "a":
                v99 := []int{(fi43 / ((fi43) | 1)), 799}
                _ = v99
            case "\n\t":
                v2 = (v2 & 0)
            default:
                v100 := GenS14{F0: false, F1: true, F2: true}
                _ = v100
            }
            var v101 int16 = int16(fi43)
            v102 := []int{}
            _ = v101
            _ = v102
        }
        _ = fi43
        _ = v44
    }
    v103 := v3[(v2 + (v2 + 65535))]
    for fi104 := 0; fi104 < 5; fi104++ {
        v42 = (v42 / ((v42) | 1))
        for fi105 := 0; fi105 < 11; fi105++ {
            fmt.Println(v1, fi104)
            _ = fi105
        }
        v103 -= v103
        _ = fi104
    }
    var v106 float32 = 0.0
    v1 = (v1 - 18446744073709551615)
    v107 := []int{9223372036854775807, -9223372036854775808, 205, -153}
    v103 = v103
    var v108 bool = ((!(!true)) || (true && (930 != 42)))
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v106:", v106)
    fmt.Println("slicelen v107:", len(v107), v107)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v103
    _ = v106
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v42
}
