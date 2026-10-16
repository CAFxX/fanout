package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
func genFn0(p0 string, p1 string, p2 uint8) int64 {
    var v6 float64 = 1e10
    if false {
        v7 := GenS1{F0: "foo bar", F1: 42}
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        if (!false) {
            v6 = v6
            fmt.Println(p2, p1)
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            fmt.Println(p1)
            var v10 uint64 = 4294967295
            fmt.Println(fi9, fi8)
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: "hello", F1: 1}
        _ = fi8
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        p1 = "\n\t"
        p0 = p1
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 uint16) int16 {
    fmt.Println(p0)
    switch "" {
    case "\n\t":
        for fi13 := 0; fi13 < 11; fi13++ {
            if false {
                for fi14 := 0; fi14 < 2; fi14++ {
                    switch (fi14 / ((fi13) | 1)) {
                    case 9223372036854775807:
                        fmt.Println(p0, fi13, p0)
                    default:
                        switch "hello" {
                        case "\n\t":
                            v15 := GenS2{F0: 42}
                            _ = v15
                        }
                    }
                    if false {
                        p0 = 642
                        switch "" {
                        case "\n\t":
                            var v16 float64 = 0.0
                            _ = v16
                        }
                    } else {
                        p0 = (p0 ^ 989)
                        p0 += p0
                        p0 += (p0 / ((p0) | 1))
                        var v17 uint16 = p0
                        fmt.Println(p0, fi14, fi13)
                        _ = v17
                    }
                    fmt.Println(fi14)
                    var v18 uint64 = 42
                    _ = fi14
                    _ = v18
                }
            }
            switch -2147483648 {
            case -9223372036854775808:
                v19 := GenS3{F0: 0}
                _ = v19
            default:
                v20 := GenS4{F0: (p0 >> ((0) & 63))}
                _ = v20
            }
            _ = fi13
        }
    default:
        var v21 int32 = 1
        _ = v21
    }
    _ = p0
    return 32767
}

func main() {
    var v1 float32 = -2.25
    var v2 int8 = 127
    var v3 float64 = -2.25
    var v4 bool = (!true)
    v5 := []int{}
    var v22 float64 = 0.1
    var v23 uint8 = 20
    v1 = v1
    v2 = 100
    v1 = v1
    v24 := GenS5{F0: 2147483647}
    var v25 string = ("" + "a")
    v4 = v4
    fmt.Println(v4)
    fmt.Println(v3, v25)
    var v26 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
