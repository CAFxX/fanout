package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string }

func main() {
    var v1 uint8 = 62
    var v2 uint8 = (v1 + 253)
    var v3 uint8 = 1
    v4 := []string{"foo bar"}
    v5 := make(map[uint32]int)
    v6 := []string{"\n\t", "hello", "a", "\n\t"}
    var v7 int64 = -885
    if true {
        var v8 int = 32767
        var v9 uint = 1
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 12; fi10++ {
            if false {
                v11 := GenS1{F0: "a"}
                for fi12 := 0; fi12 < 19; fi12++ {
                    v13 := GenS1{F0: "a"}
                    v14 := GenS2{F0: (-994 == (fi12 + (fi10 + fi12)))}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                var v15 float32 = -2.25
                _ = v11
                _ = v15
            }
            _ = fi10
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS1{F0: ""}
            fmt.Println(v3, v1, v7)
            _ = fi16
            _ = v17
        }
    }
    fmt.Println(v1)
    v3 = v3
    v1 += v3
    v2 = 42
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 uint32 = 1
        _ = fi18
        _ = v19
    }
    if (42 > 99) {
        fmt.Println(v2, v1, v1)
    }
    var v20 uint = uint(v2)
    v21 := GenS3{F0: (v2 << ((v20) & 63)), F1: uint64(v1), F2: 89}
    for fi22 := 0; fi22 < 13; fi22++ {
        switch ("hello" + "a") {
        case "a":
            v2 = (v1 ^ v3)
        default:
            if true {
                for fi23 := 0; fi23 < 17; fi23++ {
                    v1 += 255
                    var v24 string = ""
                    switch ((v24 + ("a" + v24)) + v24) {
                    case "\n\t":
                        var v25 int64 = (v7 | v7)
                        var v26 uint64 = 729
                        if false {
                            if ((v26 | uint64(v2)) <= 894) {
                                var v27 int16 = -32768
                                v28 := GenS3{F0: v2, F1: (v26 >> (((v20 - uint(v1))) & 63)), F2: 0}
                                v29 := GenS3{F0: v2, F1: v26, F2: 0}
                                _ = v27
                                _ = v28
                                _ = v29
                            } else {
                                v24 = ("" + "foo bar")
                            }
                            if (false && true) {
                                v1 -= (v2 % ((19) | 1))
                                switch -9223372036854775808 {
                                case 9223372036854775807:
                                    v2 = 251
                                case 32768:
                                    switch v24 {
                                    case "\n\t":
                                        if (v2 >= v1) {
                                            if (!(!false)) {
                                                delete(v5, 2147483648)
                                                var v30 int = 32767
                                                _ = v30
                                            }
                                            v5[42] = (-(fi22))
                                            switch v24 {
                                            case "foo bar":
                                                switch v24 {
                                                case "hello":
                                                    if false {
                                                        v31 := v5[2]
                                                        if false {
                                                            var v32 int8 = int8(v7)
                                                            v7 = (v25 - (v7 ^ 315))
                                                            switch v24 {
                                                            case "":
                                                                v33 := v5[42]
                                                                _ = v33
                                                            case "a":
                                                                switch v24 {
                                                                case "foo bar":
                                                                    fmt.Println(v2, v3, v3)
                                                                }
                                                            case "\n\t":
                                                                if true {
                                                                    v2 = (v3 * 2)
                                                                }
                                                            default:
                                                                v32 -= (v32 / (((v32 - int8(v2))) | 1))
                                                            }
                                                            var v34 string = (v24 + (v24 + "foo bar"))
                                                            var v35 uint = v20
                                                            var v36 int = int(v26)
                                                            delete(v5, 697)
                                                            _ = v32
                                                            _ = v34
                                                            _ = v35
                                                            _ = v36
                                                        } else {
                                                            switch v24 {
                                                            case "foo bar":
                                                                fmt.Println(v3, v26)
                                                            case "":
                                                                var v37 float32 = (3.14159 * 1e10)
                                                                _ = v37
                                                            default:
                                                                if true {
                                                                    var v38 bool = true
                                                                    _ = v38
                                                                }
                                                            }
                                                            var v39 int64 = -9223372036854775808
                                                            v40 := GenS1{F0: v24}
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                        _ = v31
                                                    }
                                                case "":
                                                    v2 = uint8(v26)
                                                case "foo bar":
                                                    v2 -= (v1 / ((228) | 1))
                                                }
                                            case "\n\t":
                                                v3 = v3
                                            case "hello":
                                                var v41 string = (v24 + v24)
                                                _ = v41
                                            default:
                                                v42 := GenS4{F0: (!true), F1: 0}
                                                _ = v42
                                            }
                                        } else {
                                            v43 := v5[2147483647]
                                            v44 := GenS4{F0: true, F1: -1}
                                            v45 := GenS5{F0: (v2 << ((v20) & 63))}
                                            fmt.Println(v7, v20, v26)
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                        }
                                    case "hello":
                                        var v46 int32 = 0
                                        _ = v46
                                    default:
                                        v26 -= v26
                                    }
                                    if true {
                                        var v47 int8 = 15
                                        if (("foo bar" + v24) < (v24 + (v24 + ("\n\t" + "a")))) {
                                            if false {
                                                v20 -= (v20 + 42)
                                            }
                                            var v48 int64 = v25
                                            _ = v48
                                        }
                                        var v49 uint64 = (v26 & v26)
                                        _ = v47
                                        _ = v49
                                    }
                                }
                                v24 = v24
                            } else {
                                v1 -= 131
                                var v50 int16 = 1000
                                _ = v50
                            }
                        } else {
                            if (-32768 != 0) {
                                var v51 int = fi23
                                v1 -= uint8(v20)
                                switch (fi22 | fi23) {
                                case 316:
                                    if true {
                                        v26 = v26
                                        var v52 int64 = v7
                                        fmt.Println(v3)
                                        _ = v52
                                    }
                                case 127:
                                    switch 4294967295 {
                                    case -2147483648:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case -9223372036854775808:
                                        fmt.Println(v25, v2)
                                    case -434:
                                        var v53 int8 = 0
                                        _ = v53
                                    }
                                case 1:
                                    var v54 float64 = 1.5
                                    _ = v54
                                }
                                _ = v51
                            } else {
                                if true {
                                    var v55 uint = 18446744073709551615
                                    _ = v55
                                } else {
                                    v56 := GenS4{F0: (-2.25 < (1.5 + ((1e10 / -2.25) / 1.0))), F1: 42}
                                    _ = v56
                                }
                            }
                            if (32767 != 32767) {
                                var v57 uint32 = 2
                                switch (fi23 << (((v20 | (^(v20)))) & 63)) {
                                case 901:
                                    var v58 float32 = 0.0
                                    _ = v58
                                case 32768:
                                    var v59 float32 = 3.14159
                                    _ = v59
                                case 4294967295:
                                    var v60 int = 32767
                                    _ = v60
                                }
                                _ = v57
                            }
                            var v61 int8 = 1
                            _ = v61
                        }
                        _ = v25
                        _ = v26
                    case "a":
                        var v62 uint32 = 42
                        _ = v62
                    case "":
                        _ = len(v5)
                    }
                    _ = fi23
                    _ = v24
                }
                switch (fi22 ^ 4294967295) {
                case -9223372036854775808:
                    for fi63 := 0; fi63 < 2; fi63++ {
                        var v64 uint8 = 122
                        v1 += (^(v64))
                        _ = len(v5)
                        _ = fi63
                        _ = v64
                    }
                case 9223372036854775807:
                    switch "\n\t" {
                    case "\n\t":
                        if (2147483647 < 127) {
                            v3 = 0
                            var v65 uint64 = 4294967295
                            var v66 uint = (v20 & v20)
                            _ = v65
                            _ = v66
                        }
                    }
                default:
                    switch (fi22 | fi22) {
                    case 42:
                        switch ("a" + "hello") {
                        case "foo bar":
                            for fi67 := 0; fi67 < 6; fi67++ {
                                var v68 int64 = (v7 >> ((v20) & 63))
                                var v69 int = 1
                                var v70 int8 = -1
                                _ = fi67
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                        case "":
                            var v71 int64 = (v7 << (((v20 % ((v20) | 1))) & 63))
                            _ = v71
                        case "hello":
                            if (v20 >= v20) {
                                v72 := GenS4{F0: true, F1: 1}
                                switch (-(fi22)) {
                                case 2147483647:
                                    if (("foo bar" + "\n\t") < "hello") {
                                        delete(v5, 2147483647)
                                        if (1 != 9223372036854775807) {
                                            if (-925 >= (fi22 & 127)) {
                                                v1 += (v3 << ((65535) & 63))
                                                var v73 int16 = -32768
                                                if true {
                                                    v74 := GenS6{F0: 255, F1: 0}
                                                    if len(v4) > 0 {
                                                        v4[0] = ("a" + "foo bar")
                                                    }
                                                    var v75 float64 = 1e10
                                                    _ = v74
                                                    _ = v75
                                                } else {
                                                    for fi76 := 0; fi76 < 7; fi76++ {
                                                        var v77 float64 = 0.0
                                                        if false {
                                                            v7 += (^(v7))
                                                        } else {
                                                            fmt.Println(v2)
                                                        }
                                                        _ = fi76
                                                        _ = v77
                                                    }
                                                }
                                                _ = v73
                                            }
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                var v78 int8 = 100
                                                _ = v78
                                            case "foo bar":
                                                var v79 bool = true
                                                _ = v79
                                            }
                                            v80 := GenS1{F0: ""}
                                            _ = v80
                                        }
                                    } else {
                                        var v81 float32 = ((float32((float64((((3.14159 - float32(0.0)) + ((float32(((0.1 - (1.5 - (0.0 / -2.25))) - -2.25)) + ((3.14159 * (3.14159 / (float32(1.5) / 1e10))) - -2.25)) - 0.1)) * 1e10)) * 3.14159)) + 3.14159) + (((1.5 / 3.14159) * (1e10 / (0.1 * -2.25))) + (float32(0.0) * -2.25)))
                                        _ = v81
                                    }
                                case -515:
                                    for fi82 := 0; fi82 < 10; fi82++ {
                                        fmt.Println(v1, fi22)
                                        _ = fi82
                                    }
                                }
                                fmt.Println(v20, v3, fi22)
                                _ = v72
                            }
                        }
                    }
                }
                v83 := GenS7{F0: 170}
                _ = v83
            }
        }
        if ((-(v7)) != (v7 + (v7 >> (((v20 % (((v20 + uint(v1))) | 1))) & 63)))) {
            delete(v5, 1)
            var v84 int16 = 42
            _ = v84
        } else {
            v3 = 1
            if false {
                v85 := GenS3{F0: v1, F1: 815, F2: 4294967295}
                _ = v85
            } else {
                v20 = v20
                v7 -= 303
            }
            for fi86 := 0; fi86 < 16; fi86++ {
                v20 = v20
                _ = fi86
            }
        }
        switch (fi22 / (((fi22 + fi22)) | 1)) {
        case 465:
            v87 := v5[65535]
            _ = v87
        case 2147483647:
            fmt.Println(v20, v7)
        case 42:
            for fi88 := 0; fi88 < 10; fi88++ {
                v6 = append(v6, "foo bar")
                v20 += v20
                v2 += 128
                _ = fi88
            }
        }
        _ = fi22
    }
    v1 -= v1
    var v89 int = 53
    for fi90 := 0; fi90 < 1; fi90++ {
        var v91 float64 = 0.1
        switch v89 {
        case 4294967295:
            _ = len(v6)
        case 976:
            for fi92 := 0; fi92 < 17; fi92++ {
                var v93 bool = true
                fmt.Println(v7)
                v94 := GenS8{F0: "\n\t"}
                _ = fi92
                _ = v93
                _ = v94
            }
        default:
            var v95 uint16 = 255
            _ = v95
        }
        _ = fi90
        _ = v91
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v89
}
