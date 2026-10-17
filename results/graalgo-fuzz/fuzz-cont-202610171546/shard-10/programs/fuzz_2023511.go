package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint; F1 uint32 }

func main() {
    var v1 int64 = 1
    v2 := make(map[int]bool)
    var v3 float64 = -2.25
    v4 := make(map[int]string)
    var v5 uint16 = 255
    var v6 uint8 = 27
    v3 += float64((float32(float64(((0.1 - float32((v3 + (v3 - 0.0)))) * (3.14159 / -2.25)))) / 1e10))
    v7 := GenS1{F0: 1}
    fmt.Println(v6, v6, v5)
    switch "foo bar" {
    case "a":
        var v8 uint64 = 18446744073709551615
        _ = v8
    case "":
        if true {
            fmt.Println(v1, v1)
            var v9 uint16 = v5
            _ = v9
        }
    default:
        var v10 int64 = v1
        _ = v10
    }
    if false {
        if (4294967295 != 18446744073709551615) {
            v3 = v3
            var v11 int = 32767
            _ = v11
        } else {
            v12 := []int{0}
            v3 = (((1.5 / 1.0) - float64(((float32(((1.5 + (v3 / (v3 / (v3 / v3)))) + (((((float64(float32(((0.0 * v3) - 0.1))) - float64(0.0)) - (v3 + v3)) * v3) + 1.5) + ((v3 - v3) - v3)))) * (1e10 / (1e10 * 3.14159))) - (((-2.25 - 1.5) * -2.25) * 0.1)))) - 3.14159)
            _ = v12
        }
        v13 := GenS1{F0: -32768}
        fmt.Println(v5)
        _ = v13
    } else {
        var v14 int16 = 355
        v6 += v6
        v15 := GenS2{F0: 2147483647, F1: 42, F2: (v5 - 1)}
        _ = v14
        _ = v15
    }
    if (0 != 0) {
        var v16 float64 = v3
        var v17 int8 = -128
        v18 := GenS3{F0: -372, F1: 9223372036854775807}
        _ = v16
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 7; fi19++ {
            for fi20 := 0; fi20 < 1; fi20++ {
                v5 -= v5
                var v21 int8 = -52
                var v22 int16 = int16(fi19)
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi19
        }
    }
    if false {
        switch ("a" + "foo bar") {
        case "\n\t":
            v1 += 228
        case "a":
            switch 0 {
            case 1:
                if false {
                    v5 -= v5
                    v5 += v5
                    if (681 > 2147483647) {
                        if false {
                            var v23 uint = 314
                            _ = v23
                        }
                        var v24 float64 = v3
                        _ = v24
                    }
                } else {
                    if ("" < "a") {
                        fmt.Println(v5)
                    }
                }
            case -39:
                v1 += (v1 + (v1 | v1))
            default:
                v25 := []int{-2147483648, 2147483647}
                _ = v25
            }
        case "hello":
            v1 -= v1
        }
        var v26 int32 = 0
        var v27 uint = 18446744073709551615
        _ = v26
        _ = v27
    }
    switch "foo bar" {
    case "a":
        var v28 int32 = 2147483647
        _ = v28
    case "":
        switch -2147483648 {
        case 1:
            v29 := GenS3{F0: -1, F1: 32768}
            _ = v29
        default:
            var v30 int8 = 1
            _ = v30
        }
    case "foo bar":
        switch -568 {
        case 32768:
            v31 := []int{4294967295, -9223372036854775808, int(v5), 733}
            _ = v31
        case 2147483647:
            var v32 bool = (v1 <= (v1 - v1))
            _ = v32
        case 9223372036854775807:
            for fi33 := 0; fi33 < 5; fi33++ {
                for fi34 := 0; fi34 < 19; fi34++ {
                    v35 := GenS1{F0: -1}
                    v5 = v5
                    v5 = (v5 & v5)
                    _ = fi34
                    _ = v35
                }
                _ = fi33
            }
        default:
            var v36 int16 = -1
            _ = v36
        }
    default:
        v37 := GenS4{F0: 103, F1: 42}
        _ = v37
    }
    v38 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
