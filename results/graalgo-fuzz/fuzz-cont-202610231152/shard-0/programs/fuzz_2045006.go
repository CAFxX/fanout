package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int }
func genFn0(p0 int16, p1 string, p2 int) uint {
    fmt.Println(p2, p1)
    var v7 uint16 = 230
    if (128 <= 183) {
        fmt.Println(p1, p1, p2)
    }
    switch "\n\t" {
    case "hello":
        var v8 bool = true
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 567
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    v3 := make(map[int]string)
    var v4 string = ""
    var v5 int8 = -1
    var v6 uint = 18446744073709551615
    if len(v2) > 0 {
        v2[0] = 65535
    }
    switch v4 {
    case "hello":
        var v9 int16 = -1
        _ = v9
    }
    v10 := GenS1{F0: -1, F1: "foo bar", F2: (v6 & (^(v6)))}
    _ = len(v2)
    var v11 int = 1
    fmt.Println(v5)
    v2 = append(v2, uint32(v6))
    v12 := GenS1{F0: 0, F1: "foo bar", F2: v6}
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := GenS1{F0: -9223372036854775808, F1: v4, F2: (v6 - 621)}
            _ = fi14
            _ = v15
        }
        switch (fi13 >> ((65535) & 63)) {
        case 2147483647:
            switch v11 {
            case -281:
                fmt.Println(v11)
            }
        case -831:
            for fi16 := 0; fi16 < 18; fi16++ {
                v17 := GenS2{F0: 255}
                _ = fi16
                _ = v17
            }
        case 9223372036854775807:
            var v18 uint16 = 32767
            _ = v18
        }
        if false {
            delete(v3, (fi13 >> (((v6 % ((v6) | 1))) & 63)))
            var v19 int32 = -952
            _ = v19
        } else {
            switch "\n\t" {
            case "hello":
                fmt.Println(v6)
            case "":
                fmt.Println(fi13, v4)
            default:
                if true {
                    var v20 uint32 = 2
                    _ = v20
                } else {
                    var v21 int64 = 1
                    _ = v21
                }
            }
            switch fi13 {
            case 128:
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 string = ("foo bar" + (v4 + (v4 + "\n\t")))
                    _ = fi22
                    _ = v23
                }
            default:
                for fi24 := 0; fi24 < 2; fi24++ {
                    v25 := GenS3{F0: -109, F1: (^(v11))}
                    fmt.Println(v5, v6)
                    var v26 uint = (v6 & 387)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
        }
        _ = fi13
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
