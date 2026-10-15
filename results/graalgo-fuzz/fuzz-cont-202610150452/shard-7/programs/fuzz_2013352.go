package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint8, p2 uint64) uint16 {
    if (!(!true)) {
        p1 -= (^(p0))
        p1 += p1
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(p0, p0)
        var v6 int32 = 42
        fmt.Println(p1)
        _ = fi5
        _ = v6
    }
    switch ("" + ("" + "hello")) {
    case "":
        v7 := GenS1{F0: 42}
        _ = v7
    case "hello":
        var v8 uint64 = p2
        _ = v8
    case "foo bar":
        v9 := []int{1, 32768, -1, 127}
        _ = v9
    default:
        v10 := GenS2{F0: 1}
        _ = v10
    }
    if (0 < 0) {
        for fi11 := 0; fi11 < 16; fi11++ {
            fmt.Println(fi11, fi11)
            _ = fi11
        }
    }
    var v12 uint8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 168
}

func main() {
    var v1 string = ("foo bar" + ("" + (("\n\t" + (("\n\t" + "a") + "")) + "\n\t")))
    var v2 float32 = (((1e10 - 0.1) + ((0.0 * 0.1) + 0.0)) * -2.25)
    v3 := make(map[string]uint32)
    var v4 uint16 = 363
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS3{F0: -32768}
        _ = fi13
        _ = v14
    }
    if true {
        fmt.Println(v4, v4)
        var v15 int64 = -923
        _ = v15
    } else {
        v16 := []int{}
        _ = v16
    }
    v1 = ((("" + "hello") + v1) + v1)
    var v17 int = -308
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 string = ("\n\t" + v1)
        v2 += (1.5 + 1.5)
        _ = fi18
        _ = v19
    }
    var v20 float64 = float64(v2)
    if (v4 != (v4 % (((v4 * v4)) | 1))) {
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 bool = true
            _ = fi21
            _ = v22
        }
        fmt.Println(v1, v2, v1)
        switch "foo bar" {
        case "":
            var v23 string = v1
            _ = v23
        case "\n\t":
            fmt.Println(v20, v20)
        }
    }
    v24 := GenS2{F0: 0}
    for fi25 := 0; fi25 < 18; fi25++ {
        fmt.Println(v1)
        _ = fi25
    }
    v2 += (((float32((((float64(-2.25) / (-2.25 - 1e10)) + (v20 / v20)) - 1e10)) * v2) / v2) + (1e10 + (v2 - v2)))
    if (v17 == v17) {
        var v26 bool = (35 < -9223372036854775808)
        for fi27 := 0; fi27 < 12; fi27++ {
            fmt.Println(v4, v26, fi27)
            _ = fi27
        }
        _ = len(v3)
        _ = v26
    }
    if false {
        var v28 float32 = (0.0 / ((0.0 + 1.5) / v2))
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
