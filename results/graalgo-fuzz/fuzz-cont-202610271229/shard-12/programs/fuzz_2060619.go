package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 int16; F1 int8; F2 string }
type GenS6 struct { F0 int16; F1 int8; F2 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint8; F1 bool; F2 bool }
type GenS9 struct { F0 int32; F1 bool; F2 uint8 }
type GenS10 struct { F0 int16; F1 uint8; F2 string }
type GenS11 struct { F0 int; F1 int16; F2 uint }

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]bool)
    var v3 float32 = 1e10
    var v4 int32 = 1
    var v5 string = "hello"
    if true {
        for fi6 := 0; fi6 < 8; fi6++ {
            v5 = v5
            v7 := GenS1{F0: 42, F1: v1}
            _ = fi6
            _ = v7
        }
        v1 = (v1 | (v1 & -128))
    }
    var v8 string = v5
    v4 -= (v4 - 2147483647)
    var v9 int8 = -128
    var v10 uint32 = 1
    if true {
        v3 += (v3 / ((((0.0 / ((0.1 - 1.5) - -2.25)) / v3) - v3) / v3))
    } else {
        var v11 uint8 = 36
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS1{F0: 890, F1: -128}
        _ = fi12
        _ = v13
    }
    switch v8 {
    case "hello":
        var v14 int32 = v4
        _ = v14
    }
    if false {
        v4 -= (-(v4))
        var v15 int16 = 636
        if true {
            v16 := []int{13, 32768, 9223372036854775807, -9223372036854775808}
            for fi17 := 0; fi17 < 19; fi17++ {
                if false {
                    v18 := GenS2{F0: 1, F1: fi17}
                    var v19 uint32 = 2
                    v20 := GenS3{F0: 42}
                    _ = v18
                    _ = v19
                    _ = v20
                }
                _ = fi17
            }
            if (v10 >= v10) {
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 uint8 = 0
                    var v23 uint = 65535
                    for fi24 := 0; fi24 < 14; fi24++ {
                        v2[(v10 - 42)] = false
                        v16 = append(v16, -194)
                        v25 := GenS4{F0: v1, F1: -1}
                        _ = fi24
                        _ = v25
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                if len(v16) > 0 {
                    v16[0] = 9223372036854775807
                }
                for fi26 := 0; fi26 < 2; fi26++ {
                    var v27 string = ("" + "foo bar")
                    _ = fi26
                    _ = v27
                }
            } else {
                if true {
                    v16 = append(v16, 32768)
                    v16 = append(v16, 128)
                    fmt.Println(v15, v9, v10)
                }
            }
            _ = v16
        } else {
            for fi28 := 0; fi28 < 16; fi28++ {
                fmt.Println(fi28, v15)
                var v29 bool = true
                _ = fi28
                _ = v29
            }
            var v30 uint16 = 597
            v10 = 65535
            _ = v30
        }
        _ = v15
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        if ((float64(v3) / ((((((1.5 + (1.5 - -2.25)) * 3.14159) / (float64(3.14159) / 1e10)) * 3.14159) / (3.14159 + 3.14159)) - ((1.5 * 0.0) / -2.25))) > 0.0) {
            if (false && (32767 < 255)) {
                switch "hello" {
                case "\n\t":
                    v32 := []int{}
                    _ = v32
                case "a":
                    if false {
                        if ((v9 % ((v1) | 1)) >= 42) {
                            v33 := GenS5{F0: int16(v4), F1: v1, F2: v5}
                            _ = v33
                        } else {
                            v5 = "foo bar"
                        }
                        for fi34 := 0; fi34 < 12; fi34++ {
                            if (797 != 0) {
                                fmt.Println(v1)
                            } else {
                                if false {
                                    v35 := []int{1}
                                    _ = v35
                                } else {
                                    v5 = v8
                                }
                                if true {
                                    v8 = ((v8 + v5) + (v8 + "foo bar"))
                                    if false {
                                        var v36 uint8 = 131
                                        _ = v36
                                    }
                                    v37 := []int{-9223372036854775808, (fi31 * (fi31 | (^(fi34)))), 128}
                                    _ = v37
                                } else {
                                    if (3.14159 > (1.5 + (-2.25 - (0.1 / 0.1)))) {
                                        var v38 string = ""
                                        var v39 uint8 = 255
                                        var v40 int64 = 230
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                    var v41 int32 = (v4 & v4)
                                    _ = v41
                                }
                                var v42 int32 = (v4 >> ((4294967295) & 63))
                                switch (fi31 % ((fi31) | 1)) {
                                case -2147483648:
                                    if true {
                                        if (false && (((v5 + (v5 + "foo bar")) + ("" + v5)) == "foo bar")) {
                                            switch -858 {
                                            case 0:
                                                fmt.Println(v9)
                                            }
                                            if (!false) {
                                                var v43 uint = 1
                                                _ = v43
                                            } else {
                                                switch fi31 {
                                                case 42:
                                                    v42 += 2147483647
                                                }
                                                v44 := GenS6{F0: 32767, F1: v9, F2: (!false)}
                                                v45 := GenS1{F0: -1, F1: int8(v4)}
                                                _ = v44
                                                _ = v45
                                            }
                                            v46 := GenS4{F0: (v9 + v1), F1: 1000}
                                            _ = v46
                                        } else {
                                            switch v5 {
                                            case "foo bar":
                                                var v47 uint32 = (^(v10))
                                                _ = v47
                                            case "":
                                                switch fi31 {
                                                case -9223372036854775808:
                                                    var v48 uint32 = (v10 * v10)
                                                    _ = v48
                                                case 32767:
                                                    v42 = (v4 + 42)
                                                case 128:
                                                    v4 = v42
                                                default:
                                                    if true {
                                                        var v49 int32 = v4
                                                        v4 = v42
                                                        _ = v49
                                                    } else {
                                                        if (0.1 <= 0.1) {
                                                            v4 += v4
                                                            switch -899 {
                                                            case 127:
                                                                if true {
                                                                    switch (((v8 + v8) + (v5 + "")) + ("foo bar" + (("\n\t" + v5) + "\n\t"))) {
                                                                    case "":
                                                                        var v50 uint64 = 34
                                                                        _ = v50
                                                                    case "\n\t":
                                                                        fmt.Println(fi34, v42)
                                                                    case "a":
                                                                        v9 -= (v1 ^ v1)
                                                                    default:
                                                                        v10 += (v10 << ((1) & 63))
                                                                        if false {
                                                                            v8 = "\n\t"
                                                                            v51 := []int{(^(fi34)), (fi34 ^ -1), fi31, (fi31 >> ((42) & 63))}
                                                                            v52 := GenS7{F0: "\n\t"}
                                                                            delete(v2, 1)
                                                                            _ = v51
                                                                            _ = v52
                                                                        } else {
                                                                            fmt.Println(v3, v1, v10)
                                                                            v53 := GenS8{F0: 2, F1: (("foo bar" + v5) > v8), F2: true}
                                                                            fmt.Println(v42, v10, v3)
                                                                            _ = v53
                                                                        }
                                                                    }
                                                                    switch v5 {
                                                                    case "":
                                                                        v5 = v5
                                                                    }
                                                                } else {
                                                                    var v54 int8 = (^(v9))
                                                                    _ = v54
                                                                }
                                                            case 42:
                                                                var v55 int32 = v4
                                                                _ = v55
                                                            }
                                                            v56 := []int{fi31}
                                                            _ = v56
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                v1 = (^(v9))
                                            default:
                                                var v57 uint32 = v10
                                                _ = v57
                                            }
                                            if true {
                                                if (v3 != 0.1) {
                                                    var v58 int32 = v4
                                                    if true {
                                                        switch v5 {
                                                        case "a":
                                                            v9 -= 27
                                                        case "hello":
                                                            var v59 string = ((((("foo bar" + "foo bar") + "a") + v5) + "\n\t") + "foo bar")
                                                            _ = v59
                                                        }
                                                        var v60 float32 = ((((v3 + v3) - v3) / 0.1) / v3)
                                                        _ = v60
                                                    }
                                                    _ = v58
                                                } else {
                                                    v61 := GenS9{F0: -631, F1: true, F2: 42}
                                                    if (false && (((0.1 * -2.25) * (3.14159 / 3.14159)) <= 3.14159)) {
                                                        if true {
                                                            v62 := []int{fi34, (fi31 - 278), (fi31 / ((int(v42)) | 1)), 0}
                                                            _ = v62
                                                        }
                                                    }
                                                    _ = v61
                                                }
                                                v2[(v10 + 538)] = false
                                            }
                                            switch (v5 + "\n\t") {
                                            case "\n\t":
                                                switch "\n\t" {
                                                case "hello":
                                                    v63 := []int{fi31, 42}
                                                    _ = v63
                                                case "foo bar":
                                                    delete(v2, (^(v10)))
                                                }
                                            case "a":
                                                var v64 float32 = 1e10
                                                _ = v64
                                            default:
                                                v3 += v3
                                            }
                                        }
                                        v65 := v2[(v10 + v10)]
                                        v66 := GenS4{F0: -114, F1: int16(v10)}
                                        _ = v65
                                        _ = v66
                                    } else {
                                        var v67 string = ((v5 + (((v5 + (v5 + (((v5 + ("" + "")) + "a") + ("" + v5)))) + ("\n\t" + (v5 + ""))) + ("a" + v5))) + v8)
                                        switch (("foo bar" + (v67 + "\n\t")) + v8) {
                                        case "":
                                            v68 := GenS10{F0: -986, F1: 44, F2: v5}
                                            _ = v68
                                        case "hello":
                                            var v69 uint32 = (^(v10))
                                            _ = v69
                                        case "a":
                                            var v70 uint64 = 1
                                            _ = v70
                                        default:
                                            v42 = v4
                                        }
                                        _ = v67
                                    }
                                default:
                                    var v71 int = 32767
                                    _ = v71
                                }
                                _ = v42
                            }
                            var v72 int8 = v9
                            v2[(v10 & (v10 / (((^(v10))) | 1)))] = true
                            v73 := []int{(fi31 * int(v4))}
                            _ = fi34
                            _ = v72
                            _ = v73
                        }
                        if (18446744073709551615 != 1) {
                            fmt.Println(v5)
                        } else {
                            for fi74 := 0; fi74 < 4; fi74++ {
                                v4 += v4
                                var v75 uint16 = 1
                                switch (-(fi31)) {
                                case 128:
                                    if false {
                                        _ = len(v2)
                                        switch "" {
                                        case "a":
                                            v5 = "hello"
                                        case "":
                                            v10 -= (v10 * v10)
                                        }
                                    }
                                case 4294967295:
                                    var v76 uint64 = 18446744073709551615
                                    _ = v76
                                case -962:
                                    fmt.Println(v75, fi31, fi74)
                                }
                                fmt.Println(v75, v3)
                                v77 := GenS6{F0: -68, F1: (v9 >> ((415) & 63)), F2: (2 >= 127)}
                                _ = fi74
                                _ = v75
                                _ = v77
                            }
                        }
                    }
                case "":
                    if (!(v9 > v1)) {
                        v78 := []int{(^(fi31)), fi31, 717}
                        v2[65535] = (true && false)
                        _ = v78
                    } else {
                        for fi79 := 0; fi79 < 2; fi79++ {
                            var v80 uint64 = 295
                            v3 = v3
                            v80 -= (^(v80))
                            if (!(!(true && true))) {
                                v81 := GenS10{F0: int16(v9), F1: 49, F2: v5}
                                v80 -= 18446744073709551615
                                v8 = v8
                                v80 = v80
                                _ = v81
                            }
                            _ = fi79
                            _ = v80
                        }
                    }
                default:
                    v82 := GenS11{F0: fi31, F1: -32768, F2: 1}
                    _ = v82
                }
            }
            var v83 string = (v5 + (("\n\t" + (v5 + v8)) + v5))
            _ = v83
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
