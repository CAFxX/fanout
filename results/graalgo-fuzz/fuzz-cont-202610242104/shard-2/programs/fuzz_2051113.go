package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint32; F1 int }
func genFn0() uint {
    v4 := []int{-190, 111, 128}
    var v5 float32 = 1.5
    _ = v4
    _ = v5
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 14; fi7++ {
        if true {
            var v8 int8 = 105
            fmt.Println(v8, v8)
            delete(v1, 1)
            _ = v8
        }
        switch (((("" + (("hello" + "a") + ((("\n\t" + "foo bar") + "hello") + "foo bar"))) + "a") + "") + "hello") {
        case "\n\t":
            var v9 uint32 = 422
            _ = v9
        case "foo bar":
            var v10 uint = 18446744073709551615
            _ = v10
        }
        _ = fi7
    }
    v6 = v6
    var v11 float32 = (3.14159 + (1e10 * 1e10))
    if true {
        v6 -= 22
        v11 += v11
        switch 630 {
        case 4294967295:
            v6 = v6
        case -138:
            for fi12 := 0; fi12 < 4; fi12++ {
                v11 -= v11
                _ = len(v3)
                var v13 int16 = 1
                _ = fi12
                _ = v13
            }
        case 42:
            switch ("" + "\n\t") {
            case "hello":
                for fi14 := 0; fi14 < 7; fi14++ {
                    v6 += 1
                    _ = fi14
                }
            case "\n\t":
                var v15 uint32 = 42
                _ = v15
            case "foo bar":
                fmt.Println(v11, v6)
            default:
                var v16 int8 = 0
                _ = v16
            }
        }
    }
    v6 += (^(v6))
    v11 -= v11
    fmt.Println(v11)
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 uint64 = 0
        if true {
            v19 := GenS1{F0: (!((float64((((0.1 / v11) - (-2.25 * float32(0.1))) / (0.1 + (3.14159 / 0.1)))) + (0.0 / (((1e10 / 3.14159) + -2.25) / 1.0))) >= 3.14159)), F1: -2147483648}
            v20 := GenS2{F0: 0, F1: 2147483647}
            _ = v19
            _ = v20
        }
        var v21 uint32 = 2147483648
        _ = fi17
        _ = v18
        _ = v21
    }
    switch -9223372036854775808 {
    case -2147483648:
        v6 = v6
    }
    v22 := GenS2{F0: 42, F1: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v6
}
