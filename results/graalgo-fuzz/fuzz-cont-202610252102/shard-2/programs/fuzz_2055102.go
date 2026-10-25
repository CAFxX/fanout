package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }

func main() {
    v1 := []uint32{}
    v2 := []string{}
    var v3 int = 1
    var v4 uint8 = 0
    var v5 float32 = (1.5 + (0.0 / 1e10))
    var v6 float64 = 0.1
    v6 -= (3.14159 - -2.25)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = 127
    if true {
        fmt.Println(v5, v6, v6)
        fmt.Println(v4)
        v4 += v4
    } else {
        var v7 uint = uint(v4)
        _ = v7
    }
    v4 += (v4 + 255)
    _ = len(v1)
    var v8 uint64 = uint64(v4)
    var v9 uint = 868
    var v10 int = (-(v3))
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v5, v8)
    default:
        switch "a" {
        case "a":
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := GenS1{F0: (v8 & (v8 ^ 230)), F1: (false || false)}
                if ((9223372036854775807 == 42) || (false && (((((-2.25 * v6) * ((v6 / (v6 * (3.14159 - 0.1))) + v6)) - (v6 * ((v6 / ((v6 / (((v6 + (float64(((float32(v6) / v5) * v5)) - 1e10)) + ((v6 * v6) / ((float64(1e10) + (v6 + ((v6 - v6) * (1.5 - 0.1)))) * v6))) * 1e10)) - v6)) + 0.1))) * (1.5 * v6)) != ((v6 / 1.5) * v6)))) {
                    v8 = (v8 >> ((1) & 63))
                }
                _ = fi11
                _ = v12
            }
        }
    }
    if true {
        var v13 int16 = 42
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 6; fi15++ {
            v6 += (v6 / (0.1 * v6))
            _ = fi15
        }
        switch "foo bar" {
        case "":
            var v16 float32 = v5
            _ = v16
        case "\n\t":
            var v17 string = "\n\t"
            _ = v17
        }
        _ = fi14
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := GenS1{F0: v8, F1: true}
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        fmt.Println(v6, v6, fi20)
        _ = fi20
    }
    switch (("hello" + (("hello" + "\n\t") + "")) + "\n\t") {
    case "hello":
        if (v6 <= v6) {
            switch (v3 % (((v3 + v10)) | 1)) {
            case -1:
                var v21 int16 = 0
                _ = v21
            case -9223372036854775808:
                v6 = 0.1
            case -263:
                if false {
                    v5 += (v5 - (float32(1e10) / v5))
                }
            default:
                _ = len(v1)
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case "":
        fmt.Println(v9, v8)
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = 459
        }
    default:
        var v22 bool = (-273 == 1)
        _ = v22
    }
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
