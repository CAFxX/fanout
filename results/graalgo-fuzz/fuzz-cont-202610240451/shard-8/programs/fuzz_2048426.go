package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int64 }
func genFn0(p0 int16, p1 int8, p2 uint) uint64 {
    v7 := []int{}
    if (2147483647 != 0) {
        v8 := GenS1{F0: 42, F1: -57, F2: p0}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 420
}
func genFn1(p0 int32) bool {
    p0 = 0
    p0 += p0
    fmt.Println(p0)
    v9 := []int{-1, 0, 9223372036854775807, 496}
    _ = p0
    _ = v9
    return (400 > 9223372036854775807)
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = -788
    v3 := []uint32{}
    var v4 uint8 = 15
    v5 := []uint32{0, 4294967295, 42, 2147483648}
    var v6 bool = true
    switch ((("a" + "a") + ("hello" + "hello")) + ("a" + "hello")) {
    case "foo bar":
        var v10 uint = 18446744073709551615
        _ = v10
    case "a":
        fmt.Println(v2, v2, v2)
    default:
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS2{F0: (v2 * 318), F1: uint32(v4)}
            _ = fi11
            _ = v12
        }
    }
    if (1 <= 0) {
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := GenS2{F0: v2, F1: 933}
            switch fi13 {
            case 9223372036854775807:
                v6 = (0 < 9223372036854775807)
            case 0:
                v6 = (255 == 285)
            case -87:
                v6 = (0 != 42)
            }
            var v15 string = ("\n\t" + "hello")
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 19; fi16++ {
            v4 -= v4
            fmt.Println(v2)
            _ = fi16
        }
        if (2 <= 255) {
            if (v6 && (!v6)) {
                for fi17 := 0; fi17 < 9; fi17++ {
                    v18 := GenS3{F0: 42}
                    if (!false) {
                        v19 := GenS3{F0: 0}
                        if (v6 || (128 > (v4 | (v4 << ((2) & 63))))) {
                            var v20 uint32 = 590
                            _ = v20
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v21 := v1[(("hello" + (((("foo bar" + "a") + "") + "foo bar") + "hello")) + "a")]
                            _ = v21
                        }
                        _ = v19
                    }
                    _ = fi17
                    _ = v18
                }
            }
            for fi22 := 0; fi22 < 4; fi22++ {
                v4 -= v4
                for fi23 := 0; fi23 < 8; fi23++ {
                    v24 := GenS4{F0: 2, F1: (fi23 / ((fi23) | 1))}
                    v4 -= v4
                    var v25 uint16 = 457
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi22
            }
        }
    } else {
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 int = (fi26 * (-(fi26)))
            var v28 float64 = -2.25
            var v29 int32 = -58
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        switch "" {
        case "a":
            var v30 float32 = (0.0 - 0.1)
            _ = v30
        case "":
            var v31 int64 = -1
            _ = v31
        }
        v32 := GenS2{F0: (v2 + v2), F1: 4294967295}
        _ = v32
    }
    if false {
        var v33 int64 = -9223372036854775808
        switch -2147483648 {
        case 42:
            var v34 int32 = 42
            _ = v34
        case -2147483648:
            v2 -= -868
        }
        if len(v5) > 0 {
            v5[0] = 2
        }
        _ = v33
    } else {
        fmt.Println(v4, v2, v2)
        fmt.Println(v2)
        fmt.Println(v2, v2)
    }
    if false {
        switch 2147483647 {
        case -391:
            var v35 uint8 = 255
            _ = v35
        case -96:
            v2 += 212
        default:
            for fi36 := 0; fi36 < 14; fi36++ {
                var v37 uint = 2
                var v38 int8 = 42
                _ = fi36
                _ = v37
                _ = v38
            }
        }
        if (9223372036854775807 >= 307) {
            v39 := v1["hello"]
            v39 -= (-(v39))
            _ = v39
        } else {
            if false {
                v1["hello"] = 4294967295
                switch -2147483648 {
                case 2147483647:
                    if (v6 || (18446744073709551615 <= 535)) {
                        v2 += (v2 >> ((65535) & 63))
                    } else {
                        v2 = (v2 % ((v2) | 1))
                        var v40 int8 = 107
                        _ = v40
                    }
                default:
                    v1[("\n\t" + ("hello" + "\n\t"))] = -491
                }
                v6 = true
            } else {
                v2 += v2
                _ = len(v1)
                var v41 string = ""
                _ = v41
            }
            switch 1 {
            case 128:
                var v42 uint32 = 85
                _ = v42
            case -380:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 32767:
                if v6 {
                    v43 := GenS5{F0: int64(v4)}
                    _ = v43
                } else {
                    for fi44 := 0; fi44 < 8; fi44++ {
                        var v45 uint64 = 2
                        _ = fi44
                        _ = v45
                    }
                }
            default:
                if (!true) {
                    switch "a" {
                    case "\n\t":
                        var v46 int64 = 1
                        _ = v46
                    case "a":
                        var v47 uint32 = 4294967295
                        _ = v47
                    case "hello":
                        if v6 {
                            v2 = (^(v2))
                        }
                    default:
                        if (v4 < (v4 * 2)) {
                            v4 = (^(v4))
                        }
                    }
                    for fi48 := 0; fi48 < 3; fi48++ {
                        if (34 < -128) {
                            switch "hello" {
                            case "\n\t":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                        _ = fi48
                    }
                    v2 += v2
                }
            }
        }
    } else {
        switch (("" + (("hello" + "foo bar") + "\n\t")) + "foo bar") {
        case "foo bar":
            for fi49 := 0; fi49 < 2; fi49++ {
                v50 := GenS2{F0: v2, F1: 4294967295}
                _ = fi49
                _ = v50
            }
        case "\n\t":
            v6 = (!(!(!(0 < 2))))
        }
        var v51 int = 872
        fmt.Println(v2)
        _ = v51
    }
    v2 += v2
    v2 = -1
    var v52 float64 = -2.25
    switch "foo bar" {
    case "\n\t":
        v2 = -1
    case "a":
        for fi53 := 0; fi53 < 19; fi53++ {
            for fi54 := 0; fi54 < 20; fi54++ {
                v55 := GenS1{F0: 1, F1: 75, F2: -1}
                v2 = v2
                _ = fi54
                _ = v55
            }
            var v56 float32 = 3.14159
            _ = fi53
            _ = v56
        }
    case "hello":
        v57 := GenS5{F0: 9223372036854775807}
        _ = v57
    }
    if v6 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v58 int8 = -128
        fmt.Println(v2, v6)
        _ = v58
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v2 += (v2 % (((^(v2))) | 1))
    switch 259 {
    case 127:
        if v6 {
            var v59 float64 = v52
            fmt.Println(v52, v52, v2)
            _ = v59
        }
    case 0:
        v52 = v52
    case -378:
        if (!false) {
            fmt.Println(v4, v4, v4)
            for fi60 := 0; fi60 < 3; fi60++ {
                if v6 {
                    v6 = (!true)
                    fmt.Println(fi60)
                    v52 = (v52 - float64(float32(((v52 + v52) / ((v52 + v52) * 0.0)))))
                } else {
                    v61 := GenS3{F0: 0}
                    _ = v61
                }
                if (!(2147483647 <= 0)) {
                    if (!((2 > 42) && (!v6))) {
                        fmt.Println(fi60, v2)
                        var v62 float64 = (1.5 / (v52 / (v52 + v52)))
                        _ = v62
                    }
                    v52 = v52
                } else {
                    for fi63 := 0; fi63 < 20; fi63++ {
                        switch fi63 {
                        case -9223372036854775808:
                            v52 = v52
                        case 127:
                            switch "a" {
                            case "foo bar":
                                var v64 int = 0
                                _ = v64
                            case "":
                                delete(v1, "foo bar")
                            case "hello":
                                v2 = v2
                                var v65 uint = 223
                                _ = v65
                            }
                        }
                        _ = fi63
                    }
                    var v66 uint = 621
                    _ = v66
                }
                _ = fi60
            }
        }
    default:
        var v67 float32 = -2.25
        _ = v67
    }
    fmt.Println(v6, v6)
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
