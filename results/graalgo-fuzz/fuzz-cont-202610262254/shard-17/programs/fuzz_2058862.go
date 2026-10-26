package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
type GenS2 struct { F0 bool; F1 int32; F2 int64 }
func genFn0() int {
    v5 := GenS1{F0: 152, F1: 42, F2: 0}
    var v6 bool = true
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            if true {
                var v9 uint64 = 42
                var v10 uint8 = 0
                v11 := GenS1{F0: 1, F1: 1, F2: int64(v9)}
                _ = v10
                _ = v11
                _ = v9
            } else {
                fmt.Println(fi7, fi8)
                v12 := []int{(fi7 ^ 965), (fi8 ^ fi8), (fi8 + fi8), (fi8 - (fi8 / ((fi7) | 1)))}
                v13 := GenS1{F0: 65535, F1: -32768, F2: 0}
                var v14 float64 = ((1e10 - 3.14159) / (0.1 + ((-2.25 + 0.0) + float64(0.1))))
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi8
        }
        _ = fi7
    }
    if (1000 != -231) {
        v6 = (!true)
        v15 := GenS1{F0: 570, F1: -1, F2: 9223372036854775807}
        _ = v15
    }
    _ = v5
    _ = v6
    return -163
}

func main() {
    var v1 string = ("hello" + "foo bar")
    v2 := []uint32{683}
    var v3 uint16 = 316
    var v4 int8 = -128
    v3 += (v3 / (((v3 * 0)) | 1))
    v16 := GenS1{F0: (v3 ^ v3), F1: 1000, F2: 9223372036854775807}
    if true {
        v17 := GenS2{F0: (v1 == v1), F1: 0, F2: -1}
        switch "" {
        case "":
            switch "foo bar" {
            case "a":
                var v18 uint16 = (v3 % (((v3 + (v3 | 0))) | 1))
                _ = v18
            case "\n\t":
                v1 = v1
            case "":
                v3 -= v3
            default:
                if true {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        var v20 uint = 18446744073709551615
                        if ((v20 | (^(v20))) != (^(v20))) {
                            v1 = "hello"
                        } else {
                            var v21 int8 = (-(v4))
                            v1 = "foo bar"
                            _ = v21
                        }
                        var v22 float32 = -2.25
                        _ = fi19
                        _ = v20
                        _ = v22
                    }
                    switch -868 {
                    case -2147483648:
                        var v23 uint8 = 26
                        _ = v23
                    case 966:
                        v3 += 148
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            v2 = append(v2, 4294967295)
                        case "hello":
                            if false {
                                if (!(false && (v1 >= "foo bar"))) {
                                    switch 32768 {
                                    case 1:
                                        v1 = "\n\t"
                                    case 42:
                                        v1 = v1
                                    case -182:
                                        if ((false || false) || (!(!(0 == 1)))) {
                                            if len(v2) > 0 {
                                                v2[0] = 65535
                                            }
                                            switch 32768 {
                                            case 127:
                                                switch "hello" {
                                                case "a":
                                                    v4 = v4
                                                }
                                            case -2147483648:
                                                v24 := GenS2{F0: true, F1: -1, F2: 27}
                                                _ = v24
                                            case 9223372036854775807:
                                                if len(v2) > 0 {
                                                    v2[0] = 2147483648
                                                }
                                            }
                                        }
                                    default:
                                        if false {
                                            if len(v2) > 0 {
                                                v2[0] = 2147483647
                                            }
                                        } else {
                                            for fi25 := 0; fi25 < 18; fi25++ {
                                                if len(v2) > 0 {
                                                    v2[0] = 1
                                                }
                                                v1 = "hello"
                                                v1 = ""
                                                _ = fi25
                                            }
                                        }
                                    }
                                } else {
                                    v26 := GenS1{F0: v3, F1: 42, F2: 0}
                                    _ = v26
                                }
                                v2 = append(v2, 2)
                            }
                        default:
                            v27 := GenS2{F0: (162 <= 255), F1: 42, F2: 0}
                            _ = v27
                        }
                    }
                } else {
                    var v28 uint32 = 0
                    fmt.Println(v1)
                    _ = v28
                }
            }
        default:
            var v29 uint64 = 574
            _ = v29
        }
        _ = v17
    }
    var v30 float32 = 1e10
    var v31 int8 = (v4 % ((v4) | 1))
    var v32 uint64 = 4294967295
    for fi33 := 0; fi33 < 11; fi33++ {
        _ = len(v2)
        var v34 uint8 = 0
        var v35 uint16 = 0
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := GenS2{F0: true, F1: -163, F2: int64(v32)}
    switch -1 {
    case 2147483647:
        v37 := GenS1{F0: v3, F1: 0, F2: -751}
        _ = v37
    default:
        v31 = -128
    }
    v38 := GenS1{F0: uint16(v32), F1: -305, F2: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
    _ = v38
    _ = v4
}
