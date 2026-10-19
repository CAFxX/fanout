package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 uint = 255
    var v2 uint8 = 252
    v3 := []int{-2147483648, 2147483647}
    v4 := []int{579, 2147483647}
    var v5 int8 = 100
    fmt.Println(v2, v1)
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: 2147483648, F1: (v2 + v2)}
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: 777, F1: 2}
            _ = fi8
            _ = v9
        }
        var v10 float32 = (3.14159 - ((0.1 + 1.5) + 0.0))
        _ = fi6
        _ = v10
        _ = v7
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint8 = (^(v2))
            var v14 uint32 = 2147483647
            if (float64(-2.25) >= 0.0) {
                v13 = (v2 + v13)
                fmt.Println(v13)
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 uint = 42
        _ = fi11
        _ = v15
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    v16 := GenS1{F0: 540, F1: v2}
    v2 += v2
    v2 = (^(v2))
    v5 -= (v5 & 127)
    for fi17 := 0; fi17 < 9; fi17++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if false {
            switch 4294967295 {
            case 0:
                switch "" {
                case "\n\t":
                    if true {
                        var v18 int64 = 404
                        if true {
                            var v19 int8 = int8(v18)
                            var v20 uint = 4294967295
                            _ = v19
                            _ = v20
                        }
                        _ = v18
                    } else {
                        v5 -= (-(v5))
                        fmt.Println(v5, fi17)
                    }
                }
            case -816:
                for fi21 := 0; fi21 < 7; fi21++ {
                    var v22 int16 = 1
                    v22 = v22
                    _ = fi21
                    _ = v22
                }
            default:
                var v23 bool = true
                _ = v23
            }
        } else {
            v1 -= 554
        }
        v1 += v1
        _ = fi17
    }
    if false {
        var v24 int64 = 1
        if len(v3) > 0 {
            v3[0] = -468
        }
        _ = v24
    } else {
        var v25 uint = v1
        v26 := GenS1{F0: 1, F1: v2}
        _ = v25
        _ = v26
    }
    if len(v3) > 0 {
        v3[0] = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
