package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint) uint32 {
    p0 -= 4294967295
    if true {
        var v5 string = ("hello" + "\n\t")
        _ = v5
    }
    _ = p0
    return 693
}
func genFn1(p0 uint, p1 int64) bool {
    switch "a" {
    case "foo bar":
        fmt.Println(p0)
    case "a":
        if true {
            p0 = p0
        }
    }
    var v6 int32 = -999
    _ = p0
    _ = p1
    _ = v6
    return false
}

func main() {
    var v1 int32 = -752
    var v2 int64 = 0
    v3 := []int{}
    var v4 float64 = (1e10 - 0.1)
    var v7 int32 = -554
    v4 = v4
    var v8 uint8 = 127
    if (0.0 > 0.1) {
        if true {
            var v9 int16 = -1
            fmt.Println(v4, v9, v1)
            _ = v9
        } else {
            var v10 uint16 = 42
            _ = v10
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 float64 = float64(3.14159)
            switch fi11 {
            case 42:
                if false {
                    var v13 int16 = 0
                    _ = v13
                }
            default:
                var v14 uint8 = (v8 - v8)
                _ = v14
            }
            switch (fi11 ^ 708) {
            case -2147483648:
                if false {
                    v1 += 42
                    var v15 uint16 = 255
                    v2 = v2
                    _ = v15
                } else {
                    var v16 uint = 1
                    switch fi11 {
                    case -236:
                        var v17 bool = (1 >= 100)
                        _ = v17
                    case -9223372036854775808:
                        var v18 uint16 = 272
                        _ = v18
                    }
                    _ = v16
                }
            case 127:
                switch (fi11 << ((499) & 63)) {
                case -1:
                    v7 = v1
                case 791:
                    if false {
                        _ = len(v3)
                    }
                default:
                    v4 -= v4
                }
            }
            _ = fi11
            _ = v12
        }
    }
    var v19 int32 = (-(v7))
    v20 := GenS1{F0: 32767, F1: v2, F2: -9223372036854775808}
    v21 := GenS2{F0: 0}
    var v22 float64 = v4
    fmt.Println(v4)
    v3 = append(v3, 9223372036854775807)
    v4 = 1e10
    v1 += (v19 / (((v1 | (v1 % ((v1) | 1)))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
