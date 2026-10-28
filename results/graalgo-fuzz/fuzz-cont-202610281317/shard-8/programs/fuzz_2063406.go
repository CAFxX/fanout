package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 string; F2 string }
type GenS3 struct { F0 int32; F1 int64; F2 uint32 }
type GenS4 struct { F0 int8; F1 int32 }
func genFn0() uint8 {
    var v6 int16 = 32767
    v7 := []int{47, 517, 0}
    if (true && ((v6 / ((-32768) | 1)) >= 32767)) {
        if (0 <= -128) {
            v8 := GenS1{F0: 1}
            v6 = v6
            _ = v8
        }
        _ = len(v7)
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        fmt.Println(v6)
        var v9 uint = 1
        _ = v9
    }
    var v10 uint = 4294967295
    v10 -= v10
    _ = v10
    _ = v6
    _ = v7
    return 248
}
func genFn1(p0 bool, p1 int32) uint8 {
    if p0 {
        var v11 float64 = ((0.1 - (float64((float32(3.14159) + 3.14159)) - 3.14159)) * (-2.25 - -2.25))
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            fmt.Println(p1)
            if false {
                v14 := []int{}
                v15 := GenS2{F0: 227, F1: "\n\t", F2: ("" + "a")}
                _ = v14
                _ = v15
            } else {
                v16 := GenS3{F0: p1, F1: 42, F2: 1}
                fmt.Println(p0)
                _ = v16
            }
            _ = fi13
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    return 147
}

func main() {
    var v1 uint = 65535
    var v2 int32 = 1
    var v3 string = ""
    v4 := []uint32{0}
    var v5 uint8 = 157
    var v17 int = -2147483648
    v4 = append(v4, 2147483647)
    v18 := GenS4{F0: 29, F1: v2}
    var v19 uint = 18446744073709551615
    v2 = v2
    if len(v4) > 0 {
        v4[0] = 2147483648
    }
    if false {
        fmt.Println(v2, v19, v17)
        v4 = append(v4, 1)
    }
    var v20 uint64 = 141
    _ = len(v4)
    _ = len(v4)
    if (v17 == v17) {
        if true {
            var v21 string = ((v3 + "a") + ((("hello" + v3) + (v3 + v3)) + v3))
            var v22 int32 = v2
            fmt.Println(v22)
            _ = v21
            _ = v22
        }
    } else {
        var v23 int32 = (v2 % (((v2 >> (((v19 & 1)) & 63))) | 1))
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
