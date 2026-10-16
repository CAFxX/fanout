package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int32, p1 int16, p2 uint) uint {
    if true {
        if (true || true) {
            var v7 bool = true
            if ((3.14159 / 3.14159) != (3.14159 + ((1e10 - 0.0) - (float32(3.14159) * -2.25)))) {
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 float64 = ((1e10 - 3.14159) + ((1e10 * 0.1) * 1e10))
                    _ = fi8
                    _ = v9
                }
                for fi10 := 0; fi10 < 8; fi10++ {
                    p0 += 2147483647
                    _ = fi10
                }
                v11 := GenS1{F0: 0, F1: true}
                _ = v11
            } else {
                v7 = ((!false) && (!((p0 | (p0 % ((-587) | 1))) < p0)))
                p0 = (p0 + p0)
                p1 -= 21
            }
            var v12 uint32 = 879
            _ = v12
            _ = v7
        } else {
            switch -812 {
            case 9223372036854775807:
                p1 -= p1
            }
            if false {
                p1 = p1
                v13 := GenS2{F0: 1, F1: false}
                if false {
                    switch 32768 {
                    case 127:
                        if (!((^(p1)) < 32767)) {
                            switch -9223372036854775808 {
                            case 128:
                                p2 -= 0
                            }
                        } else {
                            var v14 int16 = -32768
                            v15 := GenS1{F0: p0, F1: false}
                            _ = v14
                            _ = v15
                        }
                    }
                }
                _ = v13
            } else {
                p2 += (^(p2))
                var v16 uint8 = 228
                _ = v16
            }
            p2 = (p2 & 4294967295)
        }
    } else {
        fmt.Println(p2, p2, p1)
        fmt.Println(p1, p0, p1)
    }
    var v17 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return (^(v17))
}
func genFn1() int64 {
    var v18 bool = false
    v19 := GenS3{F0: 2147483648}
    v20 := []int{}
    v18 = (-2.25 >= 0.0)
    fmt.Println(v18)
    _ = v18
    _ = v19
    _ = v20
    return -9223372036854775808
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]string)
    var v3 string = ("foo bar" + "\n\t")
    var v4 string = v3
    v5 := make(map[int]uint32)
    var v6 uint32 = 2147483647
    fmt.Println(v6)
    v6 = v6
    v3 = v3
    v1 -= (v1 - (v1 + genFn1()))
    fmt.Println(v4, v6)
    var v21 uint32 = 4294967295
    v22 := []int{-1, 9223372036854775807, -953}
    if len(v22) > 0 {
        v22[0] = 766
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 int64 = v1
        _ = fi23
        _ = v24
    }
    if true {
        v4 = v3
        fmt.Println(v21)
        v22 = append(v22, -2147483648)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
