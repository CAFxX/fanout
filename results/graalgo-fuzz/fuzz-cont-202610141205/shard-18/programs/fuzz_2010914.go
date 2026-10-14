package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int; F2 int32 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 19; fi4++ {
        if (0 >= -1) {
            switch -571 {
            case -9223372036854775808:
                var v5 float64 = -2.25
                _ = v5
            case 47:
                fmt.Println(fi4)
            default:
                var v6 int64 = 42
                _ = v6
            }
            var v7 float32 = -2.25
            _ = v7
        }
        v8 := GenS1{F0: 1000, F1: 32767}
        _ = fi4
        _ = v8
    }
    if true {
        v9 := []int{-442, -1, 127, 32768}
        _ = v9
    }
    v10 := []int{-2147483648, 9223372036854775807}
    if true {
        var v11 uint = 483
        var v12 string = ("\n\t" + "hello")
        _ = v11
        _ = v12
    }
    var v13 uint32 = 480
    _ = v10
    _ = v13
    return 80
}
func genFn1() bool {
    var v14 int16 = 102
    v14 = 516
    var v15 string = ("\n\t" + "foo bar")
    if (false || true) {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint = 65535
            v14 -= (v14 - (^(v14)))
            switch 1 {
            case 127:
                v18 := GenS2{F0: 0, F1: fi16, F2: -1}
                _ = v18
            case -753:
                v19 := GenS2{F0: 65535, F1: (fi16 >> (((v17 % (((v17 - v17)) | 1))) & 63)), F2: -2147483648}
                _ = v19
            default:
                var v20 int16 = (v14 >> ((1) & 63))
                _ = v20
            }
            _ = fi16
            _ = v17
        }
        fmt.Println(v14, v15)
    }
    _ = v14
    _ = v15
    return true
}

func main() {
    var v1 uint8 = 42
    var v2 int = 42
    v3 := []uint32{2147483647}
    if true {
        v1 = v1
    }
    switch "foo bar" {
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = 65535
        }
    case "a":
        switch v2 {
        case 127:
            for fi21 := 0; fi21 < 14; fi21++ {
                _ = len(v3)
                _ = fi21
            }
        case -986:
            fmt.Println(v1, v1)
        case -960:
            v22 := GenS1{F0: 32767, F1: 1}
            _ = v22
        default:
            v3 = append(v3, 65535)
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v23 uint8 = (v1 << ((65535) & 63))
    v24 := GenS2{F0: 2, F1: -112, F2: -492}
    v1 += v1
    fmt.Println(v23)
    for fi25 := 0; fi25 < 7; fi25++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v26 int8 = 0
        v27 := GenS1{F0: -283, F1: 42}
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 int64 = -1
    var v29 uint8 = 55
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
}
