package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int8) uint32 {
    v4 := []int{}
    if false {
        if false {
            for fi5 := 0; fi5 < 11; fi5++ {
                for fi6 := 0; fi6 < 12; fi6++ {
                    if (42 > (fi5 - -659)) {
                        v7 := GenS1{F0: (fi5 % (((fi5 % ((128) | 1))) | 1)), F1: 18446744073709551615}
                        _ = v7
                    }
                    var v8 uint = uint(p0)
                    v9 := GenS1{F0: (fi6 ^ 4294967295), F1: 42}
                    var v10 int16 = -1
                    _ = len(v4)
                    _ = fi6
                    _ = v10
                    _ = v8
                    _ = v9
                }
                v11 := GenS2{F0: p0}
                var v12 int64 = -9223372036854775808
                _ = fi5
                _ = v11
                _ = v12
            }
            p0 = (p0 - (^(p0)))
        }
        if false {
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := GenS1{F0: fi13, F1: uint64(p0)}
                p1 = p1
                _ = fi13
                _ = v14
            }
            switch (("hello" + "\n\t") + (("a" + ("a" + "foo bar")) + ("foo bar" + "hello"))) {
            case "a":
                for fi15 := 0; fi15 < 4; fi15++ {
                    p1 -= -128
                    for fi16 := 0; fi16 < 7; fi16++ {
                        switch "" {
                        case "\n\t":
                            switch "" {
                            case "a":
                                v17 := GenS1{F0: -9223372036854775808, F1: 4294967295}
                                _ = v17
                            case "hello":
                                var v18 uint8 = 89
                                _ = v18
                            }
                        case "a":
                            if false {
                                p1 += p1
                            }
                        case "":
                            p1 -= (p1 >> ((255) & 63))
                        default:
                            if false {
                                var v19 uint32 = 0
                                _ = v19
                            }
                        }
                        fmt.Println(p1, p1)
                        fmt.Println(fi15, p1, fi15)
                        _ = fi16
                    }
                    p0 = 42
                    _ = fi15
                }
            default:
                fmt.Println(p1)
            }
            var v20 bool = false
            _ = v20
        } else {
            p1 = (p1 >> ((18446744073709551615) & 63))
            p0 = (p0 ^ 32767)
            if false {
                for fi21 := 0; fi21 < 12; fi21++ {
                    v22 := GenS2{F0: (p0 ^ 255)}
                    _ = fi21
                    _ = v22
                }
                p0 = p0
                var v23 float32 = 0.1
                _ = v23
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        var v24 int8 = p1
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        p0 = p0
        var v26 uint32 = 157
        v27 := GenS3{F0: (v26 & (^(v26)))}
        _ = fi25
        _ = v26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 int16 = 1000
    var v2 float32 = 1e10
    var v3 int64 = 9223372036854775807
    var v28 uint32 = 2
    v2 = ((v2 + (-2.25 * v2)) * v2)
    if true {
        switch "\n\t" {
        case "":
            if (18446744073709551615 < 18446744073709551615) {
                fmt.Println(v1, v1)
                v28 = genFn0(42, -1)
                switch 9223372036854775807 {
                case 594:
                    fmt.Println(v28, v2, v3)
                }
            }
        case "hello":
            v29 := []int{1}
            _ = v29
        default:
            var v30 int = 0
            _ = v30
        }
        v28 = (v28 - v28)
        v3 -= (v3 >> ((149) & 63))
    }
    v2 -= float32(((1.5 / (((0.1 + (0.1 * 1.5)) + 1.5) * 1.5)) - (float64(0.1) - 0.0)))
    for fi31 := 0; fi31 < 18; fi31++ {
        v3 = (-(v3))
        var v32 int16 = v1
        _ = fi31
        _ = v32
    }
    v33 := []int{}
    var v34 int16 = 343
    var v35 float32 = (3.14159 / (((1e10 - (float32(float64(v2)) * v2)) * v2) - -2.25))
    if false {
        fmt.Println(v2, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
