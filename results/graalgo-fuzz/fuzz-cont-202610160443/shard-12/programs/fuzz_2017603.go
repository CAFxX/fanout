package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 int; F1 int8 }
func genFn0(p0 int16, p1 string) string {
    if false {
        p1 = p1
    } else {
        var v7 int8 = 1
        _ = v7
    }
    v8 := []int{}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v8
    return (p1 + (p1 + p1))
}
func genFn1() string {
    var v9 int64 = 9223372036854775807
    var v10 float32 = 1e10
    if ((v9 > v9) || false) {
        var v11 uint32 = 2147483648
        fmt.Println(v10)
        fmt.Println(v10, v10, v9)
        _ = v11
    } else {
        v12 := []int{}
        _ = v12
    }
    _ = v10
    _ = v9
    return "foo bar"
}

func main() {
    var v1 float64 = 0.1
    v2 := []int{-9223372036854775808, 998, 128}
    var v3 uint32 = 0
    v4 := make(map[int]uint32)
    var v5 int32 = 0
    v6 := []string{"hello", ("foo bar" + "foo bar"), ("a" + (("a" + "a") + "hello"))}
    switch "\n\t" {
    case "\n\t":
        v13 := GenS1{F0: v3, F1: 1}
        _ = v13
    }
    v5 = -1
    v14 := GenS2{F0: (^(v3)), F1: (-(v5))}
    v3 = v3
    fmt.Println(v3)
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS3{F0: ("" + "hello")}
        v3 = 0
        if (300 == 65535) {
            var v17 uint32 = (v3 % (((v3 * (v3 * 529))) | 1))
            if false {
                v18 := GenS4{F0: 32768, F1: (-(fi15))}
                if false {
                    v19 := GenS3{F0: "a"}
                    v1 -= ((0.0 - v1) - ((((1.5 * 1e10) - v1) - 1e10) + 1e10))
                    _ = v19
                }
                var v20 int = fi15
                _ = v18
                _ = v20
            } else {
                var v21 float64 = v1
                _ = v21
            }
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 bool = true
                var v24 bool = v23
                v25 := GenS3{F0: ("" + ((("\n\t" + ("hello" + (("" + "") + "foo bar"))) + "a") + ""))}
                v1 = v1
                v3 += v3
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v1, v5, v3)
    v26 := GenS3{F0: "hello"}
    if true {
        if (!false) {
            v27 := GenS5{F0: 282, F1: 100}
            _ = v27
        } else {
            if true {
                fmt.Println(v1)
                var v28 uint = uint(v3)
                switch "hello" {
                case "hello":
                    v4[128] = (v3 ^ 0)
                }
                _ = v28
            } else {
                fmt.Println(v1)
                var v29 int16 = 42
                _ = v29
            }
            _ = len(v4)
        }
    } else {
        fmt.Println(v1)
        delete(v4, 42)
        var v30 float32 = (1.5 * 0.0)
        _ = v30
    }
    if true {
        for fi31 := 0; fi31 < 13; fi31++ {
            if false {
                if false {
                    v32 := v4[fi31]
                    var v33 int8 = 1
                    var v34 int64 = 341
                    _ = v32
                    _ = v33
                    _ = v34
                }
                v5 += (v5 / (((v5 | -2147483648)) | 1))
                v3 -= 65535
            }
            v1 -= v1
            v35 := GenS2{F0: v3, F1: (v5 * 1)}
            _ = fi31
            _ = v35
        }
        v1 -= float64(float32(3.14159))
    }
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
