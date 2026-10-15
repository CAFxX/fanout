package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int }
type GenS2 struct { F0 int; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 uint32 = 2147483648
    var v4 int8 = -128
    var v5 uint16 = 857
    v6 := make(map[int]uint32)
    if (0 >= 42) {
        var v7 string = ""
        _ = v7
    } else {
        for fi8 := 0; fi8 < 18; fi8++ {
            v4 -= 102
            var v9 uint8 = uint8(v3)
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 429, F1: (("hello" + ("\n\t" + "foo bar")) + "a"), F2: 490}
        switch "hello" {
        case "\n\t":
            fmt.Println(v3, v4, v4)
        case "a":
            v11 := GenS1{F0: 395, F1: "", F2: 32768}
            _ = v11
        default:
            switch ("a" + "hello") {
            case "hello":
                _ = len(v2)
            case "\n\t":
                v4 = (v4 >> ((541) & 63))
            default:
                if (-32768 != 282) {
                    v3 = v3
                    if true {
                        switch "foo bar" {
                        case "\n\t":
                            if false {
                                v5 = 2
                                v5 -= v5
                                v12 := GenS1{F0: 4294967295, F1: "foo bar", F2: 32767}
                                _ = v12
                            } else {
                                if true {
                                    v5 = 0
                                    var v13 uint16 = v5
                                    _ = v13
                                } else {
                                    v3 = v3
                                }
                                for fi14 := 0; fi14 < 2; fi14++ {
                                    for fi15 := 0; fi15 < 4; fi15++ {
                                        v5 = (v5 + (v5 / ((v5) | 1)))
                                        _ = fi15
                                    }
                                    v16 := GenS2{F0: 9223372036854775807, F1: 127, F2: 0}
                                    v3 = v3
                                    _ = fi14
                                    _ = v16
                                }
                            }
                        case "a":
                            v17 := []int{}
                            _ = v17
                        default:
                            var v18 int64 = 1
                            _ = v18
                        }
                        v4 -= 70
                    } else {
                        if false {
                            if true {
                                v1["a"] = ("a" + "hello")
                                v19 := []int{-9223372036854775808}
                                _ = v19
                            } else {
                                v2["foo bar"] = "\n\t"
                                v20 := []int{-2147483648, -2147483648, 1}
                                fmt.Println(v5)
                                _ = v20
                            }
                        }
                    }
                } else {
                    var v21 int = 128
                    _ = v21
                }
            }
        }
        _ = v10
    }
    var v22 bool = ((2147483647 < 2147483647) && false)
    v23 := []int{-268, 851, -348, -941}
    v22 = (!v22)
    if ((v4 - (v4 - v4)) > 127) {
        var v24 bool = ((v5 * (v5 | 2)) > 42)
        v4 = (v4 & v4)
        var v25 int32 = -1
        _ = v24
        _ = v25
    }
    var v26 float64 = 1e10
    v27 := GenS2{F0: -905, F1: 36, F2: 42}
    var v28 uint = 890
    v5 -= (v5 - (^(v5)))
    switch "\n\t" {
    case "":
        var v29 uint32 = 869
        _ = v29
    case "\n\t":
        v28 = (^(v28))
    }
    v30 := GenS3{F0: 65535}
    for fi31 := 0; fi31 < 10; fi31++ {
        var v32 uint = (v28 + (v28 + v28))
        _ = fi31
        _ = v32
    }
    v4 = (v4 + v4)
    v3 -= (v3 % ((v3) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
