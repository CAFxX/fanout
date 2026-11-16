package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 int16; F2 bool }

func main() {
    v1 := []uint32{}
    var v2 bool = true
    v3 := make(map[int]bool)
    var v4 int16 = 1
    var v5 int = 127
    v6 := GenS1{F0: -2147483648, F1: 4294967295}
    var v7 int = (^(v5))
    var v8 int16 = v4
    if (v2 || ((((float32(float64((((0.1 + 1e10) + 0.1) + 1.5))) - (1.5 * (-2.25 / 1.5))) * -2.25) != 0.1) || (-2147483648 <= -2147483648))) {
        v4 = (v8 * v8)
        if (2 < 1) {
            for fi9 := 0; fi9 < 16; fi9++ {
                v8 += (^(v4))
                for fi10 := 0; fi10 < 7; fi10++ {
                    if v2 {
                        var v11 int64 = 582
                        if v2 {
                            v2 = ((0.0 - float32(3.14159)) == ((-2.25 + -2.25) / float32(0.1)))
                        }
                        _ = v11
                    }
                    _ = fi10
                }
                _ = len(v3)
                _ = fi9
            }
            if true {
                var v12 uint = 662
                var v13 float64 = (0.0 - float64((1e10 + -2.25)))
                _ = v12
                _ = v13
            } else {
                var v14 int32 = -1
                _ = v14
            }
        } else {
            for fi15 := 0; fi15 < 20; fi15++ {
                v2 = (!(357 > v4))
                fmt.Println(v8, v8)
                fmt.Println(v2, fi15, v7)
                _ = fi15
            }
        }
        v7 -= 22
    }
    var v16 int32 = -2147483648
    var v17 bool = v2
    v18 := GenS2{F0: -348, F1: v8}
    v5 += v5
    fmt.Println(v7, v17)
    fmt.Println(v2, v8)
    v19 := GenS3{F0: -134}
    var v20 int32 = (v16 & v16)
    v21 := GenS1{F0: 11, F1: 4294967295}
    for fi22 := 0; fi22 < 5; fi22++ {
        v23 := v3[v7]
        _ = fi22
        _ = v23
    }
    if len(v1) > 0 {
        v1[0] = 3
    }
    v24 := GenS4{F0: -32768, F1: (v8 | (v8 >> ((163) & 63))), F2: (721 <= 65535)}
    var v25 int32 = -572
    v26 := GenS2{F0: (v8 / ((42) | 1)), F1: (-(v4))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
