package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]bool)
    var v3 float32 = 0.0
    var v4 float64 = 0.0
    v5 := GenS1{F0: (v1 - 1), F1: "foo bar"}
    v4 = v4
    v6 := v2[127]
    _ = len(v2)
    _ = len(v2)
    if v6 {
        if v6 {
            fmt.Println(v3)
            v7 := []int{9223372036854775807}
            _ = v7
        } else {
            v4 = v4
        }
    }
    if v6 {
        for fi8 := 0; fi8 < 3; fi8++ {
            switch "\n\t" {
            case "a":
                if (2147483648 != 785) {
                    delete(v2, (fi8 | fi8))
                    for fi9 := 0; fi9 < 10; fi9++ {
                        fmt.Println(fi9, fi9, v1)
                        _ = fi9
                    }
                }
            case "hello":
                v4 = ((((float64((float32(((v4 / (float64(v3) - (v4 + 1e10))) - (float64(float32(float64((v3 - v3)))) / 1.0))) * v3)) / v4) + (3.14159 - 0.1)) - (v4 / v4)) * (1e10 * (1.5 - v4)))
            }
            if ((v1 << ((94) & 63)) < -2147483648) {
                v4 -= v4
            } else {
                var v10 int64 = 262
                fmt.Println(fi8)
                if v6 {
                    switch fi8 {
                    case 0:
                        v11 := []int{fi8, (fi8 | 128), -1, (-(fi8))}
                        _ = v11
                    case 1:
                        switch -290 {
                        case -543:
                            for fi12 := 0; fi12 < 19; fi12++ {
                                v13 := []int{fi8, (fi12 * fi12)}
                                var v14 int8 = int8(fi8)
                                v15 := GenS1{F0: (v1 >> ((77) & 63)), F1: "foo bar"}
                                v6 = v6
                                _ = fi12
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                        case 2147483647:
                            var v16 string = ""
                            _ = v16
                        }
                    case -806:
                        var v17 int8 = 0
                        _ = v17
                    }
                }
                _ = v10
            }
            _ = fi8
        }
    }
    if true {
        v18 := v2[32768]
        fmt.Println(v1, v6)
        _ = v18
    } else {
        fmt.Println(v1, v1, v1)
    }
    if v6 {
        switch "hello" {
        case "foo bar":
            var v19 string = "foo bar"
            _ = v19
        default:
            v20 := []int{-886, 9223372036854775807, -999}
            _ = v20
        }
        if false {
            var v21 bool = (true || (304 <= 2))
            _ = v21
        } else {
            v6 = ((v3 - (v3 + v3)) >= 0.0)
            var v22 int8 = 42
            _ = v22
        }
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
