package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }
type GenS2 struct { F0 uint16; F1 uint8; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 uint8; F2 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64 }
func genFn0(p0 string, p1 string, p2 bool) string {
    v6 := GenS1{F0: 255, F1: (p2 && p2), F2: false}
    p1 = p1
    var v7 int8 = 42
    p1 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p1 + "")
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    v3 := []int{127, -149, 42}
    v4 := []int{2147483647, 32768}
    var v5 uint16 = 32767
    v8 := GenS2{F0: (v5 % ((65535) | 1)), F1: 127, F2: ("foo bar" != "foo bar")}
    v9 := GenS1{F0: 255, F1: ((-2.25 - 1e10) == 1e10), F2: true}
    v5 = v5
    if false {
        v10 := GenS2{F0: (v5 - (v5 + v5)), F1: 1, F2: ("a" > "foo bar")}
        _ = v10
    }
    fmt.Println(v5)
    v5 = 600
    if true {
        v11 := GenS1{F0: 42, F1: (((-2.25 * 3.14159) * float64(((-2.25 - 1e10) + (((float32(-2.25) + ((0.0 / -2.25) * 1.5)) / (((3.14159 - 0.0) + (0.0 + float32(0.1))) * 3.14159)) - 0.0)))) != (3.14159 + (((0.0 + -2.25) / (float64(-2.25) + 1e10)) * 1.5))), F2: true}
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 string = "\n\t"
        fmt.Println(fi12, v5)
        _ = fi12
        _ = v13
    }
    v5 -= 255
    var v14 int16 = 1
    switch "foo bar" {
    case "hello":
        for fi15 := 0; fi15 < 6; fi15++ {
            fmt.Println(v5, v5)
            if false {
                if len(v4) > 0 {
                    v4[0] = (fi15 | 2147483647)
                }
                var v16 int32 = -2147483648
                _ = v16
            }
            switch "\n\t" {
            case "hello":
                v14 += (-(v14))
            case "a":
                var v17 int32 = -2147483648
                _ = v17
            }
            _ = fi15
        }
    case "a":
        delete(v1, "hello")
    case "":
        v14 = v14
    default:
        v5 += (v5 + (v5 ^ (v5 | (v5 & (^(v5))))))
    }
    var v18 uint8 = 49
    _ = len(v2)
    v14 += 0
    if (9223372036854775807 >= 0) {
        fmt.Println(v14, v18)
        v5 = (v5 ^ v5)
    } else {
        if true {
            v14 = v14
            if true {
                v19 := v1["hello"]
                for fi20 := 0; fi20 < 7; fi20++ {
                    if true {
                        if (false && true) {
                            for fi21 := 0; fi21 < 12; fi21++ {
                                v22 := GenS3{F0: 2147483647}
                                var v23 int8 = 42
                                if true {
                                    if (4294967295 > 2147483647) {
                                        v24 := GenS2{F0: v5, F1: v18, F2: true}
                                        _ = v24
                                    } else {
                                        switch fi20 {
                                        case 29:
                                            v25 := GenS2{F0: (v5 << ((2) & 63)), F1: (v18 % (((^(v18))) | 1)), F2: (!(v14 >= v14))}
                                            _ = v25
                                        case 42:
                                            _ = len(v4)
                                        default:
                                            var v26 int8 = -128
                                            _ = v26
                                        }
                                        var v27 float64 = 1e10
                                        var v28 int16 = v14
                                        _ = v27
                                        _ = v28
                                    }
                                    delete(v2, fi21)
                                    if true {
                                        var v29 uint = 2
                                        v3 = append(v3, 610)
                                        fmt.Println(v19, fi21)
                                        _ = v29
                                    } else {
                                        v18 -= (v18 * v18)
                                        v30 := v2[(v19 - 698)]
                                        fmt.Println(v19, fi21, fi20)
                                        _ = v30
                                    }
                                } else {
                                    v23 -= -1
                                }
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        }
                        var v31 int = (^(v19))
                        if false {
                            for fi32 := 0; fi32 < 11; fi32++ {
                                v33 := GenS3{F0: -942}
                                switch (v31 * 128) {
                                case -434:
                                    var v34 string = "a"
                                    _ = v34
                                case -9223372036854775808:
                                    fmt.Println(v19, fi20, v14)
                                case -75:
                                    switch 1 {
                                    case -9223372036854775808:
                                        if ((v5 & v5) <= uint16(v18)) {
                                            v35 := GenS4{F0: -635}
                                            if true {
                                                if (1.5 <= (1e10 + (0.0 - (3.14159 / 3.14159)))) {
                                                    var v36 uint = 2
                                                    switch fi20 {
                                                    case 220:
                                                        v37 := GenS1{F0: 34, F1: ((v36 + v36) != v36), F2: false}
                                                        _ = v37
                                                    case -9223372036854775808:
                                                        v31 = 32767
                                                    }
                                                    _ = len(v2)
                                                    _ = v36
                                                } else {
                                                    if (42 != v5) {
                                                        if (-503 >= -9223372036854775808) {
                                                            v38 := GenS5{F0: 100, F1: 209, F2: 65535}
                                                            fmt.Println(fi20)
                                                            _ = v38
                                                        } else {
                                                            v19 = (v31 * v19)
                                                            if (fi20 <= v31) {
                                                                fmt.Println(v14, fi32)
                                                                var v39 bool = false
                                                                var v40 uint8 = (v18 - v18)
                                                                _ = v39
                                                                _ = v40
                                                            } else {
                                                                var v41 uint = 1
                                                                var v42 uint64 = 42
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                            var v43 string = ""
                                                            _ = v43
                                                        }
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v44 uint8 = (v18 ^ (v18 << ((842) & 63)))
                                                            _ = v44
                                                        case "a":
                                                            v19 += v19
                                                        case "":
                                                            var v45 uint8 = (v18 * 166)
                                                            if len(v3) > 0 {
                                                                v3[0] = (fi32 + v19)
                                                            }
                                                            v45 += (v18 | v18)
                                                            _ = v45
                                                        default:
                                                            if true {
                                                                v18 = (v18 | (v18 << ((909) & 63)))
                                                                if false {
                                                                    var v46 int64 = -781
                                                                    v46 -= 0
                                                                    var v47 float64 = float64(0.0)
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                                var v48 uint8 = v18
                                                                if false {
                                                                    fmt.Println(v48, fi20, v18)
                                                                    v49 := GenS6{F0: v31}
                                                                    _ = v49
                                                                }
                                                                _ = v48
                                                            } else {
                                                                var v50 float64 = (1.5 - -2.25)
                                                                _ = v50
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v14, v18, v31)
                                                    var v51 uint32 = 65535
                                                    _ = v51
                                                }
                                                v2[(v31 / ((v31) | 1))] = 2
                                            }
                                            switch ("\n\t" + "foo bar") {
                                            case "":
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v52 uint64 = 844
                                                    _ = v52
                                                }
                                            case "\n\t":
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v53 string = "a"
                                                    _ = v53
                                                case "":
                                                    if (!true) {
                                                        switch ("\n\t" + ("\n\t" + "foo bar")) {
                                                        case "foo bar":
                                                            v31 = (v31 - v19)
                                                        }
                                                        v5 = v5
                                                    }
                                                case "a":
                                                    var v54 uint16 = v5
                                                    _ = v54
                                                }
                                            case "hello":
                                                if len(v4) > 0 {
                                                    v4[0] = -2147483648
                                                }
                                            default:
                                                v18 = v18
                                            }
                                            var v55 float64 = float64(float32((0.0 - ((((1.5 + float64(0.1)) * (0.0 / 3.14159)) - (1e10 * 0.1)) * 3.14159))))
                                            v56 := GenS5{F0: 42, F1: (v18 % ((v18) | 1)), F2: 364}
                                            _ = v35
                                            _ = v55
                                            _ = v56
                                        }
                                    case 4294967295:
                                        v57 := GenS7{F0: fi32}
                                        _ = v57
                                    }
                                default:
                                    _ = len(v1)
                                    fmt.Println(fi32, v31)
                                    v14 = (v14 ^ v14)
                                }
                                v3 = append(v3, (fi32 >> ((uint(v5)) & 63)))
                                _ = fi32
                                _ = v33
                            }
                            switch "foo bar" {
                            case "hello":
                                var v58 uint8 = v18
                                _ = v58
                            case "":
                                v5 = (v5 / ((255) | 1))
                            case "\n\t":
                                v59 := GenS4{F0: 42}
                                _ = v59
                            default:
                                fmt.Println(v19, v14, v19)
                            }
                            for fi60 := 0; fi60 < 3; fi60++ {
                                var v61 string = ("hello" + ("" + "\n\t"))
                                _ = fi60
                                _ = v61
                            }
                        } else {
                            fmt.Println(v31, v14)
                        }
                        _ = v31
                    } else {
                        for fi62 := 0; fi62 < 5; fi62++ {
                            var v63 int64 = 0
                            v63 += (v63 + v63)
                            _ = fi62
                            _ = v63
                        }
                    }
                    if false {
                        var v64 float64 = 0.0
                        var v65 uint = 65535
                        fmt.Println(fi20, v64, v64)
                        _ = v64
                        _ = v65
                    } else {
                        _ = len(v3)
                        v14 += v14
                        var v66 string = "a"
                        _ = v66
                    }
                    if (!(!(v14 == 1000))) {
                        if (!false) {
                            v18 += 27
                            var v67 uint16 = 255
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v67
                        } else {
                            v14 -= (v14 | v14)
                        }
                        if (652 == uint(v18)) {
                            v68 := v2[v19]
                            fmt.Println(v14, v5, fi20)
                            _ = v68
                        }
                    } else {
                        for fi69 := 0; fi69 < 17; fi69++ {
                            fmt.Println(fi20)
                            _ = fi69
                        }
                    }
                    _ = fi20
                }
                _ = v19
            }
            if false {
                v5 = v5
                var v70 int16 = (^(v14))
                for fi71 := 0; fi71 < 20; fi71++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v72 uint = 255
                    _ = fi71
                    _ = v72
                }
                _ = v70
            }
        }
        var v73 uint = 18446744073709551615
        _ = v73
    }
    var v74 int32 = -2147483648
    v75 := GenS8{F0: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v74
    _ = v75
    _ = v8
    _ = v9
}
