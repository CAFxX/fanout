package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int64 }

func main() {
    var v1 int64 = 0
    v2 := []uint32{1, 1, 1, 0}
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    var v5 uint16 = 255
    v6 := GenS1{F0: v5, F1: (v5 / (((v5 & (v5 / ((v5) | 1)))) | 1)), F2: 4294967295}
    fmt.Println(v5, v1)
    _ = len(v2)
    var v7 int32 = 2147483647
    fmt.Println(v7, v5)
    var v8 uint32 = 0
    if false {
        fmt.Println(v7)
        var v9 uint32 = 2147483647
        if false {
            var v10 float64 = ((1.5 + 1e10) / (3.14159 + -2.25))
            var v11 uint16 = 255
            v10 = 3.14159
            _ = v10
            _ = v11
        }
        _ = v9
    }
    var v12 float64 = ((1e10 * 0.1) - 1e10)
    switch "\n\t" {
    case "":
        v13 := GenS1{F0: v5, F1: (v5 | v5), F2: 32767}
        _ = v13
    case "foo bar":
        if false {
            var v14 uint8 = 0
            v12 = v12
            var v15 int16 = -507
            _ = v14
            _ = v15
        } else {
            var v16 int8 = 127
            if (true || false) {
                if true {
                    v12 = 3.14159
                } else {
                    fmt.Println(v1)
                }
                var v17 bool = true
                if len(v2) > 0 {
                    v2[0] = v8
                }
                _ = v17
            }
            v5 -= v5
            _ = v16
        }
    case "\n\t":
        if (23 <= 32767) {
            var v18 int16 = 1
            if true {
                fmt.Println(v8, v1)
            } else {
                v19 := GenS1{F0: (v5 - (^(v5))), F1: 2, F2: -1}
                _ = v19
            }
            v12 -= v12
            _ = v18
        }
    }
    fmt.Println(v1, v12, v5)
    switch 0 {
    case 0:
        var v20 uint16 = (v5 % ((956) | 1))
        _ = v20
    case 32768:
        if true {
            v1 += (^(v1))
        }
    case 4294967295:
        v5 -= v5
    default:
        if true {
            v5 += v5
            switch 2147483647 {
            case 335:
                _ = len(v2)
            }
            v21 := GenS2{F0: 2}
            _ = v21
        } else {
            v8 -= uint32(v5)
        }
    }
    v12 -= v12
    v5 = (v5 / ((v5) | 1))
    v8 += 2147483648
    delete(v4, 733)
    fmt.Println(v7, v12, v1)
    switch "a" {
    case "a":
        for fi22 := 0; fi22 < 9; fi22++ {
            v8 -= v8
            _ = fi22
        }
    case "":
        var v23 int8 = 127
        _ = v23
    case "foo bar":
        switch 1 {
        case 127:
            v24 := GenS3{F0: 42}
            _ = v24
        case 1:
            if false {
                fmt.Println(v12, v8, v1)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v25 := GenS4{F0: 100, F1: (v1 + (v1 & (v1 * -303)))}
                _ = v25
            } else {
                for fi26 := 0; fi26 < 19; fi26++ {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v1 = v1
                        v7 = (^(v7))
                        v28 := GenS2{F0: 1}
                        _ = fi27
                        _ = v28
                    }
                    _ = fi26
                }
                v29 := GenS3{F0: 1000}
                _ = v29
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
