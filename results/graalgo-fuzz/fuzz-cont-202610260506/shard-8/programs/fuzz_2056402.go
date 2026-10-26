package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int16; F1 uint64; F2 int16 }
func genFn0(p0 uint32, p1 string) uint64 {
    fmt.Println(p1, p1)
    if (18446744073709551615 != 2) {
        p1 = ("" + p1)
    } else {
        for fi6 := 0; fi6 < 7; fi6++ {
            if true {
                v7 := GenS1{F0: 0}
                p0 = p0
                _ = v7
            }
            switch 32767 {
            case 534:
                v8 := []int{}
                _ = v8
            case 32768:
                if (!(-108 < 0)) {
                    v9 := GenS2{F0: -9223372036854775808}
                    fmt.Println(p1)
                    p1 = p1
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 11; fi10++ {
                        switch p1 {
                        case "hello":
                            p0 = 2147483647
                            fmt.Println(p0)
                        case "":
                            var v11 int = -123
                            var v12 int16 = -1
                            _ = v11
                            _ = v12
                        }
                        v13 := GenS3{F0: 2147483648, F1: -9223372036854775808}
                        _ = fi10
                        _ = v13
                    }
                    p1 = p1
                    var v14 int8 = 100
                    _ = v14
                }
            }
            if (168 != 246) {
                fmt.Println(p0)
                if (0.0 != 3.14159) {
                    var v15 uint32 = (p0 >> ((255) & 63))
                    var v16 int = (^(fi6))
                    _ = v15
                    _ = v16
                }
                v17 := []int{(^(fi6)), fi6, 42, fi6}
                _ = v17
            }
            _ = fi6
        }
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v19 := []int{-2147483648, fi18}
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = v20
    return 18446744073709551615
}
func genFn1(p0 uint, p1 uint16, p2 uint32) int16 {
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 15; fi22++ {
            v23 := GenS4{F0: (p2 * (p2 + (^(p2)))), F1: -9223372036854775808}
            p2 = (p2 / ((p2) | 1))
            v24 := GenS2{F0: 206}
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 uint = (^(p0))
        _ = fi21
        _ = v25
    }
    v26 := []int{1}
    v27 := GenS5{F0: 1, F1: 0, F2: 0}
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 uint64 = 2
        _ = fi28
        _ = v29
    }
    var v30 uint16 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    _ = v30
    return -1
}

func main() {
    v1 := []uint32{65535, 0, 65535}
    v2 := make(map[int]int)
    var v3 bool = (668 < 0)
    var v4 uint = 726
    var v5 uint64 = 42
    fmt.Println(v4, v3, v3)
    v5 += v5
    v31 := GenS4{F0: 662, F1: -22}
    for fi32 := 0; fi32 < 1; fi32++ {
        var v33 uint8 = 50
        v34 := GenS1{F0: (v4 / (((v4 * (v4 >> (((v4 | (v4 * 65535))) & 63)))) | 1))}
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 uint = uint(v5)
    fmt.Println(v5)
    var v36 int32 = -2147483648
    var v37 uint8 = 229
    _ = len(v2)
    v37 += v37
    for fi38 := 0; fi38 < 8; fi38++ {
        for fi39 := 0; fi39 < 11; fi39++ {
            v35 += v4
            _ = fi39
        }
        if v3 {
            fmt.Println(v5, v36)
        } else {
            for fi40 := 0; fi40 < 15; fi40++ {
                v41 := GenS2{F0: 0}
                _ = fi40
                _ = v41
            }
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        }
        fmt.Println(v5, v37, v36)
        _ = fi38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
