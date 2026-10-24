package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"foo bar", ("foo bar" + ""), "foo bar"}
    var v3 int16 = 1
    v4 := make(map[uint32]bool)
    var v5 int64 = -1
    v3 = 0
    if len(v2) > 0 {
        v2[0] = ("hello" + "")
    }
    fmt.Println(v5, v3)
    var v6 int = 607
    if false {
        v6 += (v6 >> ((648) & 63))
        if true {
            if (!((!((v6 ^ 764) > (v6 ^ v6))) || false)) {
                for fi7 := 0; fi7 < 20; fi7++ {
                    for fi8 := 0; fi8 < 11; fi8++ {
                        switch "a" {
                        case "foo bar":
                            switch (v6 + (^(fi8))) {
                            case -9223372036854775808:
                                if true {
                                    if (42 < 0) {
                                        var v9 float32 = float32(0.1)
                                        _ = v9
                                    }
                                    var v10 string = "foo bar"
                                    _ = v10
                                }
                            default:
                                if false {
                                    v3 = v3
                                }
                            }
                        default:
                            fmt.Println(fi8)
                        }
                        _ = fi8
                    }
                    fmt.Println(v3, v3)
                    _ = fi7
                }
                v3 += (v3 % ((v3) | 1))
                v3 += v3
            }
        }
    } else {
        var v11 int16 = v3
        _ = v11
    }
    if true {
        if (-2.25 > (float64((1e10 - (1e10 + float32(0.1)))) - 3.14159)) {
            var v12 float64 = (3.14159 - 3.14159)
            fmt.Println(v3)
            v13 := v4[2147483647]
            _ = v12
            _ = v13
        }
    } else {
        v14 := GenS1{F0: 42}
        var v15 int16 = (v3 ^ v3)
        if (!(!false)) {
            switch "\n\t" {
            case "":
                fmt.Println(v6, v15, v3)
            default:
                if (255 <= 150) {
                    v5 = (-(v5))
                } else {
                    if true {
                        if (!false) {
                            var v16 uint16 = 0
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v15, v5)
                            case "hello":
                                switch (-(v6)) {
                                case -2147483648:
                                    v17 := GenS1{F0: 842}
                                    _ = v17
                                default:
                                    v6 -= 2147483647
                                }
                            }
                            _ = v16
                        }
                        var v18 int8 = 0
                        v19 := GenS2{F0: 0, F1: 127}
                        _ = v18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 17; fi20++ {
                        v21 := GenS1{F0: v5}
                        var v22 int8 = -128
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
            }
        }
        _ = v14
        _ = v15
    }
    _ = len(v2)
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
