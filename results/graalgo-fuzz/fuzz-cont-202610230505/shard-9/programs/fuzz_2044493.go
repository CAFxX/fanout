package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int }
type GenS2 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint32) int16 {
    p0 = (^(p0))
    for fi4 := 0; fi4 < 20; fi4++ {
        fmt.Println(p0, p0)
        _ = fi4
    }
    v5 := GenS1{F0: 42, F1: (1 <= 1000), F2: 650}
    for fi6 := 0; fi6 < 14; fi6++ {
        p0 += (p0 & p0)
        _ = fi6
    }
    _ = p0
    _ = v5
    return -481
}
func genFn1() string {
    var v7 int64 = 9223372036854775807
    v8 := []int{-1, -9223372036854775808, -2147483648, 128}
    switch 32767 {
    case 0:
        if false {
            switch (("hello" + "foo bar") + ((("\n\t" + ("foo bar" + ("\n\t" + "hello"))) + ("a" + "foo bar")) + "\n\t")) {
            case "hello":
                var v9 uint64 = 4294967295
                _ = v9
            default:
                var v10 string = "\n\t"
                _ = v10
            }
        } else {
            fmt.Println(v7)
        }
    case -9223372036854775808:
        v8 = append(v8, 0)
    case -712:
        var v11 int32 = 42
        _ = v11
    default:
        v7 = 42
    }
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 uint = 1
    var v2 uint = (v1 ^ v1)
    var v3 int64 = 0
    var v12 uint = 2
    var v13 float64 = 0.1
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := []int{fi14, (fi14 - (fi14 & fi14))}
        switch "" {
        case "\n\t":
            for fi16 := 0; fi16 < 9; fi16++ {
                v15 = append(v15, int(v1))
                _ = fi16
            }
        case "":
            v1 = (^(v12))
        }
        _ = fi14
        _ = v15
    }
    v17 := GenS2{F0: 42, F1: (v12 / ((v2) | 1)), F2: 4294967295}
    v12 -= v12
    var v18 uint16 = 42
    var v19 int64 = v3
    fmt.Println(v19, v19)
    var v20 uint8 = 2
    v21 := GenS3{F0: -9223372036854775808}
    v13 -= v13
    var v22 float32 = (0.0 + 3.14159)
    v1 -= v2
    for fi23 := 0; fi23 < 18; fi23++ {
        if (-34 >= 0) {
            if false {
                for fi24 := 0; fi24 < 17; fi24++ {
                    if (v13 >= 1e10) {
                        v25 := []int{1, fi24, (fi24 / (((fi23 << (((v2 / (((v2 | v2)) | 1))) & 63))) | 1)), fi24}
                        var v26 uint16 = v18
                        v27 := GenS4{F0: 4294967295}
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    var v28 uint64 = 2
                    var v29 uint8 = v20
                    _ = fi24
                    _ = v28
                    _ = v29
                }
                v19 += v19
            } else {
                var v30 int = fi23
                _ = v30
            }
            v12 += 1
        }
        v3 = v19
        v31 := GenS5{F0: v20}
        _ = fi23
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
