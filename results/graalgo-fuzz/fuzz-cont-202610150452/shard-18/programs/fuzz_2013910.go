package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint64, p1 int64) uint64 {
    v7 := []int{-2147483648, -231, 32767}
    for fi8 := 0; fi8 < 14; fi8++ {
        switch "foo bar" {
        case "hello":
            var v9 float64 = ((0.0 + 0.1) * (1.5 - (1e10 - 1e10)))
            _ = v9
        default:
            p1 = p1
        }
        v10 := GenS1{F0: (p0 >> ((42) & 63)), F1: 1}
        _ = fi8
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    return p0
}
func genFn1() int64 {
    v11 := []int{}
    if true {
        var v12 float32 = 0.0
        if false {
            fmt.Println(v12)
            for fi13 := 0; fi13 < 3; fi13++ {
                v12 -= 1.5
                _ = fi13
            }
        } else {
            _ = len(v11)
            var v14 uint64 = 2
            _ = v14
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 uint8 = 2
            _ = fi15
            _ = v16
        }
        _ = v12
    }
    _ = v11
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    v3 := []uint32{1, 42, 1}
    var v4 uint64 = 42
    v5 := make(map[int]bool)
    var v6 int8 = int8(v4)
    var v17 string = ""
    fmt.Println(v4)
    v2["foo bar"] = "hello"
    v18 := GenS1{F0: v4, F1: 42}
    delete(v1, 2)
    fmt.Println(v6)
    if true {
        v4 = 2
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            v17 = "a"
            switch "hello" {
            case "foo bar":
                v4 = (v4 - v4)
                var v21 int64 = -1
                fmt.Println(v4, v4, v17)
                v22 := GenS1{F0: v4, F1: int32(fi20)}
                if true {
                    var v23 bool = true
                    fmt.Println(v4)
                    v24 := GenS2{F0: 555, F1: 0}
                    _ = v23
                    _ = v24
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v21
                _ = v22
            default:
                var v25 uint16 = 65535
                _ = v25
            }
            _ = fi20
        }
        v6 = (v6 & int8(v4))
        fmt.Println(v17, v4)
        _ = fi19
    }
    v26 := GenS1{F0: (v4 << ((255) & 63)), F1: 2147483647}
    var v27 bool = (!false)
    v28 := GenS3{F0: 42, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
