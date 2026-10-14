package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 bool; F2 bool }
func genFn0(p0 uint32) int32 {
    var v6 string = "\n\t"
    if true {
        fmt.Println(p0)
        if (!(312 != 342)) {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v6, v6)
            case "\n\t":
                v7 := []int{0, 4294967295, 0}
                _ = v7
            case "":
                v8 := []int{9223372036854775807, 9223372036854775807}
                _ = v8
            }
        } else {
            fmt.Println(p0, p0)
        }
    }
    var v9 float64 = ((((0.1 / 3.14159) * 0.1) + 0.0) + 0.1)
    var v10 int32 = 648
    _ = p0
    _ = v10
    _ = v6
    _ = v9
    return (-(v10))
}
func genFn1() int16 {
    var v11 uint32 = 454
    v12 := GenS1{F0: (42 >= 0), F1: 0}
    v13 := GenS2{F0: 255, F1: -381}
    v14 := []int{32768}
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 32767
}

func main() {
    var v1 int = 313
    var v2 int16 = 1
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    var v5 int32 = -696
    var v15 uint16 = 255
    v16 := GenS3{F0: (709 < 2147483647), F1: 1}
    v17 := GenS3{F0: true, F1: -1}
    switch v1 {
    case 42:
        var v18 int = v1
        _ = v18
    case 2147483647:
        v19 := GenS2{F0: 1, F1: -1}
        _ = v19
    case 127:
        if (880 >= 1) {
            if false {
                switch v1 {
                case 866:
                    switch "foo bar" {
                    case "foo bar":
                        if false {
                            v20 := GenS4{F0: genFn1()}
                            var v21 int16 = v2
                            var v22 uint8 = 127
                            _ = v20
                            _ = v21
                            _ = v22
                        } else {
                            v23 := GenS3{F0: true, F1: 97}
                            _ = v23
                        }
                    }
                case 0:
                    switch 128 {
                    case 36:
                        v24 := []int{(v1 - (v1 & 9223372036854775807))}
                        _ = v24
                    case 0:
                        for fi25 := 0; fi25 < 9; fi25++ {
                            v26 := []int{int(v15), v1, (fi25 >> ((255) & 63)), v1}
                            _ = fi25
                            _ = v26
                        }
                    default:
                        switch ("foo bar" + ("hello" + ("" + "a"))) {
                        case "hello":
                            switch v1 {
                            case 1:
                                var v27 int32 = v5
                                _ = v27
                            case -854:
                                var v28 uint64 = 18446744073709551615
                                _ = v28
                            }
                        case "a":
                            if (1 >= 0) {
                                if true {
                                    var v29 int = (v1 >> ((65535) & 63))
                                    v30 := GenS1{F0: false, F1: 343}
                                    _ = v29
                                    _ = v30
                                }
                                var v31 uint32 = 2147483648
                                v15 = 906
                                _ = v31
                            } else {
                                if true {
                                    for fi32 := 0; fi32 < 20; fi32++ {
                                        if true {
                                            v15 = v15
                                            v15 += v15
                                        }
                                        var v33 int = (v1 << ((1) & 63))
                                        fmt.Println(v2)
                                        _ = fi32
                                        _ = v33
                                    }
                                    var v34 int16 = (v2 % ((v2) | 1))
                                    for fi35 := 0; fi35 < 16; fi35++ {
                                        var v36 int = v1
                                        for fi37 := 0; fi37 < 8; fi37++ {
                                            var v38 int8 = -107
                                            v39 := GenS1{F0: true, F1: 42}
                                            _ = fi37
                                            _ = v38
                                            _ = v39
                                        }
                                        _ = fi35
                                        _ = v36
                                    }
                                    _ = v34
                                }
                                var v40 int8 = 42
                                _ = v40
                            }
                        case "":
                            v41 := GenS2{F0: 255, F1: -357}
                            _ = v41
                        default:
                            for fi42 := 0; fi42 < 9; fi42++ {
                                var v43 float64 = (((1e10 - (((float64((1.5 + 0.0)) * float64((3.14159 + (1.5 + -2.25)))) + (((((1.5 + (3.14159 / 0.1)) / 1.0) + -2.25) + float64(((((-2.25 + (1e10 * (0.0 * 0.1))) + ((1.5 / (float32(float64(3.14159)) - 0.1)) * 1e10)) * (0.1 + ((0.0 * float32(1e10)) - 1e10))) / 1.0))) / (1e10 - (float64(float32(3.14159)) - ((1.5 * ((float64(1.5) / float64(1.5)) + 1.5)) / 1.0))))) - 1e10)) - 1e10) - 1.5)
                                v44 := GenS5{F0: 105}
                                for fi45 := 0; fi45 < 2; fi45++ {
                                    fmt.Println(v5, v5, v1)
                                    _ = fi45
                                }
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                        }
                    }
                case 9223372036854775807:
                    var v46 uint8 = uint8(v15)
                    _ = v46
                }
                var v47 string = "a"
                _ = v47
            } else {
                switch v1 {
                case 4294967295:
                    v15 = (v15 | v15)
                case 822:
                    switch "a" {
                    case "":
                        for fi48 := 0; fi48 < 14; fi48++ {
                            v49 := GenS2{F0: 42, F1: 1}
                            _ = fi48
                            _ = v49
                        }
                    case "hello":
                        switch (("\n\t" + ("" + "")) + ("a" + "a")) {
                        case "a":
                            v2 = -32768
                        case "\n\t":
                            var v50 int32 = (v5 + (v5 & -2147483648))
                            _ = v50
                        default:
                            for fi51 := 0; fi51 < 8; fi51++ {
                                v2 += (v2 % ((-32768) | 1))
                                if true {
                                    v52 := []int{v1, 32768, 32767}
                                    switch v1 {
                                    case 32768:
                                        if ((^(v2)) >= 0) {
                                            for fi53 := 0; fi53 < 19; fi53++ {
                                                v3["a"] = "foo bar"
                                                if len(v52) >= 2 {
                                                    v52 = v52[1:len(v52)]
                                                }
                                                switch "hello" {
                                                case "":
                                                    if true {
                                                        v54 := GenS3{F0: (int(v5) > (v1 - fi53)), F1: -63}
                                                        if false {
                                                            var v55 uint64 = 2
                                                            v15 = (^(v15))
                                                            if len(v52) > 0 {
                                                                v52[0] = (v1 & fi51)
                                                            }
                                                            _ = v55
                                                        }
                                                        _ = v54
                                                    }
                                                case "a":
                                                    var v56 string = "hello"
                                                    _ = v56
                                                case "hello":
                                                    _ = len(v52)
                                                }
                                                _ = fi53
                                            }
                                        }
                                    default:
                                        var v57 int8 = 100
                                        _ = v57
                                    }
                                    for fi58 := 0; fi58 < 8; fi58++ {
                                        v59 := v3["hello"]
                                        var v60 int32 = v5
                                        _ = fi58
                                        _ = v59
                                        _ = v60
                                    }
                                    _ = v52
                                } else {
                                    v61 := GenS6{F0: (v2 - (v2 / ((genFn1()) | 1))), F1: (103 < 0), F2: (-32768 == (v2 * 686))}
                                    _ = v61
                                }
                                v1 -= (v1 | 745)
                                _ = fi51
                            }
                        }
                    case "\n\t":
                        v62 := v4[("foo bar" + "a")]
                        _ = v62
                    }
                }
            }
        }
    }
    if (!false) {
        var v63 int64 = -495
        switch v1 {
        case -187:
            v5 -= genFn0(2)
        case 0:
            var v64 int32 = genFn0(214)
            _ = v64
        }
        _ = v63
    } else {
        v15 = (v15 - v15)
        var v65 int16 = genFn1()
        _ = v65
    }
    v15 += v15
    if false {
        v66 := v3[""]
        var v67 uint32 = 65535
        var v68 int16 = v2
        _ = v66
        _ = v67
        _ = v68
    }
    if (!true) {
        var v69 bool = false
        _ = v69
    } else {
        v5 = 67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
