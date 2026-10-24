package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint64 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]string)
    var v3 string = ("hello" + "a")
    var v4 uint64 = v1
    var v5 int = 32768
    fmt.Println(v3, v5, v3)
    for fi6 := 0; fi6 < 17; fi6++ {
        fmt.Println(v5, v5, v4)
        v3 = v3
        v4 = v1
        _ = fi6
    }
    v5 += v5
    fmt.Println(v1, v4)
    fmt.Println(v4, v1, v5)
    if (1 >= 255) {
        fmt.Println(v1, v1)
    } else {
        var v7 int64 = -1
        _ = v7
    }
    fmt.Println(v3)
    switch (-(v5)) {
    case 9223372036854775807:
        switch v5 {
        case 32768:
            v4 = v1
        }
    case 128:
        for fi8 := 0; fi8 < 18; fi8++ {
            v4 += (v1 + (^(v4)))
            for fi9 := 0; fi9 < 14; fi9++ {
                v10 := GenS1{F0: 0, F1: -1, F2: 1}
                var v11 int = v5
                _ = len(v2)
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
    }
    var v12 string = "foo bar"
    v13 := GenS2{F0: 255, F1: 65535, F2: 167}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
