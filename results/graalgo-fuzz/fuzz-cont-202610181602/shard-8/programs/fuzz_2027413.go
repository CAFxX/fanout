package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int }

func main() {
    v1 := []string{"", (((("hello" + "hello") + ("\n\t" + ("foo bar" + "foo bar"))) + "a") + "\n\t"), "\n\t", "hello"}
    var v2 int16 = -1
    var v3 int64 = 9223372036854775807
    v4 := []uint32{}
    var v5 float32 = (0.0 / (0.1 * -2.25))
    var v6 float64 = 3.14159
    var v7 uint = 65535
    if true {
        switch -1 {
        case -327:
            var v8 string = ("a" + "\n\t")
            _ = v8
        case 164:
            var v9 uint8 = 0
            _ = v9
        }
    } else {
        switch "" {
        case "a":
            fmt.Println(v5, v7, v3)
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        v7 -= (v7 + 2)
        var v10 float64 = v6
        _ = v10
    }
    if (!false) {
        if true {
            var v11 string = "a"
            _ = v11
        } else {
            var v12 float32 = 1e10
            v12 = v12
            _ = v12
        }
        v5 += v5
    } else {
        v6 += (((0.0 + v6) + v6) - 1e10)
        switch ("a" + "\n\t") {
        case "":
            for fi13 := 0; fi13 < 18; fi13++ {
                v2 -= 1
                v14 := GenS1{F0: 4294967295, F1: 106, F2: (fi13 << (((v7 % ((v7) | 1))) & 63))}
                _ = fi13
                _ = v14
            }
        case "a":
            _ = len(v4)
        }
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        if true {
            fmt.Println(v7)
        } else {
            var v16 int = 32768
            _ = v16
        }
        _ = fi15
    }
    fmt.Println(v6)
    var v17 uint16 = 0
    var v18 bool = (0 > v7)
    _ = len(v1)
    v19 := GenS1{F0: uint32(v7), F1: 128, F2: -144}
    if v18 {
        v6 += (v6 * float64(float32(v6)))
        for fi20 := 0; fi20 < 15; fi20++ {
            v18 = v18
            _ = fi20
        }
    } else {
        v5 = v5
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
