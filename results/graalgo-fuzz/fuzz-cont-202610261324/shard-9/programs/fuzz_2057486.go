package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 bool; F1 int8; F2 int16 }

func main() {
    var v1 float32 = ((1.5 - (-2.25 + -2.25)) + 0.0)
    v2 := []uint32{0, 72, 2147483648, 2147483648}
    var v3 int = -322
    switch "" {
    case "hello":
        switch "hello" {
        case "foo bar":
            v3 -= -9223372036854775808
        case "":
            switch (((("" + "hello") + "\n\t") + "foo bar") + "hello") {
            case "a":
                var v4 int8 = 77
                _ = v4
            case "":
                var v5 uint8 = 128
                _ = v5
            case "\n\t":
                fmt.Println(v1)
            default:
                if false {
                    v3 = 0
                } else {
                    for fi6 := 0; fi6 < 20; fi6++ {
                        for fi7 := 0; fi7 < 1; fi7++ {
                            switch v3 {
                            case 127:
                                switch -535 {
                                case -762:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v8 := GenS1{F0: 128, F1: 435, F2: 32767}
                                    var v9 uint = 18446744073709551615
                                    if len(v2) > 0 {
                                        v2[0] = 348
                                    }
                                    switch (fi7 + -2147483648) {
                                    case -264:
                                        var v10 string = ""
                                        v11 := GenS1{F0: 141, F1: 32767, F2: 32767}
                                        _ = v10
                                        _ = v11
                                    case -9223372036854775808:
                                        if (true && true) {
                                            if (!(!true)) {
                                                if (554 != 65535) {
                                                    if false {
                                                        v1 += float32(1e10)
                                                        var v12 int = v3
                                                        _ = v12
                                                    }
                                                    v3 = fi6
                                                    if (v1 > v1) {
                                                        if true {
                                                            fmt.Println(v1, fi6)
                                                            v3 -= (-(fi7))
                                                            var v13 uint8 = 255
                                                            _ = v13
                                                        } else {
                                                            var v14 uint8 = 128
                                                            _ = v14
                                                        }
                                                        var v15 bool = (100 >= 127)
                                                        v16 := GenS1{F0: 7, F1: 1, F2: 436}
                                                        switch v3 {
                                                        case 32768:
                                                            fmt.Println(v15, v1, v9)
                                                        }
                                                        _ = v15
                                                        _ = v16
                                                    }
                                                } else {
                                                    v1 += 0.1
                                                    var v17 int8 = 42
                                                    _ = v17
                                                }
                                                v1 -= v1
                                            }
                                            var v18 string = ("hello" + (((("foo bar" + "\n\t") + ("\n\t" + "a")) + "\n\t") + "hello"))
                                            if false {
                                                if true {
                                                    if (2147483648 == 4294967295) {
                                                        if len(v2) > 0 {
                                                            v2[0] = 912
                                                        }
                                                        v1 = float32(0.0)
                                                    } else {
                                                        v3 = v3
                                                    }
                                                } else {
                                                    v9 += v9
                                                    v18 = v18
                                                    v19 := GenS1{F0: 2, F1: 255, F2: 0}
                                                    _ = v19
                                                }
                                            } else {
                                                var v20 string = ((v18 + v18) + (("hello" + "") + "\n\t"))
                                                fmt.Println(v18)
                                                var v21 uint16 = 255
                                                _ = v20
                                                _ = v21
                                            }
                                            _ = v18
                                        } else {
                                            v9 -= 0
                                            fmt.Println(v3)
                                            fmt.Println(v3, v3, v3)
                                        }
                                    case 2147483647:
                                        _ = len(v2)
                                    default:
                                        fmt.Println(fi6, fi6, fi6)
                                    }
                                    _ = v8
                                    _ = v9
                                case 4294967295:
                                    var v22 float64 = (3.14159 * 3.14159)
                                    _ = v22
                                default:
                                    v1 -= v1
                                }
                            case 128:
                                var v23 int64 = 9223372036854775807
                                _ = v23
                            case 32768:
                                if false {
                                    var v24 int16 = -461
                                    v25 := GenS1{F0: 146, F1: 915, F2: (v24 - -793)}
                                    _ = v24
                                    _ = v25
                                }
                            }
                            fmt.Println(fi7)
                            v26 := GenS2{F0: 957, F1: 255}
                            _ = fi7
                            _ = v26
                        }
                        var v27 uint16 = 1
                        _ = fi6
                        _ = v27
                    }
                }
            }
        default:
            v28 := GenS1{F0: 2, F1: 42, F2: -32768}
            _ = v28
        }
    case "\n\t":
        for fi29 := 0; fi29 < 18; fi29++ {
            v3 -= (v3 >> ((18446744073709551615) & 63))
            v1 -= v1
            _ = fi29
        }
    }
    var v30 int = 128
    for fi31 := 0; fi31 < 7; fi31++ {
        v2 = append(v2, 0)
        var v32 int8 = -94
        _ = fi31
        _ = v32
    }
    v3 += (^(v3))
    var v33 uint16 = 398
    _ = len(v2)
    if len(v2) > 0 {
        v2[0] = 65535
    }
    v34 := GenS1{F0: 0, F1: (v33 * (v33 >> ((65535) & 63))), F2: 0}
    _ = len(v2)
    fmt.Println(v1, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    v33 -= 90
    if true {
        v35 := GenS3{F0: (1e10 <= (0.0 - 1e10)), F1: -61, F2: 0}
        if false {
            var v36 float64 = -2.25
            fmt.Println(v3, v3)
            _ = v36
        }
        v30 = v30
        _ = v35
    } else {
        v2 = append(v2, 4294967295)
    }
    v37 := GenS2{F0: 42, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v37
}
