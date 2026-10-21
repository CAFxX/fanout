package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string; F1 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 42
    v3 := []uint32{2147483647, 65535}
    var v4 bool = ((0.0 * 1e10) == 0.1)
    var v5 uint8 = 1
    var v6 uint = 255
    for fi7 := 0; fi7 < 6; fi7++ {
        v4 = v4
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        var v8 int8 = (v2 / (((v2 & (v2 & v2))) | 1))
        _ = fi7
        _ = v8
    }
    fmt.Println(v5, v5)
    switch "a" {
    case "hello":
        _ = len(v1)
    case "":
        var v9 uint32 = 2147483647
        _ = v9
    case "\n\t":
        fmt.Println(v6, v4, v6)
    }
    switch 2147483647 {
    case -98:
        fmt.Println(v2)
    case 593:
        v4 = (!(!(("" != "") && (128 >= 32768))))
    default:
        v10 := GenS1{F0: 1}
        _ = v10
    }
    v11 := GenS1{F0: -2147483648}
    v6 = (v6 * (v6 ^ 1))
    for fi12 := 0; fi12 < 9; fi12++ {
        switch "foo bar" {
        case "hello":
            var v13 int8 = (v2 ^ (v2 - (v2 + -128)))
            _ = v13
        case "a":
            if true {
                v2 += v2
                var v14 float32 = ((3.14159 / float32((-2.25 + ((float64(-2.25) / 3.14159) + (((float64(0.1) - (1e10 / 1e10)) + -2.25) * 1.5))))) / (-2.25 + 1.5))
                v15 := GenS1{F0: 472}
                _ = v14
                _ = v15
            } else {
                delete(v1, 65535)
                v2 += (v2 ^ (^(v2)))
            }
        case "\n\t":
            var v16 int8 = (^(v2))
            _ = v16
        default:
            if true {
                fmt.Println(v2, fi12, v5)
            } else {
                if ((("" + ("" + ("foo bar" + ("a" + "\n\t")))) > "") || (!((v6 - 65535) != (^(v6))))) {
                    v4 = v4
                    fmt.Println(v4, v6)
                }
                var v17 int32 = 2147483647
                _ = v17
            }
        }
        if (v4 && v4) {
            v2 = v2
            if true {
                _ = len(v1)
            }
            _ = len(v3)
        }
        _ = fi12
    }
    v5 = v5
    var v18 float64 = 1.5
    fmt.Println(v6, v5, v18)
    v19 := GenS1{F0: -1000}
    for fi20 := 0; fi20 < 20; fi20++ {
        if true {
            switch -9223372036854775808 {
            case 405:
                if ("\n\t" <= ("\n\t" + "\n\t")) {
                    var v21 uint8 = (v5 << ((v6) & 63))
                    for fi22 := 0; fi22 < 8; fi22++ {
                        switch fi22 {
                        case 1:
                            v5 += v5
                        case -411:
                            var v23 int16 = 853
                            v3 = append(v3, 1)
                            _ = v23
                        case 128:
                            if (int(v6) != (fi22 & (^(fi20)))) {
                                if (32767 > 42) {
                                    var v24 bool = v4
                                    v5 = v21
                                    v2 = (-(v2))
                                    _ = v24
                                }
                            } else {
                                v25 := GenS2{F0: uint16(v6), F1: v2}
                                _ = v25
                            }
                        default:
                            var v26 int64 = 42
                            _ = v26
                        }
                        _ = fi22
                    }
                    fmt.Println(v5)
                    _ = v21
                }
            case -926:
                v4 = false
            case -2147483648:
                for fi27 := 0; fi27 < 1; fi27++ {
                    var v28 float64 = v18
                    v2 += v2
                    _ = fi27
                    _ = v28
                }
            default:
                for fi29 := 0; fi29 < 7; fi29++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v30 := GenS3{F0: "hello", F1: 2147483647}
                    v5 = (v5 - 127)
                    _ = fi29
                    _ = v30
                }
            }
        } else {
            var v31 bool = false
            v32 := v1[792]
            _ = v31
            _ = v32
        }
        _ = fi20
    }
    v4 = false
    v33 := GenS3{F0: "a", F1: 310}
    var v34 uint = (^(v6))
    if len(v3) > 0 {
        v3[0] = 65535
    }
    v18 += (0.1 * v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
