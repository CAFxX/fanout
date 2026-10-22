package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 int }
func genFn0() uint {
    switch "a" {
    case "a":
        var v7 bool = false
        _ = v7
    case "\n\t":
        var v8 uint64 = 4294967295
        _ = v8
    default:
        var v9 float64 = 0.0
        _ = v9
    }
    v10 := []int{134, -340}
    for fi11 := 0; fi11 < 1; fi11++ {
        if (false && (((((float64((0.0 + (1e10 * 3.14159))) + 0.0) >= 1e10) || false) || ((1 >= 457) && (!true))) || (9223372036854775807 != -783))) {
            fmt.Println(fi11)
        } else {
            if false {
                if (326 > 263) {
                    var v12 uint32 = 4294967295
                    if false {
                        v12 -= (v12 ^ (v12 << ((1) & 63)))
                        var v13 int64 = -9223372036854775808
                        switch (fi11 >> ((1) & 63)) {
                        case 187:
                            var v14 uint64 = 2
                            _ = v14
                        case 9223372036854775807:
                            v12 = 2147483647
                        }
                        _ = v13
                    }
                    var v15 int = -2147483648
                    _ = v12
                    _ = v15
                }
                var v16 uint32 = 4294967295
                _ = v16
            }
            switch fi11 {
            case -881:
                for fi17 := 0; fi17 < 3; fi17++ {
                    v18 := GenS1{F0: 2, F1: 432}
                    _ = fi17
                    _ = v18
                }
            default:
                var v19 uint16 = 1
                _ = v19
            }
            var v20 int = 2147483647
            _ = v20
        }
        if true {
            v21 := GenS2{F0: 1}
            _ = v21
        } else {
            var v22 uint8 = 42
            _ = v22
        }
        v23 := GenS3{F0: -1, F1: false}
        _ = fi11
        _ = v23
    }
    _ = v10
    return 65535
}
func genFn1() int8 {
    v24 := []int{32767}
    var v25 bool = (!false)
    switch ("a" + ("a" + "")) {
    case "foo bar":
        var v26 uint16 = 518
        _ = v26
    default:
        switch (("" + "foo bar") + "hello") {
        case "foo bar":
            v25 = (-681 == 42)
        case "hello":
            if (228 <= 700) {
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
            } else {
                switch 32767 {
                case -1:
                    var v27 uint8 = 128
                    _ = v27
                case 1:
                    fmt.Println(v25)
                }
            }
        case "a":
            fmt.Println(v25)
        default:
            if (!v25) {
                v24 = append(v24, 986)
            } else {
                v25 = false
                for fi28 := 0; fi28 < 8; fi28++ {
                    if (0 >= 42) {
                        fmt.Println(fi28)
                        var v29 float64 = ((1.5 + 0.1) / -2.25)
                        _ = v29
                    } else {
                        if false {
                            for fi30 := 0; fi30 < 2; fi30++ {
                                if (true && (-570 > 0)) {
                                    fmt.Println(fi28)
                                    v25 = (42 == 0)
                                    var v31 string = "a"
                                    var v32 int8 = -5
                                    _ = v31
                                    _ = v32
                                }
                                var v33 int16 = 669
                                v34 := GenS3{F0: -9223372036854775808, F1: ((2 == 42) || v25)}
                                _ = fi30
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(v25)
                        }
                    }
                    v25 = ((-370 >= -913) && (v25 || v25))
                    if len(v24) > 0 {
                        v24[0] = (-(fi28))
                    }
                    _ = fi28
                }
                fmt.Println(v25)
            }
        }
    }
    v25 = (!v25)
    switch -1 {
    case 2147483647:
        v35 := GenS3{F0: 1, F1: (65535 >= 2)}
        _ = v35
    case 127:
        fmt.Println(v25)
    default:
        fmt.Println(v25)
    }
    _ = v24
    _ = v25
    return -1
}

func main() {
    v1 := []string{"\n\t", "", "foo bar", ("foo bar" + "")}
    var v2 int64 = 42
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]bool)
    var v5 uint8 = 165
    var v6 uint16 = 2
    if (true && true) {
        v36 := GenS3{F0: (^(v2)), F1: (genFn1() < 100)}
        _ = v36
    } else {
        for fi37 := 0; fi37 < 3; fi37++ {
            var v38 uint16 = (v6 * v6)
            var v39 int = (fi37 << ((genFn0()) & 63))
            var v40 uint = 0
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
    }
    v41 := GenS4{F0: 542}
    v42 := GenS2{F0: v6}
    v2 = v2
    var v43 string = ((("foo bar" + "a") + "") + ("\n\t" + (("a" + ("foo bar" + "foo bar")) + ("" + "foo bar"))))
    if (false || (false || (-2147483648 != -2147483648))) {
        fmt.Println(v6, v43)
        if false {
            if true {
                for fi44 := 0; fi44 < 18; fi44++ {
                    var v45 float64 = float64(0.0)
                    switch v43 {
                    case "":
                        v45 += v45
                    case "a":
                        var v46 uint16 = uint16(v3)
                        _ = v46
                    default:
                        var v47 uint8 = v5
                        _ = v47
                    }
                    fmt.Println(v43)
                    _ = fi44
                    _ = v45
                }
                var v48 float32 = (1e10 - (0.0 * 0.1))
                v1 = append(v1, v43)
                _ = v48
            } else {
                for fi49 := 0; fi49 < 10; fi49++ {
                    v6 += 2
                    v6 -= (v6 << ((65535) & 63))
                    v2 = v2
                    _ = fi49
                }
                var v50 int16 = 0
                v3 += v3
                _ = v50
            }
            if (0 < 547) {
                v2 = -9223372036854775808
                for fi51 := 0; fi51 < 4; fi51++ {
                    v52 := GenS4{F0: 1}
                    for fi53 := 0; fi53 < 14; fi53++ {
                        fmt.Println(v6, v43)
                        v3 = (^(v3))
                        _ = len(v4)
                        _ = fi53
                    }
                    fmt.Println(fi51, v6)
                    _ = fi51
                    _ = v52
                }
                if false {
                    v54 := GenS3{F0: (v2 & (v2 / ((-741) | 1))), F1: false}
                    var v55 int8 = 88
                    v3 += 1
                    _ = v54
                    _ = v55
                } else {
                    var v56 int8 = 1
                    _ = v56
                }
            }
            for fi57 := 0; fi57 < 16; fi57++ {
                v58 := v4[893]
                for fi59 := 0; fi59 < 2; fi59++ {
                    v60 := GenS5{F0: (v3 ^ (v3 / (((v3 & (v3 * v3))) | 1))), F1: (-943 > -1)}
                    v61 := GenS4{F0: -882}
                    var v62 int32 = 0
                    _ = fi59
                    _ = v60
                    _ = v61
                    _ = v62
                }
                var v63 int8 = 8
                _ = fi57
                _ = v58
                _ = v63
            }
        } else {
            var v64 int = 0
            var v65 int16 = -233
            if len(v1) > 0 {
                v1[0] = "hello"
            }
            _ = v64
            _ = v65
        }
    } else {
        var v66 int16 = 804
        v67 := GenS3{F0: v2, F1: true}
        _ = v66
        _ = v67
    }
    var v68 string = v43
    if true {
        v43 = ""
    }
    fmt.Println(v5, v68)
    fmt.Println(v2)
    v68 = v68
    v1 = append(v1, ((((("" + "a") + (v43 + v43)) + "foo bar") + v43) + ((v68 + (v43 + "")) + v43)))
    v69 := GenS6{F0: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v68
    _ = v69
}
