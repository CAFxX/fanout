package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint32; F1 bool }
func genFn0(p0 int64, p1 int32, p2 bool) int32 {
    fmt.Println(p2, p1)
    p1 = p1
    v5 := []int{}
    p2 = (p2 || p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}

func main() {
    var v1 uint8 = 255
    var v2 int64 = 1
    var v3 bool = false
    v4 := make(map[string]uint32)
    for fi6 := 0; fi6 < 16; fi6++ {
        v3 = (!false)
        fmt.Println(v1)
        _ = fi6
    }
    v1 = 127
    for fi7 := 0; fi7 < 4; fi7++ {
        v2 = v2
        var v8 int16 = 999
        _ = fi7
        _ = v8
    }
    var v9 int64 = v2
    var v10 bool = false
    v3 = (0 > -128)
    switch "hello" {
    case "":
        v1 -= (^(v1))
    case "\n\t":
        var v11 bool = ((true && (-2.25 < (-2.25 + float32(1e10)))) && (1 > 42))
        _ = v11
    }
    var v12 int64 = v9
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := GenS1{F0: 1}
        fmt.Println(v1, v12, fi13)
        var v15 float64 = 1e10
        _ = fi13
        _ = v14
        _ = v15
    }
    if v3 {
        for fi16 := 0; fi16 < 15; fi16++ {
            v9 -= (v9 * -113)
            _ = fi16
        }
    } else {
        switch "\n\t" {
        case "foo bar":
            v3 = v3
        default:
            v9 += (v2 + v9)
        }
        var v17 uint64 = 4294967295
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 float32 = 0.0
        var v20 uint = 2
        _ = fi18
        _ = v19
        _ = v20
    }
    switch -2147483648 {
    case 1:
        for fi21 := 0; fi21 < 20; fi21++ {
            v12 = v2
            fmt.Println(v1)
            _ = fi21
        }
    case 301:
        v22 := GenS2{F0: (("" + "hello") + ("" + ((("a" + (("hello" + "\n\t") + "")) + "hello") + ("\n\t" + "hello")))), F1: 951}
        _ = v22
    }
    var v23 int64 = v2
    var v24 int16 = 904
    if (!(!false)) {
        v25 := GenS1{F0: v24}
        fmt.Println(v24)
        _ = v25
    }
    v26 := GenS3{F0: 2, F1: false}
    v2 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v9
}
