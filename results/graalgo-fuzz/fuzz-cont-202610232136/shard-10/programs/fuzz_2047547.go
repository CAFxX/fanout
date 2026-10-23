package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint = 65535
    v2 := make(map[string]bool)
    var v3 int32 = 2147483647
    var v4 int64 = 9223372036854775807
    switch "\n\t" {
    case "a":
        v3 -= v3
    default:
        for fi5 := 0; fi5 < 7; fi5++ {
            v4 += (v4 & v4)
            _ = fi5
        }
    }
    v2["foo bar"] = true
    var v6 bool = false
    var v7 uint64 = 0
    delete(v2, ("\n\t" + "\n\t"))
    v3 += v3
    for fi8 := 0; fi8 < 6; fi8++ {
        if (false || (v3 >= 560)) {
            if (!true) {
                for fi9 := 0; fi9 < 7; fi9++ {
                    switch (fi8 & -1) {
                    case 32768:
                        var v10 uint32 = 1
                        _ = v10
                    case -9223372036854775808:
                        v6 = (!(!v6))
                    case 316:
                        v11 := GenS1{F0: (v4 | -1), F1: -18}
                        _ = v11
                    }
                    var v12 uint16 = 65535
                    v4 = (v4 / (((v4 * v4)) | 1))
                    _ = fi9
                    _ = v12
                }
                var v13 int = 2147483647
                var v14 int8 = 42
                _ = v13
                _ = v14
            }
        } else {
            if v6 {
                for fi15 := 0; fi15 < 1; fi15++ {
                    var v16 int16 = -1
                    v17 := []int{fi8, fi15, (^(fi8)), -519}
                    v4 = 9
                    var v18 uint64 = 1
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            } else {
                v19 := GenS1{F0: -1, F1: 0}
                _ = v19
            }
            fmt.Println(v7, v1, v6)
        }
        v3 = (^(v3))
        _ = fi8
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS2{F0: 375}
        v22 := []int{fi20, (fi20 / ((int(v7)) | 1))}
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
