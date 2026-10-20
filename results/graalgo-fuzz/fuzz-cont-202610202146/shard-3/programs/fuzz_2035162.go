package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 1
    v3 := make(map[uint32]uint32)
    _ = len(v3)
    v2 -= (^(v2))
    for fi4 := 0; fi4 < 8; fi4++ {
        if (!false) {
            var v5 int64 = int64(fi4)
            _ = len(v1)
            _ = v5
        }
        switch (-(fi4)) {
        case 0:
            var v6 int8 = 41
            _ = v6
        default:
            v2 = 1
        }
        _ = fi4
    }
    v2 = (v2 % (((v2 % ((65535) | 1))) | 1))
    var v7 uint32 = 65535
    v8 := GenS1{F0: true, F1: -722}
    v7 = (v7 % (((v7 * v7)) | 1))
    v9 := GenS2{F0: -560, F1: int64(v2)}
    var v10 string = (("\n\t" + "") + ("foo bar" + "foo bar"))
    fmt.Println(v10)
    var v11 float32 = (0.1 - 1e10)
    var v12 bool = (584 >= 956)
    v2 += (^(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
