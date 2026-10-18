package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int }
type GenS2 struct { F0 int16; F1 int64; F2 int32 }
type GenS3 struct { F0 int64; F1 string; F2 string }

func main() {
    v1 := []int{-678}
    var v2 float32 = 0.1
    v3 := make(map[string]uint32)
    var v4 uint16 = 142
    var v5 float64 = ((1.5 / (((1e10 - float64(((v2 + (v2 - float32((0.1 * (-2.25 / ((1.5 * 1e10) * ((((-2.25 - 3.14159) - float64((v2 * v2))) - ((-2.25 + 0.1) - (float64(1.5) + ((0.1 - 3.14159) * 1.5)))) + 0.1))))))) - (v2 * v2)))) - 1e10) - (0.1 * 1e10))) * (float64(1e10) - 1e10))
    switch "\n\t" {
    case "hello":
        var v6 uint32 = 42
        _ = v6
    case "a":
        var v7 int8 = 127
        _ = v7
    }
    if (!(!false)) {
        v8 := GenS1{F0: 1, F1: 162, F2: -2147483648}
        _ = v8
    }
    if (!(1 != 127)) {
        if (18446744073709551615 <= 0) {
            fmt.Println(v5, v2, v4)
            var v9 uint = 2
            fmt.Println(v4, v9)
            _ = v9
        } else {
            fmt.Println(v5, v2, v4)
            for fi10 := 0; fi10 < 19; fi10++ {
                v5 = v5
                _ = len(v3)
                for fi11 := 0; fi11 < 5; fi11++ {
                    var v12 float32 = v2
                    v5 += v5
                    delete(v3, ("a" + ("\n\t" + "\n\t")))
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
        }
        switch 1 {
        case 128:
            v13 := GenS1{F0: -1, F1: 127, F2: 128}
            _ = v13
        case 32768:
            if len(v1) > 0 {
                v1[0] = -9223372036854775808
            }
        case 2147483647:
            if (!(!false)) {
                if false {
                    v14 := GenS2{F0: -547, F1: 1, F2: -2147483648}
                    _ = v14
                } else {
                    v5 = 1.5
                }
            } else {
                for fi15 := 0; fi15 < 7; fi15++ {
                    if true {
                        fmt.Println(v4, v5, fi15)
                        if len(v1) > 0 {
                            v1[0] = fi15
                        }
                    } else {
                        switch ("" + "foo bar") {
                        case "hello":
                            v16 := v3[("hello" + "foo bar")]
                            _ = v16
                        case "":
                            v17 := GenS3{F0: -1, F1: "", F2: "foo bar"}
                            _ = v17
                        default:
                            if ((("hello" + "a") + "hello") <= "a") {
                                var v18 uint64 = 569
                                v19 := GenS2{F0: -739, F1: 42, F2: 1}
                                var v20 uint64 = v18
                                _ = v18
                                _ = v19
                                _ = v20
                            } else {
                                v4 = v4
                            }
                        }
                        if (!(164 >= 745)) {
                            _ = len(v3)
                        }
                    }
                    fmt.Println(fi15, v4, v4)
                    var v21 int64 = -292
                    _ = fi15
                    _ = v21
                }
                v2 -= (-2.25 + 0.1)
                var v22 int32 = 1
                _ = v22
            }
        }
    }
    v2 += float32(v5)
    var v23 uint64 = 33
    for fi24 := 0; fi24 < 9; fi24++ {
        v23 = v23
        switch ("\n\t" + "") {
        case "a":
            v2 = (v2 - (v2 * 0.0))
        case "hello":
            delete(v3, ("hello" + "foo bar"))
        case "foo bar":
            v4 = (v4 << ((65535) & 63))
        default:
            fmt.Println(v2)
        }
        _ = fi24
    }
    v25 := GenS2{F0: 32767, F1: 359, F2: 42}
    switch "foo bar" {
    case "":
        v3["a"] = 42
    case "\n\t":
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 uint32 = 1
            _ = fi26
            _ = v27
        }
    default:
        var v28 uint64 = v23
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
