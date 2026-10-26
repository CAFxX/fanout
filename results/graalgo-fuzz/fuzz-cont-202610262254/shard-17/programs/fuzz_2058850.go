package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -128
    v3 := []uint32{0, 2147483648, 65535}
    var v4 int32 = 347
    var v5 int64 = 9223372036854775807
    var v6 int64 = (v5 | (-(v5)))
    var v7 int64 = (v6 - 0)
    v8 := GenS1{F0: 715, F1: 1, F2: "\n\t"}
    var v9 int8 = (v2 / (((v2 ^ 1)) | 1))
    var v10 int8 = (^(v9))
    if false {
        switch "foo bar" {
        case "hello":
            v2 += 34
        case "foo bar":
            v11 := GenS1{F0: -1, F1: -716, F2: "\n\t"}
            _ = v11
        case "a":
            var v12 int64 = v6
            _ = v12
        default:
            v13 := v1[2147483648]
            _ = v13
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v14 bool = (42 == 128)
    v15 := GenS1{F0: 1000, F1: 1000, F2: "a"}
    switch "a" {
    case "":
        if v14 {
            fmt.Println(v7)
            var v16 string = "a"
            var v17 float64 = 1e10
            _ = v16
            _ = v17
        }
    case "foo bar":
        v5 -= (v7 ^ -9223372036854775808)
    default:
        v10 = v9
    }
    if (((v9 - v10) == v2) && true) {
        if ((!((1e10 * 0.0) <= 3.14159)) || true) {
            v5 += v7
            v18 := GenS1{F0: 1000, F1: 0, F2: ("hello" + "foo bar")}
            _ = v18
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            v14 = (-1 <= v5)
            _ = fi19
        }
        v4 += 1
    }
    var v20 uint = 255
    if v14 {
        fmt.Println(v2, v14)
        var v21 bool = v14
        v22 := GenS1{F0: 32767, F1: -550, F2: ("foo bar" + ("" + ("hello" + "")))}
        _ = v21
        _ = v22
    } else {
        var v23 string = "a"
        _ = v23
    }
    fmt.Println(v20, v5, v9)
    var v24 uint32 = 1
    v25 := GenS1{F0: -536, F1: 42, F2: "a"}
    for fi26 := 0; fi26 < 10; fi26++ {
        if false {
            v10 = v9
            fmt.Println(v6, v6, v6)
        } else {
            var v27 int64 = 42
            var v28 bool = false
            v20 += (v20 << ((v20) & 63))
            _ = v27
            _ = v28
        }
        _ = fi26
    }
    var v29 int64 = v6
    switch "\n\t" {
    case "foo bar":
        var v30 int64 = (v5 >> ((344) & 63))
        _ = v30
    case "\n\t":
        if (((0.1 + (float64(1e10) * -2.25)) + 1e10) < 0.1) {
            v4 = (v4 * (v4 & (-(v4))))
        } else {
            for fi31 := 0; fi31 < 15; fi31++ {
                _ = len(v1)
                _ = fi31
            }
            v5 = 0
            v29 += v29
        }
    case "a":
        if v14 {
            v32 := GenS1{F0: 0, F1: 42, F2: "\n\t"}
            _ = v32
        }
    default:
        v10 += v10
    }
    v2 = (v10 % (((v10 >> ((v20) & 63))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
