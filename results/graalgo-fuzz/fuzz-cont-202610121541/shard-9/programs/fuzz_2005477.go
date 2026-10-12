package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int; F1 int; F2 uint64 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 int32; F1 string; F2 uint16 }
type GenS7 struct { F0 int16; F1 bool; F2 uint32 }
type GenS8 struct { F0 int32; F1 int16; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    var v3 int8 = 77
    v3 = -128
    var v4 uint64 = 18446744073709551615
    v4 = (v4 % ((v4) | 1))
    if true {
        v3 += (v3 >> ((2) & 63))
        var v5 uint64 = v4
        _ = v5
    } else {
        v6 := GenS1{F0: -1}
        _ = v6
    }
    switch 931 {
    case 32768:
        var v7 int64 = -825
        _ = v7
    default:
        if false {
            var v8 uint8 = 128
            _ = v8
        } else {
            v9 := GenS2{F0: v3, F1: uint32(v4)}
            for fi10 := 0; fi10 < 3; fi10++ {
                switch 32767 {
                case 0:
                    for fi11 := 0; fi11 < 2; fi11++ {
                        var v12 uint32 = 746
                        switch fi10 {
                        case 32768:
                            fmt.Println(fi11)
                            v13 := []int{}
                            v14 := GenS2{F0: v3, F1: 640}
                            _ = v13
                            _ = v14
                        case 190:
                            switch "foo bar" {
                            case "a":
                                switch ((("a" + "\n\t") + "foo bar") + ("\n\t" + "\n\t")) {
                                case "\n\t":
                                    if (!(!(false && true))) {
                                        var v15 int16 = 941
                                        v16 := GenS3{F0: fi10, F1: (fi10 % (((-(fi11))) | 1)), F2: (v4 * 248)}
                                        if true {
                                            fmt.Println(v4)
                                        } else {
                                            var v17 int16 = (^(v15))
                                            switch "foo bar" {
                                            case "\n\t":
                                                switch (^(fi10)) {
                                                case -2147483648:
                                                    if true {
                                                        v18 := GenS3{F0: (fi10 | (^(fi11))), F1: (fi10 - 431), F2: v4}
                                                        _ = v18
                                                    } else {
                                                        v2[864] = fi11
                                                        if (1.5 < 3.14159) {
                                                            fmt.Println(v15)
                                                        }
                                                        switch 9223372036854775807 {
                                                        case 720:
                                                            var v19 int64 = 1
                                                            _ = v19
                                                        case 999:
                                                            v4 = 4294967295
                                                            v15 = v17
                                                        default:
                                                            if true {
                                                                switch (fi11 * fi11) {
                                                                case 112:
                                                                    if (1e10 == (float32(1.5) - float32((((3.14159 / 1.0) - (((float64((-2.25 - 1.5)) / (3.14159 / float64(3.14159))) / 1.5) / 3.14159)) / -2.25)))) {
                                                                        v17 = (v15 - -742)
                                                                    }
                                                                case 733:
                                                                    var v20 int16 = (^(v17))
                                                                    _ = v20
                                                                case -1:
                                                                    v21 := []int{(fi10 << ((917) & 63))}
                                                                    var v22 uint = 0
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        var v23 float32 = 1e10
                                                                        _ = v23
                                                                    default:
                                                                        var v24 int = (-(fi11))
                                                                        _ = v24
                                                                    }
                                                                    switch 881 {
                                                                    case 0:
                                                                        v22 = (v22 >> ((1) & 63))
                                                                    case -2147483648:
                                                                        v12 -= v12
                                                                    case 6:
                                                                        v21 = append(v21, (fi11 >> (((v22 / (((v22 ^ v22)) | 1))) & 63)))
                                                                    }
                                                                    _ = v21
                                                                    _ = v22
                                                                }
                                                            } else {
                                                                var v25 uint16 = 42
                                                                _ = v25
                                                            }
                                                        }
                                                        var v26 string = "foo bar"
                                                        fmt.Println(v12)
                                                        _ = v26
                                                    }
                                                case 9223372036854775807:
                                                    v17 -= v15
                                                }
                                            default:
                                                v27 := []int{(fi10 + fi10)}
                                                _ = v27
                                            }
                                            _ = v17
                                        }
                                        var v28 uint64 = v4
                                        if false {
                                            var v29 string = ""
                                            var v30 int64 = 419
                                            _ = v29
                                            _ = v30
                                        } else {
                                            var v31 uint32 = (v12 & (^(v12)))
                                            switch (fi10 % ((fi11) | 1)) {
                                            case -1:
                                                var v32 float32 = (1.5 + 1.5)
                                                _ = v32
                                            case -18:
                                                v12 = (v31 + v31)
                                            default:
                                                var v33 int32 = int32(v15)
                                                _ = v33
                                            }
                                            _ = v31
                                        }
                                        v34 := v2[v12]
                                        _ = v15
                                        _ = v16
                                        _ = v28
                                        _ = v34
                                    }
                                case "":
                                    if (241 == 128) {
                                        switch fi10 {
                                        case 32767:
                                            v12 -= v12
                                        case 146:
                                            delete(v2, v12)
                                        default:
                                            fmt.Println(v12, v4, v3)
                                        }
                                        fmt.Println(fi10, fi11)
                                        switch ((((("hello" + (("" + "") + (("" + "hello") + ("foo bar" + "hello")))) + "\n\t") + "\n\t") + ("foo bar" + ("" + (("hello" + "hello") + ((("" + ("hello" + "a")) + "foo bar") + ("hello" + ("\n\t" + "foo bar"))))))) + "hello") {
                                        case "foo bar":
                                            if true {
                                                v35 := GenS1{F0: fi10}
                                                var v36 uint16 = 42
                                                _ = v35
                                                _ = v36
                                            } else {
                                                var v37 int64 = 772
                                                _ = v37
                                            }
                                        default:
                                            if ((v3 << ((2) & 63)) >= 0) {
                                                v38 := GenS4{F0: true, F1: 65535}
                                                var v39 uint16 = 65535
                                                v40 := GenS5{F0: 1, F1: 13}
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                            } else {
                                                var v41 bool = (3.14159 <= 0.1)
                                                switch (("hello" + "\n\t") + "a") {
                                                case "hello":
                                                    v12 -= v12
                                                    var v42 int64 = -1
                                                    v2[(v12 + (^(v12)))] = -2147483648
                                                    _ = v42
                                                case "foo bar":
                                                    delete(v2, (v12 % ((v12) | 1)))
                                                }
                                                v43 := GenS3{F0: fi10, F1: fi10, F2: v4}
                                                _ = v41
                                                _ = v43
                                            }
                                        }
                                    }
                                }
                            default:
                                fmt.Println(fi10, v4, fi11)
                            }
                        default:
                            v3 -= v3
                        }
                        switch "hello" {
                        case "":
                            v44 := []int{(fi11 / ((fi11) | 1)), 42}
                            _ = v44
                        case "foo bar":
                            switch "hello" {
                            case "hello":
                                v45 := GenS5{F0: 0, F1: 1}
                                _ = v45
                            case "a":
                                switch "\n\t" {
                                case "a":
                                    v46 := GenS3{F0: fi11, F1: (fi10 << ((2) & 63)), F2: (v4 / ((v4) | 1))}
                                    _ = v46
                                case "":
                                    switch (("\n\t" + "foo bar") + "foo bar") {
                                    case "a":
                                        switch 394 {
                                        case 4294967295:
                                            var v47 uint64 = v4
                                            _ = v47
                                        case 128:
                                            v48 := v2[v12]
                                            fmt.Println(fi10, fi10, v48)
                                            fmt.Println(v4)
                                            v4 -= 0
                                            switch "foo bar" {
                                            case "a":
                                                var v49 float32 = 1e10
                                                _ = v49
                                            case "foo bar":
                                                fmt.Println(v3, fi11)
                                            }
                                            _ = v48
                                        default:
                                            v12 += v12
                                        }
                                    default:
                                        var v50 uint16 = 957
                                        _ = v50
                                    }
                                }
                            }
                        case "\n\t":
                            if false {
                                if (9223372036854775807 <= 1) {
                                    fmt.Println(fi10, fi10, fi11)
                                    var v51 float64 = 0.0
                                    v52 := []int{-9223372036854775808, int(v3), 32768}
                                    v53 := GenS3{F0: fi11, F1: int(v12), F2: (v4 * v4)}
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                } else {
                                    v4 = v4
                                    if false {
                                        var v54 int64 = 0
                                        v4 = 2
                                        var v55 uint32 = (^(v12))
                                        fmt.Println(v54, v3, fi11)
                                        _ = v54
                                        _ = v55
                                    } else {
                                        v3 = v3
                                        var v56 int32 = 2147483647
                                        v57 := GenS3{F0: (fi11 * (fi11 * 42)), F1: (fi11 >> ((4294967295) & 63)), F2: (v4 + v4)}
                                        _ = v56
                                        _ = v57
                                    }
                                    v12 = v12
                                }
                                var v58 float32 = float32(0.1)
                                v59 := GenS1{F0: (fi10 >> ((744) & 63))}
                                _ = v58
                                _ = v59
                            } else {
                                var v60 float32 = (1e10 + -2.25)
                                _ = v60
                            }
                        default:
                            var v61 int16 = -125
                            _ = v61
                        }
                        if (!false) {
                            var v62 bool = false
                            if (false && ((999 >= 0) || ((v4 > (v4 << ((65535) & 63))) || v62))) {
                                v63 := v1[("hello" + "")]
                                v63 = v63
                                var v64 string = "\n\t"
                                var v65 bool = (129 < uint8(v12))
                                _ = v63
                                _ = v64
                                _ = v65
                            }
                            _ = v62
                        }
                        switch 32767 {
                        case 32767:
                            if (fi11 >= (fi11 >> ((8) & 63))) {
                                v66 := v1[""]
                                v67 := []int{-957, (fi10 - fi10)}
                                v4 = (^(v4))
                                _ = v66
                                _ = v67
                            } else {
                                var v68 string = "foo bar"
                                v69 := GenS6{F0: 42, F1: v68, F2: 771}
                                _ = v68
                                _ = v69
                            }
                        default:
                            v12 += v12
                        }
                        _ = fi11
                        _ = v12
                    }
                default:
                    var v70 int = fi10
                    _ = v70
                }
                v71 := GenS7{F0: 889, F1: false, F2: 227}
                for fi72 := 0; fi72 < 7; fi72++ {
                    v73 := GenS1{F0: 1}
                    var v74 int8 = v3
                    v75 := GenS7{F0: -32768, F1: true, F2: 834}
                    if (false && (true || (2 >= 4294967295))) {
                        switch (fi10 - 14) {
                        case 1:
                            var v76 uint8 = 1
                            var v77 uint8 = v76
                            _ = v76
                            _ = v77
                        case 32767:
                            if false {
                                var v78 uint64 = 676
                                if (0.0 > (1e10 / float64(1.5))) {
                                    var v79 int8 = v74
                                    v1["\n\t"] = ""
                                    _ = v79
                                }
                                v80 := []int{}
                                _ = v78
                                _ = v80
                            }
                        case 909:
                            switch "" {
                            case "":
                                v81 := v1[""]
                                _ = v81
                            case "hello":
                                v82 := GenS8{F0: 0, F1: 0, F2: (v4 & 0)}
                                _ = v82
                            case "a":
                                if (127 < 128) {
                                    v83 := v1["foo bar"]
                                    _ = v83
                                } else {
                                    v74 -= -58
                                }
                            default:
                                var v84 int = (fi10 & fi10)
                                _ = v84
                            }
                        }
                        v85 := []int{fi10, (fi10 % ((fi72) | 1))}
                        _ = len(v85)
                        _ = v85
                    } else {
                        var v86 uint64 = 18446744073709551615
                        v3 = (v3 % ((v3) | 1))
                        var v87 int8 = v3
                        _ = v86
                        _ = v87
                    }
                    if true {
                        if false {
                            switch "" {
                            case "":
                                switch ("\n\t" + "hello") {
                                case "hello":
                                    var v88 int8 = v74
                                    _ = v88
                                case "\n\t":
                                    v89 := []int{}
                                    _ = v89
                                default:
                                    v90 := GenS7{F0: 1, F1: (!false), F2: 231}
                                    _ = v90
                                }
                                fmt.Println(v3)
                            case "a":
                                v74 -= v3
                            }
                        }
                    }
                    v91 := []int{(fi72 - (fi10 - fi72)), fi72, fi72, fi10}
                    _ = fi72
                    _ = v73
                    _ = v74
                    _ = v75
                    _ = v91
                }
                _ = fi10
                _ = v71
            }
            for fi92 := 0; fi92 < 18; fi92++ {
                var v93 int32 = -2147483648
                _ = fi92
                _ = v93
            }
            _ = v9
        }
    }
    _ = len(v1)
    fmt.Println(v3)
    delete(v2, 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
