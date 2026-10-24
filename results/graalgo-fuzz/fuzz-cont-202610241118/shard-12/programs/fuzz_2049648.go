package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint16; F2 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint; F1 uint32; F2 bool }
type GenS9 struct { F0 string; F1 int; F2 uint16 }
type GenS10 struct { F0 uint32; F1 uint; F2 uint8 }
func genFn0(p0 uint64, p1 bool, p2 uint16) int {
    v5 := GenS1{F0: 42}
    p0 -= (p0 + p0)
    v6 := []int{461, 32768, 458, -469}
    p2 = p2
    p2 += (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 string, p1 int) bool {
    v7 := []int{-2147483648, p1, -9223372036854775808}
    v8 := GenS1{F0: 248}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (-801 > -332)
}

func main() {
    var v1 int16 = 0
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 uint = 4294967295
    for fi9 := 0; fi9 < 8; fi9++ {
        v1 += -723
        v10 := GenS1{F0: 2147483647}
        for fi11 := 0; fi11 < 20; fi11++ {
            if (false && (2 == 65535)) {
                var v12 float32 = (1.5 - 1.5)
                _ = v12
            } else {
                delete(v3, ("hello" + "\n\t"))
                fmt.Println(v1, fi9, v1)
            }
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    var v13 int8 = 100
    if ((!(!true)) || false) {
        fmt.Println(v1, v1, v4)
    } else {
        v4 = 906
        if true {
            if (v13 != v13) {
                v4 -= v4
                v4 += (^(v4))
            }
            var v14 uint16 = 32767
            _ = v14
        } else {
            v15 := GenS2{F0: 525, F1: ((0 != int(v13)) && true)}
            v16 := []int{genFn0(42, false, 65535), 539}
            var v17 uint16 = 0
            _ = v15
            _ = v16
            _ = v17
        }
        fmt.Println(v4, v1, v4)
    }
    v13 -= v13
    v18 := GenS2{F0: 0, F1: true}
    if ("a" <= "a") {
        v19 := GenS2{F0: 2, F1: ((v4 & (v4 - (v4 + v4))) < (v4 >> (((v4 | v4)) & 63)))}
        _ = len(v2)
        fmt.Println(v4, v1, v1)
        _ = v19
    } else {
        if (false || true) {
            v20 := GenS1{F0: 980}
            _ = v20
        } else {
            fmt.Println(v1, v1, v4)
        }
        v21 := GenS3{F0: v13}
        var v22 string = ("hello" + "hello")
        _ = v21
        _ = v22
    }
    var v23 int = 32768
    v1 = v1
    for fi24 := 0; fi24 < 13; fi24++ {
        if false {
            v4 += 0
            switch fi24 {
            case -44:
                var v25 float64 = float64(float32(0.0))
                _ = v25
            case 1:
                v13 = v13
            default:
                v13 -= v13
            }
            var v26 uint32 = 2147483647
            _ = v26
        } else {
            var v27 uint = (v4 % (((^(v4))) | 1))
            _ = v27
        }
        v13 = v13
        for fi28 := 0; fi28 < 16; fi28++ {
            var v29 uint64 = 2
            var v30 int = fi24
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi24
    }
    v31 := v3["hello"]
    var v32 uint32 = 89
    v1 += (v1 | v1)
    if (!v31) {
        if v31 {
            if true {
                for fi33 := 0; fi33 < 18; fi33++ {
                    switch v23 {
                    case -3:
                        if (42 >= -2147483648) {
                            v34 := []int{v23}
                            switch fi33 {
                            case 953:
                                v35 := GenS1{F0: v32}
                                _ = v35
                            case -1:
                                var v36 string = "foo bar"
                                _ = v36
                            default:
                                v37 := GenS3{F0: v13}
                                _ = v37
                            }
                            _ = v34
                        } else {
                            fmt.Println(v31)
                            if (v31 || (-2.25 == (-2.25 - 3.14159))) {
                                switch ("foo bar" + "foo bar") {
                                case "hello":
                                    v38 := GenS4{F0: (!v31)}
                                    _ = v38
                                case "\n\t":
                                    switch ("hello" + "a") {
                                    case "foo bar":
                                        v39 := GenS3{F0: v13}
                                        _ = v39
                                    default:
                                        var v40 int8 = 42
                                        _ = v40
                                    }
                                case "foo bar":
                                    var v41 uint32 = v32
                                    _ = v41
                                }
                            }
                            v42 := GenS5{F0: fi33, F1: 255, F2: 137}
                            _ = v42
                        }
                    }
                    var v43 uint = (^(v4))
                    _ = fi33
                    _ = v43
                }
                switch "foo bar" {
                case "":
                    fmt.Println(v31)
                }
            }
            if v31 {
                switch "foo bar" {
                case "a":
                    v23 -= (v23 ^ (v23 >> (((v4 ^ 255)) & 63)))
                case "foo bar":
                    var v44 uint8 = 43
                    _ = v44
                case "":
                    if (v31 && (1 < 4294967295)) {
                        v32 = v32
                        for fi45 := 0; fi45 < 1; fi45++ {
                            fmt.Println(v23, v31)
                            for fi46 := 0; fi46 < 1; fi46++ {
                                if v31 {
                                    fmt.Println(v31, v13, v32)
                                    var v47 string = "hello"
                                    fmt.Println(v47)
                                    _ = v47
                                }
                                v23 += (fi46 + fi46)
                                switch (fi45 * (fi45 << ((0) & 63))) {
                                case 127:
                                    var v48 uint64 = 563
                                    _ = v48
                                case 42:
                                    if false {
                                        switch ("hello" + "") {
                                        case "":
                                            var v49 float32 = (((float32(3.14159) - 3.14159) * float32(3.14159)) / 0.1)
                                            _ = v49
                                        default:
                                            var v50 uint8 = 255
                                            _ = v50
                                        }
                                    }
                                default:
                                    if v31 {
                                        v51 := GenS1{F0: (v32 << ((v4) & 63))}
                                        _ = v51
                                    } else {
                                        v52 := GenS6{F0: 32767}
                                        v53 := []int{(-(fi45)), v23, (fi46 | (v23 % ((fi45) | 1))), 2147483647}
                                        if true {
                                            v23 += genFn0(42, v31, 1)
                                        } else {
                                            v54 := GenS6{F0: fi46}
                                            v55 := GenS7{F0: 78}
                                            delete(v2, (-(fi45)))
                                            v56 := v2[fi45]
                                            v57 := v2[0]
                                            _ = v54
                                            _ = v55
                                            _ = v56
                                            _ = v57
                                        }
                                        _ = v52
                                        _ = v53
                                    }
                                }
                                _ = fi46
                            }
                            switch "foo bar" {
                            case "hello":
                                var v58 int32 = -1
                                _ = v58
                            }
                            _ = fi45
                        }
                        v32 -= v32
                    }
                }
                v59 := GenS4{F0: ((!true) && v31)}
                delete(v2, (v23 / ((v23) | 1)))
                _ = v59
            } else {
                switch ("\n\t" + "a") {
                case "\n\t":
                    var v60 uint32 = 65535
                    _ = v60
                case "":
                    if (!true) {
                        fmt.Println(v4)
                        if (0.0 != (0.1 - 0.0)) {
                            v61 := GenS8{F0: (v4 << (((v4 - (v4 ^ v4))) & 63)), F1: (v32 << (((v4 << ((42) & 63))) & 63)), F2: v31}
                            v62 := v2[-9223372036854775808]
                            _ = v61
                            _ = v62
                        } else {
                            delete(v3, "hello")
                            if false {
                                for fi63 := 0; fi63 < 9; fi63++ {
                                    v64 := v2[genFn0(4294967295, v31, 99)]
                                    var v65 int = 128
                                    if true {
                                        v66 := GenS9{F0: "", F1: genFn0(4294967295, (!v31), 262), F2: 61}
                                        if v31 {
                                            var v67 string = "hello"
                                            _ = v67
                                        } else {
                                            var v68 int32 = 42
                                            fmt.Println(v64, v4, v64)
                                            v68 -= -429
                                            _ = v68
                                        }
                                        _ = v66
                                    } else {
                                        var v69 int = fi63
                                        v70 := GenS7{F0: uint8(v32)}
                                        _ = v69
                                        _ = v70
                                    }
                                    _ = fi63
                                    _ = v64
                                    _ = v65
                                }
                            }
                        }
                    } else {
                        for fi71 := 0; fi71 < 12; fi71++ {
                            var v72 uint = 2
                            fmt.Println(v4, v4, v1)
                            _ = fi71
                            _ = v72
                        }
                        v73 := GenS3{F0: v13}
                        _ = v73
                    }
                case "foo bar":
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v32)
                    case "":
                        v74 := GenS5{F0: v23, F1: 42, F2: 0}
                        _ = v74
                    case "a":
                        var v75 uint32 = (v32 + v32)
                        _ = v75
                    default:
                        for fi76 := 0; fi76 < 13; fi76++ {
                            fmt.Println(fi76, v31, v4)
                            v77 := GenS6{F0: -9223372036854775808}
                            _ = fi76
                            _ = v77
                        }
                    }
                default:
                    fmt.Println(v4, v13, v13)
                }
                for fi78 := 0; fi78 < 14; fi78++ {
                    v79 := v2[v23]
                    v23 += (fi78 * 127)
                    _ = fi78
                    _ = v79
                }
                var v80 int64 = -527
                _ = v80
            }
        } else {
            fmt.Println(v4)
        }
        for fi81 := 0; fi81 < 9; fi81++ {
            if (!(v31 && v31)) {
                v31 = ((506 > 734) || (!true))
            }
            var v82 int = 42
            switch v82 {
            case -2147483648:
                v31 = v31
            case 709:
                v4 -= v4
            }
            _ = fi81
            _ = v82
        }
    }
    if v31 {
        v23 = v23
    } else {
        fmt.Println(v31, v13, v32)
        for fi83 := 0; fi83 < 3; fi83++ {
            for fi84 := 0; fi84 < 17; fi84++ {
                var v85 uint32 = v32
                if true {
                    if false {
                        if true {
                            switch 9223372036854775807 {
                            case 4294967295:
                                delete(v2, genFn0(42, v31, 2))
                            default:
                                if v31 {
                                    switch v23 {
                                    case 42:
                                        var v86 int32 = 982
                                        _ = v86
                                    case 128:
                                        v85 = (v85 ^ (v32 % (((v32 * v85)) | 1)))
                                    case 32768:
                                        fmt.Println(v31)
                                    }
                                    fmt.Println(v85)
                                    v85 -= v32
                                } else {
                                    v87 := GenS1{F0: (v32 ^ v85)}
                                    _ = v87
                                }
                            }
                            if v31 {
                                v88 := []int{fi83}
                                _ = v88
                            }
                            v89 := v2[fi84]
                            _ = v89
                        }
                        switch fi84 {
                        case 0:
                            var v90 float64 = 0.1
                            _ = v90
                        case 1:
                            var v91 string = "\n\t"
                            _ = v91
                        }
                        fmt.Println(fi84, fi84)
                        v31 = v31
                        v13 = 1
                    } else {
                        v85 += (v32 + uint32(v4))
                    }
                    fmt.Println(v4, v13, v1)
                    v92 := GenS10{F0: (v85 >> ((1) & 63)), F1: 36, F2: 127}
                    _ = v92
                }
                switch -92 {
                case -706:
                    v23 += genFn0(1, (v31 && true), 42)
                case -1:
                    v93 := GenS10{F0: 0, F1: v4, F2: 144}
                    _ = v93
                }
                fmt.Println(v1, v4, v23)
                var v94 float64 = float64(1.5)
                _ = fi84
                _ = v85
                _ = v94
            }
            v13 += (v13 << ((656) & 63))
            _ = fi83
        }
        if (v13 < (v13 << ((65535) & 63))) {
            v95 := []int{v23, (^(v23)), (v23 - (-(v23)))}
            v31 = (false || false)
            delete(v2, (-(v23)))
            _ = v95
        } else {
            delete(v3, ((("hello" + "\n\t") + "foo bar") + "a"))
            if false {
                if (!false) {
                    var v96 int16 = v1
                    for fi97 := 0; fi97 < 3; fi97++ {
                        for fi98 := 0; fi98 < 7; fi98++ {
                            var v99 int64 = 207
                            var v100 int8 = v13
                            delete(v2, -550)
                            _ = fi98
                            _ = v100
                            _ = v99
                        }
                        _ = fi97
                    }
                    _ = v96
                } else {
                    switch (v23 >> (((v4 >> ((v4) & 63))) & 63)) {
                    case 0:
                        v32 = (v32 * v32)
                    }
                    for fi101 := 0; fi101 < 14; fi101++ {
                        var v102 uint = 255
                        _ = fi101
                        _ = v102
                    }
                }
                v4 -= v4
            }
            v1 = v1
        }
    }
    fmt.Println(v31)
    for fi103 := 0; fi103 < 20; fi103++ {
        for fi104 := 0; fi104 < 11; fi104++ {
            v105 := GenS3{F0: v13}
            _ = fi104
            _ = v105
        }
        _ = fi103
    }
    v23 += v23
    var v106 int16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v106
    _ = v13
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
