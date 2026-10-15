package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int16; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS5 struct { F0 int; F1 uint32; F2 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 4294967295
    var v3 uint64 = 2
    v4 := []int{1, -2147483648, 9223372036854775807}
    var v5 uint8 = 1
    v3 = (v3 % ((v3) | 1))
    if (!true) {
        var v6 int64 = 0
        fmt.Println(v3)
        v7 := GenS1{F0: v2}
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: (v2 % (((v2 - v2)) | 1)), F1: (("\n\t" + ((("foo bar" + "\n\t") + "") + ("" + "a"))) + ((((((("a" + "") + "\n\t") + "foo bar") + ("" + "\n\t")) + "foo bar") + "a") + ""))}
    v5 += v5
    v9 := GenS2{F0: v2, F1: (((("" + "\n\t") + "foo bar") + "\n\t") + ((("foo bar" + ("\n\t" + ("a" + (("" + "hello") + "hello")))) + "a") + "\n\t"))}
    _ = len(v4)
    if false {
        for fi10 := 0; fi10 < 3; fi10++ {
            if false {
                for fi11 := 0; fi11 < 20; fi11++ {
                    _ = len(v4)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi11
                }
                fmt.Println(v3, v5, v5)
            }
            _ = fi10
        }
        fmt.Println(v3)
        fmt.Println(v3, v5, v5)
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v5 -= 149
    for fi12 := 0; fi12 < 5; fi12++ {
        if len(v4) > 0 {
            v4[0] = fi12
        }
        if (true && true) {
            var v13 int16 = 988
            v14 := GenS1{F0: (v2 ^ v2)}
            v1[2147483648] = 32768
            _ = v13
            _ = v14
        } else {
            if (((1 == 42) && false) && true) {
                switch "a" {
                case "foo bar":
                    for fi15 := 0; fi15 < 15; fi15++ {
                        var v16 int16 = 0
                        _ = fi15
                        _ = v16
                    }
                case "a":
                    v17 := GenS3{F0: 0, F1: 42, F2: 2}
                    _ = v17
                default:
                    if (((0.1 + 1e10) / (float64(1e10) / (-2.25 / 0.1))) == (3.14159 / (float64(1.5) - float64(3.14159)))) {
                        v4 = append(v4, fi12)
                        v5 -= v5
                        var v18 float64 = ((1e10 * 0.1) + (float64(0.1) * -2.25))
                        _ = v18
                    } else {
                        v19 := GenS4{F0: uint8(v2), F1: 65535, F2: 32767}
                        _ = v19
                    }
                }
                var v20 int16 = 42
                _ = v20
            } else {
                v2 = (v2 - v2)
                switch "" {
                case "hello":
                    if false {
                        v3 = v3
                        var v21 uint16 = 24
                        if true {
                            for fi22 := 0; fi22 < 8; fi22++ {
                                if true {
                                    v2 = (v2 - v2)
                                    var v23 uint8 = v5
                                    _ = v23
                                }
                                var v24 float32 = (0.0 + 0.0)
                                _ = fi22
                                _ = v24
                            }
                        } else {
                            for fi25 := 0; fi25 < 10; fi25++ {
                                var v26 int32 = 305
                                v3 += (v3 + v3)
                                v27 := v1[0]
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        }
                        _ = v21
                    }
                case "\n\t":
                    var v28 uint8 = v5
                    _ = v28
                default:
                    if true {
                        var v29 uint32 = 0
                        _ = v29
                    } else {
                        v30 := GenS1{F0: v2}
                        fmt.Println(v2, v5)
                        _ = v30
                    }
                }
                v3 += v3
            }
        }
        _ = fi12
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        v32 := GenS5{F0: (fi31 | fi31), F1: 42, F2: 523}
        _ = fi31
        _ = v32
    }
    v3 += (v3 ^ v3)
    v2 += v2
    v2 += v2
    v3 = (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
