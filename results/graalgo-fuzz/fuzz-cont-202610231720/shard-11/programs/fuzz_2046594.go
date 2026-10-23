package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint16 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 255
    v3 := []uint32{42, 1, 65535}
    var v4 int8 = -128
    var v5 uint16 = (^(v2))
    var v6 uint8 = 198
    v7 := GenS1{F0: (!(float32(-2.25) != (3.14159 - -2.25))), F1: -1, F2: 0}
    v8 := GenS1{F0: (v2 <= 65535), F1: 1, F2: v5}
    fmt.Println(v5, v4)
    v9 := GenS2{F0: ("foo bar" + "foo bar")}
    var v10 float32 = (float32((0.1 + float64((-2.25 / 0.1)))) - ((1e10 + 0.0) + float32((float64(float32((0.0 * 1e10))) - 3.14159))))
    var v11 float64 = 3.14159
    var v12 int32 = -91
    delete(v1, 4294967295)
    v13 := GenS2{F0: ""}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
