package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }
type GenS3 struct { F0 bool; F1 int32; F2 uint16 }
type GenS4 struct { F0 int64; F1 bool }
func genFn0(p0 int16, p1 uint16) int {
    v4 := []int{239, -178, 0}
    v5 := GenS1{F0: p1}
    if ((true || (1.5 <= ((0.1 / 1e10) / 1.5))) && (!true)) {
        for fi6 := 0; fi6 < 3; fi6++ {
            if len(v4) > 0 {
                v4[0] = fi6
            }
            switch 127 {
            case 32767:
                p0 += (p0 << ((42) & 63))
            case -9223372036854775808:
                var v7 float32 = 0.0
                _ = v7
            default:
                if true {
                    for fi8 := 0; fi8 < 3; fi8++ {
                        p1 -= p1
                        var v9 uint = 444
                        _ = fi8
                        _ = v9
                    }
                    if (!false) {
                        v10 := GenS1{F0: (^(p1))}
                        _ = v10
                    }
                } else {
                    switch ("a" + ("\n\t" + "hello")) {
                    case "hello":
                        switch (^(fi6)) {
                        case 42:
                            fmt.Println(p1, p0)
                        }
                    case "\n\t":
                        switch fi6 {
                        case 32768:
                            for fi11 := 0; fi11 < 5; fi11++ {
                                v12 := GenS1{F0: 0}
                                p0 += (-(p0))
                                v13 := GenS1{F0: 255}
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                        default:
                            var v14 uint64 = 42
                            _ = v14
                        }
                    case "":
                        p0 = -674
                    }
                    v4 = append(v4, fi6)
                }
            }
            fmt.Println(p1, p0)
            _ = fi6
        }
    }
    fmt.Println(p0, p0)
    switch 123 {
    case 42:
        for fi15 := 0; fi15 < 16; fi15++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi15
        }
    case 1:
        p0 -= p0
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 32768
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 4294967295
    v3 := make(map[string]bool)
    for fi16 := 0; fi16 < 5; fi16++ {
        switch "\n\t" {
        case "hello":
            var v17 float64 = (1.5 / float64(1e10))
            _ = v17
        default:
            fmt.Println(fi16)
        }
        v2 = v2
        _ = fi16
    }
    v18 := []int{42, 128}
    v19 := GenS2{F0: -876, F1: 0, F2: 9223372036854775807}
    var v20 string = "foo bar"
    v20 = "a"
    var v21 float64 = (3.14159 + -2.25)
    v22 := v1[v20]
    v23 := GenS3{F0: (true || true), F1: 2147483647, F2: 241}
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    var v24 int32 = -1
    v25 := GenS1{F0: 2}
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 uint32 = 42
        _ = fi26
        _ = v27
    }
    if true {
        var v28 float32 = (0.0 / 1.0)
        if len(v18) > 0 {
            v18[0] = 4294967295
        }
        _ = v28
    } else {
        v29 := GenS1{F0: 0}
        fmt.Println(v22, v20)
        var v30 int32 = (v24 & v24)
        _ = v29
        _ = v30
    }
    v31 := GenS4{F0: 0, F1: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
}
