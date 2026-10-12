package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int32; F2 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{556}
    var v3 int16 = 42
    v4 := []int{1, 127, 2147483647}
    switch ("hello" + "foo bar") {
    case "a":
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 int = fi5
            var v7 int = (fi5 * (fi5 + (v6 % ((32768) | 1))))
            _ = fi5
            _ = v6
            _ = v7
        }
    case "":
        for fi8 := 0; fi8 < 14; fi8++ {
            v3 -= v3
            _ = fi8
        }
    }
    v9 := GenS1{F0: 42}
    var v10 float64 = float64(1.5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (0 < 127) {
        var v11 uint16 = 65535
        for fi12 := 0; fi12 < 8; fi12++ {
            if true {
                v13 := GenS1{F0: fi12}
                v3 = v3
                _ = v13
            } else {
                fmt.Println(fi12)
                var v14 uint32 = 0
                var v15 string = "hello"
                _ = v14
                _ = v15
            }
            switch "" {
            case "hello":
                _ = len(v1)
            case "foo bar":
                var v16 int64 = 9223372036854775807
                _ = v16
            case "":
                var v17 uint = 42
                _ = v17
            default:
                for fi18 := 0; fi18 < 17; fi18++ {
                    v19 := GenS2{F0: 2, F1: 42, F2: 711}
                    _ = fi18
                    _ = v19
                }
            }
            _ = fi12
        }
        _ = v11
    } else {
        fmt.Println(v3)
        switch -834 {
        case 42:
            v20 := GenS2{F0: 127, F1: 2147483647, F2: 0}
            _ = v20
        case 128:
            if (!false) {
                v10 += float64(float32(((v10 / v10) + (0.0 / 3.14159))))
                switch "" {
                case "hello":
                    var v21 uint8 = 2
                    _ = v21
                case "a":
                    var v22 string = "hello"
                    _ = v22
                }
            }
        }
    }
    if true {
        var v23 int32 = 384
        _ = v23
    } else {
        fmt.Println(v10, v10)
        var v24 int16 = -1
        fmt.Println(v3, v3)
        _ = v24
    }
    v25 := GenS2{F0: 94, F1: -2147483648, F2: 780}
    switch 0 {
    case 4294967295:
        v3 += 788
    case 604:
        var v26 string = "\n\t"
        _ = v26
    case -2147483648:
        fmt.Println(v3, v3)
    }
    v10 -= ((((v10 / (v10 * v10)) + v10) - ((v10 / v10) * v10)) * (v10 * float64(0.1)))
    v10 = v10
    var v27 uint64 = 4294967295
    var v28 int16 = (^(v3))
    switch 9223372036854775807 {
    case 127:
        v3 = v28
    case 32767:
        var v29 int = -531
        _ = v29
    }
    v3 = (-(v3))
    v27 += 18446744073709551615
    for fi30 := 0; fi30 < 13; fi30++ {
        v31 := GenS1{F0: fi30}
        var v32 uint64 = v27
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v9
}
