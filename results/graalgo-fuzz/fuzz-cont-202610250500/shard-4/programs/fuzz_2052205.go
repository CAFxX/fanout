package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 string; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint; F1 int8; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 uint32; F1 uint8; F2 uint }
type GenS8 struct { F0 uint8 }

func main() {
    var v1 uint = 42
    v2 := make(map[uint32]string)
    var v3 uint8 = 201
    fmt.Println(v3)
    var v4 int8 = -106
    v5 := []int{741, 791}
    switch "" {
    case "foo bar":
        switch 127 {
        case -1:
            var v6 int = -9223372036854775808
            _ = v6
        case -9223372036854775808:
            v3 -= v3
        case -531:
            v7 := GenS1{F0: (v1 + v1)}
            _ = v7
        }
    default:
        v8 := GenS1{F0: v1}
        _ = v8
    }
    var v9 float64 = -2.25
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            switch fi10 {
            case 4294967295:
                if false {
                    v12 := GenS2{F0: 0, F1: ("hello" + "hello"), F2: 2147483647}
                    v1 = v1
                    if true {
                        v13 := GenS3{F0: 42, F1: v4}
                        var v14 uint64 = 815
                        v3 -= 128
                        var v15 int64 = 0
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        var v16 uint = 18446744073709551615
                        _ = v16
                    }
                    fmt.Println(v9, fi11, v1)
                    _ = v12
                } else {
                    switch int(v4) {
                    case 207:
                        var v17 uint64 = 814
                        _ = v17
                    case 1:
                        var v18 int16 = 1000
                        _ = v18
                    case 32767:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        v19 := GenS3{F0: 65535, F1: v4}
                        _ = v19
                    }
                    var v20 int32 = 627
                    _ = v20
                }
            case 127:
                fmt.Println(fi11, fi11)
            case -9223372036854775808:
                var v21 uint64 = 429
                _ = v21
            }
            v1 += v1
            _ = fi11
        }
        _ = fi10
    }
    switch 32767 {
    case -2147483648:
        if true {
            for fi22 := 0; fi22 < 13; fi22++ {
                if (128 > (-(fi22))) {
                    if true {
                        var v23 int16 = 1000
                        _ = v23
                    }
                    if false {
                        var v24 bool = (v3 >= (v3 - v3))
                        _ = v24
                    }
                    switch (fi22 + -820) {
                    case 42:
                        switch "" {
                        case "foo bar":
                            v1 = v1
                        }
                    default:
                        v4 += (v4 + (v4 >> ((v1) & 63)))
                    }
                }
                _ = fi22
            }
        } else {
            var v25 uint = v1
            for fi26 := 0; fi26 < 15; fi26++ {
                switch ("\n\t" + (("a" + "\n\t") + (("hello" + "\n\t") + "a"))) {
                case "a":
                    for fi27 := 0; fi27 < 3; fi27++ {
                        switch (^(fi27)) {
                        case 42:
                            v28 := GenS2{F0: 659, F1: "\n\t", F2: 2147483647}
                            _ = v28
                        case -979:
                            if false {
                                if (true || ((fi26 * (-(fi27))) < fi26)) {
                                    var v29 float32 = 3.14159
                                    if (-1 > -32768) {
                                        v2[293] = ("hello" + "a")
                                    } else {
                                        v1 += 1
                                    }
                                    _ = v29
                                } else {
                                    var v30 int8 = (v4 & 89)
                                    _ = v30
                                }
                                v31 := GenS1{F0: (v1 & v25)}
                                v25 -= v1
                                _ = v31
                            }
                            v32 := v2[1]
                            if true {
                                var v33 uint16 = 2
                                v34 := GenS4{F0: 2, F1: v4, F2: -1}
                                if true {
                                    if true {
                                        switch ("" + "a") {
                                        case "a":
                                            if (v25 < 2) {
                                                var v35 int8 = (v4 / ((v4) | 1))
                                                _ = v35
                                            } else {
                                                var v36 int32 = 2147483647
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v33 += uint16(v3)
                                                _ = v36
                                            }
                                        case "\n\t":
                                            v9 = v9
                                        default:
                                            v33 = (^(v33))
                                        }
                                    } else {
                                        if (838 < 4294967295) {
                                            v33 -= (^(v33))
                                            delete(v2, 404)
                                        } else {
                                            v37 := GenS2{F0: -1, F1: v32, F2: 25}
                                            switch fi26 {
                                            case 128:
                                                fmt.Println(v32)
                                            case -9223372036854775808:
                                                var v38 float64 = ((v9 - (float64((3.14159 * (3.14159 - 1.5))) + v9)) - v9)
                                                _ = v38
                                            default:
                                                var v39 uint = v1
                                                v40 := v2[960]
                                                _ = v39
                                                _ = v40
                                            }
                                            v33 = 0
                                            _ = v37
                                        }
                                        var v41 int = fi26
                                        v32 = (v32 + v32)
                                        _ = v41
                                    }
                                } else {
                                    var v42 float64 = v9
                                    v43 := GenS2{F0: 42, F1: (v32 + v32), F2: 1}
                                    switch v32 {
                                    case "\n\t":
                                        fmt.Println(fi26, v42)
                                    }
                                    _ = v42
                                    _ = v43
                                }
                                _ = v33
                                _ = v34
                            } else {
                                v44 := GenS4{F0: (v1 >> ((255) & 63)), F1: (v4 | v4), F2: v4}
                                if false {
                                    var v45 int8 = (v4 * v4)
                                    v3 += v3
                                    switch (v32 + "a") {
                                    case "a":
                                        var v46 string = "a"
                                        _ = v46
                                    }
                                    _ = v45
                                } else {
                                    v1 -= v25
                                    v25 = 2
                                    v47 := GenS2{F0: -32768, F1: (((v32 + "\n\t") + "a") + "hello"), F2: 42}
                                    var v48 int16 = 32767
                                    _ = v47
                                    _ = v48
                                }
                                v9 -= float64(3.14159)
                                _ = v44
                            }
                            _ = v32
                        default:
                            var v49 uint32 = 873
                            _ = v49
                        }
                        var v50 int16 = -1
                        fmt.Println(v1, fi27, v4)
                        fmt.Println(v4)
                        v51 := GenS1{F0: (v1 | (^(v25)))}
                        _ = fi27
                        _ = v50
                        _ = v51
                    }
                case "":
                    switch "hello" {
                    case "hello":
                        var v52 int32 = 2147483647
                        _ = v52
                    default:
                        switch fi26 {
                        case 326:
                            delete(v2, 4294967295)
                        case -9223372036854775808:
                            fmt.Println(v4, v4, v9)
                        default:
                            v4 += v4
                        }
                    }
                }
                _ = fi26
            }
            v3 += (^(v3))
            _ = v25
        }
    case 4294967295:
        var v53 uint = 255
        _ = v53
    case 412:
        v54 := GenS2{F0: -32768, F1: ("hello" + ("a" + (((("a" + ("" + "\n\t")) + "\n\t") + "hello") + "hello"))), F2: 0}
        _ = v54
    }
    var v55 uint32 = 4294967295
    v4 = (v4 - (-(v4)))
    for fi56 := 0; fi56 < 18; fi56++ {
        var v57 bool = false
        switch (fi56 + 879) {
        case -233:
            if (fi56 > 2147483647) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v58 uint16 = 698
                v59 := GenS5{F0: (v58 & v58), F1: v55, F2: (v4 ^ 1)}
                _ = v58
                _ = v59
            } else {
                var v60 float64 = (v9 * ((v9 / 1e10) / ((1e10 * (v9 * v9)) - float64((3.14159 / 1e10)))))
                if len(v5) > 0 {
                    v5[0] = fi56
                }
                _ = v60
            }
        case -174:
            var v61 int8 = (v4 | (v4 - v4))
            _ = v61
        }
        if (1 == 42) {
            var v62 uint32 = (^(v55))
            v63 := GenS1{F0: 18446744073709551615}
            v57 = (((!v57) || v57) && (v4 == 0))
            _ = v62
            _ = v63
        }
        _ = fi56
        _ = v57
    }
    if (!true) {
        if true {
            var v64 uint = (v1 - 401)
            for fi65 := 0; fi65 < 9; fi65++ {
                v66 := GenS5{F0: 255, F1: 2147483648, F2: (v4 * v4)}
                _ = fi65
                _ = v66
            }
            var v67 float32 = 0.0
            _ = v64
            _ = v67
        }
        var v68 float32 = float32(v9)
        _ = v68
    }
    v55 = (v55 | v55)
    if (v1 != v1) {
        for fi69 := 0; fi69 < 12; fi69++ {
            v1 += 672
            _ = fi69
        }
    }
    if (false || (true || ((v3 & 42) >= v3))) {
        v1 = v1
        if (1 < 2) {
            var v70 int = 1
            _ = v70
        } else {
            if (!true) {
                if (false || (0 <= 0)) {
                    if ((^(v3)) <= (v3 << ((v1) & 63))) {
                        var v71 uint64 = 986
                        var v72 int32 = 42
                        _ = v71
                        _ = v72
                    }
                    if (0 < 18446744073709551615) {
                        var v73 int32 = 2147483647
                        _ = v73
                    }
                    switch "foo bar" {
                    case "foo bar":
                        if (0.0 > 0.1) {
                            var v74 uint16 = 65535
                            _ = v74
                        } else {
                            v3 -= (v3 / ((127) | 1))
                            for fi75 := 0; fi75 < 5; fi75++ {
                                var v76 int = fi75
                                for fi77 := 0; fi77 < 11; fi77++ {
                                    var v78 string = "a"
                                    var v79 bool = (false && false)
                                    switch int(v3) {
                                    case -408:
                                        fmt.Println(fi77)
                                    case 1:
                                        var v80 uint16 = 79
                                        _ = v80
                                    case -9223372036854775808:
                                        var v81 float32 = 1.5
                                        _ = v81
                                    default:
                                        var v82 uint16 = 994
                                        _ = v82
                                    }
                                    _ = fi77
                                    _ = v78
                                    _ = v79
                                }
                                _ = fi75
                                _ = v76
                            }
                        }
                    case "a":
                        v9 = (0.0 + (1e10 - (1.5 + v9)))
                    case "hello":
                        v55 = (v55 + 921)
                    default:
                        v55 -= v55
                    }
                } else {
                    v4 += v4
                    v83 := GenS5{F0: uint16(v1), F1: 77, F2: (v4 % (((v4 >> ((v1) & 63))) | 1))}
                    _ = v83
                }
                delete(v2, (v55 << ((v1) & 63)))
                v84 := GenS6{F0: "", F1: uint16(v1)}
                _ = v84
            } else {
                if (396 <= 42) {
                    _ = len(v2)
                    v85 := GenS7{F0: (^(v55)), F1: v3, F2: v1}
                    _ = v85
                } else {
                    for fi86 := 0; fi86 < 14; fi86++ {
                        v2[(v55 >> ((1) & 63))] = ("a" + "hello")
                        _ = fi86
                    }
                }
            }
            for fi87 := 0; fi87 < 11; fi87++ {
                for fi88 := 0; fi88 < 11; fi88++ {
                    if false {
                        fmt.Println(v55)
                        if (!true) {
                            var v89 int64 = 0
                            var v90 int32 = -1
                            v4 += v4
                            _ = v89
                            _ = v90
                        }
                        var v91 int64 = -9223372036854775808
                        _ = v91
                    } else {
                        v55 = 1
                        v92 := GenS7{F0: (v55 % (((v55 ^ v55)) | 1)), F1: v3, F2: 18446744073709551615}
                        _ = v92
                    }
                    _ = fi88
                }
                v1 -= 42
                for fi93 := 0; fi93 < 9; fi93++ {
                    v94 := GenS8{F0: (v3 * 42)}
                    _ = fi93
                    _ = v94
                }
                _ = fi87
            }
            if false {
                var v95 uint16 = 65535
                var v96 float32 = 3.14159
                _ = v95
                _ = v96
            } else {
                v4 = v4
                var v97 int = 4294967295
                v98 := GenS4{F0: 13, F1: 1, F2: (v4 - int8(v55))}
                _ = v97
                _ = v98
            }
        }
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v9
}
