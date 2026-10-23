package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 string, p1 int32) uint64 {
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := GenS1{F0: 4294967295, F1: -1}
        if (true && false) {
            fmt.Println(p0, p0, p0)
            var v6 int = int(p1)
            var v7 uint64 = 4294967295
            _ = v6
            _ = v7
        }
        p1 -= (p1 >> ((65535) & 63))
        _ = fi4
        _ = v5
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint = 791
        p1 = p1
        _ = fi8
        _ = v9
    }
    v10 := []int{}
    var v11 uint64 = 741
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return v11
}
func genFn1(p0 string, p1 int32, p2 int) uint64 {
    p1 = (p1 >> ((0) & 63))
    if true {
        v12 := []int{}
        p1 = p1
        _ = v12
    }
    switch p0 {
    case "":
        v13 := GenS1{F0: 996, F1: 42}
        _ = v13
    default:
        p2 -= (-(p2))
    }
    v14 := []int{}
    switch p0 {
    case "a":
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 int64 = 42
                var v18 int32 = 0
                if true {
                    var v19 uint16 = 255
                    p2 = 128
                    var v20 int = (fi16 | fi15)
                    _ = v19
                    _ = v20
                } else {
                    var v21 uint64 = 2
                    fmt.Println(p2, v21)
                    _ = v21
                }
                if true {
                    fmt.Println(p0)
                }
                switch fi15 {
                case 42:
                    _ = len(v14)
                case -14:
                    if (32767 >= (p2 & (^(p2)))) {
                        var v22 float32 = ((1e10 + 3.14159) + 1e10)
                        _ = len(v14)
                        switch 42 {
                        case 1:
                            var v23 uint16 = 255
                            _ = v23
                        case 32767:
                            var v24 int8 = 42
                            _ = v24
                        case 928:
                            var v25 bool = (v18 < -68)
                            _ = v25
                        default:
                            p0 = (p0 + p0)
                        }
                        v18 = p1
                        _ = v22
                    }
                case 32768:
                    v26 := GenS1{F0: genFn0(p0, 0), F1: 1}
                    _ = v26
                default:
                    fmt.Println(p1)
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi15
        }
    default:
        p0 = "hello"
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 1
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]uint32)
    var v3 uint32 = 1
    var v27 uint8 = 98
    v28 := []int{32767, 42, 886}
    var v29 uint32 = 60
    if (-2.25 >= 3.14159) {
        var v30 bool = false
        v1 += uint16(v3)
        v31 := GenS2{F0: "hello", F1: -9223372036854775808}
        _ = v30
        _ = v31
    } else {
        v1 = v1
    }
    var v32 int8 = 42
    var v33 int32 = 42
    v34 := GenS1{F0: 2, F1: 79}
    var v35 uint16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
