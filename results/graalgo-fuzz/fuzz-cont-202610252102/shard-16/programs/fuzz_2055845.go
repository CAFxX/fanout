package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }

func main() {
    var v1 float64 = (-2.25 / 1.5)
    var v2 int32 = 2147483647
    var v3 int = -9223372036854775808
    var v4 int32 = 225
    var v5 uint32 = 65535
    v6 := []int{}
    v7 := GenS1{F0: 0, F1: -1}
    var v8 uint8 = 127
    var v9 uint32 = v5
    v9 = (v5 | (v5 / ((v5) | 1)))
    var v10 int64 = 9223372036854775807
    _ = len(v6)
    var v11 float32 = 1e10
    _ = len(v6)
    var v12 int = v3
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if len(v6) > 0 {
        v6[0] = (v12 - int(v9))
    }
    switch "\n\t" {
    case "":
        var v13 bool = (v10 > (v10 + (v10 % (((v10 | v10)) | 1))))
        _ = v13
    case "foo bar":
        if (false && true) {
            v6 = append(v6, v12)
            if true {
                var v14 int64 = v10
                _ = v14
            }
        } else {
            for fi15 := 0; fi15 < 5; fi15++ {
                v9 -= v9
                v5 = (v9 * 42)
                _ = fi15
            }
        }
    case "a":
        if len(v6) > 0 {
            v6[0] = (v3 ^ v12)
        }
    default:
        fmt.Println(v2, v2)
    }
    v8 = (v8 + v8)
    _ = len(v6)
    fmt.Println(v2, v2)
    if len(v6) > 0 {
        v6[0] = v12
    }
    if ((!(true || ((-9223372036854775808 <= v10) || true))) && true) {
        switch v3 {
        case -313:
            for fi16 := 0; fi16 < 20; fi16++ {
                v11 = 1.5
                _ = fi16
            }
        case 0:
            var v17 float32 = 3.14159
            _ = v17
        case 833:
            var v18 int = v3
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
