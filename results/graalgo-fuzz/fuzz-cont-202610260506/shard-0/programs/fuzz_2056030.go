package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint32; F2 int32 }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 uint64; F1 int16 }
func genFn0() uint16 {
    v7 := GenS1{F0: 1, F1: 127, F2: -32768}
    var v8 uint32 = 2147483648
    var v9 bool = true
    v8 += (v8 * v8)
    for fi10 := 0; fi10 < 2; fi10++ {
        switch ((("foo bar" + "") + ("\n\t" + "\n\t")) + ("hello" + "a")) {
        case "\n\t":
            switch fi10 {
            case -9223372036854775808:
                switch (fi10 | (fi10 | fi10)) {
                case 32767:
                    v9 = v9
                case -2147483648:
                    v9 = (-1 != -128)
                case 32768:
                    fmt.Println(v9)
                }
            case -514:
                v11 := []int{fi10, fi10, (^(fi10))}
                _ = v11
            case 32768:
                v12 := GenS2{F0: 0}
                _ = v12
            default:
                v13 := GenS1{F0: -207, F1: fi10, F2: 32767}
                _ = v13
            }
        default:
            var v14 bool = (1 < 2147483647)
            _ = v14
        }
        v15 := []int{}
        _ = len(v15)
        _ = fi10
        _ = v15
    }
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint32, p1 int8) uint32 {
    p1 = (p1 - -128)
    p1 -= p1
    _ = p0
    _ = p1
    return (p0 ^ (^(p0)))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    v3 := []int{}
    var v4 uint32 = 280
    var v5 uint8 = 0
    v6 := make(map[string]string)
    v5 += 128
    var v16 int = 128
    var v17 float64 = 3.14159
    if (true && false) {
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 int32 = -725
            var v20 float64 = ((3.14159 * (3.14159 * 0.0)) / v17)
            _ = fi18
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 8; fi21++ {
            if false {
                var v22 uint = 174
                _ = v22
            } else {
                _ = len(v3)
                var v23 int32 = 42
                switch ("hello" + ("\n\t" + (((((("hello" + "\n\t") + ("foo bar" + (("" + "foo bar") + "hello"))) + ("" + "a")) + (("foo bar" + ((("a" + ("hello" + "a")) + "foo bar") + "hello")) + "")) + "foo bar") + "hello"))) {
                case "foo bar":
                    var v24 int64 = -648
                    _ = v24
                default:
                    if (false || (false && false)) {
                        var v25 int64 = 9223372036854775807
                        var v26 string = "\n\t"
                        _ = v25
                        _ = v26
                    } else {
                        switch int(v4) {
                        case 1:
                            for fi27 := 0; fi27 < 11; fi27++ {
                                v28 := GenS2{F0: 1}
                                switch (-(fi21)) {
                                case 2147483647:
                                    fmt.Println(v5)
                                case -1:
                                    if false {
                                        switch "" {
                                        case "a":
                                            if (("hello" + "\n\t") != ("\n\t" + "a")) {
                                                var v29 uint16 = 756
                                                v30 := GenS3{F0: "hello", F1: 42}
                                                var v31 uint = 560
                                                var v32 uint32 = v4
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            }
                                        case "hello":
                                            var v33 string = ""
                                            _ = v33
                                        case "":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    }
                                default:
                                    fmt.Println(fi27, v23, v4)
                                }
                                v17 -= (-2.25 * v17)
                                _ = fi27
                                _ = v28
                            }
                        case 42:
                            v34 := GenS1{F0: 1, F1: -348, F2: 1000}
                            _ = v34
                        case 0:
                            fmt.Println(v5, v23, v5)
                        }
                    }
                }
                _ = v23
            }
            if len(v3) > 0 {
                v3[0] = v16
            }
            var v35 int16 = 1000
            _ = fi21
            _ = v35
        }
    }
    v4 = (v4 * (v4 << ((426) & 63)))
    v17 = v17
    v4 -= v4
    v5 = (^(v5))
    if true {
        fmt.Println(v4, v4, v17)
        v36 := GenS3{F0: "foo bar", F1: 9223372036854775807}
        if true {
            v37 := GenS3{F0: "\n\t", F1: 42}
            _ = v37
        }
        _ = v36
    }
    var v38 int = v16
    for fi39 := 0; fi39 < 1; fi39++ {
        var v40 string = "hello"
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 15; fi41++ {
        v4 = (v4 ^ (v4 % (((^(v4))) | 1)))
        var v42 int8 = -1
        var v43 int64 = 9223372036854775807
        _ = fi41
        _ = v42
        _ = v43
    }
    var v44 uint64 = 2
    v45 := GenS4{F0: 4294967295}
    switch v38 {
    case -1:
        if (0.1 > (1e10 / (1.5 / 1.0))) {
            v46 := GenS5{F0: 65535, F1: (v4 ^ (^(v4))), F2: 438}
            _ = v46
        } else {
            for fi47 := 0; fi47 < 1; fi47++ {
                var v48 uint32 = (v4 + uint32(v44))
                _ = fi47
                _ = v48
            }
        }
    default:
        v49 := v2[(v38 ^ v38)]
        _ = v49
    }
    if (-9223372036854775808 != 831) {
        if (!(65535 > 65535)) {
            for fi50 := 0; fi50 < 8; fi50++ {
                if true {
                    switch "foo bar" {
                    case "a":
                        if false {
                            if ((1.5 >= (0.1 - -2.25)) && (255 < 926)) {
                                switch -829 {
                                case 1:
                                    var v51 int = (fi50 % ((fi50) | 1))
                                    _ = v51
                                }
                            }
                            if (4294967295 <= 609) {
                                switch "\n\t" {
                                case "hello":
                                    for fi52 := 0; fi52 < 10; fi52++ {
                                        var v53 uint16 = 0
                                        v53 = 65535
                                        _ = fi52
                                        _ = v53
                                    }
                                default:
                                    switch "hello" {
                                    case "foo bar":
                                        if (false || (!false)) {
                                            v16 -= -9223372036854775808
                                            for fi54 := 0; fi54 < 20; fi54++ {
                                                var v55 uint = 220
                                                _ = fi54
                                                _ = v55
                                            }
                                        }
                                    case "\n\t":
                                        fmt.Println(v5, v17)
                                    case "a":
                                        for fi56 := 0; fi56 < 11; fi56++ {
                                            fmt.Println(v4)
                                            v57 := GenS3{F0: "", F1: 1}
                                            _ = fi56
                                            _ = v57
                                        }
                                    default:
                                        var v58 uint = 2
                                        _ = v58
                                    }
                                }
                            }
                        } else {
                            for fi59 := 0; fi59 < 2; fi59++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi59
                            }
                            v5 += v5
                            var v60 uint8 = 42
                            _ = v60
                        }
                    default:
                        v44 = 4294967295
                    }
                }
                v44 = v44
                for fi61 := 0; fi61 < 17; fi61++ {
                    switch 32767 {
                    case 183:
                        fmt.Println(v44, v5, v4)
                    case -451:
                        fmt.Println(v17, v44)
                    case 610:
                        if true {
                            switch ("" + "a") {
                            case "\n\t":
                                var v62 int32 = 42
                                _ = v62
                            case "a":
                                switch "a" {
                                case "\n\t":
                                    fmt.Println(v38, v38)
                                case "foo bar":
                                    v38 = fi61
                                case "hello":
                                    v63 := GenS5{F0: 878, F1: v4, F2: 109}
                                    v16 = 32768
                                    _ = v63
                                }
                            default:
                                var v64 bool = (255 <= uint16(v4))
                                _ = v64
                            }
                            v4 = v4
                            delete(v6, ((("a" + (("\n\t" + "a") + (((("a" + "a") + "") + "hello") + ("" + "")))) + ("a" + "foo bar")) + "a"))
                        }
                    default:
                        v65 := GenS6{F0: int(v44), F1: 130}
                        _ = v65
                    }
                    fmt.Println(v16)
                    v44 += 0
                    _ = fi61
                }
                _ = fi50
            }
        } else {
            v66 := GenS4{F0: 1}
            _ = v66
        }
        var v67 float64 = 1.5
        v68 := GenS7{F0: v44, F1: -32768}
        _ = v67
        _ = v68
    } else {
        v69 := GenS7{F0: uint64(v4), F1: 1000}
        _ = v69
    }
    v16 = v38
    for fi70 := 0; fi70 < 12; fi70++ {
        _ = len(v6)
        for fi71 := 0; fi71 < 19; fi71++ {
            fmt.Println(v16, v44, v44)
            _ = fi71
        }
        _ = fi70
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
