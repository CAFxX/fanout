package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint; F1 int64; F2 int64 }
func genFn0(p0 uint64) uint16 {
    switch "hello" {
    case "":
        v4 := []int{-2147483648, -734, 42}
        _ = v4
    case "foo bar":
        if false {
            if false {
                var v5 int32 = 2147483647
                _ = v5
            }
            fmt.Println(p0)
        } else {
            v6 := []int{-2147483648, 42, 820}
            p0 += p0
            for fi7 := 0; fi7 < 4; fi7++ {
                for fi8 := 0; fi8 < 12; fi8++ {
                    v9 := GenS1{F0: 1}
                    p0 = (p0 - (p0 / (((p0 * p0)) | 1)))
                    var v10 string = "foo bar"
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi7
            }
            _ = v6
        }
    }
    p0 = (p0 << ((42) & 63))
    var v11 float64 = 1e10
    _ = p0
    _ = v11
    return 42
}
func genFn1(p0 int64, p1 int16, p2 uint16) uint8 {
    var v12 uint16 = 842
    var v13 uint = 0
    if ((p1 / ((p1) | 1)) >= p1) {
        v13 = (v13 << ((v13) & 63))
    }
    switch 4294967295 {
    case 1:
        var v14 uint = v13
        _ = v14
    case 45:
        p0 += (-(p0))
    case 126:
        v15 := GenS1{F0: 2}
        _ = v15
    default:
        if (!false) {
            p1 += p1
        } else {
            if false {
                switch ("a" + "\n\t") {
                case "hello":
                    v12 += 43
                case "foo bar":
                    v16 := GenS1{F0: 216}
                    _ = v16
                case "":
                    var v17 int64 = (p0 + (^(p0)))
                    _ = v17
                default:
                    var v18 uint64 = 385
                    _ = v18
                }
                v19 := GenS1{F0: 120}
                _ = v19
            } else {
                var v20 string = "foo bar"
                switch 128 {
                case -1:
                    v21 := []int{0, 1}
                    _ = v21
                }
                _ = v20
            }
        }
    }
    fmt.Println(v13, v13, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 52
}

func main() {
    var v1 string = "a"
    v2 := []string{}
    var v3 uint = 4294967295
    v22 := GenS1{F0: 213}
    v3 = (v3 % ((v3) | 1))
    _ = len(v2)
    v23 := GenS2{F0: 1, F1: 2147483648}
    var v24 float32 = (1.5 * 0.1)
    var v25 bool = true
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 uint32 = 42
        if len(v2) > 0 {
            v2[0] = v1
        }
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: 128}
    v1 = "foo bar"
    v29 := GenS1{F0: 1}
    if true {
        for fi30 := 0; fi30 < 1; fi30++ {
            v24 += 3.14159
            var v31 int16 = -32768
            _ = fi30
            _ = v31
        }
        switch 9223372036854775807 {
        case 32768:
            if len(v2) > 0 {
                v2[0] = v1
            }
        case 949:
            v32 := GenS3{F0: v3, F1: 462, F2: 886}
            _ = v32
        }
    }
    v33 := GenS1{F0: 127}
    var v34 float64 = (1e10 * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
}
