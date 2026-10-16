package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int64; F2 uint64 }
func genFn0() int32 {
    if ((!(!false)) && false) {
        var v4 float64 = 0.1
        _ = v4
    } else {
        var v5 float32 = (3.14159 + (0.1 + 1e10))
        for fi6 := 0; fi6 < 13; fi6++ {
            fmt.Println(v5, v5)
            switch -9223372036854775808 {
            case -606:
                v7 := []int{fi6}
                _ = v7
            case 2147483647:
                switch ("foo bar" + ("hello" + "")) {
                case "":
                    v5 -= -2.25
                case "a":
                    v8 := []int{}
                    _ = v8
                default:
                    v9 := GenS1{F0: 42, F1: fi6}
                    _ = v9
                }
            }
            _ = fi6
        }
        v10 := []int{4, 9223372036854775807, 127, -571}
        _ = v10
        _ = v5
    }
    var v11 uint64 = 1
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS2{F0: 1}
        v14 := []int{9223372036854775807, -245, fi12}
        fmt.Println(fi12)
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = v11
    return 0
}

func main() {
    v1 := []int{42, -566, 4294967295}
    var v2 uint32 = 977
    var v3 int = -29
    for fi15 := 0; fi15 < 6; fi15++ {
        switch v3 {
        case 32767:
            for fi16 := 0; fi16 < 18; fi16++ {
                switch "" {
                case "a":
                    var v17 float32 = 1.5
                    _ = v17
                case "":
                    v3 -= fi16
                }
                _ = fi16
            }
        case 127:
            var v18 float64 = (1.5 / 3.14159)
            _ = v18
        default:
            for fi19 := 0; fi19 < 11; fi19++ {
                v2 -= v2
                v20 := GenS2{F0: 1}
                _ = fi19
                _ = v20
            }
        }
        v2 -= v2
        _ = fi15
    }
    fmt.Println(v2, v2)
    var v21 int32 = 1
    v21 = genFn0()
    switch (-(v3)) {
    case 0:
        switch "foo bar" {
        case "hello":
            _ = len(v1)
        case "\n\t":
            fmt.Println(v3)
        }
    case 127:
        v3 -= 4294967295
    case 32767:
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 float64 = 3.14159
            v21 = v21
            _ = fi22
            _ = v23
        }
    }
    var v24 uint16 = 65535
    var v25 uint64 = 845
    v26 := GenS2{F0: 255}
    v1 = append(v1, 1)
    fmt.Println(v21, v25, v3)
    switch ("hello" + "") {
    case "hello":
        v27 := GenS3{F0: 0, F1: 0, F2: (v25 * (v25 - (v25 ^ uint64(v2))))}
        _ = v27
    case "\n\t":
        var v28 bool = true
        _ = v28
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
