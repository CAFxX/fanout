package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 string; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 int64; F1 int16; F2 uint32 }
type GenS6 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16; F1 bool; F2 int8 }
type GenS9 struct { F0 int; F1 uint64; F2 int16 }
type GenS10 struct { F0 uint8; F1 int8 }
type GenS11 struct { F0 uint }
func genFn0(p0 string, p1 uint8, p2 int64) int8 {
    v6 := []int{4294967295, 908, 32767, 127}
    var v7 uint32 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []string{"hello", "a", "a", "hello"}
    var v2 int64 = 42
    v3 := []string{}
    var v4 float32 = (float32(3.14159) + 3.14159)
    var v5 uint32 = 1
    var v8 int8 = -1
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    var v9 int = 2147483647
    v2 = (v2 ^ (v2 + v2))
    for fi10 := 0; fi10 < 15; fi10++ {
        switch "a" {
        case "":
            v11 := GenS1{F0: (fi10 / ((128) | 1)), F1: "\n\t"}
            _ = v11
        case "hello":
            var v12 int = (fi10 & (fi10 - v9))
            _ = v12
        case "\n\t":
            fmt.Println(v2, v4, fi10)
        default:
            if (!(!false)) {
                if (false || true) {
                    v4 += ((3.14159 * 3.14159) + v4)
                    _ = len(v1)
                    var v13 uint = 565
                    _ = v13
                } else {
                    switch 139 {
                    case 9223372036854775807:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 2147483647:
                        v8 += (v8 | int8(fi10))
                    case 253:
                        var v14 string = "\n\t"
                        _ = v14
                    default:
                        for fi15 := 0; fi15 < 20; fi15++ {
                            switch -9223372036854775808 {
                            case 180:
                                var v16 int64 = v2
                                _ = v16
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (2 >= 440) {
                                v17 := GenS1{F0: (fi15 << ((65535) & 63)), F1: "a"}
                                _ = v17
                            }
                            var v18 float64 = float64(-2.25)
                            if true {
                                v9 -= -2147483648
                                var v19 int = (fi15 % (((fi10 + 4294967295)) | 1))
                                var v20 uint32 = v5
                                switch "foo bar" {
                                case "\n\t":
                                    if false {
                                        v21 := GenS2{F0: (("hello" + "") + "hello"), F1: (v20 & 555), F2: 1}
                                        var v22 float32 = 0.1
                                        v20 -= (v5 ^ (v20 | 2))
                                        _ = v21
                                        _ = v22
                                    }
                                case "":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case "foo bar":
                                    fmt.Println(v4, v19)
                                }
                                _ = len(v3)
                                _ = v19
                                _ = v20
                            } else {
                                var v23 int32 = 0
                                _ = v23
                            }
                            _ = fi15
                            _ = v18
                        }
                    }
                    var v24 int32 = 401
                    if (fi10 <= (v9 - -492)) {
                        if true {
                            var v25 uint32 = 2147483648
                            for fi26 := 0; fi26 < 2; fi26++ {
                                v27 := GenS3{F0: 9223372036854775807, F1: 0}
                                _ = fi26
                                _ = v27
                            }
                            fmt.Println(v2)
                            _ = v25
                        }
                        switch (v9 % ((4294967295) | 1)) {
                        case -449:
                            if false {
                                var v28 string = "\n\t"
                                _ = v28
                            } else {
                                if (1000 >= 0) {
                                    if (!(-2.25 >= ((-2.25 + 0.1) / 1e10))) {
                                        v4 += -2.25
                                    } else {
                                        v29 := GenS2{F0: "a", F1: (v5 >> ((4294967295) & 63)), F2: 255}
                                        var v30 int32 = -463
                                        _ = v29
                                        _ = v30
                                    }
                                    var v31 int8 = 20
                                    _ = v31
                                } else {
                                    if false {
                                        _ = len(v3)
                                        v8 = genFn0("foo bar", 122, (-(v2)))
                                        v5 = v5
                                    }
                                    for fi32 := 0; fi32 < 19; fi32++ {
                                        var v33 int32 = (v24 | (v24 * (^(v24))))
                                        var v34 int = fi32
                                        fmt.Println(v24, v34)
                                        _ = fi32
                                        _ = v33
                                        _ = v34
                                    }
                                }
                                if true {
                                    for fi35 := 0; fi35 < 17; fi35++ {
                                        var v36 uint8 = 42
                                        v37 := GenS4{F0: (v5 * (v5 >> ((42) & 63))), F1: 0}
                                        _ = fi35
                                        _ = v36
                                        _ = v37
                                    }
                                    v38 := GenS4{F0: 42, F1: 42}
                                    if true {
                                        v39 := GenS5{F0: 9223372036854775807, F1: -32768, F2: 4294967295}
                                        fmt.Println(v8, v24, v24)
                                        _ = v39
                                    }
                                    _ = v38
                                }
                                fmt.Println(v4)
                            }
                        case 32767:
                            v9 += (fi10 ^ int(v5))
                        }
                        for fi40 := 0; fi40 < 7; fi40++ {
                            fmt.Println(v24, fi10)
                            _ = fi40
                        }
                    }
                    _ = v24
                }
                var v41 uint8 = 150
                v41 = v41
                _ = v41
            }
        }
        switch v9 {
        case 127:
            if (v9 >= (v9 - (fi10 & -253))) {
                v42 := GenS6{F0: (v8 >> ((967) & 63)), F1: 2147483647, F2: 255}
                _ = v42
            } else {
                v8 += v8
                var v43 float64 = 0.0
                v44 := GenS1{F0: (^(v9)), F1: "a"}
                _ = v43
                _ = v44
            }
        default:
            if false {
                _ = len(v1)
                v9 = 128
            }
        }
        _ = fi10
    }
    var v45 uint16 = 255
    if true {
        v46 := GenS6{F0: v8, F1: v5, F2: 1}
        for fi47 := 0; fi47 < 12; fi47++ {
            v8 += v8
            _ = fi47
        }
        switch ("a" + "foo bar") {
        case "a":
            for fi48 := 0; fi48 < 17; fi48++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi49 := 0; fi49 < 19; fi49++ {
                    if len(v1) > 0 {
                        v1[0] = ("a" + "hello")
                    }
                    var v50 int8 = (v8 % ((-1) | 1))
                    _ = fi49
                    _ = v50
                }
                var v51 float32 = -2.25
                _ = fi48
                _ = v51
            }
        case "foo bar":
            var v52 int16 = -32768
            _ = v52
        case "hello":
            v5 = 1
        }
        _ = v46
    } else {
        if false {
            if (-2.25 >= (1.5 + ((0.0 - 1.5) / 1.0))) {
                v9 += -597
                if true {
                    if (!(((v8 + int8(v45)) == int8(v2)) || true)) {
                        v5 = 2
                        fmt.Println(v9, v9, v5)
                        v8 -= (v8 << ((18446744073709551615) & 63))
                    } else {
                        if (!(!("hello" != (("\n\t" + "hello") + "foo bar")))) {
                            var v53 int32 = 0
                            fmt.Println(v9, v8)
                            var v54 string = ("" + "")
                            _ = v53
                            _ = v54
                        }
                        switch v9 {
                        case 32767:
                            var v55 float64 = 3.14159
                            _ = v55
                        case -1:
                            v8 -= genFn0("a", 42, (v2 << ((580) & 63)))
                        case 2147483647:
                            switch (v9 * v9) {
                            case -39:
                                var v56 uint = 18446744073709551615
                                _ = v56
                            case 4294967295:
                                if (!true) {
                                    v57 := GenS7{F0: v5}
                                    _ = v57
                                }
                            default:
                                v58 := GenS8{F0: 0, F1: false, F2: genFn0("\n\t", 172, 42)}
                                _ = v58
                            }
                        }
                        v9 -= -459
                    }
                }
            }
        }
        switch (^(v9)) {
        case 32767:
            var v59 int = (-(v9))
            _ = v59
        case -2147483648:
            for fi60 := 0; fi60 < 10; fi60++ {
                for fi61 := 0; fi61 < 19; fi61++ {
                    var v62 float64 = float64(float32(0.0))
                    var v63 uint8 = 42
                    if (!((1000 != -249) && false)) {
                        var v64 int64 = -40
                        v65 := GenS9{F0: (v9 / ((32767) | 1)), F1: 619, F2: 32767}
                        v66 := GenS10{F0: (v63 % (((v63 << ((4294967295) & 63))) | 1)), F1: genFn0("a", v63, (-(v2)))}
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                    _ = fi61
                    _ = v62
                    _ = v63
                }
                v3 = append(v3, "a")
                var v67 string = (("\n\t" + "") + "hello")
                _ = fi60
                _ = v67
            }
        case 715:
            v68 := GenS7{F0: v5}
            _ = v68
        }
    }
    switch (v9 % (((v9 ^ (^(v9)))) | 1)) {
    case 32767:
        v45 += (^(v45))
    default:
        v69 := GenS11{F0: 561}
        _ = v69
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v8
    _ = v9
}
