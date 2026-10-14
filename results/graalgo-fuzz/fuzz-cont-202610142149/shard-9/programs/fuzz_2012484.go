package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int }
type GenS2 struct { F0 bool; F1 uint8; F2 string }
type GenS3 struct { F0 uint64; F1 int16; F2 bool }
func genFn0(p0 uint8) uint16 {
    v4 := GenS1{F0: 42, F1: false, F2: -1}
    p0 -= (p0 / ((p0) | 1))
    if true {
        fmt.Println(p0)
        switch "a" {
        case "\n\t":
            v5 := []int{-9223372036854775808, 662}
            _ = v5
        case "hello":
            p0 -= p0
        case "a":
            p0 += 127
        }
        fmt.Println(p0)
    } else {
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int16 = -374
            _ = fi6
            _ = v7
        }
        v8 := []int{32768, -9223372036854775808, 0}
        v9 := GenS1{F0: 42, F1: (853 <= 1), F2: -517}
        _ = v8
        _ = v9
    }
    var v10 int32 = 1
    _ = p0
    _ = v10
    _ = v4
    return 440
}

func main() {
    var v1 int = -1
    var v2 uint8 = 2
    var v3 uint16 = 231
    var v11 uint32 = 0
    var v12 uint64 = 18446744073709551615
    var v13 int8 = 1
    var v14 uint = 18446744073709551615
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := []int{int(v12), (v1 - 306), 32767, 360}
        if len(v16) > 0 {
            v16[0] = int(v14)
        }
        _ = fi15
        _ = v16
    }
    if false {
        v17 := GenS1{F0: 42, F1: ((true || (true && true)) || true), F2: (v1 ^ 32768)}
        v18 := []int{0}
        _ = v17
        _ = v18
    } else {
        if (42 != -1) {
            v1 = -378
            v19 := []int{}
            _ = v19
        }
        v11 -= (^(v11))
    }
    v11 -= v11
    var v20 float32 = 3.14159
    if (0 >= (v14 / ((v14) | 1))) {
        var v21 int64 = 9223372036854775807
        v22 := []int{-621}
        v12 += 641
        _ = v21
        _ = v22
    }
    v3 = v3
    for fi23 := 0; fi23 < 2; fi23++ {
        if true {
            v24 := GenS2{F0: false, F1: v2, F2: "hello"}
            var v25 int8 = int8(v3)
            _ = v24
            _ = v25
        } else {
            v26 := []int{int(v2), (-(fi23))}
            fmt.Println(v13, v1)
            _ = v26
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            switch "a" {
            case "hello":
                fmt.Println(fi23, v1)
            case "":
                switch "foo bar" {
                case "a":
                    if false {
                        v28 := GenS2{F0: false, F1: 48, F2: ""}
                        if false {
                            v2 += (^(v2))
                            v29 := GenS3{F0: (v12 / ((228) | 1)), F1: -1, F2: false}
                            switch fi27 {
                            case 42:
                                switch fi23 {
                                case 584:
                                    v20 -= v20
                                }
                            case 127:
                                fmt.Println(v12, v2)
                            }
                            _ = v29
                        }
                        var v30 uint16 = 476
                        _ = v28
                        _ = v30
                    }
                case "\n\t":
                    v12 = 710
                case "":
                    var v31 uint16 = (v3 >> ((206) & 63))
                    _ = v31
                default:
                    v14 = 458
                }
            default:
                fmt.Println(v20)
            }
            _ = fi27
        }
        var v32 int16 = 1000
        _ = fi23
        _ = v32
    }
    v33 := []int{int(v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v33
}
