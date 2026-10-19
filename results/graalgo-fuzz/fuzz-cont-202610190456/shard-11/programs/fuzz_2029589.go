package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int8; F1 string }
func genFn0(p0 uint32) uint8 {
    p0 = p0
    for fi7 := 0; fi7 < 12; fi7++ {
        p0 = (p0 << ((1) & 63))
        _ = fi7
    }
    if false {
        p0 = p0
        switch -9223372036854775808 {
        case 32768:
            v8 := GenS1{F0: -742}
            _ = v8
        default:
            fmt.Println(p0)
        }
    }
    _ = p0
    return 42
}
func genFn1(p0 uint64, p1 int16) bool {
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint8 = 127
        v11 := GenS2{F0: (fi9 / (((fi9 >> ((42) & 63))) | 1)), F1: 0}
        var v12 int32 = -2147483648
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        fmt.Println(fi13)
        v14 := []int{(-(fi13))}
        _ = fi13
        _ = v14
    }
    var v15 float64 = (float64(0.0) / -2.25)
    _ = p0
    _ = p1
    _ = v15
    return true
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    v3 := []int{-1, -2147483648, -2147483648}
    var v4 int16 = 32767
    v5 := []uint32{42, 2147483648}
    var v6 uint32 = 2
    fmt.Println(v4, v4)
    var v16 uint = 1
    fmt.Println(v16, v4, v16)
    switch "hello" {
    case "":
        var v17 uint = (^(v16))
        _ = v17
    case "\n\t":
        for fi18 := 0; fi18 < 15; fi18++ {
            v16 += 704
            for fi19 := 0; fi19 < 17; fi19++ {
                v5 = append(v5, (^(v6)))
                v20 := v2[(-(fi19))]
                switch (fi18 - -1) {
                case -2147483648:
                    v21 := GenS2{F0: fi19, F1: v20}
                    _ = v21
                case 702:
                    if (!(false && true)) {
                        var v22 float32 = 1.5
                        _ = v22
                    } else {
                        fmt.Println(fi19, v16)
                    }
                }
                switch ("" + ("" + "")) {
                case "foo bar":
                    if (3.14159 > ((((3.14159 * -2.25) / 1e10) * 0.0) / 1.5)) {
                        var v23 uint32 = (v6 % (((v20 << ((42) & 63))) | 1))
                        switch (fi18 ^ (fi18 * 4294967295)) {
                        case 128:
                            switch (fi19 / ((fi18) | 1)) {
                            case -1:
                                v24 := GenS2{F0: (fi19 * (fi18 << ((v16) & 63))), F1: v6}
                                v6 -= uint32(v16)
                                if (870 == 750) {
                                    _ = len(v2)
                                    fmt.Println(v6)
                                }
                                if false {
                                    v25 := v1[((("" + "\n\t") + "foo bar") + "hello")]
                                    if ((0 > 42) || (!false)) {
                                        v25 = v25
                                        var v26 uint16 = 163
                                        if true {
                                            v23 = uint32(v26)
                                        } else {
                                            if false {
                                                fmt.Println(v4)
                                                delete(v1, "a")
                                                if true {
                                                    fmt.Println(v25)
                                                    fmt.Println(v25, v4, v6)
                                                    switch "a" {
                                                    case "hello":
                                                        v25 = (^(v20))
                                                    case "":
                                                        var v27 int32 = 2147483647
                                                        _ = v27
                                                    }
                                                }
                                                v26 -= v26
                                            } else {
                                                v4 = int16(v26)
                                                var v28 uint8 = genFn0(v6)
                                                fmt.Println(v26, v20, fi18)
                                                if (true && (!true)) {
                                                    v3 = append(v3, fi18)
                                                    var v29 string = ("" + "hello")
                                                    if (!(-2.25 > 0.0)) {
                                                        v16 += v16
                                                        if (v16 > (v16 * uint(v28))) {
                                                            var v30 string = v29
                                                            _ = v30
                                                        } else {
                                                            v31 := GenS2{F0: fi18, F1: (v23 % (((v20 - uint32(v16))) | 1))}
                                                            _ = v31
                                                        }
                                                    }
                                                    _ = v29
                                                } else {
                                                    switch (fi19 << ((v16) & 63)) {
                                                    case -1:
                                                        fmt.Println(fi19, fi18)
                                                    case 63:
                                                        fmt.Println(v20, v4)
                                                    default:
                                                        if (((((3.14159 - 1e10) + 3.14159) * ((3.14159 / (-2.25 * 1.5)) * 0.1)) / (-2.25 / 1.0)) == float64(1.5)) {
                                                            v20 -= 2
                                                            v32 := GenS3{F0: (!(!(false && false))), F1: v20}
                                                            _ = v32
                                                        }
                                                        delete(v2, (fi19 / ((-1) | 1)))
                                                        var v33 uint16 = (v26 % (((v26 ^ uint16(v16))) | 1))
                                                        _ = v33
                                                    }
                                                    var v34 float64 = 1.5
                                                    v28 = genFn0(v20)
                                                    _ = v34
                                                }
                                                switch fi18 {
                                                case 32767:
                                                    var v35 string = (("\n\t" + ("a" + (((("\n\t" + ("hello" + "a")) + ("" + ("\n\t" + "hello"))) + "foo bar") + ("foo bar" + "\n\t")))) + "hello")
                                                    _ = v35
                                                case -2147483648:
                                                    if (!false) {
                                                        var v36 uint = (v16 >> (((^(v16))) & 63))
                                                        _ = v36
                                                    }
                                                case 0:
                                                    v20 -= 0
                                                default:
                                                    switch (fi19 >> (((v16 ^ v16)) & 63)) {
                                                    case -9223372036854775808:
                                                        _ = len(v3)
                                                    }
                                                }
                                                _ = v28
                                            }
                                        }
                                        _ = v26
                                    } else {
                                        fmt.Println(v4, fi18)
                                    }
                                    _ = v25
                                } else {
                                    v37 := GenS1{F0: -2147483648}
                                    if ((v6 << ((2) & 63)) >= v6) {
                                        var v38 uint32 = v6
                                        _ = v38
                                    }
                                    v16 = (v16 >> ((v16) & 63))
                                    _ = v37
                                }
                                _ = v24
                            default:
                                fmt.Println(fi18, v16, v20)
                            }
                        }
                        v2[(fi19 / (((fi18 | fi18)) | 1))] = (v23 % ((1) | 1))
                        _ = v23
                    }
                case "\n\t":
                    v4 += 32767
                case "hello":
                    if (0 != v20) {
                        if true {
                            _ = len(v2)
                        }
                        var v39 bool = (1 >= -1)
                        fmt.Println(v4, v4)
                        _ = v39
                    } else {
                        if false {
                            if (57 >= 394) {
                                var v40 int32 = int32(v20)
                                if false {
                                    fmt.Println(fi18, v16, v20)
                                }
                                _ = v40
                            }
                            v41 := GenS4{F0: 865}
                            _ = v41
                        } else {
                            if (float32(1e10) > 1e10) {
                                if ((0.1 + 1e10) > float32(0.1)) {
                                    var v42 int32 = 42
                                    v43 := GenS5{F0: -32768, F1: 2}
                                    if (!false) {
                                        _ = len(v2)
                                        switch (fi18 ^ (fi18 - fi19)) {
                                        case 605:
                                            switch -197 {
                                            case 4294967295:
                                                switch 597 {
                                                case -902:
                                                    v44 := GenS1{F0: (fi18 % ((-248) | 1))}
                                                    _ = v44
                                                case 0:
                                                    delete(v2, (fi19 % (((fi19 | -9223372036854775808)) | 1)))
                                                }
                                            case -1:
                                                v42 += -2147483648
                                            case -122:
                                                switch (fi19 * fi19) {
                                                case 196:
                                                    v3 = append(v3, fi19)
                                                default:
                                                    var v45 uint = (v16 >> ((392) & 63))
                                                    _ = v45
                                                }
                                            }
                                        case 0:
                                            var v46 float32 = 0.0
                                            _ = v46
                                        }
                                        fmt.Println(v4, fi18)
                                    } else {
                                        var v47 bool = (!(997 > 42))
                                        _ = v47
                                    }
                                    v42 = (v42 << ((1) & 63))
                                    _ = v42
                                    _ = v43
                                } else {
                                    var v48 uint64 = 860
                                    switch int(v20) {
                                    case 1:
                                        var v49 bool = ((true || false) || true)
                                        _ = v49
                                    case 42:
                                        if false {
                                            if ((("\n\t" + "foo bar") < "a") || true) {
                                                if (true || (598 != 819)) {
                                                    v6 -= (v6 ^ (v6 >> (((v16 % ((uint(v48)) | 1))) & 63)))
                                                }
                                                var v50 uint16 = 1
                                                _ = v50
                                            }
                                            v51 := GenS6{F0: 52}
                                            _ = v51
                                        } else {
                                            if false {
                                                v5 = append(v5, 2147483648)
                                            } else {
                                                switch fi18 {
                                                case 127:
                                                    if (623 <= 42) {
                                                        v6 -= (v6 ^ (v20 % (((^(v6))) | 1)))
                                                        var v52 int64 = 0
                                                        _ = v52
                                                    } else {
                                                        var v53 int64 = 42
                                                        _ = v53
                                                    }
                                                case 970:
                                                    switch "hello" {
                                                    case "a":
                                                        fmt.Println(fi18, v16, fi18)
                                                    case "foo bar":
                                                        v54 := GenS7{F0: -1, F1: v4}
                                                        _ = v54
                                                    default:
                                                        v55 := GenS8{F0: fi19, F1: 28}
                                                        _ = v55
                                                    }
                                                case -300:
                                                    fmt.Println(v48)
                                                }
                                                v56 := GenS9{F0: 707, F1: -695, F2: v48}
                                                v57 := GenS7{F0: 0, F1: 1}
                                                _ = v56
                                                _ = v57
                                            }
                                            v58 := GenS9{F0: (v48 >> ((v16) & 63)), F1: 1, F2: v48}
                                            v4 -= v4
                                            _ = v58
                                        }
                                        v59 := GenS7{F0: 80, F1: (v4 << ((v16) & 63))}
                                        _ = v59
                                    case 127:
                                        var v60 string = ("foo bar" + "")
                                        _ = v60
                                    }
                                    _ = v48
                                }
                                var v61 int8 = 42
                                v62 := GenS5{F0: (v4 / ((1000) | 1)), F1: genFn0(v20)}
                                _ = v61
                                _ = v62
                            }
                            v63 := v2[fi19]
                            switch ("" + ("\n\t" + ("a" + ("" + "\n\t")))) {
                            case "hello":
                                if true {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v64 uint32 = 42
                                    v64 -= (^(v20))
                                    _ = v64
                                }
                            case "foo bar":
                                v65 := GenS10{F0: 127}
                                _ = v65
                            default:
                                var v66 string = ("a" + "\n\t")
                                _ = v66
                            }
                            _ = v63
                        }
                        var v67 int8 = 109
                        if false {
                            v68 := GenS11{F0: v67, F1: (("a" + "hello") + (("a" + "") + ""))}
                            _ = v68
                        } else {
                            fmt.Println(fi18, v4)
                        }
                        _ = v67
                    }
                }
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    case "foo bar":
        v69 := GenS6{F0: 0}
        _ = v69
    default:
        v4 = v4
    }
    v4 = v4
    v70 := GenS11{F0: int8(v16), F1: "foo bar"}
    _ = len(v1)
    if (!false) {
        v16 -= v16
        v4 = 102
        v4 = -1
    } else {
        var v71 uint = (^(v16))
        v71 = (^(v16))
        _ = v71
    }
    var v72 float64 = 1.5
    v72 -= 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v72
}
