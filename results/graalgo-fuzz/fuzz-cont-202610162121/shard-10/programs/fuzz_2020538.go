package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint8; F1 int32; F2 int32 }
func genFn0(p0 bool, p1 uint32, p2 int) uint8 {
    fmt.Println(p0, p0)
    v7 := GenS1{F0: 0, F1: 0}
    v8 := GenS2{F0: 18446744073709551615, F1: p2}
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint8 = 72
        fmt.Println(p1)
        _ = fi9
        _ = v10
    }
    if (-207 <= 1) {
        fmt.Println(p0, p0)
    } else {
        var v11 uint16 = 512
        var v12 int = (p2 / (((p2 % ((p2) | 1))) | 1))
        p2 -= (p2 + (v12 | 42))
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 128
}
func genFn1() bool {
    v13 := []int{}
    var v14 uint16 = 255
    if ("hello" < ("hello" + "")) {
        switch 1 {
        case 315:
            _ = len(v13)
        case 128:
            var v15 uint = 20
            _ = v15
        }
        fmt.Println(v14)
    } else {
        var v16 string = ""
        _ = v16
    }
    var v17 float32 = -2.25
    _ = v13
    _ = v14
    _ = v17
    return (!true)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{1}
    v3 := []int{127, 0}
    var v4 uint8 = 179
    var v5 bool = true
    var v6 int32 = -502
    v18 := GenS1{F0: 9223372036854775807, F1: 18446744073709551615}
    v19 := v1[2]
    fmt.Println(v4)
    v6 -= (v6 | (-(v6)))
    v20 := GenS1{F0: -9223372036854775808, F1: 65535}
    v21 := GenS2{F0: 2, F1: 435}
    var v22 uint = 4294967295
    var v23 float32 = 0.0
    fmt.Println(v23, v19, v23)
    var v24 bool = v5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        v26 := GenS3{F0: 2, F1: (v6 | (v6 % ((2147483647) | 1))), F2: 2147483647}
        _ = fi25
        _ = v26
    }
    switch v19 {
    case "foo bar":
        _ = len(v2)
    case "a":
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 string = "a"
            fmt.Println(v23, v5)
            _ = fi27
            _ = v28
        }
    default:
        v29 := GenS3{F0: (v4 << (((v22 - 4294967295)) & 63)), F1: (^(v6)), F2: int32(v22)}
        _ = v29
    }
    var v30 int = -173
    var v31 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
