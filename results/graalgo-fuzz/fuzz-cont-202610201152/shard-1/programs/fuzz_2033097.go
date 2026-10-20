package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    var v4 int64 = 42
    var v5 int = -248
    v5 -= (v5 % ((526) | 1))
    if (0.1 == 1e10) {
        if false {
            var v6 uint16 = 65535
            _ = v6
        } else {
            var v7 float64 = 0.0
            switch "\n\t" {
            case "":
                v4 = v4
            default:
                switch "foo bar" {
                case "hello":
                    v8 := []int{(v5 ^ 32767), (v5 * (v5 & v5))}
                    _ = v8
                }
            }
            _ = v7
        }
    } else {
        v9 := GenS1{F0: true, F1: 113}
        _ = v9
    }
    v10 := []int{741, v5, 0, -9223372036854775808}
    _ = len(v3)
    v5 = (v5 >> ((0) & 63))
    v5 = (v5 | (v5 << ((4294967295) & 63)))
    v4 -= v4
    v2[2147483648] = (false || (("hello" + "") == "hello"))
    var v11 int64 = 9223372036854775807
    switch "" {
    case "hello":
        var v12 uint32 = 1
        _ = v12
    case "":
        if (!true) {
            v4 += v4
            var v13 int64 = v11
            _ = v13
        } else {
            for fi14 := 0; fi14 < 10; fi14++ {
                v4 = v11
                switch "foo bar" {
                case "foo bar":
                    var v15 float32 = 1.5
                    _ = v15
                case "a":
                    for fi16 := 0; fi16 < 5; fi16++ {
                        if false {
                            v4 = v11
                        } else {
                            v5 += -153
                            v17 := GenS1{F0: (true && (v4 == 42)), F1: -128}
                            _ = v17
                        }
                        _ = fi16
                    }
                }
                _ = fi14
            }
            v5 = (v5 - (v5 & 128))
            v4 -= v11
        }
    case "foo bar":
        switch "\n\t" {
        case "":
            var v18 uint32 = 2147483647
            _ = v18
        case "hello":
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 int = v5
                _ = fi19
                _ = v20
            }
        }
    }
    var v21 float32 = (((1.5 * 1e10) * float32(1.5)) + (-2.25 / 1.0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
