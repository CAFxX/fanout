package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 uint16 = 460
    var v2 int64 = -9223372036854775808
    var v3 int64 = v2
    var v4 int16 = 42
    v5 := make(map[string]int)
    var v6 float32 = 0.1
    if ("\n\t" != ("hello" + "a")) {
        v5[("" + "a")] = 166
        var v7 uint64 = 1
        if true {
            var v8 int = 127
            _ = v8
        } else {
            var v9 int = 127
            v7 += (v7 / (((v7 - (v7 + (^(v7))))) | 1))
            v7 = v7
            _ = v9
        }
        _ = v7
    }
    v10 := GenS1{F0: v4}
    v4 -= v4
    var v11 uint8 = 128
    switch "foo bar" {
    case "\n\t":
        v3 -= v2
    case "":
        var v12 int8 = 40
        _ = v12
    case "hello":
        var v13 uint = 255
        _ = v13
    default:
        var v14 float64 = (1.5 - 0.1)
        _ = v14
    }
    switch "\n\t" {
    case "a":
        var v15 uint8 = v11
        _ = v15
    case "":
        v6 += v6
    default:
        if true {
            v16 := []int{42}
            _ = v16
        }
    }
    v3 = 1
    v2 -= v2
    v17 := GenS2{F0: 1, F1: (v11 >> ((800) & 63))}
    fmt.Println(v3, v4, v4)
    if false {
        for fi18 := 0; fi18 < 19; fi18++ {
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 uint16 = v1
                switch (fi18 + (fi18 << ((4294967295) & 63))) {
                case -864:
                    v21 := v5[("" + "foo bar")]
                    _ = v21
                case 32767:
                    fmt.Println(v2)
                default:
                    v6 -= v6
                }
                v22 := GenS1{F0: v4}
                _ = fi19
                _ = v20
                _ = v22
            }
            _ = fi18
        }
        v23 := []int{}
        _ = v23
    } else {
        v1 -= (v1 * v1)
        if (!(v1 > (v1 & (v1 & 648)))) {
            if (false && false) {
                switch "foo bar" {
                case "":
                    var v24 float32 = ((v6 - (v6 * float32(1e10))) + float32((((float64(0.0) + (((1e10 + float64(float32((0.0 / (1e10 + 0.1))))) / (1e10 / float64((((v6 / (((((float32(-2.25) + (v6 - v6)) - -2.25) + 1e10) + v6) * (v6 * ((3.14159 * (v6 - -2.25)) - (float32(-2.25) + ((v6 + v6) - v6)))))) / (v6 * (0.0 - 3.14159))) + v6)))) * 0.0)) / 1.5) - 1e10)))
                    _ = v24
                default:
                    var v25 uint8 = 1
                    _ = v25
                }
            } else {
                v26 := []int{128, -9223372036854775808}
                _ = v26
            }
        } else {
            _ = len(v5)
        }
    }
    v1 += v1
    v27 := []int{}
    fmt.Println(v2)
    if false {
        v28 := v5["foo bar"]
        for fi29 := 0; fi29 < 17; fi29++ {
            fmt.Println(v3)
            if (231 < (v11 * (v11 ^ 127))) {
                var v30 uint64 = 4294967295
                switch ("hello" + "\n\t") {
                case "hello":
                    fmt.Println(v1, v4, v1)
                case "":
                    switch ("" + "\n\t") {
                    case "a":
                        v31 := GenS2{F0: v30, F1: 153}
                        _ = v31
                    case "hello":
                        v28 = 439
                    case "foo bar":
                        var v32 uint32 = 489
                        _ = v32
                    default:
                        var v33 int32 = -1
                        _ = v33
                    }
                case "\n\t":
                    if true {
                        if false {
                            switch int(v30) {
                            case -653:
                                switch "a" {
                                case "\n\t":
                                    v5[(("a" + ("\n\t" + ("a" + ("\n\t" + "\n\t")))) + "\n\t")] = (fi29 | v28)
                                }
                            default:
                                v34 := GenS2{F0: (v30 | (v30 >> ((35) & 63))), F1: v11}
                                _ = v34
                            }
                        }
                    } else {
                        var v35 float32 = (float32(-2.25) * -2.25)
                        _ = v35
                    }
                default:
                    for fi36 := 0; fi36 < 18; fi36++ {
                        v30 += v30
                        _ = fi36
                    }
                }
                _ = len(v5)
                _ = v30
            }
            _ = fi29
        }
        _ = v28
    }
    var v37 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
