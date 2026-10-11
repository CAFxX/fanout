package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint }

func main() {
    var v1 bool = (false && (!((255 > 143) || false)))
    var v2 uint32 = 2147483648
    v3 := []string{(("hello" + "a") + "")}
    var v4 uint32 = v2
    switch 959 {
    case 2147483647:
        for fi5 := 0; fi5 < 9; fi5++ {
            switch "foo bar" {
            case "\n\t":
                v2 += 65535
            default:
                v6 := GenS1{F0: 1, F1: 2, F2: 2}
                _ = v6
            }
            _ = fi5
        }
    default:
        v1 = (((9223372036854775807 >= -827) || (0 < 1000)) && true)
    }
    v7 := GenS1{F0: 127, F1: 2, F2: 18446744073709551615}
    v1 = (float64((((1e10 * 1e10) + 3.14159) * float32(0.1))) != 0.1)
    v3 = append(v3, ((("" + "foo bar") + ("\n\t" + ("" + "foo bar"))) + (("a" + "foo bar") + "hello")))
    v1 = false
    if (1 >= 711) {
        fmt.Println(v2, v4)
    } else {
        if len(v3) > 0 {
            v3[0] = "a"
        }
        fmt.Println(v4)
    }
    var v8 float32 = -2.25
    v8 += v8
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 int = fi9
        _ = len(v3)
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 52, F1: 4294967295}
    var v12 int16 = 42
    v3 = append(v3, (("foo bar" + "hello") + "a"))
    switch int(v12) {
    case 32767:
        var v13 float64 = (float64(v8) - 1.5)
        _ = v13
    default:
        var v14 float32 = (1.5 / (0.1 / (((-2.25 - ((v8 - v8) / 1e10)) - (3.14159 * float32(3.14159))) / float32((float64((float32(float64(float32(float64(v8)))) * 3.14159)) + float64(v8))))))
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
