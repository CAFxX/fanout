package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }

func main() {
    var v1 float32 = (3.14159 * -2.25)
    var v2 uint8 = 17
    v3 := make(map[int]int)
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := []int{fi5, fi5, 937}
        v2 -= v2
        _ = fi5
        _ = v6
    }
    v2 -= v2
    v1 = v1
    if true {
        fmt.Println(v2)
        var v7 float32 = 3.14159
        fmt.Println(v1)
        _ = v7
    } else {
        var v8 string = ""
        fmt.Println(v8, v2, v2)
        fmt.Println(v8, v1, v2)
        _ = v8
    }
    switch -9223372036854775808 {
    case -471:
        v9 := []int{145}
        _ = v9
    default:
        var v10 int8 = 1
        _ = v10
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        if false {
            var v12 int64 = -339
            v1 -= v1
            _ = v12
        }
        fmt.Println(v2)
        fmt.Println(v1, v2)
        _ = fi11
    }
    if (!(42 == 1)) {
        var v13 int = 585
        v14 := GenS1{F0: true, F1: 42}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        if (-122 == 1000) {
            var v16 int64 = 0
            v1 = (v1 * ((v1 - v1) + v1))
            fmt.Println(v16, v16)
            _ = v16
        } else {
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 uint64 = 242
                var v19 int = fi17
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 uint8 = 221
            _ = v20
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 string = ("foo bar" + "")
            _ = fi21
            _ = v22
        }
        switch fi15 {
        case -1:
            _ = len(v4)
        case -9223372036854775808:
            for fi23 := 0; fi23 < 18; fi23++ {
                v1 -= (v1 + float32((1e10 / 1e10)))
                var v24 uint8 = 128
                fmt.Println(v24, v2, v2)
                _ = fi23
                _ = v24
            }
        default:
            var v25 int64 = 9223372036854775807
            _ = v25
        }
        _ = fi15
    }
    v26 := GenS1{F0: (0 != 569), F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
