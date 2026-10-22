package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int16 }
type GenS2 struct { F0 int32; F1 uint32; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 float32 = 1.5
    var v5 int8 = 0
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: v6, F1: 0, F2: int16(v1)}
    v8 := GenS2{F0: v1, F1: 42, F2: v6}
    if true {
        if (true && false) {
            fmt.Println(v1, v1, v5)
            v9 := GenS3{F0: "\n\t"}
            _ = v9
        } else {
            fmt.Println(v4, v6, v5)
            if true {
                v1 -= v1
                v10 := []int{}
                _ = v10
            } else {
                var v11 bool = false
                _ = v11
            }
            v1 += v1
        }
    }
    v12 := GenS3{F0: ("foo bar" + "")}
    var v13 int16 = 42
    var v14 string = "\n\t"
    v6 = (v6 << ((18446744073709551615) & 63))
    for fi15 := 0; fi15 < 6; fi15++ {
        fmt.Println(v6, v13)
        v16 := []int{fi15}
        v17 := GenS3{F0: ""}
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS3{F0: (v14 + "\n\t")}
    if false {
        for fi19 := 0; fi19 < 5; fi19++ {
            if true {
                var v20 float64 = (0.0 + -2.25)
                if ((0 != 2) || (!true)) {
                    var v21 int64 = 550
                    _ = v21
                } else {
                    v6 -= (v6 + (v6 ^ (v6 % ((v6) | 1))))
                    if false {
                        v4 = 0.0
                        v22 := []int{-1, (fi19 / (((^(fi19))) | 1))}
                        v23 := GenS4{F0: v14}
                        _ = v22
                        _ = v23
                    } else {
                        var v24 uint = 18446744073709551615
                        v25 := []int{-9223372036854775808, fi19, (fi19 / (((fi19 | (fi19 * fi19))) | 1)), (fi19 - fi19)}
                        _ = v24
                        _ = v25
                    }
                }
                _ = v20
            } else {
                var v26 int8 = v5
                v27 := []int{(fi19 & (fi19 & fi19))}
                if ((v1 % (((v1 << ((275) & 63))) | 1)) == 2147483647) {
                    delete(v3, (fi19 & (fi19 >> ((255) & 63))))
                    fmt.Println(v13, v14, v1)
                }
                _ = v26
                _ = v27
            }
            v28 := GenS1{F0: v6, F1: 18446744073709551615, F2: 492}
            _ = fi19
            _ = v28
        }
        _ = len(v3)
    } else {
        switch "foo bar" {
        case "foo bar":
            v6 = (v6 << ((1) & 63))
        default:
            delete(v2, 262)
        }
    }
    switch ("" + v14) {
    case "a":
        v29 := GenS4{F0: v14}
        _ = v29
    case "hello":
        if (!(true && false)) {
            v30 := GenS5{F0: -652}
            var v31 uint = 42
            _ = v30
            _ = v31
        }
    }
    v32 := []int{}
    var v33 int64 = (v6 - (v6 * (^(v6))))
    v14 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
