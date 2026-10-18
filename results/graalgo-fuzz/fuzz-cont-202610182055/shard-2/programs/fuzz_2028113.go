package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 uint16; F2 uint16 }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 bool; F1 uint16; F2 string }

func main() {
    var v1 int64 = 0
    var v2 int = -634
    v3 := make(map[int]uint32)
    v4 := []int{}
    var v5 string = ("hello" + "foo bar")
    var v6 string = (v5 + v5)
    v2 += v2
    for fi7 := 0; fi7 < 18; fi7++ {
        v2 = (v2 - (v2 - 366))
        var v8 string = "foo bar"
        _ = fi7
        _ = v8
    }
    v5 = ""
    var v9 uint16 = 2
    if true {
        if (4294967295 >= 814) {
            if (!(true || true)) {
                fmt.Println(v5)
                v5 = ((v5 + v6) + (("hello" + "\n\t") + v6))
                var v10 uint = 977
                _ = v10
            }
            if (!(!(!false))) {
                if (-105 >= -921) {
                    var v11 uint8 = 2
                    switch (v6 + v5) {
                    case "\n\t":
                        v12 := GenS1{F0: true, F1: int8(v11)}
                        _ = v12
                    case "foo bar":
                        if (!false) {
                            v13 := GenS2{F0: (^(v9)), F1: (v5 + "\n\t")}
                            _ = v13
                        }
                    case "hello":
                        var v14 uint32 = 2147483648
                        _ = v14
                    default:
                        v4 = append(v4, v2)
                    }
                    var v15 int = 4294967295
                    _ = v11
                    _ = v15
                } else {
                    switch (v2 % (((v2 - v2)) | 1)) {
                    case -195:
                        v6 = "a"
                    case 2147483647:
                        var v16 float64 = (1.5 - 1e10)
                        _ = v16
                    case 32768:
                        v17 := GenS2{F0: v9, F1: "a"}
                        _ = v17
                    }
                    if true {
                        if true {
                            for fi18 := 0; fi18 < 2; fi18++ {
                                v5 = ("foo bar" + "foo bar")
                                _ = fi18
                            }
                            v19 := GenS1{F0: (true || (1.5 == (3.14159 * (float64((0.0 * 1e10)) - 0.0)))), F1: 100}
                            var v20 int32 = 1
                            _ = v19
                            _ = v20
                        }
                        if (42 >= 127) {
                            _ = len(v4)
                            fmt.Println(v2, v6)
                            if false {
                                if (!((-(v1)) <= v1)) {
                                    for fi21 := 0; fi21 < 10; fi21++ {
                                        if len(v4) > 0 {
                                            v4[0] = (fi21 | (fi21 & (fi21 + (v2 & int(v1)))))
                                        }
                                        _ = fi21
                                    }
                                    v4 = append(v4, (^(v2)))
                                    switch ("" + "foo bar") {
                                    case "":
                                        if false {
                                            var v22 float64 = 1.5
                                            _ = v22
                                        }
                                    default:
                                        if true {
                                            var v23 float64 = 1e10
                                            v24 := GenS3{F0: 127}
                                            var v25 uint64 = 18446744073709551615
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        } else {
                                            if (-128 < 47) {
                                                var v26 string = v5
                                                var v27 int = (-(v2))
                                                var v28 uint8 = 125
                                                _ = v26
                                                _ = v27
                                                _ = v28
                                            }
                                            var v29 uint16 = (v9 | 311)
                                            if true {
                                                if true {
                                                    if false {
                                                        v29 = 65535
                                                        var v30 int32 = -2147483648
                                                        _ = v30
                                                    } else {
                                                        fmt.Println(v6)
                                                        fmt.Println(v6, v1, v1)
                                                        v31 := GenS4{F0: 23, F1: 127, F2: -1}
                                                        _ = v31
                                                    }
                                                    if (false || (false || true)) {
                                                        var v32 int16 = -410
                                                        _ = v32
                                                    }
                                                }
                                            }
                                            _ = v29
                                        }
                                    }
                                } else {
                                    v33 := GenS1{F0: false, F1: 0}
                                    var v34 int = 2147483647
                                    var v35 uint8 = 255
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                if ((false && (1000 < 32767)) && false) {
                                    _ = len(v3)
                                    delete(v3, (v2 / ((-2147483648) | 1)))
                                } else {
                                    for fi36 := 0; fi36 < 1; fi36++ {
                                        v37 := GenS2{F0: v9, F1: v5}
                                        _ = fi36
                                        _ = v37
                                    }
                                    v1 += v1
                                    switch (((((v5 + v6) + ((((("a" + ("a" + v6)) + v6) + "\n\t") + ("\n\t" + ("a" + ""))) + (v6 + ("a" + ((("\n\t" + (v5 + v6)) + v6) + v6))))) + ((((v6 + v5) + v5) + v6) + "foo bar")) + v5) + "foo bar") {
                                    case "hello":
                                        fmt.Println(v1)
                                    default:
                                        v38 := GenS1{F0: (((v5 + v5) > "hello") && (39 == uint32(v9))), F1: -116}
                                        _ = v38
                                    }
                                }
                            }
                        } else {
                            var v39 uint64 = 572
                            v4 = append(v4, (v2 + v2))
                            _ = v39
                        }
                    } else {
                        v5 = ((("" + ("a" + (("\n\t" + ((v6 + v6) + (v5 + v6))) + "hello"))) + v5) + "hello")
                        switch 127 {
                        case 42:
                            if false {
                                fmt.Println(v9, v1, v6)
                                fmt.Println(v6, v5)
                            }
                        case 2147483647:
                            v2 -= (v2 / (((v2 - v2)) | 1))
                        case 128:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            for fi40 := 0; fi40 < 17; fi40++ {
                                v2 = v2
                                _ = fi40
                            }
                        }
                    }
                    switch 2147483647 {
                    case -8:
                        var v41 uint = 18446744073709551615
                        _ = v41
                    case -890:
                        var v42 uint16 = (v9 * (v9 % (((v9 % (((v9 | v9)) | 1))) | 1)))
                        _ = v42
                    case 1:
                        if true {
                            v43 := GenS4{F0: 18446744073709551615, F1: 255, F2: 1}
                            switch (v2 * v2) {
                            case 32768:
                                switch -2147483648 {
                                case -319:
                                    v44 := GenS2{F0: (v9 * 0), F1: "hello"}
                                    _ = v44
                                case 13:
                                    for fi45 := 0; fi45 < 6; fi45++ {
                                        v6 = "a"
                                        v1 = 1
                                        var v46 int8 = -2
                                        _ = fi45
                                        _ = v46
                                    }
                                case 32767:
                                    _ = len(v4)
                                }
                            }
                            v1 -= int64(v9)
                            _ = v43
                        } else {
                            var v47 string = v5
                            _ = v47
                        }
                    }
                }
                v5 = "\n\t"
                if (-32768 != 1) {
                    if false {
                        if true {
                            var v48 uint16 = (v9 % (((^(v9))) | 1))
                            v1 -= v1
                            v48 += v9
                            _ = v48
                        } else {
                            var v49 bool = (!false)
                            var v50 string = (((((v6 + v5) + ("a" + v5)) + ((v5 + (v5 + (v5 + v5))) + "a")) + v5) + v5)
                            for fi51 := 0; fi51 < 11; fi51++ {
                                v50 = "a"
                                _ = fi51
                            }
                            _ = v49
                            _ = v50
                        }
                        switch ("foo bar" + "a") {
                        case "":
                            var v52 int = (v2 & -903)
                            _ = v52
                        case "hello":
                            if false {
                                var v53 uint64 = 18446744073709551615
                                _ = v53
                            }
                        case "a":
                            if len(v4) > 0 {
                                v4[0] = 127
                            }
                        }
                        fmt.Println(v6)
                    }
                    var v54 uint64 = 272
                    v54 -= v54
                    _ = v54
                }
            } else {
                for fi55 := 0; fi55 < 5; fi55++ {
                    if false {
                        for fi56 := 0; fi56 < 18; fi56++ {
                            v3[(fi55 | fi56)] = uint32(v9)
                            _ = fi56
                        }
                        v4 = append(v4, (^(v2)))
                        v9 -= v9
                    }
                    v57 := GenS5{F0: v9}
                    v2 -= 2147483647
                    _ = fi55
                    _ = v57
                }
                delete(v3, 1)
            }
            fmt.Println(v5, v9)
        } else {
            switch (v2 | (v2 & v2)) {
            case 53:
                v9 -= v9
            case 32768:
                for fi58 := 0; fi58 < 19; fi58++ {
                    for fi59 := 0; fi59 < 3; fi59++ {
                        v4 = append(v4, (^(fi59)))
                        if true {
                            v4 = append(v4, fi58)
                        } else {
                            var v60 uint8 = 255
                            if (v1 >= 386) {
                                v61 := GenS4{F0: 752, F1: v60, F2: 157}
                                var v62 float32 = -2.25
                                if (0 >= 1) {
                                    fmt.Println(fi58, v62)
                                }
                                _ = v61
                                _ = v62
                            }
                            _ = v60
                        }
                        _ = fi59
                    }
                    fmt.Println(v2, v6, v1)
                    v63 := GenS2{F0: (v9 << ((42) & 63)), F1: ((v6 + (("" + "\n\t") + "\n\t")) + ((((v6 + ("foo bar" + v5)) + "") + (v6 + v6)) + v5))}
                    _ = fi58
                    _ = v63
                }
            default:
                if (true || (!(!(!true)))) {
                    v64 := GenS6{F0: -128}
                    switch 4294967295 {
                    case -128:
                        if true {
                            var v65 bool = (((4294967295 != 2147483648) && false) || true)
                            var v66 uint16 = (v9 & v9)
                            _ = v65
                            _ = v66
                        }
                    case -2147483648:
                        v67 := GenS7{F0: 2147483647, F1: (v9 * (v9 >> ((506) & 63))), F2: 755}
                        _ = v67
                    default:
                        v2 -= -291
                    }
                    v9 -= 55
                    _ = v64
                }
            }
        }
    } else {
        if true {
            var v68 uint64 = 626
            _ = v68
        }
    }
    var v69 bool = true
    var v70 int16 = -1
    switch v2 {
    case -940:
        var v71 float64 = float64(float32(0.1))
        _ = v71
    case 229:
        for fi72 := 0; fi72 < 7; fi72++ {
            fmt.Println(fi72, v69, v5)
            _ = fi72
        }
    default:
        v2 += (v2 >> ((2) & 63))
    }
    v69 = true
    if true {
        _ = len(v3)
        for fi73 := 0; fi73 < 1; fi73++ {
            v1 = 0
            v9 = v9
            _ = fi73
        }
    } else {
        switch 32768 {
        case -1:
            switch (-(v2)) {
            case 1:
                _ = len(v4)
            case 667:
                fmt.Println(v6, v69, v2)
            case 128:
                v4 = append(v4, -1)
            }
        case 128:
            if (v69 && true) {
                switch "a" {
                case "foo bar":
                    var v74 float32 = -2.25
                    _ = v74
                }
                v69 = true
                v6 = v6
            } else {
                switch (-(v2)) {
                case -2147483648:
                    for fi75 := 0; fi75 < 4; fi75++ {
                        switch (fi75 & fi75) {
                        case -9223372036854775808:
                            v1 = v1
                        default:
                            v6 = (v5 + "\n\t")
                        }
                        if (v2 != v2) {
                            if v69 {
                                for fi76 := 0; fi76 < 17; fi76++ {
                                    v1 = v1
                                    if v69 {
                                        v77 := GenS5{F0: 1}
                                        var v78 int16 = v70
                                        _ = v77
                                        _ = v78
                                    } else {
                                        var v79 uint8 = 127
                                        v80 := GenS3{F0: 255}
                                        _ = v79
                                        _ = v80
                                    }
                                    v70 = (v70 << ((1) & 63))
                                    v69 = (128 < fi75)
                                    var v81 int32 = 2147483647
                                    _ = fi76
                                    _ = v81
                                }
                            }
                        }
                        _ = fi75
                    }
                case -715:
                    v82 := GenS8{F0: 0, F1: 941}
                    _ = v82
                case -843:
                    var v83 float32 = (0.0 - -2.25)
                    _ = v83
                default:
                    v84 := v3[v2]
                    _ = v84
                }
                v1 = (^(v1))
                v6 = "a"
            }
        }
        var v85 int16 = v70
        _ = v85
    }
    if true {
        var v86 int16 = v70
        _ = v86
    }
    if (!false) {
        v87 := GenS9{F0: (v2 != (v2 ^ 128)), F1: (^(v9)), F2: ""}
        fmt.Println(v1, v70, v70)
        _ = v87
    } else {
        v69 = (v69 && false)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v9
}
