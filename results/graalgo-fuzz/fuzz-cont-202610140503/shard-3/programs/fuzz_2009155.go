package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0(p0 string, p1 uint) bool {
    if true {
        v6 := []int{}
        _ = v6
    }
    fmt.Println(p1, p1)
    var v7 uint32 = 810
    if (true && false) {
        if (280 == 0) {
            var v8 int32 = 2147483647
            _ = v8
        }
        v9 := []int{-9223372036854775808, -9223372036854775808, -148}
        _ = v9
    } else {
        v10 := GenS1{F0: 1, F1: (p1 | (p1 - (^(p1))))}
        if ((!true) && false) {
            v7 = 527
        } else {
            if true {
                var v11 float32 = ((1e10 - 1e10) / 0.1)
                var v12 string = p0
                v13 := GenS1{F0: 42, F1: 255}
                _ = v11
                _ = v12
                _ = v13
            }
        }
        v7 += (v7 + 2147483647)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    return ((42 == 1) && (true && true))
}

func main() {
    var v1 int16 = -32768
    var v2 uint = 401
    var v3 int64 = 0
    var v4 int32 = -1
    var v5 uint = v2
    var v14 int = -341
    fmt.Println(v2, v2)
    v15 := []int{}
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 float32 = 1.5
        v15 = append(v15, (fi16 - 2147483647))
        var v18 bool = (-2.25 <= -2.25)
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: 42, F1: 4294967295}
    switch -2147483648 {
    case 32768:
        if (!false) {
            if true {
                for fi20 := 0; fi20 < 8; fi20++ {
                    fmt.Println(v3)
                    _ = fi20
                }
                fmt.Println(v2, v5, v5)
            } else {
                var v21 uint8 = 230
                for fi22 := 0; fi22 < 8; fi22++ {
                    for fi23 := 0; fi23 < 16; fi23++ {
                        fmt.Println(v3, v14, fi23)
                        var v24 uint16 = 668
                        v4 += (v4 - v4)
                        _ = fi23
                        _ = v24
                    }
                    v15 = append(v15, fi22)
                    _ = fi22
                }
                v4 = 2147483647
                _ = v21
            }
            v25 := GenS1{F0: 450, F1: 42}
            _ = v25
        }
    case 42:
        var v26 bool = ((v14 * v14) > v14)
        _ = v26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v3 -= (v3 << ((0) & 63))
        fmt.Println(fi27)
        _ = fi27
    }
    v4 += 42
    v5 = v2
    var v28 string = "hello"
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    var v29 int8 = -1
    var v30 int16 = (v1 * (v1 % ((32767) | 1)))
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
