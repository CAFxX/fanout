package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint }
type GenS2 struct { F0 int; F1 uint8; F2 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = true
    var v3 int16 = 42
    var v4 uint8 = 68
    var v5 uint8 = (v4 | v4)
    var v6 float64 = 0.1
    switch "\n\t" {
    case "\n\t":
        var v7 int = 9223372036854775807
        _ = v7
    case "a":
        v8 := GenS1{F0: 18446744073709551615, F1: ("a" + "foo bar"), F2: 42}
        _ = v8
    default:
        var v9 uint32 = 2
        _ = v9
    }
    if (v2 || false) {
        var v10 bool = v2
        _ = v10
    } else {
        switch ("" + "") {
        case "":
            v2 = v2
        case "foo bar":
            if v2 {
                var v11 int8 = 102
                _ = v11
            }
        default:
            var v12 uint = 1
            _ = v12
        }
        switch 32767 {
        case -646:
            v13 := []int{127, 110, int(v5), 1}
            _ = v13
        case 9223372036854775807:
            v14 := v1[427]
            _ = v14
        }
    }
    v15 := []int{-625, 9223372036854775807}
    v16 := GenS2{F0: 9223372036854775807, F1: (v4 * (v5 % (((^(v5))) | 1))), F2: 4294967295}
    if (true && (255 > 18446744073709551615)) {
        var v17 int32 = 1
        _ = v17
    } else {
        var v18 int = 32767
        var v19 uint16 = 65535
        _ = v18
        _ = v19
    }
    switch "\n\t" {
    case "hello":
        if false {
            for fi20 := 0; fi20 < 19; fi20++ {
                if len(v15) > 0 {
                    v15[0] = fi20
                }
                _ = fi20
            }
            var v21 int8 = 42
            _ = v21
        } else {
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 uint16 = 255
                _ = len(v15)
                if (!(v2 && true)) {
                    v2 = (v2 || v2)
                } else {
                    var v24 float64 = (((0.0 / 1.0) * 0.0) * 0.1)
                    _ = v24
                }
                _ = fi22
                _ = v23
            }
            for fi25 := 0; fi25 < 10; fi25++ {
                var v26 uint = 4294967295
                v1[1] = uint32(v26)
                for fi27 := 0; fi27 < 16; fi27++ {
                    if v2 {
                        if v2 {
                            v5 += (^(v5))
                            switch ("" + "") {
                            case "":
                                if (true && (-269 <= -9223372036854775808)) {
                                    v4 = v4
                                } else {
                                    v26 = (^(v26))
                                    if (!(true || ((!(777 == 274)) || true))) {
                                        var v28 bool = v2
                                        v5 += 255
                                        _ = v28
                                    }
                                }
                            case "a":
                                v6 += -2.25
                            default:
                                v4 = v5
                            }
                        }
                        var v29 int = (fi25 / ((fi25) | 1))
                        _ = v29
                    }
                    fmt.Println(fi27, v2, v5)
                    v4 -= v5
                    _ = fi27
                }
                _ = fi25
                _ = v26
            }
        }
    case "":
        for fi30 := 0; fi30 < 5; fi30++ {
            var v31 float32 = (1e10 / (((float32(v6) / ((0.0 * (1.5 * (((1.5 - (1.5 / 1e10)) - (0.1 * (1.5 / float32(1.5)))) * 1.5))) + 1e10)) / 3.14159) + 0.0))
            _ = fi30
            _ = v31
        }
    case "foo bar":
        var v32 int32 = 42
        _ = v32
    }
    var v33 uint64 = 18446744073709551615
    v5 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
