package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint32; F1 int16; F2 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    var v3 uint64 = 42
    v4 := []string{"", "hello"}
    v5 := []string{}
    v6 := GenS1{F0: "\n\t"}
    if true {
        switch 114 {
        case 451:
            fmt.Println(v3)
        case 128:
            var v7 uint16 = 32767
            _ = v7
        case 9223372036854775807:
            v3 = v3
        default:
            v8 := GenS1{F0: ""}
            _ = v8
        }
        var v9 uint8 = 128
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 string = "\n\t"
            _ = len(v2)
            if true {
                v12 := GenS2{F0: v9, F1: v9, F2: -1}
                _ = v12
            } else {
                var v13 uint64 = (v3 * v3)
                var v14 int64 = -1
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        fmt.Println(v3)
        var v15 uint = 42
        v2[2] = "foo bar"
        _ = v15
    }
    var v16 int8 = -1
    if false {
        var v17 uint16 = 2
        _ = v17
    } else {
        v5 = append(v5, "a")
        fmt.Println(v3, v3)
    }
    v3 += v3
    switch "hello" {
    case "\n\t":
        delete(v1, 2147483647)
    case "a":
        v18 := GenS3{F0: 766, F1: 32767, F2: "hello"}
        _ = v18
    case "foo bar":
        v19 := GenS4{F0: int32(v16)}
        _ = v19
    default:
        v16 += 42
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        fmt.Println(v3, v3)
        switch fi20 {
        case 128:
            v16 -= (v16 / (((v16 & v16)) | 1))
        }
        var v21 int8 = v16
        _ = fi20
        _ = v21
    }
    var v22 uint8 = 230
    var v23 int8 = 23
    var v24 uint32 = 987
    v25 := GenS5{F0: 9223372036854775807}
    fmt.Println(v23, v22)
    if (32768 >= 2147483647) {
        v22 -= (v22 << ((0) & 63))
        var v26 string = ""
        _ = v26
    }
    fmt.Println(v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
