package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int64; F1 int64; F2 uint32 }
func genFn0(p0 string, p1 int32, p2 string) int32 {
    switch ("a" + ("\n\t" + ("foo bar" + "\n\t"))) {
    case "hello":
        v4 := []int{}
        _ = v4
    case "\n\t":
        p2 = ("a" + p2)
    default:
        var v5 float64 = float64(0.0)
        _ = v5
    }
    switch 32767 {
    case 127:
        p0 = p2
    case 252:
        if true {
            var v6 uint8 = 0
            v7 := GenS1{F0: 1, F1: 0}
            _ = v6
            _ = v7
        }
    case 458:
        if true {
            v8 := []int{127}
            if len(v8) > 0 {
                v8[0] = -1
            }
            _ = v8
        } else {
            var v9 float64 = ((3.14159 / 0.1) / 1.5)
            fmt.Println(p0, p1)
            p1 -= p1
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 6; fi10++ {
            if ((2147483647 <= 85) && true) {
                p0 = ((("foo bar" + (("\n\t" + ("foo bar" + ("\n\t" + ""))) + p0)) + p2) + (p0 + "\n\t"))
                var v11 float64 = (1.5 - 3.14159)
                if (18446744073709551615 != 2) {
                    v12 := []int{(fi10 * fi10), 32768}
                    _ = v12
                } else {
                    var v13 uint8 = 2
                    _ = v13
                }
                _ = v11
            }
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1(p0 uint16, p1 string, p2 uint64) uint {
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := GenS1{F0: -14, F1: -1}
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 0, F1: -1, F2: 399}
    v17 := []int{127, 128, 645, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 1
}

func main() {
    var v1 float32 = 1.5
    v2 := []int{42, -9223372036854775808}
    var v3 uint64 = 0
    fmt.Println(v1)
    v3 = 18446744073709551615
    fmt.Println(v3, v3)
    switch ("hello" + "") {
    case "":
        v3 = (v3 + v3)
    case "hello":
        var v18 int16 = 0
        _ = v18
    }
    var v19 int8 = 41
    v1 = (v1 * float32((((-2.25 + 1e10) - 3.14159) + (-2.25 / float64(v1)))))
    switch "a" {
    case "a":
        v20 := GenS2{F0: -386, F1: 1, F2: 2147483647}
        _ = v20
    case "hello":
        v19 += (v19 / (((v19 - v19)) | 1))
    default:
        var v21 int8 = 123
        _ = v21
    }
    v19 = v19
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3)
    _ = len(v2)
    var v22 int16 = -1
    for fi23 := 0; fi23 < 10; fi23++ {
        switch "hello" {
        case "hello":
            var v24 float32 = 0.1
            _ = v24
        case "foo bar":
            var v25 uint8 = 0
            _ = v25
        }
        _ = fi23
    }
    if len(v2) > 0 {
        v2[0] = 151
    }
    fmt.Println(v22, v22, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
}
