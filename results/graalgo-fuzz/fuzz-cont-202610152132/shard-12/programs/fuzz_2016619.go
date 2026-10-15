package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int32; F2 string }
type GenS4 struct { F0 int8; F1 bool; F2 int16 }
type GenS5 struct { F0 string; F1 string; F2 int64 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 bool; F1 int8 }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 int16 }
func genFn0(p0 bool, p1 int) uint {
    if ((255 < 2) || ((((!true) || p0) || false) && false)) {
        v6 := GenS1{F0: 170, F1: 1, F2: 42}
        p0 = (!p0)
        v7 := GenS1{F0: 428, F1: 42, F2: 0}
        _ = v6
        _ = v7
    } else {
        var v8 float32 = 1.5
        fmt.Println(p0)
        _ = v8
    }
    if true {
        v9 := GenS2{F0: 1}
        var v10 int16 = 969
        _ = v10
        _ = v9
    } else {
        fmt.Println(p1)
    }
    if (!p0) {
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 int8 = -125
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            p0 = p0
            var v14 float64 = -2.25
            p1 -= p1
            _ = fi13
            _ = v14
        }
    }
    v15 := []int{127}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v15
    return 0
}

func main() {
    v1 := []string{}
    var v2 int32 = -344
    var v3 uint16 = 726
    var v4 int64 = -1
    var v5 int64 = v4
    v3 = (^(v3))
    var v16 int16 = -32768
    var v17 int8 = 104
    if false {
        _ = len(v1)
        switch -1 {
        case 1:
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        case -1:
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS3{F0: 42, F1: v2, F2: (("" + "a") + ("" + "a"))}
                fmt.Println(v5, v4, v2)
                _ = fi18
                _ = v19
            }
        case -403:
            if (((v3 % (((v3 & (v3 % ((694) | 1)))) | 1)) <= (v3 * v3)) && (!("a" < "foo bar"))) {
                v5 = v5
                fmt.Println(v3)
                v17 = (v17 >> ((42) & 63))
            }
        default:
            v20 := GenS3{F0: 354, F1: (v2 + 42), F2: "a"}
            _ = v20
        }
    }
    _ = len(v1)
    _ = len(v1)
    v21 := GenS3{F0: 534, F1: (v2 - 0), F2: "a"}
    if (true || (792 != v3)) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v22 := GenS4{F0: (v17 + (v17 - 42)), F1: (-1 == (v5 / (((-(v4))) | 1))), F2: v16}
        _ = v22
    }
    switch -262 {
    case -902:
        if true {
            v16 = v16
        }
    case -12:
        var v23 uint8 = 134
        _ = v23
    default:
        var v24 int64 = 0
        _ = v24
    }
    var v25 float32 = 1.5
    v2 += (-(v2))
    v5 = (v5 | -1)
    for fi26 := 0; fi26 < 1; fi26++ {
        for fi27 := 0; fi27 < 2; fi27++ {
            switch "hello" {
            case "":
                v1 = append(v1, ("hello" + ("a" + ("foo bar" + ("hello" + ("hello" + "a"))))))
            case "\n\t":
                if (((v17 >> ((255) & 63)) < (v17 >> ((2) & 63))) || ((fi26 / ((fi27) | 1)) != -80)) {
                    if false {
                        if false {
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                                v16 = v16
                            } else {
                                var v28 float32 = v25
                                switch (fi26 >> ((18446744073709551615) & 63)) {
                                case -559:
                                    v29 := GenS2{F0: 238}
                                    _ = v29
                                default:
                                    if (!(v2 < 42)) {
                                        v1 = append(v1, "\n\t")
                                    } else {
                                        if (128 <= 0) {
                                            var v30 int = 128
                                            var v31 int32 = (v2 - 42)
                                            _ = v30
                                            _ = v31
                                        }
                                        var v32 int = fi26
                                        _ = v32
                                    }
                                }
                                _ = len(v1)
                                _ = v28
                            }
                            if (1 == 0) {
                                var v33 uint16 = (v3 - (v3 & v3))
                                v16 -= v16
                                switch (fi26 - 634) {
                                case -2147483648:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    v34 := GenS5{F0: "hello", F1: (("" + "\n\t") + "foo bar"), F2: -1}
                                    _ = v34
                                }
                                var v35 float32 = 1.5
                                var v36 bool = false
                                if (((v17 % (((^(v17))) | 1)) >= v17) || ((2147483647 == 2) || (-2.25 >= (float64(1.5) / 1.5)))) {
                                    v37 := GenS2{F0: 0}
                                    _ = v37
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v33
                                _ = v35
                                _ = v36
                            } else {
                                var v38 int = -466
                                _ = v38
                            }
                        } else {
                            fmt.Println(v2)
                            v39 := GenS1{F0: (^(v2)), F1: 2, F2: 5}
                            if (18446744073709551615 >= 1) {
                                v17 = v17
                                if false {
                                    var v40 uint16 = v3
                                    fmt.Println(v5)
                                    _ = v40
                                }
                            }
                            _ = v39
                        }
                        var v41 int16 = (v16 + v16)
                        v25 -= (1e10 - v25)
                        if (2147483647 == 2) {
                            if true {
                                v5 = (^(v5))
                                var v42 int64 = (v5 * (v5 << ((927) & 63)))
                                v43 := GenS5{F0: "a", F1: "a", F2: 42}
                                v3 += 182
                                fmt.Println(v41, v42, v16)
                                _ = v42
                                _ = v43
                            }
                            if ((v41 * v41) == v16) {
                                v4 = -58
                            }
                        } else {
                            var v44 int = (^(fi27))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v3 -= (v3 >> ((496) & 63))
                            _ = v44
                        }
                        _ = v41
                    }
                    v25 -= v25
                    switch fi27 {
                    case 32768:
                        var v45 int32 = 0
                        _ = v45
                    }
                }
            default:
                fmt.Println(v17)
            }
            _ = fi27
        }
        _ = fi26
    }
    v16 -= v16
    if (false || (42 >= (v5 % (((v5 % ((int64(v17)) | 1))) | 1)))) {
        for fi46 := 0; fi46 < 9; fi46++ {
            for fi47 := 0; fi47 < 6; fi47++ {
                var v48 float64 = (-2.25 / 0.1)
                var v49 uint8 = 42
                v48 -= 3.14159
                var v50 float64 = (v48 - (v48 / v48))
                var v51 uint8 = v49
                _ = len(v1)
                v1 = append(v1, "\n\t")
                _ = fi47
                _ = v48
                _ = v49
                _ = v50
                _ = v51
            }
            var v52 int32 = (v2 ^ int32(v17))
            if (!false) {
                v53 := GenS1{F0: 0, F1: 127, F2: 77}
                _ = v53
            } else {
                var v54 string = "\n\t"
                for fi55 := 0; fi55 < 8; fi55++ {
                    var v56 int64 = (v4 << ((322) & 63))
                    fmt.Println(v52)
                    v57 := GenS6{F0: 42, F1: (v4 << ((2) & 63))}
                    if ((1e10 * float64(v25)) <= 1e10) {
                        if false {
                            switch (("a" + (v54 + ((v54 + v54) + "hello"))) + (v54 + (v54 + v54))) {
                            case "\n\t":
                                fmt.Println(v2)
                            case "hello":
                                v58 := GenS7{F0: false, F1: (v17 >> ((65535) & 63))}
                                _ = v58
                            case "a":
                                v59 := GenS8{F0: 1, F1: v17}
                                _ = v59
                            }
                            if false {
                                if true {
                                    if len(v1) > 0 {
                                        v1[0] = v54
                                    }
                                    var v60 uint16 = 42
                                    _ = v60
                                } else {
                                    var v61 uint16 = v3
                                    v16 -= -362
                                    v17 = v17
                                    fmt.Println(v4, v2)
                                    _ = v61
                                }
                                var v62 uint64 = 93
                                var v63 bool = true
                                _ = v62
                                _ = v63
                            }
                            v64 := GenS7{F0: (v54 == (v54 + v54)), F1: int8(v2)}
                            _ = v64
                        } else {
                            v65 := GenS4{F0: -90, F1: false, F2: (v16 / ((v16) | 1))}
                            var v66 uint64 = 18446744073709551615
                            _ = v65
                            _ = v66
                        }
                        fmt.Println(fi46, v3)
                    }
                    v67 := GenS9{F0: 320}
                    _ = fi55
                    _ = v56
                    _ = v57
                    _ = v67
                }
                switch (fi46 / (((fi46 | fi46)) | 1)) {
                case 128:
                    v68 := GenS7{F0: (0 >= 4294967295), F1: v17}
                    _ = v68
                case 4294967295:
                    switch fi46 {
                    case 32767:
                        for fi69 := 0; fi69 < 4; fi69++ {
                            fmt.Println(fi46, v4)
                            _ = fi69
                        }
                    }
                case -2147483648:
                    for fi70 := 0; fi70 < 10; fi70++ {
                        var v71 uint64 = 2
                        fmt.Println(v3)
                        _ = fi70
                        _ = v71
                    }
                }
                _ = v54
            }
            _ = fi46
            _ = v52
        }
        if false {
            var v72 uint16 = v3
            var v73 uint8 = 221
            _ = v72
            _ = v73
        }
        switch 0 {
        case 2147483647:
            _ = len(v1)
        case 9223372036854775807:
            if true {
                fmt.Println(v17, v2)
                for fi74 := 0; fi74 < 18; fi74++ {
                    if (v2 >= v2) {
                        var v75 float32 = ((v25 * v25) + v25)
                        switch fi74 {
                        case 9223372036854775807:
                            switch (fi74 * fi74) {
                            case 112:
                                if (!false) {
                                    for fi76 := 0; fi76 < 12; fi76++ {
                                        fmt.Println(v75, v2, v75)
                                        _ = fi76
                                    }
                                    v77 := GenS3{F0: 0, F1: -354, F2: "a"}
                                    _ = v77
                                }
                            case 9223372036854775807:
                                var v78 uint16 = v3
                                _ = v78
                            case 32767:
                                if ((v16 + (v16 >> ((genFn0(true, fi74)) & 63))) < v16) {
                                    switch ("hello" + "hello") {
                                    case "a":
                                        v3 += v3
                                    case "foo bar":
                                        v79 := GenS7{F0: (!false), F1: v17}
                                        _ = v79
                                    }
                                    var v80 bool = (537 <= 4294967295)
                                    v81 := GenS7{F0: ((0.1 < float64((0.0 * v25))) || v80), F1: -128}
                                    _ = v80
                                    _ = v81
                                }
                            }
                        case -9223372036854775808:
                            fmt.Println(v2, v5)
                        default:
                            v82 := GenS8{F0: 2, F1: v17}
                            _ = v82
                        }
                        var v83 int8 = int8(fi74)
                        _ = v75
                        _ = v83
                    }
                    fmt.Println(v3, v16, v17)
                    _ = fi74
                }
                var v84 int64 = v4
                _ = v84
            }
        default:
            v85 := GenS10{F0: (v16 * 1000), F1: (^(v16))}
            _ = v85
        }
    } else {
        fmt.Println(v2)
    }
    if false {
        fmt.Println(v2)
    } else {
        switch 892 {
        case -30:
            var v86 uint32 = 65535
            _ = v86
        }
    }
    var v87 uint32 = 404
    var v88 int64 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v87
    _ = v88
}
