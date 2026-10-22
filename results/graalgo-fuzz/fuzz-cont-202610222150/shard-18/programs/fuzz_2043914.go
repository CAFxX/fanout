package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 string }
type GenS2 struct { F0 int; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int }
type GenS4 struct { F0 uint8; F1 string; F2 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 255
    v3 := make(map[string]bool)
    var v4 bool = true
    v4 = v4
    v5 := GenS1{F0: 4294967295, F1: 2147483648, F2: "foo bar"}
    fmt.Println(v2)
    var v6 int8 = 0
    v7 := []int{575, -995, 127, 128}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 uint64 = 4294967295
        switch "hello" {
        case "hello":
            v10 := GenS2{F0: fi8, F1: (v9 - v9), F2: uint32(v9)}
            _ = v10
        default:
            var v11 string = "hello"
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    if false {
        var v12 int8 = v6
        if v4 {
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 bool = v4
                for fi15 := 0; fi15 < 6; fi15++ {
                    var v16 uint8 = 1
                    var v17 int = (fi13 << ((uint(v16)) & 63))
                    v18 := GenS3{F0: 32767, F1: 65535, F2: fi15}
                    var v19 int16 = -355
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                var v20 int32 = 42
                _ = fi13
                _ = v14
                _ = v20
            }
            v21 := GenS4{F0: 2, F1: "hello", F2: 42}
            v6 = v12
            _ = v21
        }
        switch "a" {
        case "foo bar":
            fmt.Println(v2, v6)
        case "a":
            switch 32767 {
            case 32767:
                v2 -= 1
            case -9223372036854775808:
                if (9223372036854775807 < 42) {
                    v22 := GenS4{F0: 196, F1: (("\n\t" + "") + "foo bar"), F2: 838}
                    v23 := GenS3{F0: 941, F1: 982, F2: int(v6)}
                    _ = v22
                    _ = v23
                } else {
                    if v4 {
                        delete(v3, "a")
                        v6 = -1
                        v4 = (!((v6 ^ v6) <= (-(v12))))
                    }
                    v12 += (v6 * 0)
                }
            default:
                v1[4294967295] = 4294967295
            }
        }
        _ = v12
    } else {
        v4 = (1e10 <= 3.14159)
        if len(v7) > 0 {
            v7[0] = -1
        }
        var v24 string = "\n\t"
        _ = v24
    }
    if (!(0 < 9223372036854775807)) {
        fmt.Println(v6, v2)
    }
    if true {
        v6 = v6
        var v25 int8 = (v6 - (v6 + (v6 << (((v2 + (v2 >> (((v2 + (v2 * (v2 & 18446744073709551615)))) & 63)))) & 63))))
        _ = v25
    } else {
        var v26 int32 = 1
        _ = v26
    }
    v27 := GenS2{F0: -258, F1: 42, F2: 1}
    v28 := GenS3{F0: 2, F1: 718, F2: -571}
    var v29 float32 = ((0.0 / 1e10) + (1e10 + (1e10 / 3.14159)))
    for fi30 := 0; fi30 < 2; fi30++ {
        v31 := GenS1{F0: 0, F1: 777, F2: ("foo bar" + "")}
        _ = fi30
        _ = v31
    }
    var v32 bool = (104 != v6)
    var v33 uint64 = 2
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
