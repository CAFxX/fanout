package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }

func main() {
    v1 := []uint32{}
    var v2 int32 = 296
    var v3 uint32 = 976
    v4 := []uint32{v3, 509}
    var v5 int64 = 9223372036854775807
    v6 := make(map[uint32]int)
    for fi7 := 0; fi7 < 6; fi7++ {
        v4 = append(v4, 4294967295)
        fmt.Println(v3)
        _ = fi7
    }
    fmt.Println(v2, v2, v2)
    v5 = v5
    switch "hello" {
    case "foo bar":
        v2 = (v2 * (-(v2)))
    case "a":
        switch ("" + "a") {
        case "foo bar":
            if (false || true) {
                fmt.Println(v3, v5)
            }
        case "\n\t":
            v8 := GenS1{F0: -6, F1: ((v2 < (v2 & (v2 % ((v2) | 1)))) || true)}
            _ = v8
        case "a":
            _ = len(v1)
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case "hello":
        v3 -= v3
    }
    fmt.Println(v2)
    for fi9 := 0; fi9 < 14; fi9++ {
        _ = len(v4)
        _ = fi9
    }
    _ = len(v6)
    v3 = 42
    v10 := GenS1{F0: -128, F1: (658 > 255)}
    switch int(v2) {
    case 84:
        var v11 float64 = 0.0
        _ = v11
    case -393:
        fmt.Println(v2, v3, v3)
    case 32767:
        var v12 bool = true
        _ = v12
    }
    if (v3 <= v3) {
        var v13 uint8 = 127
        if false {
            for fi14 := 0; fi14 < 3; fi14++ {
                v2 += 0
                _ = fi14
            }
            var v15 float64 = 0.0
            _ = v15
        }
        v3 = 1
        _ = v13
    } else {
        fmt.Println(v5)
        var v16 int = 128
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := GenS1{F0: -71, F1: (true && (((("" + ("hello" + ("\n\t" + "hello"))) + "\n\t") + "a") == "a"))}
        v19 := GenS1{F0: 1, F1: false}
        if ("\n\t" >= ("foo bar" + ("\n\t" + "foo bar"))) {
            for fi20 := 0; fi20 < 19; fi20++ {
                fmt.Println(v3, fi17)
                v2 += (v2 + v2)
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
                _ = fi20
            }
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
