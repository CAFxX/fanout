package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
func genFn0(p0 int8, p1 uint8, p2 uint32) int32 {
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 uint = 65535
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    if (!false) {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int = (-(fi8))
            fmt.Println(p2, v9)
            var v10 float64 = float64(3.14159)
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 float32 = (3.14159 * 0.0)
        _ = v11
    } else {
        v12 := GenS1{F0: 255, F1: false}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 831
}

func main() {
    var v1 int = 32767
    var v2 bool = false
    v3 := []string{("" + ""), "foo bar", "foo bar", ("a" + "\n\t")}
    var v4 int16 = -283
    if len(v3) > 0 {
        v3[0] = ("foo bar" + (("hello" + "hello") + "foo bar"))
    }
    if (0 <= 996) {
        if v2 {
            v13 := GenS1{F0: 18446744073709551615, F1: (42 < 118)}
            _ = v13
        }
        fmt.Println(v1, v2)
        fmt.Println(v2, v1)
    }
    if len(v3) > 0 {
        v3[0] = ((("" + ("foo bar" + "")) + "hello") + ("hello" + ("a" + "")))
    }
    _ = len(v3)
    _ = len(v3)
    v1 += (v1 ^ (v1 / ((-985) | 1)))
    switch "a" {
    case "a":
        v14 := GenS1{F0: 2, F1: ((0 <= 9223372036854775807) && v2)}
        _ = v14
    case "hello":
        var v15 float64 = 1e10
        _ = v15
    default:
        for fi16 := 0; fi16 < 14; fi16++ {
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 int16 = 0
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    }
    fmt.Println(v2, v2)
    var v19 int = v1
    v1 -= 159
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
