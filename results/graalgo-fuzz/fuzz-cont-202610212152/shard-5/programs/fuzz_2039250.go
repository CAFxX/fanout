package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int; F1 uint; F2 int32 }
type GenS6 struct { F0 int16; F1 int16; F2 int64 }
type GenS7 struct { F0 int16; F1 bool; F2 int8 }
type GenS8 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS9 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 1
    var v3 float32 = ((1.5 - 1e10) * ((0.0 * (0.1 * 0.1)) + 0.1))
    var v4 int8 = -1
    if true {
        v5 := GenS1{F0: 0, F1: -9223372036854775808, F2: 0}
        for fi6 := 0; fi6 < 12; fi6++ {
            for fi7 := 0; fi7 < 7; fi7++ {
                if true {
                    v4 -= (v4 & -126)
                    fmt.Println(v3, fi6)
                }
                v8 := GenS1{F0: 4294967295, F1: -122, F2: 0}
                v9 := []int{}
                _ = fi7
                _ = v8
                _ = v9
            }
            v10 := GenS2{F0: fi6}
            _ = fi6
            _ = v10
        }
        _ = v5
    }
    v11 := GenS1{F0: 42, F1: 9223372036854775807, F2: 18446744073709551615}
    fmt.Println(v2)
    v12 := GenS3{F0: -2147483648}
    v4 += (v4 >> ((2) & 63))
    if (2147483648 > 2147483648) {
        switch -914 {
        case 32767:
            v13 := []int{-300}
            _ = v13
        case -327:
            var v14 uint16 = 0
            _ = v14
        case -907:
            v15 := GenS4{F0: 255, F1: -121}
            _ = v15
        default:
            if false {
                v16 := GenS3{F0: 128}
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 uint64 = 4294967295
            v18 = v18
            if false {
                v18 = (^(v18))
            } else {
                v19 := []int{fi17}
                v4 -= 0
                v2 = (v2 >> ((155) & 63))
                _ = v19
            }
            _ = fi17
            _ = v18
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            v21 := GenS3{F0: fi20}
            _ = fi20
            _ = v21
        }
    }
    v22 := GenS5{F0: -1, F1: 18446744073709551615, F2: 823}
    v23 := v1[("\n\t" + "")]
    if (0 <= -96) {
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := GenS6{F0: v2, F1: (v2 & (v2 | (v2 | 767))), F2: 9223372036854775807}
            v26 := v1[v23]
            _ = fi24
            _ = v25
            _ = v26
        }
        fmt.Println(v3, v4)
        v27 := GenS7{F0: v2, F1: false, F2: (^(v4))}
        _ = v27
    } else {
        v3 += ((v3 + (((v3 / -2.25) / (0.0 / ((3.14159 + v3) / v3))) + v3)) - (v3 + 1e10))
        _ = len(v1)
    }
    v28 := GenS5{F0: 0, F1: 705, F2: -418}
    v2 = -821
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 uint = 18446744073709551615
        _ = fi29
        _ = v30
    }
    switch "a" {
    case "a":
        switch -656 {
        case 128:
            fmt.Println(v3, v23)
        case 32767:
            v2 -= (v2 / (((v2 & v2)) | 1))
        default:
            if false {
                v31 := GenS8{F0: 1, F1: 21, F2: -100}
                if false {
                    var v32 int8 = -69
                    if (1 >= 691) {
                        fmt.Println(v23)
                        v33 := []int{}
                        fmt.Println(v2)
                        _ = v33
                    }
                    _ = v32
                } else {
                    v23 = (((v23 + ("hello" + v23)) + "hello") + "a")
                }
                v34 := GenS9{F0: 1, F1: -1}
                _ = v31
                _ = v34
            } else {
                v35 := v1["foo bar"]
                v4 = (v4 >> ((42) & 63))
                _ = v35
            }
        }
    default:
        delete(v1, ("hello" + ((v23 + v23) + (("foo bar" + v23) + v23))))
    }
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
}
