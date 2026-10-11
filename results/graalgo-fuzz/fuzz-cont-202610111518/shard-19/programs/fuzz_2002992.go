package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int8 }
type GenS2 struct { F0 int8; F1 uint16; F2 int }
func genFn0(p0 uint16, p1 int32, p2 int8) string {
    fmt.Println(p1, p0, p0)
    var v5 int16 = int16(p0)
    var v6 int64 = 472
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return "\n\t"
}

func main() {
    v1 := []string{"", "", ""}
    var v2 uint64 = 42
    v3 := []uint32{4294967295, 0, 2147483647, 240}
    v4 := make(map[int]uint32)
    var v7 int8 = -128
    fmt.Println(v7)
    v1 = append(v1, ("\n\t" + "hello"))
    v7 += -128
    if (0 < 42) {
        v7 += v7
    } else {
        v8 := GenS1{F0: 1, F1: 42, F2: (v7 * (v7 ^ 109))}
        _ = v8
    }
    fmt.Println(v7)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int16 = -32768
        _ = fi9
        _ = v10
    }
    if ("hello" > (("" + ("foo bar" + "foo bar")) + "\n\t")) {
        var v11 int32 = -2147483648
        _ = v11
    } else {
        v12 := GenS2{F0: v7, F1: 73, F2: 0}
        if (255 >= 1) {
            fmt.Println(v7, v2)
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 int16 = -602
                var v15 bool = false
                var v16 uint8 = 2
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = len(v4)
        } else {
            v2 = 18446744073709551615
        }
        _ = v12
    }
    var v17 int64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
