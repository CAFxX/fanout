package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint; F2 uint32 }

func main() {
    var v1 float32 = float32(1.5)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[string]string)
    var v5 float32 = ((float32(1.5) + v1) + float32(1e10))
    v5 -= -2.25
    var v6 int = -2147483648
    v7 := []int{v6}
    var v8 int16 = -965
    for fi9 := 0; fi9 < 15; fi9++ {
        _ = len(v4)
        var v10 string = (((("\n\t" + "hello") + ("hello" + "foo bar")) + "hello") + "\n\t")
        var v11 int16 = v8
        _ = fi9
        _ = v10
        _ = v11
    }
    v6 = (v6 << ((42) & 63))
    var v12 uint = 4294967295
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint8 = 117
            v1 = v5
            if (0.0 >= 1e10) {
                v16 := GenS1{F0: (v12 >> ((uint(v15)) & 63))}
                v12 += 18446744073709551615
                v8 = v8
                _ = v16
            }
            fmt.Println(v15, v6)
            _ = fi14
            _ = v15
        }
        for fi17 := 0; fi17 < 18; fi17++ {
            v12 = (v12 * (v12 & v12))
            var v18 int64 = -1
            _ = fi17
            _ = v18
        }
        _ = fi13
    }
    switch "foo bar" {
    case "a":
        v19 := GenS2{F0: 0, F1: (v12 & v12), F2: 540}
        _ = v19
    case "\n\t":
        v8 = (^(v8))
    case "foo bar":
        var v20 int = (v6 % ((-1) | 1))
        _ = v20
    }
    var v21 int64 = 1
    fmt.Println(v8, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
