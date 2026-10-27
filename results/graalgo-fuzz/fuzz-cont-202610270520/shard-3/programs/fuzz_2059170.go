package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint32; F1 bool }

func main() {
    v1 := []string{"foo bar", (("foo bar" + "foo bar") + "foo bar")}
    v2 := make(map[uint32]string)
    v3 := []int{}
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    var v6 uint = 0
    v7 := GenS1{F0: 669}
    switch "\n\t" {
    case "foo bar":
        if true {
            fmt.Println(v6)
            fmt.Println(v6)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v6 = 65535
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 uint8 = 255
                switch ("hello" + "\n\t") {
                case "\n\t":
                    delete(v5, 831)
                case "hello":
                    v9 -= v9
                default:
                    _ = len(v2)
                }
                _ = fi8
                _ = v9
            }
            delete(v2, 795)
        }
    case "hello":
        var v10 int8 = 127
        _ = v10
    default:
        v11 := GenS2{F0: 42, F1: 131}
        _ = v11
    }
    fmt.Println(v6)
    switch ("foo bar" + (("foo bar" + "foo bar") + "foo bar")) {
    case "foo bar":
        delete(v4, 814)
    case "a":
        switch "a" {
        case "":
            if len(v1) > 0 {
                v1[0] = "a"
            }
        default:
            for fi12 := 0; fi12 < 17; fi12++ {
                fmt.Println(v6, v6)
                var v13 int16 = 32767
                v14 := GenS1{F0: 65535}
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    case "hello":
        var v15 uint8 = 2
        _ = v15
    default:
        if (9223372036854775807 != 1) {
            var v16 int8 = 0
            if false {
                v17 := GenS2{F0: 0, F1: 149}
                v18 := v5[220]
                _ = v17
                _ = v18
            } else {
                var v19 bool = ((("" + "foo bar") + "hello") == (("hello" + "hello") + "foo bar"))
                _ = v19
            }
            _ = v16
        }
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        v6 = v6
        var v21 uint8 = 127
        _ = fi20
        _ = v21
    }
    v22 := GenS3{F0: 65535, F1: ((0.0 / (-2.25 - ((1e10 * 1.5) * -2.25))) <= ((1e10 / 1.0) + (((-2.25 * float64(0.1)) - -2.25) * float64(1e10))))}
    var v23 int32 = 42
    v24 := GenS1{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
