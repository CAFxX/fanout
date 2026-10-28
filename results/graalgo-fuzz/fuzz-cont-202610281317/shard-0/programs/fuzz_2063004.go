package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 int32; F1 int32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint16; F1 uint; F2 bool }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int16; F1 uint64; F2 uint8 }
func genFn0(p0 uint16, p1 int32) uint {
    p1 = (-(p1))
    var v6 int = 614
    switch (v6 * v6) {
    case 127:
        fmt.Println(p0, v6)
    }
    if false {
        var v7 int32 = p1
        var v8 uint16 = p0
        var v9 int8 = 111
        _ = v7
        _ = v8
        _ = v9
    } else {
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 int8 = -1
    var v2 string = "hello"
    v3 := make(map[string]bool)
    v4 := []string{"a", ((((v2 + v2) + ((v2 + v2) + "")) + v2) + (v2 + "a")), "hello"}
    var v5 float64 = -2.25
    fmt.Println(v1)
    if false {
        if true {
            v5 = 0.0
            v10 := GenS1{F0: 42}
            fmt.Println(v5, v5)
            _ = v10
        } else {
            v11 := GenS2{F0: 65535}
            v12 := GenS3{F0: 18446744073709551615}
            delete(v3, v2)
            _ = v11
            _ = v12
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        var v13 float32 = -2.25
        if true {
            var v14 int = 127
            if false {
                var v15 uint = 1
                for fi16 := 0; fi16 < 17; fi16++ {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        if (!false) {
                            v18 := GenS3{F0: 42}
                            if ((("a" + v2) + "foo bar") <= "hello") {
                                v1 += (v1 >> (((v15 / (((v15 / (((v15 * (v15 & v15))) | 1))) | 1))) & 63))
                                var v19 string = (v2 + v2)
                                _ = v19
                            } else {
                                if (42 == -617) {
                                    switch "foo bar" {
                                    case "hello":
                                        var v20 uint64 = 19
                                        _ = v20
                                    case "a":
                                        fmt.Println(v1, v2, v15)
                                    case "\n\t":
                                        v1 += v1
                                    default:
                                        if (int32(v15) == 0) {
                                            fmt.Println(v2, v2)
                                            v15 = v15
                                            var v21 int = fi17
                                            switch v21 {
                                            case -70:
                                                if true {
                                                    fmt.Println(v13, v15, fi16)
                                                }
                                            default:
                                                v4 = append(v4, v2)
                                            }
                                            v22 := GenS1{F0: 42}
                                            var v23 int8 = (v1 | 0)
                                            _ = v21
                                            _ = v22
                                            _ = v23
                                        } else {
                                            if false {
                                                delete(v3, "hello")
                                            }
                                            v24 := GenS4{F0: ((true && false) || true), F1: 456}
                                            _ = v24
                                        }
                                    }
                                } else {
                                    if false {
                                        v13 = (v13 - (v13 * 1.5))
                                        v25 := GenS4{F0: true, F1: 9223372036854775807}
                                        v26 := v3[v2]
                                        _ = v25
                                        _ = v26
                                    } else {
                                        var v27 float64 = (0.0 + v5)
                                        _ = v27
                                    }
                                    var v28 int64 = -9223372036854775808
                                    _ = v28
                                }
                                switch -2147483648 {
                                case 42:
                                    var v29 bool = true
                                    _ = v29
                                default:
                                    v5 -= v5
                                }
                            }
                            v30 := GenS5{F0: 2, F1: -778}
                            _ = v18
                            _ = v30
                        } else {
                            v5 = 3.14159
                            v31 := GenS6{F0: false}
                            v32 := GenS2{F0: 748}
                            _ = v31
                            _ = v32
                        }
                        _ = fi17
                    }
                    _ = fi16
                }
                switch "a" {
                case "":
                    v33 := GenS2{F0: 1}
                    _ = v33
                }
                _ = v15
            }
            _ = v14
        } else {
            for fi34 := 0; fi34 < 18; fi34++ {
                switch v2 {
                case "\n\t":
                    v2 = (v2 + v2)
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "foo bar":
                    for fi35 := 0; fi35 < 2; fi35++ {
                        var v36 string = ""
                        _ = fi35
                        _ = v36
                    }
                }
                _ = fi34
            }
            v2 = ((v2 + v2) + v2)
        }
        var v37 float32 = 1e10
        _ = v13
        _ = v37
    }
    var v38 float32 = float32(v5)
    fmt.Println(v2, v5, v2)
    fmt.Println(v1, v38, v1)
    v39 := GenS7{F0: 371, F1: -332}
    v40 := GenS8{F0: 42}
    var v41 string = "\n\t"
    var v42 int64 = 9223372036854775807
    v41 = ("a" + v2)
    switch v41 {
    case "hello":
        switch (v41 + ("a" + v41)) {
        case "":
            switch v41 {
            case "a":
                for fi43 := 0; fi43 < 17; fi43++ {
                    switch fi43 {
                    case 128:
                        v5 -= 1.5
                    }
                    var v44 float64 = v5
                    _ = fi43
                    _ = v44
                }
            case "hello":
                v45 := v3["a"]
                _ = v45
            case "\n\t":
                if len(v4) > 0 {
                    v4[0] = (("" + "a") + v2)
                }
            }
        case "\n\t":
            for fi46 := 0; fi46 < 10; fi46++ {
                if false {
                    v1 += 58
                    v42 += v42
                    v2 = v2
                } else {
                    var v47 float64 = v5
                    if false {
                        v42 = (-(v42))
                        if (true || false) {
                            v47 += (((3.14159 + v47) / ((v5 + ((float64(v38) * 1e10) - v47)) + v47)) + v5)
                            delete(v3, v41)
                        }
                    } else {
                        var v48 int64 = -1
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        _ = v48
                    }
                    _ = v47
                }
                switch v2 {
                case "hello":
                    if (!(true && true)) {
                        switch "a" {
                        case "":
                            switch 4294967295 {
                            case 128:
                                fmt.Println(v41)
                            }
                        }
                    }
                case "foo bar":
                    for fi49 := 0; fi49 < 15; fi49++ {
                        var v50 float64 = -2.25
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v1 -= (v1 + v1)
                        _ = fi49
                        _ = v50
                    }
                case "a":
                    for fi51 := 0; fi51 < 20; fi51++ {
                        fmt.Println(fi46, fi46, v5)
                        _ = fi51
                    }
                default:
                    for fi52 := 0; fi52 < 17; fi52++ {
                        if true {
                            switch -9223372036854775808 {
                            case -414:
                                if false {
                                    v41 = "\n\t"
                                } else {
                                    v53 := GenS6{F0: false}
                                    _ = v53
                                }
                            case 2147483647:
                                switch (fi52 & -2147483648) {
                                case 42:
                                    _ = len(v4)
                                case 32768:
                                    var v54 uint = genFn0(672, -1)
                                    _ = v54
                                default:
                                    v55 := GenS9{F0: 1, F1: -32768}
                                    _ = v55
                                }
                            default:
                                v38 = v38
                            }
                            if false {
                                var v56 uint64 = 2
                                delete(v3, v2)
                                v57 := GenS10{F0: -2147483648, F1: -656}
                                if len(v4) > 0 {
                                    v4[0] = v41
                                }
                                _ = v56
                                _ = v57
                            }
                        } else {
                            var v58 int32 = 2147483647
                            v59 := GenS11{F0: 556}
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v58
                            _ = v59
                        }
                        _ = fi52
                    }
                }
                fmt.Println(v42, v1)
                _ = fi46
            }
        case "foo bar":
            var v60 float32 = v38
            _ = v60
        }
    case "\n\t":
        var v61 int = 4294967295
        _ = v61
    default:
        if (!((4294967295 <= 42) && (559 > 65535))) {
            v42 -= (v42 * (^(v42)))
            if false {
                var v62 uint8 = 180
                if false {
                    v63 := GenS6{F0: false}
                    _ = v63
                } else {
                    var v64 int8 = v1
                    _ = v64
                }
                if (false && (1 <= 1)) {
                    var v65 int8 = v1
                    var v66 bool = ((v42 % (((v42 ^ 1)) | 1)) <= v42)
                    _ = v65
                    _ = v66
                } else {
                    v38 = 3.14159
                }
                _ = v62
            } else {
                v42 = 936
                v38 += 0.1
            }
            if (!(42 >= 2)) {
                var v67 int = -543
                fmt.Println(v38)
                for fi68 := 0; fi68 < 1; fi68++ {
                    if (!(!(v41 != v41))) {
                        var v69 bool = false
                        _ = v69
                    }
                    v2 = v41
                    _ = fi68
                }
                _ = v67
            }
        } else {
            delete(v3, "a")
            var v70 uint32 = 244
            _ = v70
        }
    }
    for fi71 := 0; fi71 < 17; fi71++ {
        switch ("foo bar" + v41) {
        case "":
            var v72 float64 = v5
            _ = v72
        case "foo bar":
            if true {
                for fi73 := 0; fi73 < 4; fi73++ {
                    var v74 float32 = (float32(v5) - float32(((v5 + 0.1) + (0.1 + v5))))
                    _ = fi73
                    _ = v74
                }
            }
        default:
            for fi75 := 0; fi75 < 20; fi75++ {
                var v76 uint8 = 127
                v1 += (^(v1))
                v77 := GenS4{F0: false, F1: -1}
                _ = fi75
                _ = v76
                _ = v77
            }
        }
        for fi78 := 0; fi78 < 2; fi78++ {
            var v79 int32 = 42
            _ = fi78
            _ = v79
        }
        if false {
            v80 := GenS12{F0: 32767, F1: 2, F2: true}
            v3[v2] = false
            _ = v80
        }
        _ = fi71
    }
    if true {
        for fi81 := 0; fi81 < 6; fi81++ {
            var v82 int64 = v42
            for fi83 := 0; fi83 < 18; fi83++ {
                var v84 int16 = -1
                v3["hello"] = (false || ((v84 + 32767) != v84))
                v85 := GenS13{F0: 2}
                _ = fi83
                _ = v84
                _ = v85
            }
            _ = fi81
            _ = v82
        }
        var v86 int32 = 62
        if true {
            v2 = v2
            for fi87 := 0; fi87 < 4; fi87++ {
                v88 := GenS14{F0: 1, F1: 891, F2: 26}
                _ = fi87
                _ = v88
            }
            v5 -= -2.25
        } else {
            var v89 int32 = v86
            if false {
                var v90 uint32 = 525
                _ = len(v4)
                _ = v90
            } else {
                if (!(4294967295 > 42)) {
                    for fi91 := 0; fi91 < 5; fi91++ {
                        var v92 uint64 = 0
                        fmt.Println(v38, v89)
                        var v93 uint16 = 126
                        _ = fi91
                        _ = v92
                        _ = v93
                    }
                    v86 += 101
                    v94 := GenS6{F0: false}
                    _ = v94
                } else {
                    fmt.Println(v38, v38)
                }
            }
            _ = v89
        }
        _ = v86
    } else {
        fmt.Println(v38)
    }
    var v95 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v95
}
