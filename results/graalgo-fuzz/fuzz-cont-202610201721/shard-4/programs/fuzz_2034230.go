package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
func genFn0() int8 {
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := []int{2147483647}
        var v7 uint = 0
        for fi8 := 0; fi8 < 9; fi8++ {
            if (32767 == -1) {
                v7 += (^(v7))
                _ = len(v6)
                _ = len(v6)
                v9 := GenS1{F0: 0, F1: 4294967295}
                v7 = (v7 ^ v7)
                var v10 bool = true
                _ = v10
                _ = v9
            } else {
                v7 += v7
                v7 = (v7 | v7)
            }
            v6 = append(v6, (fi5 << ((v7) & 63)))
            _ = fi8
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint16 = 0
        fmt.Println(v12, fi11)
        for fi13 := 0; fi13 < 18; fi13++ {
            if (504 <= -995) {
                if (true && (-128 >= int8(fi13))) {
                    v12 -= 32767
                }
                v14 := []int{fi13, (fi13 - 687), (fi11 | 1), (fi13 / ((fi11) | 1))}
                var v15 float64 = (-2.25 + (float64(1e10) / 1e10))
                var v16 int16 = -1
                v15 -= 1e10
                _ = len(v14)
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := GenS2{F0: (fi17 | (fi17 - (fi17 % ((-2147483648) | 1))))}
        v19 := []int{}
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 uint64 = 4294967295
            _ = fi20
            _ = v21
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        fmt.Println(fi22)
        v23 := []int{fi22}
        var v24 float64 = float64(-2.25)
        _ = fi22
        _ = v23
        _ = v24
    }
    return 127
}

func main() {
    v1 := []uint32{43, 210, 2147483648}
    var v2 string = "hello"
    v3 := make(map[string]int)
    var v4 float32 = (3.14159 - 0.1)
    var v25 uint32 = 0
    var v26 uint8 = 1
    var v27 bool = (!((!(-120 >= -1)) || true))
    v28 := GenS2{F0: 1}
    v29 := v3[""]
    v30 := GenS1{F0: 253, F1: (^(v25))}
    v31 := GenS2{F0: (v29 & (v29 & v29))}
    v32 := GenS2{F0: -2147483648}
    v33 := GenS3{F0: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
