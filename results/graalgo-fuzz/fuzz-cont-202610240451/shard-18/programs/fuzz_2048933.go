package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 string }
func genFn0(p0 int64, p1 uint8) uint32 {
    v5 := []int{-749, 32768, 128, 128}
    v6 := GenS1{F0: 1, F1: 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 string, p1 uint64) int64 {
    p1 = 1
    if ((true || (!true)) || false) {
        var v7 uint = 194
        var v8 string = "\n\t"
        _ = v7
        _ = v8
    } else {
        var v9 int32 = -782
        _ = v9
    }
    var v10 uint = 501
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    var v4 uint = 704
    var v11 uint = v4
    for fi12 := 0; fi12 < 4; fi12++ {
        switch ("\n\t" + ("hello" + (("\n\t" + ("\n\t" + "a")) + "\n\t"))) {
        case "foo bar":
            var v13 float64 = 1.5
            _ = v13
        case "hello":
            switch ("\n\t" + (("\n\t" + ("foo bar" + (("a" + "hello") + "\n\t"))) + "a")) {
            case "\n\t":
                switch 127 {
                case 828:
                    _ = len(v1)
                case -1:
                    v11 = v4
                case -2147483648:
                    if ((false && false) || false) {
                        fmt.Println(fi12)
                    } else {
                        if (!true) {
                            for fi14 := 0; fi14 < 3; fi14++ {
                                var v15 uint64 = 0
                                _ = len(v1)
                                switch "\n\t" {
                                case "hello":
                                    v1 = append(v1, 4294967295)
                                }
                                _ = fi14
                                _ = v15
                            }
                        }
                        v1 = append(v1, 2)
                    }
                default:
                    v4 -= (v11 >> ((v11) & 63))
                }
            default:
                v11 += 4294967295
            }
        case "a":
            v16 := GenS1{F0: genFn0(int64(fi12), 243), F1: (fi12 - (fi12 % ((fi12) | 1)))}
            _ = v16
        default:
            var v17 uint64 = 42
            _ = v17
        }
        v4 = (v11 ^ v11)
        _ = fi12
    }
    var v18 float64 = 0.0
    if false {
        v4 += (v11 / ((1) | 1))
    }
    v19 := GenS2{F0: 4294967295}
    fmt.Println(v4, v18)
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS2{F0: 4294967295}
        var v22 uint8 = 2
        _ = fi20
        _ = v21
        _ = v22
    }
    switch ("" + ((("hello" + "hello") + ("hello" + "foo bar")) + ((((((("" + "foo bar") + "") + "\n\t") + "hello") + ("" + ("hello" + ""))) + "hello") + "a"))) {
    case "":
        switch "\n\t" {
        case "foo bar":
            v18 -= 3.14159
        case "":
            v23 := GenS1{F0: 2147483647, F1: -432}
            _ = v23
        case "a":
            delete(v2, 968)
        default:
            v4 = 255
        }
    case "foo bar":
        _ = len(v1)
    default:
        switch "a" {
        case "a":
            var v24 float32 = 0.0
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 int64 = 0
        var v27 int32 = 217
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint32 = 42
    v28 = v28
    v28 += 65535
    fmt.Println(v28, v28, v28)
    switch 9223372036854775807 {
    case 32768:
        switch 127 {
        case -1:
            var v29 int32 = 42
            _ = v29
        case -9223372036854775808:
            for fi30 := 0; fi30 < 14; fi30++ {
                for fi31 := 0; fi31 < 11; fi31++ {
                    v32 := GenS3{F0: -2147483648, F1: "foo bar"}
                    v11 -= 4294967295
                    var v33 int8 = 0
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                var v34 float32 = 0.0
                _ = fi30
                _ = v34
            }
        }
    }
    var v35 int32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v4
}
