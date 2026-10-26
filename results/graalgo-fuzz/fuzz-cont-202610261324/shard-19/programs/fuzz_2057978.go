package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint32 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int32; F1 string; F2 int32 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS5 struct { F0 uint; F1 int64; F2 uint16 }
func genFn0() int8 {
    var v7 bool = (65535 < 2147483648)
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint = 65535
        _ = fi8
        _ = v9
    }
    v7 = (0.1 > 1.5)
    v10 := GenS1{F0: -1, F1: ("foo bar" + "\n\t"), F2: 4294967295}
    _ = v10
    _ = v7
    return -1
}
func genFn1(p0 int16) int8 {
    var v11 int64 = 42
    var v12 float64 = (0.0 - ((((1e10 + (((0.0 / 1.5) - -2.25) + 0.0)) - 3.14159) + 1.5) - 0.0))
    var v13 uint8 = 1
    var v14 uint32 = 635
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 34
}

func main() {
    var v1 uint16 = 32767
    var v2 uint16 = 953
    var v3 float64 = 1.5
    var v4 int32 = -531
    var v5 int = int(v2)
    var v6 string = ("\n\t" + (("a" + "\n\t") + ("" + "foo bar")))
    v3 += 1e10
    for fi15 := 0; fi15 < 20; fi15++ {
        v5 -= fi15
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 float32 = 0.1
            v3 = v3
            if true {
                switch fi15 {
                case -347:
                    if false {
                        if true {
                            var v18 bool = false
                            fmt.Println(v1, v4, fi16)
                            _ = v18
                        }
                        fmt.Println(v4)
                        var v19 uint = 65535
                        _ = v19
                    } else {
                        if true {
                            fmt.Println(v4)
                            var v20 uint16 = v2
                            _ = v20
                        } else {
                            v6 = (v6 + v6)
                            v6 = v6
                        }
                    }
                }
                v3 = (1e10 - ((v3 - v3) * float64(v17)))
                var v21 uint16 = (v2 ^ 0)
                _ = v21
            } else {
                switch (fi15 | v5) {
                case 9223372036854775807:
                    v22 := GenS2{F0: 18446744073709551615, F1: 9223372036854775807}
                    switch v6 {
                    case "a":
                        v6 = (("hello" + "foo bar") + "a")
                    default:
                        var v23 uint32 = 42
                        _ = v23
                    }
                    _ = v22
                case 2147483647:
                    v4 -= (v4 % ((v4) | 1))
                }
                var v24 bool = (!((fi16 >> ((1) & 63)) < fi15))
                v1 = (v1 & (v2 % ((v2) | 1)))
                _ = v24
            }
            var v25 string = (v6 + (v6 + ("" + "foo bar")))
            _ = fi16
            _ = v17
            _ = v25
        }
        _ = fi15
    }
    var v26 int8 = 42
    v1 = 2
    for fi27 := 0; fi27 < 3; fi27++ {
        v2 = 65535
        v4 += 42
        var v28 int = -274
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        fmt.Println(fi29)
        _ = fi29
    }
    switch ("foo bar" + "hello") {
    case "a":
        v30 := GenS1{F0: -678, F1: "hello", F2: 0}
        _ = v30
    default:
        var v31 int8 = v26
        _ = v31
    }
    v26 += (-(v26))
    v3 -= v3
    var v32 int = v5
    v2 = 259
    v33 := []int{}
    v34 := GenS3{F0: int32(v5), F1: "\n\t", F2: v4}
    for fi35 := 0; fi35 < 14; fi35++ {
        switch "hello" {
        case "foo bar":
            v36 := GenS4{F0: 135, F1: 267, F2: 0}
            _ = v36
        case "a":
            var v37 int = (v5 % ((4294967295) | 1))
            _ = v37
        case "\n\t":
            if (-1 != 32767) {
                if true {
                    if false {
                        var v38 float32 = 1e10
                        for fi39 := 0; fi39 < 12; fi39++ {
                            v38 += (v38 * v38)
                            var v40 uint32 = 65535
                            if true {
                                var v41 bool = true
                                _ = v41
                            }
                            _ = fi39
                            _ = v40
                        }
                        _ = v38
                    } else {
                        var v42 string = v6
                        _ = v42
                    }
                    v43 := GenS1{F0: 32767, F1: (v6 + v6), F2: 1}
                    for fi44 := 0; fi44 < 5; fi44++ {
                        var v45 int32 = (v4 | (v4 << ((0) & 63)))
                        v46 := GenS3{F0: v45, F1: v6, F2: (v4 ^ 42)}
                        switch (((v6 + v6) + v6) + (v6 + ("hello" + v6))) {
                        case "hello":
                            fmt.Println(fi44, v1)
                        }
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                    _ = v43
                } else {
                    for fi47 := 0; fi47 < 17; fi47++ {
                        var v48 bool = (("foo bar" + ("\n\t" + "foo bar")) > v6)
                        _ = fi47
                        _ = v48
                    }
                }
                fmt.Println(v1, v32, v4)
            }
        default:
            if false {
                v2 -= 396
            } else {
                var v49 int16 = -874
                v33 = append(v33, v5)
                for fi50 := 0; fi50 < 2; fi50++ {
                    var v51 int8 = v26
                    var v52 uint32 = 2147483647
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                _ = v49
            }
        }
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        _ = fi35
    }
    for fi53 := 0; fi53 < 17; fi53++ {
        for fi54 := 0; fi54 < 6; fi54++ {
            v4 -= 2147483647
            v55 := GenS3{F0: (v4 + v4), F1: v6, F2: 2147483647}
            var v56 int = (v32 + (v32 / (((v32 | -826)) | 1)))
            _ = fi54
            _ = v55
            _ = v56
        }
        if false {
            switch v5 {
            case 32767:
                v57 := GenS5{F0: 2, F1: 95, F2: (^(v2))}
                _ = v57
            case 0:
                for fi58 := 0; fi58 < 6; fi58++ {
                    if (v6 <= "\n\t") {
                        var v59 float32 = float32(v3)
                        var v60 int8 = -17
                        _ = v59
                        _ = v60
                    } else {
                        if true {
                            switch (v6 + v6) {
                            case "a":
                                v4 = (v4 & v4)
                            case "hello":
                                v4 = v4
                            default:
                                if len(v33) > 0 {
                                    v33[0] = (fi53 | v32)
                                }
                            }
                        } else {
                            fmt.Println(v5, v1, fi58)
                            if false {
                                v61 := GenS4{F0: 566, F1: 18446744073709551615, F2: 885}
                                _ = v61
                            } else {
                                fmt.Println(v6, fi53, v26)
                                var v62 bool = true
                                v33 = append(v33, v5)
                                _ = v62
                            }
                            v63 := GenS5{F0: 65535, F1: -1, F2: v1}
                            _ = v63
                        }
                        switch 1 {
                        case -1:
                            v6 = "foo bar"
                        case 356:
                            v64 := GenS2{F0: 37, F1: -345}
                            _ = v64
                        }
                    }
                    v65 := GenS3{F0: (v4 >> ((2) & 63)), F1: ("foo bar" + ("a" + v6)), F2: v4}
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                    _ = fi58
                    _ = v65
                }
            }
        }
        _ = fi53
    }
    v2 = (^(v1))
    if true {
        v3 = v3
        var v66 bool = true
        v33 = append(v33, v32)
        _ = v66
    }
    v67 := GenS1{F0: 0, F1: (v6 + "\n\t"), F2: 317}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v67
}
