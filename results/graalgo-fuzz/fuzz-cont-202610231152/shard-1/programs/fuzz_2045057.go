package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint }
type GenS5 struct { F0 int32; F1 int64; F2 uint16 }

func main() {
    var v1 string = ("a" + "hello")
    v2 := make(map[uint32]bool)
    var v3 float64 = 1.5
    v4 := []string{}
    fmt.Println(v3)
    v1 = "foo bar"
    for fi5 := 0; fi5 < 16; fi5++ {
        v3 -= 0.0
        var v6 string = (v1 + v1)
        v1 = "hello"
        _ = fi5
        _ = v6
    }
    var v7 uint = 65535
    fmt.Println(v3, v3, v1)
    var v8 int8 = 125
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := GenS1{F0: (false && (1 >= -572))}
        switch v1 {
        case "\n\t":
            fmt.Println(v3)
        case "foo bar":
            fmt.Println(v7)
        case "a":
            v4 = append(v4, ((v1 + v1) + (v1 + "hello")))
        default:
            switch v1 {
            case "":
                v11 := GenS1{F0: ((true || (false || true)) && true)}
                _ = v11
            case "\n\t":
                fmt.Println(fi9, fi9, v7)
            case "hello":
                v7 = 4294967295
            }
        }
        v12 := GenS2{F0: 255, F1: (fi9 & fi9), F2: 100}
        _ = fi9
        _ = v10
        _ = v12
    }
    v7 = v7
    v4 = append(v4, v1)
    if false {
        _ = len(v2)
        v13 := GenS3{F0: (v8 ^ v8)}
        if (779 <= 330) {
            v3 += 1e10
            v14 := GenS4{F0: 0, F1: 9223372036854775807, F2: (v7 | (v7 << ((228) & 63)))}
            _ = v14
        } else {
            _ = len(v4)
        }
        _ = v13
    } else {
        v8 = (v8 - v8)
    }
    v15 := GenS1{F0: (-821 <= 0)}
    v8 += v8
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint32 = 65535
        v1 = ("" + v1)
        _ = fi16
        _ = v17
    }
    _ = len(v4)
    fmt.Println(v3)
    var v18 int64 = -630
    v7 += 42
    if false {
        for fi19 := 0; fi19 < 14; fi19++ {
            fmt.Println(v8, v7, v7)
            _ = fi19
        }
        if (-937 == 0) {
            var v20 int16 = 0
            _ = v20
        }
    } else {
        for fi21 := 0; fi21 < 17; fi21++ {
            switch v1 {
            case "":
                v8 -= (v8 >> ((v7) & 63))
            default:
                if true {
                    v22 := GenS5{F0: -2147483648, F1: (^(v18)), F2: 2}
                    for fi23 := 0; fi23 < 7; fi23++ {
                        if false {
                            if (42 <= 42) {
                                var v24 uint64 = 4294967295
                                if (float64((1e10 - -2.25)) < v3) {
                                    switch "" {
                                    case "":
                                        v8 += (v8 + v8)
                                    }
                                    _ = len(v4)
                                }
                                _ = v24
                            }
                            fmt.Println(fi23, v7)
                        }
                        _ = fi23
                    }
                    _ = v22
                }
            }
            if true {
                var v25 float32 = (3.14159 / -2.25)
                _ = v25
            }
            _ = fi21
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
