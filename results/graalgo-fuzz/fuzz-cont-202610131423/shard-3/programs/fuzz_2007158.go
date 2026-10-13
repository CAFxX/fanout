package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int8; F1 string }
func genFn0(p0 int8, p1 uint64, p2 int8) int8 {
    var v5 uint = 793
    v5 += v5
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -128
}
func genFn1(p0 uint8, p1 uint16, p2 bool) uint32 {
    p0 = p0
    v6 := []int{4294967295, -9223372036854775808, 819, int(p1)}
    for fi7 := 0; fi7 < 18; fi7++ {
        p1 = p1
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v8 int64 = 1
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483648
}

func main() {
    v1 := []string{"\n\t", "foo bar", "foo bar"}
    var v2 bool = false
    var v3 string = ("\n\t" + "")
    v4 := []int{-444, 42, 2147483647}
    v2 = (0 >= 0)
    switch "hello" {
    case "\n\t":
        v3 = "foo bar"
    }
    var v9 float64 = 3.14159
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 = 0.0
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 float64 = 0.1
        var v12 float64 = 3.14159
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int32 = 2147483647
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint32 = 2147483647
        fmt.Println(v2, v2)
        fmt.Println(v2, v15, fi14)
        _ = fi14
        _ = v15
    }
    fmt.Println(v9, v13, v3)
    for fi16 := 0; fi16 < 11; fi16++ {
        v3 = v3
        v17 := GenS1{F0: 140, F1: 0}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi16
        _ = v17
    }
    if (false && v2) {
        v2 = v2
        if v2 {
            v9 = v9
        }
        v18 := GenS1{F0: 18446744073709551615, F1: 127}
        _ = v18
    } else {
        if false {
            for fi19 := 0; fi19 < 3; fi19++ {
                v3 = "hello"
                _ = fi19
            }
            v20 := GenS1{F0: 4294967295, F1: -105}
            _ = v20
        } else {
            switch v3 {
            case "foo bar":
                var v21 string = (v3 + v3)
                _ = v21
            case "\n\t":
                v4 = append(v4, 127)
            case "":
                v9 -= (3.14159 - 0.0)
            default:
                for fi22 := 0; fi22 < 12; fi22++ {
                    if len(v4) > 0 {
                        v4[0] = (fi22 | (fi22 | 4294967295))
                    }
                    _ = fi22
                }
            }
        }
        var v23 int = 42
        fmt.Println(v2, v23, v2)
        _ = v23
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        if v2 {
            v9 = 3.14159
            for fi25 := 0; fi25 < 17; fi25++ {
                v4 = append(v4, 32768)
                v1 = append(v1, "\n\t")
                v3 = ("hello" + (v3 + v3))
                if len(v4) > 0 {
                    v4[0] = fi25
                }
                _ = fi25
            }
        } else {
            if (-521 >= 1) {
                _ = len(v4)
            }
        }
        fmt.Println(v13, v2)
        if (v3 <= (v3 + v3)) {
            fmt.Println(v9)
            v13 = -480
        } else {
            if true {
                v9 += 3.14159
                if (0.0 >= float32(((0.1 + v9) + v9))) {
                    var v26 uint64 = 2
                    _ = v26
                } else {
                    v27 := GenS2{F0: 0, F1: ("" + v3)}
                    if false {
                        v1 = append(v1, v3)
                        switch "" {
                        case "a":
                            if (!(genFn0(-128, 0, genFn0(127, 4294967295, -1)) < 0)) {
                                if true {
                                    var v28 uint = 1
                                    var v29 uint32 = 0
                                    for fi30 := 0; fi30 < 8; fi30++ {
                                        fmt.Println(v13, v28)
                                        var v31 int32 = v13
                                        _ = fi30
                                        _ = v31
                                    }
                                    _ = v28
                                    _ = v29
                                }
                                if false {
                                    fmt.Println(v3, fi24)
                                    for fi32 := 0; fi32 < 15; fi32++ {
                                        v9 += ((v9 * 3.14159) / 1.0)
                                        _ = fi32
                                    }
                                } else {
                                    v13 += 1
                                    switch ("foo bar" + ("hello" + v3)) {
                                    case "foo bar":
                                        v33 := GenS2{F0: 127, F1: (v3 + (v3 + v3))}
                                        _ = v33
                                    default:
                                        switch ("a" + v3) {
                                        case "hello":
                                            var v34 uint16 = 2
                                            _ = v34
                                        case "\n\t":
                                            var v35 uint = 4294967295
                                            _ = v35
                                        default:
                                            if (!true) {
                                                for fi36 := 0; fi36 < 10; fi36++ {
                                                    switch v3 {
                                                    case "foo bar":
                                                        var v37 int16 = 1
                                                        _ = v37
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    if (32767 >= 255) {
                                                        v3 = ""
                                                        var v38 int64 = 1
                                                        fmt.Println(v38, v38, v2)
                                                        _ = v38
                                                    } else {
                                                        if len(v4) > 0 {
                                                            v4[0] = fi24
                                                        }
                                                        v13 = (v13 >> ((4294967295) & 63))
                                                        v13 = v13
                                                    }
                                                    _ = fi36
                                                }
                                                v3 = (v3 + v3)
                                            }
                                        }
                                    }
                                }
                            }
                        case "\n\t":
                            v3 = (v3 + (v3 + "foo bar"))
                        case "":
                            if ((4294967295 != 2) && v2) {
                                v39 := GenS1{F0: 18446744073709551615, F1: 127}
                                var v40 int64 = 1
                                v9 += (v9 * v9)
                                _ = v39
                                _ = v40
                            } else {
                                fmt.Println(fi24, v2)
                                v41 := GenS2{F0: 117, F1: "foo bar"}
                                _ = v41
                            }
                        default:
                            v2 = ((genFn1(0, 540, false) <= genFn1(127, 899, v2)) || v2)
                        }
                        if len(v1) > 0 {
                            v1[0] = (v3 + "a")
                        }
                    }
                    var v42 uint16 = 0
                    _ = v27
                    _ = v42
                }
                if len(v4) > 0 {
                    v4[0] = fi24
                }
            }
            var v43 uint16 = 32767
            if (502 == 0) {
                v44 := GenS2{F0: -128, F1: v3}
                _ = v44
            } else {
                if true {
                    var v45 uint8 = 0
                    var v46 int16 = -32768
                    _ = v45
                    _ = v46
                }
            }
            _ = v43
        }
        _ = fi24
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        if (1e10 >= float64((1.5 / 1e10))) {
            v48 := GenS1{F0: 448, F1: -1}
            if ((false || (fi47 != (fi47 ^ fi47))) && (!false)) {
                v9 = 0.0
            } else {
                if (!true) {
                    var v49 int32 = v13
                    _ = v49
                }
            }
            _ = v48
        }
        switch v3 {
        case "":
            for fi50 := 0; fi50 < 8; fi50++ {
                var v51 uint = 18446744073709551615
                _ = fi50
                _ = v51
            }
        default:
            var v52 float32 = (3.14159 - 1.5)
            _ = v52
        }
        _ = fi47
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    v3 = (v3 + (v3 + v3))
    v1 = append(v1, "foo bar")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
