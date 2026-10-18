package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }

func main() {
    v1 := []string{"hello", (("" + "hello") + "a"), "hello"}
    var v2 int64 = 0
    var v3 int8 = -128
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]bool)
    var v6 uint = 760
    switch 822 {
    case 528:
        var v7 bool = false
        _ = v7
    case 127:
        var v8 int8 = v3
        _ = v8
    default:
        switch -1 {
        case 32768:
            switch (("" + "a") + (((("\n\t" + ("a" + ("\n\t" + "hello"))) + "\n\t") + (("a" + "foo bar") + "")) + "hello")) {
            case "":
                v1 = append(v1, "")
            case "foo bar":
                var v9 float64 = ((-2.25 - 1e10) / 3.14159)
                _ = v9
            case "hello":
                for fi10 := 0; fi10 < 16; fi10++ {
                    fmt.Println(v6, v6, v3)
                    var v11 int64 = v2
                    _ = fi10
                    _ = v11
                }
            }
        case 129:
            v12 := GenS1{F0: 128, F1: (v2 / ((-1) | 1))}
            _ = v12
        case 9223372036854775807:
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := v4[(fi13 >> ((255) & 63))]
                _ = fi13
                _ = v14
            }
        }
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 int = 2147483647
        switch ("foo bar" + "\n\t") {
        case "hello":
            fmt.Println(v2)
        default:
            var v17 int64 = (v2 - 566)
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v2, v2, v6)
    if (false && false) {
        v18 := GenS1{F0: 148, F1: v2}
        var v19 bool = ((1 > 42) && true)
        v4[-650] = 1
        _ = v18
        _ = v19
    }
    switch 565 {
    case 32768:
        fmt.Println(v6, v6)
    case -2147483648:
        v4[42] = 65535
    case -231:
        var v20 string = "a"
        _ = v20
    default:
        v2 = (v2 % ((v2) | 1))
    }
    fmt.Println(v2, v2, v6)
    fmt.Println(v6, v6)
    v3 = 100
    fmt.Println(v3, v6, v2)
    if false {
        switch "hello" {
        case "\n\t":
            v3 += v3
        case "":
            v6 += 255
        default:
            fmt.Println(v6, v6)
        }
    } else {
        _ = len(v4)
        switch "hello" {
        case "foo bar":
            _ = len(v1)
        case "":
            var v21 uint = 2
            _ = v21
        }
    }
    var v22 uint32 = 65535
    v22 -= 4294967295
    v2 = (v2 + int64(v3))
    var v23 float64 = -2.25
    v4[-935] = uint32(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
