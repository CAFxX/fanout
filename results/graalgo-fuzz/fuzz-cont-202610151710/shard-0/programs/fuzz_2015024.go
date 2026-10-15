package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int8 }
func genFn0(p0 string, p1 bool, p2 uint) int64 {
    var v6 int8 = 0
    v6 = 0
    var v7 uint32 = 2147483647
    var v8 bool = ((true && p1) && false)
    for fi9 := 0; fi9 < 16; fi9++ {
        switch p0 {
        case "\n\t":
            switch 309 {
            case -2:
                p2 = p2
            case 0:
                p2 += p2
            default:
                p2 = (p2 / ((65535) | 1))
            }
        default:
            v7 = v7
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 float32 = 0.0
    var v2 uint8 = 128
    var v3 uint = 65535
    v4 := make(map[uint32]bool)
    v5 := make(map[string]uint32)
    for fi10 := 0; fi10 < 7; fi10++ {
        switch "a" {
        case "\n\t":
            if (0 <= 2) {
                fmt.Println(v2, v1)
                switch "" {
                case "hello":
                    v1 += (v1 / v1)
                }
            } else {
                v3 += (v3 / (((v3 >> ((v3) & 63))) | 1))
            }
        case "foo bar":
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 int8 = -128
                var v13 uint8 = 1
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            v2 = 169
        }
        _ = fi10
    }
    var v14 float64 = (1e10 / 3.14159)
    v1 += (v1 / -2.25)
    fmt.Println(v3, v14, v14)
    v2 -= (v2 * 128)
    v15 := v4[4294967295]
    for fi16 := 0; fi16 < 14; fi16++ {
        if (false && ((v15 || v15) || v15)) {
            if false {
                if v15 {
                    if v15 {
                        v14 -= (0.1 - v14)
                    }
                    fmt.Println(v15, v2, v15)
                    v17 := GenS1{F0: ("a" + "hello"), F1: 290, F2: -38}
                    _ = v17
                } else {
                    v15 = true
                    v2 = v2
                }
                for fi18 := 0; fi18 < 14; fi18++ {
                    if (!v15) {
                        v19 := GenS1{F0: ("foo bar" + "hello"), F1: 42, F2: 127}
                        if v15 {
                            v5["foo bar"] = 361
                            v2 += v2
                            v20 := []int{(^(fi18)), (fi18 >> (((v3 ^ v3)) & 63)), (^(fi18)), (fi18 + (fi18 * (fi16 | (fi18 << (((v3 & 0)) & 63)))))}
                            _ = v20
                        } else {
                            fmt.Println(v14)
                            var v21 uint32 = 2147483647
                            var v22 uint = 756
                            _ = v21
                            _ = v22
                        }
                        _ = v19
                    }
                    v3 = (^(v3))
                    _ = fi18
                }
                v2 += v2
            }
            v23 := GenS1{F0: "", F1: 2, F2: -1}
            _ = v23
        }
        _ = fi16
    }
    if ((("foo bar" + "foo bar") + "hello") > ("foo bar" + ("" + ("\n\t" + "\n\t")))) {
        for fi24 := 0; fi24 < 15; fi24++ {
            for fi25 := 0; fi25 < 18; fi25++ {
                delete(v5, "")
                _ = fi25
            }
            v26 := GenS1{F0: (("" + ("" + "a")) + ("a" + "foo bar")), F1: 42, F2: -1}
            _ = fi24
            _ = v26
        }
    } else {
        fmt.Println(v2, v2, v14)
    }
    var v27 uint = (v3 & 2)
    fmt.Println(v1, v15)
    if v15 {
        var v28 int16 = -116
        _ = v28
    } else {
        var v29 int = 42
        _ = v29
    }
    fmt.Println(v14, v1, v2)
    fmt.Println(v14, v3)
    for fi30 := 0; fi30 < 5; fi30++ {
        if v15 {
            v2 -= 241
        } else {
            var v31 int64 = 1
            _ = v31
        }
        var v32 uint32 = 197
        _ = fi30
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
