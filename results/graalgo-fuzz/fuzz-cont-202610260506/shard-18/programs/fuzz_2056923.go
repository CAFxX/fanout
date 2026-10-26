package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 bool; F2 uint64 }
func genFn0(p0 int8, p1 uint8) int {
    v7 := []int{-72, 0, 0}
    v7 = append(v7, 9223372036854775807)
    for fi8 := 0; fi8 < 2; fi8++ {
        _ = len(v7)
        _ = fi8
    }
    p0 -= p0
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v7
    return 127
}
func genFn1(p0 uint64, p1 int64, p2 int8) uint16 {
    var v9 int8 = 100
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            v9 = (v9 & p2)
            var v12 float64 = (1.5 / 0.1)
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 uint8 = 1
            v15 := GenS1{F0: (p1 % ((p1) | 1)), F1: (v14 | (v14 | v14)), F2: 42}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi10
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := []int{(fi16 ^ -366), fi17, fi17}
            var v19 uint64 = p0
            _ = fi17
            _ = v18
            _ = v19
        }
        if true {
            switch "\n\t" {
            case "a":
                for fi20 := 0; fi20 < 3; fi20++ {
                    v21 := []int{(-(fi16)), fi16, (^(fi16)), (^(fi16))}
                    fmt.Println(fi16, fi16, p1)
                    if true {
                        fmt.Println(v9)
                        var v22 float64 = -2.25
                        v23 := GenS2{F0: p0}
                        v21 = append(v21, (^(fi20)))
                        _ = v22
                        _ = v23
                    } else {
                        v24 := GenS2{F0: p0}
                        if (!true) {
                            p1 += (p1 / ((int64(v9)) | 1))
                            v25 := GenS2{F0: 337}
                            v26 := GenS2{F0: (p0 << ((918) & 63))}
                            _ = v25
                            _ = v26
                        }
                        _ = v24
                    }
                    _ = fi20
                    _ = v21
                }
            case "hello":
                v27 := []int{(fi16 | genFn0((^(p2)), 255)), (-(fi16)), fi16, 32768}
                _ = v27
            }
        }
        _ = fi16
    }
    v28 := []int{-571, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v9
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 160
    var v3 int16 = -653
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]int)
    for fi29 := 0; fi29 < 2; fi29++ {
        fmt.Println(fi29, v2, v2)
        v30 := []int{4294967295, v2, fi29}
        var v31 float32 = ((0.1 - 3.14159) / 1.0)
        _ = fi29
        _ = v30
        _ = v31
    }
    v32 := []int{v2, v2, genFn0(100, 42), v2}
    if (65535 < 2) {
        fmt.Println(v2, v3)
        v2 -= (v2 / (((v2 % ((127) | 1))) | 1))
    } else {
        switch "foo bar" {
        case "\n\t":
            var v33 int64 = -472
            _ = v33
        default:
            v2 -= (v2 >> ((2) & 63))
        }
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        v2 = fi34
        if ("foo bar" != "foo bar") {
            switch (("\n\t" + "\n\t") + "\n\t") {
            case "hello":
                var v35 int64 = 9223372036854775807
                _ = v35
            case "":
                v36 := GenS2{F0: 1}
                _ = v36
            default:
                fmt.Println(v3)
            }
        }
        fmt.Println(v3, v2, v3)
        _ = fi34
    }
    switch -2147483648 {
    case -624:
        var v37 uint8 = 208
        _ = v37
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        v39 := GenS3{F0: v2, F1: (false && false), F2: 4294967295}
        _ = fi38
        _ = v39
    }
    fmt.Println(v2)
    var v40 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
