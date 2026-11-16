package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int }
func genFn0(p0 int64) uint64 {
    var v6 int64 = 728
    var v7 float32 = (1e10 * 0.1)
    var v8 int32 = -131
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := []string{("a" + ""), "hello", ("hello" + (("a" + "") + "hello"))}
    var v2 string = ""
    v3 := []uint32{1}
    var v4 int16 = -1
    var v5 uint = 504
    v4 += (v4 / ((v4) | 1))
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(v2)
        _ = fi9
    }
    v2 = "a"
    var v10 string = (("a" + "\n\t") + "foo bar")
    _ = len(v3)
    v4 -= (v4 | (v4 << ((v5) & 63)))
    fmt.Println(v10, v4)
    switch "foo bar" {
    case "a":
        fmt.Println(v5)
    }
    if (-128 < 127) {
        v3 = append(v3, 2147483647)
        v11 := GenS1{F0: 249, F1: 1}
        var v12 int8 = 42
        _ = v11
        _ = v12
    } else {
        switch v10 {
        case "\n\t":
            switch 127 {
            case -1:
                var v13 float64 = -2.25
                _ = v13
            default:
                var v14 float64 = (float64((3.14159 + 1.5)) - ((((1e10 * 1.5) - ((0.1 / (1.5 / (((-2.25 * ((0.1 / 0.1) - 1e10)) / 1.5) / (((0.0 + -2.25) * (1.5 / (float64((-2.25 * -2.25)) / (3.14159 + 1e10)))) / (1.5 - ((1e10 - 0.0) * 1e10)))))) - 3.14159)) - ((float64((((-2.25 - 3.14159) - 1e10) * ((-2.25 * 0.1) - 3.14159))) * 0.1) - 0.0)) - 0.0))
                _ = v14
            }
        }
        if true {
            if true {
                for fi15 := 0; fi15 < 20; fi15++ {
                    fmt.Println(v5)
                    fmt.Println(v5, v4, fi15)
                    _ = fi15
                }
                v1 = append(v1, (v2 + v10))
            } else {
                var v16 int16 = (v4 % (((v4 - v4)) | 1))
                var v17 int32 = 42
                _ = v16
                _ = v17
            }
            v10 = (("" + "a") + v10)
            v5 -= v5
        }
    }
    v18 := GenS2{F0: 985, F1: 32767}
    switch ((v10 + v10) + v2) {
    case "hello":
        if (-2.25 < float32(-2.25)) {
            v2 = v2
        } else {
            var v19 int32 = 2147483647
            v4 = v4
            v20 := GenS3{F0: 78}
            _ = v19
            _ = v20
        }
    }
    v1 = append(v1, v10)
    fmt.Println(v4, v10)
    v5 += v5
    v1 = append(v1, "\n\t")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
