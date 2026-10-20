package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int32 }
type GenS2 struct { F0 int8; F1 bool; F2 uint16 }
type GenS3 struct { F0 int; F1 uint; F2 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 uint8; F2 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8; F1 int64; F2 string }
type GenS8 struct { F0 int; F1 int8; F2 int8 }
func genFn0(p0 int32, p1 int, p2 int8) string {
    p0 += p0
    v7 := GenS1{F0: 257, F1: "", F2: p0}
    v8 := GenS2{F0: (p2 ^ -21), F1: false, F2: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ""
}
func genFn1(p0 uint8, p1 bool, p2 int8) bool {
    switch (("hello" + "\n\t") + (((("hello" + ("hello" + "\n\t")) + "foo bar") + "") + "\n\t")) {
    case "\n\t":
        p0 = p0
    default:
        if (!(p1 && (p2 > p2))) {
            v9 := []int{768, -2147483648, -126}
            _ = v9
        } else {
            switch 675 {
            case 790:
                v10 := GenS3{F0: 4294967295, F1: 42, F2: 65535}
                _ = v10
            case 128:
                p1 = ((-128 != p2) || p1)
            case -2147483648:
                p0 = (p0 - (^(p0)))
            default:
                switch 655 {
                case -1:
                    p1 = p1
                case 32767:
                    var v11 uint32 = 979
                    _ = v11
                default:
                    var v12 int8 = (p2 * (p2 >> ((18446744073709551615) & 63)))
                    _ = v12
                }
            }
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := []int{4294967295, fi13}
                switch -736 {
                case 1:
                    if (-1 >= 1000) {
                        v14 = append(v14, fi13)
                    } else {
                        fmt.Println(fi13, fi13, p0)
                        if (!(1 > 1)) {
                            p2 = p2
                            var v15 float64 = 1e10
                            _ = v15
                        }
                    }
                case 128:
                    p2 = (-(p2))
                case 42:
                    v16 := GenS3{F0: (fi13 >> ((255) & 63)), F1: 18446744073709551615, F2: 2147483647}
                    _ = v16
                }
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 float64 = ((1e10 * 1.5) / 3.14159)
                    p1 = p1
                    var v19 int64 = 9223372036854775807
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi13
                _ = v14
            }
            var v20 uint8 = (^(p0))
            _ = v20
        }
    }
    v21 := []int{-65, 2147483647, 32768, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return true
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int16 = -32768
    var v3 string = ("foo bar" + ("hello" + ""))
    v4 := []int{(v1 + 127), v1, (-(v1))}
    var v5 float64 = 0.0
    var v6 string = ""
    v22 := GenS1{F0: 4294967295, F1: ("\n\t" + v3), F2: int32(v1)}
    if false {
        var v23 uint16 = 1
        v23 -= v23
        for fi24 := 0; fi24 < 12; fi24++ {
            fmt.Println(v3, v6)
            if false {
                switch v6 {
                case "":
                    switch (^(v1)) {
                    case 127:
                        v23 = (^(v23))
                    case 767:
                        if ((v2 * (v2 - v2)) >= (v2 / ((v2) | 1))) {
                            var v25 int32 = 42
                            if true {
                                switch ("" + "hello") {
                                case "\n\t":
                                    v6 = (v6 + v6)
                                case "a":
                                    switch (fi24 / (((fi24 - (fi24 % ((9223372036854775807) | 1)))) | 1)) {
                                    case 2147483647:
                                        var v26 int = -2147483648
                                        _ = v26
                                    default:
                                        for fi27 := 0; fi27 < 8; fi27++ {
                                            v28 := GenS4{F0: (fi24 >> ((724) & 63))}
                                            var v29 uint8 = 2
                                            var v30 uint = 357
                                            _ = fi27
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        }
                                    }
                                default:
                                    var v31 bool = false
                                    _ = v31
                                }
                            } else {
                                for fi32 := 0; fi32 < 20; fi32++ {
                                    var v33 int = (fi24 + (fi32 ^ v1))
                                    v4 = append(v4, v33)
                                    _ = fi32
                                    _ = v33
                                }
                            }
                            _ = v25
                        } else {
                            var v34 int16 = (v2 ^ v2)
                            _ = len(v4)
                            _ = v34
                        }
                    case -2147483648:
                        var v35 bool = (false || false)
                        _ = v35
                    default:
                        var v36 uint16 = (^(v23))
                        _ = v36
                    }
                }
                v4 = append(v4, (-(fi24)))
                var v37 float64 = float64(0.1)
                _ = v37
            }
            _ = fi24
        }
        _ = v23
    }
    if (0 > 685) {
        if len(v4) > 0 {
            v4[0] = (v1 ^ v1)
        }
        v38 := GenS5{F0: 1, F1: 255, F2: (v1 & 42)}
        v39 := GenS6{F0: 128}
        _ = v38
        _ = v39
    } else {
        switch (v6 + v6) {
        case "":
            v3 = ("hello" + (v3 + "a"))
        case "a":
            v1 = (v1 - v1)
        case "foo bar":
            var v40 uint8 = 2
            _ = v40
        default:
            if (255 >= 4294967295) {
                if len(v4) > 0 {
                    v4[0] = (v1 >> ((2) & 63))
                }
            }
        }
    }
    var v41 int8 = int8(v2)
    v5 -= ((-2.25 / (v5 * (v5 * ((float64((1.5 * (1e10 + (0.0 - -2.25)))) / v5) / ((v5 + float64(0.1)) * 1e10))))) + (1e10 / (v5 - ((1.5 * ((v5 + float64(float32(v5))) / -2.25)) / 1e10))))
    var v42 float32 = 3.14159
    v5 -= (3.14159 / float64(1.5))
    for fi43 := 0; fi43 < 17; fi43++ {
        v3 = "\n\t"
        v2 = v2
        _ = fi43
    }
    for fi44 := 0; fi44 < 15; fi44++ {
        switch int(v41) {
        case 1:
            var v45 int32 = 42
            _ = v45
        case 128:
            if (!(-1 >= -1)) {
                if len(v4) > 0 {
                    v4[0] = -688
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi46 := 0; fi46 < 13; fi46++ {
                    v47 := GenS7{F0: v41, F1: 916, F2: ((v6 + ("foo bar" + "foo bar")) + (v6 + (v3 + "\n\t")))}
                    v41 -= (-(v41))
                    if len(v4) > 0 {
                        v4[0] = (-(fi46))
                    }
                    _ = fi46
                    _ = v47
                }
            }
        case 432:
            switch v6 {
            case "a":
                v6 = (v6 + v3)
            default:
                if true {
                    v48 := GenS1{F0: 42, F1: v3, F2: 2147483647}
                    v42 -= float32(0.1)
                    _ = v48
                }
            }
        default:
            v49 := GenS8{F0: fi44, F1: (^(v41)), F2: (-(v41))}
            _ = v49
        }
        v50 := GenS3{F0: (v1 ^ 32768), F1: 711, F2: 2}
        switch v1 {
        case 128:
            for fi51 := 0; fi51 < 13; fi51++ {
                var v52 float32 = (v42 + (v42 * (1e10 / 1.0)))
                _ = fi51
                _ = v52
            }
        case -347:
            for fi53 := 0; fi53 < 17; fi53++ {
                fmt.Println(v6, v42, fi44)
                if false {
                    v2 += v2
                } else {
                    if len(v4) > 0 {
                        v4[0] = (^(fi53))
                    }
                    _ = len(v4)
                    v42 += v42
                }
                v54 := GenS1{F0: 117, F1: ("hello" + "hello"), F2: 0}
                _ = fi53
                _ = v54
            }
        }
        _ = fi44
        _ = v50
    }
    var v55 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v55
    _ = v6
}
