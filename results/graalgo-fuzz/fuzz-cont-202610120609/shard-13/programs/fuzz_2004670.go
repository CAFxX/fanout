package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]string)
    var v3 uint32 = 756
    var v4 int = -311
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            v3 = v3
            var v7 string = (("hello" + "foo bar") + "\n\t")
            fmt.Println(v4, fi6, fi5)
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    if ("foo bar" <= ((("foo bar" + (("hello" + ("foo bar" + ("foo bar" + "hello"))) + ("\n\t" + "foo bar"))) + "foo bar") + (("" + "foo bar") + ("hello" + ("" + "hello"))))) {
        var v8 float32 = (1.5 * 1.5)
        var v9 uint32 = v3
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 45, F1: v1}
    var v11 uint16 = 42
    fmt.Println(v11, v11, v3)
    if true {
        v4 = v4
        switch v4 {
        case 964:
            for fi12 := 0; fi12 < 2; fi12++ {
                v2[(v3 ^ (v3 * v3))] = ("\n\t" + "a")
                var v13 uint32 = v3
                _ = fi12
                _ = v13
            }
        }
        delete(v2, v3)
    }
    v4 -= (-(v4))
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 6; fi15++ {
            v1 += 2147483647
            _ = fi15
        }
        switch 580 {
        case 9223372036854775807:
            var v16 float32 = ((((float32(3.14159) - float32(3.14159)) + -2.25) + (0.1 - 1e10)) / (((((1e10 - 1e10) + -2.25) - 1.5) * 0.0) + (3.14159 / ((3.14159 * 0.1) * ((3.14159 * float32(3.14159)) / 1.0)))))
            _ = v16
        default:
            var v17 int8 = 0
            _ = v17
        }
        _ = fi14
    }
    var v18 uint = 42
    v3 -= (v3 / ((uint32(v18)) | 1))
    v1 -= v1
    v2[v3] = ("\n\t" + "\n\t")
    var v19 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
