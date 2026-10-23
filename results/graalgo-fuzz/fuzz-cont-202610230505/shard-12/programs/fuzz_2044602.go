package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS2 struct { F0 int; F1 bool; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 string }
func genFn0() uint8 {
    var v6 float64 = 1.5
    v7 := []int{-213, 1, 42}
    switch 59 {
    case 644:
        v6 += v6
    case -9223372036854775808:
        fmt.Println(v6)
    default:
        v8 := GenS1{F0: 2147483648, F1: 42, F2: -230}
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS2{F0: (fi9 % (((fi10 - fi10)) | 1)), F1: (18446744073709551615 != 2), F2: -1}
            v12 := GenS3{F0: 42}
            _ = fi10
            _ = v11
            _ = v12
        }
        v6 -= v6
        v13 := GenS1{F0: 4294967295, F1: 2147483647, F2: 0}
        _ = fi9
        _ = v13
    }
    _ = v6
    _ = v7
    return 255
}
func genFn1() string {
    var v14 uint = 2
    switch -2147483648 {
    case 248:
        var v15 int16 = 0
        _ = v15
    case 9223372036854775807:
        v14 -= (v14 * (^(v14)))
    case -9223372036854775808:
        v16 := GenS1{F0: 253, F1: 2147483648, F2: 1}
        _ = v16
    default:
        if ((("foo bar" + "\n\t") + ("\n\t" + "hello")) <= (("foo bar" + "a") + ("hello" + "a"))) {
            v17 := []int{32767}
            if (((false && false) && ((uint64(v14) == 4294967295) || (true && (32767 >= -1)))) && (!true)) {
                v17 = append(v17, 943)
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            }
            _ = v17
        } else {
            var v18 float64 = 1.5
            v19 := []int{555, -9223372036854775808, 2147483647}
            if true {
                v14 = (v14 % (((^(v14))) | 1))
                switch "hello" {
                case "a":
                    v20 := GenS4{F0: 909}
                    _ = v20
                case "":
                    if (65535 > 0) {
                        v18 = v18
                        if true {
                            var v21 uint64 = 592
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 2; fi22++ {
                            fmt.Println(v14, fi22, v18)
                            _ = fi22
                        }
                    }
                case "foo bar":
                    v14 -= (v14 << (((^(v14))) & 63))
                }
            } else {
                _ = len(v19)
                v23 := GenS5{F0: 857, F1: (("" + "") + ("a" + (("foo bar" + (("" + "") + "hello")) + ("" + "\n\t"))))}
                _ = v23
            }
            _ = v18
            _ = v19
        }
    }
    _ = v14
    return "hello"
}

func main() {
    var v1 uint8 = 0
    var v2 float32 = 0.1
    v3 := []int{}
    v4 := make(map[uint32]uint32)
    v5 := []string{}
    switch (("foo bar" + "\n\t") + "foo bar") {
    case "foo bar":
        v2 = (((v2 / v2) - 0.0) * (v2 / 1.0))
    case "a":
        for fi24 := 0; fi24 < 19; fi24++ {
            _ = len(v4)
            _ = fi24
        }
    }
    v2 -= v2
    switch ("hello" + ("foo bar" + "a")) {
    case "hello":
        if true {
            v2 = v2
            var v25 uint32 = 603
            v26 := v4[0]
            _ = v25
            _ = v26
        } else {
            var v27 uint8 = 1
            switch 2147483647 {
            case -9223372036854775808:
                if (65535 >= 1) {
                    var v28 uint64 = 247
                    _ = v28
                }
            case 127:
                if (2 < 42) {
                    v2 += v2
                    if false {
                        var v29 uint8 = (v1 | v27)
                        if false {
                            v27 -= (v29 & (v29 - (v29 >> ((255) & 63))))
                            v30 := GenS1{F0: 42, F1: 2147483648, F2: 42}
                            v27 += 2
                            _ = v30
                        }
                        _ = v29
                    } else {
                        var v31 uint16 = 255
                        var v32 bool = (false && (!false))
                        _ = v31
                        _ = v32
                    }
                } else {
                    var v33 uint = 65535
                    _ = v33
                }
            default:
                fmt.Println(v1)
            }
            _ = v27
        }
    case "foo bar":
        var v34 int = 1
        _ = v34
    case "":
        v1 += v1
    }
    var v35 float64 = float64(v2)
    var v36 string = (("a" + "a") + "\n\t")
    v36 = (v36 + ((v36 + v36) + v36))
    if true {
        switch (v36 + v36) {
        case "hello":
            fmt.Println(v35, v36, v2)
        default:
            var v37 int8 = -128
            _ = v37
        }
        v35 = float64((float32(v35) + -2.25))
    } else {
        if len(v3) > 0 {
            v3[0] = 551
        }
    }
    v2 = (v2 / v2)
    var v38 uint8 = (v1 ^ genFn0())
    if len(v3) > 0 {
        v3[0] = -464
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        switch v36 {
        case "\n\t":
            var v40 uint = 255
            _ = v40
        default:
            var v41 float64 = v35
            _ = v41
        }
        v42 := GenS5{F0: 65535, F1: v36}
        var v43 string = "\n\t"
        _ = fi39
        _ = v42
        _ = v43
    }
    if true {
        var v44 uint16 = 32767
        _ = v44
    } else {
        var v45 int16 = -610
        var v46 float64 = v35
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 18; fi47++ {
        for fi48 := 0; fi48 < 20; fi48++ {
            if (42 < 1) {
                v49 := GenS4{F0: 1}
                _ = v49
            }
            _ = fi48
        }
        for fi50 := 0; fi50 < 19; fi50++ {
            v51 := GenS2{F0: fi47, F1: (((("" + "foo bar") + "a") <= "hello") && (true && true)), F2: -32768}
            _ = fi50
            _ = v51
        }
        v52 := v4[uint32(v1)]
        _ = fi47
        _ = v52
    }
    v36 = "a"
    if false {
        var v53 int16 = 822
        var v54 uint16 = 2
        var v55 int16 = v53
        _ = v53
        _ = v54
        _ = v55
    } else {
        if (v2 > v2) {
            for fi56 := 0; fi56 < 6; fi56++ {
                v35 = 0.1
                _ = fi56
            }
            var v57 bool = (272 > -32768)
            _ = v57
        } else {
            if false {
                _ = len(v3)
                if (-32768 == 1) {
                    var v58 string = ("hello" + "hello")
                    _ = v58
                } else {
                    if true {
                        v59 := GenS4{F0: 65535}
                        for fi60 := 0; fi60 < 12; fi60++ {
                            var v61 float64 = v35
                            for fi62 := 0; fi62 < 2; fi62++ {
                                delete(v4, 65535)
                                var v63 float64 = 1e10
                                v64 := GenS3{F0: -1}
                                _ = fi62
                                _ = v63
                                _ = v64
                            }
                            _ = fi60
                            _ = v61
                        }
                        _ = v59
                    }
                }
            } else {
                for fi65 := 0; fi65 < 8; fi65++ {
                    switch ("a" + (("hello" + v36) + v36)) {
                    case "a":
                        delete(v4, 1)
                    case "\n\t":
                        var v66 uint8 = v38
                        _ = v66
                    case "hello":
                        v38 = (v1 << ((619) & 63))
                    default:
                        v2 += v2
                    }
                    var v67 string = (((v36 + (v36 + "hello")) + v36) + v36)
                    _ = fi65
                    _ = v67
                }
                var v68 int16 = 401
                _ = v68
            }
        }
        switch int(v1) {
        case -2147483648:
            v4[42] = 208
        case 0:
            var v69 uint32 = 1
            _ = v69
        case 128:
            switch -444 {
            case 42:
                v38 += v38
            case -953:
                var v70 bool = (272 >= 65535)
                _ = v70
            default:
                v71 := v4[474]
                _ = v71
            }
        }
    }
    for fi72 := 0; fi72 < 8; fi72++ {
        v1 -= 127
        _ = fi72
    }
    var v73 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v73
}
