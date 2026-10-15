package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint8 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 string; F1 uint16; F2 int16 }
func genFn0(p0 uint, p1 int, p2 bool) uint16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(fi7, fi7, p0)
        _ = fi7
    }
    if p2 {
        if p2 {
            for fi8 := 0; fi8 < 18; fi8++ {
                p2 = true
                v9 := []int{}
                _ = len(v9)
                _ = fi8
                _ = v9
            }
            v10 := []int{}
            _ = v10
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            p0 = p0
            fmt.Println(p0)
            _ = fi11
        }
    }
    var v12 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 1
}
func genFn1(p0 bool, p1 string) uint {
    var v13 uint64 = 18446744073709551615
    v14 := GenS1{F0: 10, F1: 1000, F2: 127}
    p1 = p1
    var v15 uint16 = 901
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := []int{fi16}
        fmt.Println(fi16, fi16, v13)
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return 0
}

func main() {
    v1 := []uint32{2147483647, 2, 226, 2}
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 uint32 = 1
    var v5 uint64 = 2
    var v6 uint = uint(v5)
    if len(v1) > 0 {
        v1[0] = v4
    }
    switch "foo bar" {
    case "a":
        v5 -= v5
    case "hello":
        v18 := v3[-9223372036854775808]
        _ = v18
    case "foo bar":
        var v19 uint8 = uint8(v4)
        _ = v19
    default:
        v5 -= v5
    }
    fmt.Println(v5)
    if false {
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        var v20 int32 = 42
        v5 = v5
        _ = v20
    } else {
        fmt.Println(v4)
    }
    var v21 uint64 = (v5 | 475)
    fmt.Println(v21, v5, v21)
    v22 := GenS2{F0: 9223372036854775807, F1: true}
    switch "hello" {
    case "foo bar":
        var v23 float32 = 3.14159
        _ = v23
    default:
        var v24 uint32 = v4
        _ = v24
    }
    v25 := GenS3{F0: "\n\t", F1: 65535, F2: 32767}
    var v26 bool = ((!false) && (true && false))
    for fi27 := 0; fi27 < 9; fi27++ {
        if (-608 == 32767) {
            if (!(!v26)) {
                switch "" {
                case "":
                    if v26 {
                        fmt.Println(v26, v26, v26)
                    }
                case "foo bar":
                    v28 := v3[127]
                    _ = v28
                case "hello":
                    switch ("" + "hello") {
                    case "foo bar":
                        fmt.Println(fi27)
                    case "\n\t":
                        v5 -= (v21 % ((v21) | 1))
                    default:
                        if true {
                            var v29 int8 = 1
                            var v30 bool = ((1.5 != 1.5) || v26)
                            for fi31 := 0; fi31 < 9; fi31++ {
                                var v32 bool = ((!(v30 && (!true))) || true)
                                _ = fi31
                                _ = v32
                            }
                            _ = v29
                            _ = v30
                        }
                    }
                default:
                    var v33 bool = v26
                    _ = v33
                }
                v2[1] = (fi27 + fi27)
            }
        }
        v34 := v2[(^(v4))]
        fmt.Println(v34, v21, v4)
        _ = fi27
        _ = v34
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v35 int32 = 42
    v36 := GenS2{F0: -1, F1: true}
    for fi37 := 0; fi37 < 14; fi37++ {
        var v38 uint = v6
        if false {
            var v39 float32 = -2.25
            if false {
                var v40 uint16 = 7
                _ = v40
            } else {
                v3[fi37] = false
            }
            _ = v39
        }
        _ = fi37
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
