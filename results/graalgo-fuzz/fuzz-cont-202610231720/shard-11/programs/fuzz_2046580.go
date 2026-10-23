package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }

func main() {
    var v1 int32 = -462
    var v2 int = 128
    v3 := make(map[int]int)
    if (false && true) {
        fmt.Println(v1)
        for fi4 := 0; fi4 < 15; fi4++ {
            v2 = 730
            v5 := GenS1{F0: int16(v1)}
            var v6 string = "hello"
            _ = fi4
            _ = v5
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 14; fi7++ {
            v3[(fi7 % ((v2) | 1))] = v2
            if (!false) {
                var v8 string = "foo bar"
                _ = v8
            }
            _ = fi7
        }
        if false {
            v2 = (v2 & v2)
            v9 := []int{9223372036854775807}
            v10 := GenS2{F0: 1}
            _ = v10
            _ = v9
        } else {
            v1 -= (v1 % (((v1 * v1)) | 1))
        }
    }
    v1 -= 42
    if true {
        v11 := GenS1{F0: 42}
        v12 := []int{-420}
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            if true {
                var v15 uint32 = 0
                var v16 uint64 = 2
                _ = v15
                _ = v16
            } else {
                v2 -= (fi14 * (v2 * fi13))
                v2 -= (fi13 >> ((2) & 63))
            }
            _ = fi14
        }
        var v17 int64 = 1
        for fi18 := 0; fi18 < 19; fi18++ {
            v17 = v17
            var v19 int32 = (v1 & -1)
            var v20 int = 128
            switch fi18 {
            case 128:
                v1 -= (v19 * -1)
            default:
                if true {
                    if (!(2 == 255)) {
                        v21 := GenS1{F0: 1000}
                        var v22 string = "hello"
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS2{F0: (v17 << ((168) & 63))}
                    v24 := GenS3{F0: 1, F1: 0, F2: 2}
                    var v25 uint32 = 0
                    _ = v23
                    _ = v24
                    _ = v25
                }
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi13
        _ = v17
    }
    v26 := v3[v2]
    _ = len(v3)
    var v27 int64 = 1
    v28 := GenS4{F0: -215}
    var v29 int8 = 103
    var v30 float64 = 1.5
    if (true || (!(!((v29 * 127) == (v29 >> ((255) & 63)))))) {
        v31 := GenS5{F0: (^(v2))}
        fmt.Println(v1, v27, v27)
        _ = v31
    }
    var v32 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
}
