package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS5 struct { F0 bool; F1 int16; F2 int16 }
type GenS6 struct { F0 uint }
func genFn0(p0 int64) int32 {
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(p0, fi6)
        v7 := []int{(-(fi6)), fi6}
        for fi8 := 0; fi8 < 5; fi8++ {
            if false {
                v7 = append(v7, (fi6 + (fi8 * 32767)))
            }
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    v9 := GenS1{F0: 1, F1: 2, F2: 70}
    _ = p0
    _ = v9
    return -369
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 int32 = -350
    var v4 uint = 65535
    var v5 float64 = float64(0.1)
    v10 := GenS1{F0: 2147483647, F1: uint64(v4), F2: 0}
    var v11 uint64 = 0
    for fi12 := 0; fi12 < 13; fi12++ {
        if true {
            v13 := []int{}
            var v14 string = "a"
            var v15 uint32 = 2147483648
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    v11 += (^(v11))
    var v16 float64 = v5
    fmt.Println(v11)
    for fi17 := 0; fi17 < 10; fi17++ {
        if (uint8(v4) == 255) {
            fmt.Println(v11, fi17, v5)
            var v18 int16 = 1000
            if true {
                v3 = genFn0(-469)
            }
            _ = v18
        } else {
            switch "a" {
            case "":
                var v19 int32 = (v3 * v3)
                _ = v19
            case "a":
                for fi20 := 0; fi20 < 11; fi20++ {
                    switch 2147483647 {
                    case -1:
                        switch "foo bar" {
                        case "\n\t":
                            var v21 string = "a"
                            _ = v21
                        case "foo bar":
                            fmt.Println(v4, v3)
                        case "hello":
                            if false {
                                v16 -= (v5 / ((3.14159 * ((float64((3.14159 + 0.0)) * v5) + 1.5)) / ((float64(((((0.0 / 1.5) - 3.14159) * -2.25) - (3.14159 + ((-2.25 + 1e10) - (-2.25 + ((0.1 - 1.5) / (((1.5 - 3.14159) - (((-2.25 - float32(v16)) * (float32(((0.1 * v5) - v16)) / ((-2.25 / 1.0) - 1e10))) - (((1.5 * float32(v5)) / (-2.25 - 1e10)) - 1.5))) / 1e10))))))) * v5) * (v16 + v5))))
                                v16 -= ((v5 + 1.5) + v16)
                                v5 = v5
                            }
                        default:
                            v3 -= v3
                        }
                    case 301:
                        if (!true) {
                            var v22 string = "a"
                            _ = v22
                        }
                    case 127:
                        if (-1 <= -817) {
                            switch ("\n\t" + "foo bar") {
                            case "":
                                v16 += v5
                            }
                        }
                    }
                    _ = fi20
                }
            }
            if (3.14159 != 0.1) {
                v23 := []int{(^(fi17)), 127}
                for fi24 := 0; fi24 < 3; fi24++ {
                    fmt.Println(fi24, fi24)
                    var v25 uint32 = 65535
                    v3 = (v3 % ((v3) | 1))
                    var v26 uint8 = 128
                    v27 := GenS1{F0: 65535, F1: 42, F2: v26}
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                v3 -= (^(v3))
                _ = v23
            }
        }
        v3 = v3
        v2[2147483648] = fi17
        _ = fi17
    }
    v28 := GenS1{F0: 0, F1: 819, F2: 127}
    v29 := v2[2147483648]
    if true {
        switch v29 {
        case 128:
            v11 = (^(v11))
        case 1:
            switch 127 {
            case -575:
                v30 := GenS1{F0: 2, F1: (v11 ^ 18446744073709551615), F2: 72}
                _ = v30
            default:
                if ((42 == 1) || true) {
                    switch (v29 / ((-252) | 1)) {
                    case 42:
                        _ = len(v1)
                    case 1:
                        if true {
                            switch (v29 * (v29 | (v29 / (((v29 / ((v29) | 1))) | 1)))) {
                            case -227:
                                fmt.Println(v16, v29)
                            }
                        }
                    }
                    v5 = float64(1e10)
                } else {
                    switch "hello" {
                    case "a":
                        var v31 bool = (!(-2.25 == ((((0.0 / 1.0) + float32((float64((3.14159 + 0.0)) / 1.5))) + (3.14159 / 1.5)) * float32(v5))))
                        _ = v31
                    case "hello":
                        v29 -= int(v11)
                    case "\n\t":
                        var v32 uint64 = (v11 << ((v4) & 63))
                        _ = v32
                    default:
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            if false {
                                _ = len(v2)
                                switch ("hello" + "") {
                                case "\n\t":
                                    v33 := GenS1{F0: 42, F1: (v11 ^ uint64(v4)), F2: 1}
                                    _ = v33
                                case "hello":
                                    if true {
                                        v34 := GenS2{F0: -1, F1: 42}
                                        v11 += (v11 | (^(v11)))
                                        var v35 uint = (v4 >> ((v4) & 63))
                                        _ = v34
                                        _ = v35
                                    }
                                default:
                                    for fi36 := 0; fi36 < 14; fi36++ {
                                        var v37 float64 = v16
                                        v11 += 657
                                        var v38 uint32 = 0
                                        _ = fi36
                                        _ = v37
                                        _ = v38
                                    }
                                }
                            }
                        case "hello":
                            for fi39 := 0; fi39 < 5; fi39++ {
                                for fi40 := 0; fi40 < 9; fi40++ {
                                    var v41 bool = (42 == -1)
                                    _ = fi40
                                    _ = v41
                                }
                                v29 += (v29 + 32768)
                                var v42 int8 = -1
                                _ = fi39
                                _ = v42
                            }
                        case "foo bar":
                            for fi43 := 0; fi43 < 13; fi43++ {
                                for fi44 := 0; fi44 < 16; fi44++ {
                                    switch "a" {
                                    case "hello":
                                        v11 -= (v11 >> (((v4 + (v4 ^ v4))) & 63))
                                    case "\n\t":
                                        v5 += 3.14159
                                    }
                                    switch "\n\t" {
                                    case "":
                                        if true {
                                            if ("a" > "hello") {
                                                v3 = (v3 * (^(v3)))
                                                v45 := []int{(v29 / ((fi44) | 1)), (fi43 ^ 9223372036854775807)}
                                                fmt.Println(v11, v16)
                                                var v46 int16 = -1
                                                var v47 int = 9223372036854775807
                                                if ((!((0.0 + 0.1) >= 3.14159)) && true) {
                                                    var v48 string = "foo bar"
                                                    v29 = (v29 | fi44)
                                                    var v49 bool = (!((v4 & uint(v11)) != 4294967295))
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    v45 = append(v45, fi44)
                                                }
                                                _ = v45
                                                _ = v46
                                                _ = v47
                                            } else {
                                                if ((!false) && true) {
                                                    v50 := GenS3{F0: 9223372036854775807}
                                                    v51 := []int{}
                                                    _ = v50
                                                    _ = v51
                                                }
                                                delete(v2, 1)
                                                if true {
                                                    v52 := []int{(fi43 | (fi44 * 4294967295))}
                                                    var v53 bool = true
                                                    v29 += fi44
                                                    if len(v52) > 0 {
                                                        v52[0] = (fi44 ^ (fi43 ^ (fi43 % ((fi44) | 1))))
                                                    }
                                                    _ = v52
                                                    _ = v53
                                                } else {
                                                    var v54 uint8 = 1
                                                    switch -922 {
                                                    case 677:
                                                        fmt.Println(v3, v16, fi44)
                                                    case 969:
                                                        _ = len(v1)
                                                    case 390:
                                                        fmt.Println(v54, v11, v4)
                                                    }
                                                    _ = v54
                                                }
                                            }
                                            v55 := GenS3{F0: -1}
                                            switch ((((("" + "a") + "foo bar") + "\n\t") + ("a" + "a")) + "") {
                                            case "hello":
                                                if (true && true) {
                                                    v11 += (v11 % ((v11) | 1))
                                                    var v56 int64 = -774
                                                    v1[(v29 - 2147483647)] = 1
                                                    _ = v56
                                                }
                                            case "foo bar":
                                                var v57 uint8 = 2
                                                _ = v57
                                            }
                                            _ = v55
                                        } else {
                                            if (892 == -1) {
                                                var v58 uint = v4
                                                v59 := GenS1{F0: 2, F1: 767, F2: 255}
                                                _ = v58
                                                _ = v59
                                            } else {
                                                if false {
                                                    v11 = v11
                                                    delete(v1, (-(fi44)))
                                                    delete(v2, 2147483647)
                                                }
                                                v3 -= int32(v4)
                                                if ((false && (32768 != fi43)) && (true || (((0 >= uint64(v4)) || false) || (!false)))) {
                                                    if true {
                                                        if true {
                                                            v5 = (0.1 + 1.5)
                                                            var v60 string = (("a" + "foo bar") + "")
                                                            fmt.Println(v3)
                                                            _ = v60
                                                        }
                                                        var v61 float32 = (1.5 / 3.14159)
                                                        v4 -= (v4 >> ((v4) & 63))
                                                        _ = v61
                                                    }
                                                    fmt.Println(v11, v29)
                                                    var v62 uint = (v4 + v4)
                                                    v63 := []int{(fi44 * (fi44 - (fi43 | (-(fi44))))), (v29 & -378), fi43}
                                                    var v64 float64 = 0.0
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                }
                                            }
                                            v3 += v3
                                        }
                                    default:
                                        v65 := GenS2{F0: 42, F1: -465}
                                        _ = v65
                                    }
                                    if (v5 != v5) {
                                        if false {
                                            switch fi43 {
                                            case 2147483647:
                                                v66 := GenS4{F0: 4294967295, F1: (853 >= 42), F2: 127}
                                                _ = v66
                                            case 4294967295:
                                                delete(v1, 428)
                                            }
                                        }
                                        _ = len(v1)
                                    } else {
                                        var v67 uint = v4
                                        v5 = 0.1
                                        _ = v67
                                    }
                                    if (-1 >= 31) {
                                        v4 -= v4
                                    }
                                    _ = fi44
                                }
                                fmt.Println(v11)
                                v16 -= -2.25
                                _ = fi43
                            }
                        }
                    }
                }
            }
        }
    }
    for fi68 := 0; fi68 < 8; fi68++ {
        v5 = v5
        _ = fi68
    }
    for fi69 := 0; fi69 < 16; fi69++ {
        var v70 int64 = 42
        if true {
            if false {
                v71 := v2[655]
                _ = v71
            } else {
                v72 := GenS5{F0: true, F1: 0, F2: 42}
                switch (("hello" + ("\n\t" + "a")) + "\n\t") {
                case "a":
                    if (2 > 0) {
                        var v73 float64 = ((float64(-2.25) * float64(-2.25)) - (v5 - v5))
                        _ = v73
                    }
                case "":
                    v74 := []int{}
                    _ = v74
                }
                var v75 int = int(v4)
                _ = v72
                _ = v75
            }
        }
        _ = fi69
        _ = v70
    }
    v29 -= v29
    v76 := GenS6{F0: v4}
    v77 := []int{(v29 & 2147483647), v29, v29, 128}
    var v78 uint = (v4 ^ 0)
    delete(v2, 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v76
    _ = v77
    _ = v78
}
