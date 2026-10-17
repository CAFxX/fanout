package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 uint8; F2 string }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 uint32; F1 string }
type GenS9 struct { F0 uint64; F1 string; F2 uint64 }
type GenS10 struct { F0 bool; F1 uint32 }
type GenS11 struct { F0 int8; F1 int8 }
func genFn0(p0 int, p1 uint64, p2 int8) int32 {
    v4 := GenS1{F0: -604, F1: 255}
    var v5 bool = (!(2 < 357))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 661
}
func genFn1(p0 uint64, p1 uint64) uint {
    v6 := GenS1{F0: -9223372036854775808, F1: 127}
    v7 := GenS2{F0: 9223372036854775807, F1: 527}
    if (true || true) {
        v8 := GenS3{F0: 2}
        _ = v8
    }
    if true {
        if true {
            v9 := GenS2{F0: 9223372036854775807, F1: 346}
            _ = v9
        }
        var v10 int32 = 2147483647
        var v11 string = ((("a" + "\n\t") + "hello") + "hello")
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 42, F1: 186}
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    _ = v7
    return 65535
}

func main() {
    v1 := []int{73, 32767, 852}
    var v2 int32 = 762
    v3 := make(map[string]bool)
    for fi13 := 0; fi13 < 9; fi13++ {
        if true {
            v2 += -341
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(fi13, fi13, fi14)
                v1 = append(v1, (fi13 >> ((65535) & 63)))
                _ = fi14
            }
            switch "hello" {
            case "hello":
                v15 := GenS2{F0: -9223372036854775808, F1: 0}
                _ = v15
            case "a":
                var v16 uint = 1
                _ = v16
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = (-(fi13))
                }
            default:
                for fi17 := 0; fi17 < 6; fi17++ {
                    v2 += (v2 & v2)
                    v2 += (v2 >> ((4294967295) & 63))
                    _ = fi17
                }
            }
        } else {
            v2 = genFn0(32767, 479, -1)
            v18 := GenS1{F0: 2147483647, F1: 159}
            var v19 int32 = v2
            _ = v18
            _ = v19
        }
        if (((!((float64((0.1 * (1e10 + -2.25))) / 1.0) <= 1e10)) && false) && (!false)) {
            v2 = 0
            v2 += 2147483647
            if true {
                v2 = 2147483647
                fmt.Println(fi13, v2)
                switch 2147483647 {
                case 101:
                    var v20 int = (fi13 << ((genFn1(4294967295, 26)) & 63))
                    _ = v20
                case -296:
                    switch (fi13 % (((fi13 & fi13)) | 1)) {
                    case 127:
                        v3["hello"] = true
                    case 4294967295:
                        v21 := v3[("a" + "hello")]
                        _ = v21
                    case 751:
                        if false {
                            v2 = v2
                            var v22 float32 = 0.0
                            var v23 uint32 = 65535
                            _ = v22
                            _ = v23
                        } else {
                            fmt.Println(v2)
                            for fi24 := 0; fi24 < 16; fi24++ {
                                v25 := GenS4{F0: -2147483648, F1: 255}
                                v26 := GenS5{F0: (fi24 - (fi24 | (fi13 | (-(fi24)))))}
                                if len(v1) > 0 {
                                    v1[0] = fi13
                                }
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            v2 -= 0
                        }
                    }
                case -55:
                    v27 := GenS6{F0: -9223372036854775808, F1: 1, F2: ("\n\t" + ((("\n\t" + ("hello" + "foo bar")) + (("hello" + ("\n\t" + "foo bar")) + "hello")) + "a"))}
                    _ = v27
                default:
                    var v28 uint32 = 442
                    _ = v28
                }
            }
        } else {
            switch "a" {
            case "a":
                var v29 uint8 = 255
                _ = v29
            case "":
                v2 += (v2 % (((-(v2))) | 1))
            default:
                v2 -= 0
            }
            v2 -= 42
            v1 = append(v1, fi13)
        }
        var v30 uint8 = 29
        _ = fi13
        _ = v30
    }
    fmt.Println(v2)
    var v31 uint = 0
    v32 := GenS7{F0: 42, F1: -1}
    for fi33 := 0; fi33 < 17; fi33++ {
        v2 -= v2
        switch (("" + "hello") + "a") {
        case "hello":
            var v34 int8 = 1
            _ = v34
        }
        v35 := GenS8{F0: 306, F1: "\n\t"}
        _ = fi33
        _ = v35
    }
    v36 := GenS9{F0: 2, F1: "", F2: 599}
    switch 1 {
    case 4294967295:
        var v37 int8 = -16
        _ = v37
    case 2147483647:
        if false {
            var v38 float64 = (-2.25 - 0.1)
            _ = v38
        }
    default:
        if true {
            for fi39 := 0; fi39 < 11; fi39++ {
                var v40 uint8 = 24
                for fi41 := 0; fi41 < 19; fi41++ {
                    var v42 bool = false
                    fmt.Println(v2, v40)
                    _ = fi41
                    _ = v42
                }
                if (912 == 1) {
                    switch ("a" + "foo bar") {
                    case "foo bar":
                        v31 = (^(v31))
                    case "":
                        v43 := GenS10{F0: (!(1 < 0)), F1: 0}
                        _ = v43
                    }
                    if false {
                        for fi44 := 0; fi44 < 20; fi44++ {
                            if ((9223372036854775807 <= 9223372036854775807) && false) {
                                switch fi39 {
                                case -374:
                                    if (0.0 > 0.1) {
                                        var v45 int32 = int32(fi44)
                                        if (32767 > 695) {
                                            v40 = v40
                                        } else {
                                            v46 := GenS7{F0: 127, F1: -1}
                                            v47 := GenS10{F0: ((v40 >> ((2) & 63)) == (v40 + (v40 << ((v31) & 63)))), F1: 65535}
                                            var v48 uint64 = uint64(v40)
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        }
                                        switch ("\n\t" + "a") {
                                        case "\n\t":
                                            var v49 float32 = 1.5
                                            _ = v49
                                        }
                                        _ = v45
                                    } else {
                                        var v50 int16 = 385
                                        fmt.Println(v40, fi39)
                                        _ = v50
                                    }
                                default:
                                    fmt.Println(fi39, v31, fi44)
                                }
                                if true {
                                    var v51 int64 = -1
                                    _ = v51
                                } else {
                                    fmt.Println(fi39)
                                }
                                v52 := GenS5{F0: (fi44 * (fi39 >> (((v31 >> (((v31 >> ((1) & 63))) & 63))) & 63)))}
                                _ = v52
                            }
                            var v53 uint64 = 1
                            _ = fi44
                            _ = v53
                        }
                    } else {
                        fmt.Println(v31, v31)
                        v31 = v31
                        fmt.Println(v40, fi39)
                    }
                } else {
                    _ = len(v1)
                    _ = len(v1)
                    for fi54 := 0; fi54 < 16; fi54++ {
                        v31 += uint(v40)
                        _ = fi54
                    }
                }
                _ = fi39
                _ = v40
            }
            for fi55 := 0; fi55 < 17; fi55++ {
                switch -112 {
                case -9223372036854775808:
                    v31 -= (^(v31))
                case -1:
                    var v56 int32 = v2
                    _ = v56
                case 804:
                    v2 = (v2 & v2)
                }
                _ = fi55
            }
            for fi57 := 0; fi57 < 11; fi57++ {
                if len(v1) > 0 {
                    v1[0] = 127
                }
                _ = fi57
            }
        } else {
            var v58 string = ("a" + "hello")
            v59 := GenS11{F0: -93, F1: 127}
            switch 47 {
            case -329:
                _ = len(v3)
            default:
                var v60 int = 32768
                _ = v60
            }
            _ = v58
            _ = v59
        }
    }
    v31 = (v31 | (v31 >> (((v31 & v31)) & 63)))
    var v61 int8 = -89
    var v62 float64 = ((0.0 * 0.0) / (float64(0.1) - 0.0))
    var v63 int = -110
    var v64 uint16 = 42
    switch "hello" {
    case "hello":
        v2 += genFn0(v63, 4294967295, (-(v61)))
    case "":
        v31 -= 42
    }
    v3["a"] = false
    _ = len(v3)
    switch (v63 % ((-930) | 1)) {
    case 4294967295:
        v64 = (v64 / ((v64) | 1))
    case 32768:
        var v65 int8 = v61
        _ = v65
    case 846:
        for fi66 := 0; fi66 < 6; fi66++ {
            if false {
                for fi67 := 0; fi67 < 3; fi67++ {
                    var v68 int16 = int16(v61)
                    _ = fi67
                    _ = v68
                }
                for fi69 := 0; fi69 < 16; fi69++ {
                    v70 := GenS8{F0: 0, F1: ""}
                    if (v61 <= int8(v64)) {
                        v2 = (^(v2))
                        var v71 int = (fi69 * 127)
                        v2 = 0
                        _ = v71
                    }
                    v72 := GenS9{F0: 54, F1: "\n\t", F2: 815}
                    _ = fi69
                    _ = v70
                    _ = v72
                }
            } else {
                v73 := v3["a"]
                _ = v73
            }
            _ = fi66
        }
    default:
        var v74 int8 = v61
        _ = v74
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v61
    _ = v62
    _ = v63
    _ = v64
}
