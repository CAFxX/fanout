package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0() uint8 {
    v4 := GenS1{F0: 265, F1: 2}
    v5 := GenS2{F0: 58, F1: -44}
    var v6 uint64 = 1
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{fi7, (fi7 / ((fi7) | 1)), fi7, (fi7 * (fi7 >> ((42) & 63)))}
        var v9 float32 = float32(1.5)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v4
    _ = v5
    _ = v6
    return 233
}
func genFn1() string {
    v10 := []int{127}
    var v11 bool = false
    fmt.Println(v11)
    v11 = v11
    _ = v10
    _ = v11
    return "a"
}

func main() {
    v1 := []int{}
    var v2 int = -2147483648
    v3 := make(map[string]uint32)
    var v12 uint = 18446744073709551615
    var v13 int = (v2 >> (((v12 - v12)) & 63))
    v2 -= v2
    if false {
        if (-1 <= -109) {
            if (65535 >= (^(v12))) {
                v13 = 9223372036854775807
                if (1.5 < 1.5) {
                    var v14 bool = true
                    _ = v14
                } else {
                    v13 = (v2 % (((v13 ^ v13)) | 1))
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v2)
            }
            v15 := GenS3{F0: 88, F1: false}
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 int8 = -76
                _ = fi16
                _ = v17
            }
            _ = v15
        }
    }
    v2 += v13
    var v18 uint = v12
    for fi19 := 0; fi19 < 13; fi19++ {
        v3["hello"] = 198
        _ = fi19
    }
    var v20 string = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}
