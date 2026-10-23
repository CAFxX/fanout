package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint }
func genFn0(p0 int32, p1 uint, p2 uint8) uint64 {
    if (1 != 2147483648) {
        var v5 int64 = 549
        _ = v5
    }
    var v6 uint = (p1 << (((p1 % ((p1) | 1))) & 63))
    v7 := GenS1{F0: (p0 >> ((v6) & 63)), F1: -2147483648}
    if true {
        if true {
            var v8 float32 = (-2.25 - 0.0)
            v9 := []int{621, 179}
            _ = v8
            _ = v9
        }
        var v10 int = 127
        fmt.Println(p0, v6, v10)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 128
    v3 := []int{-1, 0}
    var v4 float64 = -2.25
    var v11 float32 = 0.0
    var v12 int = -259
    if (!(!(float32((v4 + v4)) <= 0.0))) {
        switch "\n\t" {
        case "foo bar":
            var v13 uint = 65535
            _ = v13
        case "a":
            v14 := GenS2{F0: 81, F1: (v2 + (v2 >> ((110) & 63))), F2: int16(v12)}
            _ = v14
        case "\n\t":
            var v15 int64 = 9223372036854775807
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 uint16 = 690
            var v18 int8 = 85
            for fi19 := 0; fi19 < 1; fi19++ {
                switch "hello" {
                case "a":
                    var v20 uint32 = 1
                    _ = v20
                }
                _ = fi19
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        v21 := GenS1{F0: -2147483648, F1: v12}
        _ = v21
    }
    v22 := GenS2{F0: 0, F1: 8, F2: 42}
    var v23 uint8 = (v2 + v2)
    v24 := GenS3{F0: 667}
    v4 = v4
    var v25 int16 = int16(v12)
    var v26 int16 = v25
    for fi27 := 0; fi27 < 11; fi27++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
