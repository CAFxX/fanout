package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int; F2 int16 }
type GenS3 struct { F0 bool; F1 uint32; F2 int }
func genFn0(p0 int, p1 int64, p2 string) int8 {
    p1 = p1
    var v5 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 97
}

func main() {
    v1 := []int{-29, 42, 128}
    var v2 bool = (9223372036854775807 > 686)
    v3 := []string{}
    var v4 float32 = 0.1
    v4 += (v4 - v4)
    switch 42 {
    case -992:
        v6 := GenS1{F0: 42}
        _ = v6
    default:
        v2 = false
    }
    v7 := GenS2{F0: 1, F1: 127, F2: -1}
    switch "hello" {
    case "foo bar":
        var v8 string = "foo bar"
        _ = v8
    case "hello":
        var v9 uint16 = 42
        _ = v9
    default:
        if v2 {
            v2 = v2
            _ = len(v3)
            var v10 int16 = -1
            _ = v10
        } else {
            var v11 int8 = 62
            _ = len(v1)
            _ = v11
        }
    }
    if v2 {
        switch "foo bar" {
        case "\n\t":
            var v12 float64 = 0.1
            _ = v12
        case "a":
            v4 = v4
        }
    } else {
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(fi13)
            _ = fi13
        }
        v14 := GenS1{F0: -866}
        _ = v14
    }
    var v15 uint32 = 2
    var v16 int = 832
    for fi17 := 0; fi17 < 4; fi17++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch fi17 {
        case 32768:
            if (!((v2 || (!v2)) || ((true && false) && (!v2)))) {
                switch fi17 {
                case 127:
                    v18 := GenS2{F0: 42, F1: 322, F2: 18}
                    _ = v18
                case 846:
                    if (v2 || (!(-2.25 > 1e10))) {
                        v19 := GenS1{F0: 267}
                        _ = v19
                    } else {
                        var v20 uint32 = (v15 & v15)
                        v2 = (v2 || ((v4 != 1.5) && v2))
                        _ = v20
                    }
                default:
                    v21 := GenS3{F0: v2, F1: 2147483647, F2: (-(v16))}
                    _ = v21
                }
            }
        }
        _ = fi17
    }
    if (!(-2.25 >= 1e10)) {
        if len(v1) > 0 {
            v1[0] = (v16 | v16)
        }
        switch v16 {
        case 32768:
            var v22 uint64 = 1
            _ = v22
        case 9223372036854775807:
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 uint64 = 42
                _ = fi23
                _ = v24
            }
        }
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        _ = len(v1)
        var v26 uint32 = v15
        _ = fi25
        _ = v26
    }
    v15 += v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
