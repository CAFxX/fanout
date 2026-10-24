package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint8 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint8 }

func main() {
    v1 := []int{-403, 42}
    var v2 uint32 = 425
    var v3 uint32 = v2
    v4 := GenS1{F0: -1, F1: 229, F2: 255}
    var v5 uint16 = 65535
    v3 += (v2 | (^(v3)))
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if len(v1) > 0 {
            v1[0] = -223
        }
        for fi6 := 0; fi6 < 5; fi6++ {
            fmt.Println(v3, v5)
            var v7 float64 = -2.25
            if ((v2 << ((1) & 63)) != 507) {
                var v8 uint = 59
                _ = v8
            }
            _ = fi6
            _ = v7
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v9 uint64 = 2
        _ = v9
    }
    fmt.Println(v3, v2)
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint16 = 759
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v11 -= v11
        }
        _ = fi10
        _ = v11
    }
    v5 += (v5 % ((v5) | 1))
    var v12 uint = 255
    v5 -= v5
    fmt.Println(v3, v3)
    fmt.Println(v5, v12, v12)
    var v13 int16 = 32767
    v1 = append(v1, -659)
    if true {
        switch "foo bar" {
        case "":
            switch "foo bar" {
            case "foo bar":
                v2 += 42
            case "\n\t":
                v14 := GenS2{F0: 117, F1: 100}
                _ = v14
            default:
                v13 += v13
            }
        case "hello":
            fmt.Println(v2, v12)
        }
    }
    v1 = append(v1, 2147483647)
    var v15 string = "\n\t"
    fmt.Println(v12)
    if true {
        switch "\n\t" {
        case "":
            v12 = (^(v12))
        case "\n\t":
            _ = len(v1)
        case "foo bar":
            v16 := GenS3{F0: 243, F1: -102, F2: 0}
            _ = v16
        default:
            var v17 string = (((v15 + v15) + "a") + v15)
            _ = v17
        }
        v3 = 0
        v1 = append(v1, 0)
    } else {
        if true {
            if len(v1) > 0 {
                v1[0] = 128
            }
        } else {
            var v18 int32 = 2147483647
            if false {
                v1 = append(v1, 4294967295)
                var v19 string = v15
                v2 = v2
                _ = v19
            }
            _ = v18
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
