package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint; F1 int; F2 int8 }

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{805, 65535, 42, 770}
    var v3 uint8 = 127
    v4 := make(map[uint32]bool)
    var v5 int64 = -1
    v5 -= (v5 >> ((65535) & 63))
    var v6 int32 = 908
    switch "\n\t" {
    case "":
        fmt.Println(v1, v5)
    case "foo bar":
        var v7 uint16 = 0
        _ = v7
    }
    fmt.Println(v3)
    v8 := GenS1{F0: 151, F1: 128}
    if len(v2) > 0 {
        v2[0] = 0
    }
    if (!false) {
        var v9 int8 = -128
        _ = v9
    } else {
        v10 := GenS1{F0: 42, F1: 127}
        v2 = append(v2, 665)
        if true {
            v11 := GenS1{F0: 2, F1: 2147483647}
            v12 := GenS1{F0: 42, F1: 42}
            switch "\n\t" {
            case "hello":
                v1 = 0.0
            case "a":
                var v13 int16 = 42
                _ = v13
            default:
                if (0 <= 31) {
                    _ = len(v2)
                    switch "a" {
                    case "foo bar":
                        v1 = ((((1e10 - v1) + v1) - v1) / v1)
                    case "a":
                        _ = len(v2)
                    case "hello":
                        v14 := GenS2{F0: (-(v5)), F1: -1}
                        _ = v14
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    var v15 float32 = 0.1
                    _ = v15
                } else {
                    delete(v4, 2)
                }
            }
            _ = v11
            _ = v12
        } else {
            fmt.Println(v3, v6, v5)
            var v16 bool = (v6 <= v6)
            _ = v16
        }
        _ = v10
    }
    if true {
        var v17 string = ("a" + "a")
        _ = v17
    } else {
        v1 = v1
        if false {
            v3 += v3
            var v18 string = "a"
            _ = v18
        }
    }
    switch ("" + "") {
    case "hello":
        var v19 uint8 = 255
        _ = v19
    case "":
        var v20 uint = uint(v3)
        _ = v20
    case "foo bar":
        if (v1 <= ((0.0 * v1) - v1)) {
            v4[2] = (-2147483648 >= 168)
        }
    default:
        for fi21 := 0; fi21 < 20; fi21++ {
            if (v1 > -2.25) {
                v6 = 42
                var v22 uint16 = 129
                v6 = -1
                _ = v22
            } else {
                var v23 string = (("hello" + "a") + "\n\t")
                for fi24 := 0; fi24 < 15; fi24++ {
                    switch (fi24 + (-(fi24))) {
                    case -181:
                        v25 := GenS3{F0: 916, F1: fi21, F2: int8(v3)}
                        _ = v25
                    case -9223372036854775808:
                        fmt.Println(v23, v6)
                    default:
                        if false {
                            v26 := v4[509]
                            _ = v26
                        }
                    }
                    _ = fi24
                }
                _ = v23
            }
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            if (true || false) {
                v5 -= (v5 % (((v5 >> ((402) & 63))) | 1))
                _ = len(v2)
            }
            _ = fi21
        }
    }
    var v27 uint16 = uint16(v3)
    v4[uint32(v27)] = false
    if len(v2) > 0 {
        v2[0] = 65535
    }
    v5 -= v5
    fmt.Println(v3, v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
