package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint32; F2 uint }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 int; F1 bool; F2 int32 }
type GenS9 struct { F0 uint; F1 uint16; F2 int8 }
type GenS10 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (!(false || (true && ((862 < 65535) && true))))
    var v3 int64 = 42
    var v4 uint64 = 18446744073709551615
    var v5 int8 = -88
    var v6 float64 = (-2.25 - 0.0)
    var v7 int = 32768
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float32 = 3.14159
        if v2 {
            for fi10 := 0; fi10 < 10; fi10++ {
                if v2 {
                    var v11 int64 = (v3 | -1)
                    _ = v11
                } else {
                    switch fi8 {
                    case 32767:
                        v6 = v6
                    case 42:
                        if (-2147483648 != 2147483647) {
                            v5 = 42
                            v4 -= 4294967295
                            if false {
                                v12 := GenS1{F0: (v3 & (v3 | (v3 ^ 657)))}
                                var v13 float32 = v9
                                var v14 int16 = -32768
                                _ = v12
                                _ = v13
                                _ = v14
                            } else {
                                var v15 float64 = (v6 - float64(v9))
                                _ = v15
                            }
                        } else {
                            v7 = -2147483648
                        }
                    case -430:
                        v5 = v5
                    default:
                        switch "" {
                        case "foo bar":
                            if (false && v2) {
                                v16 := GenS1{F0: v3}
                                if v2 {
                                    _ = len(v1)
                                    v1[969] = ("hello" + "\n\t")
                                    var v17 uint8 = 2
                                    _ = v17
                                } else {
                                    var v18 float64 = float64((1.5 + v9))
                                    if false {
                                        v4 += (v4 * (v4 - v4))
                                        var v19 uint = 1
                                        _ = v19
                                    } else {
                                        v2 = (18446744073709551615 != (v4 * (^(v4))))
                                    }
                                    _ = v18
                                }
                                if v2 {
                                    v20 := GenS1{F0: (v3 * v3)}
                                    v4 -= 76
                                    _ = v20
                                }
                                v21 := GenS1{F0: -738}
                                v22 := GenS1{F0: (v3 & 0)}
                                _ = v16
                                _ = v21
                                _ = v22
                            } else {
                                v4 += 4294967295
                                if ((false || (1 >= 2)) || false) {
                                    var v23 int = (^(v7))
                                    _ = v23
                                }
                            }
                        }
                    }
                }
                v4 = v4
                v2 = (v2 || v2)
                switch (fi10 | v7) {
                case 2147483647:
                    if v2 {
                        v4 = (v4 % ((1) | 1))
                        if (!(((!false) && v2) && ((v3 | v3) == -1))) {
                            switch "foo bar" {
                            case "foo bar":
                                if (((v5 - (-(v5))) <= -5) || v2) {
                                    if (int64(v4) == (v3 | 42)) {
                                        v3 -= (v3 | -702)
                                        var v24 int32 = 2147483647
                                        _ = v24
                                    } else {
                                        v3 -= v3
                                        var v25 int8 = (v5 / (((v5 ^ (-(v5)))) | 1))
                                        _ = v25
                                    }
                                    v26 := GenS2{F0: 2147483647}
                                    var v27 float64 = (v6 / v6)
                                    _ = v26
                                    _ = v27
                                }
                            case "":
                                if true {
                                    if v2 {
                                        v28 := GenS1{F0: (v3 ^ (v3 ^ (^(v3))))}
                                        _ = v28
                                    } else {
                                        v9 -= (v9 * 3.14159)
                                        switch "a" {
                                        case "\n\t":
                                            switch fi8 {
                                            case -79:
                                                var v29 uint64 = 1
                                                _ = v29
                                            default:
                                                v30 := GenS3{F0: "a", F1: 2147483648, F2: 18446744073709551615}
                                                _ = v30
                                            }
                                        case "a":
                                            var v31 uint64 = v4
                                            _ = v31
                                        case "hello":
                                            v32 := []int{(^(fi10))}
                                            _ = v32
                                        }
                                        v33 := GenS4{F0: 42, F1: 42}
                                        _ = v33
                                    }
                                }
                            default:
                                v34 := GenS4{F0: 537, F1: 32767}
                                _ = v34
                            }
                            var v35 int32 = 42
                            _ = len(v1)
                            v36 := []int{fi8, (v7 / ((932) | 1))}
                            _ = v35
                            _ = v36
                        } else {
                            v4 += (v4 << ((18446744073709551615) & 63))
                        }
                    }
                case -1:
                    var v37 uint64 = (v4 - (v4 >> ((65535) & 63)))
                    _ = v37
                case 32768:
                    v2 = v2
                    switch "foo bar" {
                    case "foo bar":
                        v38 := v1[998]
                        _ = v38
                    default:
                        var v39 bool = (4294967295 >= 255)
                        _ = v39
                    }
                    v6 -= 1.5
                default:
                    var v40 uint64 = 152
                    _ = v40
                }
                _ = fi10
            }
        } else {
            for fi41 := 0; fi41 < 20; fi41++ {
                if v2 {
                    if (1 >= 128) {
                        if false {
                            v42 := GenS5{F0: "\n\t"}
                            fmt.Println(fi8, v2, fi41)
                            _ = v42
                        } else {
                            var v43 int32 = 0
                            delete(v1, 990)
                            _ = v43
                        }
                    }
                } else {
                    v2 = (909 > 255)
                    if (!false) {
                        switch fi41 {
                        case 32767:
                            var v44 uint8 = 79
                            _ = v44
                        case -9223372036854775808:
                            var v45 uint = 2
                            _ = v45
                        }
                        var v46 int32 = -1
                        if v2 {
                            fmt.Println(fi41)
                            _ = len(v1)
                            if ((-2.25 / ((v6 + 0.0) - v6)) <= 3.14159) {
                                fmt.Println(v9, v2)
                                if ((v4 >= (v4 ^ v4)) && (!(!(!(false && v2))))) {
                                    v2 = false
                                }
                                if false {
                                    if v2 {
                                        if (((v6 - (0.1 * v6)) / v6) == v6) {
                                            fmt.Println(v3, v9, fi41)
                                            v47 := []int{}
                                            _ = v47
                                        } else {
                                            switch ("foo bar" + "a") {
                                            case "\n\t":
                                                var v48 int32 = v46
                                                _ = v48
                                            }
                                            var v49 uint = 94
                                            v7 -= 249
                                            _ = v49
                                        }
                                    } else {
                                        if v2 {
                                            v1[2147483648] = ("hello" + "a")
                                            if true {
                                                switch "\n\t" {
                                                case "":
                                                    switch ("hello" + ("" + "a")) {
                                                    case "a":
                                                        v50 := GenS6{F0: (v3 ^ (v3 << ((202) & 63))), F1: 32767, F2: -9223372036854775808}
                                                        _ = v50
                                                    case "hello":
                                                        switch "a" {
                                                        case "a":
                                                            v46 = 2147483647
                                                        }
                                                    }
                                                }
                                                v51 := GenS7{F0: v4, F1: 1}
                                                var v52 int = (^(v7))
                                                _ = v51
                                                _ = v52
                                            }
                                            v53 := GenS7{F0: (v4 | 176), F1: -78}
                                            _ = v53
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            v4 -= 2
                                        case "foo bar":
                                            var v54 int16 = 32767
                                            _ = v54
                                        }
                                        switch v7 {
                                        case 976:
                                            v55 := v1[1]
                                            _ = v55
                                        case -940:
                                            v56 := GenS8{F0: fi8, F1: v2, F2: (v46 ^ v46)}
                                            _ = v56
                                        case 4294967295:
                                            v57 := GenS9{F0: 2, F1: 32767, F2: (v5 << ((1) & 63))}
                                            _ = v57
                                        }
                                    }
                                    v58 := GenS3{F0: "\n\t", F1: 0, F2: 2}
                                    fmt.Println(v46, fi41)
                                    _ = v58
                                }
                            }
                        } else {
                            var v59 uint16 = 65535
                            _ = v59
                        }
                        v60 := []int{}
                        _ = v46
                        _ = v60
                    } else {
                        v3 = v3
                        fmt.Println(v6, fi8, v5)
                    }
                }
                if true {
                    v61 := GenS3{F0: "a", F1: 2147483647, F2: 255}
                    var v62 string = ""
                    _ = v61
                    _ = v62
                } else {
                    v63 := []int{-897, (^(v7)), (^(fi41)), 32767}
                    switch fi8 {
                    case -7:
                        v9 -= (v9 * ((v9 / (v9 / ((-2.25 / float32(v6)) - (v9 + (v9 * 0.1))))) * 0.1))
                    }
                    v9 = v9
                    _ = v63
                }
                var v64 int8 = v5
                _ = fi41
                _ = v64
            }
            v9 += ((v9 / v9) * ((((1e10 - 1.5) * v9) * 1.5) + (v9 + v9)))
            v2 = v2
        }
        _ = fi8
        _ = v9
    }
    v65 := GenS10{F0: 1, F1: 2}
    var v66 int8 = v5
    v67 := []int{}
    fmt.Println(v6, v66, v3)
    var v68 bool = true
    v69 := v1[2147483648]
    v2 = (false && (v6 > ((3.14159 + v6) * (float64(float32(v6)) / float64((1.5 + 3.14159))))))
    var v70 string = v69
    switch (v70 + (v69 + v69)) {
    case "\n\t":
        v71 := GenS4{F0: 224, F1: -1}
        _ = v71
    case "":
        var v72 uint = 42
        _ = v72
    case "a":
        v2 = ("\n\t" < v69)
    }
    var v73 float64 = ((-2.25 / (v6 * (((v6 - ((-2.25 * (float64(((1.5 - 0.0) / ((1e10 * float32(0.1)) - float32((0.0 + v6))))) + v6)) + v6)) * (v6 * ((1.5 - (v6 + ((v6 * (v6 * 1e10)) / (v6 / ((v6 / v6) / 0.1))))) / 0.1))) + -2.25))) * 1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v7
    _ = v70
    _ = v73
}
