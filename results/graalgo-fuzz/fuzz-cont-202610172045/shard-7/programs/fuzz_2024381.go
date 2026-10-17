package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 uint }
type GenS4 struct { F0 uint8; F1 int; F2 uint8 }
type GenS5 struct { F0 uint }
func genFn0(p0 int32, p1 uint, p2 uint16) uint8 {
    for fi4 := 0; fi4 < 11; fi4++ {
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 uint32 = 1
            var v7 int = fi5
            _ = fi5
            _ = v6
            _ = v7
        }
        var v8 int64 = -1
        _ = fi4
        _ = v8
    }
    if true {
        var v9 uint32 = uint32(p1)
        _ = v9
    } else {
        v10 := GenS1{F0: true, F1: 255, F2: 32768}
        _ = v10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        if (1 == (p0 - (p0 / ((p0) | 1)))) {
            var v12 uint32 = 1
            for fi13 := 0; fi13 < 12; fi13++ {
                fmt.Println(v12, v12, fi11)
                v14 := []int{42, fi13, fi13}
                _ = fi13
                _ = v14
            }
            _ = v12
        } else {
            switch 0 {
            case 32768:
                v15 := []int{(fi11 & 1)}
                _ = v15
            case -526:
                p2 += p2
            default:
                if false {
                    if false {
                        fmt.Println(p1, p1)
                        p0 = 0
                    }
                    switch fi11 {
                    case -1:
                        v16 := GenS1{F0: false, F1: 2, F2: (fi11 >> ((p1) & 63))}
                        _ = v16
                    case 994:
                        p0 -= (p0 - (p0 << ((2) & 63)))
                    }
                    var v17 int32 = (p0 * int32(p1))
                    _ = v17
                } else {
                    var v18 int8 = 100
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 11; fi19++ {
                fmt.Println(p1, fi19)
                _ = fi19
            }
            v20 := GenS2{F0: 972}
            _ = v20
        }
        v21 := []int{int(p2)}
        _ = fi11
        _ = v21
    }
    v22 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 128
}
func genFn1(p0 int32) int8 {
    if false {
        switch "hello" {
        case "\n\t":
            v23 := []int{127, 9223372036854775807, 32768}
            _ = v23
        case "a":
            v24 := GenS3{F0: (1 == p0), F1: 0, F2: 376}
            _ = v24
        }
        v25 := []int{-2147483648, -2147483648, 128, -567}
        p0 -= -1
        _ = v25
    }
    v26 := []int{}
    p0 = p0
    fmt.Println(p0)
    _ = p0
    _ = v26
    return -1
}

func main() {
    v1 := []string{"hello", "hello"}
    var v2 int16 = 1000
    v3 := make(map[string]int)
    if len(v1) > 0 {
        v1[0] = ""
    }
    v2 -= 0
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 int16 = v2
        var v29 uint64 = 0
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println(v2)
    v2 -= 998
    var v30 int8 = 42
    var v31 int64 = -1
    var v32 int = 515
    v33 := GenS4{F0: 2, F1: (v32 + v32), F2: genFn0(-1, 255, 42)}
    v2 += (v2 + v2)
    v34 := GenS3{F0: true, F1: v30, F2: 0}
    fmt.Println(v31, v31)
    v35 := GenS4{F0: 151, F1: (v32 & (v32 * v32)), F2: genFn0(0, 1, 32767)}
    v36 := GenS2{F0: 1}
    fmt.Println(v2)
    v37 := GenS5{F0: 18446744073709551615}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
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
    _ = v37
}
