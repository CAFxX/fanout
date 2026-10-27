package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 bool; F1 bool }
type GenS11 struct { F0 uint }
func genFn0() int16 {
    var v5 bool = true
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: v5, F1: -1}
        v8 := GenS2{F0: 2}
        var v9 int32 = -2147483648
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(v5)
    _ = v5
    return 9
}
func genFn1(p0 int64) int8 {
    if true {
        v10 := GenS3{F0: 1}
        var v11 uint64 = 42
        _ = v10
        _ = v11
    }
    var v12 int16 = 929
    var v13 uint16 = 32767
    _ = p0
    _ = v12
    _ = v13
    return -103
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int8 = -128
    v3 := []string{}
    var v4 uint = 111
    var v14 int32 = -103
    var v15 int32 = 0
    var v16 uint64 = 669
    v17 := GenS1{F0: false, F1: (v15 - (v14 % (((v14 >> ((v4) & 63))) | 1)))}
    var v18 int8 = (v2 * (^(v2)))
    v19 := GenS4{F0: 255, F1: 0}
    v20 := GenS4{F0: 0, F1: int16(v4)}
    var v21 string = "\n\t"
    switch int(v15) {
    case 42:
        fmt.Println(v18, v1, v21)
    default:
        if true {
            switch v21 {
            case "foo bar":
                var v22 int8 = genFn1(596)
                _ = v22
            case "":
                v18 = (v2 - (v2 / ((v18) | 1)))
            default:
                if ((true || (!((v16 << (((v4 + 80)) & 63)) >= v16))) && (!(!false))) {
                    v21 = ((v21 + v21) + v21)
                    var v23 bool = false
                    _ = v23
                }
            }
        }
    }
    var v24 uint32 = 931
    v15 -= v15
    v25 := GenS5{F0: v21, F1: v15}
    if (!false) {
        if true {
            var v26 string = ((v21 + v21) + v21)
            var v27 int16 = genFn0()
            if false {
                v28 := GenS6{F0: (v15 ^ (v14 % ((v14) | 1)))}
                if true {
                    if ((true && false) || ((true && ((502 <= 1) && (v27 == (v27 >> ((255) & 63))))) && (true && ((v15 & (v14 | (^(v15)))) != v14)))) {
                        v1 = uint64(v4)
                    }
                    var v29 uint32 = (v24 / (((v24 + v24)) | 1))
                    if true {
                        v30 := GenS6{F0: v15}
                        v26 = ""
                        var v31 float64 = (1.5 + 3.14159)
                        _ = v30
                        _ = v31
                    } else {
                        _ = len(v3)
                        var v32 uint = v4
                        _ = v32
                    }
                    _ = v29
                }
                v21 = "a"
                _ = v28
            }
            _ = v26
            _ = v27
        } else {
            var v33 uint8 = 42
            v15 = (v15 + (v15 & 2147483647))
            _ = v33
        }
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v21, v4)
        }
        switch -107 {
        case 128:
            if len(v3) > 0 {
                v3[0] = v21
            }
        case 1:
            var v34 int16 = -32768
            _ = v34
        default:
            var v35 float32 = -2.25
            _ = v35
        }
    } else {
        for fi36 := 0; fi36 < 19; fi36++ {
            v37 := GenS7{F0: 991}
            v38 := GenS8{F0: true, F1: 591}
            var v39 uint16 = 32767
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        fmt.Println(v1, v2, v14)
        if ((!false) || true) {
            var v40 int = int(v1)
            if (((!false) && ((437 <= 32767) || ((v14 < (v14 % ((-1) | 1))) && (true || false)))) || true) {
                if (!(v21 >= v21)) {
                    var v41 uint16 = 1
                    _ = len(v3)
                    _ = v41
                } else {
                    if (!true) {
                        v18 = genFn1(42)
                        switch (v40 / ((0) | 1)) {
                        case -118:
                            var v42 float32 = 0.0
                            _ = v42
                        case 128:
                            v3 = append(v3, "")
                        case 4294967295:
                            fmt.Println(v40, v14)
                        default:
                            v15 = v14
                        }
                    }
                    v43 := GenS2{F0: (^(v4))}
                    for fi44 := 0; fi44 < 15; fi44++ {
                        v15 += v15
                        for fi45 := 0; fi45 < 10; fi45++ {
                            v46 := GenS9{F0: ((v18 << ((42) & 63)) >= 127)}
                            _ = fi45
                            _ = v46
                        }
                        v15 = (v15 + v14)
                        _ = fi44
                    }
                    _ = v43
                }
                var v47 float64 = 3.14159
                v3 = append(v3, "\n\t")
                _ = v47
            } else {
                if true {
                    _ = len(v3)
                } else {
                    var v48 uint64 = 0
                    v48 = v16
                    _ = v48
                }
            }
            _ = v40
        }
    }
    if true {
        var v49 uint32 = (v24 >> (((v4 | (v4 >> ((v4) & 63)))) & 63))
        v50 := GenS10{F0: false, F1: false}
        var v51 int16 = genFn0()
        _ = v49
        _ = v50
        _ = v51
    }
    for fi52 := 0; fi52 < 11; fi52++ {
        v2 += (v2 * v2)
        var v53 string = (v21 + v21)
        switch fi52 {
        case -134:
            for fi54 := 0; fi54 < 5; fi54++ {
                var v55 string = v21
                if false {
                    if (v4 < v4) {
                        switch (fi52 * (fi52 << (((v4 / ((v4) | 1))) & 63))) {
                        case -709:
                            if true {
                                v56 := GenS11{F0: uint(v16)}
                                switch 2147483647 {
                                case 679:
                                    var v57 string = v21
                                    _ = v57
                                case -9223372036854775808:
                                    fmt.Println(v2, v16, fi54)
                                case 4294967295:
                                    _ = len(v3)
                                }
                                _ = v56
                            }
                        case 42:
                            v18 = (-(v2))
                        case -9223372036854775808:
                            switch (("hello" + ("a" + (v53 + v21))) + "") {
                            case "":
                                v18 = (v2 >> (((^(v4))) & 63))
                            default:
                                var v58 uint16 = 321
                                _ = v58
                            }
                        }
                    } else {
                        fmt.Println(v24, v24, v1)
                        var v59 int32 = v15
                        _ = v59
                    }
                    fmt.Println(v14, v21, v4)
                    v14 = (^(v15))
                }
                v60 := GenS10{F0: true, F1: (!(!(187 > 0)))}
                _ = fi54
                _ = v55
                _ = v60
            }
        case 0:
            for fi61 := 0; fi61 < 20; fi61++ {
                fmt.Println(v15, v15, fi52)
                v3 = append(v3, v21)
                _ = fi61
            }
        case 32768:
            v4 -= v4
        }
        _ = fi52
        _ = v53
    }
    switch -91 {
    case 128:
        if false {
            switch v21 {
            case "hello":
                for fi62 := 0; fi62 < 3; fi62++ {
                    v24 = v24
                    v3 = append(v3, v21)
                    _ = fi62
                }
            }
            for fi63 := 0; fi63 < 18; fi63++ {
                fmt.Println(v14, v15)
                var v64 int8 = (v2 % ((v2) | 1))
                if false {
                    if (v1 == (v16 - (^(v1)))) {
                        if (!(-9223372036854775808 <= -9223372036854775808)) {
                            switch "" {
                            case "":
                                var v65 float64 = float64((float32(((1e10 / 1.0) + (float64(1e10) * (0.1 / 1.0)))) / -2.25))
                                _ = v65
                            case "\n\t":
                                var v66 int32 = v14
                                _ = v66
                            default:
                                fmt.Println(v4, v1, v2)
                            }
                            for fi67 := 0; fi67 < 12; fi67++ {
                                var v68 float32 = (float32((float64(((1.5 + float32(-2.25)) - 0.0)) * 3.14159)) - 1.5)
                                v14 = v14
                                _ = fi67
                                _ = v68
                            }
                        }
                    }
                }
                _ = fi63
                _ = v64
            }
            fmt.Println(v18, v2)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
