package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint32 }
func genFn0() uint32 {
    v4 := GenS1{F0: 429, F1: 65535}
    var v5 float32 = 0.1
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 int64, p1 uint16) uint16 {
    fmt.Println(p0, p1)
    var v6 uint16 = p1
    var v7 int = -2147483648
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 int64 = 266
    var v2 int = 0
    var v3 uint32 = 2147483648
    v3 = 65535
    if true {
        var v8 string = "\n\t"
        _ = v8
    }
    fmt.Println(v1, v1)
    fmt.Println(v3, v3)
    v9 := GenS2{F0: 1, F1: true, F2: (v3 + 697)}
    v2 -= (-(v2))
    var v10 int8 = 0
    var v11 uint32 = genFn0()
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 float64 = (3.14159 / 1.5)
        v10 = v10
        var v14 float64 = (v13 / (v13 + float64(0.0)))
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint8 = 255
    var v16 int8 = v10
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint64 = 4294967295
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: 18446744073709551615, F1: 2}
    v20 := []int{}
    fmt.Println(v10, v3)
    if false {
        if true {
            for fi21 := 0; fi21 < 14; fi21++ {
                if true {
                    v11 = (v3 + genFn0())
                    v2 = fi21
                } else {
                    fmt.Println(v2, fi21)
                }
                if false {
                    v10 = v16
                    var v22 string = "a"
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 2; fi23++ {
                        fmt.Println(fi23, v10, v15)
                        v1 -= (-(v1))
                        v1 -= (v1 % ((v1) | 1))
                        _ = fi23
                    }
                }
                if (false || (!(uint32(v15) <= 2147483648))) {
                    fmt.Println(fi21)
                    var v24 int = int(v1)
                    for fi25 := 0; fi25 < 11; fi25++ {
                        var v26 bool = false
                        v11 = 0
                        var v27 uint16 = 0
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    _ = v24
                } else {
                    v1 -= 411
                }
                _ = fi21
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v9
}
