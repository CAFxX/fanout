package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint; F1 uint16; F2 uint64 }
func genFn0(p0 int32, p1 int) uint16 {
    v5 := []int{}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    p1 = (p1 + 778)
    p1 -= p1
    var v6 float64 = ((0.0 + (-2.25 + (-2.25 / 3.14159))) - ((((0.0 + float64(1e10)) + 0.0) + 1.5) / ((1.5 - (3.14159 - 3.14159)) * 1.5)))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint16 = 2
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]int)
    fmt.Println(v2, v3)
    fmt.Println(v2, v3, v2)
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            if true {
                var v9 uint32 = 65535
                var v10 string = "hello"
                v9 += (v9 - (v9 - (^(v9))))
                _ = v10
                _ = v9
            } else {
                v11 := []int{fi7, fi7}
                if true {
                    switch 1 {
                    case 2147483647:
                        v12 := GenS1{F0: 42}
                        _ = v12
                    case 461:
                        var v13 uint16 = (v2 + (v2 % (((v2 / (((v2 + v2)) | 1))) | 1)))
                        _ = v13
                    case 127:
                        var v14 int64 = 259
                        _ = v14
                    }
                } else {
                    v1 = v1
                    v2 = (v2 * (^(v2)))
                    var v15 int8 = 127
                    v15 += (v15 | v15)
                    _ = v15
                }
                _ = v11
            }
            var v16 float64 = 1.5
            _ = fi8
            _ = v16
        }
        v2 = genFn0((v1 ^ -731), 674)
        v17 := GenS1{F0: -9223372036854775808}
        _ = fi7
        _ = v17
    }
    v3 += (v3 ^ 245)
    var v18 int64 = 9223372036854775807
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := v4[fi19]
        v21 := GenS2{F0: (v3 - 79), F1: "hello"}
        switch (v20 * v20) {
        case 53:
            var v22 string = ("foo bar" + "hello")
            _ = v22
        case 128:
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 int32 = (v1 | (-(v1)))
                v3 -= v3
                var v25 int16 = 0
                _ = fi23
                _ = v24
                _ = v25
            }
        default:
            fmt.Println(fi19)
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    if (32768 != 1) {
        v26 := GenS3{F0: 255, F1: 65535, F2: (v3 % (((v3 << ((414) & 63))) | 1))}
        _ = v26
    } else {
        var v27 uint64 = 69
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        for fi29 := 0; fi29 < 12; fi29++ {
            v30 := GenS3{F0: 4294967295, F1: v2, F2: v3}
            var v31 int32 = v1
            _ = fi29
            _ = v30
            _ = v31
        }
        fmt.Println(v2, v3, v1)
        _ = fi28
    }
    v32 := []int{0}
    v32 = append(v32, -2147483648)
    var v33 uint = 255
    var v34 uint32 = 0
    fmt.Println(v34, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
