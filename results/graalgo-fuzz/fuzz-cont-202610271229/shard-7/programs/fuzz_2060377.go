package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int }
type GenS2 struct { F0 string }

func main() {
    var v1 float32 = 0.1
    v2 := []int{128, -522, 2147483647}
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 254
    var v5 int8 = -128
    var v6 uint64 = 0
    var v7 uint16 = 140
    switch 42 {
    case -576:
        var v8 int8 = v5
        _ = v8
    case 2147483647:
        v9 := GenS1{F0: (2 == 54), F1: uint64(v4), F2: 832}
        _ = v9
    case 616:
        switch ("foo bar" + ((("foo bar" + "") + "") + ("hello" + "foo bar"))) {
        case "hello":
            v4 += (v4 - 2)
        }
    }
    if ((!true) && false) {
        var v10 int32 = 42
        _ = v10
    } else {
        v1 = float32(((-2.25 * (((-2.25 * float64((v1 + v1))) * 0.0) * 0.0)) + 0.0))
    }
    var v11 uint8 = 255
    v4 = v4
    for fi12 := 0; fi12 < 17; fi12++ {
        v5 += v5
        for fi13 := 0; fi13 < 8; fi13++ {
            switch ("\n\t" + ("hello" + ("" + ("hello" + "\n\t")))) {
            case "foo bar":
                fmt.Println(v4)
            }
            _ = fi13
        }
        _ = fi12
    }
    v14 := v3[v4]
    v15 := GenS2{F0: (("\n\t" + "") + "foo bar")}
    var v16 bool = (true && false)
    v1 -= v1
    _ = len(v2)
    var v17 int64 = 1
    v18 := GenS1{F0: (((0.1 + 1e10) * 0.0) != 1e10), F1: 793, F2: 32768}
    v7 -= (v7 >> ((4294967295) & 63))
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            v4 = (v14 >> ((1) & 63))
            if (true || v16) {
                switch (fi20 + fi20) {
                case -34:
                    v6 = v6
                default:
                    if (0.1 < 1.5) {
                        var v21 float32 = float32(0.1)
                        v5 = (v5 ^ 14)
                        v11 = (^(v11))
                        _ = v21
                    } else {
                        if (-304 != int16(v11)) {
                            fmt.Println(v5)
                        } else {
                            v17 = (v17 >> ((1) & 63))
                            v11 -= v11
                        }
                        switch "" {
                        case "hello":
                            var v22 int8 = v5
                            _ = v22
                        case "":
                            v6 = v6
                        case "a":
                            v23 := GenS2{F0: ("foo bar" + "\n\t")}
                            v1 -= (((((1e10 - v1) + (0.1 - float32(1.5))) + v1) - 0.0) + ((3.14159 * 1e10) - v1))
                            v5 = v5
                            fmt.Println(v7, v5)
                            _ = len(v3)
                            _ = v23
                        default:
                            if true {
                                if (18446744073709551615 <= 160) {
                                    delete(v3, v14)
                                    v7 = (^(v7))
                                } else {
                                    v24 := GenS2{F0: ("\n\t" + ("\n\t" + ""))}
                                    _ = v24
                                }
                            }
                        }
                    }
                    v3[v4] = (v4 * uint32(v6))
                }
            }
            v6 = v6
            _ = fi20
        }
        if true {
            v25 := GenS2{F0: (("" + "") + "a")}
            switch "foo bar" {
            case "":
                var v26 int8 = 42
                _ = v26
            case "hello":
                var v27 int64 = (v17 / ((v17) | 1))
                _ = v27
            case "a":
                var v28 uint64 = (v6 | (v6 ^ v6))
                _ = v28
            default:
                fmt.Println(v17, v17, v11)
            }
            _ = v25
        }
        _ = fi19
    }
    fmt.Println(v4, v11, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
