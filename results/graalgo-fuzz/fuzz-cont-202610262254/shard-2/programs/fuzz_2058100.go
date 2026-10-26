package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint, p1 int16, p2 int32) uint8 {
    switch 0 {
    case 0:
        var v7 uint8 = 32
        _ = v7
    case 367:
        if (142 == uint32(p0)) {
            p0 = (p0 << (((p0 ^ p0)) & 63))
            fmt.Println(p1, p0, p0)
        }
    case 42:
        v8 := []int{-607}
        _ = v8
    }
    switch "\n\t" {
    case "\n\t":
        var v9 uint = (p0 / (((p0 % ((p0) | 1))) | 1))
        _ = v9
    case "foo bar":
        var v10 float32 = 0.1
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := []int{fi11, 672, 32767, 32768}
            _ = fi11
            _ = v12
        }
    }
    var v13 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint64 = 2
    var v3 int = -838
    var v4 int8 = 42
    var v5 uint16 = 255
    v6 := []int{v3, (-(v3)), (v3 >> ((65535) & 63)), v3}
    for fi14 := 0; fi14 < 20; fi14++ {
        if false {
            v1 = 907
            v15 := GenS1{F0: 32767, F1: 2, F2: (v4 << ((uint(v2)) & 63))}
            v4 += 39
            _ = v15
        } else {
            v2 = 2
        }
        var v16 string = (("foo bar" + ((("" + ("hello" + "foo bar")) + ((("foo bar" + "") + "\n\t") + (("foo bar" + "\n\t") + "\n\t"))) + "hello")) + "foo bar")
        _ = fi14
        _ = v16
    }
    if len(v6) > 0 {
        v6[0] = (v3 & v3)
    }
    fmt.Println(v1, v2)
    switch (v3 >> ((1) & 63)) {
    case -2147483648:
        var v17 uint = 42
        _ = v17
    case 9223372036854775807:
        var v18 string = "foo bar"
        _ = v18
    case 127:
        v4 += (v4 | v4)
    }
    switch v3 {
    case -964:
        v19 := GenS1{F0: (v3 << ((0) & 63)), F1: 870, F2: v4}
        _ = v19
    case -1:
        _ = len(v6)
    default:
        fmt.Println(v4, v2)
    }
    fmt.Println(v5, v5, v4)
    fmt.Println(v3)
    v20 := GenS2{F0: 1, F1: (v4 >> ((0) & 63))}
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := GenS3{F0: uint16(v2)}
        _ = fi21
        _ = v22
    }
    v23 := GenS4{F0: 2}
    v3 = (v3 >> ((691) & 63))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 uint8 = 255
        _ = fi24
        _ = v25
    }
    var v26 int8 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
