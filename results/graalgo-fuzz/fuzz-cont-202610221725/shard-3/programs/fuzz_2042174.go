package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint32, p1 uint8) uint {
    var v7 float64 = -2.25
    var v8 uint8 = 82
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint64 = 1
        switch "\n\t" {
        case "a":
            v10 = v10
        }
        _ = fi9
        _ = v10
    }
    if false {
        fmt.Println(p0, v7, v8)
        var v11 int8 = 127
        v12 := []int{}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 uint32, p1 int32) int32 {
    p1 = p1
    switch "foo bar" {
    case "hello":
        fmt.Println(p0, p1)
    case "":
        p0 = p0
    }
    v13 := GenS1{F0: p1}
    switch "a" {
    case "hello":
        fmt.Println(p0)
    case "\n\t":
        var v14 float64 = 3.14159
        _ = v14
    case "a":
        p0 += (^(p0))
    default:
        if (166 >= 304) {
            var v15 float32 = float32((0.0 / (1.5 + -2.25)))
            _ = v15
        } else {
            for fi16 := 0; fi16 < 8; fi16++ {
                p1 = (p1 / ((p1) | 1))
                _ = fi16
            }
            p1 = 2147483647
        }
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v13
    return (^(p1))
}

func main() {
    var v1 int32 = 667
    v2 := make(map[uint32]int)
    var v3 uint8 = 2
    v4 := make(map[int]bool)
    var v5 uint32 = 31
    v6 := make(map[string]uint32)
    if false {
        fmt.Println(v5, v1)
    } else {
        var v17 string = "a"
        switch v17 {
        case "":
            v5 += 861
        case "foo bar":
            v18 := []int{127}
            _ = v18
        }
        _ = v17
    }
    v5 = v5
    v3 -= (v3 / (((v3 >> ((629) & 63))) | 1))
    for fi19 := 0; fi19 < 15; fi19++ {
        v5 += v5
        _ = fi19
    }
    v20 := []int{}
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            _ = len(v20)
            if false {
                v5 = 405
            } else {
                v23 := GenS1{F0: v1}
                _ = v23
            }
            _ = fi22
        }
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = fi21
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        fmt.Println(fi24, v1, v1)
        v1 += v1
        for fi25 := 0; fi25 < 19; fi25++ {
            switch "hello" {
            case "foo bar":
                var v26 uint32 = 65535
                _ = v26
            case "\n\t":
                var v27 int64 = 9223372036854775807
                _ = v27
            case "":
                fmt.Println(v3)
            }
            var v28 uint16 = 42
            _ = fi25
            _ = v28
        }
        _ = fi24
    }
    var v29 int8 = 42
    for fi30 := 0; fi30 < 12; fi30++ {
        var v31 float64 = -2.25
        _ = fi30
        _ = v31
    }
    var v32 int = -9223372036854775808
    if true {
        v33 := GenS1{F0: int32(v32)}
        var v34 int64 = 0
        _ = v33
        _ = v34
    }
    v29 = (v29 / ((v29) | 1))
    for fi35 := 0; fi35 < 16; fi35++ {
        var v36 uint = 176
        _ = fi35
        _ = v36
    }
    switch ("" + "foo bar") {
    case "a":
        v37 := GenS2{F0: 1, F1: (v29 << ((131) & 63))}
        _ = v37
    }
    v38 := GenS1{F0: v1}
    v29 -= 0
    for fi39 := 0; fi39 < 4; fi39++ {
        v40 := GenS1{F0: 42}
        if (!true) {
            v41 := GenS1{F0: v1}
            v1 = int32(v29)
            for fi42 := 0; fi42 < 7; fi42++ {
                v1 -= (-(v1))
                _ = fi42
            }
            _ = v41
        }
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
