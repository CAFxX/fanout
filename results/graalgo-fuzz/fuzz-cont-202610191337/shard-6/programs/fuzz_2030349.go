package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint8 }
func genFn0(p0 int16, p1 int16) int32 {
    p0 = 813
    var v4 int8 = 127
    switch "a" {
    case "":
        switch 782 {
        case 32768:
            for fi5 := 0; fi5 < 13; fi5++ {
                if (1 < 816) {
                    fmt.Println(p1)
                } else {
                    v4 = v4
                    if ((v4 | (v4 + v4)) < -128) {
                        v4 += v4
                        if false {
                            v4 += (v4 ^ (^(v4)))
                        } else {
                            if true {
                                fmt.Println(p1, v4)
                                switch 32767 {
                                case -775:
                                    fmt.Println(fi5)
                                default:
                                    var v6 int16 = p1
                                    _ = v6
                                }
                                p0 = (p0 & p0)
                            }
                            fmt.Println(p0, p0, fi5)
                            fmt.Println(v4, v4, v4)
                        }
                    } else {
                        v7 := []int{}
                        _ = v7
                    }
                    if (280 >= 250) {
                        fmt.Println(v4, p1)
                    }
                }
                _ = fi5
            }
        case 32767:
            switch "" {
            case "\n\t":
                v4 += (v4 ^ (v4 * 42))
            default:
                p1 = (^(p1))
            }
        case -227:
            var v8 int = 2147483647
            _ = v8
        default:
            var v9 string = "foo bar"
            _ = v9
        }
    case "a":
        fmt.Println(v4, p1, v4)
    }
    v10 := []int{-833, 2147483647, 32768}
    var v11 uint64 = 2
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v4
    return 42
}
func genFn1(p0 int64, p1 string) bool {
    p0 = 1
    switch 2147483647 {
    case -426:
        var v12 uint = 0
        _ = v12
    case 1:
        if true {
            p1 = p1
        } else {
            v13 := GenS1{F0: 4294967295}
            if (true || false) {
                p1 = p1
            } else {
                switch p1 {
                case "hello":
                    var v14 int = 4294967295
                    _ = v14
                case "":
                    var v15 bool = false
                    _ = v15
                }
                for fi16 := 0; fi16 < 15; fi16++ {
                    switch "\n\t" {
                    case "hello":
                        v17 := GenS2{F0: (-(fi16)), F1: 64}
                        _ = v17
                    case "":
                        switch fi16 {
                        case 4294967295:
                            p1 = "\n\t"
                        case -1:
                            v18 := []int{(-(fi16)), fi16, (fi16 | (fi16 * fi16)), fi16}
                            _ = v18
                        case 32768:
                            for fi19 := 0; fi19 < 1; fi19++ {
                                var v20 uint64 = 18446744073709551615
                                v21 := []int{(fi16 << ((2) & 63))}
                                var v22 string = p1
                                _ = fi19
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        }
                    case "foo bar":
                        p0 -= p0
                    }
                    for fi23 := 0; fi23 < 19; fi23++ {
                        fmt.Println(fi23)
                        v24 := []int{}
                        _ = fi23
                        _ = v24
                    }
                    _ = fi16
                }
            }
            _ = v13
        }
    default:
        var v25 int32 = 255
        _ = v25
    }
    _ = p0
    _ = p1
    return ((true || (0 < 1)) && (false || false))
}

func main() {
    v1 := []string{"a"}
    v2 := make(map[int]bool)
    var v3 int16 = 197
    if (false && (!(!true))) {
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := GenS2{F0: fi26, F1: 24}
            _ = fi26
            _ = v27
        }
    } else {
        fmt.Println(v3)
        var v28 bool = true
        _ = v28
    }
    fmt.Println(v3)
    switch ("a" + ((((("" + (("\n\t" + "\n\t") + "foo bar")) + ("foo bar" + "\n\t")) + "a") + "") + "")) {
    case "hello":
        v3 -= -1
    case "foo bar":
        var v29 uint8 = 1
        _ = v29
    case "a":
        fmt.Println(v3)
    }
    var v30 bool = (!true)
    _ = len(v2)
    _ = len(v2)
    var v31 int = 0
    v1 = append(v1, (("hello" + "foo bar") + "a"))
    var v32 int16 = (v3 / (((^(v3))) | 1))
    v33 := v2[v31]
    if true {
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 uint32 = 0
            _ = fi34
            _ = v35
        }
        var v36 uint64 = 42
        switch "" {
        case "a":
            var v37 float32 = float32((0.1 * 1e10))
            _ = v37
        }
        _ = v36
    } else {
        for fi38 := 0; fi38 < 15; fi38++ {
            if v30 {
                v39 := GenS3{F0: 153, F1: 42, F2: 0}
                _ = v39
            }
            _ = fi38
        }
        fmt.Println(v30)
        _ = len(v1)
    }
    for fi40 := 0; fi40 < 9; fi40++ {
        _ = len(v2)
        var v41 float32 = float32(1.5)
        delete(v2, 128)
        _ = fi40
        _ = v41
    }
    v32 -= v3
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
