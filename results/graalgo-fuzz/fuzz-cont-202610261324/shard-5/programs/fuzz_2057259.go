package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS7 struct { F0 int; F1 bool; F2 string }
type GenS8 struct { F0 uint; F1 uint }
type GenS9 struct { F0 int32; F1 uint64 }
type GenS10 struct { F0 uint8; F1 int16; F2 int }
type GenS11 struct { F0 uint64; F1 uint32; F2 uint16 }

func main() {
    var v1 int16 = 42
    v2 := make(map[int]bool)
    var v3 float32 = 0.0
    v4 := []string{}
    v5 := []uint32{}
    var v6 int16 = v1
    v7 := GenS1{F0: (!false), F1: ("\n\t" + ("a" + ("" + (("" + "hello") + ((("hello" + "") + "\n\t") + "hello"))))), F2: 2147483648}
    v2[-1] = false
    if (false || true) {
        var v8 bool = true
        v6 -= 0
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 uint32 = 2147483648
            v11 := GenS2{F0: int8(v10)}
            var v12 uint32 = (v10 - (v10 + (v10 ^ (v10 >> ((18446744073709551615) & 63)))))
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = v8
    }
    v5 = append(v5, 42)
    fmt.Println(v1, v6)
    switch -483 {
    case 128:
        var v13 bool = true
        _ = v13
    case 2147483647:
        var v14 float64 = (0.0 - (1e10 / (((3.14159 + 0.0) + 1e10) * 3.14159)))
        _ = v14
    }
    v1 = (v1 & 42)
    var v15 uint8 = 255
    var v16 float32 = -2.25
    var v17 string = "hello"
    v18 := GenS2{F0: 127}
    v19 := GenS1{F0: false, F1: v17, F2: 0}
    v3 = ((v16 + v16) - -2.25)
    for fi20 := 0; fi20 < 16; fi20++ {
        if false {
            if (!(false || false)) {
                v17 = v17
                for fi21 := 0; fi21 < 3; fi21++ {
                    v3 = v3
                    _ = fi21
                }
            } else {
                if (!(!(!true))) {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        switch "foo bar" {
                        case "a":
                            v23 := v2[128]
                            _ = v23
                        case "":
                            switch (fi20 - 202) {
                            case 194:
                                _ = len(v5)
                            }
                        case "\n\t":
                            _ = len(v2)
                        }
                        delete(v2, fi22)
                        _ = fi22
                    }
                    v24 := GenS3{F0: 0, F1: -1}
                    _ = v24
                }
                var v25 int64 = -1
                _ = v25
            }
            switch (v17 + "\n\t") {
            case "hello":
                for fi26 := 0; fi26 < 2; fi26++ {
                    v17 = (v17 + "")
                    fmt.Println(v6, v3, v15)
                    v2[fi20] = true
                    _ = fi26
                }
            case "a":
                v27 := GenS4{F0: 805, F1: 42, F2: 517}
                _ = v27
            case "":
                for fi28 := 0; fi28 < 12; fi28++ {
                    fmt.Println(v16)
                    var v29 uint8 = (v15 & 255)
                    v30 := v2[(fi28 >> ((65535) & 63))]
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            default:
                var v31 int = fi20
                _ = v31
            }
            _ = len(v2)
        } else {
            if false {
                if true {
                    for fi32 := 0; fi32 < 19; fi32++ {
                        if true {
                            var v33 int = (fi32 ^ (^(fi20)))
                            if (!true) {
                                if false {
                                    switch fi32 {
                                    case 42:
                                        if (2 > 18446744073709551615) {
                                            v2[(fi20 ^ (fi20 << ((18446744073709551615) & 63)))] = false
                                            v5 = append(v5, 994)
                                            switch v17 {
                                            case "hello":
                                                v34 := GenS1{F0: false, F1: ("a" + (v17 + v17)), F2: 150}
                                                _ = v34
                                            case "foo bar":
                                                _ = len(v2)
                                            }
                                        }
                                    }
                                    v2[680] = true
                                } else {
                                    v17 = v17
                                    v17 = v17
                                    var v35 int8 = 42
                                    _ = v35
                                }
                                if (!false) {
                                    v36 := GenS5{F0: 42}
                                    fmt.Println(fi20, fi20)
                                    if (!(true && (943 <= 4294967295))) {
                                        v2[-346] = true
                                        if len(v5) > 0 {
                                            v5[0] = 508
                                        }
                                    } else {
                                        fmt.Println(fi20)
                                        var v37 int8 = int8(v6)
                                        var v38 int8 = v37
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = v36
                                }
                                if ((v15 != (v15 & 26)) && true) {
                                    var v39 int8 = 100
                                    if ((v15 | 0) >= (v15 & v15)) {
                                        var v40 int32 = int32(fi20)
                                        var v41 bool = (4294967295 == 495)
                                        v42 := GenS2{F0: 105}
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v1 = v6
                                        v43 := GenS6{F0: (v15 * (v15 / ((1) | 1))), F1: 2, F2: v15}
                                        _ = v43
                                    }
                                    fmt.Println(fi32)
                                    v6 += 550
                                    var v44 float64 = 0.0
                                    _ = v39
                                    _ = v44
                                } else {
                                    v5 = append(v5, 42)
                                    var v45 uint32 = 2
                                    v45 = 2147483648
                                    v46 := GenS7{F0: (^(v33)), F1: (820 != 42), F2: "\n\t"}
                                    var v47 uint32 = 2147483648
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                }
                            } else {
                                var v48 int32 = 42
                                _ = v48
                            }
                            _ = v33
                        } else {
                            switch fi32 {
                            case -2147483648:
                                var v49 int16 = v6
                                var v50 int = (fi32 & fi20)
                                switch (v50 + v50) {
                                case -9223372036854775808:
                                    var v51 uint8 = v15
                                    _ = v51
                                case 42:
                                    var v52 int8 = -69
                                    _ = v52
                                }
                                if true {
                                    fmt.Println(v16, v15)
                                    var v53 uint32 = 0
                                    _ = v53
                                } else {
                                    v2[4294967295] = (v17 < v17)
                                    var v54 float32 = (v16 * v3)
                                    if (v16 <= v3) {
                                        if true {
                                            if false {
                                                switch "hello" {
                                                case "\n\t":
                                                    var v55 float64 = -2.25
                                                    _ = v55
                                                case "foo bar":
                                                    fmt.Println(fi32, v3)
                                                case "a":
                                                    var v56 int16 = v6
                                                    _ = v56
                                                }
                                                var v57 string = v17
                                                _ = v57
                                            } else {
                                                v58 := GenS8{F0: 614, F1: 65535}
                                                _ = v58
                                            }
                                        }
                                        v50 += fi32
                                        switch "a" {
                                        case "foo bar":
                                            if false {
                                                switch v17 {
                                                case "foo bar":
                                                    var v59 int64 = -9223372036854775808
                                                    _ = v59
                                                case "":
                                                    v6 += (^(v6))
                                                    v60 := GenS9{F0: 1, F1: 254}
                                                    _ = v60
                                                default:
                                                    switch fi20 {
                                                    case 317:
                                                        v61 := GenS7{F0: -9223372036854775808, F1: (88 < 109), F2: v17}
                                                        _ = v61
                                                    case 42:
                                                        switch v50 {
                                                        case 0:
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        }
                                                    case 4294967295:
                                                        fmt.Println(fi20)
                                                    }
                                                }
                                                v49 += v6
                                            }
                                        default:
                                            var v62 int16 = v49
                                            _ = v62
                                        }
                                    } else {
                                        if ((300 >= 42) && (229 < 220)) {
                                            v63 := GenS9{F0: -708, F1: 42}
                                            var v64 float32 = v54
                                            _ = v63
                                            _ = v64
                                        }
                                        v54 = -2.25
                                        v3 -= (((0.1 + (0.0 * ((-2.25 * 3.14159) / ((0.0 * 3.14159) / v54)))) * v16) / float32(1.5))
                                        var v65 int16 = v6
                                        fmt.Println(fi20, v49, v15)
                                        _ = v65
                                    }
                                    _ = v54
                                }
                                fmt.Println(v6)
                                _ = v49
                                _ = v50
                            case 632:
                                v2[32768] = true
                            default:
                                if false {
                                    var v66 uint32 = 2
                                    v67 := GenS6{F0: (^(v15)), F1: 978, F2: (v15 >> ((1) & 63))}
                                    _ = v66
                                    _ = v67
                                }
                            }
                            v16 -= v3
                            if true {
                                var v68 uint = 1
                                _ = v68
                            } else {
                                if true {
                                    var v69 uint8 = (v15 & v15)
                                    _ = v69
                                } else {
                                    v70 := GenS10{F0: (v15 | 127), F1: (v6 ^ (v6 * (v6 - v6))), F2: -22}
                                    var v71 float64 = 0.0
                                    _ = v70
                                    _ = v71
                                }
                            }
                        }
                        var v72 uint8 = v15
                        _ = fi32
                        _ = v72
                    }
                    v6 = 0
                }
                v73 := v2[-310]
                _ = v73
            } else {
                for fi74 := 0; fi74 < 12; fi74++ {
                    fmt.Println(v17, v17, v17)
                    switch (fi20 >> ((721) & 63)) {
                    case -1:
                        var v75 float32 = 0.1
                        _ = v75
                    case -506:
                        var v76 int32 = -1
                        _ = v76
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi74
                }
            }
            var v77 string = v17
            _ = v77
        }
        var v78 uint = 2
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi20
        _ = v78
    }
    v79 := GenS11{F0: 42, F1: 751, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v79
}
