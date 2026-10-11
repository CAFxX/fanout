package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0(p0 uint, p1 uint8, p2 uint32) bool {
    var v5 int64 = 0
    if (!(p0 >= uint(p2))) {
        var v6 int8 = -128
        _ = v6
    } else {
        var v7 int32 = 42
        var v8 bool = true
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := []int{fi9, fi9, (fi9 % ((-9223372036854775808) | 1)), (fi9 & (fi9 / ((32768) | 1)))}
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    }
    var v11 uint64 = 1
    switch "hello" {
    case "\n\t":
        var v12 float32 = ((-2.25 / 3.14159) / (-2.25 - (float32((((((3.14159 - -2.25) - float64(1.5)) * 0.0) / 1.0) + 0.0)) + 0.0)))
        _ = v12
    case "a":
        p0 += (^(p0))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v5
    return true
}

func main() {
    var v1 float32 = (1e10 / 3.14159)
    var v2 uint8 = 128
    var v3 float64 = (((1e10 / -2.25) * 1e10) * (1e10 + (0.1 - 3.14159)))
    v4 := make(map[int]int)
    if false {
        v2 = (v2 & v2)
        fmt.Println(v2, v3)
    }
    switch "a" {
    case "foo bar":
        v13 := v4[127]
        _ = v13
    default:
        for fi14 := 0; fi14 < 7; fi14++ {
            switch ("hello" + (((((("hello" + ("a" + (("a" + "\n\t") + "hello"))) + "a") + ("a" + "hello")) + "a") + "\n\t") + "hello")) {
            case "foo bar":
                v1 -= v1
            case "hello":
                var v15 float64 = (3.14159 + 0.0)
                _ = v15
            }
            _ = fi14
        }
    }
    v16 := []int{42, 42}
    v2 += v2
    switch "foo bar" {
    case "":
        v3 -= v3
    case "a":
        fmt.Println(v2, v2)
    }
    v17 := GenS1{F0: (("foo bar" + ((("foo bar" + "hello") + "\n\t") + "hello")) + "foo bar"), F1: 100}
    switch "\n\t" {
    case "foo bar":
        if (!(65535 != 2)) {
            var v18 int = -2147483648
            _ = v18
        } else {
            for fi19 := 0; fi19 < 15; fi19++ {
                v3 += (v3 * v3)
                _ = fi19
            }
            if (-1 == -1) {
                v3 = (float64(v1) / 1.0)
            }
        }
    case "\n\t":
        if true {
            v1 = 0.0
            v4[-9223372036854775808] = -961
        } else {
            var v20 uint8 = (v2 / (((v2 % (((v2 << ((1) & 63))) | 1))) | 1))
            v3 = v3
            v21 := GenS1{F0: ("a" + "foo bar"), F1: 42}
            _ = v20
            _ = v21
        }
    default:
        if false {
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            var v22 float64 = v3
            v16 = append(v16, -330)
            _ = v22
        }
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
