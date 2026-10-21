package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 int16 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint; F1 int64; F2 uint32 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 int; F1 int }
type GenS9 struct { F0 uint32; F1 int64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 string; F1 int }
type GenS12 struct { F0 string; F1 int8 }
func genFn0(p0 int8, p1 uint) uint32 {
    v4 := GenS1{F0: 228}
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 string = ""
        v6 = ("hello" + v6)
        _ = fi5
        _ = v6
    }
    v7 := []int{243, 4294967295}
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 168
}
func genFn1(p0 int16, p1 uint16, p2 uint64) int64 {
    v8 := []int{32767}
    p1 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -1
}

func main() {
    var v1 int64 = 42
    v2 := []string{"foo bar", "a", ("a" + "foo bar")}
    var v3 int8 = 100
    var v9 int16 = 1
    switch "a" {
    case "":
        var v10 bool = (false && true)
        _ = v10
    case "\n\t":
        var v11 uint = 461
        _ = v11
    default:
        if (1 == 18446744073709551615) {
            v1 += (v1 & (-(v1)))
            switch -9223372036854775808 {
            case -2147483648:
                fmt.Println(v3)
            }
            v3 += v3
        } else {
            v1 = genFn1((-(v9)), 2, 0)
        }
    }
    v9 -= 0
    var v12 bool = false
    v13 := GenS2{F0: 0, F1: 2147483647, F2: (-(v9))}
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 int32 = -907
        var v16 string = ("" + "hello")
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS2{F0: 0, F1: 795, F2: 1000}
        var v19 uint16 = 2
        _ = fi17
        _ = v18
        _ = v19
    }
    v3 = (v3 * (v3 >> ((4294967295) & 63)))
    if v12 {
        var v20 int = -360
        _ = v20
    } else {
        v12 = v12
        for fi21 := 0; fi21 < 5; fi21++ {
            for fi22 := 0; fi22 < 10; fi22++ {
                if v12 {
                    if (((1e10 + 3.14159) < (1.5 / ((0.0 - (-2.25 * 0.0)) - 1e10))) && true) {
                        v23 := GenS1{F0: 69}
                        v9 += v9
                        _ = v23
                    } else {
                        var v24 int = (-(fi22))
                        v25 := GenS2{F0: 18446744073709551615, F1: -1, F2: 388}
                        var v26 int8 = v3
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS3{F0: (-(fi21)), F1: 18446744073709551615}
                    _ = v27
                }
                if (875 > 4294967295) {
                    v3 = (^(v3))
                    v28 := GenS4{F0: 42, F1: 415, F2: 550}
                    v1 -= v1
                    _ = v28
                }
                if v12 {
                    switch "" {
                    case "foo bar":
                        fmt.Println(v1)
                    }
                    fmt.Println(v12, fi21, fi21)
                } else {
                    v1 = v1
                }
                _ = fi22
            }
            if (!(!(v3 <= 41))) {
                if v12 {
                    var v29 uint16 = 953
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    if true {
                        for fi30 := 0; fi30 < 17; fi30++ {
                            var v31 int64 = (v1 * (^(v1)))
                            _ = fi30
                            _ = v31
                        }
                    }
                    _ = v29
                } else {
                    if v12 {
                        switch fi21 {
                        case 9223372036854775807:
                            if v12 {
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                                if (!("\n\t" < "")) {
                                    switch 128 {
                                    case 127:
                                        switch -2147483648 {
                                        case -2147483648:
                                            fmt.Println(v3, fi21, v9)
                                        case 42:
                                            for fi32 := 0; fi32 < 11; fi32++ {
                                                var v33 int = fi32
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v3 = v3
                                                case "a":
                                                    v34 := GenS3{F0: v33, F1: 18446744073709551615}
                                                    _ = v34
                                                }
                                                _ = fi32
                                                _ = v33
                                            }
                                        case 128:
                                            for fi35 := 0; fi35 < 17; fi35++ {
                                                v36 := GenS5{F0: 4294967295, F1: (v1 * v1), F2: 2147483648}
                                                _ = fi35
                                                _ = v36
                                            }
                                        default:
                                            _ = len(v2)
                                        }
                                    case 975:
                                        v3 = v3
                                    default:
                                        v3 = (-(v3))
                                    }
                                    for fi37 := 0; fi37 < 20; fi37++ {
                                        v38 := GenS6{F0: (v1 / (((v1 ^ 1)) | 1)), F1: v9}
                                        fmt.Println(v12, v12)
                                        v39 := GenS7{F0: v3, F1: 869}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v40 uint32 = 653
                                        _ = fi37
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                }
                                switch (fi21 % (((fi21 ^ 1)) | 1)) {
                                case 127:
                                    fmt.Println(v3)
                                case 2147483647:
                                    var v41 int16 = v9
                                    _ = v41
                                case -2147483648:
                                    v3 -= v3
                                default:
                                    v12 = (677 < 255)
                                }
                            }
                        case 42:
                            switch "foo bar" {
                            case "hello":
                                var v42 int16 = 0
                                _ = v42
                            case "\n\t":
                                v43 := GenS8{F0: fi21, F1: int(v3)}
                                _ = v43
                            case "":
                                if v12 {
                                    if v12 {
                                        switch (fi21 - (^(fi21))) {
                                        case 1:
                                            v9 = (v9 + v9)
                                        case 32768:
                                            switch fi21 {
                                            case -1:
                                                if v12 {
                                                    fmt.Println(fi21, v1)
                                                }
                                            default:
                                                for fi44 := 0; fi44 < 6; fi44++ {
                                                    if v12 {
                                                        v3 = (v3 ^ v3)
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v45 := GenS5{F0: 18446744073709551615, F1: v1, F2: genFn0(v3, 2)}
                                                            if (602 > 0) {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v46 := GenS9{F0: 4294967295, F1: v1}
                                                                v1 += 1
                                                                _ = v46
                                                            } else {
                                                                var v47 uint8 = 42
                                                                var v48 int16 = v9
                                                                v47 += (v47 | (v47 * 128))
                                                                v49 := GenS10{F0: ("hello" + "a")}
                                                                _ = v47
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                            var v50 string = ""
                                                            _ = v45
                                                            _ = v50
                                                        }
                                                    } else {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = "a"
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = fi44
                                                }
                                            }
                                        case 127:
                                            switch fi21 {
                                            case 1:
                                                var v51 uint64 = 0
                                                _ = v51
                                            }
                                        default:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        v2 = append(v2, "")
                                    } else {
                                        v52 := GenS11{F0: "a", F1: 127}
                                        _ = v52
                                    }
                                    if (true && (!v12)) {
                                        v9 = -446
                                    } else {
                                        v9 = -1
                                        switch "a" {
                                        case "\n\t":
                                            switch (fi21 >> ((301) & 63)) {
                                            case 0:
                                                v1 = v1
                                            case 1:
                                                v3 += (v3 & v3)
                                            }
                                        default:
                                            _ = len(v2)
                                        }
                                    }
                                }
                            default:
                                if false {
                                    v3 -= v3
                                    if ((200 > genFn0(v3, 0)) && true) {
                                        v9 += 1000
                                    }
                                }
                            }
                        case 232:
                            v2 = append(v2, "a")
                        default:
                            _ = len(v2)
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v9 -= (v9 + (v9 >> ((142) & 63)))
                }
            }
            _ = fi21
        }
        var v53 int = -579
        _ = v53
    }
    v54 := GenS8{F0: 32767, F1: 2147483647}
    v12 = false
    v55 := GenS12{F0: (("foo bar" + ("hello" + "foo bar")) + "foo bar"), F1: v3}
    if ((-(v9)) > -520) {
        v2 = append(v2, ("a" + ("foo bar" + "\n\t")))
        fmt.Println(v3)
    }
    v1 = (v1 % ((genFn1(1000, 0, 652)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v9
}
