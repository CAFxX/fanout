package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint; F1 uint32; F2 int8 }

func main() {
    var v1 int8 = 1
    var v2 float32 = 3.14159
    var v3 bool = true
    var v4 uint = 4294967295
    var v5 float32 = (v2 + (v2 + (v2 - (v2 * float32(1e10)))))
    v5 = (v5 - 3.14159)
    v6 := GenS1{F0: 1, F1: 2147483647}
    switch -995 {
    case -1:
        var v7 int64 = 1
        _ = v7
    case 32767:
        v4 += 2
    case 1:
        if v3 {
            var v8 int8 = 0
            _ = v8
        }
    default:
        fmt.Println(v4)
    }
    v9 := []int{9223372036854775807, 9223372036854775807, 0}
    v2 = v2
    v10 := GenS2{F0: -809, F1: 23}
    fmt.Println(v4, v2)
    switch "\n\t" {
    case "foo bar":
        var v11 uint8 = 2
        _ = v11
    case "":
        switch ("" + ("" + "foo bar")) {
        case "\n\t":
            v5 = float32(0.1)
        }
    case "\n\t":
        for fi12 := 0; fi12 < 7; fi12++ {
            fmt.Println(v3)
            if (!v3) {
                if v3 {
                    v13 := GenS3{F0: 32767, F1: v1}
                    switch "" {
                    case "foo bar":
                        var v14 int = (fi12 & -416)
                        _ = v14
                    case "a":
                        fmt.Println(v2, v3)
                    case "\n\t":
                        _ = len(v9)
                    }
                    _ = v13
                }
            } else {
                fmt.Println(v3, v4)
                v4 = (v4 % ((v4) | 1))
            }
            _ = fi12
        }
    }
    if v3 {
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 uint = v4
            v17 := GenS4{F0: (^(v16)), F1: 0, F2: v1}
            fmt.Println(fi15, v5, v16)
            _ = fi15
            _ = v16
            _ = v17
        }
        v3 = false
        v3 = v3
    }
    var v18 int64 = -790
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
