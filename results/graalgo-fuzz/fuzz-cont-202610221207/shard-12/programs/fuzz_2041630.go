package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint }
type GenS5 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 uint32 }

func main() {
    v1 := []string{"a", ""}
    v2 := []string{"foo bar"}
    var v3 int16 = -67
    var v4 uint64 = 18446744073709551615
    var v5 int32 = int32(v3)
    if true {
        var v6 uint16 = 747
        var v7 int64 = int64(v3)
        _ = v6
        _ = v7
    }
    v3 = 0
    var v8 int32 = 374
    v5 += (v8 % (((v8 ^ 2147483647)) | 1))
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: (fi9 >> ((0) & 63)), F1: 941}
        if len(v2) > 0 {
            v2[0] = "a"
        }
        fmt.Println(v4, v8)
        _ = fi9
        _ = v10
    }
    if true {
        if ((!(42 >= -752)) && (1 == 1)) {
            fmt.Println(v5, v3, v8)
            v8 -= v5
        } else {
            var v11 uint32 = 2147483647
            v2 = append(v2, "\n\t")
            _ = v11
        }
        var v12 uint16 = uint16(v4)
        var v13 uint = 255
        _ = v12
        _ = v13
    } else {
        v5 -= (v5 + (v8 << ((303) & 63)))
        var v14 int32 = (v8 | (v5 ^ v5))
        fmt.Println(v4, v14, v5)
        _ = v14
    }
    fmt.Println(v3)
    var v15 int16 = v3
    var v16 bool = (false || (true || true))
    v17 := GenS2{F0: 128, F1: -613}
    if (!(!v16)) {
        if true {
            if v16 {
                for fi18 := 0; fi18 < 17; fi18++ {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        if true {
                            v5 += (v8 ^ (v8 & v5))
                            if (3.14159 == (float64(0.1) * -2.25)) {
                                _ = len(v2)
                                v2 = append(v2, "\n\t")
                            }
                        } else {
                            if ((v4 | 42) <= (v4 - 1)) {
                                var v20 int32 = (v5 - v5)
                                v8 = -2147483648
                                _ = v20
                            } else {
                                if v16 {
                                    switch ("a" + ("foo bar" + "a")) {
                                    case "":
                                        fmt.Println(v16)
                                    case "\n\t":
                                        v21 := GenS1{F0: (fi19 >> ((2) & 63)), F1: 9223372036854775807}
                                        var v22 int16 = v3
                                        _ = len(v1)
                                        _ = v21
                                        _ = v22
                                    }
                                } else {
                                    var v23 uint32 = 42
                                    fmt.Println(v23, v5)
                                    _ = v23
                                }
                            }
                            v15 -= (v3 / ((v15) | 1))
                        }
                        if v16 {
                            v4 = v4
                        } else {
                            v24 := GenS3{F0: 65535, F1: 115}
                            _ = v24
                        }
                        _ = fi19
                    }
                    switch fi18 {
                    case 1:
                        if len(v1) > 0 {
                            v1[0] = ("foo bar" + ("" + "a"))
                        }
                    case -9223372036854775808:
                        if v16 {
                            if (v5 < v8) {
                                for fi25 := 0; fi25 < 11; fi25++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                    v3 += v3
                                    _ = fi25
                                }
                            }
                        } else {
                            for fi26 := 0; fi26 < 20; fi26++ {
                                v27 := GenS2{F0: 234, F1: (fi18 ^ fi18)}
                                v15 += v3
                                var v28 int32 = (v8 >> ((807) & 63))
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            v3 -= (v15 % ((v15) | 1))
                            for fi29 := 0; fi29 < 4; fi29++ {
                                var v30 int = (-(fi18))
                                _ = fi29
                                _ = v30
                            }
                        }
                    case 32768:
                        if len(v1) > 0 {
                            v1[0] = (("hello" + ("foo bar" + "")) + ("foo bar" + "\n\t"))
                        }
                    }
                    var v31 uint64 = 42
                    _ = fi18
                    _ = v31
                }
                v8 = 1
            } else {
                v32 := GenS4{F0: 0, F1: (v4 << ((18446744073709551615) & 63)), F2: 65535}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v33 int64 = -9223372036854775808
                _ = v32
                _ = v33
            }
            fmt.Println(v4, v8)
        } else {
            if len(v2) > 0 {
                v2[0] = "a"
            }
            var v34 uint = 255
            for fi35 := 0; fi35 < 20; fi35++ {
                v15 -= (v15 ^ (v3 ^ (v3 * (v3 * 42))))
                if v16 {
                    for fi36 := 0; fi36 < 19; fi36++ {
                        v4 -= 393
                        fmt.Println(v16, v15, v15)
                        var v37 int = 395
                        _ = fi36
                        _ = v37
                    }
                    switch (-(fi35)) {
                    case -1:
                        if (!(42 > -110)) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            switch fi35 {
                            case 772:
                                var v38 uint8 = 127
                                _ = v38
                            case -2147483648:
                                if ((true && (0.0 > 0.1)) || v16) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            case 128:
                                var v39 int64 = 0
                                _ = v39
                            }
                            v40 := GenS2{F0: 42, F1: (-(fi35))}
                            if (348 > 635) {
                                if ((3.14159 * float64(float32(-2.25))) == 0.1) {
                                    var v41 uint8 = 141
                                    _ = v41
                                } else {
                                    if (false && (v34 >= 42)) {
                                        v5 += v5
                                    } else {
                                        var v42 uint16 = 42
                                        _ = v42
                                    }
                                    if (v5 < v5) {
                                        if (1 == -1) {
                                            v43 := GenS5{F0: 255, F1: ((fi35 > -9223372036854775808) && v16), F2: 1}
                                            for fi44 := 0; fi44 < 3; fi44++ {
                                                if v16 {
                                                    v5 = (v5 & v8)
                                                    var v45 int16 = v15
                                                    v45 = -1
                                                    _ = v45
                                                } else {
                                                    switch "hello" {
                                                    case "hello":
                                                        if len(v2) > 0 {
                                                            v2[0] = "hello"
                                                        }
                                                    case "a":
                                                        var v46 int8 = -1
                                                        _ = v46
                                                    case "\n\t":
                                                        v47 := GenS3{F0: 1, F1: 9223372036854775807}
                                                        _ = v47
                                                    default:
                                                        if true {
                                                            var v48 uint = (v34 * (v34 - (v34 + (^(v34)))))
                                                            if false {
                                                                v15 -= v15
                                                                v1 = append(v1, ("a" + ("a" + (((("\n\t" + "") + ("" + "a")) + "a") + "\n\t"))))
                                                            } else {
                                                                var v49 string = "hello"
                                                                v50 := GenS2{F0: 255, F1: 128}
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                            _ = v48
                                                        }
                                                    }
                                                    switch "" {
                                                    case "":
                                                        var v51 uint = (v34 % (((v34 % ((755) | 1))) | 1))
                                                        _ = v51
                                                    }
                                                    v5 = (v8 * 987)
                                                }
                                                if true {
                                                    v1 = append(v1, ("" + "a"))
                                                } else {
                                                    v52 := GenS5{F0: 65535, F1: v16, F2: 973}
                                                    _ = v52
                                                }
                                                _ = fi44
                                            }
                                            switch ("foo bar" + "a") {
                                            case "a":
                                                var v53 uint8 = 127
                                                _ = v53
                                            }
                                            _ = v43
                                        } else {
                                            for fi54 := 0; fi54 < 11; fi54++ {
                                                v8 = v5
                                                switch (-(fi54)) {
                                                case 0:
                                                    v15 -= v15
                                                case -9223372036854775808:
                                                    v55 := GenS3{F0: 122, F1: -9223372036854775808}
                                                    _ = v55
                                                }
                                                _ = fi54
                                            }
                                        }
                                        v4 = (v4 << ((v34) & 63))
                                    } else {
                                        v56 := GenS6{F0: (v34 % ((255) | 1)), F1: v34}
                                        var v57 int8 = 36
                                        _ = v56
                                        _ = v57
                                    }
                                }
                                v58 := GenS5{F0: 1, F1: ((^(v3)) < v3), F2: 914}
                                var v59 uint8 = 255
                                _ = v58
                                _ = v59
                            } else {
                                var v60 int64 = 322
                                _ = v60
                            }
                            _ = v40
                        }
                    }
                    var v61 bool = ((1 <= 245) && (v8 > int32(v34)))
                    _ = v61
                }
                switch "" {
                case "":
                    for fi62 := 0; fi62 < 4; fi62++ {
                        if v16 {
                            v3 += (v15 / ((v3) | 1))
                            fmt.Println(v4, v3, v3)
                            fmt.Println(v4)
                        }
                        _ = fi62
                    }
                }
                _ = fi35
            }
            _ = v34
        }
        if len(v1) > 0 {
            v1[0] = ("\n\t" + (("foo bar" + (("hello" + "foo bar") + "")) + "a"))
        }
        v63 := GenS5{F0: 384, F1: (!(!v16)), F2: 42}
        _ = v63
    } else {
        var v64 int64 = -15
        fmt.Println(v3)
        _ = v64
    }
    var v65 uint = 2
    if v16 {
        var v66 uint = v65
        v67 := GenS7{F0: 2}
        v15 += 184
        _ = v66
        _ = v67
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v8
}
