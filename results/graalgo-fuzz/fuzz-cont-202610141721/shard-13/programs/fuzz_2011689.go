package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
func genFn0(p0 int16, p1 int, p2 string) int16 {
    v4 := []int{}
    var v5 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p0 << ((2) & 63))
}

func main() {
    v1 := []int{461, 4294967295, 32768, 9223372036854775807}
    var v2 int16 = -410
    var v3 int = 462
    fmt.Println(v2, v2)
    var v6 float64 = float64(((1.5 / float32(3.14159)) * 1e10))
    v6 -= v6
    v2 = (v2 ^ 1)
    v6 += v6
    if len(v1) > 0 {
        v1[0] = (v3 + (-(v3)))
    }
    switch "" {
    case "\n\t":
        v7 := GenS1{F0: 1, F1: 18446744073709551615}
        _ = v7
    case "":
        v8 := GenS1{F0: 2147483648, F1: 65535}
        _ = v8
    case "a":
        var v9 float32 = -2.25
        _ = v9
    }
    var v10 int32 = 0
    v3 = v3
    var v11 int64 = -99
    if (!(true && (-128 >= -128))) {
        if true {
            v11 = v11
            for fi12 := 0; fi12 < 15; fi12++ {
                v13 := GenS1{F0: 65535, F1: 441}
                _ = fi12
                _ = v13
            }
        } else {
            fmt.Println(v3, v2, v11)
            switch (v3 - -1) {
            case 21:
                v6 = float64(0.0)
            case -9223372036854775808:
                var v14 bool = ((v10 + -1) != v10)
                _ = v14
            default:
                switch ("\n\t" + "foo bar") {
                case "hello":
                    var v15 int = (v3 % ((int(v10)) | 1))
                    _ = v15
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
}
