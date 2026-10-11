package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint }
type GenS2 struct { F0 uint; F1 int; F2 uint }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 int16; F1 int; F2 uint64 }
type GenS6 struct { F0 int8 }
func genFn0() string {
    var v6 string = (("" + ("foo bar" + ("foo bar" + "hello"))) + (("foo bar" + (("" + ("" + (("hello" + "a") + ("" + "\n\t")))) + (("\n\t" + "hello") + ("" + "")))) + ("a" + (("hello" + ("\n\t" + ("\n\t" + ""))) + ("hello" + ("" + ""))))))
    var v7 float64 = 3.14159
    _ = v6
    _ = v7
    return "foo bar"
}
func genFn1() int16 {
    v8 := []int{341}
    if false {
        var v9 int64 = 1
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: 1, F1: 9223372036854775807, F2: 144}
            for fi12 := 0; fi12 < 6; fi12++ {
                v8 = append(v8, 4294967295)
                v9 -= 735
                var v13 float32 = -2.25
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        var v14 int = -2147483648
        _ = v14
        _ = v9
    }
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 1
    var v3 float64 = float64(-2.25)
    v4 := []uint32{679}
    var v5 uint8 = 128
    _ = len(v1)
    switch 1 {
    case 32768:
        switch ("\n\t" + "\n\t") {
        case "":
            var v15 float32 = (1.5 + 0.1)
            _ = v15
        default:
            for fi16 := 0; fi16 < 20; fi16++ {
                switch "hello" {
                case "\n\t":
                    v4 = append(v4, 852)
                case "foo bar":
                    var v17 uint32 = 931
                    _ = v17
                case "hello":
                    if false {
                        v18 := GenS2{F0: 618, F1: fi16, F2: 18446744073709551615}
                        var v19 uint8 = 127
                        if (float64(1e10) < v3) {
                            v3 += v3
                            for fi20 := 0; fi20 < 18; fi20++ {
                                var v21 float64 = (1e10 * v3)
                                switch ("\n\t" + "a") {
                                case "":
                                    if len(v4) > 0 {
                                        v4[0] = 2
                                    }
                                    v22 := GenS1{F0: -630, F1: -1, F2: 65535}
                                    _ = v22
                                case "\n\t":
                                    _ = len(v1)
                                default:
                                    v23 := GenS3{F0: (fi16 % (((-(fi20))) | 1)), F1: 65535}
                                    _ = v23
                                }
                                var v24 uint = 65535
                                _ = fi20
                                _ = v21
                                _ = v24
                            }
                            fmt.Println(v2)
                        }
                        _ = v18
                        _ = v19
                    } else {
                        v3 = v3
                    }
                }
                if true {
                    if false {
                        v3 -= 1.5
                    }
                } else {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        if (true || (((3.14159 * (0.0 * ((v3 + (v3 - v3)) * (1.5 / (0.1 * v3))))) / v3) != v3)) {
                            var v26 uint = 18446744073709551615
                            var v27 int64 = 1
                            _ = v26
                            _ = v27
                        }
                        switch -492 {
                        case 42:
                            var v28 int16 = 1
                            _ = v28
                        case 2147483647:
                            v3 = 0.0
                            if (float64((((3.14159 + 1e10) - (((1.5 / -2.25) + 0.0) - -2.25)) * -2.25)) == 1.5) {
                                v29 := GenS4{F0: ("hello" + "a"), F1: -251}
                                switch fi16 {
                                case 128:
                                    if true {
                                        var v30 uint16 = (v2 / (((v2 % (((v2 >> ((431) & 63))) | 1))) | 1))
                                        switch -438 {
                                        case -2147483648:
                                            var v31 uint32 = 373
                                            _ = v31
                                        default:
                                            var v32 float32 = 1.5
                                            _ = v32
                                        }
                                        v5 -= v5
                                        _ = v30
                                    }
                                }
                                _ = v29
                            } else {
                                var v33 uint8 = (v5 / ((v5) | 1))
                                _ = v33
                            }
                            if false {
                                v34 := GenS4{F0: "\n\t", F1: -1}
                                _ = v34
                            }
                        case 127:
                            if (204 <= 57) {
                                v35 := GenS1{F0: 2147483647, F1: 0, F2: 1}
                                if true {
                                    var v36 int8 = int8(v2)
                                    v37 := GenS4{F0: "hello", F1: 549}
                                    _ = v36
                                    _ = v37
                                } else {
                                    fmt.Println(v5, fi25)
                                    fmt.Println(v5)
                                    v5 = (^(v5))
                                }
                                var v38 uint8 = (^(v5))
                                var v39 uint16 = v2
                                v39 = v2
                                _ = v35
                                _ = v38
                                _ = v39
                            } else {
                                v4 = append(v4, 1)
                                var v40 int32 = 42
                                _ = v40
                            }
                        }
                        v41 := GenS1{F0: -236, F1: 936, F2: 992}
                        v2 -= (v2 >> ((42) & 63))
                        var v42 float32 = (3.14159 + 1.5)
                        v5 += 2
                        _ = fi25
                        _ = v41
                        _ = v42
                    }
                    var v43 uint = 42
                    _ = v43
                }
                _ = fi16
            }
        }
    case -467:
        if len(v4) > 0 {
            v4[0] = 2147483648
        }
    }
    fmt.Println(v3, v5, v3)
    switch 794 {
    case -9223372036854775808:
        v3 += v3
    case 523:
        v5 += v5
    case 42:
        switch "hello" {
        case "hello":
            var v44 int32 = 0
            _ = v44
        }
    default:
        for fi45 := 0; fi45 < 9; fi45++ {
            fmt.Println(v3)
            _ = fi45
        }
    }
    if false {
        var v46 int64 = -9223372036854775808
        switch -1 {
        case 128:
            if true {
                fmt.Println(v5, v46)
            }
        }
        if false {
            v47 := GenS5{F0: 42, F1: -2147483648, F2: 1}
            _ = v47
        } else {
            for fi48 := 0; fi48 < 17; fi48++ {
                _ = len(v4)
                _ = fi48
            }
            for fi49 := 0; fi49 < 15; fi49++ {
                v46 = 1
                for fi50 := 0; fi50 < 4; fi50++ {
                    v2 -= v2
                    if (!true) {
                        var v51 bool = (true && (!false))
                        var v52 uint8 = (v5 << ((296) & 63))
                        var v53 int16 = 542
                        _ = v51
                        _ = v52
                        _ = v53
                    } else {
                        v54 := GenS3{F0: (fi49 << ((42) & 63)), F1: 2147483647}
                        delete(v1, 32767)
                        v3 += (v3 - v3)
                        if (false && true) {
                            if false {
                                fmt.Println(v2)
                                v46 = 9223372036854775807
                                if true {
                                    v4 = append(v4, 4294967295)
                                } else {
                                    v4 = append(v4, 2147483648)
                                }
                            } else {
                                v55 := GenS3{F0: int(v5), F1: uint32(v5)}
                                v5 = (^(v5))
                                _ = v55
                            }
                        }
                        var v56 int64 = v46
                        _ = v54
                        _ = v56
                    }
                    if (true && true) {
                        v1[31] = 2
                    }
                    _ = fi50
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi49
            }
        }
        _ = v46
    } else {
        v1[32768] = 659
        for fi57 := 0; fi57 < 13; fi57++ {
            var v58 uint64 = 0
            v5 += v5
            _ = fi57
            _ = v58
        }
        v59 := GenS3{F0: 857, F1: 2}
        _ = v59
    }
    v60 := GenS4{F0: "foo bar", F1: -1}
    switch ("hello" + "\n\t") {
    case "\n\t":
        switch "\n\t" {
        case "":
            v3 = 0.0
        case "\n\t":
            var v61 float32 = -2.25
            _ = v61
        case "foo bar":
            fmt.Println(v2, v3)
        }
    case "a":
        var v62 uint32 = 65535
        _ = v62
    default:
        v5 = 255
    }
    var v63 string = ("" + "hello")
    var v64 float64 = ((-2.25 / v3) / v3)
    var v65 float32 = -2.25
    switch "hello" {
    case "\n\t":
        var v66 int = -346
        _ = v66
    case "foo bar":
        v67 := GenS4{F0: ((("" + "hello") + v63) + ""), F1: -1}
        _ = v67
    case "":
        for fi68 := 0; fi68 < 5; fi68++ {
            _ = len(v4)
            _ = fi68
        }
    default:
        var v69 string = "a"
        _ = v69
    }
    for fi70 := 0; fi70 < 8; fi70++ {
        v71 := GenS3{F0: fi70, F1: 65535}
        var v72 int64 = 9223372036854775807
        for fi73 := 0; fi73 < 9; fi73++ {
            var v74 string = "a"
            var v75 uint32 = uint32(v2)
            v76 := GenS6{F0: -1}
            _ = fi73
            _ = v74
            _ = v75
            _ = v76
        }
        _ = fi70
        _ = v71
        _ = v72
    }
    v1[-1] = 0
    if (2 == 58) {
        fmt.Println(v65, v5)
        fmt.Println(v3, v2, v2)
        var v77 float64 = (float64((v65 - v65)) + 0.1)
        _ = v77
    }
    var v78 uint64 = 18446744073709551615
    var v79 bool = true
    var v80 uint16 = (v2 - v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v63
    _ = v64
    _ = v65
    _ = v78
    _ = v79
    _ = v80
}
