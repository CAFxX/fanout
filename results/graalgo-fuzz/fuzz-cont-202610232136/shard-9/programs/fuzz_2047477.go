package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint }
type GenS4 struct { F0 int64; F1 bool; F2 int16 }
func genFn0(p0 uint) int {
    if (-2 != -87) {
        fmt.Println(p0)
    } else {
        p0 = 42
    }
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 uint16 = 42
        _ = fi4
        _ = v5
    }
    p0 -= 255
    _ = p0
    return -1
}

func main() {
    var v1 uint = 1
    var v2 uint8 = 127
    var v3 int16 = 42
    v1 = v1
    switch "a" {
    case "hello":
        v6 := GenS1{F0: 127}
        _ = v6
    case "":
        for fi7 := 0; fi7 < 18; fi7++ {
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := GenS1{F0: fi7}
                var v10 int8 = 100
                fmt.Println(v10)
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        fmt.Println(v1)
        _ = fi11
    }
    fmt.Println(v3)
    v3 = v3
    switch 9223372036854775807 {
    case -9223372036854775808:
        v1 = (^(v1))
    case 1:
        switch "foo bar" {
        case "a":
            v2 -= v2
        case "":
            fmt.Println(v2, v3, v2)
        }
    }
    switch 32768 {
    case 127:
        v3 -= 463
    case 2147483647:
        fmt.Println(v3, v3)
    }
    fmt.Println(v1, v1, v3)
    fmt.Println(v1)
    v12 := []int{-2147483648}
    var v13 int = 128
    if true {
        fmt.Println(v1)
        switch (v13 * (v13 ^ (v13 % (((v13 / ((int(v2)) | 1))) | 1)))) {
        case -1:
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 int8 = -104
                for fi16 := 0; fi16 < 8; fi16++ {
                    var v17 uint8 = (v2 % (((^(v2))) | 1))
                    v18 := GenS2{F0: (v15 + 100), F1: 865}
                    if true {
                        v19 := GenS3{F0: 2147483647, F1: 837, F2: (v1 ^ 114)}
                        fmt.Println(fi14, v17, v1)
                        _ = v19
                    }
                    if ("\n\t" >= "a") {
                        var v20 uint = v1
                        var v21 int = (fi16 - v13)
                        v22 := GenS2{F0: 107, F1: 2}
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS2{F0: v15, F1: 2}
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v23
                }
                fmt.Println(v15, v2)
                _ = fi14
                _ = v15
            }
        }
        _ = len(v12)
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 uint32 = 2147483648
        if false {
            switch (v13 / (((v13 - int(v2))) | 1)) {
            case 32767:
                var v26 uint32 = (v25 + v25)
                _ = v26
            case -2147483648:
                v27 := GenS4{F0: 1, F1: (false || ((-2.25 * 0.0) > (3.14159 - ((0.1 / 1.0) / ((-2.25 * 0.0) - ((float32(0.0) - (3.14159 * float32(float64(3.14159)))) - (((float32(float64(((float32((1.5 * 1.5)) + 0.0) * 0.0))) / ((3.14159 / 3.14159) * float32(1e10))) * (((1.5 - -2.25) * 3.14159) * 3.14159)) + -2.25))))))), F2: v3}
                _ = v27
            default:
                v12 = append(v12, genFn0(v1))
            }
            fmt.Println(fi24, fi24, v1)
        }
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
