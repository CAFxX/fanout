package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
func genFn0() bool {
    var v5 uint = 18446744073709551615
    var v6 bool = false
    _ = v5
    _ = v6
    return ((128 != 42) && v6)
}

func main() {
    var v1 int64 = 0
    var v2 uint = 892
    v3 := []uint32{}
    var v4 uint = 65535
    fmt.Println(v1, v2)
    fmt.Println(v2, v2)
    fmt.Println(v4)
    v4 = v2
    for fi7 := 0; fi7 < 16; fi7++ {
        if true {
            v1 = int64(fi7)
        }
        _ = fi7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int32 = -483
        _ = fi8
        _ = v9
    }
    v1 = v1
    if (((0.1 - 1.5) >= 3.14159) || ((2 < 2) && true)) {
        if len(v3) > 0 {
            v3[0] = 0
        }
        switch 0 {
        case -824:
            if true {
                v10 := GenS1{F0: -664, F1: 2}
                switch -929 {
                case -2147483648:
                    v3 = append(v3, 2147483647)
                default:
                    v4 = v4
                }
                _ = v10
            } else {
                var v11 uint32 = 42
                switch ("hello" + "") {
                case "":
                    v11 = (v11 & v11)
                case "a":
                    fmt.Println(v11, v4)
                case "\n\t":
                    switch "a" {
                    case "hello":
                        if len(v3) > 0 {
                            v3[0] = 4294967295
                        }
                    default:
                        var v12 int32 = 1
                        _ = v12
                    }
                default:
                    _ = len(v3)
                }
                v11 = v11
                _ = v11
            }
        default:
            for fi13 := 0; fi13 < 5; fi13++ {
                for fi14 := 0; fi14 < 3; fi14++ {
                    switch "foo bar" {
                    case "":
                        if (false || true) {
                            v3 = append(v3, 42)
                            switch (-(fi14)) {
                            case 575:
                                var v15 int32 = 0
                                if (0 <= v15) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 346
                                    }
                                }
                                _ = v15
                            default:
                                fmt.Println(fi13, v4, v1)
                            }
                        }
                    case "hello":
                        switch "hello" {
                        case "foo bar":
                            v4 -= (v4 | v4)
                        case "\n\t":
                            v4 = (v4 - v2)
                        case "hello":
                            if false {
                                v1 -= 9223372036854775807
                                fmt.Println(v4, fi13, v2)
                                v16 := GenS1{F0: 858, F1: 42}
                                var v17 uint8 = 216
                                _ = v16
                                _ = v17
                            } else {
                                if ((("\n\t" + "\n\t") + ("\n\t" + "")) > (("foo bar" + "foo bar") + ("foo bar" + "hello"))) {
                                    fmt.Println(v4, v2, v1)
                                    v18 := GenS1{F0: 207, F1: 819}
                                    v2 = 33
                                    _ = v18
                                }
                                fmt.Println(v4, v2)
                            }
                            v3 = append(v3, 0)
                            v4 -= 255
                        }
                        _ = len(v3)
                    default:
                        fmt.Println(v1)
                    }
                    _ = fi14
                }
                v1 = -1
                v19 := GenS1{F0: -1, F1: 747}
                _ = fi13
                _ = v19
            }
        }
        switch ("" + "a") {
        case "foo bar":
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 int16 = -711
                _ = fi20
                _ = v21
            }
        case "\n\t":
            fmt.Println(v1, v2)
        case "":
            if ((!true) || (-861 >= -455)) {
                var v22 uint = (^(v2))
                var v23 uint32 = 4294967295
                switch 42 {
                case 9223372036854775807:
                    var v24 uint16 = 50
                    _ = v24
                case 2147483647:
                    fmt.Println(v1)
                case -555:
                    fmt.Println(v22)
                }
                _ = v22
                _ = v23
            } else {
                if (!(int(v1) >= 4294967295)) {
                    var v25 float32 = 1e10
                    v4 -= 1
                    v25 = 0.0
                    _ = v25
                }
            }
        default:
            v26 := GenS1{F0: -32768, F1: 0}
            _ = v26
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = 9223372036854775807
    }
    v27 := GenS1{F0: 668, F1: 886}
    v28 := GenS1{F0: -425, F1: 788}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
