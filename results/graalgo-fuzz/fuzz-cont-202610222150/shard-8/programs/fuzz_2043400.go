package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 int16; F1 int16; F2 string }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 uint8; F1 uint64 }
type GenS11 struct { F0 uint8 }
func genFn0(p0 int) uint8 {
    switch "a" {
    case "\n\t":
        switch p0 {
        case -603:
            var v6 uint64 = 91
            _ = v6
        case 0:
            v7 := GenS1{F0: 2147483648}
            _ = v7
        case 42:
            switch (p0 | p0) {
            case 32767:
                var v8 float32 = 1.5
                _ = v8
            case 412:
                p0 = 2147483647
            }
        }
    case "hello":
        p0 -= 128
    default:
        v9 := []int{}
        _ = v9
    }
    var v10 float32 = 3.14159
    fmt.Println(v10)
    _ = p0
    _ = v10
    return 2
}
func genFn1(p0 uint32, p1 uint64, p2 int32) int16 {
    fmt.Println(p1)
    v11 := []int{2147483647, 1, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 32767
}

func main() {
    var v1 bool = true
    var v2 uint16 = 822
    var v3 int = 128
    var v4 uint32 = 572
    v5 := make(map[string]uint32)
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int16 = 32767
        if (!(((!true) && v1) || false)) {
            switch ("foo bar" + ("a" + ("\n\t" + "hello"))) {
            case "foo bar":
                fmt.Println(v1, v1, v1)
            case "a":
                v14 := GenS1{F0: v4}
                _ = v14
            case "\n\t":
                switch 283 {
                case 42:
                    var v15 int32 = -1
                    _ = v15
                case 32768:
                    v1 = (-125 != 97)
                case 32767:
                    if (1 <= 1) {
                        for fi16 := 0; fi16 < 7; fi16++ {
                            v4 = v4
                            switch "hello" {
                            case "a":
                                v17 := GenS1{F0: (^(v4))}
                                _ = v17
                            case "\n\t":
                                var v18 int64 = -448
                                _ = v18
                            case "":
                                v19 := GenS2{F0: v4}
                                _ = v19
                            default:
                                v1 = (0 != 4294967295)
                            }
                            _ = fi16
                        }
                        if (!(false && v1)) {
                            var v20 float32 = 1.5
                            _ = v20
                        }
                    } else {
                        if v1 {
                            v21 := v5["a"]
                            fmt.Println(v13, v2)
                            _ = v21
                        }
                        fmt.Println(v3, v3, v13)
                        var v22 uint8 = genFn0(0)
                        _ = v22
                    }
                }
            default:
                v1 = false
            }
        } else {
            if ("\n\t" > "") {
                var v23 uint64 = 168
                v24 := GenS3{F0: 65535}
                if v1 {
                    if v1 {
                        v25 := []int{fi12, fi12, (v3 - fi12)}
                        v26 := GenS2{F0: v4}
                        v13 = genFn1(2147483648, 42, -1)
                        _ = v25
                        _ = v26
                    } else {
                        v1 = (-914 > 1)
                    }
                } else {
                    v13 = v13
                    switch (("foo bar" + ("foo bar" + (("" + (("hello" + "a") + "foo bar")) + ""))) + "hello") {
                    case "\n\t":
                        v27 := GenS3{F0: 18446744073709551615}
                        _ = v27
                    case "hello":
                        v13 = (v13 ^ 0)
                    case "foo bar":
                        _ = len(v5)
                    default:
                        if (!(42 > 241)) {
                            var v28 int32 = -1
                            _ = v28
                        }
                    }
                    if (0.1 > 1.5) {
                        v29 := []int{v3, (v3 ^ -2147483648), 286, fi12}
                        _ = v29
                    } else {
                        var v30 int = 127
                        v31 := GenS1{F0: v4}
                        _ = v30
                        _ = v31
                    }
                }
                _ = v23
                _ = v24
            } else {
                var v32 uint16 = (^(v2))
                switch "" {
                case "a":
                    if (18446744073709551615 == 0) {
                        var v33 uint8 = 127
                        _ = v33
                    }
                case "hello":
                    if v1 {
                        fmt.Println(v3, v3)
                        var v34 uint32 = (v4 | (v4 >> ((0) & 63)))
                        _ = v34
                    } else {
                        var v35 int = (v3 / (((v3 | (fi12 >> ((65535) & 63)))) | 1))
                        _ = v35
                    }
                case "\n\t":
                    v36 := GenS4{F0: 42, F1: -755}
                    _ = v36
                default:
                    if true {
                        v32 -= (^(v2))
                    }
                }
                var v37 uint = 4294967295
                _ = v32
                _ = v37
            }
            if (v1 && (!(!(!true)))) {
                var v38 uint = 42
                v39 := GenS1{F0: 42}
                v40 := GenS5{F0: -724}
                _ = v38
                _ = v39
                _ = v40
            }
        }
        if v1 {
            var v41 bool = true
            if true {
                if (-2.25 <= (0.1 + ((0.1 - float32((1.5 * 1e10))) + (3.14159 - 1e10)))) {
                    switch (v3 ^ (v3 - (v3 | 127))) {
                    case 4294967295:
                        var v42 string = ""
                        _ = v42
                    case -2147483648:
                        v13 += genFn1(v4, 42, -518)
                    default:
                        var v43 uint8 = 59
                        _ = v43
                    }
                    switch v3 {
                    case 0:
                        for fi44 := 0; fi44 < 7; fi44++ {
                            v41 = v41
                            if true {
                                v13 -= (v13 | 42)
                                v45 := GenS6{F0: "foo bar", F1: 18446744073709551615}
                                _ = v45
                            } else {
                                var v46 float32 = 3.14159
                                _ = v46
                            }
                            _ = fi44
                        }
                    case 6:
                        switch (fi12 >> ((2) & 63)) {
                        case 0:
                            var v47 uint8 = uint8(v4)
                            _ = v47
                        }
                    case 9223372036854775807:
                        fmt.Println(v1, v3)
                    default:
                        fmt.Println(v4, v1, v4)
                    }
                }
                v48 := []int{v3, 701, v3}
                _ = v48
            }
            for fi49 := 0; fi49 < 13; fi49++ {
                v4 = (v4 - (v4 << ((983) & 63)))
                v50 := GenS2{F0: v4}
                var v51 int64 = -1
                _ = fi49
                _ = v50
                _ = v51
            }
            _ = v41
        } else {
            if (!(4294967295 > 0)) {
                var v52 int64 = -342
                _ = v52
            } else {
                v1 = true
                delete(v5, (("" + "a") + "foo bar"))
                var v53 float64 = 0.1
                _ = v53
            }
            var v54 bool = true
            var v55 int16 = v13
            _ = v54
            _ = v55
        }
        _ = fi12
        _ = v13
    }
    if v1 {
        v3 += v3
    } else {
        v56 := GenS6{F0: "a", F1: 760}
        if (9223372036854775807 <= -1) {
            v2 -= 0
            switch (("" + "foo bar") + "a") {
            case "a":
                if false {
                    v4 -= v4
                    v57 := GenS7{F0: 42, F1: -524, F2: ("a" + "a")}
                    switch ("foo bar" + "\n\t") {
                    case "a":
                        v58 := GenS8{F0: -1}
                        _ = v58
                    case "hello":
                        v2 = (v2 >> ((42) & 63))
                    case "foo bar":
                        fmt.Println(v4, v2, v4)
                    default:
                        var v59 int16 = 561
                        _ = v59
                    }
                    _ = v57
                } else {
                    for fi60 := 0; fi60 < 19; fi60++ {
                        v3 = (-(fi60))
                        _ = fi60
                    }
                    var v61 uint32 = v4
                    var v62 uint16 = (^(v2))
                    _ = v61
                    _ = v62
                }
            default:
                v63 := GenS6{F0: ("foo bar" + "\n\t"), F1: 164}
                _ = v63
            }
        } else {
            for fi64 := 0; fi64 < 3; fi64++ {
                for fi65 := 0; fi65 < 14; fi65++ {
                    _ = len(v5)
                    _ = fi65
                }
                _ = len(v5)
                for fi66 := 0; fi66 < 12; fi66++ {
                    v67 := GenS1{F0: (v4 / (((v4 & 2147483647)) | 1))}
                    fmt.Println(v2, fi64, v3)
                    _ = fi66
                    _ = v67
                }
                _ = fi64
            }
            var v68 int = v3
            _ = v68
        }
        _ = v56
    }
    if (470 == 590) {
        var v69 int8 = -128
        v4 -= v4
        _ = v69
    }
    var v70 int8 = 0
    for fi71 := 0; fi71 < 12; fi71++ {
        if v1 {
            fmt.Println(v70)
        } else {
            switch fi71 {
            case -9223372036854775808:
                v72 := []int{128, -2147483648, -1}
                _ = v72
            case 4294967295:
                v73 := []int{(^(v3)), (v3 << ((688) & 63))}
                _ = v73
            case 32767:
                if v1 {
                    if v1 {
                        fmt.Println(v2)
                    }
                    if v1 {
                        if (v1 || (float32(1.5) != (((0.1 + -2.25) * ((1e10 / 1.0) / 1.5)) - (1e10 - (1.5 + ((1e10 / (3.14159 + 1e10)) - (0.1 - 0.0))))))) {
                            for fi74 := 0; fi74 < 11; fi74++ {
                                fmt.Println(v1, fi74)
                                _ = fi74
                            }
                            var v75 int = (fi71 >> ((67) & 63))
                            v2 = 42
                            _ = v75
                        } else {
                            if true {
                                v3 += (fi71 - v3)
                                v1 = (((1 != 87) || v1) && (1 > 0))
                            }
                        }
                        fmt.Println(v4, v3)
                    }
                    switch (^(fi71)) {
                    case -351:
                        if v1 {
                            v2 += (v2 - 0)
                            v1 = true
                            var v76 int = (-(v3))
                            _ = v76
                        }
                    case 1:
                        v77 := GenS9{F0: (v2 % ((v2) | 1)), F1: (59 != 18446744073709551615)}
                        _ = v77
                    default:
                        v78 := []int{42, (^(fi71))}
                        _ = v78
                    }
                }
            }
            if v1 {
                var v79 float32 = (1e10 - 0.1)
                _ = v79
            } else {
                v4 = 825
                v80 := GenS10{F0: 127, F1: 84}
                var v81 uint8 = 255
                _ = v80
                _ = v81
            }
        }
        _ = fi71
    }
    v1 = false
    var v82 float64 = float64(0.0)
    v70 = v70
    v1 = v1
    v83 := GenS11{F0: 63}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v70:", v70)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v82
    _ = v83
}
