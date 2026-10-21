package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int }
type GenS3 struct { F0 int64; F1 bool; F2 int }
type GenS4 struct { F0 int32 }

func main() {
    var v1 int16 = 1
    v2 := make(map[string]uint32)
    var v3 float64 = 0.1
    var v4 float64 = 1.5
    var v5 uint32 = 71
    var v6 uint64 = 0
    fmt.Println(v6)
    v1 += v1
    v7 := GenS1{F0: "", F1: 18, F2: 42}
    fmt.Println(v4, v1)
    if false {
        v8 := []int{}
        for fi9 := 0; fi9 < 1; fi9++ {
            switch (fi9 / ((fi9) | 1)) {
            case 4294967295:
                switch fi9 {
                case 32767:
                    for fi10 := 0; fi10 < 19; fi10++ {
                        if (-941 >= -586) {
                            v11 := GenS1{F0: "", F1: 2, F2: 2}
                            _ = v11
                        }
                        var v12 uint16 = 255
                        var v13 uint = 0
                        _ = fi10
                        _ = v12
                        _ = v13
                    }
                case 0:
                    v5 = (v5 - (v5 | (v5 ^ v5)))
                case -9223372036854775808:
                    if (!false) {
                        if false {
                            v1 += v1
                        }
                        fmt.Println(v1, v3)
                    }
                default:
                    var v14 uint16 = 674
                    _ = v14
                }
            case 32767:
                v2["hello"] = v5
            }
            var v15 float32 = 1e10
            _ = fi9
            _ = v15
        }
        v16 := GenS2{F0: 202, F1: 32767, F2: 576}
        _ = v16
        _ = v8
    } else {
        if false {
            v4 = (v3 / (0.0 - float64((1e10 * 1.5))))
        } else {
            if true {
                v17 := []int{1, 0, 4294967295}
                fmt.Println(v5, v5, v1)
                _ = v17
            }
            var v18 int64 = -9223372036854775808
            v19 := GenS3{F0: (^(v18)), F1: true, F2: 32768}
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 int16 = 1
            v22 := []int{fi20, (fi20 << ((4294967295) & 63))}
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v23 := GenS1{F0: "", F1: 0, F2: 2}
    v4 = 0.0
    v3 = (v4 / v3)
    fmt.Println(v4, v5)
    v3 -= 0.1
    switch "\n\t" {
    case "a":
        v24 := GenS3{F0: 42, F1: false, F2: -1}
        _ = v24
    case "":
        switch (("" + "hello") + "\n\t") {
        case "\n\t":
            v6 -= (v6 ^ 4294967295)
        case "foo bar":
            v25 := GenS4{F0: -2147483648}
            _ = v25
        case "hello":
            v1 = 0
        }
    case "\n\t":
        v4 = 0.0
    default:
        v26 := GenS1{F0: "", F1: 235, F2: 0}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
