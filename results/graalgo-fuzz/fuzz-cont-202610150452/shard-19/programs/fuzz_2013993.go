package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint = 312
    v2 := []string{"\n\t", ""}
    v3 := make(map[int]uint32)
    v4 := []uint32{}
    var v5 uint8 = 42
    var v6 int64 = 1
    v5 = 201
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: v6, F1: -128}
        if (!false) {
            var v9 int64 = v6
            v10 := GenS2{F0: -1}
            v11 := GenS1{F0: -859, F1: 100}
            _ = v10
            _ = v11
            _ = v9
        }
        if len(v2) > 0 {
            v2[0] = ("hello" + "\n\t")
        }
        _ = fi7
        _ = v8
    }
    v5 = 127
    v12 := GenS1{F0: v6, F1: -1}
    v6 = -1
    for fi13 := 0; fi13 < 8; fi13++ {
        if (!(!true)) {
            fmt.Println(v5, v5)
        }
        v14 := GenS2{F0: -3}
        _ = fi13
        _ = v14
    }
    var v15 int64 = v6
    var v16 string = "a"
    if false {
        switch -9223372036854775808 {
        case 32768:
            var v17 uint64 = 286
            _ = v17
        case 128:
            fmt.Println(v1)
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                if (!(2 > 1)) {
                    v20 := GenS2{F0: 1}
                    var v21 int = (fi19 / (((fi18 & -1)) | 1))
                    _ = v20
                    _ = v21
                }
                v15 = (v6 << ((2) & 63))
                _ = fi19
            }
            var v22 int8 = int8(v15)
            _ = fi18
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
