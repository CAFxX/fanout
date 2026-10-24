package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 string; F1 int16; F2 uint32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int16 = -674
    var v2 string = ("hello" + (("hello" + ("foo bar" + "foo bar")) + ""))
    v3 := []uint32{2147483647, 2147483648, 2147483647}
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]string)
    var v6 uint32 = 2
    fmt.Println(v1, v2, v6)
    var v7 uint8 = 255
    v1 -= 42
    fmt.Println(v1, v2)
    switch "hello" {
    case "a":
        fmt.Println(v6)
    case "":
        v6 += v6
    case "foo bar":
        switch v2 {
        case "":
            fmt.Println(v7, v1, v1)
        case "foo bar":
            var v8 float64 = (-2.25 - float64((0.1 + float32(0.1))))
            _ = v8
        case "a":
            if false {
                if (!false) {
                    switch v2 {
                    case "foo bar":
                        if ((v2 <= ("hello" + "a")) || false) {
                            if (42 != 255) {
                                switch ((v2 + v2) + ("foo bar" + v2)) {
                                case "a":
                                    switch 1 {
                                    case 189:
                                        for fi9 := 0; fi9 < 19; fi9++ {
                                            v6 -= v6
                                            _ = fi9
                                        }
                                    }
                                }
                                var v10 uint8 = 146
                                _ = v10
                            } else {
                                v2 = ("hello" + v2)
                            }
                            var v11 uint16 = 32767
                            if len(v3) > 0 {
                                v3[0] = (^(v6))
                            }
                            _ = v11
                        }
                    case "\n\t":
                        v3 = append(v3, (^(v6)))
                    default:
                        switch (v2 + v2) {
                        case "\n\t":
                            var v12 float32 = -2.25
                            _ = v12
                        default:
                            for fi13 := 0; fi13 < 5; fi13++ {
                                delete(v4, int(v1))
                                _ = fi13
                            }
                        }
                    }
                    var v14 int = -728
                    delete(v4, int(v7))
                    _ = v14
                }
                fmt.Println(v2, v1, v2)
                if (true && (v2 > (v2 + "foo bar"))) {
                    v15 := GenS1{F0: 42, F1: 0, F2: 1}
                    v2 = v2
                    _ = v15
                } else {
                    var v16 int32 = 42
                    switch -2147483648 {
                    case 32768:
                        if (int16(v16) > v1) {
                            switch 32768 {
                            case -361:
                                v6 += v6
                            case -578:
                                for fi17 := 0; fi17 < 16; fi17++ {
                                    v18 := GenS1{F0: v16, F1: 0, F2: (v7 >> ((263) & 63))}
                                    _ = fi17
                                    _ = v18
                                }
                            case -368:
                                switch "" {
                                case "\n\t":
                                    var v19 bool = false
                                    _ = v19
                                case "":
                                    if (true || false) {
                                        v20 := GenS2{F0: ("foo bar" + (v2 + "foo bar")), F1: (v1 & 32767), F2: (^(v6))}
                                        v16 -= 1
                                        _ = v20
                                    }
                                }
                            default:
                                _ = len(v5)
                            }
                            v1 = 335
                        } else {
                            var v21 bool = (0 != 79)
                            _ = v21
                        }
                    default:
                        for fi22 := 0; fi22 < 6; fi22++ {
                            for fi23 := 0; fi23 < 5; fi23++ {
                                v24 := GenS1{F0: 42, F1: 1, F2: (v7 % ((v7) | 1))}
                                switch int(v1) {
                                case -2147483648:
                                    v6 -= v6
                                case 128:
                                    var v25 uint = 255
                                    _ = v25
                                case 0:
                                    fmt.Println(v7)
                                default:
                                    v7 = 42
                                }
                                v4[403] = (v6 * 42)
                                v16 += int32(v6)
                                _ = fi23
                                _ = v24
                            }
                            var v26 uint = 18446744073709551615
                            v2 = v2
                            _ = fi22
                            _ = v26
                        }
                    }
                    _ = v16
                }
            }
        }
    default:
        if false {
            var v27 float64 = (1.5 / 1e10)
            v27 -= v27
            v6 += v6
            _ = v27
        }
    }
    v28 := GenS1{F0: -2147483648, F1: 18446744073709551615, F2: v7}
    for fi29 := 0; fi29 < 14; fi29++ {
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 uint = 65535
            var v32 float64 = 0.1
            _ = fi30
            _ = v31
            _ = v32
        }
        fmt.Println(v2, v2, v2)
        _ = fi29
    }
    v7 += v7
    var v33 int16 = (v1 * v1)
    if ((0.1 <= 1.5) || false) {
        fmt.Println(v33)
    }
    var v34 uint16 = 2
    if len(v3) > 0 {
        v3[0] = (v6 % ((280) | 1))
    }
    v35 := GenS3{F0: -740}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
