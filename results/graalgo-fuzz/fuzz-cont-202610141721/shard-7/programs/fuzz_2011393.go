package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = false
    v3 := make(map[string]bool)
    var v4 int = -9223372036854775808
    v5 := []int{(v4 + -523), v4, (v4 * (^(v4)))}
    v4 = (v4 / (((-(v4))) | 1))
    var v6 string = ("\n\t" + "a")
    if (3.14159 != 1.5) {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint = 931
            var v9 int64 = 1
            if true {
                if (0 == int8(v9)) {
                    v10 := GenS1{F0: 42, F1: 42}
                    v2 = false
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 15; fi11++ {
                        v8 -= v8
                        _ = fi11
                    }
                }
            } else {
                var v12 uint = (v8 / ((18446744073709551615) | 1))
                _ = v12
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi13 := 0; fi13 < 1; fi13++ {
            switch "hello" {
            case "hello":
                if v2 {
                    v14 := GenS1{F0: 42, F1: 65535}
                    _ = v14
                }
            case "a":
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 int = 0
                    if (-1 <= 0) {
                        v1[2] = 65535
                        if ((3.14159 == 1e10) || v2) {
                            var v17 int16 = -32768
                            if false {
                                v18 := GenS2{F0: 127, F1: 2}
                                switch v16 {
                                case 2147483647:
                                    fmt.Println(v6)
                                case 814:
                                    var v19 int8 = 17
                                    _ = v19
                                }
                                _ = v18
                            }
                            _ = v17
                        } else {
                            v4 += 109
                        }
                    } else {
                        v20 := GenS3{F0: ((!(0 == 42)) || false), F1: 42}
                        var v21 float64 = 1e10
                        v22 := GenS1{F0: 1000, F1: 255}
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    _ = fi15
                    _ = v16
                }
            default:
                v2 = (1 == 9223372036854775807)
            }
            fmt.Println(fi13)
            _ = fi13
        }
    }
    if (false && (0 < 42)) {
        if v2 {
            v23 := GenS3{F0: false, F1: 118}
            switch "a" {
            case "\n\t":
                v24 := GenS4{F0: 42}
                _ = v24
            case "foo bar":
                if (3.14159 < (((-2.25 - 0.1) / ((1.5 * (1.5 - 0.1)) - float32(-2.25))) + ((1.5 + 1.5) * (float32(float64((1.5 * 3.14159))) - 0.1)))) {
                    if (!true) {
                        switch "foo bar" {
                        case "foo bar":
                            v6 = ""
                        }
                        v6 = (v6 + "foo bar")
                    }
                    v25 := GenS4{F0: 1}
                    _ = v25
                } else {
                    fmt.Println(v2, v6, v6)
                    var v26 float64 = (-2.25 * -2.25)
                    _ = v26
                }
            default:
                v2 = v2
            }
            var v27 string = ""
            _ = v23
            _ = v27
        }
    } else {
        var v28 int8 = -128
        var v29 uint32 = 957
        v28 += (v28 << ((0) & 63))
        _ = v28
        _ = v29
    }
    v5 = append(v5, v4)
    var v30 int = v4
    var v31 int = (v30 ^ (^(v30)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
