package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint8; F1 int32; F2 int16 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint32; F1 bool; F2 int }

func main() {
    var v1 uint16 = 65535
    var v2 int64 = 42
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    v5 := GenS1{F0: 2147483647, F1: v1, F2: v2}
    var v6 uint8 = 114
    switch 9223372036854775807 {
    case -735:
        var v7 uint16 = 6
        _ = v7
    case -311:
        fmt.Println(v6)
    }
    v8 := []int{128, -637, 0, 9223372036854775807}
    v2 += v2
    var v9 int8 = -128
    if (!true) {
        v10 := GenS1{F0: 4294967295, F1: 0, F2: v2}
        var v11 bool = (((("\n\t" + "\n\t") + ((("a" + ("" + ("foo bar" + "foo bar"))) + "hello") + "foo bar")) + "") == ("" + "\n\t"))
        _ = v10
        _ = v11
    }
    if len(v8) > 0 {
        v8[0] = -493
    }
    v12 := GenS2{F0: (v6 & 240), F1: 2147483647, F2: 0}
    _ = len(v3)
    fmt.Println(v1, v1)
    fmt.Println(v1, v9)
    if (true || (1e10 >= (1.5 + float32(1e10)))) {
        v6 = (v6 >> ((4294967295) & 63))
    } else {
        for fi13 := 0; fi13 < 15; fi13++ {
            switch (fi13 & fi13) {
            case 236:
                v2 += 0
            default:
                var v14 uint32 = 893
                _ = v14
            }
            v6 -= (v6 % (((^(v6))) | 1))
            switch ("hello" + "a") {
            case "hello":
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            case "a":
                var v15 int8 = v9
                _ = v15
            }
            _ = fi13
        }
    }
    var v16 uint16 = v1
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v1 = 255
    if (v9 >= v9) {
        v17 := GenS3{F0: 0, F1: -9223372036854775808}
        _ = v17
    } else {
        var v18 int8 = v9
        var v19 uint8 = (^(v6))
        v20 := GenS4{F0: 2147483647, F1: true, F2: -2147483648}
        _ = v18
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
