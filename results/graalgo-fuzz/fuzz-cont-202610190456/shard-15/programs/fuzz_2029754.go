package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 int8; F2 uint8 }
type GenS7 struct { F0 bool; F1 uint; F2 int8 }
func genFn0(p0 bool, p1 uint16) uint8 {
    for fi5 := 0; fi5 < 13; fi5++ {
        if p0 {
            var v6 int32 = -2147483648
            p1 += p1
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 uint64 = 0
                _ = fi7
                _ = v8
            }
            _ = v6
        } else {
            v9 := GenS1{F0: 178}
            var v10 float32 = (((1e10 * (((float32(0.1) * (3.14159 - 0.0)) + -2.25) * 1.5)) / (0.0 - 1.5)) + (((((1.5 + 0.1) * (0.0 + float32((0.0 * ((1.5 - 0.1) - 1e10))))) - float32(0.1)) + 1.5) / (0.1 - (-2.25 * 0.0))))
            _ = v10
            _ = v9
        }
        _ = fi5
    }
    var v11 int32 = -2147483648
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS1{F0: 1}
            v11 = (^(v11))
            _ = fi13
            _ = v14
        }
        p0 = (p0 && p0)
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v11
    return 118
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 4294967295
    var v3 bool = true
    var v4 int8 = -31
    for fi15 := 0; fi15 < 7; fi15++ {
        delete(v1, "\n\t")
        v16 := GenS2{F0: false, F1: "foo bar"}
        for fi17 := 0; fi17 < 11; fi17++ {
            fmt.Println(fi15, fi17, fi15)
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    if true {
        _ = len(v1)
    } else {
        if (!(!((-(v4)) == (^(v4))))) {
            v3 = v3
        } else {
            v3 = v3
            _ = len(v1)
            if v3 {
                v18 := GenS3{F0: 0}
                v2 += v2
                v4 += (-(v4))
                _ = v18
            }
        }
        if true {
            if v3 {
                v19 := GenS2{F0: ((v3 || (true && (((-2.25 * ((float32(0.0) + (((0.1 - (3.14159 / -2.25)) + ((3.14159 / ((0.0 + 3.14159) + 0.0)) - 0.1)) - -2.25)) / (-2.25 + 0.1))) / 1.5) > (float32(((3.14159 * 1.5) - (((-2.25 - -2.25) - float64((1e10 * (1e10 + 1.5)))) * 1.5))) + ((0.1 * (0.0 / float32(1e10))) * ((float32((-2.25 + 0.1)) + 0.0) + 1.5)))))) || false), F1: "a"}
                for fi20 := 0; fi20 < 12; fi20++ {
                    v3 = true
                    _ = fi20
                }
                if ((18446744073709551615 != 65535) || (22 <= 1000)) {
                    v4 = v4
                } else {
                    fmt.Println(v3)
                    _ = len(v1)
                    var v21 string = ("foo bar" + "")
                    _ = v21
                }
                _ = v19
            }
            v22 := []int{32768, 9223372036854775807, 42}
            _ = v22
        }
    }
    v3 = ((v2 << ((4294967295) & 63)) != (v2 - (v2 ^ (v2 | (^(v2))))))
    var v23 int8 = 56
    switch 128 {
    case 32768:
        v4 = v4
    case -1:
        var v24 int = -989
        _ = v24
    }
    v23 += (v23 - (v23 ^ (-(v23))))
    switch 1 {
    case 128:
        if true {
            v25 := GenS1{F0: 128}
            _ = v25
        } else {
            var v26 uint16 = 102
            switch "hello" {
            case "\n\t":
                for fi27 := 0; fi27 < 18; fi27++ {
                    if v3 {
                        v28 := []int{42}
                        for fi29 := 0; fi29 < 9; fi29++ {
                            if (!v3) {
                                switch ("" + "\n\t") {
                                case "a":
                                    v4 -= (^(v23))
                                    var v30 float32 = 1e10
                                    var v31 int32 = -909
                                    _ = v30
                                    _ = v31
                                case "\n\t":
                                    if (v3 || (!(4294967295 <= 622))) {
                                        v26 = (v26 | 65535)
                                        var v32 bool = ((fi27 - fi29) > -621)
                                        v33 := GenS4{F0: 128, F1: 194}
                                        _ = v32
                                        _ = v33
                                    }
                                default:
                                    switch -2147483648 {
                                    case 9223372036854775807:
                                        v1[("foo bar" + "a")] = ((^(fi29)) != (fi27 & (fi29 / ((307) | 1))))
                                    case 32767:
                                        if ("a" == ("hello" + "a")) {
                                            var v34 int8 = -113
                                            _ = v34
                                        }
                                        if true {
                                            var v35 float32 = 0.1
                                            v4 -= 0
                                            _ = v35
                                        }
                                        fmt.Println(fi29)
                                    case 4294967295:
                                        v23 += (v23 << ((736) & 63))
                                    default:
                                        var v36 int = fi29
                                        _ = v36
                                    }
                                }
                                switch (("foo bar" + ("foo bar" + (("" + (("hello" + "\n\t") + ("a" + ""))) + ("foo bar" + "foo bar")))) + "foo bar") {
                                case "\n\t":
                                    var v37 string = "foo bar"
                                    _ = v37
                                case "foo bar":
                                    v38 := GenS1{F0: 1}
                                    _ = v38
                                case "hello":
                                    var v39 int32 = -602
                                    _ = v39
                                default:
                                    var v40 int8 = (v23 / (((-(v4))) | 1))
                                    _ = v40
                                }
                            }
                            _ = fi29
                        }
                        _ = v28
                    }
                    if v3 {
                        v41 := v1[(("a" + "hello") + "foo bar")]
                        _ = v41
                    } else {
                        if v3 {
                            v4 = int8(v2)
                            if (120 <= 1) {
                                for fi42 := 0; fi42 < 12; fi42++ {
                                    var v43 string = "\n\t"
                                    v2 += v2
                                    _ = fi42
                                    _ = v43
                                }
                            }
                        }
                    }
                    v3 = ((v4 << ((2) & 63)) <= 1)
                    _ = fi27
                }
            case "a":
                var v44 string = "hello"
                _ = v44
            case "hello":
                for fi45 := 0; fi45 < 2; fi45++ {
                    v46 := []int{32767, (^(fi45))}
                    if ((!(!false)) || false) {
                        for fi47 := 0; fi47 < 10; fi47++ {
                            var v48 int8 = (v23 * -128)
                            var v49 uint32 = 0
                            switch ("\n\t" + ("\n\t" + (("" + (("a" + (("" + ("foo bar" + "hello")) + (("\n\t" + ("" + ("hello" + ""))) + "foo bar"))) + "\n\t")) + "a"))) {
                            case "a":
                                v49 += v49
                            case "foo bar":
                                v23 -= (v4 >> ((255) & 63))
                            default:
                                var v50 uint32 = (v49 / ((v49) | 1))
                                _ = v50
                            }
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                    } else {
                        var v51 int = -1
                        for fi52 := 0; fi52 < 19; fi52++ {
                            delete(v1, ("hello" + ("hello" + ("\n\t" + ""))))
                            v2 = v2
                            _ = fi52
                        }
                        fmt.Println(v51, v3)
                        _ = v51
                    }
                    _ = fi45
                    _ = v46
                }
            }
            v53 := GenS5{F0: ""}
            _ = v26
            _ = v53
        }
    case 4294967295:
        var v54 uint64 = 42
        _ = v54
    }
    var v55 uint = 1
    if v3 {
        v56 := GenS6{F0: v4, F1: v23, F2: 42}
        var v57 float64 = 3.14159
        fmt.Println(v57, v55)
        _ = v56
        _ = v57
    } else {
        v1[("foo bar" + "a")] = ((((-1 == -1) || ((^(v2)) <= 18446744073709551615)) || v3) || ((v3 && (!false)) || (!(127 > 128))))
        if ((-2147483648 < 1) && v3) {
            switch "" {
            case "hello":
                v58 := GenS1{F0: 128}
                _ = v58
            case "":
                v59 := []int{}
                _ = v59
            default:
                v60 := []int{}
                _ = v60
            }
            if (-128 == (-(v23))) {
                var v61 uint = v55
                _ = v61
            }
            var v62 int16 = -32768
            _ = v62
        }
    }
    switch "a" {
    case "hello":
        var v63 string = ""
        _ = v63
    default:
        v64 := GenS7{F0: false, F1: (v55 + (v55 << ((42) & 63))), F2: v23}
        _ = v64
    }
    var v65 int16 = 32767
    if ((2147483647 >= -2147483648) && (!(3.14159 > 3.14159))) {
        for fi66 := 0; fi66 < 19; fi66++ {
            v2 += (v2 * 1)
            v55 = v55
            _ = fi66
        }
        var v67 int64 = -1
        _ = v67
    } else {
        v23 = -128
        v68 := []int{-910, 128, 351, 127}
        _ = v68
    }
    for fi69 := 0; fi69 < 19; fi69++ {
        var v70 string = ""
        delete(v1, (v70 + "\n\t"))
        v71 := v1[""]
        _ = fi69
        _ = v70
        _ = v71
    }
    _ = len(v1)
    v72 := []int{-2147483648, 128}
    if ((0 <= 0) && v3) {
        v73 := GenS3{F0: 255}
        if len(v72) > 0 {
            v72[0] = -2147483648
        }
        _ = v73
    } else {
        for fi74 := 0; fi74 < 2; fi74++ {
            var v75 int16 = 0
            _ = fi74
            _ = v75
        }
        if ((v3 && v3) || v3) {
            v65 += v65
        } else {
            for fi76 := 0; fi76 < 1; fi76++ {
                if (!(!(v55 <= v55))) {
                    v55 = v55
                }
                var v77 int16 = v65
                _ = fi76
                _ = v77
            }
            switch ("\n\t" + "a") {
            case "foo bar":
                _ = len(v1)
            default:
                v2 = v2
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v72:", len(v72), v72)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v55
    _ = v65
    _ = v72
}
