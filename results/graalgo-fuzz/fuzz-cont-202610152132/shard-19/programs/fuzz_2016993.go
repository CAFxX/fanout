package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint32; F1 int16; F2 int8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint32; F1 string; F2 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 int32; F2 int32 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint, p1 string) uint64 {
    fmt.Println(p1, p1)
    for fi6 := 0; fi6 < 2; fi6++ {
        p1 = ""
        if false {
            switch (p1 + (p1 + "\n\t")) {
            case "\n\t":
                for fi7 := 0; fi7 < 5; fi7++ {
                    v8 := []int{}
                    p1 = (p1 + "foo bar")
                    _ = fi7
                    _ = v8
                }
            case "":
                fmt.Println(p1, fi6)
            case "foo bar":
                var v9 int64 = -1
                _ = v9
            default:
                if (false && true) {
                    p0 = p0
                    for fi10 := 0; fi10 < 12; fi10++ {
                        var v11 uint32 = 2147483647
                        v12 := []int{}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    switch (fi6 | (fi6 >> ((4294967295) & 63))) {
                    case 128:
                        fmt.Println(fi6, p1, p1)
                    }
                } else {
                    var v13 float64 = 0.1
                    var v14 int8 = 42
                    for fi15 := 0; fi15 < 1; fi15++ {
                        var v16 float32 = 0.1
                        _ = fi15
                        _ = v16
                    }
                    _ = v13
                    _ = v14
                }
            }
        } else {
            var v17 int = 0
            fmt.Println(v17, p0)
            p0 -= p0
            _ = v17
        }
        _ = fi6
    }
    var v18 bool = (42 < 1)
    _ = p0
    _ = p1
    _ = v18
    return 2
}
func genFn1() string {
    if (!false) {
        switch "" {
        case "foo bar":
            var v19 int32 = 42
            _ = v19
        case "\n\t":
            v20 := GenS1{F0: 42, F1: 2}
            _ = v20
        }
        v21 := []int{}
        var v22 float64 = (((1.5 - -2.25) / 0.1) + ((float64(3.14159) * float64((1.5 / 1e10))) - 1.5))
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: 152, F1: 1000, F2: -87}
    _ = v23
    return ("foo bar" + "foo bar")
}

func main() {
    v1 := []int{-1, 128, 689, 998}
    var v2 int16 = 1
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    var v5 bool = false
    var v24 uint32 = 2147483647
    switch -9223372036854775808 {
    case 128:
        delete(v3, "a")
    }
    if len(v1) > 0 {
        v1[0] = 763
    }
    fmt.Println(v24, v5, v5)
    if (!true) {
        switch "hello" {
        case "a":
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 float32 = 1.5
                switch "a" {
                case "a":
                    fmt.Println(fi25)
                default:
                    if ((!(!(-2147483648 != int32(v2)))) && (127 < 42)) {
                        var v27 int8 = -128
                        v28 := GenS1{F0: 471, F1: 1}
                        _ = v27
                        _ = v28
                    } else {
                        var v29 string = "a"
                        v24 = 0
                        v30 := v4[v24]
                        _ = v29
                        _ = v30
                    }
                }
                _ = fi25
                _ = v26
            }
        case "\n\t":
            v2 = -32768
        case "hello":
            var v31 uint8 = 42
            _ = v31
        default:
            if len(v1) > 0 {
                v1[0] = 906
            }
        }
    } else {
        fmt.Println(v2, v2, v24)
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
    }
    v1 = append(v1, -1)
    for fi32 := 0; fi32 < 8; fi32++ {
        v5 = ((("foo bar" + "") + "\n\t") >= "foo bar")
        v5 = true
        _ = fi32
    }
    var v33 uint16 = 223
    v2 = v2
    fmt.Println(v24, v2, v5)
    var v34 bool = v5
    v5 = (v34 || v5)
    var v35 int64 = 42
    for fi36 := 0; fi36 < 17; fi36++ {
        switch (fi36 * 2147483647) {
        case -982:
            switch "" {
            case "hello":
                v37 := GenS2{F0: 640, F1: v2, F2: 121}
                _ = v37
            case "a":
                if v5 {
                    v38 := GenS2{F0: (v24 ^ uint32(v33)), F1: 0, F2: 96}
                    v33 = (v33 * v33)
                    for fi39 := 0; fi39 < 9; fi39++ {
                        v2 = 0
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi39
                    }
                    _ = v38
                }
            }
        case 32768:
            var v40 uint8 = 42
            _ = v40
        case -1000:
            v41 := GenS3{F0: v35, F1: 695}
            _ = v41
        }
        if v34 {
            v34 = ("\n\t" < "")
            v42 := GenS3{F0: 0, F1: 341}
            switch "" {
            case "a":
                fmt.Println(v24)
            }
            _ = v42
        } else {
            v5 = (0 >= 2)
            fmt.Println(v24, v34)
        }
        for fi43 := 0; fi43 < 10; fi43++ {
            var v44 uint64 = 1
            v45 := GenS1{F0: (v33 << ((877) & 63)), F1: 0}
            var v46 int16 = (^(v2))
            _ = fi43
            _ = v44
            _ = v45
            _ = v46
        }
        _ = fi36
    }
    for fi47 := 0; fi47 < 3; fi47++ {
        if (1 > 1) {
            if true {
                var v48 int64 = (v35 + 0)
                v2 = v2
                for fi49 := 0; fi49 < 3; fi49++ {
                    v1 = append(v1, fi49)
                    v1 = append(v1, fi49)
                    var v50 int64 = (v48 * 9223372036854775807)
                    _ = fi49
                    _ = v50
                }
                _ = v48
            }
            var v51 uint16 = (v33 % (((v33 ^ (v33 - (v33 * (v33 % ((v33) | 1)))))) | 1))
            _ = v51
        } else {
            for fi52 := 0; fi52 < 6; fi52++ {
                var v53 int32 = -698
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch "\n\t" {
                case "foo bar":
                    v54 := GenS2{F0: v24, F1: v2, F2: 27}
                    _ = v54
                default:
                    v55 := v3["foo bar"]
                    _ = v55
                }
                var v56 float64 = 0.0
                if false {
                    v2 = 42
                    v33 += (v33 ^ 60)
                    switch fi52 {
                    case 2147483647:
                        var v57 uint64 = 1
                        _ = v57
                    case 4294967295:
                        v33 += (v33 & v33)
                    case 42:
                        var v58 float64 = v56
                        _ = v58
                    default:
                        v59 := GenS4{F0: v24, F1: "hello", F2: int32(v2)}
                        _ = v59
                    }
                    v5 = (!(!(255 == 255)))
                    fmt.Println(v56, v35)
                    v60 := v3[("" + "hello")]
                    _ = v60
                } else {
                    switch fi47 {
                    case -1:
                        var v61 uint = 255
                        _ = v61
                    case 2147483647:
                        v53 += (^(v53))
                    case -427:
                        if false {
                            v62 := GenS3{F0: (v35 * (v35 | 0)), F1: 2}
                            v63 := GenS1{F0: v33, F1: 844}
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                v3["hello"] = true
                                fmt.Println(v35, v24, v2)
                            }
                            _ = v62
                            _ = v63
                        }
                    }
                    var v64 string = "hello"
                    _ = v64
                }
                var v65 float64 = float64(0.0)
                v66 := GenS5{F0: fi52}
                _ = fi52
                _ = v53
                _ = v56
                _ = v65
                _ = v66
            }
            var v67 uint32 = 79
            var v68 string = "a"
            _ = v67
            _ = v68
        }
        v1 = append(v1, -1)
        v69 := GenS1{F0: v33, F1: 413}
        _ = fi47
        _ = v69
    }
    for fi70 := 0; fi70 < 11; fi70++ {
        if (uint16(v24) < (v33 % ((255) | 1))) {
            if v34 {
                var v71 int8 = 100
                for fi72 := 0; fi72 < 3; fi72++ {
                    v73 := GenS2{F0: (v24 + 0), F1: (v2 & (v2 % ((1) | 1))), F2: -1}
                    _ = fi72
                    _ = v73
                }
                _ = v71
            } else {
                var v74 uint8 = 0
                v75 := GenS2{F0: 0, F1: (v2 & 1), F2: -128}
                _ = len(v1)
                _ = v74
                _ = v75
            }
            v76 := GenS4{F0: (v24 >> ((799) & 63)), F1: "a", F2: 0}
            if v34 {
                switch 2147483647 {
                case 629:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 9223372036854775807:
                    v5 = v34
                case 0:
                    v77 := GenS3{F0: v35, F1: 42}
                    _ = v77
                }
                v1 = append(v1, (^(fi70)))
            } else {
                for fi78 := 0; fi78 < 12; fi78++ {
                    var v79 uint32 = v24
                    _ = fi78
                    _ = v79
                }
                for fi80 := 0; fi80 < 19; fi80++ {
                    switch 425 {
                    case -34:
                        v1 = append(v1, fi80)
                    case 127:
                        switch (("a" + "hello") + ("a" + ("" + "a"))) {
                        case "a":
                            var v81 uint64 = 2
                            _ = v81
                        case "foo bar":
                            var v82 float64 = ((1.5 / 1.0) * 1.5)
                            _ = v82
                        default:
                            var v83 int64 = -1
                            _ = v83
                        }
                    }
                    var v84 int64 = int64(v24)
                    _ = fi80
                    _ = v84
                }
                v85 := GenS2{F0: uint32(v33), F1: v2, F2: 127}
                _ = v85
            }
            _ = v76
        } else {
            for fi86 := 0; fi86 < 14; fi86++ {
                if v34 {
                    var v87 bool = false
                    _ = v87
                } else {
                    var v88 float32 = 1e10
                    v2 = -538
                    v24 -= (v24 | (v24 & v24))
                    _ = v88
                }
                v89 := GenS6{F0: int32(v33), F1: int32(v35), F2: 42}
                v34 = v5
                _ = fi86
                _ = v89
            }
            var v90 int8 = int8(fi70)
            _ = v90
        }
        fmt.Println(fi70)
        for fi91 := 0; fi91 < 14; fi91++ {
            v92 := GenS7{F0: v34, F1: (v33 * v33)}
            v93 := v3["\n\t"]
            switch fi91 {
            case 252:
                if v5 {
                    fmt.Println(fi91, v5)
                    v4[v24] = 42
                } else {
                    switch "" {
                    case "hello":
                        var v94 uint16 = v33
                        _ = v94
                    case "":
                        v95 := GenS8{F0: 158, F1: v33}
                        _ = v95
                    case "\n\t":
                        v93 = v34
                    }
                    delete(v3, "hello")
                }
            case -73:
                var v96 string = ("\n\t" + "")
                _ = v96
            default:
                var v97 bool = ("hello" <= "")
                _ = v97
            }
            _ = fi91
            _ = v92
            _ = v93
        }
        _ = fi70
    }
    var v98 int16 = (-(v2))
    fmt.Println(v98, v34, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v98
}
