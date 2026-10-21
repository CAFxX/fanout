package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
func genFn0() int32 {
    v7 := []int{-2147483648, 1, 4294967295, 98}
    switch 757 {
    case -629:
        var v8 bool = false
        _ = v8
    case 127:
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: 128, F1: fi9, F2: 42}
            var v11 float32 = (-2.25 * (((float32(-2.25) - 0.0) - 1e10) * 1e10))
            _ = fi9
            _ = v10
            _ = v11
        }
    case 128:
        v12 := GenS2{F0: 345}
        _ = v12
    }
    var v13 uint8 = 2
    v14 := GenS2{F0: 0}
    _ = v13
    _ = v14
    _ = v7
    return -1
}
func genFn1(p0 uint, p1 int64, p2 int64) uint64 {
    fmt.Println(p0, p2, p1)
    if (42 < 1) {
        switch "hello" {
        case "hello":
            v15 := []int{2147483647, 4294967295}
            _ = v15
        case "a":
            var v16 int8 = -128
            _ = v16
        case "foo bar":
            var v17 uint64 = 18446744073709551615
            _ = v17
        default:
            v18 := GenS3{F0: 1, F1: 91, F2: 816}
            _ = v18
        }
        p0 -= (p0 - 65535)
        p0 -= p0
    } else {
        p1 = (-(p2))
        for fi19 := 0; fi19 < 19; fi19++ {
            p0 = p0
            p2 -= 42
            p2 = p2
            _ = fi19
        }
        if true {
            if (0.1 < 3.14159) {
                p1 = (p1 - p2)
                var v20 uint = (^(p0))
                _ = v20
            }
            v21 := []int{733, 127, 4294967295}
            p2 = (^(p2))
            _ = v21
        } else {
            v22 := GenS2{F0: 541}
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 int16 = -32768
                _ = fi23
                _ = v24
            }
            switch "" {
            case "":
                var v25 bool = true
                _ = v25
            default:
                switch 4294967295 {
                case 9223372036854775807:
                    fmt.Println(p0)
                case -2147483648:
                    var v26 uint8 = 82
                    _ = v26
                }
            }
            _ = v22
        }
    }
    var v27 uint = 197
    fmt.Println(p0, p1, v27)
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 471
}

func main() {
    var v1 bool = (!((-32768 <= 1000) || (1 > 255)))
    var v2 int64 = -956
    var v3 uint = 12
    v4 := make(map[int]string)
    var v5 uint8 = 42
    var v6 string = "hello"
    var v28 int = -9223372036854775808
    var v29 float64 = (0.1 - float64(0.1))
    v30 := []int{}
    if len(v30) > 0 {
        v30[0] = v28
    }
    if len(v30) > 0 {
        v30[0] = (-(v28))
    }
    v1 = v1
    for fi31 := 0; fi31 < 19; fi31++ {
        if v1 {
            var v32 uint32 = 2
            _ = v32
        }
        if false {
            v2 -= -1
            v33 := GenS1{F0: 127, F1: (fi31 % ((fi31) | 1)), F2: (v2 ^ v2)}
            fmt.Println(v5)
            _ = v33
        } else {
            var v34 float32 = -2.25
            var v35 float64 = ((float64(v34) + v29) / v29)
            _ = v34
            _ = v35
        }
        switch (v28 ^ v28) {
        case -9223372036854775808:
            var v36 uint = (^(v3))
            _ = v36
        case 9223372036854775807:
            v1 = (("foo bar" < "hello") || v1)
        }
        _ = fi31
    }
    if v1 {
        v37 := v4[v28]
        for fi38 := 0; fi38 < 7; fi38++ {
            var v39 int32 = 42
            _ = fi38
            _ = v39
        }
        v29 += (1.5 + (1e10 * ((v29 - (-2.25 - 1.5)) + 0.1)))
        _ = v37
    } else {
        var v40 float32 = 1e10
        v41 := GenS3{F0: 1, F1: v5, F2: 18446744073709551615}
        fmt.Println(v6, v2)
        _ = v40
        _ = v41
    }
    v42 := GenS4{F0: (v28 % ((v28) | 1))}
    for fi43 := 0; fi43 < 11; fi43++ {
        var v44 uint16 = 220
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 10; fi45++ {
        v30 = append(v30, v28)
        _ = fi45
    }
    if ((!v1) && (2 < 0)) {
        var v46 int16 = -1
        _ = v46
    } else {
        if v1 {
            v2 = (v2 | -207)
            var v47 int32 = -1
            v48 := GenS5{F0: (v2 - (v2 / (((v2 * v2)) | 1)))}
            _ = v47
            _ = v48
        } else {
            v4[(-(v28))] = ((v6 + "") + "")
        }
    }
    var v49 int32 = 42
    fmt.Println(v2, v28, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v42
    _ = v49
    _ = v5
    _ = v6
}
