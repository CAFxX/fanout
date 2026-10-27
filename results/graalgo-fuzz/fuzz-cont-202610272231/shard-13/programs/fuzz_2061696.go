package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 int32; F1 uint }
type GenS9 struct { F0 int; F1 int32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int16; F1 uint64; F2 uint32 }
func genFn0(p0 uint8, p1 string, p2 uint) uint32 {
    p1 = "hello"
    var v5 int64 = 9223372036854775807
    var v6 uint64 = 1
    v7 := []int{197}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 2147483648
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    var v4 int32 = -213
    if (((v4 + v4) < (-(v4))) && false) {
        var v8 uint16 = 2
        _ = v8
    } else {
        v4 = (-(v4))
        var v9 uint = 943
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int8 = 1
            var v13 float64 = 1.5
            v13 = float64(((0.1 + (0.0 / (((float32((float64(3.14159) + v13)) / (float32(v13) + 0.1)) + 1e10) / (-2.25 / 1.5)))) * 0.0))
            var v14 uint8 = 255
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        if (42 >= 42) {
            v4 = (v4 >> ((0) & 63))
        }
        var v15 uint8 = 2
        _ = fi10
        _ = v15
    }
    if false {
        if (!true) {
            delete(v3, ("\n\t" + ("hello" + "\n\t")))
            switch "foo bar" {
            case "hello":
                v4 = -1
            case "foo bar":
                fmt.Println(v4)
            default:
                v16 := []int{-9223372036854775808, 0, 128}
                _ = v16
            }
            if false {
                if false {
                    var v17 float64 = 0.0
                    v17 -= (3.14159 / v17)
                    var v18 int = 9223372036854775807
                    _ = v17
                    _ = v18
                }
                var v19 uint8 = 255
                if (((!true) && true) && false) {
                    if false {
                        v20 := GenS1{F0: 448}
                        var v21 float32 = 0.0
                        _ = v20
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 5; fi22++ {
                            fmt.Println(fi22)
                            var v23 uint64 = 1
                            v4 = v4
                            _ = fi22
                            _ = v23
                        }
                    }
                    fmt.Println(v19, v4)
                    v24 := GenS2{F0: 2, F1: -32768}
                    _ = v24
                } else {
                    _ = len(v1)
                    if false {
                        for fi25 := 0; fi25 < 10; fi25++ {
                            delete(v3, "hello")
                            v26 := GenS3{F0: 9223372036854775807}
                            v27 := GenS1{F0: 0}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        v19 += v19
                    }
                }
                _ = v19
            } else {
                for fi28 := 0; fi28 < 12; fi28++ {
                    if (true && true) {
                        switch (("" + "\n\t") + ("\n\t" + ("foo bar" + ("a" + "")))) {
                        case "\n\t":
                            if false {
                                v29 := GenS1{F0: 899}
                                if (false || true) {
                                    switch "hello" {
                                    case "":
                                        for fi30 := 0; fi30 < 2; fi30++ {
                                            v31 := v3[""]
                                            v2[128] = 0
                                            _ = fi30
                                            _ = v31
                                        }
                                    default:
                                        var v32 bool = false
                                        _ = v32
                                    }
                                    v33 := GenS2{F0: 255, F1: 652}
                                    v4 = (v4 - 2147483647)
                                    _ = v33
                                }
                                _ = v29
                            }
                        }
                        v34 := GenS1{F0: -457}
                        v35 := GenS2{F0: 980, F1: -32768}
                        _ = v34
                        _ = v35
                    } else {
                        var v36 float64 = 3.14159
                        v4 = v4
                        for fi37 := 0; fi37 < 2; fi37++ {
                            v38 := GenS4{F0: ("\n\t" + "a"), F1: 42}
                            if ((((("" + ("foo bar" + ("\n\t" + ""))) + "a") + "\n\t") >= "") || false) {
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    if ((1 == 128) || false) {
                                        var v39 int = (fi28 << ((156) & 63))
                                        v40 := GenS4{F0: "hello", F1: 65535}
                                        _ = v39
                                        _ = v40
                                    } else {
                                        var v41 int16 = 854
                                        if (!(v36 < (((v36 + v36) * 0.0) / -2.25))) {
                                            v42 := GenS5{F0: 127, F1: 2}
                                            fmt.Println(v36, fi28)
                                            _ = v42
                                        } else {
                                            var v43 uint64 = 1
                                            if (("\n\t" >= ("a" + "hello")) && false) {
                                                v44 := GenS3{F0: -998}
                                                _ = len(v1)
                                                v45 := GenS6{F0: 65535, F1: ""}
                                                _ = v44
                                                _ = v45
                                            }
                                            _ = v43
                                        }
                                        switch (fi28 >> ((4294967295) & 63)) {
                                        case 128:
                                            var v46 uint16 = 353
                                            _ = v46
                                        case -9223372036854775808:
                                            v47 := GenS5{F0: 128, F1: 127}
                                            _ = v47
                                        case 32768:
                                            v41 = v41
                                        default:
                                            v48 := GenS7{F0: 128, F1: "\n\t"}
                                            _ = v48
                                        }
                                        _ = v41
                                    }
                                case "":
                                    v36 += (0.1 / float64(0.1))
                                case "foo bar":
                                    if (0.0 > float32((v36 - 3.14159))) {
                                        fmt.Println(fi28, v4, fi28)
                                    }
                                }
                            }
                            var v49 bool = ("" == "foo bar")
                            if v49 {
                                if v49 {
                                    _ = len(v2)
                                    var v50 int32 = v4
                                    _ = v50
                                }
                                var v51 int8 = 0
                                var v52 int16 = 532
                                _ = v51
                                _ = v52
                            } else {
                                v53 := []int{(fi28 & 32767), 336, 128, fi28}
                                v49 = v49
                                _ = v53
                            }
                            var v54 uint32 = 2147483647
                            _ = fi37
                            _ = v38
                            _ = v49
                            _ = v54
                        }
                        _ = v36
                    }
                    _ = fi28
                }
            }
        } else {
            for fi55 := 0; fi55 < 12; fi55++ {
                if false {
                    _ = len(v2)
                    if false {
                        v56 := []int{4294967295}
                        _ = v56
                    } else {
                        v4 = (v4 << ((16) & 63))
                        v4 = -2147483648
                    }
                } else {
                    switch (fi55 & (fi55 + fi55)) {
                    case 32768:
                        if true {
                            var v57 int = fi55
                            if false {
                                switch v57 {
                                case 972:
                                    if (43 > 785) {
                                        var v58 int = (fi55 >> ((612) & 63))
                                        v57 -= -1
                                        _ = v58
                                    }
                                case 128:
                                    v59 := []int{(fi55 % (((fi55 / ((fi55) | 1))) | 1))}
                                    _ = v59
                                case -640:
                                    var v60 float64 = 0.0
                                    _ = v60
                                default:
                                    for fi61 := 0; fi61 < 11; fi61++ {
                                        v4 = v4
                                        v62 := GenS8{F0: v4, F1: 65535}
                                        v57 = (fi61 - 42)
                                        _ = fi61
                                        _ = v62
                                    }
                                }
                                v4 = v4
                                var v63 uint = 18446744073709551615
                                _ = v63
                            }
                            var v64 int8 = 127
                            _ = v57
                            _ = v64
                        } else {
                            switch (("" + "") + "hello") {
                            case "\n\t":
                                for fi65 := 0; fi65 < 13; fi65++ {
                                    v4 -= v4
                                    _ = fi65
                                }
                            case "a":
                                var v66 string = (("hello" + "\n\t") + "\n\t")
                                _ = v66
                            case "foo bar":
                                for fi67 := 0; fi67 < 12; fi67++ {
                                    switch "foo bar" {
                                    case "hello":
                                        v4 = v4
                                    case "foo bar":
                                        var v68 int32 = (-(v4))
                                        _ = v68
                                    case "\n\t":
                                        fmt.Println(fi67)
                                    default:
                                        v69 := []int{(fi55 * (fi67 >> ((601) & 63))), (fi67 >> ((2) & 63)), fi55}
                                        _ = v69
                                    }
                                    delete(v2, (fi55 | fi55))
                                    var v70 bool = (true && true)
                                    _ = fi67
                                    _ = v70
                                }
                            default:
                                if false {
                                    for fi71 := 0; fi71 < 19; fi71++ {
                                        v72 := GenS1{F0: 9223372036854775807}
                                        v4 += (v4 << ((398) & 63))
                                        v4 = (^(v4))
                                        _ = fi71
                                        _ = v72
                                    }
                                    v4 += -1
                                    for fi73 := 0; fi73 < 10; fi73++ {
                                        v74 := []int{170}
                                        _ = fi73
                                        _ = v74
                                    }
                                } else {
                                    var v75 float64 = 1.5
                                    v75 += v75
                                    var v76 int16 = 587
                                    _ = v75
                                    _ = v76
                                }
                            }
                        }
                    case 1:
                        switch ("\n\t" + "hello") {
                        case "foo bar":
                            v4 = (v4 >> ((4294967295) & 63))
                        case "\n\t":
                            var v77 bool = true
                            _ = v77
                        case "":
                            fmt.Println(v4, v4)
                        default:
                            fmt.Println(v4, v4)
                        }
                    case 2147483647:
                        var v78 uint32 = 2147483648
                        _ = v78
                    }
                }
                _ = fi55
            }
        }
        for fi79 := 0; fi79 < 4; fi79++ {
            for fi80 := 0; fi80 < 15; fi80++ {
                v4 -= (-(v4))
                v81 := GenS9{F0: fi80, F1: (v4 >> ((255) & 63))}
                var v82 int32 = 42
                var v83 int16 = 1
                _ = fi80
                _ = v81
                _ = v82
                _ = v83
            }
            _ = fi79
        }
        v84 := []int{}
        _ = v84
    }
    for fi85 := 0; fi85 < 18; fi85++ {
        var v86 bool = false
        _ = fi85
        _ = v86
    }
    switch "hello" {
    case "foo bar":
        v4 += (v4 << ((42) & 63))
    default:
        switch 0 {
        case -24:
            if false {
                for fi87 := 0; fi87 < 1; fi87++ {
                    v4 -= 236
                    _ = fi87
                }
                v88 := GenS2{F0: 255, F1: 1}
                for fi89 := 0; fi89 < 5; fi89++ {
                    for fi90 := 0; fi90 < 6; fi90++ {
                        v91 := GenS10{F0: 1}
                        _ = fi90
                        _ = v91
                    }
                    _ = fi89
                }
                _ = v88
            } else {
                var v92 uint = 0
                _ = v92
            }
        case 978:
            v4 += v4
        }
    }
    v93 := GenS11{F0: 42, F1: 752, F2: 42}
    _ = len(v3)
    var v94 uint32 = genFn0(94, "\n\t", 255)
    var v95 uint32 = (v94 & (v94 / (((v94 % (((v94 ^ v94)) | 1))) | 1)))
    var v96 uint8 = 127
    var v97 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v93
    _ = v94
    _ = v95
    _ = v96
    _ = v97
}
