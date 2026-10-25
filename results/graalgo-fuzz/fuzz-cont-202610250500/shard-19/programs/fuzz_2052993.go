package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS5 struct { F0 int16; F1 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 string = (("foo bar" + ("foo bar" + ("\n\t" + ""))) + "hello")
    v3 := []uint32{2147483648, 65535, 0, 2}
    var v4 uint16 = 845
    var v5 int = 1
    var v6 int = v5
    v7 := GenS1{F0: false, F1: 92}
    if false {
        var v8 int8 = -103
        v6 -= (v6 * 9223372036854775807)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 17; fi9++ {
            v1 += 42
            _ = fi9
        }
    }
    v1 = v1
    switch v6 {
    case -9223372036854775808:
        switch (v2 + (v2 + (v2 + "a"))) {
        case "hello":
            var v10 uint32 = 420
            _ = v10
        case "":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            var v11 uint32 = 628
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 13; fi12++ {
            switch (v2 + v2) {
            case "\n\t":
                switch 127 {
                case -567:
                    v13 := GenS1{F0: true, F1: -708}
                    _ = v13
                case 127:
                    v6 -= (-(v6))
                default:
                    fmt.Println(v2, v2, v1)
                }
            case "hello":
                if true {
                    v14 := GenS2{F0: 2147483647, F1: v1}
                    if false {
                        var v15 uint16 = 1
                        _ = v15
                    }
                    _ = v14
                } else {
                    if true {
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                        var v16 int32 = 688
                        switch v6 {
                        case 0:
                            v17 := GenS2{F0: 809, F1: v1}
                            _ = v17
                        }
                        _ = v16
                    }
                    switch (v5 & v6) {
                    case -2147483648:
                        v18 := GenS3{F0: 435}
                        _ = v18
                    case 0:
                        var v19 int16 = -32768
                        _ = v19
                    case 1:
                        if false {
                            var v20 int8 = 0
                            var v21 uint32 = 452
                            _ = v20
                            _ = v21
                        } else {
                            v22 := GenS4{F0: 14, F1: 530, F2: 1}
                            var v23 int32 = 42
                            var v24 string = (((v2 + ((((v2 + v2) + v2) + v2) + ((v2 + "\n\t") + v2))) + ("" + "a")) + ("foo bar" + v2))
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    default:
                        if true {
                            v5 = v5
                            if false {
                                var v25 float64 = (1.5 * float64(3.14159))
                                if true {
                                    for fi26 := 0; fi26 < 11; fi26++ {
                                        var v27 bool = false
                                        switch v2 {
                                        case "foo bar":
                                            var v28 int64 = -9223372036854775808
                                            _ = v28
                                        case "":
                                            v3 = append(v3, 0)
                                        case "hello":
                                            var v29 uint8 = 209
                                            _ = v29
                                        default:
                                            if true {
                                                v2 = "a"
                                            }
                                        }
                                        _ = fi26
                                        _ = v27
                                    }
                                    var v30 uint64 = uint64(v4)
                                    for fi31 := 0; fi31 < 17; fi31++ {
                                        v4 = (v4 | v4)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = fi31
                                    }
                                    _ = v30
                                }
                                _ = v25
                            } else {
                                for fi32 := 0; fi32 < 19; fi32++ {
                                    var v33 float64 = (((float64((-2.25 * -2.25)) * -2.25) * ((1e10 - 0.0) + 1.5)) * (1e10 / (3.14159 + -2.25)))
                                    if false {
                                        v3 = append(v3, 936)
                                        v34 := GenS3{F0: -1}
                                        _ = v34
                                    }
                                    v35 := GenS3{F0: -608}
                                    _ = fi32
                                    _ = v33
                                    _ = v35
                                }
                                for fi36 := 0; fi36 < 8; fi36++ {
                                    var v37 uint = (v1 - (v1 | v1))
                                    _ = fi36
                                    _ = v37
                                }
                            }
                            switch (v2 + "") {
                            case "\n\t":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                    }
                }
            default:
                if false {
                    var v38 uint64 = 0
                    var v39 int16 = 32767
                    _ = v38
                    _ = v39
                } else {
                    switch (v6 * -818) {
                    case 1:
                        var v40 uint32 = 175
                        _ = v40
                    case 127:
                        var v41 bool = true
                        _ = v41
                    case 0:
                        var v42 int16 = -343
                        _ = v42
                    default:
                        var v43 float32 = 1.5
                        _ = v43
                    }
                    fmt.Println(v5, v2, fi12)
                }
            }
            for fi44 := 0; fi44 < 1; fi44++ {
                if true {
                    v6 += (fi44 & fi12)
                }
                var v45 float32 = (0.1 - 0.1)
                var v46 uint32 = 2147483647
                switch (((v2 + "foo bar") + "\n\t") + (v2 + "")) {
                case "foo bar":
                    var v47 int32 = -2147483648
                    _ = v47
                default:
                    _ = len(v3)
                }
                _ = fi44
                _ = v45
                _ = v46
            }
            if len(v3) > 0 {
                v3[0] = 768
            }
            _ = fi12
        }
    }
    var v48 int32 = -2147483648
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    v48 = (v48 - -413)
    var v49 int = v6
    v50 := GenS4{F0: 42, F1: 2, F2: 0}
    fmt.Println(v4, v49)
    var v51 bool = false
    var v52 float64 = 1.5
    var v53 string = v2
    var v54 uint32 = 1
    v3 = append(v3, (v54 << ((v1) & 63)))
    if v51 {
        v55 := GenS5{F0: 1000, F1: (v54 & (^(v54)))}
        fmt.Println(v4, v51)
        _ = v55
    }
    fmt.Println(v52, v5)
    if ((!v51) || (v54 <= v54)) {
        var v56 uint = (v1 % ((65535) | 1))
        _ = len(v3)
        _ = v56
    } else {
        var v57 int32 = v48
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
    _ = v7
}
