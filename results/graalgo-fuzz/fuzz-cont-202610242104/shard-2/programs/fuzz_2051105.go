package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 bool; F1 bool; F2 uint16 }
func genFn0(p0 uint64, p1 int8) int64 {
    p0 = p0
    for fi4 := 0; fi4 < 20; fi4++ {
        if (42 == 803) {
            fmt.Println(p0, fi4)
            p1 = p1
        } else {
            for fi5 := 0; fi5 < 2; fi5++ {
                v6 := []int{fi5}
                _ = fi5
                _ = v6
            }
            if false {
                v7 := GenS1{F0: (fi4 * 823), F1: (fi4 | (-(fi4)))}
                _ = v7
            }
            v8 := GenS1{F0: 32767, F1: fi4}
            _ = v8
        }
        v9 := GenS2{F0: 163, F1: p0}
        _ = fi4
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS2{F0: 65535, F1: (p0 & (^(p0)))}
        var v12 uint8 = 127
        for fi13 := 0; fi13 < 10; fi13++ {
            p0 = p0
            p1 = (p1 & p1)
            _ = fi13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []int{371, 1, 42}
    v2 := []uint32{483, 2, 693, 0}
    v3 := make(map[int]uint32)
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 int32 = 2147483647
        _ = fi14
        _ = v15
    }
    var v16 int64 = -676
    var v17 uint16 = 42
    v17 = v17
    var v18 uint32 = 1
    v17 = (v17 & v17)
    var v19 uint = 65535
    v20 := GenS3{F0: false, F1: (!(!false)), F2: (v17 % (((v17 & (v17 << (((^(v19))) & 63)))) | 1))}
    var v21 float32 = 0.0
    v22 := v3[127]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
