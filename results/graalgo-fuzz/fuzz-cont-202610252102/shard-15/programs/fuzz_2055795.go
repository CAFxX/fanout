package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int16 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]string)
    v3 := []string{}
    v4 := []uint32{2, 65535, 2147483647, 42}
    var v5 int16 = -32768
    var v6 string = ("a" + "")
    var v7 int = -1
    v8 := GenS1{F0: (-(v7))}
    var v9 uint64 = 0
    var v10 int32 = 42
    v1 = 1e10
    v3 = append(v3, "\n\t")
    var v11 int32 = (v10 >> ((323) & 63))
    v6 = "\n\t"
    var v12 uint8 = 127
    v13 := GenS2{F0: v6, F1: (v5 >> ((268) & 63))}
    var v14 int32 = 1
    v10 -= v14
    delete(v2, v6)
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 int32 = v11
        v17 := GenS2{F0: v6, F1: (v5 / (((v5 >> ((42) & 63))) | 1))}
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(v9)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v10 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
