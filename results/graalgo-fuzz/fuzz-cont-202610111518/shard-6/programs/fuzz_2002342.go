package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint32 = 150
    var v2 int16 = -892
    var v3 int64 = -9223372036854775808
    v4 := []string{}
    var v5 float32 = 0.0
    var v6 int16 = v2
    v7 := GenS1{F0: (v6 - (^(v6))), F1: -27}
    v2 = 1
    switch 565 {
    case 382:
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
    case 42:
        var v8 float64 = 3.14159
        _ = v8
    case 32767:
        for fi9 := 0; fi9 < 8; fi9++ {
            switch ("" + ("a" + (("hello" + "a") + ("" + "\n\t")))) {
            case "a":
                v10 := GenS1{F0: v6, F1: 1}
                _ = v10
            default:
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 int = fi11
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v1 += 797
                    _ = fi11
                    _ = v12
                }
            }
            _ = fi9
        }
    default:
        for fi13 := 0; fi13 < 16; fi13++ {
            switch (fi13 & (^(fi13))) {
            case 9223372036854775807:
                if false {
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v15 := GenS1{F0: v2, F1: -14}
                        var v16 int = fi13
                        var v17 float32 = 1e10
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    var v18 int32 = 595
                    v3 -= v3
                    _ = v18
                } else {
                    _ = len(v4)
                    switch "" {
                    case "hello":
                        v1 += v1
                    default:
                        var v19 int = (fi13 | (fi13 ^ (fi13 + (-(fi13)))))
                        _ = v19
                    }
                }
            default:
                v6 = 17
            }
            _ = fi13
        }
    }
    v5 = v5
    v20 := GenS2{F0: 0}
    v4 = append(v4, "foo bar")
    for fi21 := 0; fi21 < 5; fi21++ {
        v2 = (v2 & (v2 >> ((18446744073709551615) & 63)))
        _ = fi21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 int16 = v2
        for fi24 := 0; fi24 < 8; fi24++ {
            fmt.Println(v2, v1)
            _ = fi24
        }
        fmt.Println(v5, v5)
        _ = fi22
        _ = v23
    }
    if false {
        fmt.Println(v2, v5)
        if (32767 == v2) {
            var v25 int64 = 9223372036854775807
            var v26 float64 = (((3.14159 / (-2.25 / (3.14159 / 3.14159))) - float64(v5)) + ((-2.25 * 1e10) * (3.14159 + (float64(v5) / 1.0))))
            for fi27 := 0; fi27 < 4; fi27++ {
                _ = len(v4)
                v28 := GenS1{F0: 1000, F1: 0}
                for fi29 := 0; fi29 < 19; fi29++ {
                    var v30 int8 = -1
                    v4 = append(v4, ("\n\t" + "\n\t"))
                    _ = fi29
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
            _ = v25
            _ = v26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
