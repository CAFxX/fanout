package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 bool }

func main() {
    var v1 int = 127
    v2 := make(map[int]bool)
    v3 := []int{(v1 & v1), (v1 & v1)}
    var v4 int32 = 1
    var v5 float64 = 0.1
    v6 := []string{}
    var v7 int32 = v4
    _ = len(v2)
    var v8 uint64 = 1
    if len(v6) > 0 {
        v6[0] = ("foo bar" + ("hello" + "a"))
    }
    switch "foo bar" {
    case "hello":
        v5 = float64((1.5 / 1e10))
    case "\n\t":
        v1 += (v1 ^ v1)
    default:
        var v9 uint8 = 2
        _ = v9
    }
    v8 += (v8 >> ((18446744073709551615) & 63))
    fmt.Println(v8)
    switch ("hello" + "\n\t") {
    case "":
        var v10 uint16 = 2
        _ = v10
    case "a":
        switch (v1 & (v1 * 128)) {
        case 7:
            v11 := GenS1{F0: 128, F1: (v8 | 1), F2: ((!false) || ((0 >= 1) && ((0.1 + (0.0 * float64((1e10 / -2.25)))) > v5)))}
            _ = v11
        case 172:
            _ = len(v2)
        case 4294967295:
            for fi12 := 0; fi12 < 9; fi12++ {
                switch v1 {
                case -219:
                    var v13 int32 = (v7 & (v4 * v4))
                    _ = v13
                }
                for fi14 := 0; fi14 < 15; fi14++ {
                    v1 = fi12
                    v15 := GenS1{F0: 0, F1: (v8 + v8), F2: (!(((v5 + 3.14159) - v5) == v5))}
                    _ = fi14
                    _ = v15
                }
                v1 = v1
                _ = fi12
            }
        }
    }
    v16 := v2[-2147483648]
    fmt.Println(v16)
    if v16 {
        for fi17 := 0; fi17 < 10; fi17++ {
            v2[fi17] = v16
            if len(v6) > 0 {
                v6[0] = ("" + ((("a" + "a") + "foo bar") + ""))
            }
            _ = fi17
        }
        var v18 int32 = -1
        fmt.Println(v18)
        _ = v18
    } else {
        for fi19 := 0; fi19 < 2; fi19++ {
            v1 -= (fi19 | -9223372036854775808)
            _ = fi19
        }
    }
    v5 += (3.14159 / (v5 + 0.0))
    var v20 float32 = (float32(float64(((1.5 - 3.14159) - 1.5))) + float32(float64(-2.25)))
    var v21 float32 = v20
    v1 = -9223372036854775808
    var v22 uint32 = 4294967295
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
