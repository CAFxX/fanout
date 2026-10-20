package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 float32 = 1.5
    var v5 uint32 = 551
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{fi6}
            v4 -= 0.1
            v5 = (^(v5))
            _ = fi7
            _ = v8
        }
        fmt.Println(fi6)
        _ = fi6
    }
    var v9 uint32 = (^(v5))
    v10 := []int{2147483647, -921, int(v9), 128}
    for fi11 := 0; fi11 < 1; fi11++ {
        v5 = (^(v9))
        var v12 uint16 = 1
        if false {
            switch (-(fi11)) {
            case 983:
                v4 = v4
            case 803:
                fmt.Println(fi11, v12, v9)
            default:
                v4 = v4
            }
        }
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        if ((true && (true && false)) || false) {
            v14 := GenS1{F0: 940, F1: -1}
            v9 -= v5
            var v15 int = (^(fi13))
            _ = v14
            _ = v15
        }
        switch (fi13 / (((fi13 >> ((0) & 63))) | 1)) {
        case -2147483648:
            if false {
                switch fi13 {
                case 775:
                    var v16 int64 = -79
                    _ = v16
                default:
                    for fi17 := 0; fi17 < 12; fi17++ {
                        var v18 int8 = 105
                        _ = fi17
                        _ = v18
                    }
                }
                fmt.Println(v5, fi13, v4)
            }
        case 19:
            if true {
                v19 := GenS2{F0: 0}
                v4 += 1.5
                if (168 == (v9 * 1)) {
                    switch (fi13 + -2147483648) {
                    case -1:
                        v20 := GenS2{F0: 2}
                        _ = v20
                    case 127:
                        v4 += 1e10
                    }
                    _ = len(v1)
                    if (!false) {
                        fmt.Println(v5, fi13)
                        v9 += v5
                    } else {
                        if len(v10) > 0 {
                            v10[0] = fi13
                        }
                    }
                } else {
                    if false {
                        v5 -= v5
                    } else {
                        v5 = v5
                    }
                }
                _ = v19
            } else {
                if len(v10) > 0 {
                    v10[0] = (fi13 | fi13)
                }
            }
        case -9223372036854775808:
            v4 += v4
        }
        var v21 int16 = 1000
        _ = fi13
        _ = v21
    }
    var v22 float32 = v4
    v23 := GenS3{F0: true, F1: 25}
    fmt.Println(v9)
    var v24 uint64 = 311
    v4 = (v22 + v22)
    var v25 uint32 = 169
    v4 = float32((1.5 + 0.1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
