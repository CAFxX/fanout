package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 string }

func main() {
    v1 := []int{197, -518}
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 uint32 = 2147483648
    var v5 uint8 = 1
    var v6 int = 1
    var v7 int32 = 979
    v4 = uint32(v5)
    var v8 uint = 42
    var v9 bool = false
    _ = len(v3)
    fmt.Println(v9, v4)
    v10 := GenS1{F0: (v4 ^ v4), F1: ("\n\t" + "foo bar")}
    v9 = false
    v4 -= v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 uint32 = 312
    fmt.Println(v8, v11)
    v11 = v4
    switch v6 {
    case 138:
        if (-106 > 104) {
            v12 := v2[-2147483648]
            _ = v12
        }
    case 32767:
        var v13 uint = v8
        _ = v13
    case 639:
        v1 = append(v1, v6)
    }
    if (true || true) {
        _ = len(v3)
        for fi14 := 0; fi14 < 1; fi14++ {
            for fi15 := 0; fi15 < 13; fi15++ {
                fmt.Println(v5, fi14, v5)
                v16 := GenS2{F0: "foo bar"}
                var v17 uint64 = 1
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi14
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
