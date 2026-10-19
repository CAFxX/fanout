package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 bool; F1 uint }
func genFn0(p0 uint, p1 int32, p2 int64) uint {
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := GenS1{F0: 342, F1: 0}
        var v6 float32 = (((3.14159 + (-2.25 - -2.25)) + 1e10) + ((3.14159 + (((0.1 + 0.1) * (((-2.25 - 0.0) / float32(1e10)) * (float32((((1e10 * ((-2.25 * 1e10) / ((1.5 - 1e10) - (1.5 * (-2.25 * 1e10))))) - (-2.25 - float64(1e10))) / 1.0)) / 0.1))) - ((-2.25 - 0.0) + 0.1))) * (-2.25 * (1.5 * float32(0.0)))))
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        switch (^(fi7)) {
        case 84:
            p1 = p1
        }
        p2 += (p2 - (p2 >> (((^(p0))) & 63)))
        _ = fi7
    }
    var v8 int16 = 196
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p0
}
func genFn1(p0 bool, p1 int32, p2 int64) uint64 {
    fmt.Println(p2, p1, p0)
    var v9 int = 42
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 42
}

func main() {
    var v1 int64 = 382
    var v2 int64 = v1
    var v3 int32 = -940
    v10 := GenS1{F0: 2147483648, F1: -32768}
    if false {
        fmt.Println(v1)
    }
    var v11 int16 = 973
    var v12 int64 = (v1 - (v2 % ((-14) | 1)))
    if ((654 > 2) || (42 == 797)) {
        v3 += v3
        var v13 int16 = (v11 % (((v11 / (((v11 % ((v11) | 1))) | 1))) | 1))
        _ = v13
    } else {
        var v14 int64 = (v12 + 1)
        var v15 bool = ((1.5 + ((3.14159 / ((0.1 / (3.14159 / -2.25)) / -2.25)) / 1.0)) <= 3.14159)
        _ = v14
        _ = v15
    }
    var v16 int32 = v3
    fmt.Println(v16)
    var v17 float32 = (((-2.25 + ((1e10 / ((1e10 / 3.14159) + -2.25)) / -2.25)) + 1.5) * (1e10 / 1e10))
    v18 := GenS1{F0: 2, F1: -526}
    var v19 float32 = v17
    if (255 >= 1) {
        var v20 float64 = float64(v17)
        v20 = float64(v17)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 15; fi21++ {
            if false {
                fmt.Println(fi21)
                if true {
                    v22 := GenS2{F0: false, F1: 621}
                    if false {
                        var v23 float32 = ((v17 * 1.5) / v19)
                        if ((v19 / 1.5) <= v19) {
                            v2 = (^(v2))
                        } else {
                            var v24 uint64 = 4294967295
                            var v25 uint64 = (v24 << ((425) & 63))
                            var v26 string = ("a" + "")
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        _ = v23
                    }
                    _ = v22
                }
            }
            v27 := []int{}
            v16 += v16
            _ = fi21
            _ = v27
        }
        var v28 uint16 = 1
        _ = v28
    }
    switch 2147483647 {
    case -9223372036854775808:
        v29 := []int{-2147483648}
        _ = v29
    case 127:
        fmt.Println(v11)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
