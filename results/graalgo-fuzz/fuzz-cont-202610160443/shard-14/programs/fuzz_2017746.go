package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int8; F2 string }

func main() {
    v1 := []uint32{2147483648, 0, 4294967295}
    var v2 float64 = 3.14159
    var v3 float64 = v2
    var v4 uint = 42
    v5 := GenS1{F0: 1, F1: "", F2: 0}
    var v6 uint8 = 0
    v7 := GenS2{F0: 2147483648}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        if true {
            v8 := GenS1{F0: 1, F1: "hello", F2: -443}
            _ = v8
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            v6 = (v6 % ((v6) | 1))
            if (!true) {
                var v10 bool = (true || true)
                v6 = 2
                if len(v1) > 0 {
                    v1[0] = 1
                }
                _ = v10
            }
            v11 := GenS2{F0: 236}
            _ = fi9
            _ = v11
        }
        v6 -= 116
    }
    v12 := GenS1{F0: 1, F1: "foo bar", F2: -9223372036854775808}
    var v13 int32 = 185
    v14 := GenS3{F0: 966, F1: -128, F2: ("foo bar" + "")}
    var v15 bool = false
    if v15 {
        switch -286 {
        case 679:
            v1 = append(v1, 42)
        }
    } else {
        switch 811 {
        case -259:
            var v16 uint8 = v6
            _ = v16
        case 42:
            for fi17 := 0; fi17 < 19; fi17++ {
                v4 -= v4
                var v18 float32 = 0.1
                switch fi17 {
                case -144:
                    var v19 uint16 = 2
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
        case 442:
            v1 = append(v1, uint32(v6))
        }
        v4 += 749
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        if (!(((v6 << ((526) & 63)) <= 128) && v15)) {
            fmt.Println(v15, v15)
            v13 = 2147483647
        }
        if len(v1) > 0 {
            v1[0] = 1
        }
        var v21 uint32 = 65535
        _ = fi20
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
