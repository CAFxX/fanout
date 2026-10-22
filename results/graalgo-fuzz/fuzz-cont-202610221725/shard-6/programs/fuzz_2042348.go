package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint; F2 uint32 }
type GenS3 struct { F0 string }
func genFn0() uint32 {
    if false {
        v4 := GenS1{F0: 2147483647}
        _ = v4
    }
    var v5 uint64 = 2
    switch 590 {
    case 2147483647:
        fmt.Println(v5)
    case 1:
        var v6 uint8 = 128
        _ = v6
    case 4294967295:
        v5 = 705
    }
    fmt.Println(v5)
    _ = v5
    return 65535
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]int)
    var v3 int32 = 209
    v1 = append(v1, 4294967295)
    switch -1 {
    case 9223372036854775807:
        _ = len(v1)
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    v3 = 663
    v3 = (v3 / (((v3 >> ((42) & 63))) | 1))
    v7 := GenS2{F0: "", F1: 907, F2: 0}
    if (1.5 < 3.14159) {
        var v8 int64 = 691
        _ = v8
    } else {
        if ((-2.25 + -2.25) > -2.25) {
            var v9 int32 = v3
            for fi10 := 0; fi10 < 14; fi10++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi10
            }
            fmt.Println(v9, v9)
            _ = v9
        }
    }
    switch ("" + ("a" + "\n\t")) {
    case "\n\t":
        v3 = v3
    case "foo bar":
        v11 := GenS3{F0: ("" + "\n\t")}
        _ = v11
    }
    if ((2147483647 < 65535) && true) {
        var v12 uint8 = 255
        v12 = (v12 % ((v12) | 1))
        _ = v12
    }
    var v13 string = "\n\t"
    var v14 uint = 255
    delete(v2, 439)
    var v15 string = v13
    var v16 uint = v14
    var v17 bool = (0 >= -767)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
}
