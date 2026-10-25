package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint16; F1 uint; F2 int32 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint32; F1 int64 }
type GenS12 struct { F0 int }

func main() {
    var v1 uint = 1
    v2 := []uint32{}
    v3 := []int{0, -1, 32768}
    var v4 int = 266
    var v5 uint32 = 42
    var v6 uint8 = 207
    var v7 bool = (0 < 32767)
    if true {
        var v8 float64 = 0.0
        switch 2147483647 {
        case 87:
            switch (v4 * (v4 | 32768)) {
            case 136:
                if (v4 >= v4) {
                    v1 = (v1 | (v1 % ((v1) | 1)))
                    var v9 int32 = 1
                    var v10 string = "hello"
                    _ = v10
                    _ = v9
                }
            case -2147483648:
                for fi11 := 0; fi11 < 6; fi11++ {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        var v13 uint16 = 65535
                        _ = fi12
                        _ = v13
                    }
                    var v14 uint = 65535
                    v5 = v5
                    _ = fi11
                    _ = v14
                }
            case -407:
                v5 = (^(v5))
            default:
                fmt.Println(v1)
            }
        case 295:
            var v15 uint16 = 1
            _ = v15
        }
        _ = v8
    } else {
        v1 = uint(v5)
        v16 := GenS1{F0: (v5 >> ((v1) & 63))}
        v5 += v5
        _ = v16
    }
    v17 := GenS2{F0: 532, F1: (v1 >> (((v1 + (v1 % (((v1 >> ((v1) & 63))) | 1)))) & 63)), F2: 1}
    if false {
        switch v4 {
        case 1:
            for fi18 := 0; fi18 < 13; fi18++ {
                if false {
                    var v19 bool = ((^(fi18)) != (fi18 * 127))
                    if (((float64(float32(float64(float32(0.0)))) * ((3.14159 - float64(1e10)) - 3.14159)) <= (-2.25 * 0.1)) || (!false)) {
                        var v20 int = (^(v4))
                        _ = v20
                    } else {
                        v21 := GenS3{F0: 6, F1: ("a" + "hello")}
                        _ = v21
                    }
                    var v22 bool = v7
                    _ = v19
                    _ = v22
                } else {
                    var v23 int32 = -413
                    if v7 {
                        fmt.Println(fi18, v5, v4)
                        if ("\n\t" <= "") {
                            var v24 int = v4
                            for fi25 := 0; fi25 < 7; fi25++ {
                                v2 = append(v2, (v5 / ((2) | 1)))
                                var v26 uint16 = 856
                                if (42 < fi18) {
                                    fmt.Println(v4)
                                    v7 = false
                                } else {
                                    v23 = v23
                                }
                                var v27 uint = 0
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            _ = v24
                        }
                        var v28 int32 = v23
                        _ = v28
                    } else {
                        var v29 int = (fi18 & int(v23))
                        _ = len(v2)
                        var v30 bool = v7
                        _ = v29
                        _ = v30
                    }
                    _ = v23
                }
                if v7 {
                    v7 = ((v1 == 255) && (false && (!v7)))
                }
                _ = fi18
            }
        default:
            if (4294967295 >= v1) {
                v1 = 165
            }
        }
        if (true && (v1 <= v1)) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if false {
                var v31 int8 = 100
                v4 = 1
                if (v7 && ((v4 - v4) <= v4)) {
                    for fi32 := 0; fi32 < 20; fi32++ {
                        v6 -= v6
                        v5 -= uint32(v1)
                        var v33 int8 = (v31 * 77)
                        _ = fi32
                        _ = v33
                    }
                } else {
                    v4 -= -1
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if v7 {
                        for fi34 := 0; fi34 < 1; fi34++ {
                            for fi35 := 0; fi35 < 15; fi35++ {
                                var v36 float32 = ((-2.25 * 3.14159) - float32(-2.25))
                                _ = fi35
                                _ = v36
                            }
                            switch "hello" {
                            case "foo bar":
                                var v37 uint64 = uint64(v1)
                                _ = v37
                            case "a":
                                v38 := GenS1{F0: (v5 | 0)}
                                _ = v38
                            case "\n\t":
                                v1 = 2
                            }
                            _ = fi34
                        }
                        if v7 {
                            switch "" {
                            case "hello":
                                fmt.Println(v1, v7, v4)
                            default:
                                v5 = 1
                            }
                        }
                        v39 := GenS4{F0: 2, F1: (v1 << ((149) & 63)), F2: 1}
                        _ = v39
                    } else {
                        fmt.Println(v7, v1)
                        if true {
                            fmt.Println(v1, v6)
                            v1 = (^(v1))
                            fmt.Println(v4, v4, v4)
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            for fi40 := 0; fi40 < 8; fi40++ {
                                if len(v3) > 0 {
                                    v3[0] = (fi40 | 69)
                                }
                                var v41 float64 = 3.14159
                                _ = fi40
                                _ = v41
                            }
                            switch v4 {
                            case -9223372036854775808:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                    }
                }
                _ = v31
            } else {
                var v42 int16 = -32768
                switch -875 {
                case 0:
                    switch "hello" {
                    case "foo bar":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                default:
                    v42 = (v42 % (((v42 - 32767)) | 1))
                }
                _ = v42
            }
        } else {
            v43 := GenS5{F0: 354, F1: 504}
            _ = v43
        }
        fmt.Println(v6, v7, v1)
    } else {
        switch v4 {
        case 32768:
            v44 := GenS2{F0: 42, F1: 184, F2: 202}
            _ = v44
        }
    }
    if ((v4 >> ((42) & 63)) <= (v4 & (v4 & v4))) {
        v6 -= 42
    }
    if true {
        v7 = (4294967295 > 2)
        for fi45 := 0; fi45 < 2; fi45++ {
            if (!true) {
                switch fi45 {
                case -731:
                    fmt.Println(v5, v4, v6)
                case 32767:
                    switch ((("hello" + "a") + "hello") + (("\n\t" + "foo bar") + (("hello" + "hello") + "hello"))) {
                    case "a":
                        fmt.Println(fi45, v6, v6)
                    case "foo bar":
                        fmt.Println(v5, fi45, v6)
                    case "hello":
                        fmt.Println(fi45)
                    }
                case 809:
                    v46 := GenS4{F0: 2, F1: (v1 * 345), F2: 0}
                    _ = v46
                default:
                    if v7 {
                        v7 = ((3.14159 + 3.14159) < (3.14159 / (1.5 / -2.25)))
                        switch ("" + "") {
                        case "\n\t":
                            var v47 string = "hello"
                            _ = v47
                        case "foo bar":
                            v48 := GenS6{F0: 255}
                            _ = v48
                        case "hello":
                            v4 = fi45
                        default:
                            for fi49 := 0; fi49 < 14; fi49++ {
                                v50 := GenS4{F0: 1, F1: (v1 - v1), F2: 0}
                                if true {
                                    if (v5 < 695) {
                                        var v51 float64 = 3.14159
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v51
                                    } else {
                                        v52 := GenS7{F0: 127, F1: v4}
                                        v53 := GenS8{F0: (v6 % (((v6 - (v6 & (v6 % (((v6 << ((v1) & 63))) | 1))))) | 1)), F1: v7, F2: v5}
                                        _ = v52
                                        _ = v53
                                    }
                                    var v54 uint8 = 128
                                    v1 -= v1
                                    var v55 float32 = ((1e10 * 0.1) / 1e10)
                                    switch 503 {
                                    case 32768:
                                        v55 -= v55
                                    case 127:
                                        var v56 bool = v7
                                        _ = v56
                                    case 2147483647:
                                        _ = len(v2)
                                    }
                                    _ = v54
                                    _ = v55
                                }
                                if len(v2) > 0 {
                                    v2[0] = (v5 / ((v5) | 1))
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v57 uint64 = 18446744073709551615
                                _ = fi49
                                _ = v50
                                _ = v57
                            }
                        }
                        v6 -= 255
                    } else {
                        if true {
                            _ = len(v2)
                        } else {
                            switch v4 {
                            case -270:
                                v4 += (fi45 | 1)
                            case -333:
                                if true {
                                    for fi58 := 0; fi58 < 3; fi58++ {
                                        v59 := GenS6{F0: v6}
                                        fmt.Println(v4)
                                        _ = fi58
                                        _ = v59
                                    }
                                    for fi60 := 0; fi60 < 9; fi60++ {
                                        if len(v3) > 0 {
                                            v3[0] = v4
                                        }
                                        v61 := GenS9{F0: (v1 | (^(v1)))}
                                        var v62 uint16 = 266
                                        _ = fi60
                                        _ = v61
                                        _ = v62
                                    }
                                    v63 := GenS2{F0: 561, F1: (v1 << ((288) & 63)), F2: 323}
                                    _ = v63
                                }
                            default:
                                switch v4 {
                                case -613:
                                    switch "a" {
                                    case "foo bar":
                                        _ = len(v3)
                                    default:
                                        for fi64 := 0; fi64 < 2; fi64++ {
                                            switch "hello" {
                                            case "a":
                                                if v7 {
                                                    v65 := GenS1{F0: v5}
                                                    v66 := GenS5{F0: 42, F1: 2}
                                                    _ = v65
                                                    _ = v66
                                                } else {
                                                    v6 = v6
                                                    if v7 {
                                                        switch ("a" + "a") {
                                                        case "\n\t":
                                                            fmt.Println(v5, v1, v1)
                                                        case "":
                                                            v7 = (-32768 <= 42)
                                                        }
                                                    } else {
                                                        var v67 uint32 = (v5 & v5)
                                                        v4 -= 9223372036854775807
                                                        var v68 uint32 = (v5 >> (((^(v1))) & 63))
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                }
                                            case "\n\t":
                                                if v7 {
                                                    var v69 int64 = int64(v5)
                                                    if v7 {
                                                        if false {
                                                            var v70 int8 = 42
                                                            var v71 uint8 = 128
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v70
                                                            _ = v71
                                                        } else {
                                                            var v72 string = ("a" + "\n\t")
                                                            _ = v72
                                                        }
                                                    }
                                                    fmt.Println(v5, v5, v69)
                                                    _ = v69
                                                } else {
                                                    var v73 uint16 = 1
                                                    _ = v73
                                                }
                                            default:
                                                var v74 uint32 = 0
                                                _ = v74
                                            }
                                            var v75 int32 = 0
                                            v76 := GenS10{F0: 42}
                                            _ = fi64
                                            _ = v75
                                            _ = v76
                                        }
                                    }
                                case -2147483648:
                                    if (!v7) {
                                        for fi77 := 0; fi77 < 16; fi77++ {
                                            var v78 int16 = 575
                                            var v79 float64 = float64(3.14159)
                                            v7 = (4294967295 >= (^(v1)))
                                            _ = fi77
                                            _ = v78
                                            _ = v79
                                        }
                                    }
                                default:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            }
                            v2 = append(v2, v5)
                        }
                    }
                }
            } else {
                v80 := GenS11{F0: v5, F1: 9223372036854775807}
                for fi81 := 0; fi81 < 2; fi81++ {
                    v1 -= v1
                    var v82 int16 = 1000
                    var v83 uint64 = 4294967295
                    _ = fi81
                    _ = v82
                    _ = v83
                }
                var v84 int32 = 0
                _ = v80
                _ = v84
            }
            _ = fi45
        }
        for fi85 := 0; fi85 < 18; fi85++ {
            var v86 string = "hello"
            var v87 bool = v7
            v6 = v6
            _ = fi85
            _ = v86
            _ = v87
        }
    }
    var v88 float32 = (0.1 / (((1.5 - ((3.14159 * (((3.14159 + (1.5 / 1e10)) + float32((1e10 + 3.14159))) / 1.0)) / 3.14159)) + float32(float64(-2.25))) * (-2.25 + 1.5)))
    switch ("foo bar" + "") {
    case "a":
        _ = len(v3)
    case "":
        var v89 uint = 2
        _ = v89
    default:
        switch 32767 {
        case 632:
            v7 = v7
        case -740:
            var v90 uint64 = 4294967295
            _ = v90
        default:
            for fi91 := 0; fi91 < 9; fi91++ {
                v2 = append(v2, 2)
                fmt.Println(v4, v7)
                for fi92 := 0; fi92 < 7; fi92++ {
                    switch fi91 {
                    case 687:
                        var v93 int8 = -128
                        _ = v93
                    case -2147483648:
                        if len(v3) > 0 {
                            v3[0] = 32768
                        }
                    }
                    fmt.Println(v7, v88)
                    switch 0 {
                    case 32768:
                        fmt.Println(v5, v4)
                        var v94 int = -415
                        _ = v94
                    case 0:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -899:
                        var v95 string = ""
                        _ = v95
                    default:
                        v96 := GenS12{F0: v4}
                        _ = v96
                    }
                    _ = fi92
                }
                _ = fi91
            }
        }
    }
    var v97 float64 = 0.0
    if v7 {
        fmt.Println(v6)
    } else {
        for fi98 := 0; fi98 < 11; fi98++ {
            if len(v2) > 0 {
                v2[0] = 65535
            }
            _ = fi98
        }
        fmt.Println(v88, v1)
        if v7 {
            if v7 {
                v97 = ((1e10 * 3.14159) * v97)
            }
            var v99 int32 = int32(v4)
            _ = v99
        } else {
            v4 = (v4 - v4)
            for fi100 := 0; fi100 < 12; fi100++ {
                v6 += v6
                var v101 uint = (v1 / ((v1) | 1))
                _ = fi100
                _ = v101
            }
            var v102 uint64 = 2
            _ = v102
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v88:", v88)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v88
    _ = v97
}
