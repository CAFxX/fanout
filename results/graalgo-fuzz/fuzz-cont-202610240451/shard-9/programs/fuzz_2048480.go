package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 string; F2 int64 }
func genFn0() int {
    var v5 uint16 = 1
    switch "foo bar" {
    case "":
        v5 -= v5
    default:
        v5 = v5
    }
    v6 := []int{}
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 uint16) int32 {
    if false {
        switch genFn0() {
        case 128:
            switch "hello" {
            case "":
                var v7 uint16 = 255
                _ = v7
            case "a":
                for fi8 := 0; fi8 < 16; fi8++ {
                    if true {
                        v9 := GenS1{F0: -9223372036854775808}
                        v10 := []int{-681, fi8, fi8, (fi8 | (fi8 / ((fi8) | 1)))}
                        _ = v10
                        _ = v9
                    }
                    _ = fi8
                }
            default:
                if false {
                    p0 += (p0 >> ((255) & 63))
                }
            }
        }
        switch "\n\t" {
        case "hello":
            v11 := []int{794}
            _ = v11
        default:
            p0 = 65535
        }
        var v12 uint16 = (p0 & p0)
        _ = v12
    } else {
        for fi13 := 0; fi13 < 20; fi13++ {
            switch "" {
            case "\n\t":
                var v14 int64 = -9223372036854775808
                _ = v14
            case "a":
                var v15 uint32 = 0
                _ = v15
            case "":
                fmt.Println(fi13, fi13)
            }
            p0 = p0
            v16 := GenS1{F0: -1}
            _ = fi13
            _ = v16
        }
        var v17 uint8 = 232
        _ = v17
    }
    var v18 int = -2147483648
    p0 -= p0
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int8 = -124
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = v18
    return 982
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 1
    var v3 int32 = -1
    v4 := make(map[string]int)
    switch 32768 {
    case 2147483647:
        fmt.Println(v2, v3)
    case 0:
        v2 = v2
    case 13:
        if ((127 < 9223372036854775807) && (true || false)) {
            v3 += 2147483647
            if false {
                _ = len(v1)
                if (-1 != (v3 & 2147483647)) {
                    fmt.Println(v3, v3)
                    for fi21 := 0; fi21 < 14; fi21++ {
                        if true {
                            switch "" {
                            case "hello":
                                v2 = 88
                            case "foo bar":
                                v3 -= (v3 + genFn1(525))
                            }
                            fmt.Println(v3, v3, v3)
                            var v22 int8 = 127
                            _ = v22
                        } else {
                            var v23 bool = ((v2 ^ v2) >= v2)
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 14; fi24++ {
                            v1[2147483647] = (2 <= 42)
                            v25 := GenS2{F0: 42, F1: "hello", F2: -9223372036854775808}
                            _ = fi24
                            _ = v25
                        }
                        _ = fi21
                    }
                    var v26 int64 = 160
                    _ = v26
                }
            }
            v3 = -395
        } else {
            var v27 int32 = (-(v3))
            if false {
                v27 = v27
            }
            var v28 string = "foo bar"
            _ = v27
            _ = v28
        }
    }
    fmt.Println(v2, v3)
    switch genFn0() {
    case 32767:
        var v29 int16 = 812
        _ = v29
    }
    v30 := []int{9223372036854775807, int(v2)}
    var v31 uint16 = 255
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 float32 = 3.14159
        _ = fi32
        _ = v33
    }
    if ("" < "\n\t") {
        for fi34 := 0; fi34 < 16; fi34++ {
            var v35 uint = 18446744073709551615
            var v36 string = "\n\t"
            _ = fi34
            _ = v35
            _ = v36
        }
        v37 := GenS1{F0: 42}
        _ = v37
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
