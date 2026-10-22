package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0(p0 int16, p1 int32, p2 string) uint32 {
    v4 := []int{-1, -1, 688, 32768}
    for fi5 := 0; fi5 < 19; fi5++ {
        if (false && true) {
            var v6 uint32 = 759
            _ = v6
        }
        p0 = -345
        _ = fi5
    }
    var v7 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 155
}
func genFn1(p0 int32, p1 int32, p2 int32) uint16 {
    var v8 uint8 = 128
    fmt.Println(v8, p1)
    v9 := GenS1{F0: 2, F1: "hello"}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 2
    v3 := make(map[uint32]string)
    fmt.Println(v2)
    if (((0.0 / (((0.1 * 1e10) / 3.14159) / 3.14159)) - 1.5) < 3.14159) {
        if (true || (0 >= -32768)) {
            var v10 int32 = 1
            for fi11 := 0; fi11 < 16; fi11++ {
                fmt.Println(v10, v10, v10)
                if false {
                    if (fi11 == fi11) {
                        v10 = 42
                        v10 += 0
                        var v12 uint8 = 128
                        _ = v12
                    }
                    var v13 uint32 = 146
                    _ = v13
                }
                if true {
                    var v14 bool = true
                    v15 := GenS1{F0: 2, F1: "foo bar"}
                    for fi16 := 0; fi16 < 16; fi16++ {
                        var v17 int32 = (v10 / (((-(v10))) | 1))
                        _ = fi16
                        _ = v17
                    }
                    _ = v14
                    _ = v15
                }
                _ = fi11
            }
            var v18 uint = 2
            _ = v10
            _ = v18
        } else {
            v2 = v2
        }
        v2 -= v2
        v2 = v2
    }
    if false {
        v2 = 427
    } else {
        v19 := GenS1{F0: 887, F1: "hello"}
        _ = v19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        fmt.Println(v2, fi20)
        _ = fi20
    }
    v2 = 42
    v2 += (v2 - (v2 - v2))
    v21 := v1[-853]
    var v22 int16 = 1
    if false {
        v23 := GenS1{F0: 255, F1: "a"}
        _ = v23
    } else {
        fmt.Println(v21)
    }
    v21 += v21
    v21 = v21
    var v24 int32 = 1
    switch 343 {
    case -779:
        var v25 float32 = (3.14159 - -2.25)
        _ = v25
    case 127:
        v2 += (v2 / ((255) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
}
