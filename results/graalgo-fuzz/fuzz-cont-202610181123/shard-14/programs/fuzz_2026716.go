package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 bool; F1 string; F2 int32 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int16 = -607
    v2 := make(map[uint32]string)
    var v3 int16 = (v1 | -1)
    var v4 int64 = -9223372036854775808
    var v5 uint = 683
    var v6 int16 = (v1 / ((0) | 1))
    var v7 uint64 = 67
    if true {
        v8 := v2[598]
        fmt.Println(v8, v7)
        _ = v8
    } else {
        v9 := []int{-2147483648, -1, 32768}
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: 128, F1: false}
        v12 := []int{411, fi10}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int16 = (v1 >> ((v5) & 63))
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := v2[uint32(v7)]
        for fi16 := 0; fi16 < 5; fi16++ {
            delete(v2, 65535)
            _ = fi16
        }
        if (v6 < 42) {
            switch v15 {
            case "":
                if false {
                    if (-9223372036854775808 == v4) {
                        v13 -= (^(v1))
                        _ = len(v2)
                    } else {
                        var v17 int16 = v6
                        _ = v17
                    }
                    v18 := []int{fi14, (^(fi14))}
                    _ = v18
                } else {
                    v19 := GenS2{F0: false, F1: "\n\t", F2: -119}
                    _ = v19
                }
            case "\n\t":
                v20 := GenS3{F0: (v7 / (((v7 * v7)) | 1))}
                _ = v20
            }
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 uint32 = 2
                _ = fi21
                _ = v22
            }
        }
        _ = fi14
        _ = v15
    }
    var v23 int64 = (v4 * v4)
    fmt.Println(v23, v4)
    if false {
        if true {
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 float32 = float32(float64((0.0 / 3.14159)))
                _ = fi24
                _ = v25
            }
            var v26 int = -2147483648
            _ = v26
        } else {
            var v27 int64 = -1
            v5 += 329
            _ = v27
        }
    } else {
        v28 := GenS3{F0: (v7 / ((v7) | 1))}
        if true {
            var v29 uint = 2
            v6 = v3
            _ = v29
        } else {
            for fi30 := 0; fi30 < 3; fi30++ {
                var v31 float64 = (3.14159 * 1e10)
                var v32 uint64 = (v7 - v7)
                _ = fi30
                _ = v31
                _ = v32
            }
            v33 := GenS3{F0: (v7 & 733)}
            _ = v33
        }
        var v34 int8 = 100
        _ = v28
        _ = v34
    }
    v1 = -32768
    v13 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
