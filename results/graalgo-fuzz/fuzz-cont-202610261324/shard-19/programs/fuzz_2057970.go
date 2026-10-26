package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 int64; F1 int16 }
func genFn0(p0 int64, p1 uint16) uint {
    if false {
        fmt.Println(p1)
        var v6 uint64 = 2
        _ = v6
    }
    p1 = 0
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint64 = 4294967295
        var v9 int32 = 470
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    var v1 int32 = -1
    var v2 uint16 = 32767
    var v3 uint64 = 107
    var v4 string = "foo bar"
    v5 := []string{"hello"}
    var v10 int64 = 1
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS1{F0: v2, F1: 32767, F2: 3}
        v5 = append(v5, "\n\t")
        fmt.Println(v2, v4)
        _ = fi11
        _ = v12
    }
    var v13 uint = genFn0((-(v10)), (^(v2)))
    if false {
        fmt.Println(v13, v13, v3)
        v2 -= 255
        if ((v1 % ((-1) | 1)) > -1) {
            fmt.Println(v4, v3)
        }
    } else {
        for fi14 := 0; fi14 < 8; fi14++ {
            v3 = v3
            _ = fi14
        }
    }
    if false {
        if false {
            v15 := GenS2{F0: int(v10)}
            v16 := GenS1{F0: v2, F1: -2147483648, F2: uint32(v2)}
            var v17 int = 128
            _ = v15
            _ = v16
            _ = v17
        }
        switch int(v13) {
        case 32767:
            var v18 uint16 = uint16(v13)
            _ = v18
        case 42:
            _ = len(v5)
        case -9223372036854775808:
            v13 = (v13 & v13)
        }
        var v19 uint64 = (v3 | 1)
        _ = v19
    } else {
        v20 := GenS3{F0: 4294967295}
        _ = len(v5)
        var v21 float32 = (float32(1.5) + (1.5 - 0.1))
        _ = v20
        _ = v21
    }
    v13 -= (v13 % ((v13) | 1))
    switch 127 {
    case -2147483648:
        v4 = ("" + (("a" + ((("" + v4) + (v4 + (v4 + v4))) + v4)) + ((v4 + v4) + ((v4 + v4) + "\n\t"))))
    case 0:
        switch 127 {
        case 9223372036854775807:
            if (2147483647 > uint32(v13)) {
                for fi22 := 0; fi22 < 12; fi22++ {
                    v13 += 711
                    _ = fi22
                }
                for fi23 := 0; fi23 < 12; fi23++ {
                    v10 -= (v10 | (v10 * (v10 >> (((v13 / ((v13) | 1))) & 63))))
                    var v24 int16 = 42
                    _ = fi23
                    _ = v24
                }
            } else {
                var v25 uint8 = 0
                _ = v25
            }
        case 32768:
            v5 = append(v5, (((v4 + ("\n\t" + v4)) + "hello") + ("foo bar" + (("a" + v4) + v4))))
        case -9223372036854775808:
            for fi26 := 0; fi26 < 19; fi26++ {
                if true {
                    fmt.Println(v2)
                }
                v13 += 18446744073709551615
                _ = fi26
            }
        }
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 int32 = v1
        var v29 uint64 = (v3 << ((4294967295) & 63))
        switch fi27 {
        case 32768:
            var v30 float64 = (3.14159 * -2.25)
            _ = v30
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    v4 = (v4 + "a")
    if (int32(v13) == v1) {
        switch v4 {
        case "\n\t":
            switch "hello" {
            case "a":
                if true {
                    v3 = uint64(v2)
                    switch 42 {
                    case -697:
                        var v31 float32 = 1e10
                        _ = v31
                    default:
                        fmt.Println(v13, v4)
                    }
                } else {
                    var v32 string = v4
                    v33 := GenS1{F0: v2, F1: 362, F2: 42}
                    _ = v32
                    _ = v33
                }
            case "":
                var v34 string = v4
                _ = v34
            case "foo bar":
                switch int(v1) {
                case 127:
                    if false {
                        v10 = -1
                        _ = len(v5)
                    } else {
                        v13 += (v13 / ((183) | 1))
                        switch v4 {
                        case "a":
                            var v35 uint16 = v2
                            _ = v35
                        case "foo bar":
                            v36 := GenS1{F0: (v2 >> ((4294967295) & 63)), F1: -839, F2: 2}
                            _ = v36
                        default:
                            var v37 int = 768
                            _ = v37
                        }
                    }
                case 0:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case 2147483647:
                    v38 := GenS3{F0: 1}
                    _ = v38
                default:
                    for fi39 := 0; fi39 < 6; fi39++ {
                        v40 := GenS1{F0: uint16(v3), F1: -1, F2: 1}
                        _ = fi39
                        _ = v40
                    }
                }
            default:
                switch -97 {
                case 32767:
                    if true {
                        for fi41 := 0; fi41 < 8; fi41++ {
                            v10 = v10
                            _ = fi41
                        }
                    } else {
                        _ = len(v5)
                    }
                }
            }
        case "foo bar":
            var v42 uint8 = 255
            _ = v42
        }
        var v43 int16 = 42
        _ = v43
    }
    if true {
        switch 1 {
        case -2147483648:
            v13 = genFn0(v10, uint16(v3))
        }
        for fi44 := 0; fi44 < 4; fi44++ {
            if true {
                var v45 uint = v13
                var v46 string = ""
                var v47 uint16 = (^(v2))
                _ = v45
                _ = v46
                _ = v47
            }
            for fi48 := 0; fi48 < 4; fi48++ {
                var v49 uint16 = (v2 ^ 65535)
                _ = fi48
                _ = v49
            }
            v50 := GenS3{F0: v3}
            _ = fi44
            _ = v50
        }
    }
    if ((!(false && (1e10 <= (((((-2.25 + ((1e10 + ((3.14159 + ((float32((-2.25 + 0.1)) + ((0.1 + 1.5) - -2.25)) * float32(1e10))) + -2.25)) * float32(-2.25))) + ((0.1 * float32(3.14159)) * (1.5 / 1e10))) - 3.14159) * -2.25) + float32(3.14159))))) || (!true)) {
        v51 := GenS4{F0: v4}
        switch -730 {
        case 32768:
            v1 -= (v1 - (v1 + (v1 ^ v1)))
        }
        if false {
            v1 = 0
            if (((1.5 + (3.14159 + 3.14159)) + 0.0) >= 3.14159) {
                for fi52 := 0; fi52 < 4; fi52++ {
                    v53 := GenS5{F0: v13, F1: "foo bar"}
                    for fi54 := 0; fi54 < 12; fi54++ {
                        var v55 int8 = 127
                        v5 = append(v5, "foo bar")
                        _ = fi54
                        _ = v55
                    }
                    _ = fi52
                    _ = v53
                }
            } else {
                v10 += 834
            }
            switch -2147483648 {
            case 2147483647:
                if len(v5) > 0 {
                    v5[0] = "foo bar"
                }
            case -360:
                v56 := GenS6{F0: uint32(v3), F1: 137}
                _ = v56
            default:
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
            }
        } else {
            var v57 float64 = float64(float32((3.14159 + 3.14159)))
            _ = v57
        }
        _ = v51
    } else {
        switch 1 {
        case -787:
            v58 := GenS7{F0: 0, F1: int16(v1)}
            _ = v58
        case 9223372036854775807:
            _ = len(v5)
        case -1:
            var v59 int64 = 604
            _ = v59
        default:
            v1 += 1
        }
        var v60 uint = v13
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
