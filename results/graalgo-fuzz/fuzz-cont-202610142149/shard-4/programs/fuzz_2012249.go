package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS3 struct { F0 string; F1 int8; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8 }
func genFn0(p0 string, p1 int16, p2 int64) string {
    var v6 string = "hello"
    v7 := GenS1{F0: p2}
    switch "hello" {
    case "":
        fmt.Println(v6, p2, p0)
    case "hello":
        p0 = "foo bar"
    default:
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (v6 + "foo bar")
}

func main() {
    var v1 uint16 = 0
    var v2 int32 = 42
    var v3 int = 32767
    v4 := make(map[uint32]bool)
    v5 := []int{(v3 << ((0) & 63)), 32768}
    fmt.Println(v2)
    for fi8 := 0; fi8 < 17; fi8++ {
        if false {
            v3 = v3
            v9 := GenS1{F0: 1}
            var v10 int64 = -9223372036854775808
            _ = v10
            _ = v9
        }
        _ = fi8
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int64 = 106
        for fi13 := 0; fi13 < 13; fi13++ {
            if false {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v3, v2)
                case "hello":
                    if true {
                        _ = len(v4)
                        var v14 uint = 713
                        v1 = v1
                        v2 = (v2 + (v2 * 0))
                        _ = v14
                    }
                    v1 = (v1 % ((v1) | 1))
                    v15 := GenS2{F0: 1, F1: 2147483647, F2: v1}
                    _ = v15
                }
                v3 = (fi13 >> ((18446744073709551615) & 63))
            } else {
                v16 := GenS3{F0: "foo bar", F1: -1, F2: "hello"}
                v1 += v1
                var v17 uint32 = 42
                _ = v16
                _ = v17
            }
            _ = fi13
        }
        var v18 bool = (!(2 != 128))
        _ = fi11
        _ = v12
        _ = v18
    }
    v19 := v4[952]
    for fi20 := 0; fi20 < 6; fi20++ {
        v21 := v4[2147483648]
        v2 -= (v2 << ((1) & 63))
        _ = fi20
        _ = v21
    }
    fmt.Println(v19)
    v2 += 17
    v1 = v1
    v3 = (v3 | v3)
    for fi22 := 0; fi22 < 17; fi22++ {
        delete(v4, 2147483647)
        for fi23 := 0; fi23 < 6; fi23++ {
            v24 := GenS4{F0: fi23}
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    v25 := GenS5{F0: 1}
    var v26 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
