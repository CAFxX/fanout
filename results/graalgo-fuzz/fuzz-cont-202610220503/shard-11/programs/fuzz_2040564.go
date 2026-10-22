package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int }
func genFn0(p0 uint32, p1 int64, p2 uint8) uint16 {
    fmt.Println(p2)
    var v6 uint64 = 1
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint8 = 194
        var v9 int32 = -2147483648
        for fi10 := 0; fi10 < 3; fi10++ {
            if ((fi7 >> ((0) & 63)) <= (fi7 / ((fi10) | 1))) {
                fmt.Println(p0, v9, p1)
                switch fi7 {
                case 2147483647:
                    var v11 uint = 0
                    var v12 uint32 = 42
                    _ = v11
                    _ = v12
                case 9223372036854775807:
                    var v13 int16 = 1
                    _ = v13
                case 984:
                    fmt.Println(v8, p1, p2)
                }
                var v14 int64 = (p1 - (^(p1)))
                _ = v14
            }
            v8 += p2
            _ = fi10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if (18446744073709551615 >= 255) {
        fmt.Println(v6, p1, v6)
    }
    var v15 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v6
    return 65535
}

func main() {
    v1 := []string{}
    var v2 string = ""
    var v3 int8 = -69
    v4 := make(map[int]int)
    v5 := make(map[uint32]int)
    v16 := GenS1{F0: 2147483647, F1: genFn0(0, 42, 128), F2: 853}
    v17 := GenS2{F0: -576, F1: 374}
    switch "a" {
    case "\n\t":
        v18 := GenS3{F0: 42, F1: 144, F2: 1}
        _ = v18
    case "hello":
        v2 = "\n\t"
    }
    fmt.Println(v3)
    if true {
        v19 := GenS2{F0: 42, F1: 2147483648}
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 string = v2
            v1 = append(v1, (v21 + "a"))
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    if true {
        v22 := GenS1{F0: 2147483648, F1: 2, F2: 1}
        var v23 uint32 = 2
        if false {
            var v24 int64 = -9223372036854775808
            _ = v24
        }
        _ = v22
        _ = v23
    }
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        v3 = (v3 | v3)
        v25 := GenS3{F0: -2147483648, F1: 42, F2: 255}
        v26 := GenS4{F0: 1, F1: 128}
        _ = v25
        _ = v26
    } else {
        if false {
            var v27 string = v2
            _ = v27
        }
    }
    var v28 float32 = float32((1e10 * 1e10))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
