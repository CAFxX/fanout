package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int = 0
    var v2 bool = false
    var v3 string = "hello"
    v2 = true
    var v4 int16 = 857
    v5 := []int{}
    switch v1 {
    case 397:
        var v6 uint = 2
        _ = v6
    case -290:
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 14; fi8++ {
                if v2 {
                    switch v1 {
                    case 32768:
                        v4 = (v4 | (v4 << ((42) & 63)))
                    case 4294967295:
                        var v9 uint64 = 649
                        _ = v9
                    default:
                        v2 = (!v2)
                    }
                    v10 := GenS1{F0: v3, F1: 529}
                    _ = v10
                }
                _ = len(v5)
                _ = fi8
            }
            switch fi7 {
            case 9223372036854775807:
                var v11 int16 = (v4 + 1)
                _ = v11
            default:
                v1 -= (^(v1))
            }
            _ = len(v5)
            _ = fi7
        }
    }
    switch 0 {
    case 0:
        v1 -= (v1 / ((v1) | 1))
    case 9223372036854775807:
        v4 = (v4 | (v4 >> ((18446744073709551615) & 63)))
    case 42:
        v3 = (v3 + v3)
    }
    v1 = v1
    v12 := GenS1{F0: v3, F1: 414}
    v1 += 802
    v3 = ""
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS1{F0: v3, F1: 4294967295}
        if (-1 <= 42) {
            v3 = "foo bar"
            switch v1 {
            case -343:
                switch v3 {
                case "hello":
                    v5 = append(v5, v1)
                case "foo bar":
                    for fi15 := 0; fi15 < 15; fi15++ {
                        var v16 float32 = float32(0.0)
                        _ = fi15
                        _ = v16
                    }
                case "a":
                    fmt.Println(v1)
                }
            case -362:
                var v17 uint32 = 879
                _ = v17
            case -450:
                for fi18 := 0; fi18 < 4; fi18++ {
                    v3 = "\n\t"
                    var v19 string = (v3 + v3)
                    switch (v1 - (fi18 + -2147483648)) {
                    case 128:
                        fmt.Println(v2, v2)
                    }
                    _ = fi18
                    _ = v19
                }
            default:
                switch (fi13 * v1) {
                case -81:
                    fmt.Println(v2)
                default:
                    v20 := GenS2{F0: 979, F1: (-818 <= 9223372036854775807)}
                    _ = v20
                }
            }
            for fi21 := 0; fi21 < 5; fi21++ {
                v3 = v3
                _ = fi21
            }
        } else {
            var v22 int8 = -128
            fmt.Println(v22)
            v23 := GenS1{F0: v3, F1: 42}
            _ = v22
            _ = v23
        }
        _ = fi13
        _ = v14
    }
    v4 -= (v4 | v4)
    switch v3 {
    case "\n\t":
        v3 = "a"
    case "hello":
        switch -432 {
        case -492:
            if (-646 > 0) {
                switch 127 {
                case -380:
                    switch v3 {
                    case "a":
                        if (!(3.14159 != ((0.0 - 1e10) / 0.1))) {
                            fmt.Println(v1, v3, v1)
                            v3 = v3
                        }
                    case "\n\t":
                        v3 = ""
                    case "foo bar":
                        var v24 uint16 = 32767
                        _ = v24
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
            } else {
                switch v3 {
                case "foo bar":
                    if true {
                        var v25 uint8 = 127
                        if len(v5) > 0 {
                            v5[0] = 42
                        }
                        _ = v25
                    } else {
                        if (!v2) {
                            v26 := GenS1{F0: "a", F1: 2}
                            _ = v26
                        }
                        if true {
                            for fi27 := 0; fi27 < 11; fi27++ {
                                v4 -= v4
                                v3 = "a"
                                _ = fi27
                            }
                            v5 = append(v5, v1)
                            var v28 float32 = (float32((0.1 - ((3.14159 / 1.5) + 3.14159))) + 3.14159)
                            _ = v28
                        }
                    }
                default:
                    v3 = "foo bar"
                }
            }
        case 2147483647:
            for fi29 := 0; fi29 < 15; fi29++ {
                _ = len(v5)
                _ = fi29
            }
        default:
            var v30 uint8 = 0
            _ = v30
        }
    default:
        var v31 bool = v2
        _ = v31
    }
    fmt.Println(v4, v4)
    v32 := GenS3{F0: (v4 % (((v4 >> ((65535) & 63))) | 1))}
    var v33 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
