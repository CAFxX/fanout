package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint8 }
func genFn0() string {
    if false {
        for fi6 := 0; fi6 < 14; fi6++ {
            for fi7 := 0; fi7 < 6; fi7++ {
                if true {
                    var v8 uint = 526
                    v9 := []int{238, 1}
                    v10 := GenS1{F0: 4294967295, F1: 42, F2: 32767}
                    _ = v10
                    _ = v8
                    _ = v9
                }
                if ("" < "a") {
                    var v11 int = fi7
                    v11 = 127
                    _ = v11
                }
                _ = fi7
            }
            var v12 int64 = -415
            v12 = (v12 * int64(fi6))
            _ = fi6
            _ = v12
        }
    }
    if false {
        var v13 float32 = 0.1
        v14 := []int{917, 32767, 914}
        v13 += 0.1
        _ = v13
        _ = v14
    } else {
        var v15 uint = 2
        var v16 int32 = -565
        _ = v15
        _ = v16
    }
    v17 := []int{1}
    var v18 int = -2147483648
    _ = v17
    _ = v18
    return "a"
}
func genFn1(p0 bool, p1 int8, p2 uint64) bool {
    switch "a" {
    case "hello":
        switch "foo bar" {
        case "\n\t":
            fmt.Println(p0)
        default:
            switch "foo bar" {
            case "a":
                v19 := []int{}
                _ = v19
            case "hello":
                p1 = -128
            case "\n\t":
                v20 := []int{}
                _ = v20
            }
        }
    default:
        fmt.Println(p1, p0)
    }
    p2 += (p2 + (^(p2)))
    _ = p0
    _ = p1
    _ = p2
    return (!(!(false || p0)))
}

func main() {
    v1 := []int{2147483647, -576}
    v2 := []string{}
    v3 := make(map[uint32]int)
    var v4 uint16 = 255
    var v5 int = -2147483648
    v4 += v4
    if (v5 > 9223372036854775807) {
        var v21 int64 = -9223372036854775808
        _ = v21
    } else {
        for fi22 := 0; fi22 < 20; fi22++ {
            if (false && true) {
                v5 = fi22
            }
            fmt.Println(fi22, v5, v4)
            if true {
                v23 := GenS1{F0: 323, F1: 360, F2: v4}
                switch ("a" + "") {
                case "\n\t":
                    v5 += fi22
                case "":
                    for fi24 := 0; fi24 < 5; fi24++ {
                        var v25 uint8 = 42
                        delete(v3, 0)
                        if ((2 <= (v25 + v25)) && (-254 == -1)) {
                            var v26 int32 = -147
                            v27 := GenS1{F0: 2, F1: int32(v25), F2: (v4 >> ((4294967295) & 63))}
                            _ = v26
                            _ = v27
                        }
                        fmt.Println(fi24)
                        _ = fi24
                        _ = v25
                    }
                default:
                    v28 := GenS2{F0: 128}
                    _ = v28
                }
                fmt.Println(v5)
                _ = v23
            }
            _ = fi22
        }
    }
    if true {
        switch "\n\t" {
        case "a":
            fmt.Println(v5, v5)
        default:
            v4 += 199
        }
        v29 := GenS1{F0: 1, F1: -639, F2: v4}
        fmt.Println(v4, v5)
        _ = v29
    } else {
        v30 := GenS3{F0: -1, F1: 0}
        fmt.Println(v4, v4)
        var v31 uint64 = 2
        _ = v30
        _ = v31
    }
    v5 += v5
    for fi32 := 0; fi32 < 11; fi32++ {
        fmt.Println(v4)
        _ = fi32
    }
    fmt.Println(v4, v4)
    fmt.Println(v4, v5)
    v4 = (v4 & 255)
    fmt.Println(v4, v5)
    switch v5 {
    case -9223372036854775808:
        switch ("" + "") {
        case "a":
            var v33 uint32 = 1
            _ = v33
        }
    case 32768:
        switch "a" {
        case "foo bar":
            var v34 int32 = -1
            _ = v34
        case "":
            v5 = (v5 / ((-686) | 1))
        default:
            fmt.Println(v4)
        }
    default:
        if len(v1) > 0 {
            v1[0] = v5
        }
    }
    _ = len(v3)
    fmt.Println(v5, v5)
    v3[2147483647] = (v5 ^ -1)
    var v35 int64 = -9223372036854775808
    if (true || ((-(v5)) < v5)) {
        var v36 float32 = 0.1
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
