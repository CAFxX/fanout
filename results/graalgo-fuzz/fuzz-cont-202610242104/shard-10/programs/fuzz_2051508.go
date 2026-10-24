package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int; F1 int16; F2 uint32 }
func genFn0(p0 uint32) bool {
    var v5 uint8 = 255
    var v6 string = (("hello" + "") + "foo bar")
    v7 := GenS1{F0: 924, F1: 18446744073709551615, F2: -128}
    switch ("\n\t" + v6) {
    case "hello":
        var v8 int64 = -1
        _ = v8
    default:
        v5 = (^(v5))
    }
    v9 := []int{80, 213}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v9
    return true
}

func main() {
    var v1 uint16 = 65535
    v2 := []string{"", ("\n\t" + "hello"), "\n\t"}
    var v3 int = 0
    var v4 float32 = 3.14159
    var v10 string = "hello"
    v11 := GenS1{F0: v3, F1: 0, F2: 0}
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS1{F0: (^(v3)), F1: 0, F2: 127}
        if (4294967295 < 1) {
            v4 = (v4 / ((1e10 - ((1.5 / v4) * 0.0)) / v4))
        } else {
            switch v10 {
            case "foo bar":
                var v14 int = 1
                _ = v14
            default:
                var v15 uint64 = 18446744073709551615
                _ = v15
            }
            v2 = append(v2, "hello")
            v1 = v1
        }
        _ = fi12
        _ = v13
    }
    v16 := GenS1{F0: (v3 >> ((18446744073709551615) & 63)), F1: 271, F2: 127}
    if (-1 >= 0) {
        var v17 uint = 2
        _ = v17
    }
    v1 = v1
    v18 := GenS2{F0: 0, F1: 32767}
    if true {
        v19 := GenS3{F0: v3, F1: 1, F2: 2147483648}
        v20 := GenS1{F0: v3, F1: uint64(v1), F2: 0}
        _ = v19
        _ = v20
    } else {
        if (654 < (v1 % ((921) | 1))) {
            switch "" {
            case "foo bar":
                switch "a" {
                case "a":
                    v10 = v10
                default:
                    var v21 int = v3
                    _ = v21
                }
            default:
                v22 := GenS3{F0: v3, F1: -32768, F2: 332}
                _ = v22
            }
        }
    }
    v2 = append(v2, ("" + "a"))
    var v23 float32 = ((float32(0.0) / 1.5) / 1.5)
    var v24 string = v10
    var v25 float32 = v4
    var v26 float64 = 0.0
    v2 = append(v2, (("foo bar" + "foo bar") + v24))
    var v27 uint32 = 4294967295
    v28 := GenS3{F0: (v3 >> ((65535) & 63)), F1: 1000, F2: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
