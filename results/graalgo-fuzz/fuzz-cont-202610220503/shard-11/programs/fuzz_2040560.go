package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int; F1 bool; F2 int8 }
type GenS3 struct { F0 bool; F1 uint8; F2 int }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint8; F1 uint32 }
func genFn0(p0 string, p1 int32, p2 uint) uint16 {
    switch 0 {
    case 32768:
        var v6 float64 = 1e10
        _ = v6
    case 32767:
        var v7 bool = false
        _ = v7
    default:
        v8 := GenS1{F0: 413, F1: -1000}
        _ = v8
    }
    v9 := []int{128, -1, 0, 31}
    var v10 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 uint = 606
    v2 := make(map[int]uint32)
    var v3 int64 = 0
    v4 := make(map[uint32]string)
    var v5 int32 = 577
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS2{F0: (fi11 | fi11), F1: (128 == 152), F2: -1}
        for fi13 := 0; fi13 < 3; fi13++ {
            v1 = (v1 - v1)
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 uint = (^(v1))
            var v17 float32 = 0.1
            _ = fi15
            _ = v16
            _ = v17
        }
        if false {
            v3 = 9223372036854775807
            switch "" {
            case "":
                if false {
                    if (103 > 42) {
                        v1 += (v1 - 469)
                        v18 := GenS1{F0: 32767, F1: -480}
                        v2[-9223372036854775808] = 0
                        _ = v18
                    }
                    v19 := GenS2{F0: -9223372036854775808, F1: ("hello" < "foo bar"), F2: -1}
                    _ = v19
                } else {
                    switch fi14 {
                    case -2147483648:
                        v20 := GenS2{F0: (fi14 - -46), F1: (((-1 != fi14) || ((127 > 127) && false)) || true), F2: 0}
                        _ = v20
                    }
                    v21 := []int{819}
                    _ = v21
                }
            case "\n\t":
                fmt.Println(fi14, fi14, fi14)
            }
        } else {
            var v22 int32 = v5
            switch "\n\t" {
            case "":
                v22 = 1
            }
            switch ("\n\t" + "\n\t") {
            case "hello":
                v23 := GenS3{F0: false, F1: 0, F2: (fi14 * (fi14 * -231))}
                _ = v23
            case "\n\t":
                for fi24 := 0; fi24 < 4; fi24++ {
                    v25 := GenS4{F0: 42, F1: 42}
                    _ = fi24
                    _ = v25
                }
            case "":
                v3 = v3
            }
            _ = v22
        }
        v3 = v3
        _ = fi14
    }
    var v26 float32 = 0.1
    v27 := GenS3{F0: (162 < 2), F1: 1, F2: -9223372036854775808}
    var v28 int64 = v3
    v2[-1] = 0
    v29 := GenS5{F0: 2, F1: 2147483647}
    v30 := []int{-921}
    v28 = (v3 + v3)
    for fi31 := 0; fi31 < 4; fi31++ {
        fmt.Println(v1, v5, v1)
        delete(v2, 467)
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
