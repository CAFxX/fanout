package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int64 }
type GenS2 struct { F0 string; F1 int32; F2 int16 }

func main() {
    var v1 int64 = -214
    var v2 string = "a"
    v3 := []int{32768}
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]string)
    var v6 int32 = 1
    v1 += v1
    var v7 bool = false
    v8 := GenS1{F0: 42, F1: ((((("" + v2) + v2) + ("foo bar" + (v2 + ""))) + v2) + (v2 + "foo bar")), F2: 1}
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v2)
        _ = fi9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS1{F0: v6, F1: v2, F2: (^(v1))}
            var v13 int8 = 100
            v3 = append(v3, fi10)
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    var v14 string = v2
    v7 = v7
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 uint = 4294967295
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        if false {
            v18 := v4[65535]
            _ = v18
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        switch v2 {
        case "foo bar":
            fmt.Println(v2, v6)
        case "hello":
            v20 := GenS2{F0: ("" + v14), F1: v6, F2: -32768}
            _ = v20
        default:
            v1 = 1
        }
        fmt.Println(v14, v14)
        v6 = (v6 % (((v6 >> ((0) & 63))) | 1))
        _ = fi19
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        v7 = (((3.14159 / 3.14159) >= 0.0) || v7)
        for fi22 := 0; fi22 < 19; fi22++ {
            switch "foo bar" {
            case "hello":
                v3 = append(v3, 1)
            case "foo bar":
                fmt.Println(v14, v7)
            }
            _ = fi22
        }
        _ = fi21
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        if (!true) {
            var v24 float32 = 1.5
            fmt.Println(v2, v7, v7)
            switch fi23 {
            case 0:
                delete(v4, 2147483648)
            }
            _ = v24
        }
        _ = fi23
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 string = (((("hello" + ("\n\t" + (v2 + ((v2 + v14) + v14)))) + (v14 + "hello")) + v2) + "hello")
        switch fi25 {
        case 4294967295:
            fmt.Println(v26)
        case -2147483648:
            fmt.Println(v26)
        case -41:
            v7 = (!true)
        }
        v27 := GenS2{F0: "\n\t", F1: v6, F2: 1000}
        _ = fi25
        _ = v26
        _ = v27
    }
    if (((v14 + "foo bar") != v14) && ((0 == -1) && v7)) {
        var v28 float64 = (0.0 - (float64(-2.25) - 0.0))
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
