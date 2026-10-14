package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int; F1 uint; F2 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 1
    var v3 uint = 0
    var v4 int32 = -2147483648
    var v5 uint16 = 32767
    var v6 int16 = 205
    v5 = (^(v5))
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := []int{v2, (^(fi7))}
        v2 += fi7
        _ = fi7
        _ = v8
    }
    if true {
        if false {
            v2 += v2
            v4 -= (v4 ^ v4)
        }
        v6 = (v6 | v6)
        fmt.Println(v3, v2)
    } else {
        v6 += (v6 << (((v3 & (v3 % ((v3) | 1)))) & 63))
        if (true || true) {
            var v9 uint = v3
            fmt.Println(v5, v2)
            if (!true) {
                fmt.Println(v9, v4)
                if ((float64(((float32(0.0) + 3.14159) - (1e10 * (3.14159 - 0.1)))) + 0.0) >= 3.14159) {
                    var v10 uint32 = 744
                    _ = v10
                } else {
                    v11 := []int{v2, (v2 - 4294967295), 32768}
                    v12 := v1[v2]
                    for fi13 := 0; fi13 < 7; fi13++ {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        _ = fi13
                    }
                    _ = v11
                    _ = v12
                }
                if false {
                    v6 = 42
                    fmt.Println(v5, v2)
                }
            }
            _ = v9
        } else {
            v6 -= v6
            v14 := GenS1{F0: (v6 - -350), F1: 0}
            _ = v14
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 uint64 = 0
            v17 := GenS2{F0: false}
            v6 = int16(v2)
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    v4 = (v4 / ((v4) | 1))
    var v18 int16 = v6
    v3 = v3
    v5 = 25
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS3{F0: v5, F1: 1}
        for fi21 := 0; fi21 < 3; fi21++ {
            v22 := GenS4{F0: -353, F1: v3, F2: int8(fi19)}
            fmt.Println(v18)
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    var v23 uint16 = 2
    v2 = 763
    var v24 uint32 = 813
    v25 := GenS3{F0: v23, F1: v3}
    var v26 int16 = v18
    v6 -= v18
    v27 := []int{v2, v2, 90}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
