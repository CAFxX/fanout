package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 string }
type GenS2 struct { F0 int8 }

func main() {
    var v1 float64 = (float64(0.0) / 0.1)
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    v2["foo bar"] = -2147483648
    v1 += 0.0
    var v4 uint = 65535
    v5 := GenS1{F0: -1, F1: 127, F2: "hello"}
    var v6 int8 = 42
    var v7 int8 = 100
    v4 = (v4 + (v4 * (v4 & 1)))
    var v8 int = -2147483648
    switch (v8 / (((v8 % ((v8) | 1))) | 1)) {
    case 0:
        switch -834 {
        case -1:
            v8 = 4294967295
        }
    case 4294967295:
        var v9 float32 = ((0.1 + float32((((3.14159 / float64(float32(((0.1 + (((v1 * 3.14159) - 3.14159) * 0.0)) * (v1 - float64(float32(1e10))))))) * 3.14159) - (v1 - ((3.14159 * (((3.14159 * float64(float32(v1))) + 1e10) - -2.25)) * v1))))) * (-2.25 * ((float32(v1) / 1.0) / 1.0)))
        _ = v9
    }
    var v10 uint16 = 65535
    var v11 uint = 18446744073709551615
    v12 := GenS2{F0: 1}
    switch "foo bar" {
    case "\n\t":
        for fi13 := 0; fi13 < 11; fi13++ {
            for fi14 := 0; fi14 < 16; fi14++ {
                if (0 < 42) {
                    delete(v2, ("hello" + "hello"))
                    v4 -= v11
                    v11 += (v4 / (((v4 ^ (v4 | v11))) | 1))
                }
                _ = fi14
            }
            _ = fi13
        }
    case "":
        var v15 float64 = 0.0
        _ = v15
    case "a":
        var v16 string = ("foo bar" + "foo bar")
        _ = v16
    }
    if true {
        v7 = 127
    } else {
        v10 += (^(v10))
        var v17 int64 = -9223372036854775808
        _ = v17
    }
    delete(v3, (v8 % ((v8) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
