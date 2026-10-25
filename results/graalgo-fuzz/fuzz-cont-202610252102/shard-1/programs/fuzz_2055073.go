package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []int{128, 576, -1, -616}
    var v2 string = ""
    var v3 uint16 = 65535
    var v4 int16 = 0
    var v5 uint16 = v3
    if true {
        v6 := GenS1{F0: 9223372036854775807, F1: 300}
        fmt.Println(v5, v5, v4)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int64 = -1
        v8 = v8
        var v9 uint8 = 128
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        if (!true) {
            fmt.Println(v3, v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v11 := GenS2{F0: 775, F1: v5, F2: 42}
            var v12 uint64 = 1
            _ = v11
            _ = v12
        }
        var v13 uint8 = 2
        _ = fi10
        _ = v13
    }
    if false {
        v14 := GenS3{F0: 42}
        v4 += (v4 >> ((42) & 63))
        _ = v14
    } else {
        v2 = "a"
        for fi15 := 0; fi15 < 8; fi15++ {
            v1 = append(v1, fi15)
            _ = fi15
        }
    }
    var v16 uint8 = 0
    _ = len(v1)
    v17 := GenS3{F0: -388}
    switch 128 {
    case 470:
        switch -2147483648 {
        case 128:
            v3 = (v3 >> ((uint(v16)) & 63))
        case 9223372036854775807:
            v4 += (^(v4))
        default:
            fmt.Println(v3, v3)
        }
    case 128:
        v16 = (^(v16))
    default:
        fmt.Println(v2, v2)
    }
    var v18 uint64 = 0
    for fi19 := 0; fi19 < 4; fi19++ {
        fmt.Println(v16)
        _ = fi19
    }
    fmt.Println(v3)
    fmt.Println(v2)
    v16 -= 127
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
