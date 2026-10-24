package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }
type GenS2 struct { F0 int32; F1 int; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64, p1 uint32, p2 uint64) uint32 {
    v7 := GenS1{F0: 4294967295, F1: 52, F2: 42}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}

func main() {
    var v1 int64 = 157
    var v2 uint16 = 255
    var v3 int32 = 0
    v4 := []uint32{287, 1, 902, 225}
    var v5 bool = (!false)
    var v6 bool = v5
    fmt.Println(v2, v2)
    fmt.Println(v5)
    v9 := GenS2{F0: v3, F1: -2147483648, F2: 65535}
    if len(v4) > 0 {
        v4[0] = 2147483648
    }
    if false {
        var v10 uint8 = 0
        _ = v10
    }
    var v11 uint16 = (v2 ^ (^(v2)))
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 bool = true
        if v13 {
            var v14 float32 = (-2.25 * ((1.5 + float32((1.5 * -2.25))) * 0.0))
            if len(v4) > 0 {
                v4[0] = 495
            }
            _ = v14
        } else {
            var v15 float64 = (1.5 + 3.14159)
            v13 = ((!true) || false)
            if ((-(v1)) != 9223372036854775807) {
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := GenS3{F0: 4294967295}
                    v11 += (v2 % ((v11) | 1))
                    v3 += v3
                    if v6 {
                        v18 := GenS4{F0: v6}
                        _ = v18
                    }
                    _ = fi16
                    _ = v17
                }
            }
            _ = v15
        }
        if v5 {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi12
        _ = v13
    }
    var v19 float32 = 3.14159
    var v20 float64 = (1.5 - float64(v19))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
