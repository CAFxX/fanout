package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint16, p1 uint, p2 uint16) uint32 {
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(p1, p0, p1)
        p0 = (p0 % ((1) | 1))
        if (2147483647 <= 127) {
            p2 = (p2 & (p2 ^ p0))
            fmt.Println(p2, fi6, fi6)
            switch "\n\t" {
            case "foo bar":
                v7 := []int{}
                _ = v7
            case "":
                v8 := []int{2147483647, fi6, (^(fi6)), fi6}
                _ = v8
            }
        } else {
            p1 = (p1 ^ (^(p1)))
            if (!(!false)) {
                p1 = p1
                v9 := GenS1{F0: 127, F1: 0}
                _ = v9
            }
            var v10 int16 = -716
            _ = v10
        }
        _ = fi6
    }
    v11 := []int{1}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 594
}
func genFn1() uint32 {
    var v12 int8 = 1
    var v13 float32 = -2.25
    var v14 int64 = 42
    _ = v12
    _ = v13
    _ = v14
    return 1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v4 int64 = -1
    v5 := []uint32{505, 167}
    v4 = 42
    var v15 int16 = 1
    fmt.Println(v4, v4)
    for fi16 := 0; fi16 < 14; fi16++ {
        fmt.Println(v15, v15, v15)
        v3[2147483647] = true
        _ = fi16
    }
    var v17 uint16 = 555
    for fi18 := 0; fi18 < 18; fi18++ {
        fmt.Println(v4, v4, v17)
        var v19 int64 = (-(v4))
        v15 += v15
        _ = fi18
        _ = v19
    }
    var v20 int = 32768
    if true {
        switch v20 {
        case 9223372036854775807:
            for fi21 := 0; fi21 < 8; fi21++ {
                v5 = append(v5, 4294967295)
                _ = fi21
            }
        }
        var v22 uint8 = 2
        _ = v22
    } else {
        for fi23 := 0; fi23 < 15; fi23++ {
            v15 += v15
            var v24 uint16 = v17
            _ = fi23
            _ = v24
        }
        v25 := GenS1{F0: 88, F1: v17}
        delete(v2, "\n\t")
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
