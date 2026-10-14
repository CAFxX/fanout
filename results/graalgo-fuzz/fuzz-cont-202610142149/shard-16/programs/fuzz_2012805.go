package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint32; F1 int8 }
func genFn0(p0 int8, p1 uint32, p2 int16) int32 {
    p2 += p2
    var v7 int8 = (p0 | (p0 % ((p0) | 1)))
    var v8 uint8 = 1
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 uint, p1 bool) uint16 {
    var v9 int16 = 1000
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v9
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = false
    var v3 string = ""
    var v4 uint = 42
    var v5 uint = (^(v4))
    var v6 string = (v3 + v3)
    fmt.Println(v2, v4, v4)
    var v10 bool = true
    if v2 {
        var v11 bool = false
        if v10 {
            _ = len(v1)
            if (2 != v4) {
                for fi12 := 0; fi12 < 2; fi12++ {
                    v4 -= 2
                    v13 := []int{fi12, 13}
                    _ = len(v1)
                    _ = fi12
                    _ = v13
                }
            }
        }
        _ = v11
    } else {
        v1[(v3 + ((v3 + "") + (("" + v6) + v3)))] = v2
        var v14 uint16 = 65535
        v15 := GenS1{F0: -128, F1: -2147483648}
        _ = v14
        _ = v15
    }
    fmt.Println(v6, v5)
    v16 := GenS1{F0: 100, F1: 0}
    if false {
        v10 = true
        v3 = v3
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS1{F0: -1, F1: 0}
        v19 := GenS1{F0: 1, F1: fi17}
        for fi20 := 0; fi20 < 7; fi20++ {
            if false {
                var v21 float32 = 0.0
                _ = v21
            } else {
                if false {
                    fmt.Println(v2)
                } else {
                    if (1 > -1) {
                        _ = len(v1)
                        var v22 int8 = 100
                        v23 := GenS2{F0: 65535, F1: v22}
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(v3, v3)
                    v10 = v2
                }
                var v24 int16 = -663
                var v25 int32 = int32(fi20)
                var v26 int64 = 42
                _ = v24
                _ = v25
                _ = v26
            }
            var v27 uint16 = 1
            _ = fi20
            _ = v27
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    v28 := GenS1{F0: -99, F1: 32768}
    fmt.Println(v2, v3, v4)
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 uint32 = 0
        v3 = (("foo bar" + v6) + v6)
        _ = fi29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
