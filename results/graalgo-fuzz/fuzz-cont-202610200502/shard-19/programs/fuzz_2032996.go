package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 string }
func genFn0(p0 int64) int8 {
    var v7 uint64 = 1
    v8 := GenS1{F0: ("a" + "hello"), F1: "", F2: -1}
    fmt.Println(v7, v7)
    v7 = v7
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1() bool {
    v9 := []int{-9223372036854775808}
    for fi10 := 0; fi10 < 13; fi10++ {
        if (3.14159 == -2.25) {
            var v11 uint = 420
            if (0 > 761) {
                v11 = 649
                var v12 bool = (-1 != 208)
                _ = v12
            }
            _ = v11
        }
        v13 := GenS1{F0: "\n\t", F1: "hello", F2: 2147483647}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi10
        _ = v13
    }
    var v14 int8 = -128
    if (-1 > -9223372036854775808) {
        for fi15 := 0; fi15 < 14; fi15++ {
            v14 += v14
            _ = fi15
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 uint = 2
                _ = fi17
                _ = v18
            }
            v14 = (v14 << ((410) & 63))
            fmt.Println(v14)
            _ = fi16
        }
        if true {
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 uint = 255
                v14 -= v14
                if (!false) {
                    switch ("hello" + "") {
                    case "foo bar":
                        if true {
                            if (!(47 > 42)) {
                                fmt.Println(v14, v20, v20)
                                v9 = append(v9, fi19)
                                for fi21 := 0; fi21 < 12; fi21++ {
                                    var v22 uint32 = 0
                                    if true {
                                        var v23 uint8 = 128
                                        _ = v23
                                    } else {
                                        v22 = (v22 / ((544) | 1))
                                        var v24 uint8 = 132
                                        if (3.14159 == ((0.0 / 3.14159) * 3.14159)) {
                                            v14 = v14
                                            if true {
                                                v14 = 0
                                            }
                                        } else {
                                            if (708 == uint64(v20)) {
                                                fmt.Println(fi21, v20, v24)
                                            }
                                            fmt.Println(v24, v14)
                                            v22 = (^(v22))
                                        }
                                        _ = v24
                                    }
                                    _ = fi21
                                    _ = v22
                                }
                            }
                            switch ("\n\t" + "hello") {
                            case "\n\t":
                                for fi25 := 0; fi25 < 13; fi25++ {
                                    if len(v9) > 0 {
                                        v9[0] = (fi25 / ((1) | 1))
                                    }
                                    var v26 uint = v20
                                    v26 = v26
                                    switch ((("\n\t" + ("a" + "hello")) + "\n\t") + "\n\t") {
                                    case "a":
                                        if (!false) {
                                            v27 := GenS2{F0: 1}
                                            var v28 uint64 = 2
                                            if ((true && false) || (0 > 0)) {
                                                if true {
                                                    v29 := GenS1{F0: ("hello" + "foo bar"), F1: ("" + ("foo bar" + "hello")), F2: 1}
                                                    fmt.Println(fi19, v20, v14)
                                                    _ = v29
                                                } else {
                                                    v14 -= genFn0(1)
                                                    v20 = v26
                                                }
                                                v30 := GenS2{F0: 278}
                                                _ = v30
                                            } else {
                                                v28 = (v28 - v28)
                                            }
                                            _ = v27
                                            _ = v28
                                        }
                                    case "foo bar":
                                        _ = len(v9)
                                    case "hello":
                                        if len(v9) > 0 {
                                            v9[0] = fi25
                                        }
                                    default:
                                        switch fi19 {
                                        case 2147483647:
                                            v31 := GenS1{F0: ("foo bar" + ("" + "\n\t")), F1: "hello", F2: -54}
                                            _ = v31
                                        case 1:
                                            fmt.Println(fi19, v26)
                                        case -1:
                                            v32 := GenS1{F0: "hello", F1: "\n\t", F2: 2147483647}
                                            switch "\n\t" {
                                            case "foo bar":
                                                v14 += genFn0(592)
                                            default:
                                                v33 := GenS2{F0: 2147483647}
                                                _ = v33
                                            }
                                            _ = v32
                                        }
                                    }
                                    _ = fi25
                                    _ = v26
                                }
                            }
                            var v34 uint32 = uint32(v20)
                            _ = v34
                        } else {
                            switch "" {
                            case "a":
                                if (true || (0.0 <= 1e10)) {
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                } else {
                                    fmt.Println(fi19, v20, v20)
                                }
                            case "\n\t":
                                var v35 int8 = v14
                                _ = v35
                            case "":
                                for fi36 := 0; fi36 < 3; fi36++ {
                                    v9 = append(v9, (^(fi36)))
                                    v20 = 0
                                    _ = len(v9)
                                    _ = fi36
                                }
                            default:
                                if true {
                                    var v37 bool = false
                                    for fi38 := 0; fi38 < 16; fi38++ {
                                        var v39 int16 = -802
                                        var v40 uint = v20
                                        _ = fi38
                                        _ = v39
                                        _ = v40
                                    }
                                    if false {
                                        switch fi19 {
                                        case 32768:
                                            var v41 int8 = 100
                                            _ = v41
                                        case 0:
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                        case -1:
                                            var v42 uint = (v20 & (v20 & v20))
                                            _ = v42
                                        default:
                                            v14 -= v14
                                        }
                                        for fi43 := 0; fi43 < 9; fi43++ {
                                            fmt.Println(v20, v20)
                                            if v37 {
                                                if len(v9) > 0 {
                                                    v9[0] = fi43
                                                }
                                                v44 := GenS2{F0: 1}
                                                _ = v44
                                            }
                                            var v45 uint64 = 42
                                            _ = fi43
                                            _ = v45
                                        }
                                    } else {
                                        var v46 int32 = int32(v20)
                                        _ = v46
                                    }
                                    _ = v37
                                } else {
                                    v20 -= (v20 | v20)
                                    if len(v9) > 0 {
                                        v9[0] = int(v20)
                                    }
                                }
                            }
                            var v47 uint32 = 4294967295
                            var v48 int16 = 456
                            _ = v47
                            _ = v48
                        }
                    default:
                        if ("hello" >= ("foo bar" + ((("hello" + "foo bar") + "hello") + ("a" + "hello")))) {
                            v49 := GenS3{F0: 0}
                            v14 -= v14
                            _ = v49
                        }
                    }
                } else {
                    for fi50 := 0; fi50 < 18; fi50++ {
                        v51 := GenS2{F0: int32(v20)}
                        v52 := GenS3{F0: 9223372036854775807}
                        if (-277 == 815) {
                            var v53 int16 = 1000
                            if (!false) {
                                if (!false) {
                                    var v54 int16 = -1
                                    var v55 float32 = 1.5
                                    switch (("" + "foo bar") + "\n\t") {
                                    case "foo bar":
                                        var v56 int = (-(fi50))
                                        _ = v56
                                    case "":
                                        var v57 uint32 = 1
                                        _ = v57
                                    default:
                                        var v58 int32 = 1
                                        _ = v58
                                    }
                                    switch 32767 {
                                    case 4294967295:
                                        var v59 uint = 18446744073709551615
                                        if len(v9) > 0 {
                                            v9[0] = fi19
                                        }
                                        _ = v59
                                    case 42:
                                        if len(v9) > 0 {
                                            v9[0] = (-(fi50))
                                        }
                                    }
                                    switch 42 {
                                    case 128:
                                        if true {
                                            var v60 uint8 = 0
                                            _ = v60
                                        }
                                    case 1:
                                        var v61 float32 = 1e10
                                        _ = v61
                                    }
                                    _ = v54
                                    _ = v55
                                }
                                v14 -= (v14 - (v14 >> (((v20 + v20)) & 63)))
                                fmt.Println(fi19)
                                fmt.Println(fi19)
                                v14 -= 127
                            }
                            _ = v53
                        }
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    v62 := GenS2{F0: 1}
                    v63 := GenS1{F0: ("hello" + "\n\t"), F1: "", F2: 42}
                    _ = v62
                    _ = v63
                }
                _ = fi19
                _ = v20
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            var v64 int32 = 320
            _ = v64
        } else {
            var v65 float32 = 1.5
            _ = v65
        }
    } else {
        var v66 int = 1
        if len(v9) > 0 {
            v9[0] = v66
        }
        _ = v66
    }
    fmt.Println(v14)
    _ = v14
    _ = v9
    return false
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = -2147483648
    var v3 uint = 42
    v4 := make(map[int]uint32)
    var v5 bool = false
    v6 := make(map[string]string)
    v67 := GenS4{F0: 9223372036854775807, F1: 1, F2: 42}
    v5 = (-100 < -128)
    var v68 int16 = -1
    v3 = (v3 % (((v3 / (((v3 + v3)) | 1))) | 1))
    switch 4294967295 {
    case 4294967295:
        switch "hello" {
        case "\n\t":
            v5 = v5
        }
    case -14:
        for fi69 := 0; fi69 < 14; fi69++ {
            if v5 {
                v5 = true
                for fi70 := 0; fi70 < 14; fi70++ {
                    v2 -= v2
                    fmt.Println(v68, v2)
                    if v5 {
                        var v71 uint8 = 96
                        if (!v5) {
                            v72 := GenS5{F0: (v68 | int16(v2))}
                            _ = v72
                        } else {
                            v71 -= (v71 + (v71 % (((v71 * (v71 >> ((v3) & 63)))) | 1)))
                        }
                        _ = v71
                    }
                    _ = fi70
                }
            }
            _ = fi69
        }
    case -361:
        fmt.Println(v68, v3)
    }
    v73 := []int{2147483647}
    v2 -= (v2 * 2147483647)
    v74 := GenS3{F0: 371}
    fmt.Println(v2, v2, v5)
    v75 := GenS6{F0: v5}
    v76 := GenS1{F0: "hello", F1: "", F2: (v2 << ((255) & 63))}
    if ((true || (0 >= uint8(v3))) && false) {
        for fi77 := 0; fi77 < 11; fi77++ {
            var v78 string = ""
            fmt.Println(v68)
            _ = fi77
            _ = v78
        }
        v5 = (!(0 <= 0))
        if v5 {
            switch -2147483648 {
            case -446:
                var v79 string = ("foo bar" + "foo bar")
                _ = v79
            }
            v80 := GenS7{F0: (^(v3)), F1: 42}
            if v5 {
                fmt.Println(v5, v3, v68)
                for fi81 := 0; fi81 < 5; fi81++ {
                    var v82 int = 2147483647
                    var v83 bool = v5
                    _ = fi81
                    _ = v82
                    _ = v83
                }
            }
            _ = v80
        } else {
            v3 = v3
            var v84 uint16 = 65535
            if v5 {
                var v85 int8 = -1
                _ = v85
            } else {
                var v86 uint64 = 1
                _ = v86
            }
            _ = v84
        }
    }
    switch "" {
    case "a":
        for fi87 := 0; fi87 < 7; fi87++ {
            v2 = v2
            v68 -= (v68 / ((-1) | 1))
            _ = fi87
        }
    }
    switch "hello" {
    case "":
        if ((!(-835 == -143)) && (1 <= 128)) {
            var v88 float64 = 0.0
            v3 += v3
            for fi89 := 0; fi89 < 8; fi89++ {
                v2 = 792
                _ = fi89
            }
            _ = v88
        }
    case "hello":
        fmt.Println(v2, v2)
    case "foo bar":
        if (0 <= 255) {
            switch -125 {
            case -9223372036854775808:
                v68 = 32767
            case 849:
                for fi90 := 0; fi90 < 7; fi90++ {
                    if (v5 || ((v68 / ((v68) | 1)) <= (v68 >> ((v3) & 63)))) {
                        for fi91 := 0; fi91 < 1; fi91++ {
                            v92 := GenS2{F0: -1}
                            v5 = v5
                            _ = fi91
                            _ = v92
                        }
                        if v5 {
                            for fi93 := 0; fi93 < 9; fi93++ {
                                switch "hello" {
                                case "foo bar":
                                    var v94 int64 = 704
                                    _ = v94
                                case "":
                                    if v5 {
                                        v2 += v2
                                        v95 := GenS1{F0: "\n\t", F1: "\n\t", F2: 2147483647}
                                        v96 := GenS3{F0: -557}
                                        _ = v95
                                        _ = v96
                                    }
                                case "hello":
                                    _ = len(v73)
                                }
                                var v97 int16 = (v68 & v68)
                                v98 := v1[("" + ("foo bar" + ("hello" + "foo bar")))]
                                _ = fi93
                                _ = v97
                                _ = v98
                            }
                            _ = len(v1)
                        }
                    }
                    _ = fi90
                }
            case -458:
                for fi99 := 0; fi99 < 5; fi99++ {
                    fmt.Println(v3, fi99, v3)
                    v100 := GenS3{F0: 1}
                    for fi101 := 0; fi101 < 7; fi101++ {
                        if (false && v5) {
                            if false {
                                var v102 uint8 = 127
                                if (v5 || ((fi101 - fi99) == 9223372036854775807)) {
                                    switch (fi101 | fi99) {
                                    case -520:
                                        v102 -= (v102 >> ((0) & 63))
                                    }
                                    if true {
                                        _ = len(v1)
                                        if v5 {
                                            fmt.Println(v3, fi101)
                                        }
                                    } else {
                                        v102 += (v102 | (v102 >> (((^(v3))) & 63)))
                                    }
                                } else {
                                    v5 = (v5 && true)
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    v103 := GenS3{F0: -384}
                                    _ = v103
                                case "a":
                                    v104 := GenS5{F0: (v68 | (v68 ^ (v68 % (((^(v68))) | 1))))}
                                    _ = v104
                                case "hello":
                                    if ((("foo bar" + ("hello" + ("\n\t" + ("\n\t" + "\n\t")))) != "hello") && false) {
                                        var v105 float32 = -2.25
                                        var v106 int32 = v2
                                        var v107 int16 = v68
                                        v108 := GenS6{F0: (v5 && v5)}
                                        _ = v105
                                        _ = v106
                                        _ = v107
                                        _ = v108
                                    } else {
                                        v109 := v1["a"]
                                        v2 = v2
                                        _ = v109
                                    }
                                }
                                var v110 uint64 = 2
                                if true {
                                    fmt.Println(fi99, v102)
                                } else {
                                    var v111 float32 = (float32(((-2.25 * (-2.25 * 0.1)) / (1e10 / 1.5))) * -2.25)
                                    var v112 int16 = (^(v68))
                                    fmt.Println(v110, v112, fi99)
                                    _ = v111
                                    _ = v112
                                }
                                _ = v102
                                _ = v110
                            }
                        }
                        _ = fi101
                    }
                    _ = fi99
                    _ = v100
                }
            default:
                v113 := GenS8{F0: (v5 && v5), F1: 65535}
                _ = v113
            }
            v114 := v6["hello"]
            _ = v114
        } else {
            v115 := GenS3{F0: 0}
            for fi116 := 0; fi116 < 17; fi116++ {
                switch fi116 {
                case -938:
                    v117 := GenS1{F0: "hello", F1: ("hello" + ""), F2: v2}
                    _ = v117
                case -9223372036854775808:
                    var v118 uint = 0
                    _ = v118
                default:
                    v5 = true
                }
                var v119 uint32 = 65535
                _ = fi116
                _ = v119
            }
            v120 := GenS2{F0: v2}
            _ = v115
            _ = v120
        }
    }
    switch 1 {
    case 127:
        var v121 string = ("hello" + "\n\t")
        _ = v121
    case 32767:
        var v122 int16 = 1
        _ = v122
    }
    for fi123 := 0; fi123 < 1; fi123++ {
        v2 = (v2 * (^(v2)))
        if (42 <= 53) {
            switch 32767 {
            case -839:
                for fi124 := 0; fi124 < 19; fi124++ {
                    fmt.Println(fi124, v2)
                    v125 := GenS9{F0: ((("hello" + "hello") + (("a" + "\n\t") + "\n\t")) + "hello")}
                    v126 := GenS2{F0: v2}
                    _ = fi124
                    _ = v125
                    _ = v126
                }
            default:
                var v127 int64 = 1
                _ = v127
            }
            delete(v4, 4294967295)
        } else {
            var v128 float64 = float64(1e10)
            fmt.Println(fi123, fi123, fi123)
            _ = v128
        }
        _ = fi123
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v73:", len(v73), v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v73
    _ = v74
    _ = v75
    _ = v76
}
