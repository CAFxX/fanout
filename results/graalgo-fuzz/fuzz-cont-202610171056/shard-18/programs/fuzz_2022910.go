package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint8; F1 int16 }
func genFn0(p0 bool, p1 int64) int64 {
    var v7 string = ""
    var v8 float64 = 3.14159
    var v9 int = -2147483648
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (p1 * (p1 ^ 583))
}

func main() {
    v1 := []string{"\n\t", ((((("" + "a") + ("" + ("hello" + ("hello" + (("\n\t" + "a") + ((("" + "") + "") + "hello")))))) + "hello") + ("a" + "foo bar")) + "a"), ((("a" + "foo bar") + ("hello" + "foo bar")) + ("\n\t" + "foo bar"))}
    var v2 uint16 = 87
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 uint = 1
    v6 := make(map[int]string)
    switch "\n\t" {
    case "hello":
        v10 := GenS1{F0: -22, F1: 100}
        _ = v10
    case "foo bar":
        switch ("" + "\n\t") {
        case "\n\t":
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := v4[(fi11 - 1)]
                delete(v3, fi11)
                var v13 uint = uint(v2)
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            v2 = (v2 | 882)
        }
    case "a":
        v14 := GenS2{F0: 2147483647, F1: 127}
        _ = v14
    }
    switch 1 {
    case 32768:
        var v15 uint16 = v2
        _ = v15
    case -951:
        if (true || true) {
            var v16 int64 = 42
            fmt.Println(v2, v5)
            _ = v16
        }
    case 127:
        switch "a" {
        case "hello":
            if (!true) {
                v17 := v4[9223372036854775807]
                v2 = v2
                _ = v17
            }
        default:
            fmt.Println(v2, v5)
        }
    }
    var v18 int8 = 127
    fmt.Println(v18, v5, v5)
    _ = len(v4)
    if (!(((3.14159 * (1e10 + 0.0)) >= 3.14159) && true)) {
        v19 := GenS3{F0: 1, F1: -32768}
        v20 := v6[1]
        v5 = 128
        _ = v19
        _ = v20
    }
    v2 = uint16(v5)
    if false {
        v21 := GenS2{F0: -906, F1: -2147483648}
        if false {
            var v22 int8 = v18
            _ = v22
        } else {
            var v23 int64 = 9223372036854775807
            if true {
                if false {
                    var v24 int64 = (^(v23))
                    v18 -= 127
                    _ = len(v1)
                    _ = v24
                }
                fmt.Println(v18, v23, v5)
                var v25 int16 = int16(v5)
                _ = v25
            } else {
                _ = len(v3)
            }
            _ = v23
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v21
    } else {
        for fi26 := 0; fi26 < 14; fi26++ {
            fmt.Println(v18)
            _ = fi26
        }
    }
    var v27 float32 = 3.14159
    _ = len(v4)
    var v28 float32 = v27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
