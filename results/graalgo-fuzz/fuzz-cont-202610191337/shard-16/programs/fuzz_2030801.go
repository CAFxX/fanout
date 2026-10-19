package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int64; F2 uint32 }

func main() {
    v1 := []string{"a", "foo bar", (("\n\t" + (("hello" + (("hello" + "foo bar") + "\n\t")) + "a")) + "\n\t")}
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    var v4 int32 = 160
    v5 := []string{"\n\t"}
    var v6 int16 = 1
    fmt.Println(v6)
    v4 += (v4 >> ((2) & 63))
    v7 := GenS1{F0: true}
    v4 = (v4 + (v4 % ((v4) | 1)))
    fmt.Println(v6, v6)
    if (42 == 1) {
        v8 := GenS1{F0: (0.1 <= 0.1)}
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int16 = (v6 % (((v6 << ((255) & 63))) | 1))
            fmt.Println(v10, v10, v10)
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            if true {
                fmt.Println(v4, fi11)
                fmt.Println(v4, v4, v6)
                v4 -= (v4 / (((-(v4))) | 1))
            }
            _ = fi11
        }
        _ = v8
    } else {
        v4 += v4
        v12 := GenS2{F0: 2, F1: 42, F2: 65535}
        _ = v12
    }
    var v13 int32 = v4
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
