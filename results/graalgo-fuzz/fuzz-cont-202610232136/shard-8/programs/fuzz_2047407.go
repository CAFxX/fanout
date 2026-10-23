package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int64 }

func main() {
    var v1 bool = (-9223372036854775808 < -811)
    var v2 uint = 65535
    var v3 float64 = 0.1
    var v4 float64 = ((v3 + v3) * 0.1)
    if (-1 >= -194) {
        v5 := []int{}
        if true {
            _ = len(v5)
            var v6 string = ("hello" + "\n\t")
            _ = v6
        }
        _ = v5
    } else {
        switch 257 {
        case 9223372036854775807:
            v7 := []int{1}
            _ = v7
        case 42:
            v8 := []int{42, -175}
            _ = v8
        case 32767:
            for fi9 := 0; fi9 < 1; fi9++ {
                v10 := []int{(fi9 | fi9), (fi9 * fi9)}
                if len(v10) > 0 {
                    v10[0] = -794
                }
                v11 := GenS1{F0: "", F1: 2, F2: 0}
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        v12 := []int{}
        var v13 int64 = 606
        _ = v12
        _ = v13
    }
    switch "a" {
    case "":
        if true {
            if v1 {
                if true {
                    v3 += (v3 + (0.0 - (v3 / v3)))
                }
                var v14 uint32 = 966
                v15 := GenS1{F0: "", F1: 2, F2: 289}
                _ = v14
                _ = v15
            }
        } else {
            var v16 int64 = -1
            v17 := GenS2{F0: 9223372036854775807}
            v1 = (!v1)
            _ = v16
            _ = v17
        }
    default:
        v18 := GenS1{F0: "hello", F1: 32767, F2: 42}
        _ = v18
    }
    switch 1 {
    case 9223372036854775807:
        v19 := GenS2{F0: -1}
        _ = v19
    case 908:
        fmt.Println(v1)
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        v4 = v3
        var v21 uint8 = 1
        if v1 {
            fmt.Println(fi20, v2)
            v22 := []int{fi20, (fi20 >> ((v2) & 63))}
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    v2 = v2
    switch 9223372036854775807 {
    case 42:
        v23 := GenS2{F0: 32767}
        _ = v23
    case 127:
        v2 = v2
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := GenS2{F0: (-(fi24))}
        _ = fi24
        _ = v25
    }
    switch 2147483647 {
    case 2147483647:
        v26 := GenS3{F0: 42, F1: 464}
        _ = v26
    case 128:
        fmt.Println(v4, v2)
    }
    var v27 string = ""
    switch 965 {
    case 128:
        fmt.Println(v2, v2, v27)
    case 32767:
        v27 = ("foo bar" + "foo bar")
    case 386:
        v28 := []int{128, -344, 4294967295}
        _ = v28
    }
    var v29 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v4
}
