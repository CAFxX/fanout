package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
func genFn0() int {
    v5 := []int{-1, -1, 32768, 0}
    var v6 int32 = 0
    var v7 uint = 586
    _ = v5
    _ = v6
    _ = v7
    return 59
}
func genFn1(p0 int8) int {
    var v8 int64 = 1
    switch "a" {
    case "\n\t":
        fmt.Println(p0, v8)
    case "foo bar":
        fmt.Println(p0, v8)
    case "":
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int8 = (p0 * -1)
            fmt.Println(p0)
            var v11 uint = 65535
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 uint8 = 255
    v12 = v12
    _ = p0
    _ = v12
    _ = v8
    return 2147483647
}

func main() {
    var v1 string = (("" + "hello") + "foo bar")
    var v2 uint16 = 32767
    var v3 uint8 = 1
    var v4 bool = (833 < -32768)
    fmt.Println(v1, v1, v4)
    if v4 {
        if v4 {
            v4 = (((v2 & v2) != 0) && ((0 <= 283) && (833 == 42)))
            fmt.Println(v3, v2, v4)
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := []int{}
            _ = fi13
            _ = v14
        }
        var v15 uint16 = v2
        _ = v15
    }
    v16 := []int{-2147483648, 291}
    if (false && v4) {
        fmt.Println(v1)
    } else {
        var v17 string = v1
        for fi18 := 0; fi18 < 1; fi18++ {
            v2 = (v2 - 42)
            v17 = "foo bar"
            _ = fi18
        }
        _ = v17
    }
    v3 = v3
    fmt.Println(v4, v2, v3)
    fmt.Println(v2, v1, v4)
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 float32 = 1e10
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        _ = fi19
        _ = v20
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    var v21 float64 = 1e10
    switch "" {
    case "hello":
        var v22 uint = 18446744073709551615
        _ = v22
    case "foo bar":
        switch "a" {
        case "\n\t":
            var v23 int32 = 2147483647
            _ = v23
        case "hello":
            var v24 int = genFn0()
            _ = v24
        case "a":
            fmt.Println(v4, v1, v2)
        default:
            var v25 int16 = 0
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        fmt.Println(v2)
        v27 := GenS1{F0: -609}
        _ = fi26
        _ = v27
    }
    v28 := GenS2{F0: 100}
    var v29 float64 = v21
    v21 = (v29 / 1.5)
    var v30 int64 = 920
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
