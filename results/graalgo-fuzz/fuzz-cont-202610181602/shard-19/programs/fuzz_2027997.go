package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    v3 := []int{42, 2147483647, -1, 32768}
    for fi4 := 0; fi4 < 19; fi4++ {
        delete(v1, 65535)
        _ = fi4
    }
    var v5 string = ""
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int32 = 42
        var v8 uint64 = 858
        v8 -= v8
        _ = fi6
        _ = v7
        _ = v8
    }
    v5 = (v5 + v5)
    if true {
        v3 = append(v3, 4294967295)
    }
    var v9 int32 = -184
    if (!true) {
        if false {
            v5 = v5
            v5 = ""
        }
        v10 := GenS1{F0: 255, F1: 0}
        _ = v10
    }
    var v11 int = 128
    fmt.Println(v9, v11, v9)
    var v12 float64 = 0.0
    var v13 string = ((v5 + v5) + "foo bar")
    if (!false) {
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 uint64 = 4294967295
            _ = fi14
            _ = v15
        }
        if false {
            if (1 > 255) {
                fmt.Println(v5, v12)
                fmt.Println(v9)
                v11 -= v11
            }
            var v16 int32 = v9
            _ = v16
        } else {
            v9 += -2147483648
            v5 = v5
            var v17 uint64 = 18446744073709551615
            _ = v17
        }
        v9 = (v9 * v9)
    } else {
        switch (v11 << ((18446744073709551615) & 63)) {
        case 32768:
            delete(v2, v11)
        case 127:
            fmt.Println(v5, v12, v5)
        case 643:
            if true {
                if len(v3) > 0 {
                    v3[0] = -1
                }
            }
        }
        _ = len(v3)
    }
    fmt.Println(v9)
    fmt.Println(v9, v5)
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := GenS2{F0: v11}
        _ = fi18
        _ = v19
    }
    v9 += -1
    if (!false) {
        v13 = v13
        v3 = append(v3, (-(v11)))
        var v20 uint32 = 2147483648
        _ = v20
    } else {
        fmt.Println(v13, v5)
        v11 = (v11 & v11)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
