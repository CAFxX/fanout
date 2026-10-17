package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 bool; F1 int8; F2 int8 }
func genFn0(p0 uint, p1 uint) string {
    fmt.Println(p0, p1)
    fmt.Println(p1, p1)
    v6 := []int{1}
    p0 -= (p1 % ((p1) | 1))
    v7 := GenS1{F0: 1, F1: 0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ("a" + "")
}
func genFn1(p0 int8) bool {
    if false {
        v8 := []int{42, 1}
        _ = v8
    } else {
        if true {
            v9 := GenS2{F0: (false || true), F1: (p0 * 127), F2: (p0 | (p0 % (((-(p0))) | 1)))}
            _ = v9
        }
        fmt.Println(p0)
        if (false || true) {
            var v10 uint32 = 65535
            var v11 uint8 = 2
            _ = v10
            _ = v11
        }
    }
    switch "hello" {
    case "\n\t":
        v12 := []int{2147483647}
        _ = v12
    default:
        v13 := []int{}
        _ = v13
    }
    _ = p0
    return (false || ((true && (4294967295 < 780)) || false))
}

func main() {
    var v1 int = 42
    v2 := make(map[int]uint32)
    var v3 uint = 42
    var v4 float32 = 3.14159
    var v5 int16 = -810
    switch v1 {
    case 2147483647:
        fmt.Println(v4, v4)
    case -1:
        v14 := GenS1{F0: 42, F1: 2147483647}
        _ = v14
    case 9223372036854775807:
        v3 = 4294967295
    default:
        v5 = v5
    }
    switch "a" {
    case "foo bar":
        switch (v1 ^ (v1 ^ v1)) {
        case 32767:
            var v15 int16 = (^(v5))
            _ = v15
        default:
            v5 = (v5 | (v5 << ((v3) & 63)))
        }
    case "hello":
        var v16 int8 = -121
        _ = v16
    default:
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 bool = (v3 > (v3 ^ v3))
            var v19 int64 = 1
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    v3 = (v3 / (((v3 ^ 0)) | 1))
    var v20 string = ("foo bar" + ("" + "a"))
    fmt.Println(v4, v4, v4)
    v21 := GenS2{F0: (-1 > 114), F1: 100, F2: 0}
    switch ((("hello" + v20) + ("hello" + v20)) + ((v20 + "") + (v20 + v20))) {
    case "":
        v5 = v5
    case "\n\t":
        v22 := GenS1{F0: 100, F1: -1}
        _ = v22
    case "foo bar":
        v1 += 127
    }
    var v23 float32 = v4
    for fi24 := 0; fi24 < 15; fi24++ {
        if (0 > 4294967295) {
            var v25 uint32 = 4294967295
            _ = v25
        }
        var v26 uint32 = 65535
        switch (v20 + ("hello" + (v20 + v20))) {
        case "foo bar":
            v26 += v26
        case "":
            for fi27 := 0; fi27 < 7; fi27++ {
                var v28 int16 = -355
                _ = fi27
                _ = v28
            }
        case "a":
            v29 := GenS1{F0: -1, F1: -1}
            _ = v29
        default:
            fmt.Println(v23, v1, v5)
        }
        _ = fi24
        _ = v26
    }
    v5 = 1
    fmt.Println(v20)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
