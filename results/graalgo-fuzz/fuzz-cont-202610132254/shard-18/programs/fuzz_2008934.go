package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int }
type GenS3 struct { F0 uint16; F1 int }

func main() {
    var v1 bool = (true && (0 < 100))
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    v4 := GenS1{F0: -59, F1: 680}
    switch "" {
    case "\n\t":
        for fi5 := 0; fi5 < 18; fi5++ {
            fmt.Println(fi5)
            _ = fi5
        }
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        delete(v2, (fi6 << ((65535) & 63)))
        _ = fi6
    }
    v7 := GenS2{F0: 2, F1: 2, F2: 2147483647}
    var v8 uint64 = 521
    if v1 {
        v8 = (v8 + 42)
    }
    switch 9223372036854775807 {
    case 128:
        var v9 bool = (32767 <= 0)
        _ = v9
    }
    var v10 bool = false
    var v11 int = 127
    for fi12 := 0; fi12 < 3; fi12++ {
        if v1 {
            switch v11 {
            case 42:
                v8 += (v8 >> ((42) & 63))
            default:
                delete(v3, ("hello" + "\n\t"))
            }
        } else {
            switch fi12 {
            case 652:
                v13 := GenS3{F0: 1, F1: -142}
                _ = v13
            default:
                var v14 int32 = -1
                _ = v14
            }
            var v15 int = 42
            v16 := GenS3{F0: 2, F1: (fi12 | (fi12 % ((0) | 1)))}
            _ = v15
            _ = v16
        }
        fmt.Println(v10, v11, v10)
        if (1 > 59) {
            var v17 uint64 = (^(v8))
            _ = v17
        } else {
            v18 := v2[(fi12 + (^(v11)))]
            _ = v18
        }
        _ = fi12
    }
    v10 = (2 >= 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
