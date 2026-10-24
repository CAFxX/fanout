package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int16 }

func main() {
    var v1 uint = 42
    v2 := []uint32{0, 2, 65535, 65535}
    var v3 int = 127
    v4 := make(map[string]int)
    var v5 int16 = 1
    v6 := make(map[uint32]string)
    var v7 string = "hello"
    v5 += v5
    v3 = v3
    fmt.Println(v3)
    if (!(true || false)) {
        for fi8 := 0; fi8 < 6; fi8++ {
            v5 += (-(v5))
            var v9 uint64 = 42
            _ = fi8
            _ = v9
        }
    }
    var v10 string = v7
    v11 := GenS1{F0: (v3 >> (((v1 * v1)) & 63)), F1: v10, F2: v5}
    delete(v6, 42)
    var v12 float64 = 0.1
    v7 = (((("" + "") + ("hello" + (((v7 + ("foo bar" + (v10 + (v10 + (v10 + ((("" + v7) + v10) + (v10 + v10))))))) + (v7 + v7)) + (v7 + (v10 + v10))))) + v10) + v7)
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint16 = 65535
        var v15 uint32 = uint32(v1)
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
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
}
