package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
func genFn0(p0 uint32) uint16 {
    var v4 uint64 = 979
    var v5 uint16 = 42
    fmt.Println(v4, v4, v5)
    if (1 >= 287) {
        fmt.Println(p0, p0, v5)
    }
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v5
    return (^(v5))
}
func genFn1(p0 uint16, p1 uint, p2 int64) uint {
    var v6 int16 = int16(p0)
    var v7 bool = false
    v8 := []int{-1}
    if (127 == 1) {
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: 42}
            for fi11 := 0; fi11 < 8; fi11++ {
                p2 -= (-(p2))
                if v7 {
                    fmt.Println(p0)
                } else {
                    var v12 uint16 = p0
                    p0 -= p0
                    v8 = append(v8, (fi11 * (fi11 - fi9)))
                    _ = v12
                }
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
        if false {
            v7 = v7
            switch "foo bar" {
            case "foo bar":
                for fi13 := 0; fi13 < 11; fi13++ {
                    if ((!(65535 > 127)) || true) {
                        for fi14 := 0; fi14 < 9; fi14++ {
                            v6 = 26
                            p1 = p1
                            switch ((("\n\t" + (("" + ("" + "foo bar")) + "foo bar")) + "") + "foo bar") {
                            case "hello":
                                var v15 uint8 = 128
                                var v16 uint = p1
                                var v17 int = fi13
                                _ = v15
                                _ = v16
                                _ = v17
                            case "\n\t":
                                v18 := GenS1{F0: -8}
                                fmt.Println(fi14, p2)
                                if (true && v7) {
                                    var v19 int64 = 42
                                    if len(v8) > 0 {
                                        v8[0] = -713
                                    }
                                    v20 := GenS1{F0: 42}
                                    _ = v19
                                    _ = v20
                                } else {
                                    v21 := GenS1{F0: (v6 ^ -84)}
                                    _ = v21
                                }
                                _ = v18
                            case "":
                                fmt.Println(p2, p0, p2)
                            default:
                                var v22 float64 = (((((-2.25 * (-2.25 * ((0.1 - 3.14159) * 1.5))) + (((3.14159 + 0.1) * 1e10) / 1e10)) - -2.25) * (1.5 * (float64(float32((0.1 * 0.1))) * 1e10))) * 3.14159)
                                _ = v22
                            }
                            _ = fi14
                        }
                    } else {
                        var v23 int = (fi13 >> (((^(p1))) & 63))
                        var v24 int = 4294967295
                        _ = v23
                        _ = v24
                    }
                    p2 = p2
                    _ = fi13
                }
            }
        } else {
            if (0 <= 105) {
                switch ("foo bar" + "a") {
                case "":
                    p2 = 626
                default:
                    var v25 bool = v7
                    _ = v25
                }
            } else {
                for fi26 := 0; fi26 < 3; fi26++ {
                    fmt.Println(p1, v7)
                    _ = fi26
                }
                switch 827 {
                case -9223372036854775808:
                    var v27 uint = 42
                    _ = v27
                case 122:
                    v28 := GenS1{F0: 1}
                    _ = v28
                case -1:
                    var v29 uint32 = 42
                    _ = v29
                }
                v30 := GenS2{F0: p0, F1: v6}
                _ = v30
            }
        }
    } else {
        if (1 >= 204) {
            if (p0 <= (^(p0))) {
                var v31 int32 = 685
                if (1.5 == 1.5) {
                    var v32 string = "a"
                    _ = v32
                }
                _ = v31
            }
            if (-9223372036854775808 != p2) {
                if (1 > -224) {
                    p1 = p1
                    for fi33 := 0; fi33 < 1; fi33++ {
                        v7 = (!(("hello" + "") <= (("a" + "foo bar") + ("hello" + "\n\t"))))
                        _ = fi33
                    }
                } else {
                    for fi34 := 0; fi34 < 15; fi34++ {
                        fmt.Println(v7, fi34)
                        _ = fi34
                    }
                    v35 := GenS3{F0: ((^(p1)) > 255)}
                    fmt.Println(p1, p0, v6)
                    _ = v35
                }
                v6 = (v6 & 32767)
            } else {
                switch "foo bar" {
                case "hello":
                    var v36 uint = p1
                    _ = v36
                case "":
                    var v37 uint16 = 481
                    _ = v37
                case "\n\t":
                    v38 := GenS3{F0: (579 != 358)}
                    _ = v38
                }
            }
            var v39 uint = p1
            _ = v39
        }
        var v40 uint64 = 18446744073709551615
        _ = v40
    }
    var v41 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v41
    _ = v6
    _ = v7
    _ = v8
    return p1
}

func main() {
    var v1 uint16 = 0
    var v2 uint32 = 228
    v3 := make(map[string]int)
    fmt.Println(v1, v2)
    var v42 int16 = 999
    switch "a" {
    case "a":
        var v43 uint8 = 127
        _ = v43
    case "":
        delete(v3, "\n\t")
    case "foo bar":
        v44 := GenS1{F0: (v42 >> ((genFn1((^(v1)), 18446744073709551615, -9223372036854775808)) & 63))}
        _ = v44
    default:
        v45 := GenS2{F0: 1, F1: 0}
        _ = v45
    }
    if (842 != 0) {
        v46 := GenS4{F0: "a"}
        v47 := GenS2{F0: v1, F1: v42}
        v1 -= (v1 * v1)
        _ = v46
        _ = v47
    }
    if (false || false) {
        switch 32768 {
        case 13:
            fmt.Println(v2, v42)
        default:
            var v48 bool = false
            _ = v48
        }
    } else {
        v1 -= v1
        v49 := []int{int(v2), 1, -231, -2147483648}
        _ = v49
    }
    delete(v3, "")
    var v50 uint8 = 127
    if true {
        var v51 int16 = (v42 ^ (v42 << ((51) & 63)))
        fmt.Println(v42, v2, v42)
        _ = v51
    }
    var v52 float32 = -2.25
    if (1e10 == 0.1) {
        for fi53 := 0; fi53 < 12; fi53++ {
            fmt.Println(v2, v2)
            _ = fi53
        }
        var v54 uint16 = 42
        _ = v54
    }
    var v55 string = "foo bar"
    var v56 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v50
    _ = v52
    _ = v55
    _ = v56
}
