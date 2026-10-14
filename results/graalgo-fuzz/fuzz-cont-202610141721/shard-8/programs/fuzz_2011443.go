package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint32; F1 uint; F2 string }
type GenS3 struct { F0 bool; F1 int16; F2 int32 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 1
    var v3 float64 = (1e10 - (1e10 - 0.0))
    var v4 float32 = 0.0
    var v5 string = (("hello" + "") + ("\n\t" + ""))
    v6 := GenS1{F0: 0, F1: v5}
    if false {
        delete(v1, 290)
        var v7 int64 = 1
        var v8 int32 = 2147483647
        _ = v7
        _ = v8
    }
    if (42 < 1) {
        v9 := []int{-845}
        switch 420 {
        case -9223372036854775808:
            v10 := GenS2{F0: 42, F1: 42, F2: (v5 + (v5 + v5))}
            _ = v10
        case 0:
            if len(v9) > 0 {
                v9[0] = 42
            }
        default:
            var v11 int = 127
            _ = v11
        }
        _ = v9
    } else {
        v12 := GenS3{F0: true, F1: 42, F2: 42}
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS4{F0: 65535, F1: 590}
            switch (fi13 + 1) {
            case 9223372036854775807:
                if false {
                    var v15 float64 = float64(((((((v4 / 1e10) * (v4 * v4)) * 3.14159) / v4) * (((((v4 / v4) / v4) / float32(v3)) - float32(v3)) + (((v4 * v4) * 0.0) + float32(v3)))) - 1.5))
                    v16 := GenS4{F0: 255, F1: 9223372036854775807}
                    _ = v15
                    _ = v16
                }
            default:
                v5 = (v5 + v5)
            }
            fmt.Println(fi13, fi13)
            _ = fi13
            _ = v14
        }
        var v17 int64 = 923
        _ = v12
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS3{F0: (fi18 != fi18), F1: -914, F2: 1}
        _ = fi18
        _ = v19
    }
    var v20 uint = 42
    v21 := []int{}
    v22 := GenS5{F0: 0}
    v5 = (v5 + v5)
    v23 := GenS5{F0: 4294967295}
    if false {
        for fi24 := 0; fi24 < 15; fi24++ {
            v21 = append(v21, -509)
            _ = fi24
        }
        v4 = 1e10
        v3 += v3
    }
    v20 += v20
    for fi25 := 0; fi25 < 16; fi25++ {
        if true {
            if (-27 == -199) {
                fmt.Println(fi25, v4, fi25)
                var v26 uint16 = 1
                v20 = 0
                _ = v26
            } else {
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
                var v27 int8 = 70
                _ = v27
            }
            var v28 uint32 = 1
            _ = v28
        } else {
            if len(v21) > 0 {
                v21[0] = (fi25 & (fi25 | 9223372036854775807))
            }
            v20 -= (v20 >> ((65535) & 63))
            var v29 float32 = 1e10
            _ = v29
        }
        switch v5 {
        case "foo bar":
            if true {
                fmt.Println(v20, v2)
                v2 += 734
            }
        }
        v20 += 2
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
