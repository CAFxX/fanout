package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int; F1 int32; F2 uint }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int = 0
    v2 := make(map[string]bool)
    var v3 string = "hello"
    var v4 int32 = 2147483647
    delete(v2, "")
    v3 = "a"
    v3 = v3
    v3 = v3
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: 42, F1: (v1 * v1)}
        v2[(v3 + v3)] = false
        _ = fi5
        _ = v6
    }
    if (!true) {
        v7 := GenS2{F0: v1, F1: 803, F2: 276}
        switch v3 {
        case "a":
            v8 := []int{(v1 >> ((487) & 63)), v1}
            _ = v8
        }
        v9 := []int{(v1 ^ (v1 >> ((4294967295) & 63))), (v1 & (v1 % (((^(v1))) | 1))), (v1 % ((v1) | 1))}
        _ = v7
        _ = v9
    } else {
        v3 = "hello"
        var v10 float64 = (3.14159 / 1e10)
        _ = v10
    }
    if true {
        fmt.Println(v3)
        if false {
            v11 := []int{(v1 << ((702) & 63)), (v1 - v1), v1, v1}
            switch v1 {
            case -2147483648:
                v3 = (v3 + v3)
            case -873:
                for fi12 := 0; fi12 < 16; fi12++ {
                    v3 = v3
                    _ = fi12
                }
            case 4294967295:
                _ = len(v11)
            }
            _ = v11
        } else {
            for fi13 := 0; fi13 < 7; fi13++ {
                for fi14 := 0; fi14 < 2; fi14++ {
                    switch fi13 {
                    case -113:
                        var v15 int16 = -690
                        v1 -= (-(fi14))
                        var v16 int64 = -351
                        var v17 int64 = 9223372036854775807
                        _ = v15
                        _ = v16
                        _ = v17
                    case -9223372036854775808:
                        v4 -= v4
                    default:
                        v4 = v4
                    }
                    _ = fi14
                }
                for fi18 := 0; fi18 < 14; fi18++ {
                    if false {
                        fmt.Println(v4, v1)
                    } else {
                        v19 := GenS2{F0: v1, F1: v4, F2: 42}
                        v20 := v2[v3]
                        _ = v19
                        _ = v20
                    }
                    v4 = (v4 % (((v4 << ((492) & 63))) | 1))
                    switch ("\n\t" + "hello") {
                    case "":
                        fmt.Println(fi18, fi13)
                    }
                    _ = fi18
                }
                _ = fi13
            }
        }
        v4 -= v4
    }
    if true {
        v21 := GenS2{F0: (v1 & v1), F1: v4, F2: 18446744073709551615}
        _ = v21
    } else {
        switch (v3 + v3) {
        case "a":
            var v22 uint8 = 128
            _ = v22
        case "\n\t":
            for fi23 := 0; fi23 < 17; fi23++ {
                var v24 int16 = 0
                _ = fi23
                _ = v24
            }
        case "":
            v25 := GenS1{F0: 42, F1: v1}
            _ = v25
        }
        var v26 float32 = float32(((1.5 * ((0.1 * 1.5) - 1.5)) / 1.5))
        _ = v26
    }
    switch "\n\t" {
    case "foo bar":
        v27 := GenS3{F0: -76}
        _ = v27
    default:
        fmt.Println(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
