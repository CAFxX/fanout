package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int64 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    var v3 int32 = 42
    var v4 int = 594
    v5 := make(map[string]bool)
    if (!(-2.25 < ((1.5 / (0.0 - 1.5)) / 0.1))) {
        if true {
            for fi6 := 0; fi6 < 11; fi6++ {
                _ = len(v1)
                _ = fi6
            }
            v7 := GenS1{F0: 0, F1: (v4 | v4), F2: 558}
            var v8 int = (^(v4))
            _ = v7
            _ = v8
        }
    }
    fmt.Println(v4, v3)
    v9 := GenS1{F0: -29, F1: v4, F2: 635}
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint64 = 4294967295
        _ = fi10
        _ = v11
    }
    if (!(-22 == -1)) {
        if (979 <= 0) {
            switch (("\n\t" + "hello") + "a") {
            case "hello":
                v12 := GenS2{F0: v3}
                _ = v12
            }
            v3 -= (v3 % (((v3 & (v3 % ((v3) | 1)))) | 1))
        }
        v3 = (v3 - v3)
        var v13 bool = (0 < 849)
        _ = v13
    }
    v14 := GenS2{F0: -1}
    v5["foo bar"] = ((((v4 & 9223372036854775807) != (v4 << ((255) & 63))) && true) || false)
    v15 := []int{-9223372036854775808, v4}
    var v16 float32 = 0.1
    var v17 int64 = -1
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 6; fi19++ {
            fmt.Println(fi19)
            fmt.Println(v17, fi18)
            _ = fi19
        }
        _ = fi18
    }
    v15 = append(v15, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
