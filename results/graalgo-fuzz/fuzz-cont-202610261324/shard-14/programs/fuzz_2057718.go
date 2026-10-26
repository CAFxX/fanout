package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint32, p1 uint16) int16 {
    for fi4 := 0; fi4 < 14; fi4++ {
        fmt.Println(p0, p0, p1)
        var v5 uint64 = 2
        _ = fi4
        _ = v5
    }
    p1 = 659
    v6 := []int{4294967295}
    switch "\n\t" {
    case "hello":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "a":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "\n\t":
        var v7 string = "foo bar"
        _ = v7
    default:
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 float32 = 3.14159
    var v2 int64 = 9223372036854775807
    var v3 bool = (65535 >= 255)
    var v8 int64 = (v2 * v2)
    if (!v3) {
        v9 := GenS1{F0: 1, F1: -1}
        if (331 <= 42) {
            var v10 uint = 255
            _ = v10
        }
        v3 = v3
        _ = v9
    } else {
        fmt.Println(v1)
    }
    var v11 uint = 227
    var v12 uint64 = 1
    for fi13 := 0; fi13 < 3; fi13++ {
        v2 = (^(v2))
        _ = fi13
    }
    switch "a" {
    case "\n\t":
        switch "" {
        case "":
            fmt.Println(v1)
        default:
            var v14 uint = (v11 * v11)
            _ = v14
        }
    case "hello":
        v11 -= v11
    case "a":
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 uint = (^(v11))
            if (true || (genFn0(4294967295, 42) == -1)) {
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 uint32 = 1
                    fmt.Println(v18, v16)
                    fmt.Println(v3, v12, v8)
                    _ = fi17
                    _ = v18
                }
                fmt.Println(v2, v3)
            } else {
                v12 += 294
                v19 := []int{-755, fi15, (fi15 % ((int(v12)) | 1))}
                _ = v19
            }
            _ = fi15
            _ = v16
        }
    default:
        var v20 float32 = v1
        _ = v20
    }
    v21 := []int{-481, -88}
    v1 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v8
}
