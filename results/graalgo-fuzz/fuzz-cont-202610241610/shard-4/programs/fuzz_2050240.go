package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int; F2 uint8 }

func main() {
    var v1 int64 = 8
    var v2 float32 = 1.5
    var v3 int64 = v1
    var v4 uint64 = 4294967295
    v5 := []uint32{2, 297}
    fmt.Println(v1, v4, v4)
    v5 = append(v5, 917)
    var v6 float64 = float64((v2 / v2))
    var v7 int8 = 92
    for fi8 := 0; fi8 < 2; fi8++ {
        v5 = append(v5, 40)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        fmt.Println(v2)
        if (2147483648 < 2) {
            var v10 int = 32768
            fmt.Println(v2, v3)
            _ = v10
        } else {
            switch ("\n\t" + "foo bar") {
            case "":
                v11 := GenS1{F0: 9}
                _ = v11
            case "hello":
                switch "\n\t" {
                case "":
                    fmt.Println(fi9, v6, v7)
                case "a":
                    var v12 float32 = v2
                    _ = v12
                }
            default:
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := GenS2{F0: 2147483647, F1: 893, F2: 0}
                    _ = fi13
                    _ = v14
                }
            }
            v4 = (^(v4))
        }
        switch (fi9 ^ fi9) {
        case -702:
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS2{F0: 42, F1: (fi15 ^ fi9), F2: 127}
                if len(v5) > 0 {
                    v5[0] = 2147483647
                }
                _ = fi15
                _ = v16
            }
        case 42:
            v17 := GenS2{F0: 42, F1: 643, F2: 127}
            _ = v17
        default:
            fmt.Println(v1, v3)
        }
        _ = fi9
    }
    var v18 uint = 410
    v5 = append(v5, 2147483647)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v19 := GenS1{F0: 255}
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            v18 = (v18 << ((v18) & 63))
            fmt.Println(v6, v3)
            _ = fi21
        }
        _ = fi20
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 float32 = ((3.14159 - 1.5) + v2)
        switch (("hello" + "hello") + (("" + "") + "hello")) {
        case "":
            var v24 int64 = v1
            _ = v24
        }
        switch fi22 {
        case 42:
            fmt.Println(v18, fi22)
        case 127:
            var v25 int32 = -521
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
