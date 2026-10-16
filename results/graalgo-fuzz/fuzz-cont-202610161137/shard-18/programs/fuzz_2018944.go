package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint16; F2 uint }
type GenS6 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS7 struct { F0 uint32; F1 uint; F2 string }
type GenS8 struct { F0 int16; F1 uint16 }
type GenS9 struct { F0 int16; F1 uint; F2 int32 }

func main() {
    var v1 int16 = 42
    var v2 int32 = 0
    var v3 int64 = -9223372036854775808
    var v4 int8 = 0
    v5 := make(map[uint32]int)
    var v6 uint32 = 65535
    v2 -= (v2 + -1)
    _ = len(v5)
    var v7 uint = 4294967295
    if true {
        var v8 int = 32767
        _ = v8
    }
    var v9 int16 = v1
    for fi10 := 0; fi10 < 12; fi10++ {
        switch fi10 {
        case 42:
            var v11 int32 = (^(v2))
            _ = v11
        case 9223372036854775807:
            for fi12 := 0; fi12 < 15; fi12++ {
                switch "" {
                case "foo bar":
                    v3 = (v3 / ((int64(v7)) | 1))
                case "hello":
                    v4 += (^(v4))
                case "":
                    v13 := []int{}
                    _ = v13
                default:
                    if (!true) {
                        if ((726 != 42) && (true && (!true))) {
                            if false {
                                fmt.Println(fi12)
                                var v14 float64 = 3.14159
                                _ = v14
                            } else {
                                v3 = -9223372036854775808
                                var v15 uint32 = (^(v6))
                                var v16 int = 32768
                                v17 := []int{}
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            if false {
                                var v18 bool = (65535 == 42)
                                v19 := GenS1{F0: v6}
                                v20 := GenS2{F0: (true && v18), F1: fi10}
                                v6 = (v6 - 318)
                                var v21 int = fi12
                                v1 = (v1 / ((-32768) | 1))
                                _ = v18
                                _ = v19
                                _ = v20
                                _ = v21
                            } else {
                                var v22 uint8 = 170
                                v23 := GenS1{F0: (v6 << ((42) & 63))}
                                var v24 int64 = -1
                                _ = len(v5)
                                fmt.Println(v3, v7, v24)
                                v25 := []int{fi12, (fi12 * (fi12 ^ (fi12 / (((fi10 / ((-2147483648) | 1))) | 1)))), (fi12 / (((^(fi12))) | 1)), 42}
                                _ = v22
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            var v26 float32 = (1e10 * (float32((0.1 / 1e10)) + 1e10))
                            _ = v26
                        }
                        switch (fi10 / (((fi12 * (fi10 >> ((v7) & 63)))) | 1)) {
                        case 32768:
                            v27 := GenS2{F0: true, F1: (^(fi10))}
                            _ = v27
                        }
                        v28 := []int{128, fi12}
                        _ = len(v28)
                        var v29 uint8 = 1
                        _ = v28
                        _ = v29
                    } else {
                        v30 := GenS2{F0: false, F1: fi12}
                        _ = v30
                    }
                }
                v2 -= (^(v2))
                var v31 int8 = v4
                v3 -= v3
                v4 = (v4 & (v31 ^ -1))
                v31 = (v31 ^ v31)
                if ((v1 < v9) || false) {
                    switch (fi10 - fi12) {
                    case -9223372036854775808:
                        if (255 == 42) {
                            v3 += v3
                        }
                    default:
                        switch fi10 {
                        case -236:
                            v32 := GenS1{F0: v6}
                            _ = v32
                        case 15:
                            var v33 uint8 = 168
                            _ = v33
                        case 937:
                            switch 1 {
                            case 127:
                                v34 := GenS2{F0: false, F1: 9223372036854775807}
                                _ = v34
                            case 2147483647:
                                var v35 int = (fi10 * fi10)
                                _ = v35
                            default:
                                var v36 float32 = (-2.25 * 1.5)
                                _ = v36
                            }
                        default:
                            var v37 int = (fi12 << ((4294967295) & 63))
                            var v38 int64 = v3
                            _ = v37
                            _ = v38
                        }
                    }
                    var v39 int = fi10
                    fmt.Println(v2)
                    _ = v39
                } else {
                    var v40 int = (fi12 | (fi10 ^ fi10))
                    var v41 uint32 = (v6 >> (((^(v7))) & 63))
                    v31 -= (v31 - v4)
                    _ = v40
                    _ = v41
                }
                _ = fi12
                _ = v31
            }
        }
        v42 := []int{(fi10 + (^(fi10))), 127, (fi10 - fi10)}
        if (!(65 < 42)) {
            switch ("" + "hello") {
            case "hello":
                for fi43 := 0; fi43 < 12; fi43++ {
                    var v44 uint16 = 1
                    _ = fi43
                    _ = v44
                }
            case "foo bar":
                v3 += v3
            default:
                var v45 float32 = (1.5 / 1e10)
                _ = v45
            }
            if (((true || true) && (!(!true))) && true) {
                v4 += v4
                fmt.Println(v1)
                switch fi10 {
                case -749:
                    v7 = (v7 * v7)
                case -811:
                    v1 -= (^(v1))
                default:
                    fmt.Println(fi10)
                }
            } else {
                fmt.Println(v4, v6)
            }
            if (!(float32((1e10 / 1.0)) == 1e10)) {
                var v46 int16 = (-(v1))
                for fi47 := 0; fi47 < 3; fi47++ {
                    switch fi47 {
                    case -85:
                        var v48 uint = (v7 >> ((42) & 63))
                        _ = v48
                    default:
                        _ = len(v5)
                    }
                    if false {
                        switch fi47 {
                        case -955:
                            v5[(^(v6))] = fi10
                        case -198:
                            v3 -= int64(v7)
                        case 2147483647:
                            v49 := GenS1{F0: (v6 + v6)}
                            _ = v49
                        }
                        fmt.Println(v1, fi47, v2)
                        fmt.Println(v1, v1, v9)
                        if false {
                            _ = len(v42)
                        }
                        if true {
                            fmt.Println(fi47, fi47)
                            var v50 uint64 = 707
                            switch "\n\t" {
                            case "foo bar":
                                v5[v6] = (fi10 + fi47)
                            case "a":
                                v1 -= 822
                                if (float64(0.0) > (((float64(0.1) / (3.14159 + 1.5)) * 1e10) / -2.25)) {
                                    if false {
                                        if true {
                                            var v51 float64 = 1e10
                                            v52 := GenS3{F0: (v6 >> (((^(v7))) & 63)), F1: (v2 - v2), F2: 255}
                                            _ = v51
                                            _ = v52
                                        } else {
                                            var v53 uint8 = 211
                                            var v54 float32 = 0.1
                                            v42 = append(v42, (fi47 | (fi47 % ((fi47) | 1))))
                                            _ = v53
                                            _ = v54
                                        }
                                        fmt.Println(v7, v46, v7)
                                        var v55 int32 = v2
                                        _ = v55
                                    } else {
                                        v46 += (v46 | 1)
                                        if len(v42) >= 2 {
                                            v42 = v42[1:len(v42)]
                                        }
                                    }
                                    if (v46 != -408) {
                                        _ = len(v5)
                                        v42 = append(v42, fi47)
                                    } else {
                                        var v56 string = "\n\t"
                                        var v57 uint = v7
                                        switch v56 {
                                        case "hello":
                                            v57 = (^(v7))
                                        case "a":
                                            v50 = (^(v50))
                                        case "\n\t":
                                            fmt.Println(v6)
                                        default:
                                            var v58 string = v56
                                            _ = v58
                                        }
                                        _ = v56
                                        _ = v57
                                    }
                                    fmt.Println(fi10, v46)
                                } else {
                                    fmt.Println(v7)
                                }
                                fmt.Println(v9, v9)
                            default:
                                fmt.Println(v3, v50, v4)
                            }
                            _ = v50
                        } else {
                            v5[4294967295] = 32767
                            var v59 string = "hello"
                            _ = v59
                        }
                    }
                    _ = fi47
                }
                _ = v46
            } else {
                v6 = uint32(v7)
            }
        } else {
            fmt.Println(v1, v2, v3)
            var v60 uint16 = 466
            fmt.Println(v9)
            _ = v60
        }
        _ = fi10
        _ = v42
    }
    for fi61 := 0; fi61 < 11; fi61++ {
        if (!false) {
            if true {
                v62 := GenS2{F0: (!false), F1: 1}
                _ = v62
            } else {
                for fi63 := 0; fi63 < 6; fi63++ {
                    var v64 int = (fi61 % ((fi63) | 1))
                    v9 += int16(v2)
                    v64 = fi63
                    v7 += v7
                    v65 := GenS2{F0: (("a" > "") || false), F1: int(v3)}
                    _ = fi63
                    _ = v64
                    _ = v65
                }
            }
        } else {
            for fi66 := 0; fi66 < 8; fi66++ {
                v67 := GenS4{F0: (v4 * v4)}
                _ = fi66
                _ = v67
            }
            if true {
                fmt.Println(v9, v9, v4)
                for fi68 := 0; fi68 < 1; fi68++ {
                    v9 -= v9
                    v1 -= (v1 - 1)
                    _ = fi68
                }
                fmt.Println(v7, v4)
            } else {
                v69 := GenS5{F0: fi61, F1: 255, F2: (v7 - (v7 + 0))}
                _ = v69
            }
            var v70 uint16 = 32767
            _ = v70
        }
        v3 = v3
        _ = fi61
    }
    fmt.Println(v3, v2, v2)
    v2 = v2
    v71 := GenS5{F0: 490, F1: 255, F2: 65535}
    for fi72 := 0; fi72 < 10; fi72++ {
        v2 = v2
        v73 := GenS1{F0: v6}
        _ = fi72
        _ = v73
    }
    v74 := GenS2{F0: (3.14159 == ((0.0 + 0.0) + 1.5)), F1: -9223372036854775808}
    v75 := GenS6{F0: v1, F1: 65535, F2: 2}
    v3 += (v3 | -9223372036854775808)
    switch 1 {
    case -9223372036854775808:
        v76 := GenS7{F0: v6, F1: 0, F2: "a"}
        _ = v76
    case -518:
        switch "foo bar" {
        case "a":
            fmt.Println(v2, v3, v4)
        case "":
            var v77 string = (("" + "\n\t") + "a")
            _ = v77
        case "\n\t":
            v9 -= (v9 & -489)
        default:
            for fi78 := 0; fi78 < 5; fi78++ {
                if (!false) {
                    v79 := GenS8{F0: v9, F1: 32767}
                    var v80 uint8 = 133
                    _ = v79
                    _ = v80
                }
                if true {
                    var v81 string = "hello"
                    v3 -= v3
                    _ = v81
                } else {
                    v6 += (v6 | v6)
                }
                switch "foo bar" {
                case "":
                    switch ("foo bar" + "hello") {
                    case "a":
                        fmt.Println(fi78, fi78)
                    case "hello":
                        v82 := GenS9{F0: (v1 * (-(v1))), F1: (v7 * (v7 | 556)), F2: -295}
                        _ = v82
                    default:
                        fmt.Println(v1, v3, v9)
                    }
                case "hello":
                    v6 -= 2147483647
                case "foo bar":
                    var v83 int64 = v3
                    _ = v83
                default:
                    fmt.Println(v9, v9, v7)
                }
                _ = fi78
            }
        }
    }
    var v84 bool = (false && (false || true))
    fmt.Println(v9, v3, v6)
    v85 := []int{128, 32767, 32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v84:", v84)
    fmt.Println("slicelen v85:", len(v85), v85)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v71
    _ = v74
    _ = v75
    _ = v84
    _ = v85
    _ = v9
}
