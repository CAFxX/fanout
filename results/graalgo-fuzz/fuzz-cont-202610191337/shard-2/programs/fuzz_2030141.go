package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 bool }
func genFn0() int {
    var v5 uint64 = 0
    v6 := []int{32767, 1, 32768, -1}
    _ = v5
    _ = v6
    return 693
}
func genFn1(p0 uint64, p1 bool) uint8 {
    fmt.Println(p0)
    v7 := GenS1{F0: 1}
    if p1 {
        fmt.Println(p0, p0)
        var v8 int16 = -32768
        v9 := GenS2{F0: 1}
        _ = v8
        _ = v9
    }
    if p1 {
        v10 := GenS2{F0: p0}
        if (827 > 63) {
            var v11 bool = (65535 < 794)
            fmt.Println(v11, p1)
            _ = v11
        } else {
            p0 -= p0
            for fi12 := 0; fi12 < 10; fi12++ {
                fmt.Println(fi12, p0, p0)
                _ = fi12
            }
            var v13 uint32 = 4294967295
            _ = v13
        }
        _ = v10
    }
    var v14 uint32 = 0
    _ = p0
    _ = p1
    _ = v14
    _ = v7
    return 230
}

func main() {
    v1 := []int{-1, 9223372036854775807, 128, 32768}
    v2 := []string{"", ""}
    v3 := []string{"a", "hello", "", ("" + "foo bar")}
    var v4 uint32 = 0
    v15 := GenS3{F0: 0, F1: true}
    v3 = append(v3, "")
    v4 = 0
    var v16 int64 = 0
    switch (("a" + ("" + "foo bar")) + ("\n\t" + (("foo bar" + "foo bar") + "a"))) {
    case "foo bar":
        var v17 float64 = 1e10
        _ = v17
    case "hello":
        switch "foo bar" {
        case "hello":
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 uint = 904
                var v20 int8 = 100
                var v21 int64 = (v16 >> ((226) & 63))
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        case "\n\t":
            switch ("foo bar" + "") {
            case "hello":
                for fi22 := 0; fi22 < 16; fi22++ {
                    fmt.Println(v16, v4, v16)
                    _ = fi22
                }
            case "\n\t":
                var v23 float64 = -2.25
                _ = v23
            default:
                var v24 float32 = 1e10
                _ = v24
            }
        case "":
            v4 += (v4 / (((v4 % ((988) | 1))) | 1))
        default:
            var v25 uint16 = 32767
            _ = v25
        }
    case "\n\t":
        var v26 int8 = 44
        _ = v26
    default:
        var v27 float64 = -2.25
        _ = v27
    }
    fmt.Println(v16, v16)
    v4 -= v4
    v28 := GenS2{F0: uint64(v4)}
    v4 += v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = v4
    var v29 uint16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
