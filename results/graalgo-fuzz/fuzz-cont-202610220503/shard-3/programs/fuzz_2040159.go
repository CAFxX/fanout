package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS10 struct { F0 uint16; F1 int; F2 int32 }
type GenS11 struct { F0 bool; F1 uint32 }
type GenS12 struct { F0 string; F1 uint; F2 int8 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int; F1 uint64 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 float32 = 0.0
        _ = fi4
        _ = v5
    }
    switch "a" {
    case "hello":
        var v6 float64 = 1.5
        _ = v6
    default:
        var v7 bool = false
        _ = v7
    }
    switch "\n\t" {
    case "":
        if (((0.1 <= 0.1) || true) && false) {
            switch -1 {
            case 31:
                var v8 uint8 = 42
                _ = v8
            }
            var v9 int32 = -180
            _ = v9
        } else {
            var v10 uint32 = 4294967295
            _ = v10
        }
    case "\n\t":
        var v11 string = "\n\t"
        _ = v11
    case "a":
        if (true && false) {
            v12 := GenS1{F0: 65535}
            v13 := []int{2147483647, 0}
            var v14 bool = true
            _ = v12
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 2; fi15++ {
                if ((!(42 == -1)) && (!false)) {
                    var v16 uint32 = 694
                    _ = v16
                }
                _ = fi15
            }
        }
    }
    return 704
}
func genFn1() int8 {
    var v17 float64 = 3.14159
    v18 := GenS2{F0: 65535, F1: 126, F2: 2}
    v19 := GenS2{F0: 2147483647, F1: 950, F2: 82}
    _ = v17
    _ = v18
    _ = v19
    return 56
}

func main() {
    var v1 int8 = -20
    v2 := []uint32{42}
    var v3 int = 362
    var v20 uint64 = 42
    switch "foo bar" {
    case "":
        v2 = append(v2, 0)
    case "a":
        var v21 int8 = (v1 / ((v1) | 1))
        _ = v21
    case "foo bar":
        v1 = (^(v1))
    default:
        switch v3 {
        case 37:
            var v22 uint32 = 4294967295
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        v3 -= int(v20)
        switch ((("\n\t" + "a") + ("\n\t" + "a")) + "a") {
        case "\n\t":
            for fi24 := 0; fi24 < 11; fi24++ {
                switch v3 {
                case -438:
                    switch "foo bar" {
                    case "hello":
                        var v25 uint16 = 1
                        _ = v25
                    case "a":
                        switch (fi23 * fi24) {
                        case 32767:
                            v3 += (v3 + (v3 % ((32767) | 1)))
                        default:
                            switch -9223372036854775808 {
                            case 9223372036854775807:
                                var v26 uint16 = 32767
                                _ = v26
                            }
                        }
                        fmt.Println(v1, v20, fi23)
                    }
                case 127:
                    switch ((((("" + "hello") + "") + "hello") + "\n\t") + "a") {
                    case "":
                        _ = len(v2)
                    case "\n\t":
                        if (((float32(1.5) >= 1e10) || false) && (!true)) {
                            fmt.Println(fi23)
                        } else {
                            switch (fi23 - (fi23 | fi23)) {
                            case 0:
                                var v27 float64 = ((float64(0.0) / 1.5) - 1e10)
                                _ = v27
                            case 421:
                                v20 += 4294967295
                            default:
                                var v28 uint16 = 649
                                _ = v28
                            }
                            v29 := GenS2{F0: 2147483647, F1: 0, F2: 139}
                            _ = v29
                        }
                    default:
                        var v30 int8 = (v1 >> ((255) & 63))
                        _ = v30
                    }
                    if (false && (4294967295 != 2147483647)) {
                        v20 += v20
                        if (2 <= 740) {
                            var v31 string = ""
                            if true {
                                v32 := GenS2{F0: 2147483647, F1: 65535, F2: 2}
                                _ = v32
                            }
                            _ = v31
                        }
                        if (!false) {
                            v33 := GenS1{F0: 65535}
                            v34 := GenS1{F0: 32767}
                            _ = v33
                            _ = v34
                        } else {
                            var v35 bool = true
                            switch -2147483648 {
                            case -2147483648:
                                _ = len(v2)
                            case -1:
                                v36 := GenS3{F0: v1, F1: genFn0(), F2: 1000}
                                _ = v36
                            }
                            v37 := GenS4{F0: 42}
                            _ = v35
                            _ = v37
                        }
                    }
                    var v38 int32 = 42
                    _ = v38
                default:
                    fmt.Println(v20)
                }
                switch "hello" {
                case "foo bar":
                    fmt.Println(fi24)
                default:
                    v20 -= 18446744073709551615
                }
                switch fi24 {
                case 164:
                    fmt.Println(v20)
                }
                if ((true && false) || (136 <= 2)) {
                    if true {
                        if true {
                            var v39 int16 = -1
                            _ = v39
                        }
                    } else {
                        if true {
                            var v40 uint16 = 1
                            var v41 uint64 = (v20 / ((0) | 1))
                            v42 := GenS1{F0: v40}
                            _ = v40
                            _ = v41
                            _ = v42
                        } else {
                            var v43 int64 = -825
                            _ = v43
                        }
                        v44 := GenS5{F0: 20}
                        _ = v44
                    }
                    v20 = 18446744073709551615
                    if (true && false) {
                        var v45 uint8 = 173
                        v46 := GenS4{F0: 255}
                        fmt.Println(v3, fi23)
                        _ = v45
                        _ = v46
                    } else {
                        v47 := GenS6{F0: -1, F1: (v1 & 0)}
                        v3 = (v3 / ((fi24) | 1))
                        _ = v47
                    }
                } else {
                    if (-904 > -232) {
                        if false {
                            switch fi23 {
                            case 680:
                                v3 += 128
                            case 127:
                                v2 = append(v2, 530)
                            case 4294967295:
                                var v48 uint32 = 4294967295
                                _ = v48
                            default:
                                v3 -= fi24
                            }
                            v2 = append(v2, 42)
                            var v49 bool = (42 >= 2)
                            _ = v49
                        }
                        v3 = fi24
                        if (!true) {
                            _ = len(v2)
                        } else {
                            var v50 int64 = 951
                            switch -1 {
                            case -265:
                                fmt.Println(v50, v1)
                                if (v50 > v50) {
                                    v51 := GenS7{F0: 93}
                                    v52 := GenS8{F0: ("foo bar" + ("" + ("foo bar" + "hello")))}
                                    fmt.Println(v1, v3)
                                    var v53 uint8 = 163
                                    var v54 int16 = 32767
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                                var v55 int32 = 0
                                var v56 uint8 = 1
                                fmt.Println(v20, v50)
                                _ = v55
                                _ = v56
                            case 0:
                                fmt.Println(v3)
                                var v57 float64 = 1e10
                                if (1 != 4294967295) {
                                    if (!(!true)) {
                                        v1 = 100
                                        switch (("\n\t" + ((("foo bar" + "") + "a") + "\n\t")) + "\n\t") {
                                        case "a":
                                            switch ("" + ("" + ((("hello" + "") + "") + ("hello" + ("a" + ("hello" + "\n\t")))))) {
                                            case "hello":
                                                _ = len(v2)
                                            default:
                                                switch v3 {
                                                case -811:
                                                    var v58 int = int(v50)
                                                    _ = v58
                                                case -2147483648:
                                                    var v59 uint16 = 32767
                                                    _ = v59
                                                }
                                            }
                                        }
                                        v60 := GenS9{F0: 0, F1: (v1 << ((795) & 63)), F2: 1}
                                        _ = v60
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            switch 32768 {
                                            case -101:
                                                if true {
                                                    var v61 uint = 913
                                                    _ = v61
                                                } else {
                                                    fmt.Println(v3, v20, fi24)
                                                    v2 = append(v2, genFn0())
                                                }
                                            default:
                                                var v62 float32 = 0.1
                                                _ = v62
                                            }
                                        }
                                        fmt.Println(v57)
                                    }
                                    v2 = append(v2, 522)
                                    v57 -= v57
                                } else {
                                    if false {
                                        fmt.Println(v57)
                                    } else {
                                        var v63 float64 = (v57 / 1.0)
                                        _ = v63
                                    }
                                }
                                _ = v57
                            }
                            _ = v50
                        }
                    } else {
                        v1 -= v1
                        v64 := GenS10{F0: 42, F1: 0, F2: -2147483648}
                        _ = v64
                    }
                    v65 := GenS6{F0: -32768, F1: v1}
                    v66 := GenS4{F0: 0}
                    _ = v65
                    _ = v66
                }
                _ = fi24
            }
        case "hello":
            switch (v3 ^ (-(v3))) {
            case 4294967295:
                for fi67 := 0; fi67 < 20; fi67++ {
                    var v68 uint = 0
                    _ = fi67
                    _ = v68
                }
            case -546:
                switch (((((("\n\t" + "") + "hello") + "") + "") + "hello") + "foo bar") {
                case "hello":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "a":
                    for fi69 := 0; fi69 < 1; fi69++ {
                        v70 := GenS11{F0: (-2147483648 != 0), F1: 42}
                        var v71 uint8 = 42
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                case "":
                    v2 = append(v2, 692)
                default:
                    var v72 uint32 = 433
                    _ = v72
                }
            default:
                v2 = append(v2, 742)
            }
        case "":
            for fi73 := 0; fi73 < 8; fi73++ {
                v1 = v1
                _ = fi73
            }
        }
        v3 = (fi23 ^ (fi23 >> ((795) & 63)))
        _ = fi23
    }
    var v74 int = v3
    v3 = v3
    fmt.Println(v3)
    v1 -= v1
    var v75 int16 = 0
    for fi76 := 0; fi76 < 7; fi76++ {
        v77 := GenS2{F0: 4294967295, F1: 2, F2: 127}
        v78 := GenS12{F0: ("a" + "hello"), F1: 2, F2: v1}
        if false {
            var v79 uint32 = 182
            v1 = v1
            var v80 string = ""
            _ = v79
            _ = v80
        } else {
            if false {
                v81 := GenS12{F0: "\n\t", F1: 2, F2: v1}
                _ = v81
            } else {
                v3 += 1
                var v82 float64 = float64(1.5)
                _ = v82
            }
            v83 := GenS10{F0: 2, F1: -444, F2: 2147483647}
            _ = v83
        }
        _ = fi76
        _ = v77
        _ = v78
    }
    switch int(v75) {
    case -616:
        switch 1 {
        case -9223372036854775808:
            v84 := GenS8{F0: "hello"}
            _ = v84
        case 484:
            var v85 uint = 42
            _ = v85
        default:
            v2 = append(v2, 2147483647)
        }
    }
    v1 += 1
    for fi86 := 0; fi86 < 11; fi86++ {
        v3 += (v3 << ((1) & 63))
        for fi87 := 0; fi87 < 8; fi87++ {
            v2 = append(v2, genFn0())
            switch "foo bar" {
            case "hello":
                var v88 string = "\n\t"
                _ = v88
            }
            _ = fi87
        }
        for fi89 := 0; fi89 < 6; fi89++ {
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            if (0 >= 65535) {
                v20 += (v20 >> ((42) & 63))
                var v90 uint32 = 0
                v91 := GenS13{F0: 245}
                _ = v90
                _ = v91
            }
            switch (v3 * (v74 % ((32768) | 1))) {
            case -954:
                fmt.Println(v3, v74)
            default:
                fmt.Println(v3, fi86)
                v92 := GenS14{F0: -1, F1: v20}
                _ = v92
            }
            _ = fi89
        }
        _ = fi86
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v74
    _ = v75
}
