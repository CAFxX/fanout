package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 string }
func genFn0(p0 int32, p1 string) bool {
    var v5 bool = (true || true)
    v5 = v5
    v6 := []int{4294967295, int(p0), 128}
    v7 := GenS1{F0: 128}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (v5 || (v5 && v5))
}
func genFn1() int16 {
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int8 = 0
        if (false || (2147483648 <= 2147483648)) {
            v9 = (v9 - (v9 ^ (v9 ^ -86)))
            var v10 int64 = 726
            v11 := []int{4294967295, (fi8 - fi8), fi8}
            _ = v10
            _ = v11
        }
        var v12 float64 = 0.0
        _ = fi8
        _ = v12
        _ = v9
    }
    var v13 bool = true
    _ = v13
    return -32768
}

func main() {
    var v1 uint = 255
    var v2 uint32 = 0
    var v3 uint8 = 255
    var v4 string = ""
    v14 := []int{-9223372036854775808, 42, 1}
    v4 = "foo bar"
    v14 = append(v14, 4294967295)
    v15 := GenS2{F0: -517, F1: 47}
    fmt.Println(v3, v3)
    for fi16 := 0; fi16 < 4; fi16++ {
        v3 = (^(v3))
        _ = fi16
    }
    v14 = append(v14, int(v2))
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 uint32 = 2147483647
        _ = fi17
        _ = v18
    }
    v19 := GenS3{F0: ("foo bar" + ("foo bar" + "a"))}
    v1 = (v1 ^ 554)
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 9; fi21++ {
            switch "" {
            case "hello":
                v3 += v3
            case "foo bar":
                fmt.Println(fi21)
            case "":
                v1 -= 42
            default:
                v14 = append(v14, (fi20 * fi21))
            }
            var v22 int64 = 230
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v1 -= (v1 % (((v1 & (v1 << (((v1 << (((v1 * (v1 / ((v1) | 1)))) & 63))) & 63)))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
