package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 bool }
type GenS2 struct { F0 int8; F1 uint; F2 uint64 }
func genFn0(p0 uint16, p1 int, p2 int) uint64 {
    v5 := []int{(p1 & (-(p1)))}
    var v6 int8 = 42
    v6 = (v6 * v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 630
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]bool)
    var v3 int64 = -1
    var v4 bool = false
    var v7 float32 = 3.14159
    if ((0.0 < v7) || ((^(v3)) > v3)) {
        fmt.Println(v4, v4)
        var v8 int16 = 1000
        _ = v8
    } else {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{(fi9 >> ((65535) & 63)), -687, fi9}
            _ = fi9
            _ = v10
        }
        if v4 {
            v11 := []int{}
            if v4 {
                v4 = (v4 || (!(v7 <= v7)))
                var v12 uint16 = 255
                _ = v12
            } else {
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 uint64 = genFn0(2, -1, fi13)
                    v4 = false
                    _ = fi13
                    _ = v14
                }
                var v15 float64 = 0.1
                _ = v15
            }
            _ = v11
        } else {
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 uint = 691
                _ = fi16
                _ = v17
            }
            var v18 int = -192
            var v19 uint8 = 1
            _ = v18
            _ = v19
        }
        v20 := GenS1{F0: 1, F1: -1, F2: true}
        _ = v20
    }
    var v21 float32 = ((v7 / 1.0) + v7)
    v1 = ((((v1 + v1) + (v1 + v1)) + (v1 + "")) + "hello")
    v22 := GenS2{F0: 0, F1: 0, F2: 42}
    var v23 uint32 = 0
    var v24 int8 = 0
    v25 := []int{27, 32768, 32768}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v7
}
