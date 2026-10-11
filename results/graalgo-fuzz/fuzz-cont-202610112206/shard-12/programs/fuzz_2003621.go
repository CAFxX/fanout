package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int64; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int, p1 uint8, p2 bool) int64 {
    var v7 uint64 = uint64(p1)
    var v8 uint64 = (v7 & (v7 / (((v7 % ((18446744073709551615) | 1))) | 1)))
    fmt.Println(v8, v7, v8)
    v9 := GenS1{F0: (p1 << ((4294967295) & 63)), F1: -2147483648}
    var v10 int16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -953
}
func genFn1(p0 int64, p1 uint16) int64 {
    fmt.Println(p1)
    v11 := []int{}
    p1 += 462
    p0 -= (p0 << ((65535) & 63))
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v11
    return 9223372036854775807
}

func main() {
    var v1 int64 = -796
    var v2 float64 = 0.0
    v3 := []string{"\n\t", "a", ("\n\t" + "a"), "hello"}
    var v4 uint64 = 1
    v5 := make(map[int]int)
    var v6 uint64 = v4
    var v12 uint64 = (v4 >> ((5) & 63))
    var v13 uint32 = 1
    fmt.Println(v13, v4, v1)
    var v14 int = 32767
    var v15 float32 = 0.1
    _ = len(v3)
    var v16 float32 = v15
    v2 -= (float64((v16 / ((v15 + ((1e10 + v15) - v15)) - 1.5))) + v2)
    v17 := GenS1{F0: 1, F1: -257}
    switch "hello" {
    case "hello":
        var v18 uint64 = v6
        _ = v18
    default:
        v3 = append(v3, "a")
    }
    switch -2147483648 {
    case 32768:
        var v19 string = ("" + (("\n\t" + "foo bar") + ("hello" + "foo bar")))
        _ = v19
    case 936:
        v12 += v12
    default:
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 int8 = 100
            if (((v2 - (3.14159 * 1.5)) * float64((-2.25 - ((((1.5 - ((v16 - float32(1e10)) * float32(3.14159))) + (3.14159 - 0.1)) * 0.0) * v15)))) >= v2) {
                for fi22 := 0; fi22 < 1; fi22++ {
                    var v23 uint64 = (v6 | 416)
                    _ = fi22
                    _ = v23
                }
                v24 := GenS1{F0: 76, F1: 151}
                _ = v24
            } else {
                v21 -= (v21 - v21)
                if true {
                    if (v2 >= float64((v15 + 1.5))) {
                        for fi25 := 0; fi25 < 7; fi25++ {
                            v26 := GenS1{F0: 172, F1: 42}
                            _ = fi25
                            _ = v26
                        }
                        var v27 int8 = (v21 >> ((255) & 63))
                        var v28 float32 = v15
                        _ = v27
                        _ = v28
                    }
                    var v29 int64 = genFn1(v1, 32767)
                    fmt.Println(v12)
                    _ = v29
                }
            }
            for fi30 := 0; fi30 < 7; fi30++ {
                v2 = (float64(v15) / v2)
                v12 = (v12 % (((v6 & v6)) | 1))
                var v31 int8 = (v21 | (v21 >> ((18446744073709551615) & 63)))
                _ = fi30
                _ = v31
            }
            _ = fi20
            _ = v21
        }
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        v6 = (v4 | v12)
        var v33 int64 = v1
        _ = fi32
        _ = v33
    }
    v5[v14] = v14
    v34 := GenS2{F0: v1, F1: 32767, F2: 2147483648}
    switch "hello" {
    case "hello":
        if true {
            if false {
                if true {
                    if (128 != 0) {
                        var v35 int8 = 7
                        _ = v35
                    } else {
                        var v36 int = v14
                        _ = v36
                    }
                }
                if true {
                    v3 = append(v3, "hello")
                    fmt.Println(v6)
                } else {
                    v37 := GenS2{F0: -1, F1: -317, F2: v13}
                    v15 = (v16 - 1e10)
                    _ = v37
                }
            }
        } else {
            switch "foo bar" {
            case "":
                _ = len(v3)
            case "\n\t":
                v3 = append(v3, "a")
            case "a":
                var v38 uint16 = 65535
                _ = v38
            }
            v1 -= v1
        }
    case "":
        for fi39 := 0; fi39 < 5; fi39++ {
            fmt.Println(v2, v14)
            switch -2147483648 {
            case -562:
                fmt.Println(v15)
            case 0:
                if false {
                    if false {
                        var v40 int64 = (v1 ^ (v1 & 9223372036854775807))
                        v15 += v16
                        if true {
                            switch fi39 {
                            case 2147483647:
                                if (!true) {
                                    for fi41 := 0; fi41 < 7; fi41++ {
                                        v40 = v1
                                        if true {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        fmt.Println(v1)
                                        _ = fi41
                                    }
                                }
                            case 9223372036854775807:
                                for fi42 := 0; fi42 < 7; fi42++ {
                                    var v43 uint64 = (v4 - v4)
                                    if true {
                                        v44 := v5[(fi39 - 131)]
                                        var v45 uint16 = 270
                                        if ((v45 / ((v45) | 1)) != v45) {
                                            switch "foo bar" {
                                            case "":
                                                if false {
                                                    var v46 int8 = 85
                                                    fmt.Println(fi42)
                                                    fmt.Println(v15, v46)
                                                    _ = v46
                                                }
                                            case "a":
                                                fmt.Println(v44, v12)
                                                v47 := GenS3{F0: 32, F1: 42}
                                                var v48 uint64 = v4
                                                _ = v47
                                                _ = v48
                                            default:
                                                v12 = (v43 + 1)
                                            }
                                            var v49 int32 = 0
                                            fmt.Println(v6)
                                            _ = v49
                                        } else {
                                            if (false || (!("foo bar" != "hello"))) {
                                                _ = len(v3)
                                                var v50 int32 = 42
                                                v51 := GenS4{F0: 0, F1: (v40 ^ v40)}
                                                _ = v50
                                                _ = v51
                                            } else {
                                                v4 -= 1
                                                var v52 uint16 = (v45 + 65535)
                                                _ = v52
                                            }
                                            v2 = ((v2 + v2) / v2)
                                        }
                                        v40 += (^(v40))
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                        _ = v44
                                        _ = v45
                                    } else {
                                        switch fi42 {
                                        case 32768:
                                            fmt.Println(v40, v4)
                                        default:
                                            v15 = (float32(v2) * -2.25)
                                        }
                                        var v53 float64 = 1e10
                                        var v54 int32 = 42
                                        _ = v53
                                        _ = v54
                                    }
                                    _ = fi42
                                    _ = v43
                                }
                            }
                            v55 := GenS5{F0: 0}
                            _ = v55
                        }
                        _ = v40
                    } else {
                        if ((v2 + (v2 * v2)) > v2) {
                            if ("foo bar" <= "hello") {
                                if (true || false) {
                                    for fi56 := 0; fi56 < 2; fi56++ {
                                        var v57 int8 = 42
                                        switch ("" + "a") {
                                        case "foo bar":
                                            v58 := GenS4{F0: 127, F1: 648}
                                            _ = v58
                                        case "\n\t":
                                            var v59 uint32 = 1
                                            v60 := GenS6{F0: 481, F1: "hello"}
                                            _ = v59
                                            _ = v60
                                        case "hello":
                                            v61 := GenS7{F0: (v4 ^ v12)}
                                            _ = v61
                                        }
                                        var v62 uint = 444
                                        _ = fi56
                                        _ = v57
                                        _ = v62
                                    }
                                } else {
                                    v63 := GenS1{F0: 0, F1: 1}
                                    for fi64 := 0; fi64 < 14; fi64++ {
                                        fmt.Println(v15, v6)
                                        _ = fi64
                                    }
                                    _ = v63
                                }
                                _ = len(v5)
                                var v65 uint32 = (v13 & 65535)
                                _ = v65
                            } else {
                                fmt.Println(v13, v2)
                            }
                            v66 := GenS3{F0: 1, F1: 352}
                            switch "" {
                            case "\n\t":
                                switch "\n\t" {
                                case "\n\t":
                                    v67 := GenS7{F0: 2}
                                    _ = v67
                                default:
                                    var v68 float32 = (v15 - (v15 + -2.25))
                                    _ = v68
                                }
                            case "hello":
                                for fi69 := 0; fi69 < 3; fi69++ {
                                    fmt.Println(v12, v14, v13)
                                    if (47 <= 42) {
                                        v16 += (float32(1.5) + 0.1)
                                        var v70 uint64 = (v12 - v12)
                                        var v71 int8 = 42
                                        _ = v70
                                        _ = v71
                                    } else {
                                        var v72 uint32 = 345
                                        _ = v72
                                    }
                                    _ = fi69
                                }
                            case "":
                                var v73 bool = (56 < 0)
                                _ = v73
                            }
                            _ = v66
                        }
                        v74 := GenS1{F0: 32, F1: -1}
                        _ = v74
                    }
                    if true {
                        fmt.Println(fi39)
                    }
                }
            default:
                var v75 uint64 = (^(v4))
                _ = v75
            }
            _ = fi39
        }
    default:
        for fi76 := 0; fi76 < 9; fi76++ {
            v14 += (v14 - fi76)
            if true {
                var v77 int64 = (v1 & 1)
                var v78 uint8 = 128
                switch (v14 - 150) {
                case 32768:
                    if (v6 >= v6) {
                        switch "" {
                        case "":
                            _ = len(v3)
                        }
                        switch (v14 % ((v14) | 1)) {
                        case -9223372036854775808:
                            var v79 int32 = -629
                            _ = v79
                        }
                    }
                case -9223372036854775808:
                    v80 := GenS1{F0: (v78 * v78), F1: -771}
                    _ = v80
                case 127:
                    v13 = (v13 % ((v13) | 1))
                default:
                    var v81 uint16 = 255
                    _ = v81
                }
                _ = v77
                _ = v78
            }
            _ = fi76
        }
    }
    if false {
        fmt.Println(v4, v15)
    } else {
        var v82 int64 = (^(v1))
        fmt.Println(v4, v13, v13)
        _ = v82
    }
    fmt.Println(v1, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
