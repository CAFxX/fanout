package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint32 }
func genFn0() uint16 {
    v4 := []int{}
    var v5 int = -1
    v5 = 128
    if ((3.14159 != 3.14159) || false) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v5)
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 float64 = (1e10 + -2.25)
            if len(v4) > 0 {
                v4[0] = -41
            }
            _ = fi6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 19; fi8++ {
            if true {
                var v9 int8 = 42
                var v10 bool = true
                _ = v10
                _ = v9
            } else {
                v11 := GenS1{F0: 255}
                v12 := GenS1{F0: 1}
                fmt.Println(fi8)
                _ = v11
                _ = v12
            }
            var v13 uint16 = 42
            _ = fi8
            _ = v13
        }
        var v14 bool = true
        var v15 uint8 = 44
        _ = v14
        _ = v15
    }
    fmt.Println(v5)
    _ = v4
    _ = v5
    return 857
}

func main() {
    var v1 int16 = 26
    var v2 int64 = -1
    v3 := make(map[int]uint32)
    switch 718 {
    case 1:
        var v16 int64 = (v2 % (((v2 >> ((1) & 63))) | 1))
        _ = v16
    case -9223372036854775808:
        fmt.Println(v2, v1)
    case 42:
        if (42 == 649) {
            var v17 int = -9223372036854775808
            if (v17 < v17) {
                v18 := GenS1{F0: 1}
                _ = v18
            } else {
                if ((v17 % ((v17) | 1)) > v17) {
                    v19 := GenS2{F0: 855}
                    var v20 uint64 = 982
                    _ = v19
                    _ = v20
                }
                v21 := GenS3{F0: 4294967295, F1: 65535, F2: 9223372036854775807}
                v1 = v1
                _ = v21
            }
            _ = v17
        }
    }
    v1 += 303
    v22 := GenS2{F0: 399}
    v2 -= (v2 >> ((381) & 63))
    v2 += v2
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 float32 = 1.5
        if false {
            var v25 float64 = (-2.25 / (float64((v24 - v24)) / 1.0))
            for fi26 := 0; fi26 < 3; fi26++ {
                switch fi26 {
                case 682:
                    v2 = v2
                }
                v27 := GenS2{F0: 2}
                fmt.Println(v2, v24)
                _ = fi26
                _ = v27
            }
            fmt.Println(v2, v2, v1)
            _ = v25
        } else {
            switch "hello" {
            case "hello":
                var v28 float32 = (v24 + float32((-2.25 - 1.5)))
                _ = v28
            }
            if true {
                v24 += 0.0
                switch ("a" + ("\n\t" + "a")) {
                case "hello":
                    v29 := GenS4{F0: 127, F1: 2, F2: 2}
                    _ = v29
                }
            }
        }
        fmt.Println(v24)
        _ = fi23
        _ = v24
    }
    var v30 uint8 = 0
    v31 := GenS1{F0: 998}
    v32 := []int{}
    if false {
        v33 := v3[int(v30)]
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
