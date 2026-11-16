package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int64; F1 uint32 }

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]uint32)
    var v3 float32 = ((1e10 * float32((((float64(3.14159) * 3.14159) / 3.14159) * 1e10))) * (0.0 + float32(0.0)))
    fmt.Println(v3, v1)
    switch "hello" {
    case "a":
        _ = len(v2)
    case "":
        var v4 int32 = 0
        _ = v4
    case "foo bar":
        var v5 uint32 = 908
        _ = v5
    }
    v1 -= v1
    fmt.Println(v1, v1)
    v1 -= v1
    v6 := GenS1{F0: uint16(v1)}
    v7 := GenS1{F0: 32767}
    var v8 uint = 0
    v9 := GenS1{F0: 42}
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(v1)
        switch (fi10 - 68) {
        case -516:
            for fi11 := 0; fi11 < 5; fi11++ {
                _ = len(v2)
                _ = len(v2)
                switch 381 {
                case 90:
                    v1 = v1
                case 128:
                    v12 := GenS2{F0: (^(v8)), F1: 566}
                    _ = v12
                case -180:
                    v1 = v1
                default:
                    v8 -= 255
                }
                _ = fi11
            }
        case 353:
            v13 := GenS2{F0: v8, F1: -1}
            _ = v13
        }
        _ = fi10
    }
    var v14 uint32 = 4294967295
    fmt.Println(v14, v8)
    var v15 string = "hello"
    v3 = (v3 * (v3 * v3))
    switch (v15 + "foo bar") {
    case "a":
        delete(v2, (("a" + v15) + "hello"))
    case "hello":
        _ = len(v2)
    case "\n\t":
        v14 -= uint32(v1)
    }
    var v16 int64 = 1
    v17 := GenS3{F0: (v16 % (((v16 / ((v16) | 1))) | 1)), F1: (v14 - (v14 ^ 55))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
