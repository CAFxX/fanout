package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint32; F1 int32; F2 int8 }
type GenS4 struct { F0 uint16; F1 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -462
    var v3 uint = 181
    v2 += (v2 << (((v3 * 2)) & 63))
    for fi4 := 0; fi4 < 2; fi4++ {
        v3 += 42
        v5 := GenS1{F0: (fi4 ^ int(v3)), F1: -32768}
        for fi6 := 0; fi6 < 20; fi6++ {
            _ = len(v1)
            switch "a" {
            case "":
                var v7 float64 = float64(3.14159)
                _ = v7
            case "a":
                v3 = v3
                var v8 uint64 = 42
                v2 -= v2
                _ = v8
            case "hello":
                fmt.Println(v3, fi6)
            default:
                if (1 != 647) {
                    var v9 int8 = 42
                    var v10 int8 = (-(v9))
                    v9 += (v9 / ((-14) | 1))
                    switch "" {
                    case "hello":
                        v11 := GenS1{F0: -356, F1: 1000}
                        _ = v11
                    case "foo bar":
                        v10 = 0
                    default:
                        var v12 int = fi6
                        _ = v12
                    }
                    if ((^(v9)) != v10) {
                        v10 = 42
                        var v13 int16 = -32768
                        _ = v13
                    }
                    var v14 uint8 = 42
                    _ = v10
                    _ = v14
                    _ = v9
                }
            }
            v15 := []int{32767, fi6}
            _ = fi6
            _ = v15
        }
        _ = fi4
        _ = v5
    }
    v16 := GenS2{F0: 56, F1: 0}
    switch (("hello" + ((("a" + "foo bar") + "foo bar") + ("" + "foo bar"))) + "a") {
    case "foo bar":
        var v17 uint = 2
        _ = v17
    case "a":
        v2 -= v2
    default:
        if true {
            fmt.Println(v3, v3)
            for fi18 := 0; fi18 < 6; fi18++ {
                v19 := GenS1{F0: (fi18 / (((-(fi18))) | 1)), F1: 1}
                var v20 uint16 = 841
                for fi21 := 0; fi21 < 13; fi21++ {
                    v22 := []int{}
                    var v23 float32 = 1e10
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        } else {
            for fi24 := 0; fi24 < 3; fi24++ {
                v25 := GenS3{F0: 0, F1: 0, F2: 1}
                v26 := GenS4{F0: 1, F1: 897}
                _ = fi24
                _ = v25
                _ = v26
            }
            v2 += (v2 + (v2 + v2))
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 string = "a"
                _ = fi27
                _ = v28
            }
        }
    }
    v29 := []int{42, 128, 1}
    for fi30 := 0; fi30 < 18; fi30++ {
        v3 = (v3 ^ v3)
        v2 = v2
        var v31 int16 = -1
        _ = fi30
        _ = v31
    }
    v1[-972] = -2147483648
    var v32 int32 = 0
    _ = len(v1)
    var v33 bool = false
    var v34 uint32 = 0
    v1[9223372036854775807] = 646
    for fi35 := 0; fi35 < 2; fi35++ {
        for fi36 := 0; fi36 < 11; fi36++ {
            _ = len(v29)
            _ = fi36
        }
        var v37 uint8 = 130
        var v38 float32 = (3.14159 / 1e10)
        _ = fi35
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        fmt.Println(v2)
        fmt.Println(v34, v2, v2)
        _ = fi39
    }
    var v40 float64 = 1e10
    var v41 uint16 = 255
    var v42 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v40
    _ = v41
    _ = v42
}
