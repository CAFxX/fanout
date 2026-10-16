package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint32; F1 string; F2 int8 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 int16; F1 uint32; F2 int8 }
func genFn0(p0 int32, p1 string, p2 int16) int64 {
    v5 := []int{32767, 9223372036854775807}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 int64, p1 uint8) int64 {
    p1 -= 1
    for fi6 := 0; fi6 < 12; fi6++ {
        p1 += (p1 & 187)
        _ = fi6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint32 = 695
        var v9 bool = (genFn0(int32(v8), "", -518) > int64(fi7))
        fmt.Println(v8, p0)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return (p0 | p0)
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"a", ""}
    var v3 int64 = 1
    v4 := []int{}
    switch ((("hello" + ("" + ((("\n\t" + "a") + (("hello" + "foo bar") + "")) + "\n\t"))) + ("\n\t" + ("hello" + "\n\t"))) + "foo bar") {
    case "hello":
        var v10 bool = false
        _ = v10
    case "\n\t":
        var v11 uint64 = 0
        _ = v11
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if (!true) {
        switch "" {
        case "":
            if (("" + (("hello" + "") + "a")) > "a") {
                if (289 != 42) {
                    v3 -= -662
                    fmt.Println(v3)
                    v3 = (v3 | 42)
                } else {
                    switch ("" + "foo bar") {
                    case "a":
                        if len(v4) > 0 {
                            v4[0] = 1
                        }
                    default:
                        v3 -= (-(v3))
                    }
                }
                if len(v4) > 0 {
                    v4[0] = int(v3)
                }
            }
        case "hello":
            switch "\n\t" {
            case "hello":
                if len(v4) > 0 {
                    v4[0] = 9223372036854775807
                }
            case "\n\t":
                if (-32768 <= 32767) {
                    var v12 uint8 = 255
                    switch 9223372036854775807 {
                    case -2147483648:
                        var v13 uint8 = (v12 / (((^(v12))) | 1))
                        _ = v13
                    default:
                        switch 1 {
                        case 32768:
                            var v14 int64 = v3
                            _ = v14
                        case 128:
                            v12 = (v12 ^ (v12 ^ v12))
                        case 785:
                            v15 := v1[-253]
                            _ = v15
                        }
                    }
                    _ = v12
                }
            }
        case "foo bar":
            if false {
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 int8 = -128
                    for fi18 := 0; fi18 < 11; fi18++ {
                        if true {
                            switch ("" + "foo bar") {
                            case "":
                                switch (("foo bar" + ("hello" + "foo bar")) + ("a" + "")) {
                                case "foo bar":
                                    if (!((v17 / ((1) | 1)) == 42)) {
                                        v19 := GenS1{F0: 982, F1: 1}
                                        delete(v1, (^(fi16)))
                                        _ = v19
                                    }
                                case "a":
                                    if (3.14159 <= ((((3.14159 * (float32((float64(0.1) - -2.25)) * -2.25)) * (0.1 + 1e10)) - (float32((0.0 + (float64(0.0) * ((1.5 * 0.1) + 1e10)))) * (1.5 / 3.14159))) - (float32((float64(-2.25) + -2.25)) / 1.0))) {
                                        if true {
                                            var v20 string = "hello"
                                            _ = v20
                                        }
                                        if (!((float32(1.5) - ((3.14159 * 1e10) * (float32(float64((((-2.25 - 0.0) * (0.0 / ((0.0 / 3.14159) + 1e10))) - float32((0.1 / 1.0))))) / 3.14159))) >= -2.25)) {
                                            if (3.14159 < (((0.0 + 1e10) - -2.25) - (1e10 - -2.25))) {
                                                fmt.Println(v17)
                                                if (!true) {
                                                    fmt.Println(fi18)
                                                    var v21 bool = (237 <= 65535)
                                                    switch (fi18 * 4294967295) {
                                                    case 127:
                                                        v21 = (42 > 4294967295)
                                                    case -53:
                                                        switch "hello" {
                                                        case "foo bar":
                                                            fmt.Println(v21)
                                                        default:
                                                            var v22 int32 = 42
                                                            _ = v22
                                                        }
                                                        v3 = (v3 >> ((4294967295) & 63))
                                                    default:
                                                        v23 := GenS2{F0: 65535}
                                                        _ = v23
                                                    }
                                                    _ = v21
                                                }
                                            }
                                        } else {
                                            fmt.Println(fi16, fi18)
                                            fmt.Println(v3)
                                            if false {
                                                fmt.Println(v17, fi16)
                                                switch (fi16 >> ((1) & 63)) {
                                                case 9223372036854775807:
                                                    v3 = v3
                                                case -1:
                                                    v3 = 0
                                                default:
                                                    var v24 int16 = -32768
                                                    _ = v24
                                                }
                                            } else {
                                                delete(v1, fi16)
                                            }
                                            v17 = -33
                                        }
                                    } else {
                                        v17 = v17
                                    }
                                default:
                                    var v25 bool = true
                                    v17 = v17
                                    _ = v25
                                }
                                var v26 uint16 = 788
                                _ = v26
                            case "hello":
                                var v27 int32 = 775
                                _ = v27
                            case "foo bar":
                                fmt.Println(fi18, v3)
                            default:
                                if (2 >= 42) {
                                    var v28 bool = (!false)
                                    _ = v28
                                } else {
                                    fmt.Println(v3, v17, v3)
                                    if ((v3 << ((787) & 63)) == (v3 ^ (v3 % (((^(v3))) | 1)))) {
                                        if true {
                                            var v29 int = fi16
                                            _ = v29
                                        } else {
                                            var v30 uint8 = 0
                                            _ = v30
                                        }
                                        switch (fi18 - (fi16 << ((255) & 63))) {
                                        case -78:
                                            v17 = int8(fi18)
                                        case 32767:
                                            if len(v4) > 0 {
                                                v4[0] = -1
                                            }
                                        case 32768:
                                            if len(v4) > 0 {
                                                v4[0] = fi16
                                            }
                                        }
                                    } else {
                                        var v31 float64 = (-2.25 * ((3.14159 * (0.0 + 1.5)) * 1.5))
                                        var v32 uint = 694
                                        _ = v31
                                        _ = v32
                                    }
                                    var v33 float32 = -2.25
                                    v34 := GenS3{F0: 0, F1: 4294967295, F2: (v3 + (v3 + (v3 / ((1) | 1))))}
                                    v17 += -1
                                    _ = v33
                                    _ = v34
                                }
                            }
                            v3 += int64(fi16)
                        }
                        _ = fi18
                    }
                    var v35 int16 = 193
                    _ = fi16
                    _ = v17
                    _ = v35
                }
                var v36 int8 = 42
                _ = v36
            } else {
                fmt.Println(v3)
            }
        }
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        v38 := GenS1{F0: 42, F1: 994}
        _ = fi37
        _ = v38
    }
    v39 := GenS3{F0: 430, F1: 986, F2: -9223372036854775808}
    for fi40 := 0; fi40 < 14; fi40++ {
        for fi41 := 0; fi41 < 10; fi41++ {
            var v42 int = fi40
            fmt.Println(v3, v42)
            _ = fi41
            _ = v42
        }
        var v43 uint32 = 193
        for fi44 := 0; fi44 < 15; fi44++ {
            v45 := GenS3{F0: 18446744073709551615, F1: 18446744073709551615, F2: (v3 & (v3 | (v3 >> ((255) & 63))))}
            var v46 uint32 = 1
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = fi40
        _ = v43
    }
    var v47 uint8 = 42
    for fi48 := 0; fi48 < 3; fi48++ {
        v49 := GenS4{F0: 1, F1: "a", F2: -128}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v3)
        _ = fi48
        _ = v49
    }
    v50 := v1[-326]
    v51 := GenS5{F0: 2147483647, F1: 127}
    switch 41 {
    case 9223372036854775807:
        v52 := GenS3{F0: 317, F1: 42, F2: (v3 | (v3 >> ((4294967295) & 63)))}
        _ = v52
    case 0:
        fmt.Println(v3)
    }
    v3 = v3
    if (42 >= 1) {
        v53 := GenS6{F0: v3, F1: 2}
        var v54 uint8 = v47
        v3 = v3
        _ = v53
        _ = v54
    }
    switch -2147483648 {
    case 9223372036854775807:
        var v55 int16 = -592
        _ = v55
    default:
        switch "hello" {
        case "":
            if v50 {
                var v56 bool = v50
                v56 = true
                _ = v56
            } else {
                switch "" {
                case "hello":
                    fmt.Println(v50)
                case "":
                    fmt.Println(v50, v3, v47)
                }
                v57 := GenS7{F0: 0, F1: 32767}
                _ = v57
            }
        }
    }
    for fi58 := 0; fi58 < 17; fi58++ {
        var v59 uint64 = 42
        v60 := GenS3{F0: 255, F1: 2, F2: genFn0(-542, (("hello" + "") + "\n\t"), 0)}
        _ = fi58
        _ = v59
        _ = v60
    }
    v50 = ((((((0.1 / 1.0) - 0.0) - -2.25) * 3.14159) + 1.5) < float64(3.14159))
    if false {
        var v61 int16 = 1
        switch "hello" {
        case "":
            v62 := GenS8{F0: (v61 + v61), F1: 2147483648, F2: 57}
            _ = v62
        }
        var v63 float32 = 0.1
        _ = v61
        _ = v63
    } else {
        var v64 string = "a"
        fmt.Println(v64)
        fmt.Println(v3)
        _ = v64
    }
    v1[991] = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v47
    _ = v50
    _ = v51
}
