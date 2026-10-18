package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int64; F1 uint16; F2 string }

func main() {
    var v1 string = (("a" + "hello") + "\n\t")
    var v2 int32 = -1
    v3 := make(map[uint32]uint32)
    var v4 int8 = 100
    switch 4294967295 {
    case 42:
        var v5 bool = false
        _ = v5
    case 0:
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 int32 = v2
            _ = fi6
            _ = v7
        }
    case 128:
        v2 -= (^(v2))
    }
    if (v1 >= ((("a" + "a") + "\n\t") + "")) {
        v1 = "\n\t"
    } else {
        fmt.Println(v4, v4)
        if (!true) {
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := GenS1{F0: 255, F1: 770, F2: 563}
                if ((v4 & v4) == 100) {
                    var v10 int8 = v4
                    _ = v10
                } else {
                    v11 := GenS2{F0: 42, F1: v2}
                    v1 = v1
                    switch v1 {
                    case "":
                        switch "" {
                        case "":
                            v1 = v1
                        }
                    }
                    _ = v11
                }
                if false {
                    v1 = v1
                }
                _ = fi8
                _ = v9
            }
        } else {
            v12 := GenS3{F0: 9223372036854775807, F1: 1, F2: ((v1 + "a") + "")}
            v1 = v1
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS2{F0: 0, F1: v2}
        _ = fi13
        _ = v14
    }
    if false {
        v4 += (v4 >> ((672) & 63))
    }
    switch 42 {
    case -56:
        for fi15 := 0; fi15 < 15; fi15++ {
            fmt.Println(v1, fi15)
            _ = fi15
        }
    case 447:
        v2 = v2
    default:
        for fi16 := 0; fi16 < 13; fi16++ {
            v4 = 100
            _ = len(v3)
            _ = fi16
        }
    }
    v17 := GenS1{F0: 65535, F1: 397, F2: 18446744073709551615}
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 int64 = 9223372036854775807
        var v20 bool = (1000 >= 900)
        fmt.Println(v4, fi18)
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint16 = 1
    v22 := []int{128}
    delete(v3, 179)
    for fi23 := 0; fi23 < 19; fi23++ {
        fmt.Println(v2)
        if (!false) {
            var v24 int16 = 42
            delete(v3, 347)
            v2 += -2147483648
            _ = v24
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
