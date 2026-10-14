package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int8; F2 int8 }
func genFn0(p0 string, p1 int) int16 {
    var v7 int = p1
    if (true || (3.14159 <= 1.5)) {
        p0 = "a"
    } else {
        for fi8 := 0; fi8 < 15; fi8++ {
            switch (((("a" + (p0 + "a")) + p0) + (p0 + ((p0 + "foo bar") + ("" + p0)))) + (p0 + (p0 + (("" + p0) + p0)))) {
            case "foo bar":
                if (!true) {
                    switch "a" {
                    case "\n\t":
                        switch 9223372036854775807 {
                        case 0:
                            for fi9 := 0; fi9 < 20; fi9++ {
                                fmt.Println(p1, fi8, p0)
                                switch fi9 {
                                case 4294967295:
                                    v10 := GenS1{F0: true, F1: 2, F2: 867}
                                    _ = v10
                                case -672:
                                    v11 := []int{p1}
                                    _ = v11
                                default:
                                    if false {
                                        fmt.Println(v7, p0, p0)
                                        fmt.Println(p1, v7)
                                        if (p0 >= "a") {
                                            p1 = (-(p1))
                                            fmt.Println(fi9)
                                            var v12 int16 = -32768
                                            _ = v12
                                        } else {
                                            if (402 <= 124) {
                                                v13 := GenS1{F0: (true || true), F1: 0, F2: 2}
                                                v14 := []int{(fi8 | (fi8 + fi8)), -1, fi8, fi8}
                                                v15 := GenS1{F0: false, F1: 0, F2: 65535}
                                                _ = v13
                                                _ = v14
                                                _ = v15
                                            }
                                            v16 := GenS2{F0: -451}
                                            var v17 uint32 = 0
                                            _ = v16
                                            _ = v17
                                        }
                                    } else {
                                        fmt.Println(fi8)
                                        var v18 int8 = -62
                                        _ = v18
                                    }
                                }
                                v19 := GenS3{F0: -299, F1: 100, F2: -128}
                                _ = fi9
                                _ = v19
                            }
                        case 9223372036854775807:
                            fmt.Println(p0)
                        default:
                            v20 := []int{p1, (^(v7)), (p1 / (((p1 & v7)) | 1)), p1}
                            _ = v20
                        }
                    case "hello":
                        var v21 int16 = -709
                        _ = v21
                    case "a":
                        for fi22 := 0; fi22 < 11; fi22++ {
                            p0 = p0
                            v23 := GenS3{F0: 0, F1: 18, F2: 100}
                            _ = fi22
                            _ = v23
                        }
                    }
                    for fi24 := 0; fi24 < 6; fi24++ {
                        v7 -= 9223372036854775807
                        fmt.Println(v7)
                        fmt.Println(v7, v7)
                        _ = fi24
                    }
                    switch v7 {
                    case -9223372036854775808:
                        p1 += (fi8 | 32768)
                    }
                }
            }
            for fi25 := 0; fi25 < 12; fi25++ {
                if ((-1 == 700) || false) {
                    if (!false) {
                        fmt.Println(v7, fi8, v7)
                    } else {
                        v26 := []int{fi25, p1, p1, -699}
                        var v27 uint16 = 839
                        _ = v26
                        _ = v27
                    }
                }
                _ = fi25
            }
            p1 += -6
            _ = fi8
        }
        var v28 float64 = 1.5
        _ = v28
    }
    switch p0 {
    case "":
        v29 := GenS1{F0: (!(0.1 == 1.5)), F1: 42, F2: 4294967295}
        _ = v29
    }
    fmt.Println(v7, p0)
    _ = p0
    _ = p1
    _ = v7
    return 1000
}

func main() {
    var v1 uint8 = 1
    var v2 int = 4294967295
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]uint32)
    var v5 uint64 = 0
    var v6 string = "\n\t"
    var v30 uint = 18446744073709551615
    var v31 int8 = 0
    var v32 bool = (uint16(v5) <= 42)
    var v33 int = (v2 & (-(v2)))
    v34 := []int{}
    var v35 float64 = ((-2.25 - 0.1) + ((-2.25 / (((0.1 * 0.1) * 1e10) - float64(1.5))) / 1.0))
    if false {
        v34 = append(v34, v2)
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
    }
    var v36 int32 = int32(v31)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
