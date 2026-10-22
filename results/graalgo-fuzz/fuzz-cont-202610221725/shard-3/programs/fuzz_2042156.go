package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 string; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint8; F1 int8; F2 int64 }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 int64 }
func genFn0(p0 uint32) uint16 {
    var v5 float32 = 0.1
    var v6 uint = 2
    v7 := []int{-567}
    for fi8 := 0; fi8 < 12; fi8++ {
        v5 = v5
        _ = fi8
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int, p1 int16) int32 {
    if ((("" + ((("" + "hello") + (("a" + "hello") + "\n\t")) + "hello")) + "foo bar") > "foo bar") {
        var v9 int32 = -1
        _ = v9
    } else {
        if (!true) {
            var v10 uint32 = 2147483647
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 string = "a"
                fmt.Println(fi11)
                v12 = (v12 + v12)
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        p1 = (p1 << ((0) & 63))
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 uint32 = 4294967295
            var v15 uint64 = 42
            v15 = v15
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    if true {
        v16 := []int{128, 127}
        if (-1 >= -9223372036854775808) {
            fmt.Println(p1, p1)
        } else {
            _ = len(v16)
        }
        _ = v16
    } else {
        p1 = 0
        p0 -= (p0 * p0)
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    v1 := []uint32{2147483647, 0}
    var v2 int8 = 0
    var v3 uint16 = 255
    var v4 uint64 = 18446744073709551615
    var v17 string = ""
    if true {
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := GenS1{F0: 42, F1: 2147483648}
            _ = fi18
            _ = v19
        }
        if len(v1) > 0 {
            v1[0] = 0
        }
        switch 127 {
        case 127:
            var v20 int64 = -256
            _ = v20
        case 32767:
            if ((1e10 > (3.14159 + (1.5 / (1.5 / float64(((-2.25 / 3.14159) * 0.1)))))) || (true || true)) {
                if false {
                    if false {
                        v21 := GenS2{F0: v17}
                        v22 := GenS2{F0: v17}
                        var v23 float64 = ((((float64(1e10) / 1e10) / 1e10) / 0.1) - 0.1)
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        switch (("" + (("foo bar" + "") + (((v17 + "foo bar") + "foo bar") + v17))) + v17) {
                        case "":
                            v24 := GenS2{F0: ""}
                            _ = v24
                        case "hello":
                            fmt.Println(v3)
                        case "foo bar":
                            var v25 int16 = 0
                            _ = v25
                        default:
                            for fi26 := 0; fi26 < 7; fi26++ {
                                var v27 int8 = (v2 + v2)
                                v27 = (v2 >> ((4294967295) & 63))
                                _ = fi26
                                _ = v27
                            }
                        }
                        v3 += 65535
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v17 = v17
                }
            } else {
                for fi28 := 0; fi28 < 18; fi28++ {
                    v29 := GenS1{F0: -142, F1: 2}
                    var v30 float64 = float64(0.1)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                var v31 int = 223
                switch (v17 + "\n\t") {
                case "":
                    if false {
                        for fi32 := 0; fi32 < 3; fi32++ {
                            var v33 float32 = 1e10
                            fmt.Println(v4, v2)
                            _ = fi32
                            _ = v33
                        }
                        var v34 float32 = ((0.1 / 1e10) * (3.14159 * (0.1 + 0.1)))
                        _ = v34
                    } else {
                        v3 = 0
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi35 := 0; fi35 < 11; fi35++ {
                            if (!(!((0.1 <= (0.1 * 0.1)) || (false || ((fi35 >= (v31 * int(v2))) && false))))) {
                                v4 = v4
                            }
                            if false {
                                switch v17 {
                                case "a":
                                    var v36 int8 = v2
                                    _ = v36
                                default:
                                    var v37 int8 = v2
                                    _ = v37
                                }
                            } else {
                                v38 := GenS1{F0: -412, F1: 4294967295}
                                v1 = append(v1, 346)
                                _ = v38
                            }
                            _ = fi35
                        }
                    }
                case "\n\t":
                    v2 -= (v2 + (-(v2)))
                case "foo bar":
                    v39 := GenS3{F0: 255}
                    _ = v39
                default:
                    var v40 float32 = 3.14159
                    _ = v40
                }
                _ = v31
            }
        case 529:
            _ = len(v1)
        }
    }
    v17 = v17
    v41 := GenS4{F0: int32(v3)}
    for fi42 := 0; fi42 < 8; fi42++ {
        if false {
            var v43 string = (v17 + v17)
            for fi44 := 0; fi44 < 8; fi44++ {
                v45 := GenS5{F0: v4, F1: v17, F2: -1}
                _ = fi44
                _ = v45
            }
            _ = v43
        } else {
            for fi46 := 0; fi46 < 6; fi46++ {
                switch (fi42 % (((fi46 - (fi42 >> ((uint(v3)) & 63)))) | 1)) {
                case -2147483648:
                    var v47 uint16 = genFn0(42)
                    _ = v47
                case 149:
                    v48 := GenS5{F0: (v4 * (v4 >> ((960) & 63))), F1: "a", F2: 0}
                    _ = v48
                case -97:
                    v49 := GenS6{F0: 0}
                    _ = v49
                default:
                    v17 = (("foo bar" + "a") + ((v17 + v17) + v17))
                }
                if false {
                    if false {
                        if len(v1) > 0 {
                            v1[0] = 65535
                        }
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v4 += (v4 % ((42) | 1))
                }
                _ = len(v1)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v17 = (("\n\t" + (v17 + "")) + ((v17 + "") + "foo bar"))
                _ = fi46
            }
            switch (^(fi42)) {
            case 42:
                if (!false) {
                    if ((v4 * v4) <= 42) {
                        if (4294967295 <= (v4 | (v4 & v4))) {
                            var v50 int = (fi42 - -747)
                            v3 += genFn0(987)
                            if true {
                                v3 += (^(v3))
                            } else {
                                if true {
                                    var v51 uint = 0
                                    _ = v51
                                } else {
                                    switch (v17 + "\n\t") {
                                    case "a":
                                        v52 := GenS2{F0: v17}
                                        _ = v52
                                    case "":
                                        fmt.Println(v50, v4)
                                    }
                                }
                                v4 += v4
                                if true {
                                    var v53 int64 = int64(v50)
                                    fmt.Println(v2, v4, v50)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v53
                                }
                            }
                            _ = v50
                        } else {
                            switch -115 {
                            case 128:
                                switch fi42 {
                                case 2147483647:
                                    var v54 float64 = 3.14159
                                    _ = v54
                                default:
                                    switch (^(fi42)) {
                                    case 127:
                                        v17 = v17
                                    case 230:
                                        v2 -= (^(v2))
                                    case 9223372036854775807:
                                        var v55 int = -2147483648
                                        _ = v55
                                    default:
                                        fmt.Println(v3)
                                    }
                                }
                            case 42:
                                _ = len(v1)
                            case -808:
                                v56 := GenS7{F0: 2, F1: (v2 % (((v2 % ((0) | 1))) | 1)), F2: 0}
                                _ = v56
                            }
                        }
                    }
                    var v57 int = (-(fi42))
                    _ = v57
                }
            case 32767:
                v58 := GenS8{F0: v2, F1: v2}
                _ = v58
            default:
                var v59 bool = false
                _ = v59
            }
            v1 = append(v1, 42)
        }
        fmt.Println(v3, v17)
        _ = fi42
    }
    if len(v1) > 0 {
        v1[0] = 428
    }
    if false {
        if true {
            _ = len(v1)
        }
        var v60 uint16 = v3
        _ = v60
    }
    var v61 bool = ((v2 >> ((344) & 63)) < -1)
    v1 = append(v1, 65535)
    var v62 string = v17
    for fi63 := 0; fi63 < 11; fi63++ {
        fmt.Println(v3, v61, v62)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if v61 {
            _ = len(v1)
            v1 = append(v1, 555)
            v64 := GenS2{F0: (v17 + v17)}
            _ = v64
        } else {
            for fi65 := 0; fi65 < 18; fi65++ {
                var v66 string = "a"
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
                v67 := GenS9{F0: 1}
                v2 = 100
                v68 := GenS3{F0: 42}
                _ = fi65
                _ = v66
                _ = v67
                _ = v68
            }
            var v69 uint32 = 0
            _ = v69
        }
        _ = fi63
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v61
    _ = v62
}
