package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }

func main() {
    v1 := []string{"", "hello", ("\n\t" + "hello")}
    var v2 int64 = -1
    v3 := make(map[uint32]uint32)
    var v4 bool = (true || true)
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 int64 = (v2 - v2)
            var v8 uint8 = 42
            _ = fi6
            _ = v7
            _ = v8
        }
        fmt.Println(v4)
        for fi9 := 0; fi9 < 1; fi9++ {
            v2 -= -9223372036854775808
            v10 := GenS1{F0: 2, F1: "hello"}
            v4 = (!true)
            _ = fi9
            _ = v10
        }
        _ = fi5
    }
    var v11 float64 = 3.14159
    v2 = v2
    for fi12 := 0; fi12 < 15; fi12++ {
        if (32767 < 198) {
            var v13 int64 = -9223372036854775808
            if (!v4) {
                v14 := GenS1{F0: 411, F1: "a"}
                _ = v14
            } else {
                var v15 float32 = 0.0
                v16 := v3[4294967295]
                _ = v15
                _ = v16
            }
            _ = v13
        }
        _ = fi12
    }
    var v17 int8 = 127
    var v18 float32 = ((3.14159 * float32(float64(((0.0 + float32((v11 - v11))) / ((3.14159 - 3.14159) + 1.5))))) - float32((v11 - 3.14159)))
    v2 += v2
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    _ = len(v1)
    v2 += (v2 ^ (v2 - 42))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
