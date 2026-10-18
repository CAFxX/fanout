package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int8; F1 uint64; F2 uint32 }

func main() {
    var v1 uint16 = 992
    var v2 uint64 = 137
    v3 := make(map[int]int)
    v4 := GenS1{F0: 65535, F1: 940}
    v5 := GenS1{F0: uint(v2), F1: 473}
    v2 = (v2 - 42)
    switch 203 {
    case 128:
        if true {
            v3[-858] = 127
            var v6 int = 42
            v2 = (^(v2))
            _ = v6
        }
    default:
        if true {
            if (((^(v1)) > 1) || ((1 != 42) || true)) {
                v7 := []int{32768, 1, 2147483647, 2147483647}
                _ = v7
            }
            v8 := []int{}
            switch 173 {
            case 4294967295:
                fmt.Println(v2)
            }
            _ = v8
        } else {
            for fi9 := 0; fi9 < 16; fi9++ {
                switch "a" {
                case "foo bar":
                    _ = len(v3)
                case "\n\t":
                    v10 := []int{fi9, -2147483648}
                    _ = v10
                default:
                    var v11 float64 = (1.5 - 1e10)
                    _ = v11
                }
                for fi12 := 0; fi12 < 14; fi12++ {
                    _ = len(v3)
                    _ = fi12
                }
                _ = fi9
            }
        }
    }
    v13 := GenS1{F0: 255, F1: 65535}
    _ = len(v3)
    if true {
        var v14 int32 = -970
        _ = v14
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS1{F0: 60, F1: 864}
            fmt.Println(v2, v1, fi15)
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 float64 = 0.1
                v19 := GenS2{F0: 207, F1: 130}
                v20 := GenS3{F0: 30, F1: fi17}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            _ = fi15
            _ = v16
        }
    }
    switch 4294967295 {
    case -1:
        var v21 int64 = 0
        _ = v21
    }
    v2 = v2
    for fi22 := 0; fi22 < 14; fi22++ {
        v3[fi22] = fi22
        v3[(^(fi22))] = fi22
        var v23 int = fi22
        _ = fi22
        _ = v23
    }
    v2 = (v2 + 2)
    var v24 uint32 = 2147483648
    if (v24 != v24) {
        var v25 uint = 637
        v25 = 65535
        var v26 float64 = 1e10
        _ = v25
        _ = v26
    } else {
        v27 := GenS4{F0: 1000, F1: -1}
        var v28 int8 = 42
        v2 += (^(v2))
        _ = v27
        _ = v28
    }
    v29 := GenS5{F0: 94, F1: (^(v2)), F2: 42}
    var v30 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
