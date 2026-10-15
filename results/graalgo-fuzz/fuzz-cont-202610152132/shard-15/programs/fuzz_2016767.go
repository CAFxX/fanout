package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int32; F2 int32 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 int; F1 int16; F2 int8 }
func genFn0(p0 int32, p1 uint16) uint32 {
    switch ("hello" + "") {
    case "hello":
        var v4 bool = false
        _ = v4
    case "a":
        v5 := GenS1{F0: p1}
        _ = v5
    }
    p0 = p0
    p0 = p0
    v6 := []int{-9223372036854775808, 42, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v6
    return uint32(p1)
}

func main() {
    v1 := []uint32{}
    var v2 float32 = 0.0
    v3 := make(map[uint32]bool)
    var v7 uint8 = 224
    v8 := GenS1{F0: 97}
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint64 = 0
        _ = fi9
        _ = v10
    }
    switch "foo bar" {
    case "hello":
        for fi11 := 0; fi11 < 12; fi11++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi11
        }
    case "\n\t":
        if ("\n\t" != "a") {
            var v12 int64 = -133
            var v13 float32 = v2
            fmt.Println(v2, v13, v13)
            _ = v12
            _ = v13
        }
    default:
        var v14 uint8 = v7
        _ = v14
    }
    v7 += v7
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 uint = 1
        for fi17 := 0; fi17 < 11; fi17++ {
            switch (("" + ("" + "hello")) + "\n\t") {
            case "hello":
                switch ((((((((("foo bar" + "foo bar") + ("" + "")) + ("\n\t" + "hello")) + "\n\t") + "\n\t") + "") + (("hello" + ("hello" + ("a" + ("hello" + "")))) + "a")) + "a") + "foo bar") {
                case "hello":
                    fmt.Println(v16)
                case "\n\t":
                    fmt.Println(v16, v2)
                case "foo bar":
                    v16 = v16
                default:
                    var v18 int = fi17
                    _ = v18
                }
            case "a":
                v7 = v7
            case "":
                v7 = (^(v7))
            default:
                v19 := GenS1{F0: 0}
                _ = v19
            }
            var v20 uint8 = (^(v7))
            switch "hello" {
            case "\n\t":
                var v21 float32 = ((float32((float64(1e10) * -2.25)) * ((-2.25 / (1e10 + ((v2 * (float32((float64(((v2 * v2) - (3.14159 * v2))) / -2.25)) - v2)) * (0.0 * v2)))) * (((v2 - (v2 * 1e10)) / (v2 / float32(((-2.25 / (1.5 - (0.1 / 1e10))) + 0.0)))) + 1e10))) / v2)
                _ = v21
            default:
                v7 = (v20 / (((v20 / ((v7) | 1))) | 1))
            }
            _ = fi17
            _ = v20
        }
        v22 := GenS2{F0: 4294967295, F1: -850, F2: -494}
        _ = fi15
        _ = v16
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        for fi24 := 0; fi24 < 6; fi24++ {
            switch (("hello" + "") + "") {
            case "":
                v2 -= float32(-2.25)
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = 254
                }
                v25 := GenS2{F0: 2147483648, F1: 42, F2: -1}
                _ = v25
            default:
                if len(v1) > 0 {
                    v1[0] = 65535
                }
            }
            _ = fi24
        }
        v2 = v2
        var v26 bool = true
        _ = fi23
        _ = v26
    }
    v27 := GenS2{F0: 0, F1: -1, F2: -2147483648}
    if false {
        if false {
            if false {
                v7 = v7
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                var v28 uint32 = 0
                v29 := GenS3{F0: "", F1: 65535}
                _ = len(v3)
                _ = v28
                _ = v29
            }
            var v30 int64 = -9223372036854775808
            if false {
                v7 -= 42
                fmt.Println(v30)
                var v31 uint32 = 4294967295
                _ = v31
            }
            _ = v30
        } else {
            var v32 float64 = ((1.5 * 1.5) * 0.1)
            for fi33 := 0; fi33 < 2; fi33++ {
                v32 += ((1.5 / v32) - (v32 - -2.25))
                fmt.Println(v7, fi33)
                switch fi33 {
                case 86:
                    var v34 bool = true
                    _ = v34
                }
                _ = fi33
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v32
        }
        if (!(0 >= 9223372036854775807)) {
            for fi35 := 0; fi35 < 20; fi35++ {
                v7 = v7
                _ = fi35
            }
        } else {
            fmt.Println(v7)
            if true {
                var v36 uint16 = 646
                v37 := GenS4{F0: 1}
                _ = v36
                _ = v37
            } else {
                v38 := GenS5{F0: (^(v7)), F1: "a"}
                _ = v38
            }
            switch "" {
            case "foo bar":
                for fi39 := 0; fi39 < 14; fi39++ {
                    v2 = v2
                    _ = fi39
                }
            default:
                if (32768 != 4294967295) {
                    var v40 string = "hello"
                    _ = v40
                } else {
                    v7 = (^(v7))
                    var v41 bool = false
                    for fi42 := 0; fi42 < 18; fi42++ {
                        delete(v3, 1)
                        var v43 string = "hello"
                        _ = fi42
                        _ = v43
                    }
                    _ = v41
                }
            }
        }
        if ((((true && ((1 > 1) || false)) && false) && (42 <= 1)) && (-1 == 1)) {
            var v44 uint64 = 4294967295
            _ = v44
        }
    } else {
        for fi45 := 0; fi45 < 14; fi45++ {
            if false {
                v7 += v7
                v46 := GenS6{F0: 0}
                _ = v46
            }
            _ = fi45
        }
        for fi47 := 0; fi47 < 13; fi47++ {
            var v48 float64 = 0.1
            var v49 int = fi47
            for fi50 := 0; fi50 < 18; fi50++ {
                switch "" {
                case "\n\t":
                    v51 := GenS5{F0: (v7 % (((v7 ^ (v7 << ((42) & 63)))) | 1)), F1: ((("a" + "\n\t") + "hello") + "foo bar")}
                    _ = v51
                case "hello":
                    if (421 <= 2) {
                        v52 := GenS4{F0: -18}
                        v53 := GenS7{F0: 1, F1: 0}
                        v2 = (v2 * v2)
                        _ = v52
                        _ = v53
                    } else {
                        v54 := GenS8{F0: fi47, F1: 32767, F2: 127}
                        _ = v54
                    }
                }
                if true {
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                }
                if (!true) {
                    var v55 uint64 = 18446744073709551615
                    switch ("\n\t" + "hello") {
                    case "\n\t":
                        v56 := GenS3{F0: "foo bar", F1: 88}
                        _ = v56
                    case "foo bar":
                        delete(v3, 2147483647)
                    case "a":
                        _ = len(v3)
                    default:
                        switch fi47 {
                        case -962:
                            v57 := GenS5{F0: (v7 * 42), F1: "foo bar"}
                            v58 := v3[4294967295]
                            _ = v57
                            _ = v58
                        case -306:
                            var v59 uint16 = 481
                            _ = v59
                        case 884:
                            v60 := GenS6{F0: 0}
                            _ = v60
                        }
                    }
                    _ = v55
                }
                _ = fi50
            }
            _ = fi47
            _ = v48
            _ = v49
        }
        var v61 uint = 1
        _ = v61
    }
    for fi62 := 0; fi62 < 16; fi62++ {
        if (!false) {
            for fi63 := 0; fi63 < 13; fi63++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v64 := GenS5{F0: v7, F1: (("" + (("hello" + "foo bar") + (("a" + "\n\t") + "a"))) + "\n\t")}
                v2 = v2
                _ = fi63
                _ = v64
            }
        }
        delete(v3, 578)
        var v65 int = -2147483648
        _ = fi62
        _ = v65
    }
    fmt.Println(v2, v7)
    for fi66 := 0; fi66 < 11; fi66++ {
        v7 += v7
        _ = fi66
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v7
    _ = v8
}
