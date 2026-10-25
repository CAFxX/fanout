package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint }
func genFn0(p0 uint16, p1 bool) uint8 {
    v6 := []int{-9223372036854775808, 1, -9223372036854775808, -400}
    v6 = append(v6, -88)
    _ = p0
    _ = p1
    _ = v6
    return 2
}

func main() {
    var v1 int64 = 42
    var v2 string = "hello"
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]int)
    var v5 bool = ((4294967295 == 451) && true)
    v5 = (!false)
    var v7 int16 = 0
    if (((v2 + "hello") + v2) == "hello") {
        fmt.Println(v2, v5, v2)
        fmt.Println(v5, v7, v3)
    }
    var v8 uint16 = 65535
    var v9 bool = false
    v10 := []int{-1, 26, 497, -2147483648}
    var v11 uint32 = 2147483647
    var v12 int8 = 127
    var v13 uint = 289
    var v14 int8 = (v12 << ((v13) & 63))
    if (!true) {
        _ = len(v10)
        for fi15 := 0; fi15 < 15; fi15++ {
            fmt.Println(fi15, v1, v12)
            switch ("a" + v2) {
            case "":
                v16 := GenS1{F0: (^(v13)), F1: (^(v14))}
                _ = v16
            case "\n\t":
                for fi17 := 0; fi17 < 6; fi17++ {
                    fmt.Println(v5, v8)
                    if false {
                        v18 := GenS1{F0: v13, F1: (v12 | int8(v7))}
                        _ = v18
                    } else {
                        v19 := GenS1{F0: (v13 ^ (v13 << ((uint(v3)) & 63))), F1: (v14 >> ((v13) & 63))}
                        fmt.Println(v5)
                        _ = len(v4)
                        _ = v19
                    }
                    _ = fi17
                }
            default:
                switch v2 {
                case "a":
                    var v20 bool = v9
                    _ = v20
                case "\n\t":
                    v21 := GenS1{F0: (v13 * (v13 * 65535)), F1: int8(v11)}
                    _ = v21
                }
            }
            _ = fi15
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 float64 = 1.5
            var v24 float64 = v23
            _ = fi22
            _ = v23
            _ = v24
        }
    } else {
        if v9 {
            if (float64((1.5 - ((3.14159 / 1.5) / ((1.5 + (1.5 - 0.1)) - 1.5)))) > 1.5) {
                v25 := GenS1{F0: 42, F1: -45}
                v26 := GenS2{F0: v8, F1: 0, F2: (v8 & (v8 - (v8 & (v8 ^ (v8 | uint16(v3))))))}
                switch (("a" + v2) + v2) {
                case "":
                    v8 -= uint16(v3)
                }
                _ = v25
                _ = v26
            }
            v27 := GenS1{F0: (^(v13)), F1: v12}
            _ = v27
        } else {
            v28 := GenS1{F0: uint(v11), F1: int8(v1)}
            for fi29 := 0; fi29 < 14; fi29++ {
                var v30 int64 = -1
                if false {
                    v31 := GenS1{F0: v13, F1: (^(v12))}
                    _ = len(v4)
                    v3 = (v3 + (v3 + (^(v3))))
                    _ = v31
                }
                for fi32 := 0; fi32 < 8; fi32++ {
                    fmt.Println(fi29, v1)
                    fmt.Println(v2, v30)
                    _ = fi32
                }
                _ = fi29
                _ = v30
            }
            _ = v28
        }
        var v33 int32 = 2147483647
        _ = v33
    }
    if true {
        for fi34 := 0; fi34 < 17; fi34++ {
            switch fi34 {
            case 9223372036854775807:
                for fi35 := 0; fi35 < 3; fi35++ {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = fi35
                }
            }
            for fi36 := 0; fi36 < 10; fi36++ {
                _ = len(v10)
                fmt.Println(v12, v5, v2)
                _ = fi36
            }
            fmt.Println(v3, v5, v9)
            _ = fi34
        }
    }
    v37 := GenS3{F0: (v1 / (((v1 + (v1 + 366))) | 1)), F1: (v13 / (((v13 >> ((v13) & 63))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
