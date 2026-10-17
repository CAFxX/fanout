package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 int) int16 {
    fmt.Println(p0)
    switch "foo bar" {
    case "foo bar":
        if (!(91 <= -128)) {
            switch p0 {
            case 9223372036854775807:
                v7 := GenS1{F0: 1000, F1: -1}
                _ = v7
            case 228:
                var v8 int32 = 42
                _ = v8
            case 127:
                for fi9 := 0; fi9 < 10; fi9++ {
                    var v10 int64 = 42
                    if (-1 < 83) {
                        v11 := GenS1{F0: 1, F1: 2147483647}
                        _ = v11
                    } else {
                        fmt.Println(p0, v10, v10)
                        for fi12 := 0; fi12 < 19; fi12++ {
                            var v13 uint = 474
                            var v14 bool = (((v10 - 1) < (v10 / ((-1) | 1))) || ((42 != 1) && false))
                            fmt.Println(v13, v13, p0)
                            fmt.Println(fi9, p0, fi9)
                            if v14 {
                                var v15 uint16 = 25
                                _ = v15
                            } else {
                                v16 := GenS2{F0: "", F1: "\n\t"}
                                _ = v16
                            }
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                    }
                    _ = fi9
                    _ = v10
                }
            default:
                for fi17 := 0; fi17 < 19; fi17++ {
                    switch ("" + ("hello" + "\n\t")) {
                    case "hello":
                        v18 := []int{}
                        _ = v18
                    case "":
                        p0 -= (fi17 + -620)
                    default:
                        fmt.Println(fi17, p0)
                    }
                    var v19 uint8 = 1
                    switch (fi17 | -2147483648) {
                    case -706:
                        var v20 int8 = 0
                        _ = v20
                    }
                    _ = fi17
                    _ = v19
                }
            }
        }
    case "":
        for fi21 := 0; fi21 < 8; fi21++ {
            fmt.Println(fi21)
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 float64 = (0.0 / (0.1 + 1e10))
                v24 := GenS3{F0: 1, F1: ""}
                var v25 int64 = -531
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi21
        }
    case "a":
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 bool = (255 != 0)
            fmt.Println(fi26)
            _ = fi26
            _ = v27
        }
    default:
        if false {
            if (1.5 < (3.14159 - (0.0 - 3.14159))) {
                p0 = (p0 ^ p0)
            }
        }
    }
    if (!(!false)) {
        fmt.Println(p0)
        p0 += p0
    }
    p0 = (p0 / ((-2147483648) | 1))
    _ = p0
    return 350
}
func genFn1(p0 int16, p1 uint32, p2 int32) bool {
    var v28 float64 = -2.25
    v29 := []int{-617, int(p0), 324}
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v29
    return (("a" + "hello") != "\n\t")
}

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    var v3 int32 = -2147483648
    v4 := make(map[string]int)
    v5 := make(map[int]bool)
    var v6 float64 = float64(0.0)
    var v30 int = 127
    for fi31 := 0; fi31 < 3; fi31++ {
        fmt.Println(v6, fi31, v3)
        _ = fi31
    }
    v6 += v6
    v32 := v2[int(v3)]
    v30 += v30
    for fi33 := 0; fi33 < 7; fi33++ {
        for fi34 := 0; fi34 < 11; fi34++ {
            v30 = fi34
            _ = fi34
        }
        fmt.Println(v30, v32)
        _ = fi33
    }
    var v35 bool = false
    var v36 float64 = 1e10
    v37 := v4[(v32 + (v32 + "hello"))]
    switch v37 {
    case 127:
        v37 = v30
    case 32767:
        fmt.Println(v35, v32, v35)
    default:
        switch "hello" {
        case "a":
            var v38 int8 = -1
            _ = v38
        case "":
            v30 -= 32767
        case "\n\t":
            for fi39 := 0; fi39 < 6; fi39++ {
                fmt.Println(v3)
                _ = fi39
            }
        }
    }
    switch v32 {
    case "":
        var v40 int32 = 77
        _ = v40
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        if v35 {
            fmt.Println(v32)
            var v41 uint8 = 28
            var v42 float64 = 0.0
            _ = v41
            _ = v42
        } else {
            _ = len(v4)
        }
    }
    fmt.Println(v32, v6, v36)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
