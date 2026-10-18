package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int64; F2 int64 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int64; F1 int8 }
func genFn0() uint16 {
    var v5 uint16 = 1
    v5 += 155
    v5 = v5
    if (1000 >= 970) {
        var v6 float64 = (0.0 / 1e10)
        if true {
            var v7 int16 = -711
            v7 = (v7 >> ((936) & 63))
            _ = v7
        }
        var v8 int32 = -165
        _ = v6
        _ = v8
    } else {
        v5 += 539
        v5 = (v5 * (v5 * (v5 << ((164) & 63))))
    }
    v9 := []int{4294967295, 9223372036854775807, -590, 327}
    _ = v5
    _ = v9
    return 2
}

func main() {
    var v1 uint64 = 704
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    var v4 int16 = 32767
    _ = len(v2)
    v10 := v2[2147483647]
    var v11 uint8 = 0
    switch "a" {
    case "\n\t":
        v11 = v11
    case "hello":
        v12 := GenS1{F0: 4294967295}
        _ = v12
    case "foo bar":
        v13 := GenS1{F0: 65535}
        _ = v13
    default:
        v14 := v3[((((("" + ("" + ("a" + ("" + "")))) + "hello") + "") + "foo bar") + "hello")]
        _ = v14
    }
    var v15 float64 = 0.0
    var v16 int16 = v4
    if (!true) {
        switch (("a" + "a") + (("\n\t" + "a") + "foo bar")) {
        case "hello":
            fmt.Println(v16, v15, v10)
        case "a":
            v17 := GenS2{F0: 18446744073709551615, F1: 9223372036854775807, F2: 42}
            _ = v17
        default:
            v1 -= (v1 & v1)
        }
        if (!false) {
            if true {
                if true {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        for fi19 := 0; fi19 < 1; fi19++ {
                            var v20 int = 669
                            _ = fi19
                            _ = v20
                        }
                        var v21 uint64 = v1
                        var v22 string = ("" + "foo bar")
                        _ = fi18
                        _ = v21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 10; fi23++ {
                        v16 = (v4 | (v16 + (v4 * v4)))
                        var v24 uint = 994
                        _ = fi23
                        _ = v24
                    }
                } else {
                    if (!true) {
                        v25 := GenS3{F0: 290, F1: 42}
                        var v26 bool = (("\n\t" + ("a" + (((("\n\t" + "a") + "hello") + ((("\n\t" + "") + (("a" + ("foo bar" + "hello")) + ("" + "\n\t"))) + "")) + "hello"))) <= "foo bar")
                        _ = v25
                        _ = v26
                    } else {
                        var v27 uint64 = 4294967295
                        v28 := GenS4{F0: 316, F1: 84}
                        _ = v27
                        _ = v28
                    }
                }
            } else {
                var v29 float64 = v15
                v30 := v3["\n\t"]
                _ = v29
                _ = v30
            }
        }
    } else {
        for fi31 := 0; fi31 < 13; fi31++ {
            for fi32 := 0; fi32 < 18; fi32++ {
                fmt.Println(v4)
                fmt.Println(v16)
                v33 := GenS5{F0: 1, F1: 100}
                _ = fi32
                _ = v33
            }
            v16 = -32768
            var v34 uint64 = (v1 - (v1 % ((2) | 1)))
            _ = fi31
            _ = v34
        }
        v4 -= (v4 * (v4 >> ((1) & 63)))
        v35 := GenS3{F0: 154, F1: 817}
        _ = v35
    }
    v36 := GenS3{F0: 0, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
