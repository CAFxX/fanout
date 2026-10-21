package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS2 struct { F0 int; F1 uint32; F2 uint16 }

func main() {
    var v1 int8 = 42
    var v2 uint64 = 2
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    var v5 uint64 = (^(v2))
    var v6 string = ("" + ("a" + "foo bar"))
    var v7 bool = (-569 > 184)
    delete(v3, "foo bar")
    v8 := GenS1{F0: 0, F1: -1, F2: 41}
    var v9 float64 = (0.0 * (((float64(0.1) / 1.5) - (-2.25 + 1e10)) * 0.0))
    var v10 uint32 = 0
    v11 := []int{-9223372036854775808, 0, 9223372036854775807}
    if true {
        v11 = append(v11, -9223372036854775808)
        var v12 uint32 = (v10 % (((v10 + 2147483647)) | 1))
        if (false || v7) {
            v1 = (^(v1))
            fmt.Println(v10, v2, v10)
            switch -2147483648 {
            case 1:
                if (!v7) {
                    switch 2147483647 {
                    case 127:
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    default:
                        for fi13 := 0; fi13 < 13; fi13++ {
                            var v14 string = ""
                            if (((fi13 - 4294967295) < -149) || v7) {
                                switch v14 {
                                case "\n\t":
                                    if len(v11) > 0 {
                                        v11[0] = 42
                                    }
                                }
                            }
                            fmt.Println(v10, v6)
                            _ = fi13
                            _ = v14
                        }
                    }
                    v9 = 1e10
                    var v15 string = v6
                    _ = v15
                }
            case 9223372036854775807:
                if true {
                    _ = len(v3)
                } else {
                    var v16 float64 = v9
                    if v7 {
                        var v17 string = v6
                        v12 -= 109
                        switch 9223372036854775807 {
                        case 42:
                            v7 = v7
                        case 127:
                            v12 = v10
                        default:
                            var v18 int64 = 9223372036854775807
                            _ = v18
                        }
                        _ = v17
                    }
                    var v19 int = -392
                    _ = v16
                    _ = v19
                }
            case 0:
                if v7 {
                    if ((v5 ^ v2) == (v5 / ((782) | 1))) {
                        for fi20 := 0; fi20 < 7; fi20++ {
                            v12 = v10
                            var v21 uint8 = 0
                            _ = fi20
                            _ = v21
                        }
                    }
                    var v22 float64 = v9
                    var v23 int64 = -1
                    _ = v22
                    _ = v23
                }
            }
        } else {
            fmt.Println(v9, v9)
            v24 := GenS1{F0: 128, F1: -2147483648, F2: 255}
            switch 87 {
            case -152:
                fmt.Println(v7, v6)
            default:
                var v25 bool = v7
                _ = v25
            }
            _ = v24
        }
        _ = v12
    }
    switch v6 {
    case "foo bar":
        if (!true) {
            _ = len(v11)
            v26 := GenS2{F0: 4294967295, F1: (^(v10)), F2: 2}
            var v27 int = 0
            _ = v26
            _ = v27
        }
    case "":
        fmt.Println(v10, v2)
    case "hello":
        for fi28 := 0; fi28 < 3; fi28++ {
            for fi29 := 0; fi29 < 5; fi29++ {
                v1 = v1
                v10 += (v10 / (((v10 + (v10 >> ((0) & 63)))) | 1))
                var v30 uint64 = v5
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
    }
    v31 := GenS2{F0: 128, F1: (v10 & (^(v10))), F2: 419}
    _ = len(v4)
    var v32 uint8 = 7
    fmt.Println(v7, v6, v9)
    delete(v3, (v6 + (((v6 + (("\n\t" + v6) + (("a" + v6) + (v6 + "")))) + ("foo bar" + ("" + ""))) + v6)))
    for fi33 := 0; fi33 < 18; fi33++ {
        switch (fi33 << ((uint(v5)) & 63)) {
        case -9223372036854775808:
            v34 := GenS1{F0: v32, F1: 0, F2: 190}
            _ = v34
        case -1:
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        case 32768:
            var v35 int32 = -2147483648
            _ = v35
        default:
            v32 -= 199
        }
        _ = fi33
    }
    if len(v11) > 0 {
        v11[0] = 774
    }
    v11 = append(v11, -979)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
