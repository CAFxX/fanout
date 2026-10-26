package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 0.0
    var v3 uint16 = 2
    switch 4294967295 {
    case 699:
        v3 -= (v3 << ((568) & 63))
    case 9223372036854775807:
        v2 += ((v2 + v2) * v2)
    }
    v2 += v2
    v3 = (v3 ^ 1)
    v4 := GenS1{F0: 1}
    switch 9223372036854775807 {
    case 127:
        var v5 int64 = 0
        _ = v5
    }
    fmt.Println(v2, v3)
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            v2 += -2.25
            fmt.Println(fi6, v3, v2)
            var v8 int8 = 0
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v9 := GenS1{F0: 42}
    v10 := []int{}
    if len(v10) > 0 {
        v10[0] = 42
    }
    if false {
        var v11 uint32 = uint32(v3)
        _ = v11
    } else {
        fmt.Println(v2)
    }
    switch "a" {
    case "hello":
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS1{F0: 219}
            fmt.Println(v2, v2, fi12)
            switch (fi12 << ((0) & 63)) {
            case -1:
                _ = len(v10)
            case 1:
                fmt.Println(v3, fi12)
            case 32767:
                v14 := v1[-266]
                _ = v14
            }
            _ = fi12
            _ = v13
        }
    case "":
        v2 -= v2
    }
    v10 = append(v10, 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
