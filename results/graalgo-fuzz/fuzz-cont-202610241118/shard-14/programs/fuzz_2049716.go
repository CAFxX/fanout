package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 float64 = 3.14159
    var v2 uint64 = 4294967295
    var v3 float64 = (v1 - float64(1.5))
    var v4 float32 = 1e10
    var v5 uint = 4294967295
    if (65535 <= 65535) {
        fmt.Println(v1, v5)
        v3 -= 0.1
        var v6 string = "foo bar"
        _ = v6
    } else {
        v5 -= (^(v5))
    }
    v7 := []int{4294967295, 843, 9223372036854775807, -11}
    _ = len(v7)
    v8 := GenS1{F0: v2}
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int = 0
        var v11 float64 = float64((v4 * v4))
        v11 -= v11
        _ = fi9
        _ = v10
        _ = v11
    }
    v5 = 65535
    if false {
        var v12 float32 = v4
        if true {
            if len(v7) > 0 {
                v7[0] = -2147483648
            }
            fmt.Println(v1, v2)
            v13 := GenS1{F0: 4294967295}
            _ = v13
        } else {
            v1 = (((v1 + ((((v1 * v3) - 0.0) * v1) / 1.5)) + 1e10) - 1.5)
        }
        _ = v12
    } else {
        switch "hello" {
        case "":
            switch "foo bar" {
            case "a":
                v5 -= (v5 / ((v5) | 1))
            default:
                v14 := GenS1{F0: v2}
                _ = v14
            }
        default:
            var v15 uint32 = 893
            _ = v15
        }
    }
    if ((v2 << ((uint(v2)) & 63)) < 0) {
        v1 = (3.14159 * ((float64(float32((float64(((0.1 - v4) - -2.25)) + (v1 * -2.25)))) - v3) * v1))
        v4 += v4
    } else {
        v1 = v1
    }
    v2 = 18446744073709551615
    v7 = append(v7, 299)
    if ((((false || (42 < -2147483648)) || ((42 >= -128) || false)) || (0.0 != 1e10)) || true) {
        v3 -= 1.5
        if false {
            v4 = (3.14159 / (1e10 * (v4 / v4)))
            var v16 string = "\n\t"
            var v17 int = 4294967295
            _ = v16
            _ = v17
        } else {
            v4 = v4
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
    } else {
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 uint = (v5 % ((v5) | 1))
            _ = fi18
            _ = v19
        }
        var v20 uint64 = (v2 + (v2 % (((v2 / (((v2 ^ v2)) | 1))) | 1)))
        switch "a" {
        case "\n\t":
            switch "foo bar" {
            case "a":
                var v21 int64 = 520
                _ = v21
            }
        default:
            for fi22 := 0; fi22 < 14; fi22++ {
                _ = len(v7)
                v23 := GenS1{F0: v20}
                _ = fi22
                _ = v23
            }
        }
        _ = v20
    }
    v2 += (v2 - (v2 << (((v5 * (v5 | 1))) & 63)))
    v4 -= 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
