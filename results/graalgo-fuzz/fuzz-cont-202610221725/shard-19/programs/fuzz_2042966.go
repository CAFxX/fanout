package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint }
type GenS2 struct { F0 int32; F1 string }
func genFn0(p0 uint16) uint32 {
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: 18446744073709551615, F1: 930, F2: 4294967295}
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    for fi8 := 0; fi8 < 7; fi8++ {
        p0 -= (p0 | 42)
        if (-1 != -1) {
            if len(v7) > 0 {
                v7[0] = 32767
            }
        } else {
            var v9 bool = true
            var v10 int32 = 1
            var v11 bool = (v9 && ((1e10 >= -2.25) || (((1e10 / float64(-2.25)) - 3.14159) >= (float64(1e10) - -2.25))))
            _ = v10
            _ = v11
            _ = v9
        }
        var v12 int8 = 100
        _ = fi8
        _ = v12
    }
    _ = p0
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 1.5
    var v3 uint64 = 0
    var v4 int = 9223372036854775807
    v3 = v3
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 uint32 = 0
        v15 := GenS2{F0: -2147483648, F1: ("\n\t" + "hello")}
        v3 -= (v3 & v3)
        _ = fi13
        _ = v14
        _ = v15
    }
    v3 -= (v3 & v3)
    if (443 == 32767) {
        v16 := GenS2{F0: 0, F1: "\n\t"}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 3; fi17++ {
            if (true || (false && true)) {
                v2 -= v2
                v18 := GenS1{F0: (v3 >> ((255) & 63)), F1: 65535, F2: 964}
                fmt.Println(v3, fi17)
                _ = v18
            }
            _ = fi17
        }
        v19 := GenS2{F0: -2147483648, F1: "a"}
        _ = v19
    }
    var v20 uint8 = 42
    v4 -= (v4 & v4)
    if false {
        fmt.Println(v4, v2, v20)
        var v21 int = v4
        _ = v21
    } else {
        switch "\n\t" {
        case "\n\t":
            if false {
                for fi22 := 0; fi22 < 14; fi22++ {
                    v23 := GenS2{F0: -323, F1: "a"}
                    if (65535 == 255) {
                        switch (((("a" + ((((("\n\t" + "foo bar") + ("foo bar" + "a")) + "hello") + "\n\t") + ("foo bar" + "\n\t"))) + "foo bar") + (("" + "\n\t") + (("a" + "a") + "foo bar"))) + "hello") {
                        case "":
                            for fi24 := 0; fi24 < 18; fi24++ {
                                switch (^(fi24)) {
                                case -1:
                                    fmt.Println(v4)
                                case 487:
                                    v2 = v2
                                case 42:
                                    v20 += v20
                                default:
                                    v2 = v2
                                }
                                _ = fi24
                            }
                        default:
                            v25 := []int{42}
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 17; fi26++ {
                            var v27 string = ""
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        var v28 uint8 = (v20 & (v20 - 128))
                        delete(v1, genFn0(381))
                        _ = v28
                    }
                    _ = fi22
                    _ = v23
                }
                for fi29 := 0; fi29 < 17; fi29++ {
                    var v30 int32 = -2147483648
                    var v31 int8 = 127
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                var v32 float32 = (-2.25 * 0.1)
                _ = v32
            }
        default:
            var v33 int32 = 42
            _ = v33
        }
    }
    var v34 string = "hello"
    v34 = v34
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v4
}
