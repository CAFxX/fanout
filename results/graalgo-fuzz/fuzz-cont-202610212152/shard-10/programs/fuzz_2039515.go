package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint8 }

func main() {
    var v1 uint64 = 42
    var v2 uint64 = v1
    var v3 float32 = (3.14159 * (0.1 / 0.1))
    v4 := make(map[int]int)
    var v5 uint16 = 0
    v2 = v1
    var v6 uint16 = (^(v5))
    var v7 bool = true
    for fi8 := 0; fi8 < 9; fi8++ {
        v3 = ((float32(0.0) + float32(float64((v3 + ((v3 * float32((1.5 - (1.5 / 1.0)))) / v3))))) * float32((((((-2.25 * -2.25) - 0.1) / ((-2.25 / float64(v3)) / -2.25)) - (float64((v3 / (v3 - 1.5))) * -2.25)) + 0.0)))
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS1{F0: -392}
            v11 := GenS2{F0: 0, F1: v7}
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint = 65535
        _ = fi8
        _ = v12
    }
    v13 := GenS3{F0: 1}
    var v14 int = 279
    if v7 {
        v5 = v6
        delete(v4, (v14 * (v14 ^ 1)))
    }
    var v15 uint = 4294967295
    for fi16 := 0; fi16 < 5; fi16++ {
        if true {
            var v17 uint = v15
            _ = v17
        } else {
            fmt.Println(v2)
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            if false {
                if v7 {
                    if false {
                        v19 := []int{}
                        fmt.Println(v15, fi18)
                        if (v7 && true) {
                            v20 := GenS4{F0: v6, F1: v1, F2: 128}
                            _ = v20
                        } else {
                            v6 += v5
                            v21 := GenS4{F0: v5, F1: (v1 / ((v2) | 1)), F2: 51}
                            _ = v21
                        }
                        _ = v19
                    }
                    delete(v4, fi16)
                    var v22 uint64 = v1
                    _ = v22
                }
            }
            var v23 string = "a"
            var v24 bool = v7
            _ = fi18
            _ = v23
            _ = v24
        }
        _ = fi16
    }
    v25 := []int{-596, (v14 * 1), (v14 * v14)}
    var v26 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
