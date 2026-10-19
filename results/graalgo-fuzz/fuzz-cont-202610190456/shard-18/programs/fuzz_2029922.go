package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int32; F1 int }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]bool)
    var v3 int64 = -1
    var v4 uint32 = 0
    v5 := GenS1{F0: v4}
    if true {
        v3 = 42
        var v6 int32 = -2147483648
        _ = v6
    } else {
        v4 = (v4 << (((^(v1))) & 63))
        for fi7 := 0; fi7 < 2; fi7++ {
            v1 += 255
            v8 := GenS1{F0: v4}
            var v9 string = "a"
            _ = fi7
            _ = v8
            _ = v9
        }
        v4 += (v4 * (v4 >> ((583) & 63)))
    }
    v10 := GenS1{F0: (v4 >> ((v1) & 63))}
    var v11 float64 = 0.0
    var v12 uint16 = 568
    v11 = ((0.1 - v11) / (v11 / 3.14159))
    v4 += v4
    v13 := GenS1{F0: v4}
    for fi14 := 0; fi14 < 2; fi14++ {
        v3 = v3
        var v15 int8 = 0
        var v16 uint64 = 382
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := []int{-9223372036854775808}
    var v18 int16 = -32768
    if (2147483647 != 32768) {
        switch ("hello" + "\n\t") {
        case "\n\t":
            v2[(^(v4))] = ((2 >= 401) || false)
        case "a":
            v19 := GenS2{F0: false, F1: v12}
            _ = v19
        }
        v20 := GenS1{F0: uint32(v1)}
        fmt.Println(v3)
        _ = v20
    }
    if len(v17) > 0 {
        v17[0] = -991
    }
    if true {
        switch "foo bar" {
        case "a":
            switch "hello" {
            case "\n\t":
                for fi21 := 0; fi21 < 2; fi21++ {
                    v2[2] = (!true)
                    delete(v2, (v4 - (v4 * v4)))
                    for fi22 := 0; fi22 < 7; fi22++ {
                        if false {
                            v18 = v18
                            if (("hello" + "hello") == "hello") {
                                switch ("foo bar" + "a") {
                                case "foo bar":
                                    if false {
                                        var v23 uint = v1
                                        v24 := GenS3{F0: -1, F1: fi21}
                                        if len(v17) > 0 {
                                            v17[0] = (fi21 * (fi22 * (fi22 * (fi21 * fi21))))
                                        }
                                        _ = v23
                                        _ = v24
                                    } else {
                                        switch "foo bar" {
                                        case "a":
                                            fmt.Println(v12, v18, v3)
                                        case "":
                                            fmt.Println(v1, fi22, fi22)
                                        }
                                        _ = len(v17)
                                    }
                                case "\n\t":
                                    fmt.Println(v4)
                                case "hello":
                                    var v25 uint32 = v4
                                    _ = v25
                                default:
                                    v3 = int64(v18)
                                    v3 += int64(v18)
                                }
                                switch ("a" + ("hello" + (((((("foo bar" + "foo bar") + "foo bar") + ("foo bar" + "hello")) + "") + "\n\t") + "\n\t"))) {
                                case "hello":
                                    var v26 int16 = -634
                                    _ = v26
                                }
                            }
                            v4 -= 2147483648
                        }
                        _ = fi22
                    }
                    _ = fi21
                }
            case "foo bar":
                v4 -= (v4 % (((^(v4))) | 1))
            default:
                var v27 uint = 42
                _ = v27
            }
        case "":
            v1 = (v1 * (v1 * v1))
        case "foo bar":
            v2[42] = (!(!true))
        default:
            var v28 int16 = 1000
            _ = v28
        }
    } else {
        var v29 int64 = 0
        v30 := GenS3{F0: -2147483648, F1: 0}
        _ = v29
        _ = v30
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
