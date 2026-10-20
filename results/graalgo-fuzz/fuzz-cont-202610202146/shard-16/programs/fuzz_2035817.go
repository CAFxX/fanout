package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint64; F2 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 0
    var v3 float64 = 3.14159
    var v4 float32 = (-2.25 - 0.0)
    v5 := GenS1{F0: -508, F1: 42}
    switch 9223372036854775807 {
    case -2147483648:
        var v6 uint32 = 1
        _ = v6
    }
    v7 := []int{}
    v2 -= (v2 * v2)
    var v8 int16 = 1000
    var v9 float32 = v4
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    switch "\n\t" {
    case "a":
        v10 := v1[-2147483648]
        _ = v10
    case "hello":
        switch "\n\t" {
        case "hello":
            switch ("" + ("hello" + "a")) {
            case "\n\t":
                if true {
                    var v11 int8 = 29
                    _ = v11
                }
            case "":
                for fi12 := 0; fi12 < 6; fi12++ {
                    if (3.14159 < 1.5) {
                        if ((v2 ^ v2) < 9223372036854775807) {
                            v9 = v4
                            v13 := GenS1{F0: 0, F1: (v8 + v8)}
                            var v14 int64 = -1
                            _ = v13
                            _ = v14
                        }
                        var v15 int8 = 0
                        for fi16 := 0; fi16 < 3; fi16++ {
                            v1[32768] = (!true)
                            v4 = (v9 - float32(v3))
                            v15 = v15
                            v17 := GenS1{F0: -1, F1: (-(v8))}
                            _ = fi16
                            _ = v17
                        }
                        _ = v15
                    }
                    _ = fi12
                }
            case "foo bar":
                v4 += (0.0 - 0.1)
            default:
                v3 += float64(v9)
            }
        }
    case "":
        var v18 bool = ("hello" < "")
        _ = v18
    default:
        v19 := GenS2{F0: true}
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        if (!(!false)) {
            var v21 uint = 233
            fmt.Println(v8)
            _ = v21
        }
        var v22 float64 = v3
        switch "" {
        case "\n\t":
            fmt.Println(v8, v8, v22)
        case "hello":
            v23 := GenS1{F0: -2147483648, F1: 32767}
            _ = v23
        case "foo bar":
            v8 = int16(v2)
        }
        _ = fi20
        _ = v22
    }
    v4 = 3.14159
    var v24 int8 = -1
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 uint16 = 255
        var v27 string = "a"
        _ = fi25
        _ = v26
        _ = v27
    }
    if false {
        if (184 == 65535) {
            var v28 int16 = v8
            v7 = append(v7, 128)
            _ = v28
        }
        fmt.Println(v4)
    } else {
        var v29 int8 = -1
        _ = v29
    }
    fmt.Println(v24, v8)
    v30 := GenS2{F0: true}
    for fi31 := 0; fi31 < 14; fi31++ {
        v32 := GenS3{F0: 2147483647, F1: 18446744073709551615, F2: 2}
        v4 += (v9 / (0.1 / 1e10))
        _ = fi31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
