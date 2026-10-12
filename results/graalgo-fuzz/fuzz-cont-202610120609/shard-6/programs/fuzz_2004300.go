package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 bool; F1 string }

func main() {
    v1 := []int{-979, 32767}
    v2 := []string{"\n\t", "", "\n\t"}
    v3 := make(map[string]uint32)
    var v4 int = 0
    v5 := make(map[string]int)
    var v6 uint16 = 0
    v4 = v4
    switch v4 {
    case 32767:
        v6 -= v6
    }
    v6 = (v6 - v6)
    var v7 string = ("hello" + "\n\t")
    if true {
        fmt.Println(v4, v6, v6)
    } else {
        for fi8 := 0; fi8 < 19; fi8++ {
            switch fi8 {
            case 128:
                var v9 int = v4
                _ = v9
            }
            _ = fi8
        }
    }
    var v10 int32 = 512
    v11 := GenS1{F0: int8(v6), F1: (v6 - 65535)}
    switch v4 {
    case 127:
        var v12 uint64 = 822
        _ = v12
    case 2147483647:
        v5[("\n\t" + "")] = v4
    default:
        v7 = ""
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        fmt.Println(v4, v4, v10)
        _ = fi13
    }
    if false {
        if false {
            switch v7 {
            case "":
                if true {
                    _ = len(v5)
                    fmt.Println(v4, v4, v7)
                    v7 = "a"
                }
            case "\n\t":
                v1 = append(v1, (-(v4)))
            case "a":
                v14 := GenS2{F0: 42, F1: 18446744073709551615}
                _ = v14
            default:
                v4 = v4
            }
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint = 42
                if (962 <= 42) {
                    var v17 float64 = (0.1 * (1e10 + float64(((1e10 * 0.1) + float32((1e10 - float64(1.5)))))))
                    var v18 string = "hello"
                    _ = v17
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        fmt.Println(v7)
                        delete(v3, "a")
                        var v20 int64 = 9223372036854775807
                        fmt.Println(v20, fi15)
                        _ = fi19
                        _ = v20
                    }
                    fmt.Println(fi15)
                    var v21 int32 = 2147483647
                    _ = v21
                }
                _ = fi15
                _ = v16
            }
        } else {
            switch (v7 + (((("" + (("" + "foo bar") + (((v7 + (v7 + "hello")) + v7) + v7))) + v7) + v7) + v7)) {
            case "foo bar":
                v4 -= v4
            case "":
                v22 := GenS3{F0: false, F1: (v7 + v7)}
                _ = v22
            default:
                var v23 uint = 18446744073709551615
                _ = v23
            }
        }
        v24 := GenS3{F0: false, F1: v7}
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
