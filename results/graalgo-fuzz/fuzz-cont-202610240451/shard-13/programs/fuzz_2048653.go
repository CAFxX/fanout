package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 uint8; F1 string; F2 uint32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int32; F1 string }
type GenS12 struct { F0 int8; F1 int8 }
type GenS13 struct { F0 int32; F1 bool; F2 int64 }
type GenS14 struct { F0 int64; F1 int; F2 uint32 }
type GenS15 struct { F0 uint64 }
func genFn0(p0 bool) string {
    p0 = ((0.1 != float64(-2.25)) || (!(false && (p0 && (!(0 < 2))))))
    p0 = (p0 || false)
    if true {
        v6 := GenS1{F0: 2}
        if false {
            fmt.Println(p0)
            for fi7 := 0; fi7 < 11; fi7++ {
                for fi8 := 0; fi8 < 15; fi8++ {
                    var v9 uint = 0
                    var v10 uint16 = 65535
                    if p0 {
                        v11 := GenS2{F0: 1}
                        v12 := []int{(^(fi8)), fi7}
                        v13 := GenS3{F0: (v10 + 32767), F1: 42, F2: v10}
                        _ = v11
                        _ = v12
                        _ = v13
                    } else {
                        v14 := GenS4{F0: (4294967295 <= 1)}
                        fmt.Println(fi8)
                        v9 -= (v9 | 0)
                        _ = v14
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi7
            }
        } else {
            p0 = p0
            if (-1 < 42) {
                var v15 int32 = 1
                if ((true || p0) && p0) {
                    switch -649 {
                    case 4294967295:
                        var v16 int32 = v15
                        _ = v16
                    case -1:
                        v17 := GenS5{F0: 9223372036854775807, F1: (p0 && (42 < 255))}
                        _ = v17
                    case -412:
                        v18 := []int{-299, 9223372036854775807}
                        _ = v18
                    default:
                        v19 := GenS6{F0: -9223372036854775808, F1: 806, F2: 1}
                        _ = v19
                    }
                }
                switch ("\n\t" + "a") {
                case "\n\t":
                    switch 2147483647 {
                    case 874:
                        if (p0 && p0) {
                            for fi20 := 0; fi20 < 13; fi20++ {
                                for fi21 := 0; fi21 < 5; fi21++ {
                                    var v22 bool = false
                                    var v23 float64 = float64(1e10)
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                                v24 := GenS7{F0: 1}
                                _ = fi20
                                _ = v24
                            }
                        }
                    case -9223372036854775808:
                        for fi25 := 0; fi25 < 14; fi25++ {
                            v26 := GenS2{F0: 2147483647}
                            _ = fi25
                            _ = v26
                        }
                    case 4294967295:
                        v27 := GenS8{F0: 454, F1: 0}
                        _ = v27
                    }
                case "hello":
                    v28 := GenS3{F0: 921, F1: 1, F2: 255}
                    _ = v28
                default:
                    var v29 string = ""
                    _ = v29
                }
                _ = v15
            }
            switch 32768 {
            case 0:
                v30 := []int{42, 190}
                _ = v30
            case -1:
                p0 = (!(false && (42 > -508)))
            case -8:
                var v31 uint16 = 142
                _ = v31
            }
        }
        for fi32 := 0; fi32 < 20; fi32++ {
            v33 := GenS9{F0: 223, F1: "\n\t", F2: 2147483648}
            _ = fi32
            _ = v33
        }
        _ = v6
    }
    v34 := []int{-291, 0, 42}
    v35 := GenS10{F0: 100}
    _ = p0
    _ = v34
    _ = v35
    return ("foo bar" + "a")
}

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]string)
    v3 := []int{1}
    var v4 int32 = -2147483648
    var v5 int8 = 42
    var v36 float32 = ((3.14159 + 3.14159) * -2.25)
    fmt.Println(v5)
    v4 += v4
    v37 := GenS8{F0: -1, F1: 342}
    v38 := v2[213]
    for fi39 := 0; fi39 < 3; fi39++ {
        var v40 uint32 = 426
        switch v38 {
        case "\n\t":
            for fi41 := 0; fi41 < 12; fi41++ {
                v42 := GenS1{F0: 0}
                _ = fi41
                _ = v42
            }
        case "":
            v43 := GenS11{F0: (^(v4)), F1: v38}
            _ = v43
        case "hello":
            if false {
                var v44 int = fi39
                _ = v44
            }
        default:
            v45 := GenS2{F0: v40}
            _ = v45
        }
        for fi46 := 0; fi46 < 16; fi46++ {
            fmt.Println(v5, fi46, v5)
            _ = fi46
        }
        _ = fi39
        _ = v40
    }
    var v47 int = 127
    v38 = v38
    switch "a" {
    case "\n\t":
        v48 := GenS12{F0: (v5 % ((v5) | 1)), F1: 100}
        _ = v48
    case "hello":
        for fi49 := 0; fi49 < 16; fi49++ {
            v50 := GenS8{F0: 349, F1: -205}
            _ = fi49
            _ = v50
        }
    }
    for fi51 := 0; fi51 < 7; fi51++ {
        v52 := v2[4294967295]
        for fi53 := 0; fi53 < 13; fi53++ {
            _ = len(v3)
            v54 := GenS5{F0: 1, F1: (true && (!(v52 >= "a")))}
            v55 := GenS7{F0: 0}
            _ = fi53
            _ = v54
            _ = v55
        }
        v38 = (("\n\t" + ("a" + ("" + v38))) + (((("a" + (("foo bar" + v52) + v52)) + v38) + (v38 + "foo bar")) + (v52 + "\n\t")))
        _ = fi51
        _ = v52
    }
    v5 -= 1
    var v56 int32 = (^(v4))
    switch v38 {
    case "foo bar":
        if (!((v1 | (^(v1))) > (v1 << ((v1) & 63)))) {
            var v57 uint64 = 227
            var v58 string = v38
            _ = v57
            _ = v58
        }
    case "a":
        v59 := v2[523]
        _ = v59
    default:
        if false {
            if ((255 >= 255) && (192 != 2)) {
                for fi60 := 0; fi60 < 14; fi60++ {
                    v61 := GenS6{F0: -9223372036854775808, F1: 2147483648, F2: 2}
                    v62 := GenS13{F0: (-(v56)), F1: (v5 > 1), F2: -9223372036854775808}
                    _ = fi60
                    _ = v61
                    _ = v62
                }
                for fi63 := 0; fi63 < 18; fi63++ {
                    if true {
                        var v64 string = v38
                        if false {
                            if (true || false) {
                                var v65 float32 = (1.5 / ((v36 * float32(-2.25)) + (v36 + v36)))
                                _ = v65
                            } else {
                                if (!true) {
                                    var v66 uint16 = 913
                                    delete(v2, 2147483647)
                                    v66 = v66
                                    _ = v66
                                } else {
                                    fmt.Println(fi63, v56, v47)
                                }
                                var v67 int32 = 1
                                v68 := GenS1{F0: 772}
                                _ = v67
                                _ = v68
                            }
                            fmt.Println(v38, v5)
                        } else {
                            var v69 string = v38
                            _ = v69
                        }
                        var v70 int32 = (-(v56))
                        _ = v64
                        _ = v70
                    }
                    fmt.Println(v4, v56)
                    _ = fi63
                }
            }
            v56 += v4
            fmt.Println(v4, v1, v56)
        }
    }
    var v71 bool = false
    if ((0.0 * ((0.1 + (-2.25 * -2.25)) + -2.25)) >= 1.5) {
        _ = len(v3)
    } else {
        switch (-(v47)) {
        case 0:
            v36 -= v36
        default:
            if false {
                for fi72 := 0; fi72 < 20; fi72++ {
                    v4 += v4
                    _ = fi72
                }
                v73 := GenS14{F0: 811, F1: v47, F2: 4294967295}
                var v74 float32 = (v36 - v36)
                _ = v73
                _ = v74
            } else {
                for fi75 := 0; fi75 < 16; fi75++ {
                    for fi76 := 0; fi76 < 8; fi76++ {
                        v1 += (v1 % ((v1) | 1))
                        var v77 uint64 = 549
                        if v71 {
                            var v78 float32 = (v36 / v36)
                            _ = v78
                        } else {
                            switch (v38 + v38) {
                            case "hello":
                                var v79 uint = 42
                                _ = v79
                            case "\n\t":
                                fmt.Println(v36)
                            default:
                                fmt.Println(v36, v77, fi75)
                            }
                        }
                        _ = fi76
                        _ = v77
                    }
                    switch -173 {
                    case -9223372036854775808:
                        fmt.Println(fi75, v56)
                    case 2147483647:
                        var v80 bool = (3.14159 < 0.1)
                        _ = v80
                    case 32767:
                        v2[2147483647] = "a"
                    }
                    _ = fi75
                }
            }
        }
    }
    if false {
        v81 := GenS14{F0: -463, F1: (^(v47)), F2: 42}
        switch v38 {
        case "":
            v82 := GenS15{F0: 256}
            _ = v82
        case "a":
            v4 += v56
        case "\n\t":
            var v83 uint16 = 65535
            _ = v83
        }
        if v71 {
            fmt.Println(v56, v47)
        }
        _ = v81
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v47
    _ = v5
    _ = v56
    _ = v71
}
