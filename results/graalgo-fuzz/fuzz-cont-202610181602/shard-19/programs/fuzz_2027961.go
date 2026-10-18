package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 bool = false
    var v4 float64 = -2.25
    var v5 int32 = 928
    v6 := make(map[string]int)
    var v7 uint64 = 449
    if (true && true) {
        v5 = (v5 ^ (v5 ^ (-(v5))))
        if ((4294967295 != 65535) && (2147483647 <= -821)) {
            for fi8 := 0; fi8 < 11; fi8++ {
                v9 := []int{fi8, -1}
                v10 := v6[("" + "hello")]
                v11 := GenS1{F0: 236, F1: 100}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        } else {
            if (v5 <= v5) {
                var v12 float32 = 3.14159
                if (int16(v7) < -1) {
                    v13 := GenS2{F0: 42, F1: 127, F2: 0}
                    v12 -= v12
                    v4 = v4
                    _ = v13
                } else {
                    v14 := []int{}
                    for fi15 := 0; fi15 < 11; fi15++ {
                        _ = len(v1)
                        _ = fi15
                    }
                    _ = v14
                }
                v4 -= ((0.1 - v4) + float64(float32(3.14159)))
                _ = v12
            } else {
                var v16 string = ""
                switch "hello" {
                case "a":
                    fmt.Println(v16, v7)
                }
                _ = v16
            }
            fmt.Println(v7, v7, v5)
            if false {
                v17 := GenS3{F0: 42}
                delete(v2, 2147483647)
                _ = v17
            } else {
                var v18 int16 = 495
                v7 = v7
                _ = v18
            }
        }
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        v20 := GenS2{F0: 1, F1: -1, F2: 2}
        var v21 uint8 = 0
        v22 := GenS1{F0: 42, F1: 127}
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 int64 = -1
    v4 += v4
    v24 := []int{-637, 0}
    fmt.Println(v7, v23, v4)
    v4 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
