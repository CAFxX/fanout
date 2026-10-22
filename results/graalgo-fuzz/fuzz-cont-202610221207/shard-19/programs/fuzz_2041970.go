package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int16; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint; F2 int }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 uint }
func genFn0(p0 int8, p1 int8, p2 uint16) uint {
    for fi4 := 0; fi4 < 8; fi4++ {
        switch "hello" {
        case "hello":
            fmt.Println(fi4)
        default:
            for fi5 := 0; fi5 < 8; fi5++ {
                if false {
                    v6 := GenS1{F0: 184}
                    p1 = 17
                    _ = v6
                }
                v7 := []int{128, fi4, -948}
                v8 := GenS2{F0: true, F1: -32768, F2: 0}
                _ = fi5
                _ = v7
                _ = v8
            }
        }
        _ = fi4
    }
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 531
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    var v3 int = -463
    var v10 int64 = 9223372036854775807
    v11 := GenS3{F0: 1, F1: 81, F2: v3}
    v10 = (v10 + (v10 % (((v10 ^ v10)) | 1)))
    v12 := []int{}
    var v13 uint8 = 128
    for fi14 := 0; fi14 < 17; fi14++ {
        if false {
            for fi15 := 0; fi15 < 4; fi15++ {
                fmt.Println(fi14)
                switch 32768 {
                case 128:
                    switch (fi14 & (fi15 / (((-(v3))) | 1))) {
                    case 42:
                        if len(v12) > 0 {
                            v12[0] = fi14
                        }
                    }
                case 127:
                    v3 -= 32767
                case 2147483647:
                    var v16 float32 = -2.25
                    _ = v16
                default:
                    fmt.Println(v13, v13, v10)
                }
                v2[fi14] = 708
                if (!false) {
                    var v17 uint = 255
                    var v18 string = "\n\t"
                    _ = v17
                    _ = v18
                }
                _ = fi15
            }
        } else {
            if ((true && true) || (18446744073709551615 >= 379)) {
                var v19 uint16 = 658
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                var v20 uint32 = 2147483647
                _ = v19
                _ = v20
            } else {
                delete(v2, (fi14 ^ 1))
                var v21 uint8 = 127
                _ = len(v1)
                _ = v21
            }
            if true {
                v13 = 127
                fmt.Println(fi14)
            } else {
                fmt.Println(v3)
            }
            var v22 bool = true
            _ = v22
        }
        v2[fi14] = 0
        _ = fi14
    }
    if (v13 != (v13 << ((65535) & 63))) {
        fmt.Println(v13, v10, v10)
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 int16 = 32767
        v25 := GenS4{F0: 112, F1: 2147483648}
        var v26 uint = 2
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    v27 := GenS3{F0: 2, F1: 681, F2: -1}
    v28 := GenS1{F0: 42}
    if (!(!false)) {
        var v29 bool = false
        _ = v29
    }
    var v30 float64 = 0.1
    var v31 uint = 42
    v32 := GenS5{F0: v31}
    v12 = append(v12, (v3 << ((genFn0(-115, -42, 42)) & 63)))
    if false {
        v3 += v3
        fmt.Println(v30, v13)
        v2[(v3 ^ v3)] = 222
    }
    v31 += (v31 ^ (v31 | 988))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
