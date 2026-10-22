package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint32 }

func main() {
    var v1 uint32 = 1
    var v2 uint64 = 2
    var v3 string = ("foo bar" + "foo bar")
    v1 += (^(v1))
    v4 := GenS1{F0: 17}
    v5 := GenS2{F0: v1, F1: 127, F2: v1}
    switch v3 {
    case "a":
        switch 9223372036854775807 {
        case 32768:
            var v6 uint = 0
            _ = v6
        case 127:
            v7 := []int{654, 0, 1}
            _ = v7
        }
    case "hello":
        if ((v1 & (v1 ^ 682)) != 500) {
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 uint32 = v1
                var v10 uint64 = v2
                _ = fi8
                _ = v10
                _ = v9
            }
            fmt.Println(v2)
            if (128 >= 95) {
                fmt.Println(v1, v3)
                if (4294967295 < 2147483647) {
                    var v11 int64 = -418
                    var v12 uint8 = 139
                    _ = v11
                    _ = v12
                }
                v1 += (v1 | v1)
            }
        }
    case "\n\t":
        switch -423 {
        case -1:
            var v13 float64 = 0.0
            _ = v13
        case 45:
            var v14 uint8 = 210
            _ = v14
        case 905:
            var v15 float32 = (-2.25 * 1.5)
            _ = v15
        }
    default:
        if true {
            v16 := []int{-818}
            fmt.Println(v1, v3, v2)
            _ = v16
        }
    }
    v17 := GenS2{F0: (v1 % ((v1) | 1)), F1: 2, F2: (v1 / (((v1 - 4294967295)) | 1))}
    if (!(v2 == (v2 << ((uint(v2)) & 63)))) {
        fmt.Println(v2)
        v3 = v3
        for fi18 := 0; fi18 < 15; fi18++ {
            v3 = v3
            var v19 uint8 = 12
            fmt.Println(v19, fi18, v19)
            _ = fi18
            _ = v19
        }
    } else {
        v2 += v2
    }
    v1 = v1
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 int32 = 42
        if true {
            fmt.Println(v21, v1, v2)
            fmt.Println(v2, v1)
        }
        if false {
            v22 := []int{fi20}
            switch "\n\t" {
            case "\n\t":
                v22 = append(v22, -9223372036854775808)
            case "":
                v3 = "hello"
            case "hello":
                var v23 int8 = 116
                _ = v23
            }
            if ((128 >= 42) || (!(188 == 127))) {
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 bool = false
                    var v26 uint64 = (v2 % ((v2) | 1))
                    var v27 bool = (42 == 55)
                    if v27 {
                        if true {
                            switch fi24 {
                            case 127:
                                var v28 float64 = 0.0
                                _ = v28
                            default:
                                var v29 float32 = (3.14159 + 3.14159)
                                _ = v29
                            }
                            fmt.Println(fi20, v25, v27)
                            if len(v22) > 0 {
                                v22[0] = -1
                            }
                        }
                        v22 = append(v22, (fi20 | fi24))
                        v30 := GenS1{F0: 591}
                        var v31 int8 = 42
                        _ = v30
                        _ = v31
                    }
                    var v32 uint = 1
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v32
                }
            } else {
                v3 = "foo bar"
            }
            _ = v22
        } else {
            v33 := []int{fi20}
            var v34 uint64 = v2
            _ = v33
            _ = v34
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
