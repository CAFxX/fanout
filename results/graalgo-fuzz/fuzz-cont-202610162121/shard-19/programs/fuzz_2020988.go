package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int }
type GenS2 struct { F0 uint64; F1 uint64 }
func genFn0() int64 {
    if false {
        for fi5 := 0; fi5 < 8; fi5++ {
            v6 := []int{(fi5 >> ((427) & 63)), (fi5 | (fi5 ^ -915))}
            v7 := GenS1{F0: fi5, F1: 2147483648, F2: fi5}
            var v8 float64 = 0.0
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    } else {
        var v9 int8 = -1
        _ = v9
    }
    var v10 float64 = 1e10
    _ = v10
    return -1
}

func main() {
    var v1 bool = false
    var v2 int64 = -110
    v3 := []int{-148, 128, 0, 4294967295}
    var v4 uint32 = 980
    v11 := GenS1{F0: 241, F1: v4, F2: 1000}
    v4 -= 2
    var v12 int64 = -9223372036854775808
    var v13 uint = 2
    for fi14 := 0; fi14 < 3; fi14++ {
        fmt.Println(fi14, v12)
        _ = fi14
    }
    v13 = v13
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int = fi15
        if v1 {
            v16 = -837
        } else {
            v13 += v13
            if len(v3) > 0 {
                v3[0] = (fi15 ^ 1)
            }
            v17 := GenS1{F0: (fi15 >> (((v13 >> (((v13 & 65535)) & 63))) & 63)), F1: v4, F2: 4294967295}
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    if true {
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 bool = ("a" == ("foo bar" + ""))
            _ = fi18
            _ = v19
        }
    }
    switch "foo bar" {
    case "hello":
        for fi20 := 0; fi20 < 13; fi20++ {
            v12 = (v12 - v12)
            for fi21 := 0; fi21 < 10; fi21++ {
                switch fi20 {
                case 9223372036854775807:
                    var v22 float64 = 0.0
                    _ = v22
                case 42:
                    var v23 uint32 = (v4 % ((v4) | 1))
                    _ = v23
                case -445:
                    v12 = 763
                }
                v24 := GenS2{F0: 18446744073709551615, F1: 729}
                if (2 != 65535) {
                    var v25 float32 = (0.0 / 1.0)
                    var v26 string = "foo bar"
                    var v27 uint64 = 4294967295
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    var v28 int8 = 100
                    _ = v28
                }
                _ = fi21
                _ = v24
            }
            if ((v4 & v4) <= 546) {
                var v29 int8 = -113
                _ = v29
            }
            _ = fi20
        }
    default:
        v4 = 42
    }
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
