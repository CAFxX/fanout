package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
func genFn0(p0 uint32) int64 {
    p0 = 2147483647
    v5 := GenS1{F0: 127}
    v6 := GenS1{F0: 1}
    _ = p0
    _ = v5
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 uint64, p1 uint32, p2 bool) int64 {
    v7 := []int{}
    if p2 {
        fmt.Println(p2, p2, p1)
        for fi8 := 0; fi8 < 7; fi8++ {
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 uint8 = 42
                fmt.Println(p2, p2, fi9)
                v10 -= (v10 ^ (v10 & v10))
                switch "\n\t" {
                case "foo bar":
                    var v11 int16 = 538
                    _ = v11
                }
                var v12 int = 296
                _ = fi9
                _ = v10
                _ = v12
            }
            _ = fi8
        }
    }
    if p2 {
        p0 = 89
        p0 = p0
        fmt.Println(p1, p1, p0)
    } else {
        switch -1 {
        case 127:
            switch "" {
            case "":
                var v13 uint = 105
                _ = v13
            case "\n\t":
                if ((3.14159 - ((1.5 - -2.25) * (((float32(((0.1 + 1.5) + 1e10)) / ((1.5 - (float32(1e10) * 3.14159)) + 0.0)) * 0.1) / 1.0))) > 0.0) {
                    var v14 int32 = -842
                    _ = v14
                } else {
                    fmt.Println(p0, p1, p2)
                    fmt.Println(p1, p2, p2)
                }
            case "a":
                var v15 float32 = 0.0
                _ = v15
            }
        case -9223372036854775808:
            switch "foo bar" {
            case "hello":
                p1 += 778
            }
        }
        p2 = p2
    }
    var v16 string = (("\n\t" + ("foo bar" + "")) + "a")
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v7
    return -1
}

func main() {
    var v1 int8 = 0
    var v2 int32 = 2147483647
    v3 := []uint32{413, 1, 65535, 1}
    v4 := make(map[int]int)
    var v17 bool = (false || false)
    v4[128] = -836
    var v18 uint16 = 255
    fmt.Println(v2)
    v3 = append(v3, 247)
    v19 := GenS1{F0: (v1 >> ((128) & 63))}
    v4[128] = 965
    if (true && v17) {
        var v20 float64 = -2.25
        fmt.Println(v18)
        switch "hello" {
        case "":
            fmt.Println(v17, v17)
        case "hello":
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 uint64 = 2
                if v17 {
                    if v17 {
                        delete(v4, fi21)
                        _ = len(v3)
                        switch (("hello" + "hello") + "") {
                        case "foo bar":
                            fmt.Println(fi21)
                        case "a":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            v23 := GenS1{F0: int8(fi21)}
                            _ = v23
                        }
                    }
                    switch "\n\t" {
                    case "hello":
                        var v24 uint16 = (^(v18))
                        _ = v24
                    case "foo bar":
                        if v17 {
                            if (31 == 255) {
                                var v25 uint16 = v18
                                switch "\n\t" {
                                case "\n\t":
                                    switch (("" + "a") + "") {
                                    case "":
                                        fmt.Println(v18, v25)
                                    case "foo bar":
                                        v26 := GenS2{F0: (-(v1))}
                                        _ = v26
                                    }
                                }
                                var v27 uint16 = v18
                                _ = v25
                                _ = v27
                            }
                            v28 := GenS3{F0: uint(v18)}
                            _ = v28
                        } else {
                            var v29 string = "a"
                            if (1000 <= -419) {
                                if ((^(fi21)) < (fi21 | (fi21 + (fi21 - 1)))) {
                                    v2 -= -1
                                    var v30 int = fi21
                                    _ = v30
                                }
                            }
                            _ = v29
                        }
                    case "a":
                        v31 := GenS4{F0: (!((v17 && ((v2 & v2) == v2)) && true))}
                        _ = v31
                    }
                }
                _ = fi21
                _ = v22
            }
        default:
            fmt.Println(v2)
        }
        _ = v20
    }
    _ = len(v4)
    var v32 bool = v17
    var v33 int8 = (v1 | 0)
    v18 += v18
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
