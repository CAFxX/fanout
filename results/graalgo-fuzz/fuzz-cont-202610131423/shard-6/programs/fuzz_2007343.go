package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 string; F1 uint32; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8; F1 int8; F2 int8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8; F1 int32; F2 string }
func genFn0(p0 uint, p1 int) uint {
    fmt.Println(p1, p0)
    p0 = p0
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 string = ("hello" + "foo bar")
    var v2 string = v1
    var v3 uint64 = 283
    var v4 uint64 = 181
    v5 := []int{527, 128, 531, 44}
    var v6 uint16 = 2
    for fi7 := 0; fi7 < 1; fi7++ {
        switch 1 {
        case 42:
            fmt.Println(v1)
        case 1:
            v8 := GenS1{F0: -9223372036854775808, F1: 1000}
            _ = v8
        default:
            v9 := GenS2{F0: v1, F1: uint32(v6), F2: 255}
            _ = v9
        }
        var v10 uint64 = 42
        _ = fi7
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 float64 = ((3.14159 + (0.0 - float64(3.14159))) / 3.14159)
            _ = fi12
            _ = v13
        }
        fmt.Println(v2, v6, v1)
        var v14 uint16 = (v6 >> ((2) & 63))
        _ = fi11
        _ = v14
    }
    if (false && true) {
        var v15 int32 = 42
        fmt.Println(v4, v1)
        if ((v3 - 42) >= (v3 + (v3 * v3))) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            for fi16 := 0; fi16 < 1; fi16++ {
                fmt.Println(fi16, v3)
                if len(v5) > 0 {
                    v5[0] = (fi16 + fi16)
                }
                _ = fi16
            }
            for fi17 := 0; fi17 < 16; fi17++ {
                fmt.Println(v3, v6, v6)
                switch "foo bar" {
                case "":
                    v18 := GenS1{F0: -473, F1: 308}
                    _ = v18
                case "a":
                    v2 = v2
                default:
                    if true {
                        switch fi17 {
                        case -9223372036854775808:
                            var v19 int64 = -763
                            _ = v19
                        case 4294967295:
                            for fi20 := 0; fi20 < 8; fi20++ {
                                v21 := GenS1{F0: 1, F1: -687}
                                if (8 < 128) {
                                    v22 := GenS2{F0: v1, F1: 0, F2: 768}
                                    var v23 string = (v1 + v1)
                                    _ = v22
                                    _ = v23
                                }
                                if ((fi20 & (fi20 << ((65535) & 63))) > 121) {
                                    switch fi17 {
                                    case 32767:
                                        _ = len(v5)
                                    case 4294967295:
                                        v4 += (v4 ^ (v4 - (^(v4))))
                                    default:
                                        var v24 bool = false
                                        _ = v24
                                    }
                                    if true {
                                        switch (fi20 % (((fi17 + 128)) | 1)) {
                                        case 1:
                                            var v25 int8 = -1
                                            _ = v25
                                        case -1:
                                            v26 := GenS3{F0: 0}
                                            _ = v26
                                        case 42:
                                            if ((!false) || true) {
                                                var v27 int16 = 42
                                                _ = v27
                                            }
                                        }
                                        v28 := GenS4{F0: 9223372036854775807}
                                        _ = v28
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = (fi17 | -95)
                                    }
                                    fmt.Println(v1, fi20, v3)
                                } else {
                                    if ((("\n\t" + ((((((v2 + ("\n\t" + "foo bar")) + v2) + ((v1 + v2) + v1)) + "") + "foo bar") + ("a" + "a"))) + v1) < "hello") {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    } else {
                                        v29 := GenS5{F0: v3, F1: ""}
                                        fmt.Println(v15, v1)
                                        if (true || true) {
                                            v1 = "\n\t"
                                            v4 = 1
                                            v6 = v6
                                            fmt.Println(v6, v4, v3)
                                            if false {
                                                var v30 uint = 65535
                                                var v31 float64 = float64(0.0)
                                                _ = v30
                                                _ = v31
                                            } else {
                                                fmt.Println(v6)
                                                fmt.Println(v3)
                                            }
                                            var v32 uint = 1
                                            _ = v32
                                        } else {
                                            v15 = (-(v15))
                                            if len(v5) > 0 {
                                                v5[0] = (fi20 / ((fi17) | 1))
                                            }
                                            var v33 uint16 = v6
                                            _ = v33
                                        }
                                        _ = len(v5)
                                        _ = v29
                                    }
                                    if true {
                                        var v34 int64 = 9223372036854775807
                                        _ = v34
                                    } else {
                                        v35 := GenS6{F0: -85}
                                        v36 := GenS1{F0: -106, F1: 365}
                                        if (false && false) {
                                            var v37 int64 = 9223372036854775807
                                            _ = v37
                                        } else {
                                            if false {
                                                _ = len(v5)
                                                fmt.Println(v4, fi20, fi17)
                                                if len(v5) > 0 {
                                                    v5[0] = (fi20 & -9223372036854775808)
                                                }
                                                var v38 int8 = -1
                                                if false {
                                                    v1 = ((("foo bar" + (((v1 + "") + ("hello" + "\n\t")) + "foo bar")) + "") + v1)
                                                }
                                                _ = v38
                                            }
                                        }
                                        _ = v35
                                        _ = v36
                                    }
                                    fmt.Println(fi20, v2)
                                    v39 := GenS7{F0: 1, F1: 127, F2: -1}
                                    _ = v39
                                }
                                _ = fi20
                                _ = v21
                            }
                        }
                    }
                }
                _ = fi17
            }
        }
        _ = v15
    } else {
        if len(v5) > 0 {
            v5[0] = -2147483648
        }
        switch 843 {
        case 42:
            v40 := GenS6{F0: -1}
            _ = v40
        case 127:
            v2 = ("hello" + "foo bar")
        }
        v5 = append(v5, 4294967295)
    }
    var v41 float64 = 3.14159
    fmt.Println(v1, v3)
    for fi42 := 0; fi42 < 19; fi42++ {
        v43 := GenS7{F0: 107, F1: -99, F2: -41}
        var v44 bool = (!(1 >= 42))
        _ = fi42
        _ = v43
        _ = v44
    }
    v45 := GenS3{F0: 42}
    _ = len(v5)
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    var v46 float32 = 0.1
    v41 = 1.5
    if true {
        for fi47 := 0; fi47 < 6; fi47++ {
            var v48 uint8 = 51
            if true {
                v49 := GenS8{F0: 1}
                _ = v49
            } else {
                v50 := GenS2{F0: v2, F1: 2, F2: 4294967295}
                v46 -= -2.25
                var v51 int64 = 1
                _ = v50
                _ = v51
            }
            for fi52 := 0; fi52 < 5; fi52++ {
                v53 := GenS1{F0: 9223372036854775807, F1: -1}
                v54 := GenS1{F0: int64(v6), F1: -1}
                if false {
                    if ((^(fi47)) > (fi47 >> ((18446744073709551615) & 63))) {
                        if true {
                            var v55 uint64 = 953
                            v56 := GenS1{F0: -9223372036854775808, F1: -414}
                            _ = v55
                            _ = v56
                        } else {
                            var v57 uint64 = 2
                            _ = v57
                        }
                        v46 = (v46 + 1e10)
                        v58 := GenS4{F0: -9223372036854775808}
                        _ = v58
                    } else {
                        if ("foo bar" == v1) {
                            if false {
                                var v59 uint32 = 2147483648
                                _ = v59
                            }
                        } else {
                            v3 = v4
                            var v60 string = (v1 + "\n\t")
                            if false {
                                v6 = (v6 % ((v6) | 1))
                                v48 += (v48 % ((42) | 1))
                                v1 = v1
                            } else {
                                if len(v5) > 0 {
                                    v5[0] = 32767
                                }
                                var v61 uint32 = 380
                                v3 -= 2
                                _ = v61
                            }
                            var v62 int = fi47
                            switch v2 {
                            case "\n\t":
                                var v63 int64 = -782
                                _ = v63
                            default:
                                var v64 int = -9223372036854775808
                                _ = v64
                            }
                            _ = v60
                            _ = v62
                        }
                        v41 = float64(v46)
                    }
                } else {
                    if true {
                        var v65 uint16 = 255
                        v66 := GenS9{F0: v65}
                        _ = v65
                        _ = v66
                    } else {
                        switch ((("hello" + ((v2 + v1) + v2)) + v1) + "hello") {
                        case "foo bar":
                            v48 -= (v48 - v48)
                        case "a":
                            v67 := GenS10{F0: -1, F1: int32(v6), F2: (v1 + "foo bar")}
                            _ = v67
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v4 = v3
                        v5 = append(v5, fi52)
                    }
                    var v68 float32 = 1.5
                    _ = v68
                }
                var v69 float64 = (v41 / (v41 + v41))
                _ = fi52
                _ = v53
                _ = v54
                _ = v69
            }
            _ = fi47
            _ = v48
        }
    }
    if (!(2 == 65535)) {
        v5 = append(v5, 798)
        var v70 string = v1
        v2 = ((v70 + "hello") + "hello")
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
