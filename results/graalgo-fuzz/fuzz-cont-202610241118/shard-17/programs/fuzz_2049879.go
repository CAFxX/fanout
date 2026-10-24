package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    var v1 bool = true
    var v2 int = 32768
    var v3 string = ("foo bar" + ("" + ""))
    var v4 int32 = -359
    v5 := []uint32{}
    v6 := make(map[uint32]int)
    v1 = false
    v5 = append(v5, 1)
    v4 = -204
    v7 := GenS1{F0: -454, F1: 32767, F2: 9223372036854775807}
    if v1 {
        if v1 {
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
            _ = len(v5)
        }
        var v8 uint8 = 255
        _ = v8
    } else {
        v3 = v3
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 bool = (65535 == 71)
            _ = fi9
            _ = v10
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println(v4, v4, v3)
    v11 := GenS1{F0: 0, F1: 1, F2: -200}
    var v12 int32 = v4
    var v13 float64 = 3.14159
    var v14 int16 = 1000
    var v15 int32 = v12
    var v16 uint32 = 717
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 uint = 0
        if (955 <= 255) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v19 := GenS2{F0: v18, F1: 1}
            _ = v19
        } else {
            v1 = v1
        }
        var v20 int8 = -1
        _ = fi17
        _ = v18
        _ = v20
    }
    var v21 float32 = 3.14159
    switch -904 {
    case 42:
        fmt.Println(v21, v13)
    default:
        v22 := v6[v16]
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        if v1 {
            v24 := GenS2{F0: 0, F1: 1}
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 uint16 = 27
                var v27 int = -753
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = v24
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
