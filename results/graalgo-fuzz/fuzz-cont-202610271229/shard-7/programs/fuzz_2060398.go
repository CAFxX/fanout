package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint, p2 string) bool {
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            p0 += p0
            _ = fi8
        }
        _ = fi7
    }
    if (0 != 65535) {
        v9 := []int{}
        if len(v9) > 0 {
            v9[0] = -505
        }
        var v10 int8 = 1
        _ = v10
        _ = v9
    }
    var v11 uint = p1
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 string = p2
        fmt.Println(p2, p1)
        v11 = v11
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return false
}
func genFn1(p0 int32, p1 int16) uint64 {
    v14 := GenS1{F0: 599}
    v15 := GenS1{F0: 65535}
    var v16 int16 = p1
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return 2
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"hello", "\n\t", ("foo bar" + "hello"), "foo bar"}
    var v3 int = 42
    var v4 uint16 = 42
    var v5 int16 = 911
    var v6 int8 = 71
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v17 := GenS2{F0: v6, F1: -795}
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 uint64 = 69
        v20 := GenS3{F0: 255}
        v3 += (fi18 ^ v3)
        _ = fi18
        _ = v19
        _ = v20
    }
    v6 -= (v6 % (((-(v6))) | 1))
    var v21 int = (v3 << ((399) & 63))
    v22 := GenS2{F0: 127, F1: 42}
    v4 += (v4 >> ((530) & 63))
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
