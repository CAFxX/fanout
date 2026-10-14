package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int }

func main() {
    v1 := []uint32{1, 2147483647}
    var v2 uint32 = 0
    var v3 int = 128
    v4 := make(map[uint32]string)
    v3 -= 127
    v5 := GenS1{F0: 42, F1: 32767}
    v3 = v3
    v3 -= v3
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v6 int64 = -1
    for fi7 := 0; fi7 < 1; fi7++ {
        switch (v3 * 2147483647) {
        case -459:
            var v8 int8 = -1
            _ = v8
        case -9223372036854775808:
            if ((v6 - (v6 ^ v6)) == 1) {
                var v9 int = (fi7 << ((172) & 63))
                fmt.Println(v2)
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 14; fi10++ {
            v2 = (^(v2))
            _ = fi10
        }
        _ = fi7
    }
    v11 := GenS1{F0: 8, F1: 1000}
    var v12 float64 = ((3.14159 * 0.0) * -2.25)
    v13 := GenS2{F0: (v3 << ((141) & 63))}
    fmt.Println(v2, v3, v6)
    var v14 uint16 = 2
    v12 += v12
    v15 := GenS2{F0: (-(v3))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
