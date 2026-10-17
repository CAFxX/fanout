package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint16; F2 int }
type GenS4 struct { F0 int8; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint; F1 int8; F2 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 uint32; F2 bool }
type GenS8 struct { F0 uint32; F1 int32 }
type GenS9 struct { F0 uint; F1 int16 }

func main() {
    var v1 float32 = 0.0
    var v2 uint = 18446744073709551615
    v3 := []uint32{65535, 191, 2, 2147483648}
    fmt.Println(v2, v1)
    for fi4 := 0; fi4 < 20; fi4++ {
        switch (fi4 << (((v2 | 42)) & 63)) {
        case 208:
            v1 -= (1.5 * (0.0 - (v1 + (0.1 - v1))))
        }
        _ = fi4
    }
    var v5 int64 = 42
    v1 += float32(float64((1.5 + (((float32((((1.5 + (1.5 + 0.1)) * 1.5) - -2.25)) - v1) * ((float32(float64(float32(1.5))) + float32(3.14159)) - v1)) * 1.5))))
    v5 = v5
    switch "\n\t" {
    case "a":
        for fi6 := 0; fi6 < 6; fi6++ {
            if (42 >= 4294967295) {
                var v7 uint = 1
                v5 -= 9223372036854775807
                _ = v7
            }
            _ = fi6
        }
    case "foo bar":
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                switch fi9 {
                case 42:
                    var v10 int = (fi9 & (fi8 + fi8))
                    var v11 uint32 = 423
                    v12 := GenS1{F0: v2}
                    v2 = 4294967295
                    _ = v10
                    _ = v11
                    _ = v12
                case -1:
                    fmt.Println(fi9, fi9)
                default:
                    switch (fi8 - 9223372036854775807) {
                    case 32767:
                        if ((v2 * 65535) != v2) {
                            switch -526 {
                            case -9223372036854775808:
                                var v13 uint64 = 18446744073709551615
                                _ = v13
                            case 128:
                                v14 := GenS1{F0: 1}
                                _ = v14
                            case -2147483648:
                                switch -632 {
                                case 737:
                                    if len(v3) > 0 {
                                        v3[0] = 2147483648
                                    }
                                case 127:
                                    var v15 float32 = 3.14159
                                    if (334 > 0) {
                                        if (2 != 140) {
                                            fmt.Println(v1, v5)
                                        } else {
                                            v16 := GenS1{F0: v2}
                                            v1 = 1e10
                                            _ = v16
                                        }
                                        fmt.Println(v2, v15)
                                        if true {
                                            v17 := GenS1{F0: 166}
                                            var v18 uint8 = 42
                                            var v19 int16 = -1
                                            _ = v17
                                            _ = v18
                                            _ = v19
                                        } else {
                                            var v20 int32 = int32(v5)
                                            v21 := GenS2{F0: 1000}
                                            fmt.Println(v20, v20)
                                            _ = v20
                                            _ = v21
                                        }
                                        switch (fi9 & -944) {
                                        case 127:
                                            if ((-68 == -5) && true) {
                                                if true {
                                                    if false {
                                                        var v22 float64 = 0.0
                                                        _ = v22
                                                    } else {
                                                        if false {
                                                            var v23 uint32 = 4294967295
                                                            _ = v23
                                                        }
                                                        fmt.Println(v15, v1, v1)
                                                        v5 -= (v5 * (v5 - v5))
                                                    }
                                                    if false {
                                                        var v24 int64 = (v5 >> (((v2 / ((2) | 1))) & 63))
                                                        if false {
                                                            var v25 uint16 = 65535
                                                            fmt.Println(v24)
                                                            v2 = (v2 + v2)
                                                            _ = v25
                                                        } else {
                                                            v26 := GenS1{F0: v2}
                                                            _ = v26
                                                        }
                                                        v15 += ((v1 * (v1 - (1e10 - 0.0))) - v15)
                                                        _ = v24
                                                    } else {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        switch fi9 {
                                                        case 0:
                                                            fmt.Println(v1, v2, fi9)
                                                        case 32767:
                                                            var v27 bool = true
                                                            _ = v27
                                                        }
                                                    }
                                                }
                                                if (1e10 < v1) {
                                                    var v28 float64 = 0.1
                                                    _ = v28
                                                }
                                                switch "hello" {
                                                case "foo bar":
                                                    v2 = (v2 / ((90) | 1))
                                                    var v29 int = fi9
                                                    _ = v29
                                                case "a":
                                                    if (v1 == 1e10) {
                                                        if (9 < 128) {
                                                            fmt.Println(fi8)
                                                            v30 := GenS3{F0: "foo bar", F1: 203, F2: fi9}
                                                            var v31 uint32 = 2147483648
                                                            _ = v30
                                                            _ = v31
                                                        } else {
                                                            var v32 int8 = -106
                                                            if true {
                                                                v15 += v15
                                                            } else {
                                                                var v33 uint8 = 65
                                                                _ = v33
                                                            }
                                                            v32 -= (v32 << (((v2 & v2)) & 63))
                                                            v34 := GenS3{F0: "foo bar", F1: 255, F2: (fi9 - (fi8 << ((418) & 63)))}
                                                            var v35 int = fi9
                                                            _ = v32
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                        var v36 string = "\n\t"
                                                        switch (-(fi9)) {
                                                        case 127:
                                                            switch "" {
                                                            case "a":
                                                                var v37 uint = v2
                                                                _ = v37
                                                            case "foo bar":
                                                                fmt.Println(v15, v36, v36)
                                                                var v38 uint8 = 127
                                                                _ = v38
                                                            }
                                                        case -9223372036854775808:
                                                            v3 = append(v3, 2147483648)
                                                        default:
                                                            v1 = v15
                                                            fmt.Println(v1)
                                                            v39 := GenS4{F0: 42, F1: 1, F2: 2147483647}
                                                            _ = v39
                                                        }
                                                        _ = v36
                                                    }
                                                case "hello":
                                                    var v40 int16 = -300
                                                    _ = v40
                                                }
                                                if true {
                                                    v15 -= 1.5
                                                } else {
                                                    if (42 > 155) {
                                                        var v41 float32 = v15
                                                        if false {
                                                            v42 := GenS1{F0: 42}
                                                            if true {
                                                                var v43 float32 = -2.25
                                                                var v44 bool = false
                                                                v5 -= -287
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                            switch ((("\n\t" + "hello") + ("a" + "\n\t")) + "foo bar") {
                                                            case "hello":
                                                                switch fi9 {
                                                                case 0:
                                                                    var v45 int8 = 100
                                                                    _ = v45
                                                                case 32768:
                                                                    v46 := GenS1{F0: v2}
                                                                    _ = v46
                                                                case -2147483648:
                                                                    var v47 uint16 = 169
                                                                    _ = v47
                                                                }
                                                            case "":
                                                                v48 := GenS1{F0: 856}
                                                                _ = v48
                                                            case "a":
                                                                if true {
                                                                    v2 -= (v2 ^ v2)
                                                                    if true {
                                                                        v49 := GenS5{F0: v2, F1: 127, F2: 47}
                                                                        var v50 int32 = -145
                                                                        _ = v49
                                                                        _ = v50
                                                                    }
                                                                    if ((v5 ^ (v5 & int64(v2))) >= v5) {
                                                                        var v51 int = 4294967295
                                                                        _ = v51
                                                                    } else {
                                                                        v1 = 1e10
                                                                    }
                                                                    var v52 int8 = 43
                                                                    _ = v52
                                                                }
                                                            }
                                                            _ = len(v3)
                                                            v2 -= 18446744073709551615
                                                            _ = v42
                                                        } else {
                                                            switch ((("hello" + "a") + "") + "") {
                                                            case "hello":
                                                                if true {
                                                                    v15 -= v1
                                                                }
                                                            case "\n\t":
                                                                var v53 string = "foo bar"
                                                                var v54 string = ""
                                                                _ = v53
                                                                _ = v54
                                                            case "a":
                                                                v55 := GenS6{F0: -114}
                                                                _ = v55
                                                            }
                                                            fmt.Println(v1, v5)
                                                            var v56 int32 = 2147483647
                                                            _ = v56
                                                        }
                                                        _ = v41
                                                    }
                                                    v2 -= (v2 / ((65535) | 1))
                                                }
                                            } else {
                                                switch "a" {
                                                case "":
                                                    v57 := GenS7{F0: 65535, F1: 612, F2: (2 != 1)}
                                                    _ = v57
                                                case "hello":
                                                    var v58 string = ("" + "a")
                                                    _ = v58
                                                }
                                            }
                                        case -886:
                                            fmt.Println(v2, fi8)
                                            if len(v3) > 0 {
                                                v3[0] = 272
                                            }
                                        }
                                        v59 := GenS8{F0: 0, F1: 0}
                                        var v60 int16 = -853
                                        var v61 string = ""
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                    }
                                    var v62 uint16 = 1
                                    _ = v15
                                    _ = v62
                                case 9223372036854775807:
                                    v1 += (v1 / 1.0)
                                    switch fi9 {
                                    case 252:
                                        v2 -= (^(v2))
                                    case 127:
                                        v63 := GenS9{F0: 2, F1: 42}
                                        _ = v63
                                    default:
                                        v64 := GenS8{F0: 42, F1: -1}
                                        _ = v64
                                    }
                                    fmt.Println(fi8)
                                default:
                                    if true {
                                        v1 = v1
                                        var v65 uint16 = 0
                                        _ = v65
                                    }
                                }
                            }
                        }
                    case 32768:
                        v1 += (v1 / v1)
                    }
                }
                v3 = append(v3, 701)
                _ = fi9
            }
            var v66 int32 = 1
            _ = fi8
            _ = v66
        }
    }
    var v67 uint16 = 819
    switch "a" {
    case "hello":
        switch "\n\t" {
        case "hello":
            var v68 uint32 = 825
            _ = v68
        case "a":
            v69 := GenS5{F0: 18446744073709551615, F1: -128, F2: 521}
            _ = v69
        default:
            var v70 bool = (true && (v1 != 3.14159))
            _ = v70
        }
    case "foo bar":
        fmt.Println(v67, v67, v67)
    default:
        for fi71 := 0; fi71 < 3; fi71++ {
            var v72 uint = (v2 >> ((v2) & 63))
            fmt.Println(v72, fi71, v2)
            fmt.Println(v67)
            _ = fi71
            _ = v72
        }
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    v1 = 1.5
    for fi73 := 0; fi73 < 12; fi73++ {
        v3 = append(v3, 2)
        v1 += (1.5 + ((v1 / 1.0) + 0.1))
        _ = fi73
    }
    var v74 int = -9223372036854775808
    var v75 int64 = 190
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v67
    _ = v74
    _ = v75
}
