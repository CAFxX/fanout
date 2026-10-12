package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int32; F1 uint64; F2 int16 }

func main() {
    v1 := []uint32{}
    var v2 uint = 0
    v3 := make(map[string]int)
    v4 := GenS1{F0: 127, F1: 1, F2: 0}
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 string = "a"
        for fi7 := 0; fi7 < 8; fi7++ {
            switch v6 {
            case "\n\t":
                v3[v6] = (fi5 | (fi5 | (^(fi7))))
            case "":
                v2 -= (v2 / ((v2) | 1))
            case "foo bar":
                v8 := v3[""]
                _ = v8
            }
            var v9 int = (fi7 ^ fi5)
            _ = fi7
            _ = v9
        }
        v2 -= (v2 + (v2 | 2))
        _ = fi5
        _ = v6
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        if len(v1) > 0 {
            v1[0] = 915
        }
        _ = fi10
    }
    v11 := GenS1{F0: 88, F1: 178, F2: -1}
    v12 := GenS2{F0: (v2 << ((413) & 63)), F1: (75 == 127)}
    v2 -= (v2 >> ((v2) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((1 >= 1) && (-63 < 36)) {
        for fi13 := 0; fi13 < 6; fi13++ {
            _ = len(v3)
            _ = fi13
        }
        v14 := GenS3{F0: 1, F1: 42, F2: -32768}
        v2 -= (v2 + (v2 & 846))
        _ = v14
    }
    v2 += v2
    if ((!(731 != 42)) || true) {
        v2 += (v2 ^ v2)
        v1 = append(v1, 681)
        var v15 float32 = 1e10
        _ = v15
    }
    var v16 int32 = 412
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
