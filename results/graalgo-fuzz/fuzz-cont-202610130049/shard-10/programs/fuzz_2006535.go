package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }

func main() {
    var v1 int = 1
    var v2 uint = 202
    var v3 uint64 = 0
    v4 := []uint32{622, 2147483647}
    var v5 float32 = float32(0.0)
    v6 := make(map[uint32]string)
    if (-2147483648 >= 781) {
        _ = len(v4)
    }
    if true {
        v3 -= 4294967295
        v7 := v6[65535]
        fmt.Println(v1, v1, v2)
        _ = v7
    } else {
        var v8 string = (((("hello" + ("" + "")) + "") + "\n\t") + "")
        v9 := GenS1{F0: 0}
        fmt.Println(v5, v3, v8)
        _ = v8
        _ = v9
    }
    var v10 int = v1
    var v11 uint8 = 107
    fmt.Println(v2)
    v4 = append(v4, 42)
    if (true && false) {
        v3 = (v3 * (v3 ^ v3))
    } else {
        for fi12 := 0; fi12 < 16; fi12++ {
            v13 := GenS1{F0: 100}
            v11 -= v11
            delete(v6, 429)
            _ = fi12
            _ = v13
        }
        v5 -= 1.5
        var v14 uint16 = 42
        _ = v14
    }
    v5 = (((v5 + (float32(-2.25) / (1.5 * float32(1.5)))) * 0.0) - 0.0)
    v2 -= (v2 >> ((v2) & 63))
    v15 := GenS2{F0: (-(v1))}
    _ = len(v4)
    v5 = float32(0.1)
    v6[2] = (("foo bar" + "foo bar") + "a")
    fmt.Println(v3, v11, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
