package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]string)
    var v3 uint8 = 80
    var v4 int64 = 9223372036854775807
    fmt.Println(v3, v1, v4)
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 int64 = (v4 | v4)
        switch -226 {
        case -692:
            v7 := []int{32767, (fi5 % (((fi5 % (((fi5 * fi5)) | 1))) | 1)), fi5, fi5}
            _ = v7
        case -1:
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := GenS1{F0: 0, F1: 286}
                var v10 uint = 255
                _ = fi8
                _ = v10
                _ = v9
            }
        default:
            v6 = 396
        }
        _ = fi5
        _ = v6
    }
    var v11 string = "foo bar"
    v12 := GenS1{F0: 656, F1: -1}
    for fi13 := 0; fi13 < 1; fi13++ {
        switch (fi13 + (fi13 / (((fi13 + -9223372036854775808)) | 1))) {
        case 42:
            fmt.Println(v3)
        case 2147483647:
            fmt.Println(fi13, v11)
        default:
            var v14 float64 = 0.0
            _ = v14
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 uint64 = 972
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            v18 := GenS1{F0: 506, F1: 0}
            var v19 uint16 = 142
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi13
    }
    if ((v3 & (v3 | 128)) > (v3 / ((v3) | 1))) {
        v11 = ("" + v1)
        v4 += v4
    }
    fmt.Println(v3)
    if true {
        if (!((1 == 2147483647) && (4294967295 != 4294967295))) {
            var v20 uint = 2
            _ = v20
        }
    } else {
        if true {
            var v21 uint64 = 0
            _ = v21
        } else {
            var v22 uint16 = 65535
            fmt.Println(v22, v11, v22)
            v23 := v2[0]
            _ = v22
            _ = v23
        }
        var v24 int16 = 1
        _ = v24
    }
    var v25 uint8 = 42
    var v26 uint16 = 211
    var v27 float32 = ((0.1 * 1e10) / 1.0)
    v28 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
