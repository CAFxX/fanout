package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 string; F1 int; F2 int32 }
type GenS5 struct { F0 int64; F1 uint; F2 uint32 }
func genFn0(p0 string, p1 int8, p2 bool) string {
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 float64 = 1e10
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: int(p1), F1: 955}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p0 + p0)
}
func genFn1(p0 uint8, p1 uint32, p2 int32) bool {
    if false {
        switch "a" {
        case "foo bar":
            var v7 int = 42
            _ = v7
        case "\n\t":
            var v8 int32 = (p2 >> ((1) & 63))
            _ = v8
        }
        if ((0 <= -712) && (false || (!(!true)))) {
            var v9 int16 = 42
            var v10 int32 = p2
            _ = v10
            _ = v9
        }
        p1 += (^(p1))
    } else {
        v11 := GenS1{F0: 2147483647, F1: 2}
        v12 := GenS2{F0: false, F1: p1, F2: 42}
        if (!true) {
            var v13 float32 = 1.5
            v14 := []int{}
            _ = v13
            _ = v14
        }
        _ = v11
        _ = v12
    }
    switch 516 {
    case 32768:
        if false {
            v15 := GenS3{F0: (p0 & 68), F1: 18446744073709551615}
            _ = v15
        } else {
            fmt.Println(p0, p2)
            fmt.Println(p0, p2)
        }
    case -811:
        fmt.Println(p1, p2)
    case -2147483648:
        switch "foo bar" {
        case "":
            p0 += (p0 ^ p0)
        default:
            if false {
                var v16 uint8 = 127
                _ = v16
            } else {
                var v17 int64 = -9223372036854775808
                switch "\n\t" {
                case "a":
                    if (!false) {
                        for fi18 := 0; fi18 < 5; fi18++ {
                            fmt.Println(p2, p1)
                            v19 := []int{127, (fi18 % ((32768) | 1))}
                            _ = fi18
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 15; fi20++ {
                            var v21 int64 = -1
                            var v22 int = (fi20 << ((65535) & 63))
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        var v23 string = "foo bar"
                        _ = v23
                    } else {
                        var v24 uint = 503
                        _ = v24
                    }
                case "\n\t":
                    var v25 string = (("" + "") + ("foo bar" + ("\n\t" + "\n\t")))
                    _ = v25
                case "foo bar":
                    var v26 uint64 = 0
                    _ = v26
                default:
                    p1 -= p1
                }
                var v27 int64 = 9223372036854775807
                _ = v17
                _ = v27
            }
        }
    default:
        if false {
            var v28 float32 = float32(0.1)
            _ = v28
        } else {
            var v29 int16 = 933
            _ = v29
        }
    }
    if true {
        v30 := GenS4{F0: ("a" + ("a" + "")), F1: 394, F2: p2}
        switch "a" {
        case "":
            var v31 uint8 = 74
            _ = v31
        case "foo bar":
            fmt.Println(p0)
        case "hello":
            switch 709 {
            case 32768:
                if (263 != 893) {
                    fmt.Println(p1, p0, p1)
                }
            case 127:
                if false {
                    var v32 float32 = 3.14159
                    _ = v32
                } else {
                    v33 := GenS4{F0: "hello", F1: 32767, F2: p2}
                    v34 := GenS3{F0: 42, F1: 2}
                    _ = v33
                    _ = v34
                }
            default:
                v35 := []int{290, 4294967295}
                _ = v35
            }
        default:
            v36 := []int{-2147483648, 4294967295, -300}
            _ = v36
        }
        _ = v30
    } else {
        fmt.Println(p2, p2)
        p1 = (p1 >> ((255) & 63))
        var v37 int16 = -32768
        _ = v37
    }
    var v38 bool = (!((true && (false || false)) || true))
    var v39 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v38
    _ = v39
    return ((v39 >> ((1) & 63)) > v39)
}

func main() {
    var v1 int = 55
    var v2 uint32 = 124
    v3 := []int{}
    var v40 int8 = 42
    for fi41 := 0; fi41 < 9; fi41++ {
        v40 = v40
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v42 int8 = 127
        _ = fi41
        _ = v42
    }
    if true {
        var v43 int = v1
        v44 := GenS4{F0: ("foo bar" + "\n\t"), F1: v1, F2: -353}
        _ = v43
        _ = v44
    } else {
        for fi45 := 0; fi45 < 4; fi45++ {
            var v46 bool = false
            v1 += 9223372036854775807
            _ = fi45
            _ = v46
        }
    }
    v40 -= (v40 - (v40 | (-(v40))))
    for fi47 := 0; fi47 < 17; fi47++ {
        v48 := GenS4{F0: ("" + "foo bar"), F1: (fi47 ^ fi47), F2: -2147483648}
        var v49 int16 = 1000
        if len(v3) > 0 {
            v3[0] = (-(v1))
        }
        _ = fi47
        _ = v48
        _ = v49
    }
    v40 += v40
    for fi50 := 0; fi50 < 6; fi50++ {
        v51 := GenS4{F0: "hello", F1: -477, F2: -475}
        _ = fi50
        _ = v51
    }
    switch ("a" + "") {
    case "":
        fmt.Println(v1, v2)
    case "foo bar":
        for fi52 := 0; fi52 < 4; fi52++ {
            if true {
                var v53 uint8 = 184
                v3 = append(v3, fi52)
                var v54 int = fi52
                _ = v53
                _ = v54
            }
            _ = fi52
        }
    case "hello":
        var v55 int32 = 1
        _ = v55
    default:
        if true {
            v56 := GenS4{F0: ("foo bar" + "foo bar"), F1: 128, F2: -492}
            _ = v56
        } else {
            var v57 int16 = -588
            _ = v57
        }
    }
    if (!(false || (-1 == 42))) {
        var v58 uint = 1
        var v59 int8 = (v40 * v40)
        _ = v58
        _ = v59
    } else {
        v1 = v1
        v60 := GenS3{F0: 128, F1: 212}
        if (18446744073709551615 > 0) {
            fmt.Println(v40)
        } else {
            fmt.Println(v40, v1, v40)
            v1 -= v1
            v1 = (v1 / ((v1) | 1))
        }
        _ = v60
    }
    _ = len(v3)
    var v61 float32 = 1e10
    if false {
        var v62 bool = (true || true)
        for fi63 := 0; fi63 < 20; fi63++ {
            if (((4294967295 < 65535) || false) || v62) {
                v1 += 2147483647
                var v64 uint32 = 42
                _ = v64
            } else {
                fmt.Println(v62, v62)
                switch v1 {
                case 9223372036854775807:
                    v65 := GenS5{F0: 0, F1: 65535, F2: 788}
                    _ = v65
                }
                var v66 int = 128
                _ = v66
            }
            switch v1 {
            case 127:
                switch ("hello" + "") {
                case "hello":
                    for fi67 := 0; fi67 < 3; fi67++ {
                        var v68 float64 = ((0.0 - 1.5) * (-2.25 + 1e10))
                        var v69 float32 = v61
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                case "foo bar":
                    var v70 bool = ((((1.5 * (0.0 / 1e10)) / (-2.25 / 1.0)) != -2.25) && ((2147483647 <= 32767) || false))
                    _ = v70
                default:
                    var v71 int = (v1 - 2147483647)
                    _ = v71
                }
            default:
                v61 -= v61
            }
            _ = fi63
        }
        var v72 uint32 = v2
        _ = v62
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v40:", v40)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v61
}
