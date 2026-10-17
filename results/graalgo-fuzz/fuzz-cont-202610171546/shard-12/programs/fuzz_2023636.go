package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 float32 = 3.14159
    var v2 uint64 = 18446744073709551615
    var v3 uint16 = 0
    v4 := []int{32767}
    var v5 float64 = -2.25
    v6 := make(map[string]string)
    if false {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: v3, F1: 1}
            fmt.Println(v5, fi7, v1)
            _ = fi7
            _ = v8
        }
        var v9 float32 = (float32(v5) + -2.25)
        v10 := GenS1{F0: v3, F1: 310}
        _ = v10
        _ = v9
    }
    switch "\n\t" {
    case "":
        for fi11 := 0; fi11 < 16; fi11++ {
            v2 += (v2 & (v2 >> ((612) & 63)))
            var v12 float64 = v5
            var v13 int16 = 32767
            _ = fi11
            _ = v12
            _ = v13
        }
    case "a":
        var v14 string = ""
        _ = v14
    }
    switch 9223372036854775807 {
    case 32767:
        switch -2147483648 {
        case 539:
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 string = ""
                _ = fi15
                _ = v16
            }
        case 4294967295:
            fmt.Println(v2, v3, v5)
        default:
            v2 -= (v2 | (v2 / ((18446744073709551615) | 1)))
        }
    case -9223372036854775808:
        v2 -= v2
    default:
        _ = len(v6)
    }
    if (true && true) {
        var v17 uint64 = (v2 * 2)
        _ = v17
    }
    v18 := GenS2{F0: 9223372036854775807}
    var v19 int = 0
    v19 = (v19 * 712)
    if len(v4) > 0 {
        v4[0] = (v19 * 1)
    }
    if true {
        switch "a" {
        case "\n\t":
            v1 = float32(v5)
        }
        var v20 string = (("\n\t" + "a") + ("a" + "\n\t"))
        var v21 uint64 = v2
        _ = v20
        _ = v21
    } else {
        v22 := GenS3{F0: 982, F1: v2}
        _ = len(v4)
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
