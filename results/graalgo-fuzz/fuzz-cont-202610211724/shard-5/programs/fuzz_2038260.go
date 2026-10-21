package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int64; F1 string; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 2147483647
    v3 := []uint32{455, 4294967295}
    var v4 bool = (-1 < -1)
    v5 := make(map[string]string)
    v6 := make(map[int]int)
    switch "" {
    case "\n\t":
        v2 = v2
    case "hello":
        var v7 float64 = (1.5 / -2.25)
        _ = v7
    }
    _ = len(v1)
    if v4 {
        v2 += v2
        v8 := GenS1{F0: v2, F1: 453}
        _ = v8
    } else {
        v9 := GenS2{F0: 518, F1: "a", F2: 2}
        var v10 int8 = 0
        _ = v10
        _ = v9
    }
    var v11 uint64 = 2
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            v4 = ((v11 | (v11 / (((v11 - v11)) | 1))) == 210)
            v14 := GenS2{F0: int64(fi12), F1: "\n\t", F2: 65535}
            _ = fi13
            _ = v14
        }
        v11 = v11
        v11 -= 58
        _ = fi12
    }
    v11 -= v11
    v4 = ((1 < -128) || (!((-1 == 1) && (127 >= 100))))
    for fi15 := 0; fi15 < 2; fi15++ {
        if v4 {
            v16 := v5["foo bar"]
            _ = v16
        } else {
            v17 := GenS2{F0: -685, F1: ("" + "foo bar"), F2: 0}
            _ = v17
        }
        v18 := GenS2{F0: 0, F1: "foo bar", F2: 0}
        v2 = v2
        _ = fi15
        _ = v18
    }
    var v19 uint = 42
    v4 = (!(-291 == 1000))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
