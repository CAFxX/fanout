package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 bool = false
    var v2 int8 = -128
    var v3 int16 = -32768
    var v4 int = int(v2)
    v5 := []string{"a", "\n\t", "hello", ""}
    fmt.Println(v3)
    var v6 int = v4
    v7 := GenS1{F0: 943, F1: ""}
    v6 -= -9223372036854775808
    var v8 float64 = ((3.14159 - float64(1e10)) - float64(0.0))
    for fi9 := 0; fi9 < 18; fi9++ {
        v4 -= v6
        v10 := GenS2{F0: 127}
        if v1 {
            fmt.Println(v1, fi9)
        } else {
            for fi11 := 0; fi11 < 2; fi11++ {
                if true {
                    switch (fi11 + (v4 >> ((334) & 63))) {
                    case 127:
                        var v12 uint16 = 2
                        _ = v12
                    default:
                        var v13 int64 = -1
                        _ = v13
                    }
                    var v14 int32 = -942
                    _ = v14
                } else {
                    v5 = append(v5, "a")
                }
                switch (fi9 + -9223372036854775808) {
                case 42:
                    v15 := GenS2{F0: 392}
                    _ = v15
                case 679:
                    v16 := GenS3{F0: 869}
                    _ = v16
                }
                var v17 uint8 = 255
                _ = fi11
                _ = v17
            }
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v2)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v18 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
