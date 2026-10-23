package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 bool; F1 int; F2 uint16 }
type GenS7 struct { F0 string; F1 uint; F2 string }
type GenS8 struct { F0 uint64; F1 uint8 }
type GenS9 struct { F0 uint; F1 uint32; F2 int }
type GenS10 struct { F0 uint64 }
func genFn0(p0 uint16) uint16 {
    var v4 uint64 = 662
    var v5 int64 = -9223372036854775808
    _ = p0
    _ = v4
    _ = v5
    return p0
}
func genFn1(p0 int8, p1 int) int64 {
    p0 += -18
    var v6 bool = (false && false)
    v7 := []int{(p1 - p1), p1, p1}
    if len(v7) > 0 {
        v7[0] = (p1 + (-(p1)))
    }
    var v8 float32 = 3.14159
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 631
}

func main() {
    var v1 int = -825
    var v2 uint16 = 670
    var v3 uint16 = v2
    v9 := []int{-2147483648, (v1 + v1), (v1 >> ((65535) & 63))}
    v10 := GenS1{F0: 4294967295, F1: "hello"}
    fmt.Println(v1)
    fmt.Println(v3, v2)
    if true {
        v9 = append(v9, v1)
        if (0 > 128) {
            v2 -= (v2 & 42)
            var v11 int = int(v3)
            _ = v11
        } else {
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 float32 = 3.14159
                v3 += (v2 + (v2 / ((v3) | 1)))
                _ = len(v9)
                _ = fi12
                _ = v13
            }
            fmt.Println(v3)
        }
        var v14 int8 = 0
        _ = v14
    }
    v15 := GenS2{F0: -9223372036854775808, F1: 4294967295}
    for fi16 := 0; fi16 < 16; fi16++ {
        v2 += 2
        _ = fi16
    }
    if (-2.25 <= 0.0) {
        var v17 string = "hello"
        var v18 int16 = 32767
        _ = v17
        _ = v18
    }
    if true {
        var v19 int32 = 952
        v2 = genFn0((v2 << ((uint(v2)) & 63)))
        v3 = (v3 / ((v2) | 1))
        _ = v19
    }
    switch v1 {
    case 2147483647:
        var v20 float32 = float32((0.1 / float64(0.1)))
        _ = v20
    case 4294967295:
        if ((v1 | -787) != 733) {
            var v21 uint64 = 970
            for fi22 := 0; fi22 < 6; fi22++ {
                v2 = v3
                v2 -= (v2 % ((genFn0(293)) | 1))
                v1 = -1
                _ = fi22
            }
            var v23 string = "hello"
            _ = v21
            _ = v23
        } else {
            if (true && (false && true)) {
                v9 = append(v9, v1)
                if false {
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v25 := GenS3{F0: 1}
                        var v26 uint16 = (v3 | v2)
                        switch (fi24 + -902) {
                        case 127:
                            if (false || (-65 >= 1)) {
                                _ = len(v9)
                                v9 = append(v9, (fi24 | (v1 >> ((65535) & 63))))
                                if false {
                                    fmt.Println(fi24, v3, v1)
                                    if true {
                                        fmt.Println(v3, v3, v1)
                                    }
                                    var v27 float64 = 1e10
                                    _ = v27
                                }
                            } else {
                                switch (v1 % (((fi24 - v1)) | 1)) {
                                case 0:
                                    switch v1 {
                                    case 2147483647:
                                        var v28 uint8 = 118
                                        _ = v28
                                    case 637:
                                        if false {
                                            var v29 int16 = 42
                                            _ = v29
                                        }
                                    case 4294967295:
                                        var v30 string = "\n\t"
                                        _ = v30
                                    }
                                case -360:
                                    if (0 <= 610) {
                                        for fi31 := 0; fi31 < 1; fi31++ {
                                            var v32 uint8 = 128
                                            _ = fi31
                                            _ = v32
                                        }
                                    } else {
                                        v33 := GenS4{F0: 18446744073709551615, F1: 59}
                                        fmt.Println(v26)
                                        _ = v33
                                    }
                                case 1:
                                    switch 32767 {
                                    case 127:
                                        switch v1 {
                                        case 42:
                                            v34 := GenS4{F0: 212, F1: 1}
                                            _ = v34
                                        case 2147483647:
                                            fmt.Println(v3)
                                        case 4294967295:
                                            switch "foo bar" {
                                            case "\n\t":
                                                v35 := GenS5{F0: fi24, F1: -310}
                                                _ = v35
                                            }
                                        }
                                    case 128:
                                        switch (("foo bar" + "") + (("" + "\n\t") + "a")) {
                                        case "\n\t":
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                        case "":
                                            if true {
                                                var v36 float32 = float32(0.1)
                                                _ = v36
                                            } else {
                                                for fi37 := 0; fi37 < 17; fi37++ {
                                                    var v38 uint8 = 128
                                                    v9 = append(v9, v1)
                                                    if (true && false) {
                                                        var v39 uint64 = 0
                                                        fmt.Println(fi24, v3, fi37)
                                                        _ = v39
                                                    } else {
                                                        v40 := GenS4{F0: 4294967295, F1: (v38 << ((4294967295) & 63))}
                                                        _ = v40
                                                    }
                                                    _ = fi37
                                                    _ = v38
                                                }
                                                v2 = v3
                                                v1 = v1
                                            }
                                        case "a":
                                            v26 = (v26 & v2)
                                        }
                                    }
                                default:
                                    fmt.Println(v3, v2)
                                }
                                for fi41 := 0; fi41 < 7; fi41++ {
                                    v3 = 706
                                    if false {
                                        if true {
                                            _ = len(v9)
                                            switch "a" {
                                            case "":
                                                if false {
                                                    fmt.Println(v26)
                                                } else {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v42 uint8 = 42
                                                        _ = v42
                                                    }
                                                }
                                                var v43 uint8 = 1
                                                v3 = v26
                                                _ = v43
                                            case "foo bar":
                                                v1 = fi41
                                            case "a":
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v44 uint8 = 2
                                                    var v45 uint32 = 42
                                                    switch "hello" {
                                                    case "a":
                                                        var v46 int16 = -249
                                                        _ = v46
                                                    default:
                                                        v47 := GenS3{F0: -1}
                                                        fmt.Println(v1)
                                                        v26 = (v3 + 255)
                                                        switch ("hello" + (("" + "a") + "\n\t")) {
                                                        case "":
                                                            var v48 uint64 = 2
                                                            var v49 int32 = -1
                                                            switch fi41 {
                                                            case 0:
                                                                v50 := GenS6{F0: (3.14159 == 0.0), F1: v1, F2: genFn0(uint16(v44))}
                                                                _ = v50
                                                            case 128:
                                                                if (false || ((0.0 + 0.0) == (0.0 / 1.5))) {
                                                                    v45 = (^(v45))
                                                                    if false {
                                                                        v3 = (v2 | v2)
                                                                    }
                                                                    var v51 uint32 = (v45 - (v45 >> ((824) & 63)))
                                                                    _ = v51
                                                                }
                                                            default:
                                                                switch 128 {
                                                                case 9223372036854775807:
                                                                    v9 = append(v9, int(v49))
                                                                default:
                                                                    if len(v9) >= 2 {
                                                                        v9 = v9[1:len(v9)]
                                                                    }
                                                                    _ = len(v9)
                                                                }
                                                            }
                                                            _ = v48
                                                            _ = v49
                                                        case "a":
                                                            v2 = (v2 * 498)
                                                        case "\n\t":
                                                            fmt.Println(v26, v1, v44)
                                                        default:
                                                            var v52 float32 = 0.1
                                                            _ = v52
                                                        }
                                                        _ = v47
                                                    }
                                                    _ = v44
                                                    _ = v45
                                                case "foo bar":
                                                    v53 := GenS2{F0: -786, F1: 115}
                                                    _ = v53
                                                }
                                            }
                                            switch ("a" + "\n\t") {
                                            case "\n\t":
                                                v54 := GenS5{F0: fi41, F1: 42}
                                                _ = v54
                                            default:
                                                v26 = 65535
                                            }
                                        }
                                        var v55 bool = false
                                        _ = v55
                                    }
                                    v1 = (fi24 + (fi24 + -977))
                                    _ = fi41
                                }
                                v56 := GenS5{F0: fi24, F1: 2147483647}
                                _ = v56
                            }
                        case 32767:
                            if (fi24 >= 0) {
                                if (false || false) {
                                    v2 = (v2 & v3)
                                } else {
                                    fmt.Println(v2, v2, v2)
                                    if (!((v2 < (v26 << ((2) & 63))) || false)) {
                                        for fi57 := 0; fi57 < 2; fi57++ {
                                            var v58 uint8 = 125
                                            v59 := GenS7{F0: (("\n\t" + "\n\t") + "foo bar"), F1: 42, F2: "hello"}
                                            var v60 uint32 = 0
                                            _ = fi57
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                        }
                                        v26 = (^(v2))
                                    }
                                    if (!true) {
                                        v61 := GenS8{F0: uint64(v2), F1: 168}
                                        var v62 int64 = -9223372036854775808
                                        _ = v61
                                        _ = v62
                                    } else {
                                        var v63 uint8 = 127
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        _ = v63
                                    }
                                }
                                for fi64 := 0; fi64 < 6; fi64++ {
                                    if len(v9) > 0 {
                                        v9[0] = fi24
                                    }
                                    _ = fi64
                                }
                            } else {
                                v26 -= genFn0((v3 / ((65535) | 1)))
                            }
                        case 42:
                            v26 = genFn0((^(v3)))
                        }
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    var v65 uint8 = 1
                    fmt.Println(v2)
                    _ = v65
                }
            } else {
                if ((!false) || false) {
                    var v66 int64 = 9223372036854775807
                    _ = v66
                } else {
                    v67 := GenS9{F0: 384, F1: 0, F2: -2147483648}
                    var v68 float32 = 3.14159
                    v1 = v1
                    _ = v67
                    _ = v68
                }
                for fi69 := 0; fi69 < 19; fi69++ {
                    var v70 uint16 = 65535
                    _ = fi69
                    _ = v70
                }
            }
            v71 := GenS10{F0: 18446744073709551615}
            _ = v71
        }
    default:
        v3 -= 2
    }
    v2 += (v3 ^ genFn0(210))
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
