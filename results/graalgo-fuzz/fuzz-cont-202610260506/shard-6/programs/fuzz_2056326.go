package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int64; F1 uint16; F2 int32 }
type GenS5 struct { F0 string; F1 int8; F2 int32 }
type GenS6 struct { F0 int32; F1 uint16 }
func genFn0(p0 string, p1 uint8, p2 int16) int {
    if ((-2147483648 < 42) || (2 >= 42)) {
        switch "foo bar" {
        case "foo bar":
            var v6 uint32 = 65535
            _ = v6
        case "\n\t":
            v7 := []int{}
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int8 = 27
        for fi10 := 0; fi10 < 17; fi10++ {
            p2 = (-(p2))
            v11 := GenS1{F0: false}
            v12 := GenS2{F0: 18446744073709551615}
            _ = fi10
            _ = v11
            _ = v12
        }
        fmt.Println(fi8)
        _ = fi8
        _ = v9
    }
    v13 := []int{-1}
    if true {
        if true {
            v14 := GenS1{F0: (42 > 0)}
            _ = v14
        } else {
            p2 -= p2
            var v15 int8 = 0
            _ = v15
        }
        v16 := GenS1{F0: (((1.5 - float32((((3.14159 + 1.5) - 0.0) - float64(3.14159)))) > (((0.1 * 0.1) - 0.0) / 1e10)) || true)}
        _ = v16
    } else {
        if (uint32(p1) != 1) {
            if len(v13) > 0 {
                v13[0] = -908
            }
        } else {
            _ = len(v13)
            p0 = p0
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS1{F0: true}
            _ = fi17
            _ = v18
        }
        _ = len(v13)
    }
    if false {
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := GenS3{F0: (p1 * p1), F1: -1}
            _ = len(v13)
            fmt.Println(p2)
            _ = fi19
            _ = v20
        }
    } else {
        fmt.Println(p2)
        if (0.0 < (((((1e10 / float64(1.5)) / -2.25) - 3.14159) - 3.14159) / 3.14159)) {
            var v21 int64 = -1
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 4294967295
}
func genFn1(p0 uint16) int16 {
    var v22 int8 = -6
    v23 := []int{}
    v24 := GenS1{F0: (0 > 320)}
    if true {
        v23 = append(v23, 274)
    } else {
        switch 124 {
        case -865:
            switch 2147483647 {
            case 2147483647:
                switch "" {
                case "a":
                    fmt.Println(v22, v22)
                case "\n\t":
                    v25 := GenS1{F0: false}
                    _ = v25
                case "foo bar":
                    var v26 string = ("a" + ("" + ((("a" + "") + "hello") + ("\n\t" + (("\n\t" + "hello") + "a")))))
                    _ = v26
                }
            case 32768:
                fmt.Println(v22, p0)
            default:
                v23 = append(v23, 4294967295)
            }
        case 32768:
            var v27 float64 = float64((-2.25 - float32(1.5)))
            _ = v27
        case 0:
            if false {
                if (true || (32767 > 0)) {
                    var v28 uint32 = 42
                    v29 := GenS3{F0: 1, F1: 1}
                    _ = v28
                    _ = v29
                }
                switch (("foo bar" + "hello") + "\n\t") {
                case "foo bar":
                    var v30 int64 = 1
                    _ = v30
                }
            }
        default:
            v31 := GenS4{F0: 991, F1: p0, F2: 2147483647}
            _ = v31
        }
        p0 = p0
    }
    var v32 int8 = v22
    _ = p0
    _ = v22
    _ = v23
    _ = v24
    _ = v32
    return 1
}

func main() {
    v1 := []int{32768, 1, 4294967295}
    v2 := make(map[int]string)
    var v3 uint8 = 154
    var v4 float64 = ((-2.25 * 1.5) + -2.25)
    var v5 uint16 = 397
    v4 -= v4
    if (!true) {
        v5 += (v5 & v5)
    }
    var v33 int8 = 42
    var v34 string = ("" + "a")
    v3 = v3
    var v35 uint8 = 255
    if (!(true && false)) {
        switch v34 {
        case "\n\t":
            v36 := v2[42]
            _ = v36
        case "foo bar":
            v34 = "a"
        default:
            var v37 int64 = 9223372036854775807
            _ = v37
        }
        var v38 int32 = 837
        _ = v38
    }
    switch -767 {
    case 0:
        for fi39 := 0; fi39 < 10; fi39++ {
            var v40 int = (fi39 + (fi39 ^ fi39))
            if (-2147483648 != -1) {
                switch fi39 {
                case 42:
                    switch v40 {
                    case -2147483648:
                        if false {
                            if (!((v35 != (v3 + 0)) || (1 < -23))) {
                                v4 -= v4
                                for fi41 := 0; fi41 < 15; fi41++ {
                                    var v42 float64 = 0.0
                                    if false {
                                        var v43 int32 = -147
                                        v40 = (fi39 + (fi41 + genFn0("a", (^(v3)), -449)))
                                        _ = v43
                                    } else {
                                        v5 -= (^(v5))
                                    }
                                    fmt.Println(fi41, v42)
                                    _ = fi41
                                    _ = v42
                                }
                            } else {
                                v5 += v5
                            }
                        }
                    case -308:
                        switch (-(fi39)) {
                        case -816:
                            for fi44 := 0; fi44 < 1; fi44++ {
                                v3 = 0
                                _ = fi44
                            }
                        case 497:
                            var v45 int16 = -749
                            _ = v45
                        case -2147483648:
                            for fi46 := 0; fi46 < 4; fi46++ {
                                if false {
                                    var v47 int32 = 1
                                    v4 = v4
                                    _ = v47
                                }
                                v35 -= (v3 * (v35 >> ((609) & 63)))
                                if true {
                                    delete(v2, (v40 >> ((4294967295) & 63)))
                                } else {
                                    switch v34 {
                                    case "foo bar":
                                        v33 = v33
                                    case "\n\t":
                                        var v48 int16 = 32767
                                        _ = v48
                                    case "hello":
                                        fmt.Println(fi39)
                                    }
                                    v33 += (v33 & v33)
                                }
                                _ = fi46
                            }
                        }
                    case 786:
                        for fi49 := 0; fi49 < 13; fi49++ {
                            v50 := GenS4{F0: 1, F1: v5, F2: -2147483648}
                            _ = fi49
                            _ = v50
                        }
                    default:
                        v51 := GenS5{F0: v34, F1: v33, F2: 1}
                        _ = v51
                    }
                }
            } else {
                var v52 uint = 4294967295
                _ = v52
            }
            switch (v40 ^ v40) {
            case -9223372036854775808:
                v53 := GenS6{F0: 974, F1: (^(v5))}
                _ = v53
            case -493:
                for fi54 := 0; fi54 < 7; fi54++ {
                    v5 = (v5 | v5)
                    v55 := v2[(v40 - genFn0(v34, (^(v35)), -1))]
                    v4 = v4
                    _ = fi54
                    _ = v55
                }
            case 32767:
                v33 = v33
            default:
                var v56 float32 = ((((3.14159 - 1e10) * (3.14159 + 3.14159)) * (1e10 - -2.25)) * (-2.25 * 1.5))
                _ = v56
            }
            _ = fi39
            _ = v40
        }
    default:
        var v57 uint16 = (v5 ^ (v5 | (v5 + v5)))
        _ = v57
    }
    v58 := GenS2{F0: 582}
    fmt.Println(v33, v33)
    v59 := GenS2{F0: 21}
    for fi60 := 0; fi60 < 3; fi60++ {
        v33 = 1
        _ = fi60
    }
    for fi61 := 0; fi61 < 15; fi61++ {
        var v62 int = 42
        _ = fi61
        _ = v62
    }
    v63 := v2[32767]
    if true {
        for fi64 := 0; fi64 < 5; fi64++ {
            var v65 bool = false
            _ = fi64
            _ = v65
        }
        fmt.Println(v63, v34, v33)
    }
    v4 = (float64(((float32(0.1) * -2.25) - (3.14159 * (float32((v4 / float64(float32((3.14159 * 1e10))))) - 0.1)))) - -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v63
}
