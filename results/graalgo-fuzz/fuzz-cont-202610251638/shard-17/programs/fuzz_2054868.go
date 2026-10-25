package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 bool; F2 bool }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16 }

func main() {
    var v1 float64 = (((3.14159 / 1.5) * (1e10 + (1e10 / -2.25))) / (-2.25 * 0.1))
    var v2 uint64 = 0
    var v3 uint64 = (v2 | v2)
    var v4 int16 = 42
    v1 += float64(1e10)
    v2 += (v2 | 0)
    v5 := GenS1{F0: uint(v3)}
    v4 -= (v4 / (((v4 >> ((42) & 63))) | 1))
    for fi6 := 0; fi6 < 9; fi6++ {
        if ((42 <= 2) || (("" + "hello") != "")) {
            var v7 uint = 1
            v2 = (v2 / (((v2 / ((v2) | 1))) | 1))
            v8 := GenS1{F0: (^(v7))}
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 1; fi9++ {
                if (!(true && (1.5 > v1))) {
                    var v10 float64 = 3.14159
                    v11 := GenS1{F0: 1}
                    _ = v10
                    _ = v11
                } else {
                    v2 = (v2 % ((v2) | 1))
                    switch fi9 {
                    case -203:
                        v4 -= 42
                    case 4294967295:
                        if (!(42 >= 219)) {
                            switch fi9 {
                            case 771:
                                v1 -= v1
                            case -906:
                                if false {
                                    if (("" + "a") >= (("a" + "") + "a")) {
                                        fmt.Println(v4)
                                    } else {
                                        v12 := []int{(-(fi6))}
                                        v13 := GenS1{F0: 511}
                                        var v14 uint = 390
                                        _ = v12
                                        _ = v13
                                        _ = v14
                                    }
                                    if true {
                                        v15 := GenS1{F0: 4294967295}
                                        if (((0 != 9223372036854775807) && true) && (-358 <= 841)) {
                                            var v16 bool = false
                                            switch "foo bar" {
                                            case "a":
                                                fmt.Println(fi9, fi9)
                                            default:
                                                v16 = v16
                                            }
                                            switch "hello" {
                                            case "hello":
                                                switch "\n\t" {
                                                case "\n\t":
                                                    fmt.Println(v2, v1)
                                                    if (!(fi9 <= (-(fi9)))) {
                                                        v3 = 42
                                                    } else {
                                                        var v17 uint32 = uint32(v3)
                                                        _ = v17
                                                    }
                                                }
                                            default:
                                                if v16 {
                                                    var v18 uint16 = 1
                                                    _ = v18
                                                }
                                            }
                                            _ = v16
                                        } else {
                                            fmt.Println(v4, fi9, v2)
                                        }
                                        _ = v15
                                    } else {
                                        if true {
                                            v19 := []int{fi6, fi6, (fi9 & (fi6 & (-(fi6))))}
                                            v4 = v4
                                            v20 := GenS1{F0: 2}
                                            _ = v19
                                            _ = v20
                                        }
                                        if true {
                                            v21 := []int{(^(fi9))}
                                            _ = v21
                                        } else {
                                            switch (fi9 >> ((278) & 63)) {
                                            case -2147483648:
                                                var v22 float64 = v1
                                                _ = v22
                                            case 32768:
                                                var v23 uint16 = uint16(v2)
                                                _ = v23
                                            case 32767:
                                                var v24 uint8 = 1
                                                _ = v24
                                            default:
                                                var v25 int8 = int8(v3)
                                                _ = v25
                                            }
                                        }
                                        switch fi9 {
                                        case 42:
                                            var v26 int8 = 0
                                            var v27 int8 = (^(v26))
                                            v28 := GenS2{F0: 201, F1: (false && (v2 > v3)), F2: (uint(v3) <= 517)}
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        case -754:
                                            var v29 int64 = -594
                                            _ = v29
                                        default:
                                            if (3.14159 != -2.25) {
                                                fmt.Println(fi6, v1, fi6)
                                                v30 := []int{}
                                                _ = v30
                                            } else {
                                                var v31 uint64 = (v3 % (((v3 >> ((1) & 63))) | 1))
                                                v31 -= v31
                                                _ = v31
                                            }
                                        }
                                        if true {
                                            var v32 int = fi6
                                            fmt.Println(v4, v1)
                                            _ = v32
                                        }
                                    }
                                    var v33 uint8 = 127
                                    _ = v33
                                }
                            case 127:
                                if true {
                                    var v34 string = "foo bar"
                                    switch fi6 {
                                    case 0:
                                        var v35 bool = false
                                        _ = v35
                                    case 544:
                                        v1 = 3.14159
                                    }
                                    if (v34 != v34) {
                                        v36 := GenS2{F0: -1, F1: true, F2: (false || (!false))}
                                        _ = v36
                                    }
                                    _ = v34
                                } else {
                                    var v37 int32 = 2147483647
                                    var v38 int = (fi9 * fi6)
                                    var v39 uint8 = 127
                                    var v40 uint32 = 15
                                    switch (v38 * (^(fi9))) {
                                    case -18:
                                        var v41 float64 = v1
                                        _ = v41
                                    case 279:
                                        var v42 float64 = v1
                                        _ = v42
                                    default:
                                        v43 := []int{v38}
                                        _ = v43
                                    }
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                            default:
                                v44 := GenS1{F0: 42}
                                v45 := []int{}
                                _ = v44
                                _ = v45
                            }
                            var v46 uint16 = 42
                            v4 = (v4 - (-(v4)))
                            v46 = (v46 ^ (v46 / (((v46 ^ v46)) | 1)))
                            v47 := GenS1{F0: 1}
                            _ = v46
                            _ = v47
                        } else {
                            v4 -= v4
                            var v48 uint8 = 128
                            if (11 <= 1) {
                                switch (("a" + (("a" + "") + "hello")) + "hello") {
                                case "a":
                                    v49 := GenS3{F0: -470, F1: ("a" + "foo bar")}
                                    _ = v49
                                default:
                                    v3 = v2
                                }
                                if true {
                                    var v50 float32 = 3.14159
                                    v51 := GenS4{F0: (fi9 << ((2) & 63)), F1: (v48 & 33), F2: (v2 ^ v2)}
                                    _ = v50
                                    _ = v51
                                } else {
                                    var v52 uint = 397
                                    _ = v52
                                }
                                if false {
                                    var v53 string = ("foo bar" + ("foo bar" + ("" + "foo bar")))
                                    v3 += v2
                                    _ = v53
                                }
                                var v54 uint = uint(v48)
                                var v55 uint = 0
                                _ = v54
                                _ = v55
                            } else {
                                var v56 int32 = 683
                                _ = v56
                            }
                            _ = v48
                        }
                    default:
                        if true {
                            var v57 float32 = ((0.0 * 1.5) * float32((v1 - v1)))
                            switch (-(fi6)) {
                            case 396:
                                var v58 uint8 = 2
                                _ = v58
                            }
                            _ = v57
                        } else {
                            v4 -= 58
                            if ("foo bar" > ("hello" + ("" + ""))) {
                                v59 := GenS5{F0: (v4 / (((^(v4))) | 1))}
                                _ = v59
                            }
                            var v60 int32 = -1
                            _ = v60
                        }
                    }
                    var v61 float32 = 1.5
                    v4 = v4
                    switch (fi9 | 1) {
                    case 32768:
                        var v62 uint = 462
                        _ = v62
                    case 0:
                        var v63 float32 = v61
                        _ = v63
                    case 741:
                        var v64 uint16 = 65535
                        _ = v64
                    default:
                        var v65 int32 = 654
                        _ = v65
                    }
                    switch ("hello" + "\n\t") {
                    case "\n\t":
                        var v66 int64 = 0
                        _ = v66
                    }
                    v1 -= (v1 / (-2.25 / v1))
                    _ = v61
                }
                v67 := []int{fi6}
                _ = fi9
                _ = v67
            }
        }
        for fi68 := 0; fi68 < 2; fi68++ {
            var v69 int = (fi68 >> ((1) & 63))
            _ = fi68
            _ = v69
        }
        var v70 uint64 = v2
        _ = fi6
        _ = v70
    }
    v71 := []int{379, 693}
    if false {
        if (!false) {
            var v72 int = 0
            switch "hello" {
            case "":
                for fi73 := 0; fi73 < 9; fi73++ {
                    v72 = fi73
                    fmt.Println(v3, v72)
                    _ = fi73
                }
            case "foo bar":
                var v74 uint64 = 840
                _ = v74
            case "a":
                fmt.Println(v1, v4, v3)
            }
            _ = v72
        } else {
            v4 -= -461
        }
        if ((^(v2)) >= (v2 >> ((137) & 63))) {
            v75 := GenS3{F0: 0, F1: "foo bar"}
            _ = v75
        } else {
            if true {
                if len(v71) > 0 {
                    v71[0] = 127
                }
            } else {
                v76 := GenS3{F0: 1, F1: ("foo bar" + ("a" + ("\n\t" + "\n\t")))}
                var v77 uint = 372
                _ = v76
                _ = v77
            }
            if true {
                if (false && false) {
                    v1 -= v1
                }
                var v78 int16 = v4
                var v79 uint32 = 65535
                _ = v78
                _ = v79
            } else {
                var v80 uint32 = 807
                _ = len(v71)
                var v81 float32 = 1e10
                _ = v80
                _ = v81
            }
        }
    }
    var v82 uint16 = 0
    switch -331 {
    case 1:
        if true {
            if (!false) {
                v2 = (v2 * (^(v3)))
                var v83 int32 = -1
                var v84 uint8 = 2
                _ = v83
                _ = v84
            }
            v71 = append(v71, -79)
        }
    default:
        if (683 >= (^(v3))) {
            var v85 float64 = v1
            var v86 int16 = v4
            for fi87 := 0; fi87 < 4; fi87++ {
                if len(v71) > 0 {
                    v71[0] = 0
                }
                var v88 int64 = 42
                _ = fi87
                _ = v88
            }
            _ = v85
            _ = v86
        }
    }
    fmt.Println(v1, v2)
    v1 = v1
    if len(v71) > 0 {
        v71[0] = -851
    }
    var v89 uint = 296
    if (false || false) {
        if (v4 != 42) {
            v4 += (v4 * (v4 - -32768))
            v3 = (v3 | v3)
            v4 += v4
        } else {
            for fi90 := 0; fi90 < 8; fi90++ {
                if len(v71) >= 2 {
                    v71 = v71[1:len(v71)]
                }
                _ = fi90
            }
        }
        v91 := GenS6{F0: 19}
        _ = v91
    }
    v82 = v82
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v71:", len(v71), v71)
    fmt.Println("v82:", v82)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
    _ = v82
    _ = v89
}
