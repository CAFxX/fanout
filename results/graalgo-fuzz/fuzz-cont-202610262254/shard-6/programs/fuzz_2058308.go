package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint16 }
func genFn0() uint32 {
    var v5 int = 32768
    if false {
        var v6 int64 = 9223372036854775807
        _ = v6
    } else {
        v5 -= v5
    }
    fmt.Println(v5)
    _ = v5
    return 2147483647
}
func genFn1() string {
    switch "a" {
    case "\n\t":
        if true {
            var v7 float64 = 3.14159
            v7 -= 3.14159
            _ = v7
        }
    case "hello":
        v8 := GenS1{F0: -9223372036854775808}
        _ = v8
    }
    v9 := []int{}
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int8 = 100
        v11 = (v11 / ((v11) | 1))
        v11 = v11
        _ = fi10
        _ = v11
    }
    _ = v9
    return "\n\t"
}

func main() {
    var v1 uint64 = 42
    var v2 int64 = -1
    v3 := make(map[int]uint32)
    v4 := []string{}
    switch 118 {
    case 925:
        if false {
            fmt.Println(v2, v1)
        }
    }
    switch (("foo bar" + "a") + ("a" + "")) {
    case "hello":
        if false {
            fmt.Println(v1, v2)
            switch 32768 {
            case 1:
                for fi12 := 0; fi12 < 9; fi12++ {
                    v1 = 18446744073709551615
                    fmt.Println(v1)
                    _ = fi12
                }
            case -683:
                var v13 string = ""
                _ = v13
            case 42:
                switch 128 {
                case 32767:
                    switch -814 {
                    case 42:
                        _ = len(v3)
                    case 248:
                        v2 += -9223372036854775808
                    default:
                        v2 += -9223372036854775808
                    }
                case 2147483647:
                    if true {
                        if false {
                            for fi14 := 0; fi14 < 5; fi14++ {
                                for fi15 := 0; fi15 < 7; fi15++ {
                                    var v16 bool = (65535 >= 2147483647)
                                    _ = fi15
                                    _ = v16
                                }
                                var v17 int32 = 42
                                _ = fi14
                                _ = v17
                            }
                        }
                    }
                case 1:
                    v18 := GenS2{F0: 255, F1: 0}
                    _ = v18
                default:
                    v19 := GenS3{F0: 803}
                    _ = v19
                }
            default:
                switch "" {
                case "foo bar":
                    if ((float64((0.0 - 1e10)) * float64(1e10)) == (0.1 + (float64(1.5) + (0.1 - (1.5 / (-2.25 + (3.14159 + 0.0))))))) {
                        var v20 string = ("\n\t" + ((("a" + ("a" + "hello")) + "hello") + (("hello" + "a") + "foo bar")))
                        v21 := v3[32768]
                        _ = v20
                        _ = v21
                    } else {
                        var v22 int32 = 42
                        _ = v22
                    }
                case "\n\t":
                    fmt.Println(v2, v1)
                case "a":
                    var v23 int32 = 42
                    _ = v23
                }
            }
        } else {
            v2 += v2
            for fi24 := 0; fi24 < 10; fi24++ {
                v25 := GenS2{F0: 255, F1: 104}
                fmt.Println(v2)
                v2 = (v2 + 0)
                _ = fi24
                _ = v25
            }
            _ = len(v3)
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        v1 += v1
    }
    switch (("" + "") + ("hello" + "hello")) {
    case "a":
        var v26 uint64 = v1
        _ = v26
    }
    var v27 int32 = 2147483647
    v27 += (-(v27))
    v1 += (v1 % (((v1 << ((2) & 63))) | 1))
    v27 += (v27 | 798)
    var v28 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
