package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []uint32{}
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 int = 128
    var v5 int64 = 269
    _ = len(v1)
    fmt.Println(v5, v5)
    if (42 == 945) {
        v6 := GenS1{F0: 2}
        var v7 uint8 = 128
        switch (v4 / (((v4 % ((v4) | 1))) | 1)) {
        case -1:
            fmt.Println(v7, v5)
        }
        _ = v6
        _ = v7
    } else {
        if (!((v4 >> ((600) & 63)) > 771)) {
            v4 -= v4
            v5 += (v5 + (v5 * 1))
        }
        if true {
            v4 = 509
        }
    }
    var v8 int32 = -1
    fmt.Println(v5)
    var v9 float32 = -2.25
    _ = len(v1)
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
