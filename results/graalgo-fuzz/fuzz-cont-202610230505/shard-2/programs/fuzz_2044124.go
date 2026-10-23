package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int32 }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{(("\n\t" + ("foo bar" + ((((("" + "") + ("foo bar" + ("a" + "foo bar"))) + "a") + "\n\t") + ("a" + "hello")))) + ""), "hello", "\n\t", "a"}
    var v3 uint = 42
    v4 := make(map[uint32]uint32)
    var v5 float32 = (float32((((-2.25 + 1e10) * float64(float32(float64(0.0)))) + -2.25)) + float32(float64(float32(float64((0.0 / 1.5))))))
    v5 = v5
    if false {
        fmt.Println(v3)
        var v6 int = 128
        _ = v6
    } else {
        var v7 int64 = -9223372036854775808
        _ = v7
    }
    v5 = (v5 * v5)
    v8 := GenS1{F0: -2147483648, F1: -1, F2: -55}
    var v9 float64 = (0.1 / 3.14159)
    v2 = append(v2, ("hello" + ""))
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint16 = 2
        v12 := GenS2{F0: 42, F1: (false && (((("foo bar" + "\n\t") + "\n\t") + (("foo bar" + (("a" + "") + ("\n\t" + ""))) + (("" + "foo bar") + ("hello" + "\n\t")))) > ("foo bar" + ("hello" + ("foo bar" + ("hello" + "a"))))))}
        _ = fi10
        _ = v11
        _ = v12
    }
    v4[260] = 2147483647
    v5 = (v5 * 3.14159)
    fmt.Println(v3, v9, v3)
    var v13 uint16 = 42
    fmt.Println(v3, v13, v5)
    if (v13 <= 908) {
        switch -752 {
        case 42:
            for fi14 := 0; fi14 < 1; fi14++ {
                v5 = -2.25
                if (int64(v13) == 0) {
                    v13 = (v13 + uint16(v3))
                    var v15 uint8 = 0
                    v9 = ((v9 - (v9 + v9)) * 0.0)
                    _ = v15
                } else {
                    _ = len(v1)
                }
                _ = fi14
            }
        case 32767:
            var v16 uint16 = v13
            _ = v16
        case 0:
            var v17 int32 = 42
            _ = v17
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
