package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 uint; F1 uint16; F2 int }
type GenS6 struct { F0 string; F1 uint8; F2 uint32 }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 int8; F1 int32; F2 int64 }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int; F1 uint }
type GenS12 struct { F0 int8 }

func main() {
    v1 := []uint32{2147483647, 1, 4294967295, 42}
    var v2 bool = true
    v3 := []uint32{0}
    v4 := make(map[uint32]bool)
    var v5 float64 = 1.5
    switch "hello" {
    case "\n\t":
        v6 := GenS1{F0: 567, F1: 42, F2: 2}
        _ = v6
    case "":
        switch "a" {
        case "a":
            if v2 {
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
                v7 := GenS1{F0: 0, F1: -1, F2: 24}
                _ = v7
            } else {
                switch 42 {
                case -1:
                    v8 := GenS1{F0: 42, F1: -1, F2: 1}
                    _ = v8
                case 0:
                    fmt.Println(v5, v5)
                case 1:
                    if (false && v2) {
                        if false {
                            if false {
                                for fi9 := 0; fi9 < 19; fi9++ {
                                    if v2 {
                                        switch (-(fi9)) {
                                        case -9223372036854775808:
                                            v10 := GenS2{F0: ((175 <= 689) || ((fi9 % ((32768) | 1)) < (fi9 * (fi9 ^ (fi9 % ((fi9) | 1)))))), F1: "hello"}
                                            _ = v10
                                        case 32768:
                                            v5 -= 0.0
                                        case 32767:
                                            var v11 uint32 = 1
                                            _ = v11
                                        }
                                        if true {
                                            for fi12 := 0; fi12 < 10; fi12++ {
                                                v5 -= 1.5
                                                var v13 float64 = (v5 * v5)
                                                var v14 string = ("" + "\n\t")
                                                _ = fi12
                                                _ = v13
                                                _ = v14
                                            }
                                            _ = len(v3)
                                            v15 := GenS3{F0: 128, F1: "a"}
                                            _ = v15
                                        } else {
                                            switch -2147483648 {
                                            case -507:
                                                var v16 uint64 = 18446744073709551615
                                                _ = v16
                                            case 127:
                                                var v17 string = "foo bar"
                                                _ = v17
                                            }
                                        }
                                    }
                                    var v18 int32 = 42
                                    _ = fi9
                                    _ = v18
                                }
                                if false {
                                    if (-115 < 0) {
                                        if v2 {
                                            for fi19 := 0; fi19 < 14; fi19++ {
                                                v20 := GenS2{F0: (v2 || (!true)), F1: ""}
                                                switch "" {
                                                case "hello":
                                                    if v2 {
                                                        v21 := GenS4{F0: -1, F1: 65535}
                                                        _ = v21
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                default:
                                                    if false {
                                                        var v22 uint = 2
                                                        _ = v22
                                                    }
                                                }
                                                _ = fi19
                                                _ = v20
                                            }
                                        }
                                        if ((1.5 != -2.25) || ((0 > 42) && v2)) {
                                            var v23 int = 0
                                            switch -2147483648 {
                                            case -9223372036854775808:
                                                if (1000 >= 163) {
                                                    for fi24 := 0; fi24 < 6; fi24++ {
                                                        if (!v2) {
                                                            delete(v4, 0)
                                                            v5 += (float64(1e10) + 0.1)
                                                        } else {
                                                            v25 := GenS5{F0: 42, F1: 32767, F2: 774}
                                                            for fi26 := 0; fi26 < 2; fi26++ {
                                                                var v27 float64 = v5
                                                                var v28 int32 = int32(fi26)
                                                                if v2 {
                                                                    var v29 bool = v2
                                                                    var v30 uint = 830
                                                                    if (!(37 <= 13)) {
                                                                        switch v23 {
                                                                        case 783:
                                                                            var v31 int8 = -1
                                                                            _ = v31
                                                                        }
                                                                    } else {
                                                                        var v32 uint16 = 2
                                                                        if v29 {
                                                                            var v33 uint16 = (^(v32))
                                                                            if true {
                                                                                v30 = v30
                                                                                if v2 {
                                                                                    v23 = fi26
                                                                                }
                                                                                fmt.Println(fi26, v32, v27)
                                                                            } else {
                                                                                if ((203 <= 0) || (v2 && (!(!(("foo bar" < "\n\t") || true))))) {
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    var v34 float32 = 1e10
                                                                                    _ = v34
                                                                                }
                                                                                v33 = uint16(v30)
                                                                            }
                                                                            v2 = ((v30 + v30) != v30)
                                                                            _ = v33
                                                                        } else {
                                                                            v4[1] = v2
                                                                            v35 := GenS5{F0: 2, F1: v32, F2: -1}
                                                                            _ = v35
                                                                        }
                                                                        _ = v32
                                                                    }
                                                                    if (42 != -101) {
                                                                        var v36 uint16 = 183
                                                                        v37 := GenS6{F0: ("\n\t" + "a"), F1: 128, F2: 1}
                                                                        v23 += (fi26 / ((9223372036854775807) | 1))
                                                                        _ = v36
                                                                        _ = v37
                                                                    } else {
                                                                        v2 = v2
                                                                        var v38 uint = v30
                                                                        v28 = v28
                                                                        var v39 uint32 = 496
                                                                        _ = v38
                                                                        _ = v39
                                                                    }
                                                                    _ = v29
                                                                    _ = v30
                                                                } else {
                                                                    if ((^(v28)) != v28) {
                                                                        switch ("hello" + "hello") {
                                                                        case "\n\t":
                                                                            switch v23 {
                                                                            case 128:
                                                                                v40 := GenS7{F0: 42, F1: "hello"}
                                                                                _ = v40
                                                                            default:
                                                                                v41 := GenS5{F0: 98, F1: 42, F2: fi26}
                                                                                _ = v41
                                                                            }
                                                                        case "":
                                                                            switch fi26 {
                                                                            case 32767:
                                                                                var v42 uint16 = 65535
                                                                                _ = v42
                                                                            case 127:
                                                                                var v43 int16 = 1000
                                                                                _ = v43
                                                                            case 42:
                                                                                v44 := GenS8{F0: 0, F1: (^(v28)), F2: 1}
                                                                                _ = v44
                                                                            }
                                                                        default:
                                                                            v45 := GenS9{F0: 42, F1: 4294967295}
                                                                            _ = v45
                                                                        }
                                                                        if (true && false) {
                                                                            if v2 {
                                                                                var v46 int16 = 0
                                                                                var v47 int = -2147483648
                                                                                fmt.Println(v47, v5)
                                                                                v48 := GenS10{F0: true}
                                                                                v49 := GenS6{F0: "", F1: 148, F2: 742}
                                                                                v50 := v4[354]
                                                                                v51 := GenS10{F0: v2}
                                                                                _ = v46
                                                                                _ = v47
                                                                                _ = v48
                                                                                _ = v49
                                                                                _ = v50
                                                                                _ = v51
                                                                            } else {
                                                                                if true {
                                                                                    if v2 {
                                                                                        v28 = -1
                                                                                        switch v23 {
                                                                                        case 71:
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                _ = len(v4)
                                                                                            case "":
                                                                                                v27 = -2.25
                                                                                            case "hello":
                                                                                                v52 := GenS10{F0: (-9223372036854775808 != 0)}
                                                                                                _ = v52
                                                                                            }
                                                                                            v27 += (v5 - v5)
                                                                                            if true {
                                                                                                if (!(true && ((2147483648 != 2147483647) && ((fi26 * 100) != (-(v23)))))) {
                                                                                                    _ = len(v1)
                                                                                                    fmt.Println(v2, v2, v27)
                                                                                                }
                                                                                                v53 := GenS11{F0: 1, F1: 18446744073709551615}
                                                                                                _ = len(v3)
                                                                                                _ = v53
                                                                                            }
                                                                                            var v54 int16 = int16(v28)
                                                                                            if true {
                                                                                                v55 := GenS9{F0: 65535, F1: 4294967295}
                                                                                                _ = v55
                                                                                            }
                                                                                            _ = v54
                                                                                        }
                                                                                        var v56 float64 = 3.14159
                                                                                        v57 := GenS12{F0: 100}
                                                                                        v23 += (fi24 & -548)
                                                                                        _ = v56
                                                                                        _ = v57
                                                                                    } else {
                                                                                        var v58 uint16 = 32767
                                                                                        v28 -= 1
                                                                                        _ = v58
                                                                                    }
                                                                                }
                                                                            }
                                                                            if (0 < 2) {
                                                                                if true {
                                                                                    var v59 uint16 = 255
                                                                                    _ = v59
                                                                                } else {
                                                                                    if ((v28 ^ -305) < (v28 / ((-2147483648) | 1))) {
                                                                                        var v60 uint64 = 599
                                                                                        v27 -= (((((3.14159 - (((0.0 - (v5 + v27)) * (v27 / ((0.1 + (float64(0.1) + 0.1)) / ((-2.25 * 1e10) - v27)))) - (v5 * v27))) * (0.1 / ((-2.25 - ((3.14159 / float64((1.5 + (0.0 - (1e10 / 0.1))))) - v27)) - 1e10))) / v27) * v27) / 1.0)
                                                                                        var v61 bool = (((((0.1 - -2.25) == (v27 * v27)) || ((42 > 65535) && v2)) && (false || (-9223372036854775808 >= 941))) && (!(42 == 1)))
                                                                                        _ = v60
                                                                                        _ = v61
                                                                                    } else {
                                                                                        v28 = v28
                                                                                        var v62 bool = (!v2)
                                                                                        _ = v62
                                                                                    }
                                                                                    fmt.Println(v23, v5, v28)
                                                                                    var v63 int64 = 0
                                                                                    _ = v63
                                                                                }
                                                                                if v2 {
                                                                                    if (0 != -1) {
                                                                                        v64 := GenS12{F0: -49}
                                                                                        var v65 int8 = -1
                                                                                        _ = v64
                                                                                        _ = v65
                                                                                    }
                                                                                    v66 := GenS5{F0: 1, F1: 218, F2: 32768}
                                                                                    _ = v66
                                                                                }
                                                                            }
                                                                            v27 -= (v5 + (v5 / v5))
                                                                        } else {
                                                                            var v67 float32 = 0.0
                                                                            var v68 float64 = 1e10
                                                                            _ = v67
                                                                            _ = v68
                                                                        }
                                                                    }
                                                                    var v69 float32 = float32(1e10)
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 42
                                                                    }
                                                                    _ = v69
                                                                }
                                                                _ = fi26
                                                                _ = v27
                                                                _ = v28
                                                            }
                                                            _ = v25
                                                        }
                                                        v5 = (v5 + ((0.1 * v5) / float64(float32(v5))))
                                                        _ = fi24
                                                    }
                                                    for fi70 := 0; fi70 < 5; fi70++ {
                                                        var v71 float64 = -2.25
                                                        _ = fi70
                                                        _ = v71
                                                    }
                                                    var v72 uint32 = 455
                                                    _ = v72
                                                }
                                            default:
                                                var v73 uint64 = 1
                                                _ = v73
                                            }
                                            switch (("hello" + ("hello" + "foo bar")) + "hello") {
                                            case "":
                                                v2 = (0 < 0)
                                            }
                                            _ = v23
                                        } else {
                                            fmt.Println(v2)
                                            v4[2147483647] = (!v2)
                                        }
                                        switch 42 {
                                        case 128:
                                            v5 = -2.25
                                        case 9223372036854775807:
                                            for fi74 := 0; fi74 < 7; fi74++ {
                                                v75 := v4[541]
                                                _ = fi74
                                                _ = v75
                                            }
                                        default:
                                            fmt.Println(v2, v2)
                                        }
                                    }
                                } else {
                                    switch -2147483648 {
                                    case -772:
                                        v2 = v2
                                    case 42:
                                        v5 = (v5 + (3.14159 + v5))
                                    case -42:
                                        v5 += ((v5 + v5) - v5)
                                    }
                                }
                            }
                        } else {
                            v2 = (1 < -101)
                            var v76 int16 = 0
                            _ = v76
                        }
                    } else {
                        v77 := GenS12{F0: 1}
                        v2 = false
                        var v78 uint16 = 255
                        _ = v77
                        _ = v78
                    }
                default:
                    delete(v4, 4294967295)
                }
                v2 = v2
            }
        }
    default:
        v79 := GenS12{F0: 42}
        _ = v79
    }
    v2 = (1 >= 0)
    v2 = (!(!false))
    switch 32767 {
    case 128:
        v80 := GenS12{F0: -128}
        _ = v80
    }
    var v81 float32 = 0.0
    v5 = v5
    for fi82 := 0; fi82 < 10; fi82++ {
        v81 -= (float32(float64(v81)) + 0.0)
        switch 9223372036854775807 {
        case 2147483647:
            v5 -= (((0.1 - float64((float32(v5) / (float32(3.14159) + 3.14159)))) / v5) + v5)
        case -896:
            if ((!v2) || (false && (1 == 4294967295))) {
                switch ("\n\t" + ((("a" + ("\n\t" + ("a" + ("a" + "")))) + "") + ("foo bar" + "foo bar"))) {
                case "foo bar":
                    var v83 int8 = -128
                    _ = v83
                case "hello":
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v84 bool = (false && false)
                        _ = v84
                    } else {
                        var v85 uint32 = 1
                        v2 = false
                        _ = v85
                    }
                case "":
                    v86 := GenS12{F0: 1}
                    _ = v86
                }
            }
        case 4294967295:
            var v87 int8 = 28
            _ = v87
        }
        v81 = 0.1
        _ = fi82
    }
    var v88 int32 = -2147483648
    v2 = (v2 || false)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v81:", v81)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v81
    _ = v88
}
