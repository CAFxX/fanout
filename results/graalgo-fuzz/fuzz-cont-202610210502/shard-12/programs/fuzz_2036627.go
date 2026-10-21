package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint; F1 int16; F2 uint }
type GenS6 struct { F0 int16; F1 uint16; F2 uint16 }
func genFn0(p0 int) int32 {
    var v5 uint64 = 1
    v6 := GenS1{F0: -326, F1: -95}
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: int16(v5), F1: -128}
        _ = fi7
        _ = v8
    }
    p0 = p0
    if ((!(249 == 4294967295)) || true) {
        var v9 string = "hello"
        var v10 uint = 4294967295
        var v11 int32 = 0
        _ = v10
        _ = v11
        _ = v9
    } else {
        var v12 float32 = -2.25
        if ((true && (!(1 < -1))) || (18446744073709551615 == 4294967295)) {
            v12 += v12
        }
        v5 += (v5 << ((255) & 63))
        _ = v12
    }
    _ = p0
    _ = v5
    _ = v6
    return -653
}

func main() {
    v1 := []string{}
    var v2 uint32 = 65535
    var v3 uint = 65535
    var v4 uint16 = 934
    var v13 int16 = 950
    if (4294967295 > 118) {
        v4 += (v4 - v4)
        for fi14 := 0; fi14 < 3; fi14++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (1 <= genFn0((fi14 >> ((1) & 63)))) {
                v13 -= v13
                v13 = (v13 & (v13 & v13))
            } else {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
            }
            _ = fi14
        }
    } else {
        fmt.Println(v13)
    }
    v3 = 505
    if len(v1) > 0 {
        v1[0] = ("" + "\n\t")
    }
    v15 := GenS2{F0: v13}
    if true {
        var v16 string = ("foo bar" + "\n\t")
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 float32 = (3.14159 * 1e10)
            var v19 bool = ((v4 + v4) <= (v4 - 65535))
            for fi20 := 0; fi20 < 15; fi20++ {
                v21 := GenS2{F0: v13}
                v18 = v18
                _ = fi20
                _ = v21
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v16
    } else {
        if true {
            for fi22 := 0; fi22 < 2; fi22++ {
                if false {
                    var v23 int = 597
                    fmt.Println(v23, v4)
                    if false {
                        v3 = (^(v3))
                        for fi24 := 0; fi24 < 18; fi24++ {
                            if false {
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    switch ((("a" + (("hello" + "\n\t") + ("a" + ""))) + "hello") + "\n\t") {
                                    case "":
                                        v25 := GenS2{F0: (v13 << ((0) & 63))}
                                        _ = v25
                                    case "\n\t":
                                        var v26 int32 = 0
                                        _ = v26
                                    case "hello":
                                        var v27 bool = (!true)
                                        var v28 int64 = 0
                                        _ = v27
                                        _ = v28
                                    }
                                case "a":
                                    v2 = 4294967295
                                case "":
                                    if (!true) {
                                        if false {
                                            v29 := GenS1{F0: 164, F1: 127}
                                            v1 = append(v1, "")
                                            _ = v29
                                        } else {
                                            switch (-(v23)) {
                                            case -9223372036854775808:
                                                v13 = v13
                                            case 127:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case 285:
                                                switch 32768 {
                                                case -632:
                                                    _ = len(v1)
                                                    switch (fi22 & -812) {
                                                    case -9223372036854775808:
                                                        var v30 int32 = 1
                                                        _ = v30
                                                    }
                                                    v31 := GenS3{F0: ("hello" + "\n\t"), F1: 0}
                                                    v32 := GenS2{F0: (v13 | v13)}
                                                    _ = v31
                                                    _ = v32
                                                case -874:
                                                    if true {
                                                        var v33 float32 = 0.0
                                                        _ = v33
                                                    }
                                                default:
                                                    if len(v1) > 0 {
                                                        v1[0] = ("hello" + "\n\t")
                                                    }
                                                }
                                            }
                                            v1 = append(v1, "foo bar")
                                            v34 := GenS4{F0: 651, F1: -466}
                                            _ = v34
                                        }
                                        v3 = 255
                                    }
                                }
                            } else {
                                var v35 bool = (fi24 > fi24)
                                _ = v35
                            }
                            _ = fi24
                        }
                        v23 = 127
                    } else {
                        var v36 bool = ((v3 >> ((65535) & 63)) >= (v3 % (((^(v3))) | 1)))
                        var v37 uint32 = (v2 % (((^(v2))) | 1))
                        _ = v36
                        _ = v37
                    }
                    _ = v23
                }
                fmt.Println(v4, v4, v2)
                _ = fi22
            }
            var v38 uint64 = 954
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v38
        } else {
            var v39 uint = v3
            v4 = 1
            _ = v39
        }
        if true {
            v3 -= 42
            v4 -= v4
        }
        if ((2 >= 4294967295) || false) {
            v40 := GenS5{F0: (v3 ^ (v3 << ((v3) & 63))), F1: (v13 & (v13 ^ (v13 + v13))), F2: v3}
            v41 := GenS3{F0: "foo bar", F1: 637}
            for fi42 := 0; fi42 < 10; fi42++ {
                var v43 uint32 = (v2 ^ (v2 >> ((972) & 63)))
                v1 = append(v1, "")
                _ = fi42
                _ = v43
            }
            _ = v40
            _ = v41
        } else {
            for fi44 := 0; fi44 < 14; fi44++ {
                v13 += (v13 >> ((18446744073709551615) & 63))
                _ = fi44
            }
            v2 = 428
            var v45 float64 = float64(1e10)
            _ = v45
        }
    }
    switch -188 {
    case 127:
        if ((-886 >= genFn0(739)) || (!(-68 < -1))) {
            var v46 int8 = 0
            var v47 int8 = -1
            v2 = (v2 % ((910) | 1))
            _ = v46
            _ = v47
        }
    case 117:
        switch ("\n\t" + ("\n\t" + "\n\t")) {
        case "a":
            v3 += 42
        default:
            v4 = (v4 + (v4 / ((42) | 1)))
        }
    case 2147483647:
        for fi48 := 0; fi48 < 1; fi48++ {
            for fi49 := 0; fi49 < 3; fi49++ {
                fmt.Println(fi49, fi48)
                var v50 int8 = 42
                if (true && false) {
                    fmt.Println(fi49, fi49, v13)
                }
                v51 := GenS4{F0: int32(v13), F1: 42}
                var v52 int8 = v50
                _ = fi49
                _ = v50
                _ = v51
                _ = v52
            }
            _ = fi48
        }
    default:
        var v53 float32 = ((0.0 - -2.25) / -2.25)
        _ = v53
    }
    for fi54 := 0; fi54 < 18; fi54++ {
        for fi55 := 0; fi55 < 12; fi55++ {
            var v56 uint8 = 1
            var v57 bool = (215 < 255)
            var v58 int64 = -508
            _ = fi55
            _ = v56
            _ = v57
            _ = v58
        }
        v3 += v3
        v59 := GenS1{F0: (v13 + v13), F1: -128}
        _ = fi54
        _ = v59
    }
    if true {
        fmt.Println(v4, v2, v4)
        v1 = append(v1, ("hello" + "foo bar"))
        v1 = append(v1, "\n\t")
    }
    var v60 uint16 = (^(v4))
    if (true && (1e10 > ((3.14159 + 0.1) * 0.0))) {
        for fi61 := 0; fi61 < 5; fi61++ {
            if len(v1) > 0 {
                v1[0] = ("" + "foo bar")
            }
            var v62 uint16 = (v4 / (((v4 % ((v60) | 1))) | 1))
            v63 := GenS6{F0: (v13 - v13), F1: (v60 | (v60 % ((1) | 1))), F2: 0}
            _ = fi61
            _ = v62
            _ = v63
        }
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v60
}
