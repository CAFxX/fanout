package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int8 }

func main() {
    var v1 uint32 = 0
    v2 := []string{"foo bar"}
    var v3 uint8 = 42
    v4 := []uint32{v1}
    var v5 int32 = int32(v1)
    var v6 int = 839
    if true {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        v5 = 0
        v7 := GenS1{F0: -177, F1: v3, F2: 118}
        _ = v7
    }
    var v8 uint = 163
    var v9 uint8 = (v3 << ((v8) & 63))
    v8 -= 45
    fmt.Println(v3)
    for fi10 := 0; fi10 < 8; fi10++ {
        switch 128 {
        case -158:
            for fi11 := 0; fi11 < 11; fi11++ {
                fmt.Println(v6, fi10)
                _ = fi11
            }
        case -259:
            v12 := GenS1{F0: 1, F1: (v3 << (((v8 - (v8 ^ (v8 - (^(v8)))))) & 63)), F2: 42}
            _ = v12
        default:
            var v13 int = (-(v6))
            _ = v13
        }
        v14 := GenS1{F0: -1, F1: (^(v9)), F2: 127}
        _ = len(v2)
        _ = fi10
        _ = v14
    }
    var v15 float32 = (float32(0.1) - 3.14159)
    switch v6 {
    case 808:
        var v16 int8 = 0
        _ = v16
    default:
        var v17 uint32 = v1
        _ = v17
    }
    switch (("" + (("\n\t" + ((("foo bar" + ("\n\t" + "\n\t")) + "foo bar") + "")) + "hello")) + ("\n\t" + "foo bar")) {
    case "":
        v4 = append(v4, 65535)
    case "hello":
        var v18 int8 = int8(v1)
        _ = v18
    case "foo bar":
        var v19 float32 = 1.5
        _ = v19
    default:
        var v20 int64 = 42
        _ = v20
    }
    if (!(!(true || false))) {
        v21 := GenS1{F0: 284, F1: v9, F2: -128}
        if len(v4) > 0 {
            v4[0] = 65535
        }
        _ = v21
    } else {
        v9 -= v9
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v22 uint = v8
    var v23 int16 = 32767
    v23 += (v23 % ((v23) | 1))
    if len(v2) > 0 {
        v2[0] = "a"
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 float32 = float32(0.1)
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
