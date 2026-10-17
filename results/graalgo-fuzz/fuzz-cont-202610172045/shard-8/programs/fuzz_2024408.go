package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint8; F1 uint; F2 uint32 }

func main() {
    var v1 float32 = 0.0
    var v2 string = "a"
    var v3 string = ""
    v4 := make(map[string]string)
    var v5 uint64 = 1
    var v6 int = -9223372036854775808
    var v7 int16 = 999
    v8 := GenS1{F0: v5}
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 uint16 = 664
        switch fi9 {
        case -257:
            var v11 uint64 = v5
            _ = v11
        case -137:
            var v12 uint64 = v5
            _ = v12
        case 32768:
            v6 = fi9
        default:
            v13 := GenS1{F0: v5}
            _ = v13
        }
        fmt.Println(fi9, v10, v6)
        _ = fi9
        _ = v10
    }
    v14 := GenS2{F0: false, F1: 42}
    v15 := GenS3{F0: ("a" + "\n\t"), F1: (((1e10 - -2.25) + 3.14159) > float64(v1))}
    var v16 int16 = (v7 - v7)
    for fi17 := 0; fi17 < 14; fi17++ {
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 uint8 = 127
            _ = fi18
            _ = v19
        }
        v6 -= 113
        _ = fi17
    }
    if true {
        for fi20 := 0; fi20 < 9; fi20++ {
            v2 = (v2 + "\n\t")
            for fi21 := 0; fi21 < 17; fi21++ {
                fmt.Println(v3)
                v1 = (float32(((float64(float32((0.0 + (1e10 - 0.0)))) + 1e10) / -2.25)) + v1)
                var v22 int64 = 42
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
        v7 = v16
        fmt.Println(v5)
    }
    switch v2 {
    case "hello":
        for fi23 := 0; fi23 < 14; fi23++ {
            switch (fi23 & 1) {
            case 939:
                v24 := v4[("hello" + "")]
                _ = v24
            case 90:
                var v25 int32 = -1
                _ = v25
            case -9223372036854775808:
                fmt.Println(v16, v5)
            }
            fmt.Println(v7, v16, v1)
            _ = fi23
        }
    case "foo bar":
        for fi26 := 0; fi26 < 19; fi26++ {
            v27 := GenS4{F0: 128, F1: 1, F2: 65535}
            fmt.Println(v7)
            fmt.Println(fi26, v2, v6)
            _ = fi26
            _ = v27
        }
    default:
        fmt.Println(v2)
    }
    if (127 <= 0) {
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
