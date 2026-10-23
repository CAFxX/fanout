package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint; F1 int16; F2 int16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 uint64; F1 int16 }
type GenS11 struct { F0 int32; F1 string; F2 bool }
type GenS12 struct { F0 uint32 }
func genFn0(p0 string, p1 int32) uint16 {
    switch p0 {
    case "hello":
        var v6 string = (("foo bar" + ((p0 + ("a" + "a")) + p0)) + "")
        _ = v6
    case "":
        v7 := []int{-2147483648}
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: 65535, F1: true, F2: 127}
            var v10 int8 = -128
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    if true {
        p0 = (p0 + "\n\t")
        fmt.Println(p0)
    }
    var v11 string = ""
    v12 := []int{}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 42
    var v3 int = 42
    var v4 int = -2147483648
    v5 := []uint32{124, 42, 2147483648, 0}
    var v13 int64 = 0
    if (1 > 1) {
        v3 = (v3 / (((v3 % ((32767) | 1))) | 1))
        var v14 int32 = 1
        _ = v14
    } else {
        v4 -= (v3 + v3)
    }
    if true {
        if (1 != 2) {
            v2 = (v2 % ((65535) | 1))
            v15 := GenS2{F0: 2, F1: -1, F2: 0}
            _ = len(v5)
            _ = v15
        }
        v16 := GenS2{F0: 42, F1: 59, F2: 1000}
        for fi17 := 0; fi17 < 10; fi17++ {
            fmt.Println(v13, v3)
            var v18 uint16 = 255
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        if (1 < 26) {
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := GenS3{F0: 0, F1: "\n\t"}
                v1["\n\t"] = 65535
                _ = fi19
                _ = v20
            }
            if true {
                var v21 int = (v3 % ((9223372036854775807) | 1))
                v21 += v4
                var v22 uint = 42
                _ = v21
                _ = v22
            } else {
                v23 := GenS4{F0: 127, F1: true, F2: 42}
                v13 += (v13 - 42)
                _ = len(v5)
                _ = v23
            }
        } else {
            if (((42 < 4294967295) && true) && (-179 <= 0)) {
                fmt.Println(v4, v2)
                v2 = (^(v2))
                v24 := GenS4{F0: 42, F1: true, F2: 42}
                _ = v24
            }
        }
        var v25 bool = false
        var v26 string = (("" + "\n\t") + "")
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        _ = len(v1)
        var v28 uint16 = (^(v2))
        switch "a" {
        case "":
            switch "" {
            case "a":
                v29 := GenS5{F0: (!true)}
                _ = v29
            case "":
                var v30 float64 = ((((1e10 / 1.0) / 1e10) * float64(3.14159)) / -2.25)
                _ = v30
            case "\n\t":
                if (v3 <= v3) {
                    if (false && (1 != 255)) {
                        v31 := GenS6{F0: 2}
                        var v32 uint8 = 127
                        v13 = (v13 ^ (v13 ^ -353))
                        _ = v31
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 14; fi33++ {
                        v13 = 9223372036854775807
                        v4 = -9223372036854775808
                        if (!true) {
                            v3 = v3
                        }
                        _ = fi33
                    }
                }
            }
        default:
            v34 := v1["\n\t"]
            _ = v34
        }
        _ = fi27
        _ = v28
    }
    if ((-681 > -161) || ((!(false && false)) || true)) {
        delete(v1, (("hello" + (("" + "\n\t") + ("hello" + "\n\t"))) + (("a" + ("" + (("foo bar" + (("foo bar" + "") + ("\n\t" + "\n\t"))) + "hello"))) + ("\n\t" + ("foo bar" + "hello")))))
        fmt.Println(v3, v4, v13)
        v2 = genFn0((("a" + (("a" + "hello") + ("\n\t" + "hello"))) + "\n\t"), -355)
    } else {
        delete(v1, "")
    }
    v35 := v1["\n\t"]
    v3 = -9223372036854775808
    fmt.Println(v35, v13)
    v36 := GenS7{F0: 42}
    if (false && (true && false)) {
        v2 -= (^(v2))
        var v37 int = v3
        _ = v37
    }
    if (!(!true)) {
        v38 := GenS8{F0: 0, F1: v35}
        if false {
            v39 := v1["a"]
            v40 := GenS1{F0: (v35 * (v39 % ((2) | 1))), F1: (true || (((1e10 * 0.1) * 0.1) <= (-2.25 - (3.14159 / ((0.1 / -2.25) + 3.14159))))), F2: 1}
            _ = v39
            _ = v40
        }
        _ = v38
    } else {
        _ = len(v1)
        if (!(18446744073709551615 >= 4294967295)) {
            switch "\n\t" {
            case "a":
                switch "" {
                case "a":
                    for fi41 := 0; fi41 < 18; fi41++ {
                        var v42 uint32 = 1
                        _ = fi41
                        _ = v42
                    }
                case "foo bar":
                    switch v3 {
                    case 524:
                        var v43 uint8 = 127
                        _ = v43
                    case -64:
                        for fi44 := 0; fi44 < 10; fi44++ {
                            v45 := v1[""]
                            for fi46 := 0; fi46 < 20; fi46++ {
                                v3 = (v3 * 2147483647)
                                switch "foo bar" {
                                case "hello":
                                    v47 := GenS9{F0: 1, F1: ((true || (false && (!(!false)))) || false)}
                                    _ = v47
                                }
                                v3 = (fi46 % ((127) | 1))
                                _ = fi46
                            }
                            _ = fi44
                            _ = v45
                        }
                    }
                default:
                    for fi48 := 0; fi48 < 10; fi48++ {
                        v13 = v13
                        switch "\n\t" {
                        case "hello":
                            v49 := GenS10{F0: 2, F1: 905}
                            _ = v49
                        }
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            for fi50 := 0; fi50 < 17; fi50++ {
                                var v51 int8 = -1
                                switch fi48 {
                                case 42:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    fmt.Println(v35, v35)
                                case 127:
                                    v52 := GenS11{F0: -367, F1: "\n\t", F2: (!true)}
                                    _ = v52
                                case -1:
                                    if len(v5) > 0 {
                                        v5[0] = 65535
                                    }
                                }
                                _ = fi50
                                _ = v51
                            }
                            var v53 float32 = -2.25
                            _ = v53
                        }
                        _ = fi48
                    }
                }
            case "foo bar":
                var v54 int64 = 1
                _ = v54
            }
            switch int(v13) {
            case 2147483647:
                v2 += (v2 >> ((1) & 63))
            default:
                for fi55 := 0; fi55 < 18; fi55++ {
                    v5 = append(v5, 2147483647)
                    var v56 int32 = 2147483647
                    _ = fi55
                    _ = v56
                }
            }
        }
    }
    for fi57 := 0; fi57 < 18; fi57++ {
        switch "\n\t" {
        case "a":
            v2 = v2
        case "":
            var v58 bool = true
            _ = v58
        default:
            switch "a" {
            case "a":
                fmt.Println(fi57, v3)
            case "foo bar":
                v59 := GenS12{F0: v35}
                _ = v59
            default:
                for fi60 := 0; fi60 < 3; fi60++ {
                    var v61 uint = 65535
                    var v62 uint32 = 42
                    v63 := GenS10{F0: 0, F1: 32767}
                    _ = fi60
                    _ = v61
                    _ = v62
                    _ = v63
                }
            }
        }
        var v64 bool = true
        v3 += 341
        _ = fi57
        _ = v64
    }
    v65 := GenS4{F0: 0, F1: false, F2: 18446744073709551615}
    switch v4 {
    case -550:
        if (18446744073709551615 == 1) {
            _ = len(v5)
            var v66 bool = false
            _ = v66
        } else {
            var v67 bool = ((((0.0 - float64(((float32(3.14159) + (((3.14159 + float32(0.0)) / -2.25) * float32((0.0 / (0.1 / -2.25))))) - 0.1))) - float64(1e10)) * 3.14159) <= float64((-2.25 + (1.5 * ((1.5 + 3.14159) * 1.5)))))
            _ = v67
        }
    case -565:
        switch "a" {
        case "a":
            v68 := v1[("foo bar" + "a")]
            _ = v68
        case "":
            var v69 int64 = (^(v13))
            _ = v69
        }
    }
    v13 += 9223372036854775807
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v65
}
