package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 string; F2 uint32 }

func main() {
    var v1 float32 = float32(1.5)
    v2 := []string{("foo bar" + "\n\t"), "", "\n\t", ""}
    var v3 uint8 = 1
    var v4 bool = true
    v5 := []uint32{42}
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: v3}
    var v8 int16 = 1000
    _ = len(v2)
    switch 42 {
    case -1:
        v4 = (int(v3) > -273)
    case 42:
        var v9 string = ("foo bar" + "\n\t")
        _ = v9
    case -2147483648:
        v10 := GenS2{F0: 42, F1: ("\n\t" + ("" + "")), F2: 2147483647}
        _ = v10
    }
    v11 := GenS1{F0: v3}
    v3 += v3
    v12 := GenS1{F0: v3}
    fmt.Println(v1)
    v4 = (v3 != (v3 / (((v3 * 255)) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
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
