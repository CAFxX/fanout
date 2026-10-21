package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 int; F1 uint; F2 int64 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 string; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 42
    v3 := make(map[int]string)
    v4 := []string{"hello"}
    if (543 > 18446744073709551615) {
        var v5 int8 = 29
        var v6 uint = 255
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: 65535, F1: 255, F2: 0}
        _ = v7
    }
    v4 = append(v4, (("foo bar" + (("" + "a") + "")) + "foo bar"))
    if false {
        if true {
            v2 = 1
        }
        for fi8 := 0; fi8 < 15; fi8++ {
            v2 = 0
            _ = fi8
        }
    } else {
        v9 := GenS2{F0: -262, F1: 220, F2: -220}
        if (0 <= 62) {
            v10 := GenS2{F0: 1, F1: uint(v2), F2: 42}
            v11 := GenS2{F0: 32767, F1: 2, F2: 0}
            fmt.Println(v2)
            _ = v10
            _ = v11
        }
        _ = v9
    }
    v2 -= v2
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    var v12 int64 = -9223372036854775808
    v13 := GenS3{F0: 32768, F1: (v2 - v2)}
    var v14 string = "\n\t"
    if false {
        for fi15 := 0; fi15 < 8; fi15++ {
            if true {
                fmt.Println(v2, v14, v14)
                switch fi15 {
                case 127:
                    for fi16 := 0; fi16 < 8; fi16++ {
                        var v17 int64 = 0
                        fmt.Println(v12, fi15)
                        _ = fi16
                        _ = v17
                    }
                case 4294967295:
                    _ = len(v3)
                default:
                    v18 := GenS4{F0: v14, F1: 125}
                    _ = v18
                }
            }
            var v19 int = fi15
            _ = fi15
            _ = v19
        }
        v14 = "foo bar"
        v20 := GenS3{F0: -9223372036854775808, F1: v2}
        _ = v20
    } else {
        v12 = 1
    }
    var v21 int32 = 952
    if true {
        for fi22 := 0; fi22 < 19; fi22++ {
            v21 += -2147483648
            _ = fi22
        }
        if true {
            var v23 int64 = (v12 ^ 42)
            _ = v23
        } else {
            var v24 uint32 = v2
            var v25 float64 = float64(-2.25)
            var v26 uint8 = 2
            _ = v24
            _ = v25
            _ = v26
        }
        var v27 uint64 = 2
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
