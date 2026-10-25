package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 bool }

func main() {
    var v1 bool = true
    var v2 uint = 75
    var v3 int32 = 1
    v4 := []int{}
    var v5 uint8 = 1
    var v6 uint8 = v5
    v7 := GenS1{F0: true}
    var v8 uint32 = 42
    var v9 uint64 = 4294967295
    for fi10 := 0; fi10 < 2; fi10++ {
        if v1 {
            switch fi10 {
            case 167:
                v11 := GenS1{F0: false}
                _ = v11
            case 82:
                v5 = 155
            default:
                v12 := GenS1{F0: ("foo bar" < "hello")}
                _ = v12
            }
            v2 += v2
            if false {
                var v13 uint8 = v5
                v14 := GenS2{F0: fi10, F1: (v3 + v3)}
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 4; fi15++ {
                    fmt.Println(fi10, fi15)
                    var v16 int = (fi15 / ((int(v5)) | 1))
                    _ = fi15
                    _ = v16
                }
                v17 := GenS2{F0: 0, F1: (v3 - v3)}
                _ = v17
            }
        }
        _ = fi10
    }
    fmt.Println(v6, v2, v2)
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 float32 = 0.0
        v20 := GenS1{F0: (889 == 4294967295)}
        _ = fi18
        _ = v19
        _ = v20
    }
    switch "\n\t" {
    case "":
        var v21 uint16 = 1
        _ = v21
    default:
        switch "a" {
        case "\n\t":
            v4 = append(v4, -1)
        default:
            v22 := GenS1{F0: ((!(0 <= int64(v3))) && false)}
            _ = v22
        }
    }
    var v23 bool = (65535 >= 32767)
    var v24 string = "\n\t"
    if (false || v23) {
        v25 := GenS3{F0: false}
        v24 = v24
        var v26 uint32 = v8
        _ = v25
        _ = v26
    } else {
        _ = len(v4)
        fmt.Println(v24)
        v27 := GenS3{F0: v1}
        _ = v27
    }
    v28 := GenS1{F0: ((-119 <= -128) || (1.5 >= 3.14159))}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v24 = ((("\n\t" + v24) + v24) + (("hello" + (v24 + (v24 + (v24 + v24)))) + v24))
    v3 -= (v3 >> ((v2) & 63))
    var v29 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
