package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 string; F1 uint64; F2 int }
type GenS4 struct { F0 uint8; F1 int64; F2 int8 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint16, p1 string) int {
    for fi6 := 0; fi6 < 18; fi6++ {
        if true {
            var v7 int = fi6
            _ = v7
        }
        var v8 uint32 = 4294967295
        var v9 bool = false
        _ = fi6
        _ = v8
        _ = v9
    }
    var v10 float32 = -2.25
    switch 828 {
    case -2147483648:
        v11 := GenS1{F0: -391, F1: 222}
        _ = v11
    default:
        fmt.Println(v10)
    }
    v12 := GenS1{F0: -9223372036854775808, F1: 65}
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    return -2147483648
}

func main() {
    v1 := []uint32{2147483648, 42, 2147483647}
    var v2 bool = true
    v3 := []int{32767, -103}
    var v4 int16 = -1
    var v5 uint = 2
    for fi13 := 0; fi13 < 16; fi13++ {
        v4 = v4
        var v14 uint32 = 65535
        v15 := GenS2{F0: -128, F1: 65535, F2: -716}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint = v5
    var v17 int64 = 42
    v2 = v2
    for fi18 := 0; fi18 < 20; fi18++ {
        for fi19 := 0; fi19 < 5; fi19++ {
            switch ("a" + "foo bar") {
            case "hello":
                var v20 string = "\n\t"
                _ = v20
            case "":
                var v21 float64 = 3.14159
                _ = v21
            }
            var v22 int16 = v4
            if ((!v2) || true) {
                if v2 {
                    v2 = v2
                    if v2 {
                        v23 := GenS2{F0: -65, F1: 0, F2: 963}
                        fmt.Println(v4, v22)
                        v17 += (v17 * (v17 ^ v17))
                        _ = v23
                    } else {
                        v24 := GenS3{F0: (("" + ("foo bar" + "hello")) + ("\n\t" + "\n\t")), F1: 4294967295, F2: 9223372036854775807}
                        var v25 float64 = float64(0.1)
                        if (!false) {
                            var v26 uint = v16
                            _ = v26
                        }
                        _ = v24
                        _ = v25
                    }
                } else {
                    var v27 int32 = -2147483648
                    v2 = v2
                    if v2 {
                        var v28 int8 = -1
                        v29 := GenS3{F0: ("foo bar" + ""), F1: 2, F2: 1}
                        v3 = append(v3, (fi18 % (((fi18 ^ genFn0(42, ("" + "")))) | 1)))
                        _ = v28
                        _ = v29
                    }
                    _ = v27
                }
            } else {
                var v30 float64 = 3.14159
                _ = v30
            }
            _ = fi19
            _ = v22
        }
        _ = fi18
    }
    var v31 float64 = (1e10 * (1.5 + (3.14159 + 3.14159)))
    var v32 uint8 = 167
    v33 := GenS4{F0: 111, F1: -9223372036854775808, F2: -8}
    for fi34 := 0; fi34 < 7; fi34++ {
        var v35 uint32 = 512
        _ = fi34
        _ = v35
    }
    fmt.Println(v16)
    if (v16 >= v5) {
        var v36 bool = v2
        v31 = 1.5
        var v37 int32 = -2147483648
        _ = v36
        _ = v37
    } else {
        if (false && ((v2 && false) && (v2 || ((-507 <= -668) && true)))) {
            switch "a" {
            case "\n\t":
                switch genFn0(11, ("a" + "\n\t")) {
                case 1:
                    var v38 float64 = float64(((((-2.25 / 3.14159) + float32((v31 * 1e10))) * 3.14159) - 3.14159))
                    _ = v38
                default:
                    var v39 uint64 = 2
                    _ = v39
                }
            default:
                switch 2147483647 {
                case 1:
                    _ = len(v3)
                case 32767:
                    var v40 float32 = 1e10
                    _ = v40
                case 128:
                    v2 = (v32 >= (v32 & (v32 ^ 42)))
                }
            }
            v41 := GenS1{F0: -487, F1: 128}
            if v2 {
                _ = len(v3)
            } else {
                for fi42 := 0; fi42 < 20; fi42++ {
                    switch -1 {
                    case 42:
                        var v43 int32 = 910
                        _ = v43
                    case 127:
                        var v44 uint16 = 1
                        _ = v44
                    case -2147483648:
                        switch ("hello" + "\n\t") {
                        case "\n\t":
                            if (v2 && (-1 > -1)) {
                                if (0 != uint16(v5)) {
                                    var v45 uint = 155
                                    v4 = (v4 >> ((v5) & 63))
                                    _ = v45
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch "\n\t" {
                                case "":
                                    v46 := GenS5{F0: (v5 * v5), F1: 554}
                                    _ = v46
                                case "\n\t":
                                    fmt.Println(v4, fi42)
                                case "foo bar":
                                    var v47 uint = v16
                                    _ = v47
                                default:
                                    v4 -= 32767
                                }
                            }
                        case "foo bar":
                            if v2 {
                                v48 := GenS6{F0: uint32(v16)}
                                _ = v48
                            }
                        case "":
                            var v49 string = ""
                            _ = v49
                        }
                    }
                    _ = fi42
                }
                fmt.Println(v32, v4)
            }
            _ = v41
        }
        switch genFn0(1, "hello") {
        case 2147483647:
            v31 = v31
        }
        for fi50 := 0; fi50 < 6; fi50++ {
            switch -2147483648 {
            case 42:
                if (v32 <= v32) {
                    v16 = (v5 | v5)
                    if v2 {
                        v51 := GenS6{F0: 2147483648}
                        _ = v51
                    }
                } else {
                    for fi52 := 0; fi52 < 18; fi52++ {
                        fmt.Println(fi52, v17, v16)
                        if v2 {
                            var v53 int16 = (v4 * v4)
                            _ = v53
                        }
                        _ = fi52
                    }
                    v32 += v32
                    v54 := GenS6{F0: 338}
                    _ = v54
                }
            case 32767:
                var v55 uint64 = 18446744073709551615
                _ = v55
            default:
                fmt.Println(v31)
            }
            fmt.Println(v2, v16)
            _ = fi50
        }
    }
    v31 = v31
    var v56 int = -9223372036854775808
    if v2 {
        var v57 bool = (v17 >= v17)
        var v58 int16 = 1000
        _ = v57
        _ = v58
    }
    var v59 uint = uint(v32)
    var v60 uint8 = (v32 + (^(v32)))
    if len(v1) > 0 {
        v1[0] = 42
    }
    for fi61 := 0; fi61 < 3; fi61++ {
        var v62 int8 = -128
        var v63 uint32 = 742
        var v64 int32 = 2147483647
        _ = fi61
        _ = v62
        _ = v63
        _ = v64
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v56
    _ = v59
    _ = v60
}
