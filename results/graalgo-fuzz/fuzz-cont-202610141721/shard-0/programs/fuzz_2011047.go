package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint32; F2 bool }
func genFn0(p0 int8, p1 int8, p2 uint16) string {
    switch -109 {
    case 0:
        switch 32768 {
        case 4294967295:
            var v4 bool = false
            _ = v4
        case 244:
            v5 := []int{700, -819, 1}
            _ = v5
        }
    case -800:
        switch "hello" {
        case "":
            var v6 uint64 = 0
            _ = v6
        case "hello":
            fmt.Println(p0, p1, p0)
        case "a":
            var v7 int8 = p0
            _ = v7
        default:
            var v8 string = "foo bar"
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        p2 = (^(p2))
        _ = fi9
    }
    var v10 float64 = (1e10 * 1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return "hello"
}

func main() {
    v1 := make(map[string]int)
    var v2 string = ""
    var v3 float32 = 3.14159
    var v11 string = "\n\t"
    if ((true || false) && false) {
        v11 = "foo bar"
        var v12 float32 = (v3 + ((v3 / (0.0 + 3.14159)) / float32((((0.0 * float64(v3)) - -2.25) - (1e10 * (float64(v3) - 0.0))))))
        _ = v12
    }
    var v13 int = 32768
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := GenS1{F0: 18}
        var v16 float32 = ((((-2.25 * (v3 / v3)) + v3) - v3) - 1e10)
        v16 = v16
        _ = fi14
        _ = v15
        _ = v16
    }
    v13 -= -509
    _ = len(v1)
    switch (v2 + v2) {
    case "a":
        var v17 uint = 0
        _ = v17
    default:
        v3 += (float32(0.0) - ((v3 + (float32(1e10) / 3.14159)) * (float32(float64(v3)) / v3)))
    }
    var v18 bool = (42 >= 42)
    v19 := GenS2{F0: 65535, F1: 485}
    if ((127 != 65535) && v18) {
        var v20 int16 = 0
        v18 = (!(v18 || v18))
        v2 = "hello"
        _ = v20
    } else {
        v21 := []int{v13, (^(v13)), v13}
        var v22 int = -9223372036854775808
        _ = v21
        _ = v22
    }
    var v23 int64 = 42
    var v24 int64 = v23
    v24 = -1
    switch "a" {
    case "a":
        v25 := []int{(v13 >> ((18446744073709551615) & 63)), (v13 << ((0) & 63))}
        _ = v25
    case "foo bar":
        v26 := GenS3{F0: 0, F1: 65535, F2: false}
        _ = v26
    case "hello":
        if false {
            var v27 uint32 = 42
            _ = v27
        } else {
            var v28 uint64 = 4294967295
            for fi29 := 0; fi29 < 19; fi29++ {
                for fi30 := 0; fi30 < 10; fi30++ {
                    v13 = (-(fi29))
                    var v31 int16 = 0
                    delete(v1, ((v2 + v11) + v2))
                    _ = fi30
                    _ = v31
                }
                switch v2 {
                case "a":
                    v18 = (151 > 86)
                case "hello":
                    var v32 int8 = 26
                    _ = v32
                }
                _ = fi29
            }
            fmt.Println(v3)
            _ = v28
        }
    default:
        v3 -= float32(0.1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
