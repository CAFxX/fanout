package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int; F2 string }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{((("" + "a") + "a") + "\n\t"), "hello", "foo bar"}
    var v3 uint16 = 1
    var v4 uint8 = 42
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 int = (fi7 + 1)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    fmt.Println(v4)
    v3 -= 493
    v2 = append(v2, "")
    var v9 float32 = 0.0
    var v10 int = 2147483647
    v3 = 42
    v11 := GenS1{F0: -922, F1: 18446744073709551615, F2: 2147483648}
    fmt.Println(v9, v9)
    var v12 uint8 = v4
    v10 += v10
    var v13 float64 = 3.14159
    switch "a" {
    case "":
        if ((true && false) || false) {
            switch v10 {
            case 128:
                v14 := GenS2{F0: 244, F1: (v10 * v10), F2: ("\n\t" + "hello")}
                _ = v14
            default:
                fmt.Println(v3)
            }
            v3 = 833
        }
    case "a":
        v15 := GenS2{F0: uint32(v3), F1: v10, F2: "foo bar"}
        _ = v15
    }
    fmt.Println(v9)
    switch (((("\n\t" + "hello") + "hello") + ((("" + ("" + "foo bar")) + "hello") + "hello")) + "\n\t") {
    case "foo bar":
        v3 += v3
    case "\n\t":
        fmt.Println(v13, v13, v3)
    case "a":
        v9 -= (v9 / 0.1)
    default:
        switch ("foo bar" + ((((("hello" + "a") + (("" + (("a" + ((("a" + "hello") + ("foo bar" + "hello")) + "foo bar")) + "a")) + "a")) + ("\n\t" + ("\n\t" + ("foo bar" + "")))) + "foo bar") + "\n\t")) {
        case "foo bar":
            var v16 int8 = 0
            _ = v16
        case "\n\t":
            for fi17 := 0; fi17 < 15; fi17++ {
                v18 := GenS1{F0: 1, F1: 1, F2: 945}
                _ = fi17
                _ = v18
            }
        case "a":
            switch "" {
            case "foo bar":
                v19 := GenS2{F0: 956, F1: (^(v10)), F2: "hello"}
                _ = v19
            }
        default:
            for fi20 := 0; fi20 < 8; fi20++ {
                fmt.Println(v12, v9)
                _ = fi20
            }
        }
    }
    fmt.Println(v13)
    v10 += 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
