package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint, p1 bool, p2 int16) int16 {
    switch "a" {
    case "a":
        v6 := []int{-48}
        _ = v6
    case "foo bar":
        p1 = (-1 >= 1)
    }
    p2 = -896
    _ = p0
    _ = p1
    _ = p2
    return (p2 - (p2 * (p2 >> ((p0) & 63))))
}
func genFn1(p0 int16, p1 int8) uint16 {
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            if false {
                v9 := GenS1{F0: 159}
                v10 := GenS1{F0: 1}
                v11 := GenS2{F0: 42, F1: 217}
                switch (fi7 - -3) {
                case 128:
                    v12 := []int{(^(fi7))}
                    _ = v12
                case 4294967295:
                    v13 := []int{(fi7 >> ((2) & 63)), (^(fi7)), fi8}
                    _ = v13
                case 32767:
                    var v14 int64 = 42
                    _ = v14
                default:
                    var v15 string = ""
                    _ = v15
                }
                v16 := []int{(-(fi8)), (fi8 / ((-66) | 1))}
                var v17 uint64 = 4294967295
                _ = v10
                _ = v11
                _ = v16
                _ = v17
                _ = v9
            } else {
                var v18 bool = true
                if false {
                    v18 = v18
                } else {
                    var v19 float64 = (1.5 / 0.1)
                    _ = v19
                }
                v18 = v18
                _ = v18
            }
            fmt.Println(p0, fi7)
            switch fi7 {
            case 846:
                v20 := GenS3{F0: 0}
                _ = v20
            }
            _ = fi8
        }
        _ = fi7
    }
    var v21 uint16 = 255
    _ = p0
    _ = p1
    _ = v21
    return v21
}

func main() {
    v1 := []uint32{265, 682, 42}
    var v2 string = ""
    v3 := make(map[uint32]bool)
    var v4 string = (v2 + v2)
    var v5 float32 = -2.25
    var v22 bool = ((true || (-1 == genFn0(765, false, 42))) || true)
    fmt.Println(v22, v22, v22)
    if len(v1) > 0 {
        v1[0] = 611
    }
    var v23 float64 = 0.1
    if v22 {
        var v24 bool = false
        _ = v24
    } else {
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 uint64 = 737
            v27 := GenS3{F0: 1}
            _ = fi25
            _ = v26
            _ = v27
        }
        var v28 int64 = 1
        _ = v28
    }
    v23 += v23
    var v29 float32 = ((v5 / (v5 + (v5 * -2.25))) * v5)
    for fi30 := 0; fi30 < 1; fi30++ {
        v31 := GenS4{F0: 2, F1: 0}
        _ = fi30
        _ = v31
    }
    fmt.Println(v22, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
