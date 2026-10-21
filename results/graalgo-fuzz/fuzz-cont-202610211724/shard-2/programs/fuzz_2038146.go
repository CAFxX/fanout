package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint }

func main() {
    var v1 uint = 0
    v2 := []string{}
    v3 := make(map[int]uint32)
    var v4 int16 = -942
    v5 := []int{}
    var v6 int32 = -1
    for fi7 := 0; fi7 < 6; fi7++ {
        v1 = 707
        v1 -= 42
        _ = fi7
    }
    if (1e10 > 1e10) {
        fmt.Println(v4, v6, v6)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        if (v1 < 4294967295) {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 float32 = 0.1
                _ = len(v3)
                _ = fi8
                _ = v9
            }
        } else {
            v10 := GenS1{F0: (!(true && ((^(v6)) < 42))), F1: (v1 & v1)}
            _ = v10
        }
        if false {
            var v11 uint = 1
            fmt.Println(v11, v6)
            v12 := GenS1{F0: false, F1: (v1 | v1)}
            _ = v11
            _ = v12
        }
    }
    var v13 string = (("" + "\n\t") + (("" + "") + ""))
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 uint64 = 18446744073709551615
            v1 = v1
            if true {
                v6 = v6
            }
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: 0}
        if ((255 != 0) && false) {
            v13 = v13
        }
        _ = fi14
        _ = v17
    }
    v3[128] = 220
    switch -431 {
    case 127:
        var v18 int64 = -1
        _ = v18
    case 4294967295:
        for fi19 := 0; fi19 < 12; fi19++ {
            fmt.Println(v4, v6, v13)
            for fi20 := 0; fi20 < 8; fi20++ {
                fmt.Println(v6, v1, v13)
                fmt.Println(v4, v6)
                v21 := GenS2{F0: 93}
                _ = fi20
                _ = v21
            }
            v22 := GenS1{F0: (true && true), F1: (v1 - (v1 << ((613) & 63)))}
            _ = fi19
            _ = v22
        }
    case 0:
        for fi23 := 0; fi23 < 17; fi23++ {
            v6 = 42
            var v24 string = ""
            _ = fi23
            _ = v24
        }
    default:
        if true {
            var v25 int32 = -2147483648
            _ = v25
        }
    }
    fmt.Println(v4)
    v26 := GenS3{F0: 984, F1: 51, F2: v1}
    switch 32767 {
    case -813:
        if (42 != (v1 % ((v1) | 1))) {
            v1 = 244
            _ = len(v3)
            _ = len(v5)
        }
    case 734:
        v5 = append(v5, 0)
    case 127:
        var v27 int = 4294967295
        _ = v27
    default:
        var v28 uint32 = 4294967295
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
