package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 bool; F1 int8; F2 int64 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 int32 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{707, 1, 2147483648}
    var v3 string = "\n\t"
    var v4 int16 = 234
    v5 := make(map[uint32]bool)
    var v6 uint32 = 1
    if (1e10 <= 0.0) {
        var v7 string = v3
        _ = v7
    }
    fmt.Println(v6, v6)
    var v8 int8 = 127
    var v9 float64 = 0.1
    for fi10 := 0; fi10 < 5; fi10++ {
        switch v3 {
        case "hello":
            var v11 uint64 = 0
            _ = v11
        case "\n\t":
            v4 -= v4
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            if true {
                var v13 float64 = ((v9 + v9) * ((v9 - v9) - v9))
                if len(v2) > 0 {
                    v2[0] = 2
                }
                switch v3 {
                case "":
                    var v14 uint16 = 78
                    _ = v14
                case "hello":
                    var v15 bool = (-2.25 > ((1e10 + v9) + ((v13 - v9) + (((-2.25 * float64(3.14159)) * v9) + (v9 + (v9 + (float64((-2.25 + 0.0)) + 1.5)))))))
                    _ = v15
                default:
                    switch 32767 {
                    case 270:
                        var v16 uint8 = 1
                        _ = v16
                    }
                }
                _ = v13
            }
            v6 += (v6 | (^(v6)))
            switch ("a" + "hello") {
            case "hello":
                v4 += v4
            case "foo bar":
                if false {
                    var v17 uint = 2
                    _ = v17
                } else {
                    v18 := GenS1{F0: 42, F1: 2, F2: 0}
                    switch -606 {
                    case 42:
                        var v19 uint64 = 0
                        v20 := GenS2{F0: (!(true || (true || false))), F1: v8, F2: -9223372036854775808}
                        _ = v19
                        _ = v20
                    }
                    var v21 int = fi12
                    var v22 float32 = float32(v9)
                    var v23 int16 = (v4 & v4)
                    fmt.Println(v21, v6)
                    _ = v18
                    _ = v21
                    _ = v22
                    _ = v23
                }
            }
            _ = fi12
        }
        if (false || false) {
            var v24 int = fi10
            _ = v24
        }
        _ = fi10
    }
    v6 += v6
    v25 := GenS3{F0: (-(v4)), F1: false}
    for fi26 := 0; fi26 < 9; fi26++ {
        _ = len(v2)
        fmt.Println(v8)
        _ = fi26
    }
    fmt.Println(v8, v6)
    var v27 float32 = (-2.25 * 3.14159)
    v28 := GenS4{F0: 2}
    switch -9223372036854775808 {
    case -9223372036854775808:
        v29 := GenS5{F0: 126, F1: 848}
        _ = v29
    case 4294967295:
        var v30 int = 1
        _ = v30
    case 42:
        var v31 int64 = 9223372036854775807
        _ = v31
    default:
        v4 += (v4 & v4)
    }
    if true {
        var v32 float32 = -2.25
        fmt.Println(v9, v32, v32)
        v4 -= v4
        _ = v32
    } else {
        _ = len(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
