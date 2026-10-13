package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "foo bar"
    v3 := make(map[uint32]bool)
    v4 := []string{v2}
    v5 := []int{-1, -272, -957, -892}
    v6 := v1[371]
    var v7 float32 = 1e10
    for fi8 := 0; fi8 < 11; fi8++ {
        switch (-(fi8)) {
        case 959:
            fmt.Println(v6, v2)
        default:
            v7 -= v7
        }
        v7 += v7
        _ = fi8
    }
    var v9 float64 = 0.0
    if false {
        v6 = v6
        v10 := GenS1{F0: v2}
        v11 := GenS2{F0: v2}
        _ = v10
        _ = v11
    }
    var v12 int = 9223372036854775807
    if (528 > 0) {
        delete(v3, v6)
        if false {
            var v13 float32 = v7
            _ = v13
        }
    } else {
        if (42 >= 0) {
            v14 := GenS2{F0: "hello"}
            _ = v14
        } else {
            var v15 int16 = 32767
            _ = v15
        }
    }
    v7 += (1e10 - v7)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (1 > 6) {
        if (654 < 18446744073709551615) {
            v2 = v2
        }
    }
    switch v12 {
    case 127:
        v16 := GenS2{F0: (v2 + "\n\t")}
        _ = v16
    }
    fmt.Println(v2)
    fmt.Println(v2, v6, v9)
    fmt.Println(v9, v6, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
