package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string; F1 uint8; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint }
type GenS5 struct { F0 int16; F1 uint64; F2 int8 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int64) bool {
    v6 := GenS1{F0: 42, F1: 32767}
    var v7 uint32 = 4294967295
    v7 -= (v7 / (((v7 / (((^(v7))) | 1))) | 1))
    v8 := []int{-2147483648, 128}
    switch "\n\t" {
    case "\n\t":
        fmt.Println(p0, v7)
    case "":
        fmt.Println(v7, v7)
    case "hello":
        v7 = 849
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (1 != (p0 * p0))
}
func genFn1(p0 int64, p1 uint64) string {
    v9 := GenS1{F0: 956, F1: 203}
    var v10 string = "foo bar"
    var v11 int64 = p0
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return v10
}

func main() {
    var v1 uint8 = 255
    var v2 uint8 = (^(v1))
    var v3 int16 = -960
    v4 := make(map[string]string)
    v5 := make(map[uint32]int)
    for fi12 := 0; fi12 < 20; fi12++ {
        v3 = v3
        _ = fi12
    }
    v1 += v1
    v13 := GenS2{F0: ((("foo bar" + "") + (("hello" + "hello") + "\n\t")) + "hello"), F1: (v1 + (^(v1))), F2: 76}
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 bool = false
        _ = fi14
        _ = v15
    }
    var v16 uint16 = 511
    var v17 float64 = -2.25
    var v18 int32 = -1
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 int = (fi19 - 32768)
        _ = fi19
        _ = v20
    }
    v3 = (v3 >> ((1) & 63))
    v1 = v1
    switch "a" {
    case "hello":
        v21 := GenS3{F0: 1}
        _ = v21
    case "":
        var v22 string = "hello"
        _ = v22
    default:
        v2 += v2
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        v24 := GenS4{F0: (^(v16)), F1: true, F2: 255}
        v1 = v2
        _ = fi23
        _ = v24
    }
    var v25 int8 = 100
    fmt.Println(v25, v18)
    switch (("\n\t" + (("hello" + "") + "\n\t")) + (((((((("a" + (("foo bar" + "\n\t") + "a")) + "foo bar") + ("a" + "\n\t")) + "hello") + ((("a" + "\n\t") + "hello") + ((("\n\t" + "foo bar") + "\n\t") + (("a" + ("" + "")) + ("foo bar" + ""))))) + "") + "\n\t") + (("foo bar" + (("hello" + "") + "a")) + (("" + (("foo bar" + "hello") + ((("a" + (("hello" + "\n\t") + "foo bar")) + ("\n\t" + "a")) + ("hello" + "")))) + "\n\t")))) {
    case "a":
        switch "" {
        case "\n\t":
            if (18446744073709551615 >= 302) {
                v26 := []int{128, 0}
                delete(v4, ("\n\t" + ((("" + "") + ("a" + "a")) + ("a" + ("hello" + "foo bar")))))
                var v27 uint16 = (v16 / ((v16) | 1))
                _ = v26
                _ = v27
            } else {
                if false {
                    v28 := []int{0, 713, 0, -2147483648}
                    switch "" {
                    case "":
                        fmt.Println(v2)
                    case "hello":
                        v29 := GenS5{F0: 1000, F1: 42, F2: v25}
                        _ = v29
                    case "foo bar":
                        if (true || false) {
                            v30 := GenS6{F0: v2}
                            switch 0 {
                            case 42:
                                var v31 int32 = 42
                                _ = v31
                            default:
                                fmt.Println(v3)
                            }
                            for fi32 := 0; fi32 < 14; fi32++ {
                                v4[("hello" + "")] = "foo bar"
                                _ = fi32
                            }
                            _ = v30
                        } else {
                            if (false && (1e10 != 3.14159)) {
                                v2 = (v2 * v1)
                            } else {
                                switch "hello" {
                                case "hello":
                                    v17 += (((((0.0 + v17) * ((v17 + -2.25) * 1e10)) + (v17 + ((((-2.25 - v17) / 3.14159) - (float64(1.5) - float64((0.0 * float32(((v17 + v17) - ((v17 / v17) + 0.1))))))) / v17))) - 0.0) - float64(3.14159))
                                case "\n\t":
                                    fmt.Println(v18)
                                default:
                                    if (v18 >= 978) {
                                        switch -708 {
                                        case 2147483647:
                                            var v33 uint64 = 42
                                            _ = v33
                                        case -1:
                                            fmt.Println(v2)
                                        default:
                                            if false {
                                                switch 32768 {
                                                case -9223372036854775808:
                                                    var v34 uint64 = 18446744073709551615
                                                    _ = v34
                                                case 650:
                                                    fmt.Println(v17)
                                                }
                                            }
                                        }
                                        v16 += (v16 - 1)
                                        var v35 uint64 = 1
                                        _ = v35
                                    }
                                }
                                var v36 int64 = 42
                                _ = v36
                            }
                            var v37 uint16 = (^(v16))
                            _ = v37
                        }
                    default:
                        if len(v28) > 0 {
                            v28[0] = 9223372036854775807
                        }
                    }
                    var v38 int = 32767
                    _ = v28
                    _ = v38
                } else {
                    delete(v4, "a")
                }
                v39 := []int{-9223372036854775808, 105, 128, 2147483647}
                _ = v39
            }
        default:
            v40 := GenS5{F0: (v3 - -1), F1: uint64(v2), F2: (v25 - v25)}
            _ = v40
        }
    }
    v25 = int8(v1)
    v4[("foo bar" + "a")] = "\n\t"
    delete(v5, 970)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
