package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0(p0 int8) int16 {
    var v5 int8 = p0
    var v6 uint8 = 2
    fmt.Println(v5, v6, v6)
    var v7 string = ""
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1() uint64 {
    switch "" {
    case "foo bar":
        var v8 string = "hello"
        _ = v8
    case "":
        v9 := []int{-149}
        _ = v9
    case "a":
        var v10 uint32 = 2147483648
        _ = v10
    }
    var v11 int16 = 42
    var v12 int16 = 201
    _ = v11
    _ = v12
    return 133
}

func main() {
    var v1 uint16 = 32767
    var v2 string = "hello"
    var v3 float64 = ((((float64(-2.25) * 0.1) / 3.14159) * 0.1) - 0.0)
    var v4 uint64 = 1
    fmt.Println(v2, v3)
    var v13 uint64 = (v4 ^ 479)
    switch "" {
    case "\n\t":
        var v14 int32 = -1
        _ = v14
    case "":
        var v15 uint8 = 42
        _ = v15
    case "a":
        v3 = 1.5
    default:
        var v16 bool = true
        _ = v16
    }
    v1 = (v1 >> ((728) & 63))
    if true {
        v13 = v13
        v13 = v4
        v3 -= float64((3.14159 / -2.25))
    }
    var v17 float64 = v3
    for fi18 := 0; fi18 < 20; fi18++ {
        v4 = (v4 & (v13 | v13))
        for fi19 := 0; fi19 < 4; fi19++ {
            switch (-(fi18)) {
            case -1:
                var v20 bool = true
                _ = v20
            case 127:
                switch (v2 + (v2 + "a")) {
                case "foo bar":
                    v21 := []int{(fi18 >> ((699) & 63)), 763, 227, (fi19 + (fi18 << ((0) & 63)))}
                    _ = v21
                }
            case -9223372036854775808:
                var v22 bool = (!((v1 / (((v1 * v1)) | 1)) < 2))
                _ = v22
            default:
                v3 = (-2.25 - float64(float32(v17)))
            }
            _ = fi19
        }
        v13 -= 0
        _ = fi18
    }
    v23 := GenS1{F0: 0, F1: -124}
    v1 = (v1 << ((255) & 63))
    fmt.Println(v2, v17, v4)
    v24 := []int{200}
    v4 = 2
    switch v2 {
    case "foo bar":
        var v25 string = v2
        _ = v25
    case "a":
        var v26 uint64 = v13
        _ = v26
    case "\n\t":
        fmt.Println(v1)
    default:
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
    }
    var v27 float32 = (0.0 + 0.1)
    for fi28 := 0; fi28 < 17; fi28++ {
        fmt.Println(v2)
        _ = fi28
    }
    var v29 int32 = 2147483647
    if false {
        var v30 bool = true
        var v31 int8 = 100
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v29
    _ = v3
    _ = v4
}
