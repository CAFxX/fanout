package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 bool; F1 uint32 }
func genFn0(p0 bool, p1 uint64, p2 int) int {
    fmt.Println(p2, p0, p2)
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    return p2
}
func genFn1(p0 uint16, p1 uint64, p2 int64) uint {
    p2 = p2
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(p1, p1, fi6)
        fmt.Println(p2, p1, p2)
        _ = fi6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: p1, F1: 0}
        v9 := []int{236, fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: p1, F1: 42}
    switch "\n\t" {
    case "\n\t":
        for fi11 := 0; fi11 < 4; fi11++ {
            for fi12 := 0; fi12 < 4; fi12++ {
                switch 98 {
                case -9223372036854775808:
                    var v13 int8 = 0
                    _ = v13
                case 1:
                    if false {
                        fmt.Println(fi12, fi12, p1)
                    }
                case 32768:
                    v14 := GenS1{F0: 0, F1: 1}
                    if (("a" >= "a") || (-32768 == 0)) {
                        switch "hello" {
                        case "foo bar":
                            var v15 float32 = 0.1
                            _ = v15
                        case "":
                            if true {
                                var v16 int16 = 1000
                                p1 = (p1 / (((p1 / ((p1) | 1))) | 1))
                                p0 = 2
                                v16 = v16
                                _ = v16
                            }
                        case "a":
                            v17 := GenS2{F0: p0, F1: -1}
                            p1 -= p1
                            _ = v17
                        }
                        var v18 uint8 = 194
                        v18 = 170
                        _ = v18
                    }
                    switch fi11 {
                    case 99:
                        var v19 int8 = -1
                        _ = v19
                    case -447:
                        v20 := []int{}
                        _ = v20
                    case 32768:
                        var v21 int = (^(fi11))
                        _ = v21
                    }
                    _ = v14
                }
                var v22 int = fi12
                _ = fi12
                _ = v22
            }
            v23 := []int{4294967295, -2147483648, 32767}
            _ = fi11
            _ = v23
        }
    default:
        if true {
            v24 := GenS1{F0: p1, F1: 65535}
            p0 += p0
            _ = v24
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 18446744073709551615
    var v3 uint8 = 127
    var v4 bool = (!true)
    v5 := []string{}
    if v4 {
        v25 := GenS3{F0: -9223372036854775808}
        fmt.Println(v4, v2)
        fmt.Println(v3)
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        v2 = genFn1(0, 18446744073709551615, 0)
        v4 = v4
        _ = fi26
    }
    if (!false) {
        fmt.Println(v3)
    } else {
        fmt.Println(v4)
        v27 := GenS1{F0: uint64(v3), F1: 532}
        if (!v4) {
            var v28 float64 = (-2.25 - 0.0)
            v29 := v1[2]
            v5 = append(v5, "\n\t")
            _ = v28
            _ = v29
        } else {
            fmt.Println(v2, v3, v2)
            v30 := GenS4{F0: 65535, F1: 2}
            _ = v30
        }
        _ = v27
    }
    v5 = append(v5, "foo bar")
    v31 := GenS3{F0: 39}
    v3 -= 109
    if (2 < 46) {
        var v32 uint8 = 2
        _ = v32
    } else {
        for fi33 := 0; fi33 < 6; fi33++ {
            if true {
                var v34 string = "foo bar"
                _ = v34
            }
            v35 := GenS5{F0: false, F1: uint32(v3)}
            if v4 {
                var v36 float32 = 3.14159
                _ = v36
            }
            _ = fi33
            _ = v35
        }
    }
    var v37 uint = (v2 << (((v2 / (((v2 * 65535)) | 1))) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v37
    _ = v4
    _ = v5
}
