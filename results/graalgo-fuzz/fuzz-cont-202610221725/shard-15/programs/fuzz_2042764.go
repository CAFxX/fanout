package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    var v5 bool = ((0.1 / (3.14159 - (0.0 - 0.0))) > 0.1)
    if (381 < 2147483647) {
        var v6 uint64 = 42
        if (("a" + "\n\t") < "a") {
            switch -2147483648 {
            case -106:
                var v7 int64 = -9223372036854775808
                _ = v7
            default:
                for fi8 := 0; fi8 < 16; fi8++ {
                    var v9 string = ""
                    v5 = (((0.1 - 3.14159) >= 1e10) || (((100 >= 1) || false) || (100 >= 0)))
                    var v10 int16 = 32767
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
        } else {
            switch -2147483648 {
            case 4294967295:
                var v11 int8 = 100
                _ = v11
            default:
                var v12 uint64 = 42
                _ = v12
            }
            var v13 int8 = 42
            v14 := GenS1{F0: 1, F1: -128}
            _ = v13
            _ = v14
        }
        var v15 float32 = 1e10
        _ = v15
        _ = v6
    } else {
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            v5 = (false && true)
        case "foo bar":
            var v16 float64 = (float64(((0.0 + 0.1) / (float32(3.14159) * ((0.0 + (((3.14159 / ((3.14159 + (float32(0.0) - -2.25)) + float32((1.5 / (0.1 - (0.1 - -2.25)))))) * (float32((3.14159 / ((0.1 - 0.0) - 1.5))) * 1.5)) / float32(3.14159))) / 0.1)))) / (float64(1e10) / ((0.1 / 3.14159) / 1.0)))
            _ = v16
        case "":
            var v17 uint16 = 0
            _ = v17
        }
        v18 := GenS1{F0: 888, F1: 42}
        v5 = v5
        _ = v18
    }
    v19 := []int{626, 128, -250}
    for fi20 := 0; fi20 < 9; fi20++ {
        v5 = false
        var v21 int16 = 32767
        _ = fi20
        _ = v21
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    v5 = false
    v5 = (v5 && (((-1 <= -288) || (true || true)) || true))
    if true {
        fmt.Println(v5)
    }
    v5 = (0.1 >= float32(-2.25))
    if true {
        var v22 uint8 = 202
        _ = v22
    }
    v5 = true
    _ = len(v19)
    var v23 uint32 = 98
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    _ = len(v19)
    var v24 uint16 = 32767
    for fi25 := 0; fi25 < 6; fi25++ {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        for fi26 := 0; fi26 < 1; fi26++ {
            if v5 {
                var v27 int32 = -75
                var v28 float32 = -2.25
                var v29 int = (fi26 ^ fi26)
                var v30 float32 = (float32(1.5) + ((((v28 + (((float32(3.14159) - v28) - ((0.0 - (-2.25 - v28)) - v28)) * v28)) - v28) - v28) / v28))
                _ = v27
                _ = v28
                _ = v29
                _ = v30
            }
            fmt.Println(v23)
            _ = fi26
        }
        _ = fi25
    }
    fmt.Println(v23, v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
