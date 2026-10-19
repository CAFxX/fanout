package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int16 }
func genFn0(p0 bool, p1 uint64, p2 uint32) uint64 {
    for fi5 := 0; fi5 < 8; fi5++ {
        if (!p0) {
            fmt.Println(p0, p0)
        } else {
            if (255 < 122) {
                var v6 int = (fi5 % (((^(fi5))) | 1))
                v7 := []int{v6, fi5}
                _ = v6
                _ = v7
            } else {
                p2 = p2
            }
        }
        _ = fi5
    }
    switch "a" {
    case "foo bar":
        p2 = (p2 / (((p2 & (p2 ^ 42))) | 1))
    case "hello":
        p1 += p1
    case "\n\t":
        v8 := []int{32767, 0, -1}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1() uint32 {
    v9 := GenS1{F0: 0, F1: 1, F2: -1}
    var v10 int = 128
    _ = v10
    _ = v9
    return 846
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    v3 := make(map[string]string)
    var v4 uint = 42
    fmt.Println(v4)
    v4 = (v4 | v4)
    var v11 bool = false
    v12 := GenS1{F0: 2, F1: 127, F2: -380}
    switch "hello" {
    case "":
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := GenS1{F0: 107, F1: 127, F2: 1000}
            _ = fi13
            _ = v14
        }
    }
    v4 -= v4
    v15 := GenS1{F0: 255, F1: -1, F2: -32768}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 int64 = int64(fi16)
        _ = fi16
        _ = v17
    }
    if v11 {
        v18 := GenS1{F0: uint8(v4), F1: 100, F2: 1}
        for fi19 := 0; fi19 < 17; fi19++ {
            switch "foo bar" {
            case "a":
                v11 = (v11 && v11)
            }
            _ = fi19
        }
        _ = v18
    }
    v11 = v11
    var v20 uint32 = 391
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v11 = false
    var v21 uint = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
