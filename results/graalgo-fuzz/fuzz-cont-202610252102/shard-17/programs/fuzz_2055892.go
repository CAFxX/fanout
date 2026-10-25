package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int }
type GenS2 struct { F0 uint32; F1 bool; F2 int64 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS7 struct { F0 uint32 }

func main() {
    var v1 float32 = 1.5
    var v2 bool = false
    var v3 float32 = float32(3.14159)
    var v4 bool = (false && v2)
    var v5 uint32 = 1
    v2 = true
    if false {
        switch 9223372036854775807 {
        case 4294967295:
            v6 := GenS1{F0: -700, F1: true, F2: 314}
            _ = v6
        case 128:
            v4 = (v4 && true)
        }
        v7 := GenS1{F0: 0, F1: v2, F2: 508}
        var v8 uint64 = 495
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        switch fi9 {
        case -695:
            var v10 float64 = 0.0
            _ = v10
        case -14:
            switch fi9 {
            case 42:
                var v11 string = (("a" + "") + ("foo bar" + ("" + "hello")))
                _ = v11
            case -1:
                var v12 uint64 = 1
                _ = v12
            default:
                fmt.Println(fi9, v1, v4)
            }
        default:
            fmt.Println(v5, v4)
        }
        _ = fi9
    }
    var v13 uint8 = 209
    if v2 {
        fmt.Println(v3, v1)
        var v14 int = 9223372036854775807
        var v15 uint8 = v13
        _ = v14
        _ = v15
    } else {
        fmt.Println(v2)
    }
    if v2 {
        v16 := GenS1{F0: -1, F1: (v2 || ((v5 ^ (v5 / (((v5 & v5)) | 1))) != (v5 | v5))), F2: 2147483647}
        if (v2 || (("hello" <= ("a" + "")) || v4)) {
            v13 = 0
            var v17 int8 = 1
            _ = v17
        } else {
            var v18 int8 = 100
            for fi19 := 0; fi19 < 19; fi19++ {
                v20 := GenS2{F0: (v5 * (v5 ^ 4294967295)), F1: true, F2: -254}
                var v21 bool = (1 <= v18)
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = v18
        }
        v1 -= float32(((float64(((v1 / 0.1) / (0.1 / 1e10))) - 1.5) / float64((0.0 / v1))))
        _ = v16
    } else {
        v22 := GenS3{F0: 1000, F1: 42}
        v23 := []int{}
        _ = v22
        _ = v23
    }
    if (((0 > 255) || v4) && ((127 < -1) && v2)) {
        var v24 int32 = 1
        v4 = v4
        _ = v24
    } else {
        v25 := []int{130, -759, -463}
        for fi26 := 0; fi26 < 10; fi26++ {
            if v4 {
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
            }
            _ = fi26
        }
        _ = v25
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        v3 = float32(((0.1 / (float64(((((v3 - 0.1) * float32(0.1)) - (1.5 - 3.14159)) / -2.25)) * 1.5)) * 1.5))
        if v4 {
            v5 = v5
        }
        v28 := GenS4{F0: 2, F1: -32768}
        _ = fi27
        _ = v28
    }
    v5 -= 2
    v29 := GenS5{F0: 42}
    switch 4294967295 {
    case 4294967295:
        v30 := GenS5{F0: -1}
        _ = v30
    case 127:
        fmt.Println(v13, v13, v13)
    case -617:
        v31 := GenS6{F0: 1, F1: 32767, F2: (v13 / ((128) | 1))}
        _ = v31
    default:
        if false {
            var v32 float32 = (((0.0 / 0.1) + v1) / 0.1)
            _ = v32
        }
    }
    v33 := GenS4{F0: 1, F1: -1}
    for fi34 := 0; fi34 < 6; fi34++ {
        var v35 uint16 = uint16(v5)
        _ = fi34
        _ = v35
    }
    v36 := GenS7{F0: v5}
    v3 = -2.25
    v37 := []int{}
    v3 = ((((((v1 + (v3 + 1.5)) / v3) + v3) + v1) / 0.1) - (((float32((float64(1e10) - (-2.25 * -2.25))) - (v3 / 1.0)) * v3) / ((float32(0.0) + v3) / v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
