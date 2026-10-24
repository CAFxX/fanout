package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint16 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint16; F2 int16 }
type GenS5 struct { F0 int8; F1 int32; F2 uint16 }
func genFn0(p0 bool, p1 int32) int64 {
    if (p0 || (-828 >= -298)) {
        switch "hello" {
        case "":
            v5 := GenS1{F0: p0, F1: 32767, F2: 42}
            _ = v5
        case "foo bar":
            if false {
                v6 := GenS1{F0: false, F1: 0, F2: 0}
                v7 := []int{}
                for fi8 := 0; fi8 < 17; fi8++ {
                    _ = len(v7)
                    _ = fi8
                }
                _ = v6
                _ = v7
            } else {
                var v9 string = (((("\n\t" + ("" + "\n\t")) + "\n\t") + "") + "\n\t")
                _ = v9
            }
        }
        if ((624 >= 1) || ((3.14159 <= 3.14159) || p0)) {
            var v10 bool = p0
            var v11 uint8 = 231
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        p0 = (42 >= 65535)
        v13 := GenS1{F0: p0, F1: 0, F2: 255}
        fmt.Println(fi12)
        _ = fi12
        _ = v13
    }
    p1 = p1
    if ("foo bar" > "a") {
        var v14 int32 = (p1 * (p1 | (p1 << ((65535) & 63))))
        p0 = (0 != 0)
        var v15 uint = 255
        _ = v14
        _ = v15
    } else {
        p0 = p0
    }
    var v16 bool = (((p1 < (p1 - p1)) || (-1 <= 578)) || true)
    _ = p0
    _ = p1
    _ = v16
    return 1
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint8 = 255
    var v3 int64 = 1
    var v4 int8 = -1
    v3 += (-(v3))
    v1 = (^(v1))
    v4 += v4
    v2 = 128
    v17 := []int{}
    v18 := GenS2{F0: -89, F1: false}
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    v19 := GenS3{F0: 18446744073709551615}
    if len(v17) > 0 {
        v17[0] = 702
    }
    switch "" {
    case "\n\t":
        v3 = 42
    default:
        fmt.Println(v1)
    }
    switch int(v4) {
    case -9223372036854775808:
        v20 := GenS4{F0: 42, F1: 453, F2: -32768}
        _ = v20
    case 0:
        v3 = genFn0((!(!false)), -1)
    case -2147483648:
        if true {
            var v21 int = 42
            _ = v21
        }
    }
    if true {
        v22 := GenS5{F0: 42, F1: -443, F2: 259}
        var v23 uint32 = (^(v1))
        _ = v22
        _ = v23
    }
    _ = len(v17)
    var v24 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
