package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint64 }

func main() {
    var v1 float32 = (-2.25 * ((((0.0 - 0.0) * (3.14159 * 0.0)) + 1e10) + 1.5))
    var v2 float32 = (0.1 * (3.14159 / 3.14159))
    v3 := []string{(("\n\t" + ("hello" + "foo bar")) + "foo bar"), "hello", (((("" + ("\n\t" + "")) + ("a" + ("\n\t" + (("foo bar" + "a") + (("hello" + "hello") + "foo bar"))))) + ((("a" + "hello") + "a") + (("hello" + "a") + "foo bar"))) + ("" + "hello"))}
    var v4 uint = 18446744073709551615
    var v5 uint64 = 757
    var v6 uint16 = 42
    v4 += (v4 % ((65535) | 1))
    switch "" {
    case "a":
        v1 = v2
    case "hello":
        v7 := GenS1{F0: -9223372036854775808, F1: v5, F2: "\n\t"}
        _ = v7
    default:
        v5 += 0
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v5 -= (v5 - (v5 | (v5 >> ((18446744073709551615) & 63))))
        v5 += 1
        _ = fi8
    }
    var v9 float64 = 1e10
    _ = len(v3)
    v9 += 3.14159
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(v6, fi10)
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 uint64 = (v5 & (v5 & (v5 * (v5 / (((v5 * (v5 * (^(v5))))) | 1)))))
            var v13 int16 = 42
            var v14 int = (fi10 & 42)
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    v15 := GenS2{F0: false}
    var v16 uint64 = (v5 | v5)
    v1 += v2
    var v17 uint8 = 0
    if false {
        switch int(v17) {
        case 128:
            var v18 int8 = 100
            _ = v18
        case 32768:
            var v19 int = 4294967295
            _ = v19
        default:
            _ = len(v3)
        }
    } else {
        if true {
            var v20 int = 0
            v9 += v9
            _ = v20
        } else {
            switch -9223372036854775808 {
            case -165:
                _ = len(v3)
            case -360:
                v3 = append(v3, ("foo bar" + ("a" + "\n\t")))
            default:
                v3 = append(v3, "\n\t")
            }
            v6 = 1
        }
        v21 := GenS1{F0: 0, F1: v5, F2: "a"}
        v22 := GenS3{F0: -9223372036854775808, F1: v5}
        _ = v21
        _ = v22
    }
    if ((uint16(v4) != v6) && false) {
        var v23 string = "a"
        for fi24 := 0; fi24 < 20; fi24++ {
            fmt.Println(v1, v23)
            _ = fi24
        }
        _ = v23
    } else {
        v17 -= 255
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
