package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int; F2 int8 }
type GenS3 struct { F0 uint; F1 uint }

func main() {
    var v1 int8 = -128
    v2 := make(map[int]string)
    v3 := []string{"hello"}
    var v4 int8 = 70
    v5 := make(map[int]uint32)
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 uint = 1
        v8 := v2[fi6]
        var v9 bool = (((!(1e10 == 0.1)) || true) && ("\n\t" < v8))
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v4 = -3
    v4 += (v1 - v4)
    switch 4294967295 {
    case 32768:
        fmt.Println(v1, v1)
    }
    var v10 uint16 = 42
    var v11 int = 9223372036854775807
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            fmt.Println(v4, v10)
            var v14 uint8 = 127
            if (!(42 <= 280)) {
                v15 := GenS1{F0: (^(v10))}
                _ = v15
            }
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    if ((v1 ^ v4) >= v1) {
        var v16 uint64 = 415
        _ = v16
    } else {
        fmt.Println(v10, v1)
        switch ("\n\t" + "foo bar") {
        case "a":
            if (false || (283 < 1)) {
                v3 = append(v3, "")
            }
        }
    }
    if false {
        switch (^(v11)) {
        case -542:
            v4 = (v1 ^ 100)
        case 120:
            _ = len(v3)
        case -404:
            v4 = v4
        }
        var v17 uint32 = 719
        v18 := GenS2{F0: v11, F1: v11, F2: v4}
        _ = v17
        _ = v18
    }
    v1 += (v4 << ((0) & 63))
    if false {
        v19 := GenS3{F0: 2, F1: 255}
        fmt.Println(v11, v10, v11)
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
