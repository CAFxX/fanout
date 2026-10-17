package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint; F2 uint32 }

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[int]int)
    var v3 uint64 = 677
    var v4 uint32 = 125
    if true {
        v5 := []int{}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v5
    }
    var v6 float64 = (0.0 + 1e10)
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := []int{(^(v1)), v1, (fi7 % (((fi8 >> ((0) & 63))) | 1))}
            switch (fi8 ^ 42) {
            case -232:
                v10 := GenS1{F0: 466}
                var v11 string = ""
                _ = v10
                _ = v11
            case 2147483647:
                var v12 int64 = -99
                _ = v12
            case 0:
                delete(v2, (-(v1)))
                var v13 int32 = 102
                _ = v13
            }
            v1 -= (v1 << ((395) & 63))
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v14 := GenS1{F0: 9223372036854775807}
    var v15 int16 = 0
    var v16 int32 = -379
    v4 = 311
    for fi17 := 0; fi17 < 17; fi17++ {
        v15 = (v15 + (v15 % (((-(v15))) | 1)))
        _ = fi17
    }
    _ = len(v2)
    v18 := GenS2{F0: v15, F1: 1, F2: (v4 << ((42) & 63))}
    switch (v1 | v1) {
    case 32767:
        v19 := GenS1{F0: 42}
        _ = v19
    case -1:
        var v20 int16 = 42
        _ = v20
    case 0:
        v4 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
