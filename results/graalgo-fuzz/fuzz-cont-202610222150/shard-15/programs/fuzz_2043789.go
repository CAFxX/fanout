package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 int8; F1 uint; F2 uint32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint8; F1 int; F2 uint32 }
type GenS6 struct { F0 int16; F1 int32; F2 int32 }
func genFn0(p0 uint64, p1 int64, p2 uint) int {
    var v4 string = "\n\t"
    v5 := GenS1{F0: 134, F1: 1, F2: 89}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 int, p1 uint64, p2 uint64) string {
    if (65535 < 761) {
        p0 = -1
    }
    var v6 bool = (584 < 4294967295)
    var v7 uint = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 1
    v3 := make(map[int]bool)
    if false {
        fmt.Println(v2)
        for fi8 := 0; fi8 < 7; fi8++ {
            v2 = (^(v2))
            v2 -= v2
            _ = fi8
        }
        switch "" {
        case "":
            v2 -= v2
        case "foo bar":
            v9 := GenS2{F0: 49, F1: 899, F2: 22}
            _ = v9
        }
    } else {
        v2 += (v2 * v2)
        v10 := GenS3{F0: 577, F1: 1}
        var v11 bool = (4294967295 >= 1)
        _ = v10
        _ = v11
    }
    fmt.Println(v2)
    v2 -= v2
    v2 += 127
    v12 := GenS4{F0: ((127 < 82) || false), F1: 65535}
    fmt.Println(v2)
    if false {
        if ("\n\t" != "foo bar") {
            var v13 bool = false
            fmt.Println(v2, v13)
            fmt.Println(v13, v13)
            _ = v13
        }
        var v14 float32 = 0.0
        if (186 < 0) {
            if false {
                if (false || (4294967295 < 2)) {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        fmt.Println(fi15)
                        v16 := []int{}
                        fmt.Println(fi15)
                        _ = fi15
                        _ = v16
                    }
                    var v17 int32 = -955
                    fmt.Println(v14, v14, v17)
                    _ = v17
                }
            } else {
                v14 += v14
                for fi18 := 0; fi18 < 3; fi18++ {
                    v19 := []int{}
                    _ = fi18
                    _ = v19
                }
            }
        } else {
            if true {
                switch -930 {
                case 676:
                    for fi20 := 0; fi20 < 19; fi20++ {
                        v21 := []int{genFn0(2, -631, 18446744073709551615), (fi20 >> ((243) & 63)), genFn0(332, 1, 50)}
                        _ = fi20
                        _ = v21
                    }
                case -1:
                    v22 := v1[42]
                    _ = v22
                case -84:
                    fmt.Println(v2, v2)
                default:
                    var v23 uint32 = 2147483647
                    _ = v23
                }
            } else {
                switch "\n\t" {
                case "hello":
                    v14 = v14
                case "a":
                    var v24 float32 = 1.5
                    _ = v24
                }
                v25 := GenS5{F0: v2, F1: -1, F2: 2147483648}
                var v26 uint = 2
                _ = v25
                _ = v26
            }
        }
        _ = v14
    }
    if true {
        for fi27 := 0; fi27 < 3; fi27++ {
            v2 -= v2
            for fi28 := 0; fi28 < 15; fi28++ {
                fmt.Println(v2)
                _ = fi28
            }
            v29 := GenS6{F0: 1000, F1: 42, F2: -2147483648}
            _ = fi27
            _ = v29
        }
    } else {
        for fi30 := 0; fi30 < 6; fi30++ {
            v2 = 1
            if false {
                var v31 int8 = 100
                delete(v1, 4294967295)
                _ = v31
            }
            v2 = 127
            _ = fi30
        }
        v2 = v2
    }
    var v32 bool = true
    fmt.Println(v2, v32)
    v33 := []int{32768, 32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
}
