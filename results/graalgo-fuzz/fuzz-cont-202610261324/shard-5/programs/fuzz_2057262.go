package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 string; F2 bool }
func genFn0(p0 uint16, p1 uint, p2 uint) int16 {
    var v7 string = "a"
    if true {
        var v8 int = -955
        _ = v8
    } else {
        if true {
            if ((float64(0.0) * 3.14159) >= (0.1 * float64(0.1))) {
                for fi9 := 0; fi9 < 7; fi9++ {
                    fmt.Println(p0, p2, p0)
                    _ = fi9
                }
            }
        } else {
            var v10 float64 = (3.14159 / 3.14159)
            _ = v10
        }
    }
    var v11 uint64 = 155
    for fi12 := 0; fi12 < 16; fi12++ {
        if ((v11 == v11) && false) {
            if (true || (!(65535 < (p0 << ((p1) & 63))))) {
                var v13 string = "\n\t"
                v14 := GenS1{F0: p1}
                var v15 bool = (62 == -9223372036854775808)
                _ = v13
                _ = v14
                _ = v15
            }
            v7 = (v7 + v7)
            v16 := GenS1{F0: p1}
            _ = v16
        } else {
            v17 := GenS1{F0: p1}
            _ = v17
        }
        fmt.Println(p0, p0, p0)
        _ = fi12
    }
    v18 := []int{-951, 142, -697}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v18
    _ = v7
    return -1
}

func main() {
    var v1 int32 = 42
    var v2 int16 = -32768
    v3 := make(map[int]string)
    var v4 float64 = (0.1 - (3.14159 * (-2.25 - (float64(((-2.25 - float32((1.5 / 1.5))) + float32((0.0 + 0.0)))) - 1e10))))
    v5 := make(map[int]uint32)
    var v6 int8 = 0
    v1 = v1
    v2 = genFn0(388, 18446744073709551615, 0)
    for fi19 := 0; fi19 < 1; fi19++ {
        v4 = (v4 + v4)
        v20 := GenS1{F0: 0}
        _ = fi19
        _ = v20
    }
    v1 = (v1 / (((v1 | v1)) | 1))
    var v21 uint64 = 4294967295
    v4 = (v4 - v4)
    var v22 uint16 = 0
    v23 := GenS2{F0: 0, F1: ("" + "a"), F2: ((v6 & v6) > 1)}
    v1 = v1
    var v24 float32 = 0.1
    v6 = (v6 ^ v6)
    if false {
        if true {
            v2 += genFn0(v22, 42, 65535)
            if true {
                v22 -= v22
                if (!true) {
                    v25 := []int{-660, -706, 127}
                    _ = v25
                } else {
                    var v26 uint64 = 1
                    _ = v26
                }
            }
        }
    }
    v4 -= v4
    var v27 int64 = 9223372036854775807
    v28 := []int{127, int(v22), 127, 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
