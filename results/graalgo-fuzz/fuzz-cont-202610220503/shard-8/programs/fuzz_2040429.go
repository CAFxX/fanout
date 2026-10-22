package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint; F1 int8; F2 int64 }
func genFn0() int8 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 uint16 = 2
        if (!(fi5 <= fi5)) {
            v6 -= (v6 >> ((1) & 63))
            var v7 int16 = -294
            fmt.Println(v7)
            _ = v7
        } else {
            v8 := GenS1{F0: 877, F1: 42}
            _ = v8
        }
        v9 := []int{fi5, 32768, (fi5 - (fi5 / (((fi5 & (fi5 & (fi5 ^ fi5)))) | 1))), 4294967295}
        _ = fi5
        _ = v6
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 int32 = -1
        v12 := GenS1{F0: 42, F1: 1}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 bool = ((1.5 != ((((float64(0.0) + -2.25) * -2.25) - float64(3.14159)) - ((3.14159 * (((0.1 - 1.5) - (1.5 * 0.1)) * -2.25)) - (0.1 / 3.14159)))) || ((!(-1 > 0)) && true))
    var v14 bool = v13
    _ = v13
    _ = v14
    return -8
}
func genFn1() string {
    v15 := GenS1{F0: 696, F1: 0}
    for fi16 := 0; fi16 < 14; fi16++ {
        fmt.Println(fi16)
        _ = fi16
    }
    var v17 bool = (false && false)
    var v18 int64 = -9223372036854775808
    switch -720 {
    case 0:
        v19 := GenS2{F0: -476}
        _ = v19
    case -220:
        var v20 uint8 = 0
        _ = v20
    }
    _ = v15
    _ = v17
    _ = v18
    return "foo bar"
}

func main() {
    var v1 float32 = ((0.1 + (0.1 - 0.1)) / (3.14159 / (-2.25 / 1.0)))
    var v2 uint32 = 65535
    v3 := []uint32{v2, v2}
    var v4 int16 = 796
    if true {
        fmt.Println(v1, v4, v4)
    }
    v21 := GenS2{F0: 641}
    v22 := GenS3{F0: 1, F1: "foo bar"}
    var v23 float64 = 1.5
    v1 = v1
    fmt.Println(v1, v23, v23)
    var v24 int8 = genFn0()
    v25 := GenS4{F0: 255, F1: (v24 ^ -63), F2: 981}
    v3 = append(v3, (v2 + (v2 & (v2 / (((v2 / (((v2 - (v2 / ((889) | 1)))) | 1))) | 1)))))
    v26 := GenS3{F0: -187, F1: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
