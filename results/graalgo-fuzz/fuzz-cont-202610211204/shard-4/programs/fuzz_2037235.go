package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int32 }

func main() {
    var v1 int = -2147483648
    v2 := []string{}
    v3 := []string{("hello" + "foo bar")}
    v4 := []string{"a"}
    var v5 uint64 = 18446744073709551615
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 136, F1: 2}
        _ = fi6
        _ = v7
    }
    fmt.Println(v1, v5)
    if false {
        if (42 <= 1000) {
            for fi8 := 0; fi8 < 18; fi8++ {
                for fi9 := 0; fi9 < 15; fi9++ {
                    var v10 int16 = 32767
                    _ = fi9
                    _ = v10
                }
                var v11 uint = 661
                v5 = v5
                _ = fi8
                _ = v11
            }
        } else {
            var v12 uint32 = 4294967295
            _ = v12
        }
    } else {
        v13 := GenS1{F0: v1, F1: 354}
        _ = v13
    }
    var v14 float64 = 0.0
    for fi15 := 0; fi15 < 1; fi15++ {
        switch "\n\t" {
        case "hello":
            var v16 float32 = ((-2.25 - (0.1 * (1.5 * 0.1))) - 1.5)
            _ = v16
        case "foo bar":
            v17 := GenS2{F0: 32, F1: 42}
            _ = v17
        case "\n\t":
            var v18 uint64 = v5
            _ = v18
        default:
            v14 += (v14 / v14)
        }
        var v19 int = -9223372036854775808
        switch 128 {
        case 2147483647:
            for fi20 := 0; fi20 < 18; fi20++ {
                if (v14 <= float64(3.14159)) {
                    fmt.Println(v19)
                }
                var v21 int16 = -32768
                _ = fi20
                _ = v21
            }
        case 0:
            var v22 string = ""
            _ = v22
        case 32768:
            v1 = (fi15 % ((42) | 1))
        }
        _ = fi15
        _ = v19
    }
    var v23 int8 = 100
    var v24 int = 4294967295
    v24 = v24
    _ = len(v3)
    var v25 uint32 = 42
    var v26 uint32 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
