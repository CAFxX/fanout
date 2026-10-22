package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint16 }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{65535, 42, 2}
    var v3 uint16 = 650
    var v4 float64 = (((1e10 + -2.25) - 0.1) - 1e10)
    v5 := make(map[int]int)
    for fi6 := 0; fi6 < 4; fi6++ {
        if (!true) {
            if ((^(v3)) != v3) {
                var v7 uint16 = (v3 ^ v3)
                v8 := GenS1{F0: v1, F1: 0, F2: v7}
                _ = v7
                _ = v8
            }
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 int16 = -10
                _ = fi9
                _ = v10
            }
        } else {
            for fi11 := 0; fi11 < 7; fi11++ {
                fmt.Println(v4, v3)
                _ = fi11
            }
        }
        v1 = (v1 + "\n\t")
        _ = fi6
    }
    switch -375 {
    case 127:
        v3 = 255
    case 1:
        v4 -= float64(-2.25)
    }
    var v12 bool = (1.5 == (v4 + v4))
    v3 = v3
    v5[1] = -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3, v1)
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 int16 = 367
        if true {
            for fi15 := 0; fi15 < 7; fi15++ {
                fmt.Println(fi15, fi13, fi15)
                var v16 int16 = (v14 | -507)
                v4 += v4
                var v17 int64 = 0
                switch fi13 {
                case 4294967295:
                    var v18 int = fi15
                    _ = v18
                case 128:
                    var v19 uint64 = 42
                    _ = v19
                case 32768:
                    v20 := GenS1{F0: "a", F1: -811, F2: (v3 & v3)}
                    _ = v20
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            delete(v5, (fi13 >> ((255) & 63)))
            var v21 int32 = -976
            _ = v21
        }
        _ = fi13
        _ = v14
    }
    var v22 uint = 955
    var v23 int64 = 1
    switch -386 {
    case 383:
        if true {
            if v12 {
                v3 = (v3 - (^(v3)))
                v12 = true
                var v24 int16 = 1000
                _ = v24
            }
            v25 := v5[-304]
            v4 -= 1e10
            _ = v25
        } else {
            v26 := GenS2{F0: -431, F1: 2, F2: v3}
            _ = v26
        }
    }
    var v27 uint32 = 1
    v28 := GenS1{F0: "\n\t", F1: -39, F2: (v3 % ((158) | 1))}
    if false {
        var v29 int8 = 0
        if len(v2) > 0 {
            v2[0] = (v27 & 4294967295)
        }
        v1 = v1
        _ = v29
    }
    v27 -= (v27 % (((v27 >> ((v22) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
