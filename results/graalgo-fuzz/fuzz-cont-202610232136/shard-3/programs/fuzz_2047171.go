package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int16; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint16; F1 int; F2 string }
type GenS4 struct { F0 int32; F1 int32; F2 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32; F1 uint8; F2 int }
type GenS9 struct { F0 uint; F1 int; F2 int8 }
type GenS10 struct { F0 int16; F1 string; F2 string }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint8; F1 uint64; F2 int16 }
func genFn0() string {
    v6 := []int{-42, -9223372036854775808, -2147483648, 343}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    var v9 int16 = 42
    for fi10 := 0; fi10 < 3; fi10++ {
        if (v9 <= v9) {
            v9 = (v9 % ((v9) | 1))
            v9 = 32767
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int64 = -9223372036854775808
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    _ = v6
    _ = v9
    return "hello"
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    v4 := []int{-2147483648, 32768, 2147483647}
    v5 := make(map[uint32]string)
    var v13 uint16 = 350
    if true {
        switch -384 {
        case 832:
            var v14 string = "\n\t"
            _ = v14
        case 972:
            if true {
                switch "hello" {
                case "":
                    if (true || true) {
                        for fi15 := 0; fi15 < 10; fi15++ {
                            var v16 bool = true
                            _ = fi15
                            _ = v16
                        }
                    }
                case "\n\t":
                    switch "" {
                    case "\n\t":
                        var v17 float32 = ((0.0 + float32(((1e10 + 0.1) / 1.5))) - 3.14159)
                        _ = v17
                    case "a":
                        v18 := GenS1{F0: true, F1: 406}
                        _ = v18
                    case "foo bar":
                        var v19 uint = 0
                        _ = v19
                    }
                case "foo bar":
                    for fi20 := 0; fi20 < 5; fi20++ {
                        if false {
                            v13 -= v13
                        } else {
                            v21 := GenS2{F0: -32768, F1: 2147483647, F2: 32767}
                            var v22 string = (("foo bar" + "") + "foo bar")
                            _ = v21
                            _ = v22
                        }
                        _ = fi20
                    }
                }
                v1 -= 127
                switch 4294967295 {
                case 1:
                    var v23 float32 = 3.14159
                    _ = v23
                case 9223372036854775807:
                    if len(v4) > 0 {
                        v4[0] = 963
                    }
                case 32768:
                    if true {
                        v2[2147483647] = 2
                        v24 := GenS3{F0: (v13 & v13), F1: -2147483648, F2: "\n\t"}
                        _ = v24
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                default:
                    v25 := GenS2{F0: 176, F1: 2147483648, F2: 0}
                    _ = v25
                }
            } else {
                v26 := GenS2{F0: 1000, F1: 2147483647, F2: 42}
                v13 -= v13
                if (0 != 42) {
                    switch ("" + "") {
                    case "\n\t":
                        fmt.Println(v1)
                    case "a":
                        if (!(146 >= -2147483648)) {
                            v1 += v1
                            v27 := GenS4{F0: int32(v1), F1: 0, F2: -128}
                            if true {
                                v13 += v13
                                for fi28 := 0; fi28 < 13; fi28++ {
                                    var v29 float32 = float32(((0.1 + 1e10) - 0.0))
                                    _ = fi28
                                    _ = v29
                                }
                                v30 := GenS1{F0: false, F1: 42}
                                _ = v30
                            } else {
                                v4 = append(v4, 2147483647)
                                var v31 string = ("\n\t" + "a")
                                for fi32 := 0; fi32 < 5; fi32++ {
                                    v13 = (v13 << ((1) & 63))
                                    var v33 float32 = 1e10
                                    if false {
                                        if len(v4) > 0 {
                                            v4[0] = fi32
                                        }
                                        fmt.Println(v31, fi32)
                                    } else {
                                        fmt.Println(fi32)
                                    }
                                    _ = fi32
                                    _ = v33
                                }
                                _ = v31
                            }
                            _ = v27
                        } else {
                            for fi34 := 0; fi34 < 1; fi34++ {
                                v1 = v1
                                var v35 int16 = 32767
                                _ = fi34
                                _ = v35
                            }
                        }
                    default:
                        v1 += (v1 + v1)
                    }
                    var v36 string = ("hello" + "")
                    _ = v36
                } else {
                    v37 := GenS2{F0: 1000, F1: 2147483648, F2: -433}
                    _ = v37
                }
                _ = v26
            }
        case -56:
            switch "hello" {
            case "":
                if ((v13 <= v13) || true) {
                    var v38 int64 = 34
                    var v39 int16 = 42
                    _ = v38
                    _ = v39
                } else {
                    if ((true || false) && (1 > 39)) {
                        if false {
                            for fi40 := 0; fi40 < 20; fi40++ {
                                v13 = (v13 | (v13 + v13))
                                if true {
                                    var v41 uint16 = 719
                                    if false {
                                        delete(v5, 998)
                                    } else {
                                        var v42 bool = (!((v1 >> ((4294967295) & 63)) > 151))
                                        _ = v42
                                    }
                                    _ = v41
                                } else {
                                    for fi43 := 0; fi43 < 4; fi43++ {
                                        var v44 int64 = 42
                                        _ = fi43
                                        _ = v44
                                    }
                                    for fi45 := 0; fi45 < 10; fi45++ {
                                        delete(v3, (-(fi45)))
                                        var v46 float64 = 0.0
                                        if false {
                                            switch "foo bar" {
                                            case "":
                                                var v47 uint16 = (v13 ^ 2)
                                                v3[1] = ("" + ("foo bar" + ""))
                                                v1 = v1
                                                _ = v47
                                            case "a":
                                                v46 -= v46
                                            }
                                            var v48 uint16 = v13
                                            var v49 uint = 42
                                            fmt.Println(v48, v13)
                                            _ = v48
                                            _ = v49
                                        } else {
                                            if true {
                                                if (0 <= 18446744073709551615) {
                                                    v1 += v1
                                                } else {
                                                    if (false && ((^(v13)) == v13)) {
                                                        if (true && (float64(1.5) > (0.0 * 0.1))) {
                                                            switch fi40 {
                                                            case 0:
                                                                switch fi40 {
                                                                case 127:
                                                                    _ = len(v4)
                                                                }
                                                            default:
                                                                v50 := GenS2{F0: 124, F1: 89, F2: -32768}
                                                                _ = v50
                                                            }
                                                            if (742 <= 42) {
                                                                v46 = ((-2.25 - 3.14159) - float64((3.14159 * 3.14159)))
                                                                fmt.Println(v46, v13, v46)
                                                            }
                                                            v51 := GenS1{F0: ((v1 % ((v1) | 1)) < 127), F1: -813}
                                                            _ = v51
                                                        } else {
                                                            fmt.Println(v46, fi40)
                                                            if false {
                                                                fmt.Println(v13, fi45)
                                                            }
                                                            v52 := GenS3{F0: v13, F1: (-(fi40)), F2: ("\n\t" + "hello")}
                                                            _ = v52
                                                        }
                                                    }
                                                    var v53 uint32 = 42
                                                    v54 := GenS5{F0: v1}
                                                    _ = v53
                                                    _ = v54
                                                }
                                                var v55 int = fi40
                                                var v56 uint8 = 72
                                                _ = v55
                                                _ = v56
                                            } else {
                                                var v57 uint = 65535
                                                var v58 int8 = 62
                                                _ = v57
                                                _ = v58
                                            }
                                            if (v46 != (v46 + 1.5)) {
                                                var v59 string = ("" + ((("foo bar" + "hello") + "") + "a"))
                                                v1 = v1
                                                v13 -= 32767
                                                _ = v59
                                            }
                                            var v60 bool = (v13 <= 0)
                                            _ = v60
                                        }
                                        v13 -= v13
                                        if false {
                                            v46 -= v46
                                            _ = len(v2)
                                            switch ("hello" + (("a" + "a") + ("foo bar" + ("\n\t" + ("hello" + ("\n\t" + (((("hello" + "hello") + "") + (((("foo bar" + "hello") + "") + "hello") + "hello")) + "hello"))))))) {
                                            case "":
                                                if true {
                                                    v3[(fi45 | (^(fi40)))] = "\n\t"
                                                    v61 := GenS6{F0: true, F1: (v13 ^ (v13 | 65535)), F2: -9223372036854775808}
                                                    var v62 uint = 18446744073709551615
                                                    _ = v61
                                                    _ = v62
                                                }
                                            case "a":
                                                v63 := GenS1{F0: true, F1: 9223372036854775807}
                                                _ = v63
                                            case "\n\t":
                                                v64 := GenS7{F0: 18446744073709551615}
                                                _ = v64
                                            default:
                                                var v65 uint8 = v1
                                                _ = v65
                                            }
                                        }
                                        _ = fi45
                                        _ = v46
                                    }
                                }
                                _ = fi40
                            }
                            var v66 uint32 = 2
                            v67 := GenS3{F0: 0, F1: 127, F2: "a"}
                            _ = v66
                            _ = v67
                        } else {
                            v68 := v5[4294967295]
                            _ = v68
                        }
                    }
                    switch ("hello" + "foo bar") {
                    case "":
                        switch ((("a" + "foo bar") + ((("" + "foo bar") + "foo bar") + (((("hello" + "hello") + "hello") + ("foo bar" + "\n\t")) + "foo bar"))) + "foo bar") {
                        case "foo bar":
                            var v69 uint16 = (v13 << ((65535) & 63))
                            _ = v69
                        case "\n\t":
                            switch -2147483648 {
                            case -9223372036854775808:
                                for fi70 := 0; fi70 < 1; fi70++ {
                                    for fi71 := 0; fi71 < 10; fi71++ {
                                        v1 -= (^(v1))
                                        var v72 int = (fi71 / ((fi71) | 1))
                                        v1 = 127
                                        _ = fi71
                                        _ = v72
                                    }
                                    delete(v3, (fi70 >> ((569) & 63)))
                                    _ = fi70
                                }
                            case -1:
                                v73 := GenS8{F0: 0, F1: v1, F2: 128}
                                _ = v73
                            case -369:
                                v74 := GenS9{F0: 540, F1: -280, F2: 0}
                                _ = v74
                            default:
                                if true {
                                    switch 42 {
                                    case 32767:
                                        var v75 int8 = 1
                                        _ = v75
                                    case 534:
                                        v13 = v13
                                    default:
                                        v13 += 32767
                                    }
                                    for fi76 := 0; fi76 < 10; fi76++ {
                                        for fi77 := 0; fi77 < 2; fi77++ {
                                            fmt.Println(fi76, v1)
                                            v78 := GenS10{F0: 1, F1: ("foo bar" + "\n\t"), F2: "a"}
                                            var v79 int32 = 42
                                            _ = fi77
                                            _ = v78
                                            _ = v79
                                        }
                                        for fi80 := 0; fi80 < 11; fi80++ {
                                            _ = len(v5)
                                            _ = fi80
                                        }
                                        v81 := GenS11{F0: -2147483648}
                                        _ = fi76
                                        _ = v81
                                    }
                                    fmt.Println(v1, v1)
                                }
                            }
                        default:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    default:
                        v82 := GenS12{F0: (v1 ^ v1), F1: 397, F2: 96}
                        _ = v82
                    }
                }
            case "a":
                var v83 uint64 = 921
                _ = v83
            case "hello":
                for fi84 := 0; fi84 < 19; fi84++ {
                    var v85 uint32 = 0
                    _ = fi84
                    _ = v85
                }
            }
        default:
            v4 = append(v4, -9223372036854775808)
        }
        delete(v3, 9223372036854775807)
        delete(v3, 2147483647)
    } else {
        for fi86 := 0; fi86 < 4; fi86++ {
            var v87 int8 = int8(v13)
            for fi88 := 0; fi88 < 12; fi88++ {
                fmt.Println(v1, fi88)
                if (0 > 0) {
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        var v89 int16 = 1
                        _ = v89
                    }
                    v87 -= v87
                }
                var v90 int32 = -2147483648
                _ = fi88
                _ = v90
            }
            fmt.Println(v87)
            _ = fi86
            _ = v87
        }
        if false {
            for fi91 := 0; fi91 < 19; fi91++ {
                v1 -= (v1 * 20)
                var v92 string = ""
                _ = fi91
                _ = v92
            }
            if (false && (696 == 2)) {
                v13 += v13
                v1 -= v1
                var v93 uint = 65535
                _ = v93
            } else {
                v4 = append(v4, 9223372036854775807)
                v1 = v1
                var v94 string = "hello"
                _ = v94
            }
            var v95 int64 = -686
            _ = v95
        }
        var v96 uint32 = 0
        _ = v96
    }
    var v97 float32 = 0.0
    _ = len(v2)
    v5[4294967295] = (((("foo bar" + ("a" + "\n\t")) + ("a" + "hello")) + "\n\t") + "hello")
    v1 = 94
    v98 := GenS11{F0: -1}
    var v99 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v97:", v97)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v97
    _ = v98
    _ = v99
}
