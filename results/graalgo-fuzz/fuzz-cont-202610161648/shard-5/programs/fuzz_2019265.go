package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint, p1 int) uint64 {
    for fi4 := 0; fi4 < 12; fi4++ {
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 int32 = 42
            p1 += (^(fi4))
            if ((p0 - p0) <= p0) {
                switch int(p0) {
                case -2147483648:
                    v7 := GenS1{F0: ("foo bar" + "")}
                    _ = v7
                case 1:
                    var v8 int = (fi5 * (fi5 >> ((p0) & 63)))
                    _ = v8
                case 32767:
                    fmt.Println(p1, v6)
                default:
                    v9 := GenS1{F0: "a"}
                    _ = v9
                }
            } else {
                v10 := []int{}
                if (!false) {
                    var v11 uint32 = 2147483647
                    _ = v11
                } else {
                    fmt.Println(fi4, p1)
                    var v12 float64 = 0.0
                    _ = v12
                }
                _ = v10
            }
            _ = fi5
            _ = v6
        }
        v13 := GenS2{F0: fi4, F1: 90}
        _ = fi4
        _ = v13
    }
    fmt.Println(p1, p0)
    p0 = (p0 << ((p0) & 63))
    _ = p0
    _ = p1
    return 405
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v14 int32 = 546
    var v15 int32 = (^(v14))
    v16 := GenS3{F0: v14, F1: 1}
    switch 42 {
    case 9223372036854775807:
        switch "hello" {
        case "\n\t":
            if (!(!(-1 >= 1))) {
                v14 = -2147483648
                fmt.Println(v14)
            }
        case "":
            v17 := []int{734}
            _ = v17
        case "hello":
            v18 := GenS1{F0: ""}
            _ = v18
        default:
            fmt.Println(v15)
        }
    case -1:
        v19 := []int{-46}
        _ = v19
    case 128:
        v14 += -1
    default:
        var v20 float64 = 0.0
        _ = v20
    }
    v21 := v3[2147483647]
    for fi22 := 0; fi22 < 12; fi22++ {
        v15 = (v15 / ((v14) | 1))
        for fi23 := 0; fi23 < 4; fi23++ {
            _ = len(v3)
            if (false || true) {
                v1 += v1
                v24 := GenS3{F0: (v15 & v15), F1: (v21 / ((v21) | 1))}
                var v25 uint64 = 234
                if true {
                    v26 := v3[0]
                    _ = v26
                }
                v27 := []int{(v21 >> ((uint(v1)) & 63)), 922, fi23, 341}
                _ = v24
                _ = v25
                _ = v27
            }
            _ = fi23
        }
        if false {
            switch v21 {
            case -127:
                v1 = 254
            default:
                var v28 uint32 = 0
                _ = v28
            }
            fmt.Println(v14, v14)
        }
        _ = fi22
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        fmt.Println(v14, fi29, fi29)
        _ = fi29
    }
    if true {
        switch "hello" {
        case "":
            v30 := GenS4{F0: 0}
            _ = v30
        default:
            var v31 int8 = -77
            _ = v31
        }
        var v32 uint16 = 694
        v15 = (v15 + 1)
        _ = v32
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        switch "hello" {
        case "\n\t":
            fmt.Println(v15, v15, v21)
        }
        var v34 uint32 = 2147483647
        _ = fi33
        _ = v34
    }
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
}
