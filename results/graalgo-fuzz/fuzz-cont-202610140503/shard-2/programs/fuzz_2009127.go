package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint8, p1 string) int64 {
    v5 := GenS1{F0: 600}
    switch (p1 + p1) {
    case "\n\t":
        v6 := GenS2{F0: 2147483648, F1: 2}
        _ = v6
    case "a":
        if (32768 != -9223372036854775808) {
            p1 = (p1 + p1)
        } else {
            switch p1 {
            case "a":
                v7 := GenS1{F0: 9223372036854775807}
                _ = v7
            default:
                v8 := GenS1{F0: 1}
                _ = v8
            }
        }
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 int = (fi9 - fi9)
        var v11 int16 = 42
        _ = fi9
        _ = v10
        _ = v11
    }
    p1 = (((p1 + (p1 + p1)) + p1) + p1)
    var v12 int64 = 932
    _ = p0
    _ = p1
    _ = v12
    _ = v5
    return v12
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int = 4294967295
    var v3 uint32 = 65535
    var v4 uint64 = 0
    v13 := []int{(v2 & 42)}
    var v14 bool = false
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    if ((994 < 2) && (((-2.25 * (0.0 * 0.1)) * 3.14159) > 0.0)) {
        for fi15 := 0; fi15 < 13; fi15++ {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            v16 := GenS3{F0: (!(true && ((v3 * (^(v3))) > (v3 + (^(v3))))))}
            _ = fi15
            _ = v16
        }
        v13 = append(v13, -1)
    }
    var v17 uint16 = 159
    v18 := GenS3{F0: v14}
    _ = len(v13)
    var v19 int32 = 42
    var v20 uint64 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
