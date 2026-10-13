package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint8; F1 int64; F2 uint }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int8; F1 int32; F2 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS8 struct { F0 int64; F1 uint }

func main() {
    var v1 float32 = ((0.0 / 0.1) / ((-2.25 / 1.0) / 1e10))
    var v2 float64 = 1.5
    var v3 uint32 = 0
    var v4 int64 = -9223372036854775808
    var v5 bool = (42 > 1)
    v1 = v1
    v6 := GenS1{F0: 42, F1: 908, F2: 2147483647}
    switch 32767 {
    case 9223372036854775807:
        fmt.Println(v5, v1)
    case 4294967295:
        var v7 uint16 = 65535
        _ = v7
    case -9223372036854775808:
        switch 93 {
        case -140:
            var v8 uint = 895
            _ = v8
        case 32768:
            v9 := GenS1{F0: 32767, F1: 1, F2: 42}
            _ = v9
        case 622:
            var v10 int = 42
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS2{F0: 42, F1: (645 != 1)}
        for fi13 := 0; fi13 < 5; fi13++ {
            if (!(v2 != (v2 / v2))) {
                var v14 int8 = 127
                _ = v14
            } else {
                var v15 int = (fi13 / ((-459) | 1))
                var v16 float64 = v2
                _ = v15
                _ = v16
            }
            _ = fi13
        }
        fmt.Println(v4, v3, fi11)
        _ = fi11
        _ = v12
    }
    if (true && (!(false && (v4 >= 328)))) {
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 int32 = 2147483647
            v18 = v18
            _ = fi17
            _ = v18
        }
    } else {
        fmt.Println(v2)
        switch "hello" {
        case "hello":
            for fi19 := 0; fi19 < 8; fi19++ {
                v4 -= -138
                if v5 {
                    switch "foo bar" {
                    case "foo bar":
                        v1 = (v1 * v1)
                    default:
                        for fi20 := 0; fi20 < 17; fi20++ {
                            v4 = 999
                            v4 += v4
                            var v21 uint64 = 0
                            _ = fi20
                            _ = v21
                        }
                    }
                    for fi22 := 0; fi22 < 20; fi22++ {
                        v3 += (v3 * 549)
                        v23 := []int{fi22}
                        switch (fi19 + (fi22 + -549)) {
                        case 361:
                            switch fi19 {
                            case 0:
                                if len(v23) > 0 {
                                    v23[0] = 32767
                                }
                            case 32768:
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                            case -2147483648:
                                v24 := GenS3{F0: 127, F1: -1, F2: 4294967295}
                                _ = v24
                            }
                        }
                        switch (("foo bar" + "hello") + "\n\t") {
                        case "foo bar":
                            switch (fi22 % ((2147483647) | 1)) {
                            case 42:
                                v25 := GenS2{F0: 18446744073709551615, F1: (32767 != -486)}
                                _ = v25
                            case 32768:
                                v5 = v5
                            case -1:
                                if (2 > 42) {
                                    v4 -= -9223372036854775808
                                } else {
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    fmt.Println(fi22, v2, v3)
                                    v23 = append(v23, -1)
                                }
                            default:
                                if v5 {
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    v26 := GenS3{F0: 255, F1: (v4 & 852), F2: 0}
                                    _ = v26
                                }
                            }
                        }
                        var v27 bool = v5
                        switch "foo bar" {
                        case "":
                            var v28 bool = (!(42 == 42))
                            _ = v28
                        case "\n\t":
                            _ = len(v23)
                            v29 := GenS3{F0: 240, F1: (v4 * 1), F2: 255}
                            if v27 {
                                if (25 != 42) {
                                    if v27 {
                                        v30 := GenS4{F0: 2147483647, F1: 62}
                                        _ = v30
                                    }
                                } else {
                                    if len(v23) > 0 {
                                        v23[0] = -650
                                    }
                                    if v27 {
                                        v5 = v5
                                        v1 -= ((0.0 / v1) * (v1 - (v1 / 3.14159)))
                                        var v31 uint64 = 2
                                        _ = v31
                                    } else {
                                        v1 = v1
                                    }
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    v32 := GenS2{F0: 186, F1: (v27 && false)}
                                    if v5 {
                                        if v27 {
                                            if v5 {
                                                if v5 {
                                                    if true {
                                                        if v5 {
                                                            var v33 int64 = 0
                                                            fmt.Println(v5, v27)
                                                            switch ("foo bar" + (("\n\t" + "hello") + "foo bar")) {
                                                            case "a":
                                                                if (4294967295 >= 18446744073709551615) {
                                                                    var v34 uint16 = 2
                                                                    if len(v23) >= 2 {
                                                                        v23 = v23[1:len(v23)]
                                                                    }
                                                                    _ = v34
                                                                }
                                                            case "hello":
                                                                v4 = 0
                                                            }
                                                            if v5 {
                                                                if (!(!false)) {
                                                                    var v35 int32 = 42
                                                                    _ = v35
                                                                }
                                                                switch fi19 {
                                                                case 4294967295:
                                                                    var v36 bool = v5
                                                                    _ = v36
                                                                case -739:
                                                                    var v37 int = (fi19 * int(v33))
                                                                    _ = v37
                                                                case 32768:
                                                                    var v38 uint8 = 127
                                                                    _ = v38
                                                                }
                                                            }
                                                            v4 = v33
                                                            var v39 string = ("a" + "\n\t")
                                                            _ = v33
                                                            _ = v39
                                                        } else {
                                                            var v40 int32 = -848
                                                            if len(v23) >= 2 {
                                                                v23 = v23[1:len(v23)]
                                                            }
                                                            v27 = false
                                                            _ = v40
                                                        }
                                                        var v41 uint64 = 18446744073709551615
                                                        _ = v41
                                                    }
                                                }
                                                if v27 {
                                                    v42 := GenS1{F0: 42, F1: 2, F2: fi19}
                                                    var v43 uint64 = 42
                                                    if len(v23) >= 2 {
                                                        v23 = v23[1:len(v23)]
                                                    }
                                                    _ = v42
                                                    _ = v43
                                                }
                                                v44 := GenS4{F0: -1, F1: 1}
                                                var v45 uint16 = 306
                                                _ = v44
                                                _ = v45
                                            } else {
                                                if (122 > -82) {
                                                    v4 += -9223372036854775808
                                                }
                                                var v46 float64 = v2
                                                v2 -= 1e10
                                                _ = v46
                                            }
                                            v3 -= (v3 - 2)
                                            v4 += v4
                                            v47 := GenS5{F0: 1, F1: 1, F2: 115}
                                            var v48 uint8 = 255
                                            _ = v47
                                            _ = v48
                                        }
                                    } else {
                                        v23 = append(v23, 587)
                                        switch fi19 {
                                        case 960:
                                            v49 := GenS5{F0: 0, F1: -481, F2: -585}
                                            _ = v49
                                        case 128:
                                            v27 = true
                                        case 32768:
                                            switch "hello" {
                                            case "hello":
                                                fmt.Println(fi19, v5, v1)
                                            case "foo bar":
                                                if v27 {
                                                    var v50 int = (-(fi22))
                                                    _ = len(v23)
                                                    v51 := GenS5{F0: 100, F1: 96, F2: -580}
                                                    switch "foo bar" {
                                                    case "a":
                                                        v23 = append(v23, -102)
                                                    case "hello":
                                                        v5 = (!(v3 != 2147483648))
                                                    default:
                                                        fmt.Println(fi22, v1)
                                                        if len(v23) >= 2 {
                                                            v23 = v23[1:len(v23)]
                                                        }
                                                        v50 += 0
                                                    }
                                                    switch ("a" + "a") {
                                                    case "foo bar":
                                                        var v52 int16 = -343
                                                        _ = v52
                                                    }
                                                    _ = v50
                                                    _ = v51
                                                }
                                            case "a":
                                                v23 = append(v23, -904)
                                            }
                                        }
                                    }
                                    v53 := GenS3{F0: 128, F1: -1, F2: 18446744073709551615}
                                    v54 := GenS1{F0: 65535, F1: 42, F2: fi19}
                                    _ = len(v23)
                                    _ = v32
                                    _ = v53
                                    _ = v54
                                }
                                var v55 bool = false
                                if (!(!(float64((-2.25 - v1)) >= v2))) {
                                    v27 = (v2 > v2)
                                    var v56 uint64 = 18446744073709551615
                                    v5 = v27
                                    v57 := GenS3{F0: uint8(v3), F1: 0, F2: 307}
                                    _ = v56
                                    _ = v57
                                } else {
                                    v2 = (0.1 * v2)
                                    fmt.Println(v3)
                                }
                                _ = v55
                            }
                            _ = v29
                        default:
                            _ = len(v23)
                        }
                        var v58 uint64 = 42
                        if (1 > 32767) {
                            var v59 int32 = -1
                            v60 := GenS3{F0: 127, F1: 160, F2: 2}
                            v23 = append(v23, fi22)
                            _ = v59
                            _ = v60
                        }
                        _ = fi22
                        _ = v23
                        _ = v27
                        _ = v58
                    }
                } else {
                    v61 := GenS6{F0: "a"}
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v4, v4)
                    case "a":
                        if (4294967295 <= 0) {
                            var v62 bool = (v3 <= (^(v3)))
                            _ = v62
                        }
                    case "hello":
                        if true {
                            v63 := []int{(^(fi19)), fi19}
                            var v64 uint64 = 620
                            _ = v63
                            _ = v64
                        } else {
                            var v65 float64 = (v2 / ((0.1 - v2) - (v2 * (0.0 + v2))))
                            _ = v65
                        }
                    default:
                        var v66 uint8 = 1
                        _ = v66
                    }
                    _ = v61
                }
                switch 32767 {
                case 643:
                    var v67 int8 = 42
                    _ = v67
                case 1:
                    v68 := []int{}
                    _ = v68
                case 127:
                    v4 -= (v4 - -1)
                default:
                    fmt.Println(v3, v2)
                }
                _ = fi19
            }
        }
    }
    for fi69 := 0; fi69 < 10; fi69++ {
        switch (-(fi69)) {
        case 32767:
            v70 := GenS5{F0: 100, F1: 708, F2: 1}
            _ = v70
        case -1:
            if false {
                v5 = false
                var v71 int32 = -1
                _ = v71
            } else {
                v72 := GenS7{F0: 0, F1: v3, F2: 65535}
                _ = v72
            }
        default:
            v73 := []int{(fi69 + (fi69 / ((int(v4)) | 1))), 908}
            _ = v73
        }
        _ = fi69
    }
    v74 := []int{42, 686, 4294967295, 127}
    if len(v74) >= 2 {
        v74 = v74[1:len(v74)]
    }
    var v75 int = 0
    v76 := GenS8{F0: (v4 + (v4 >> ((621) & 63))), F1: 255}
    v5 = true
    v77 := GenS8{F0: 529, F1: 801}
    for fi78 := 0; fi78 < 10; fi78++ {
        _ = len(v74)
        var v79 float64 = 3.14159
        var v80 uint64 = 2
        _ = fi78
        _ = v79
        _ = v80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v74:", len(v74), v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v75
    _ = v76
    _ = v77
}
