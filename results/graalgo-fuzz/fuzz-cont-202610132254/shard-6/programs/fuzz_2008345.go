package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint16 }
func genFn0(p0 int64) string {
    var v5 int = -1
    for fi6 := 0; fi6 < 10; fi6++ {
        if false {
            var v7 uint = 18446744073709551615
            var v8 uint8 = 127
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 7; fi9++ {
            if false {
                var v10 uint32 = 65535
                _ = v10
            } else {
                switch "hello" {
                case "foo bar":
                    switch "hello" {
                    case "foo bar":
                        var v11 int64 = -785
                        _ = v11
                    case "\n\t":
                        var v12 int64 = 482
                        _ = v12
                    case "hello":
                        var v13 float32 = ((0.1 + (((((-2.25 - float32(0.1)) / (1e10 - (1.5 * (-2.25 * (((0.0 - 3.14159) / ((1e10 / 1.0) / 1.0)) + (1.5 + 1e10)))))) / 0.1) + 0.1) / -2.25)) / ((0.1 / ((float32(0.1) / 3.14159) / 1.5)) / 1e10))
                        v14 := GenS1{F0: 65535, F1: 0, F2: fi6}
                        _ = v13
                        _ = v14
                    }
                    switch ("hello" + ("\n\t" + ("foo bar" + (("" + "foo bar") + "\n\t")))) {
                    case "a":
                        var v15 uint = 0
                        _ = v15
                    case "\n\t":
                        v16 := []int{v5, (v5 - fi6)}
                        _ = v16
                    }
                case "hello":
                    v17 := GenS1{F0: 255, F1: 127, F2: 32768}
                    _ = v17
                case "":
                    v5 = fi6
                }
            }
            v18 := GenS2{F0: -1}
            v19 := GenS3{F0: 82, F1: 18446744073709551615}
            _ = fi9
            _ = v18
            _ = v19
        }
        _ = fi6
    }
    _ = p0
    _ = v5
    return "\n\t"
}

func main() {
    var v1 int32 = 2147483647
    var v2 int16 = -1
    var v3 uint = 65535
    var v4 int8 = 100
    v2 = 4
    fmt.Println(v4)
    v3 += v3
    v20 := GenS1{F0: 429, F1: (^(v4)), F2: 630}
    v21 := GenS4{F0: 9223372036854775807, F1: 255}
    v22 := []int{4294967295, 128}
    for fi23 := 0; fi23 < 14; fi23++ {
        var v24 int8 = (v4 << (((v3 - v3)) & 63))
        v24 -= v24
        if false {
            _ = len(v22)
            for fi25 := 0; fi25 < 4; fi25++ {
                v26 := GenS2{F0: (^(v4))}
                _ = fi25
                _ = v26
            }
        }
        _ = fi23
        _ = v24
    }
    var v27 int = 4294967295
    var v28 uint64 = 42
    switch 897 {
    case 127:
        if len(v22) > 0 {
            v22[0] = (v27 >> ((2) & 63))
        }
    case -611:
        fmt.Println(v3, v3)
    }
    var v29 string = ("a" + ("" + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
