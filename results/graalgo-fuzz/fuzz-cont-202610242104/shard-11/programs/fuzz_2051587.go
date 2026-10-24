package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
func genFn0(p0 uint8, p1 int16, p2 uint16) int8 {
    var v4 string = (("hello" + ("hello" + "a")) + "\n\t")
    p1 += 637
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 36
}

func main() {
    var v1 uint16 = 117
    v2 := make(map[uint32]uint32)
    var v3 string = "hello"
    switch v3 {
    case "":
        var v5 int16 = -365
        _ = v5
    case "hello":
        fmt.Println(v3, v3)
    case "a":
        v3 = v3
    }
    v6 := GenS1{F0: "a", F1: 1}
    v7 := v2[0]
    var v8 uint16 = 32767
    v8 -= (^(v8))
    fmt.Println(v3, v1)
    v8 = 96
    var v9 int16 = 1000
    v8 = (^(v8))
    v7 = 2
    v10 := v2[(^(v7))]
    switch v3 {
    case "a":
        if (742 <= 49) {
            var v11 int = 750
            if (v3 < "\n\t") {
                if (18446744073709551615 == 18446744073709551615) {
                    var v12 float32 = (((-2.25 + (1.5 / (0.0 + (-2.25 - float32((0.0 / 1.0)))))) * (float32((0.1 * 1.5)) - ((-2.25 + -2.25) * 0.0))) * 3.14159)
                    _ = v12
                } else {
                    v13 := v2[v7]
                    v14 := GenS1{F0: v3, F1: 0}
                    _ = v13
                    _ = v14
                }
                fmt.Println(v10, v10, v1)
                v9 = v9
            } else {
                v15 := GenS1{F0: v3, F1: 42}
                _ = v15
            }
            var v16 int64 = 472
            _ = v11
            _ = v16
        }
    case "\n\t":
        var v17 float32 = (3.14159 / 3.14159)
        _ = v17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 uint64 = 4294967295
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
