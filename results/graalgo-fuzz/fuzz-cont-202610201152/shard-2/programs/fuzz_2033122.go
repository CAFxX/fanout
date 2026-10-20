package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 4294967295
    var v3 float64 = 1e10
    v4 := []string{("foo bar" + ("\n\t" + "\n\t")), "\n\t", (("hello" + ("" + ("" + "a"))) + "hello"), "hello"}
    v2 += 416
    v5 := GenS1{F0: 229, F1: 132, F2: -1}
    v3 -= float64(1.5)
    if false {
        if (1 <= 18446744073709551615) {
            v6 := GenS1{F0: 201, F1: 42, F2: -1}
            _ = v6
        }
        v4 = append(v4, "\n\t")
    }
    var v7 int16 = 215
    v2 -= 258
    v8 := GenS1{F0: 0, F1: -9223372036854775808, F2: (v7 ^ (v7 ^ -655))}
    var v9 bool = (-1 > 9223372036854775807)
    var v10 uint8 = 0
    var v11 uint64 = uint64(v10)
    v4 = append(v4, "hello")
    switch -2147483648 {
    case 338:
        var v12 int32 = 52
        _ = v12
    default:
        if ((4294967295 <= 18446744073709551615) || (v10 != uint8(v11))) {
            v7 -= v7
        }
    }
    switch 32767 {
    case -2147483648:
        var v13 int16 = (v7 & v7)
        _ = v13
    default:
        if len(v4) > 0 {
            v4[0] = "a"
        }
    }
    _ = len(v1)
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 float32 = (3.14159 * (0.1 * 1.5))
        var v16 string = "hello"
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
