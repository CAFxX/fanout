package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 int; F1 uint8; F2 int8 }
func genFn0() uint8 {
    switch ("hello" + ("" + "hello")) {
    case "hello":
        var v7 string = "\n\t"
        _ = v7
    }
    var v8 uint32 = 42
    v9 := GenS1{F0: "hello", F1: 1}
    switch "\n\t" {
    case "a":
        var v10 float64 = 0.1
        _ = v10
    }
    v11 := []int{}
    _ = v11
    _ = v8
    _ = v9
    return 129
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = (-2.25 * (((3.14159 + 1e10) * 0.1) - 1e10))
    v3 := make(map[int]string)
    var v4 int32 = 1
    v5 := []int{114, -156, -2147483648, 127}
    v6 := make(map[string]int)
    var v12 int = 1
    v13 := GenS1{F0: "hello", F1: 42}
    var v14 float32 = v2
    switch (v12 * (v12 ^ (v12 >> ((4294967295) & 63)))) {
    case -9223372036854775808:
        v12 += v12
    case 32767:
        v4 -= (v4 / ((v4) | 1))
    case 469:
        switch (("foo bar" + "a") + ("foo bar" + (("" + "") + (("" + "a") + "\n\t")))) {
        case "hello":
            v14 = ((((1e10 - 0.1) * v14) + v14) - (((v14 - (1.5 * (v14 + ((v2 - 3.14159) * 1.5)))) + v14) * v2))
        case "\n\t":
            var v15 int64 = 1
            _ = v15
        }
    }
    fmt.Println(v12, v2)
    v4 += (^(v4))
    if ((0 != 1) && true) {
        v14 = float32((3.14159 + 1.5))
        v2 = (1e10 / 1.0)
        v12 = 34
    } else {
        if false {
            var v16 uint16 = 65535
            var v17 bool = false
            _ = v16
            _ = v17
        }
        var v18 float64 = (((0.0 - 1e10) + ((-2.25 / (-2.25 - (0.1 * 1e10))) + 0.0)) + -2.25)
        _ = v18
    }
    var v19 int8 = -128
    var v20 float32 = v14
    var v21 float64 = float64(3.14159)
    switch v12 {
    case 1:
        v22 := GenS2{F0: 2, F1: 47, F2: 172}
        _ = v22
    case 984:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 4294967295:
        if (0 > 1) {
            var v23 uint16 = 32767
            v24 := GenS3{F0: v12, F1: 2, F2: 13}
            _ = v23
            _ = v24
        } else {
            var v25 uint8 = 113
            for fi26 := 0; fi26 < 3; fi26++ {
                if len(v5) > 0 {
                    v5[0] = fi26
                }
                v12 = v12
                _ = fi26
            }
            _ = v25
        }
    }
    v27 := GenS2{F0: 255, F1: 154, F2: 552}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
