package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int = 861
    var v2 int = 606
    var v3 int16 = 42
    var v4 uint32 = 554
    fmt.Println(v4, v2)
    fmt.Println(v3)
    fmt.Println(v2)
    if ((v4 | v4) < 42) {
        v4 = 4294967295
        var v5 bool = false
        _ = v5
    } else {
        var v6 int = (v1 ^ v2)
        var v7 int8 = -33
        v7 += v7
        _ = v6
        _ = v7
    }
    v8 := []int{}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if (!((v2 + v2) > 85)) {
        for fi9 := 0; fi9 < 12; fi9++ {
            switch "a" {
            case "\n\t":
                fmt.Println(v2, v2, v3)
            default:
                for fi10 := 0; fi10 < 7; fi10++ {
                    if (!false) {
                        if len(v8) > 0 {
                            v8[0] = (fi10 << ((0) & 63))
                        }
                    }
                    v11 := GenS1{F0: 2, F1: -9223372036854775808, F2: 2}
                    switch "foo bar" {
                    case "":
                        v4 -= (v4 ^ v4)
                    default:
                        v4 += 2147483647
                    }
                    _ = fi10
                    _ = v11
                }
            }
            v12 := GenS2{F0: 0}
            _ = fi9
            _ = v12
        }
    }
    fmt.Println(v1, v1, v3)
    if (!true) {
        switch ("hello" + "") {
        case "foo bar":
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := GenS3{F0: 1, F1: 1}
                _ = fi13
                _ = v14
            }
        }
        fmt.Println(v3, v3, v4)
    }
    fmt.Println(v2, v3)
    if len(v8) > 0 {
        v8[0] = 128
    }
    v8 = append(v8, -199)
    var v15 int64 = 138
    if ((v15 & v15) >= v15) {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 int = v1
            for fi18 := 0; fi18 < 7; fi18++ {
                v1 -= (v17 / (((fi18 - -135)) | 1))
                _ = fi18
            }
            _ = fi16
            _ = v17
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            v3 -= v3
            _ = len(v8)
            v2 = (^(v1))
            _ = fi19
        }
    }
    _ = len(v8)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
