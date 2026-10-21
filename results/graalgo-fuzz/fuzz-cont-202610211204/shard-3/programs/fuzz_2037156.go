package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint16 = 42
    var v2 uint32 = 274
    var v3 int32 = 0
    var v4 string = "a"
    var v5 uint16 = v1
    v6 := GenS1{F0: 128}
    v7 := []int{}
    var v8 uint8 = 255
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            v8 = 247
            switch (v4 + v4) {
            case "\n\t":
                if false {
                    fmt.Println(v8, v4)
                }
            }
            fmt.Println(v8, fi9, fi9)
            _ = fi10
        }
        switch 32768 {
        case -1:
            fmt.Println(v2, v2)
        case 1:
            v4 = v4
        case 2147483647:
            for fi11 := 0; fi11 < 3; fi11++ {
                v4 = (v4 + "a")
                if (v8 >= v8) {
                    if len(v7) > 0 {
                        v7[0] = fi11
                    }
                }
                _ = fi11
            }
        default:
            v5 -= (v5 - v5)
        }
        fmt.Println(fi9)
        _ = fi9
    }
    v2 = v2
    if true {
        var v12 int16 = -1
        _ = v12
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            v5 = 1
            for fi14 := 0; fi14 < 17; fi14++ {
                v15 := GenS2{F0: (-(fi14)), F1: 1}
                _ = len(v7)
                v3 = 1
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        fmt.Println(v8, v3, v1)
    }
    v1 += uint16(v8)
    v16 := GenS2{F0: 42, F1: int64(v8)}
    if true {
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := GenS2{F0: (^(fi17)), F1: -83}
            _ = fi17
            _ = v18
        }
    } else {
        if (-1 < -62) {
            var v19 float64 = 0.1
            _ = v19
        }
        fmt.Println(v4, v5)
        var v20 uint16 = (v5 | 2)
        _ = v20
    }
    v21 := GenS3{F0: -1}
    var v22 int = 0
    fmt.Println(v2)
    v23 := GenS1{F0: 9223372036854775807}
    v24 := GenS2{F0: (v22 * v22), F1: -9223372036854775808}
    var v25 int64 = -640
    v7 = append(v7, v22)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
