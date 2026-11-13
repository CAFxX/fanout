package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 bool }
type GenS4 struct { F0 int16 }
func genFn0(p0 int64, p1 int16) uint32 {
    var v7 uint8 = 180
    fmt.Println(v7)
    var v8 int16 = (-(p1))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 10
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 1e10
    var v3 uint = 674
    var v4 int = 128
    v5 := make(map[uint32]int)
    var v6 int64 = 0
    fmt.Println(v2, v3, v4)
    v9 := []int{v4, v4, v4}
    v10 := GenS1{F0: 1}
    v11 := v5[65535]
    if (!false) {
        v3 -= (v3 + v3)
        v2 = float32(float64(1e10))
    } else {
        var v12 uint16 = 65535
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 int16 = 1000
            var v15 bool = false
            v16 := GenS1{F0: v12}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = v12
    }
    if (v6 <= 0) {
        if (false || (true && false)) {
            v11 = -1
            var v17 int32 = 2147483647
            _ = v17
        }
        if true {
            v18 := GenS2{F0: 522, F1: 1}
            _ = v18
        } else {
            v6 += (^(v6))
        }
    } else {
        var v19 float64 = (((-2.25 * -2.25) / (3.14159 - 1.5)) - -2.25)
        if false {
            v19 += (((((v19 / ((1e10 + (((1e10 + (v19 / 1.5)) - v19) * v19)) / float64(v2))) - v19) - v19) + v19) - 1.5)
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 int32 = -2147483648
                v3 -= 18446744073709551615
                var v22 int = v11
                _ = fi20
                _ = v21
                _ = v22
            }
            v6 = (v6 % (((v6 ^ 1)) | 1))
        }
        for fi23 := 0; fi23 < 6; fi23++ {
            v3 = (^(v3))
            fmt.Println(v11)
            v3 -= 760
            _ = fi23
        }
        _ = v19
    }
    _ = len(v5)
    var v24 int32 = int32(v4)
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := GenS3{F0: 255, F1: 42, F2: false}
        fmt.Println(v6, fi25)
        v11 = (v4 ^ fi25)
        _ = fi25
        _ = v26
    }
    fmt.Println(v4)
    if true {
        v27 := GenS1{F0: 0}
        fmt.Println(v4, v3, v11)
        delete(v5, 347)
        _ = v27
    } else {
        switch v4 {
        case 515:
            delete(v5, genFn0(v6, -78))
        case 402:
            fmt.Println(v2, v3)
        case 1:
            v28 := GenS2{F0: (v6 % (((v6 / ((v6) | 1))) | 1)), F1: 0}
            _ = v28
        }
        if (!(113 <= 2)) {
            _ = len(v5)
            for fi29 := 0; fi29 < 13; fi29++ {
                v30 := GenS4{F0: 1000}
                if true {
                    v31 := GenS1{F0: 68}
                    _ = v31
                } else {
                    fmt.Println(v4, v6)
                }
                _ = fi29
                _ = v30
            }
        } else {
            var v32 uint32 = 0
            for fi33 := 0; fi33 < 1; fi33++ {
                fmt.Println(fi33, v6, v6)
                _ = fi33
            }
            _ = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
