package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint64; F1 bool; F2 bool }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int64; F1 int16; F2 uint64 }
func genFn0(p0 uint8, p1 int16, p2 uint64) int {
    var v4 bool = (p0 <= (p0 | 128))
    if false {
        switch "hello" {
        case "hello":
            switch "hello" {
            case "":
                if v4 {
                    var v5 int16 = 0
                    fmt.Println(v4, v4)
                    _ = v5
                } else {
                    var v6 uint64 = p2
                    v7 := GenS1{F0: 9223372036854775807, F1: false, F2: uint32(p0)}
                    _ = v6
                    _ = v7
                }
            case "hello":
                var v8 uint8 = (^(p0))
                _ = v8
            }
        default:
            v9 := GenS1{F0: 0, F1: ((p1 * (p1 % ((p1) | 1))) != (-(p1))), F2: 2}
            _ = v9
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            if v4 {
                v11 := GenS2{F0: p2, F1: (-1 == 114), F2: (43 <= 100)}
                v12 := []int{(fi10 << ((0) & 63)), 127, (fi10 & fi10)}
                v12 = append(v12, fi10)
                _ = v11
                _ = v12
            }
            fmt.Println(v4, v4, v4)
            _ = fi10
        }
    } else {
        v13 := []int{731}
        _ = v13
    }
    if (v4 || v4) {
        if ((^(p0)) >= 1) {
            for fi14 := 0; fi14 < 14; fi14++ {
                v4 = false
                _ = fi14
            }
        } else {
            if v4 {
                v15 := GenS3{F0: 42, F1: 100}
                v16 := []int{}
                _ = v15
                _ = v16
            }
            switch 483 {
            case 114:
                v17 := GenS3{F0: -1, F1: 1}
                _ = v17
            case 1:
                var v18 bool = (v4 && (v4 || ((p2 ^ p2) < (p2 << ((65535) & 63)))))
                _ = v18
            }
            var v19 float32 = 0.0
            _ = v19
        }
    }
    p0 += p0
    for fi20 := 0; fi20 < 1; fi20++ {
        fmt.Println(p0, p1, v4)
        _ = fi20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 32767
}
func genFn1(p0 int64, p1 uint32, p2 uint64) uint32 {
    var v21 uint = 4294967295
    if (!false) {
        for fi22 := 0; fi22 < 11; fi22++ {
            v23 := []int{genFn0(239, 674, (p2 + p2)), (fi22 - fi22)}
            v24 := GenS3{F0: 100, F1: 127}
            _ = fi22
            _ = v23
            _ = v24
        }
        v25 := []int{-9223372036854775808}
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return p1
}

func main() {
    var v1 string = "a"
    var v2 uint8 = 87
    var v3 int = -602
    v3 = v3
    for fi26 := 0; fi26 < 18; fi26++ {
        fmt.Println(fi26, v3)
        var v27 uint = 2
        _ = fi26
        _ = v27
    }
    var v28 uint32 = 1
    v3 = v3
    v28 = v28
    var v29 int64 = 190
    var v30 bool = true
    v31 := GenS3{F0: -20, F1: -128}
    v2 = v2
    v30 = v30
    if (2147483647 > v3) {
        v32 := GenS4{F0: 0, F1: -1}
        v33 := GenS5{F0: -339, F1: 1000, F2: 4294967295}
        v34 := []int{(^(v3))}
        _ = v32
        _ = v33
        _ = v34
    }
    v30 = (!(v30 || v30))
    fmt.Println(v3, v29, v1)
    if false {
        fmt.Println(v30, v1, v28)
    } else {
        v35 := GenS4{F0: 853, F1: 1000}
        v30 = false
        _ = v35
    }
    v3 += (v3 >> ((42) & 63))
    fmt.Println(v2, v3, v1)
    v36 := []int{(v3 ^ v3)}
    var v37 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
}
