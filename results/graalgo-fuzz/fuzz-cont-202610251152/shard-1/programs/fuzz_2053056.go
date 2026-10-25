package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0(p0 int) string {
    for fi5 := 0; fi5 < 4; fi5++ {
        switch p0 {
        case 128:
            fmt.Println(p0, fi5)
        case 127:
            v6 := GenS1{F0: 321, F1: 9223372036854775807}
            _ = v6
        case 0:
            p0 = 2147483647
        }
        _ = fi5
    }
    fmt.Println(p0)
    var v7 int8 = 100
    _ = p0
    _ = v7
    return "a"
}

func main() {
    var v1 float64 = 3.14159
    var v2 float64 = 1.5
    v3 := []uint32{1, 4294967295}
    v4 := []string{"foo bar", "", "\n\t"}
    var v8 float32 = float32(0.0)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v9 uint = 953
    v10 := GenS1{F0: 0, F1: -806}
    if false {
        switch "a" {
        case "hello":
            var v11 float64 = v1
            _ = v11
        }
        var v12 float32 = ((v8 * 0.1) - v8)
        _ = v12
    } else {
        v4 = append(v4, ("hello" + "a"))
        v4 = append(v4, ("hello" + "\n\t"))
    }
    v9 = 2
    if (!false) {
        for fi13 := 0; fi13 < 9; fi13++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch fi13 {
            case -9223372036854775808:
                v8 = float32(1e10)
            }
            v14 := GenS1{F0: uint64(v9), F1: 0}
            _ = fi13
            _ = v14
        }
        var v15 bool = (127 < 128)
        _ = v15
    } else {
        switch 2147483647 {
        case -38:
            var v16 int16 = 339
            _ = v16
        case 4294967295:
            if (!false) {
                var v17 int = 985
                var v18 float32 = v8
                _ = len(v3)
                _ = v17
                _ = v18
            } else {
                fmt.Println(v1)
            }
        case 127:
            v19 := GenS1{F0: 1, F1: -550}
            _ = v19
        default:
            var v20 float64 = -2.25
            _ = v20
        }
    }
    switch "\n\t" {
    case "":
        switch 32768 {
        case 127:
            v2 = v1
        }
    case "\n\t":
        v21 := GenS1{F0: 42, F1: 0}
        _ = v21
    default:
        var v22 uint8 = 255
        _ = v22
    }
    switch "\n\t" {
    case "hello":
        for fi23 := 0; fi23 < 9; fi23++ {
            v8 -= 0.1
            _ = fi23
        }
    case "\n\t":
        if false {
            fmt.Println(v1)
        }
    case "foo bar":
        v24 := GenS1{F0: 1, F1: 883}
        _ = v24
    }
    fmt.Println(v1, v8, v2)
    v2 = (float64(float32(v2)) / v1)
    switch "foo bar" {
    case "foo bar":
        var v25 int64 = 578
        _ = v25
    case "a":
        switch 714 {
        case 1:
            var v26 int64 = 822
            _ = v26
        case -6:
            var v27 int16 = 0
            _ = v27
        case 2147483647:
            _ = len(v4)
        }
    default:
        for fi28 := 0; fi28 < 16; fi28++ {
            fmt.Println(v9, v9)
            var v29 int = fi28
            v1 -= v2
            _ = fi28
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
