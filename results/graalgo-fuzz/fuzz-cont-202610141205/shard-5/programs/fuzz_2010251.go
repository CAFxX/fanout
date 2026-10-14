package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64 }
func genFn0() int8 {
    switch -1 {
    case -1:
        v7 := GenS1{F0: 575, F1: 1}
        _ = v7
    case -327:
        var v8 uint32 = 2147483647
        _ = v8
    }
    var v9 int16 = -64
    v10 := []int{309, 4294967295, -1, 42}
    for fi11 := 0; fi11 < 18; fi11++ {
        if len(v10) > 0 {
            v10[0] = fi11
        }
        var v12 int8 = -117
        _ = fi11
        _ = v12
    }
    if (false || true) {
        var v13 uint32 = 2
        var v14 int64 = 0
        _ = v13
        _ = v14
    } else {
        if ("hello" == "hello") {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                if false {
                    v16 := GenS2{F0: 2, F1: 128, F2: 393}
                    _ = v16
                }
                for fi17 := 0; fi17 < 12; fi17++ {
                    var v18 uint64 = 4294967295
                    fmt.Println(v9, v18, fi15)
                    _ = fi17
                    _ = v18
                }
                _ = fi15
            }
        }
        v9 = -1
    }
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 int64 = -393
    var v2 int8 = 100
    v3 := []int{9223372036854775807, 32768}
    var v4 bool = (false && (!true))
    v5 := make(map[uint32]bool)
    v6 := make(map[string]bool)
    for fi19 := 0; fi19 < 7; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            if ((!true) && (false && ((v2 % ((genFn0()) | 1)) == v2))) {
                if (true || v4) {
                    v1 -= -809
                    v21 := GenS1{F0: 2147483648, F1: 2147483647}
                    _ = v21
                } else {
                    switch 127 {
                    case -9223372036854775808:
                        v4 = (!true)
                        if v4 {
                            var v22 uint = 42
                            var v23 int16 = 42
                            _ = v22
                            _ = v23
                        } else {
                            var v24 string = (("hello" + (("hello" + "foo bar") + "\n\t")) + "foo bar")
                            v24 = v24
                            _ = v24
                        }
                    case -2147483648:
                        var v25 int8 = v2
                        _ = v25
                    default:
                        switch "\n\t" {
                        case "a":
                            var v26 uint16 = 65535
                            _ = v26
                        }
                    }
                }
                if (!false) {
                    v2 += (v2 | v2)
                    v27 := GenS1{F0: 1, F1: 0}
                    _ = v27
                }
                v2 -= v2
            }
            var v28 uint64 = 0
            v29 := GenS1{F0: 2, F1: 1}
            _ = fi20
            _ = v28
            _ = v29
        }
        v30 := GenS2{F0: 128, F1: 148, F2: 65535}
        _ = fi19
        _ = v30
    }
    v31 := GenS2{F0: 2, F1: 0, F2: 2147483648}
    v2 -= genFn0()
    for fi32 := 0; fi32 < 6; fi32++ {
        var v33 int8 = (-(v2))
        _ = fi32
        _ = v33
    }
    if (false || (!(-219 < 42))) {
        var v34 int16 = int16(v1)
        switch 4294967295 {
        case 597:
            if (-2.25 == 0.0) {
                var v35 int32 = -2147483648
                v4 = v4
                _ = v35
            } else {
                v2 += (v2 >> ((255) & 63))
                switch 42 {
                case 616:
                    if (-2.25 <= 3.14159) {
                        switch 42 {
                        case 32767:
                            var v36 uint = 18446744073709551615
                            _ = v36
                        case -633:
                            if v4 {
                                v37 := GenS2{F0: 42, F1: 229, F2: 0}
                                _ = v37
                            }
                        }
                    }
                }
                v38 := GenS3{F0: (v1 - (v1 + -578))}
                _ = v38
            }
        case -839:
            if len(v3) > 0 {
                v3[0] = 9223372036854775807
            }
        default:
            v3 = append(v3, 127)
        }
        _ = v34
    }
    if (((((0.0 - float32((1e10 * 1e10))) + 1e10) * 1.5) == 1.5) || v4) {
        var v39 uint32 = 4294967295
        v39 = (^(v39))
        _ = v39
    }
    var v40 float64 = (-2.25 / (0.1 / -2.25))
    switch "foo bar" {
    case "\n\t":
        var v41 uint8 = 128
        _ = v41
    default:
        var v42 uint = 255
        _ = v42
    }
    if len(v3) > 0 {
        v3[0] = -1
    }
    var v43 int32 = -2147483648
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v43
    _ = v5
    _ = v6
}
