package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint8; F1 int16; F2 int }
func genFn0(p0 uint8) int8 {
    fmt.Println(p0)
    v6 := []int{}
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 70
}
func genFn1(p0 int64) int {
    p0 = p0
    p0 = (p0 << ((612) & 63))
    var v7 int32 = -493
    _ = p0
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 bool = (0 > 0)
    var v5 int32 = 2147483647
    _ = len(v1)
    switch "a" {
    case "hello":
        if (!((v4 || (v4 || false)) && (false || v4))) {
            fmt.Println(v4, v4)
            v1["a"] = true
            _ = len(v1)
        }
    case "foo bar":
        var v8 string = ("foo bar" + "")
        _ = v8
    case "\n\t":
        var v9 int = -9223372036854775808
        _ = v9
    }
    var v10 string = ""
    if (127 >= 127) {
        var v11 string = v10
        v11 = v10
        v4 = (1 < 311)
        _ = v11
    } else {
        fmt.Println(v4, v5, v10)
    }
    v12 := GenS1{F0: (v10 + "\n\t")}
    v10 = v10
    var v13 uint16 = 1
    for fi14 := 0; fi14 < 14; fi14++ {
        fmt.Println(v13, v10, v5)
        v15 := []int{(fi14 + genFn1(-403)), (-(fi14))}
        var v16 int32 = (v5 % (((-(v5))) | 1))
        _ = fi14
        _ = v15
        _ = v16
    }
    switch v10 {
    case "\n\t":
        if v4 {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 float64 = ((((float64(1.5) / (((0.1 * 1e10) * 0.1) + 3.14159)) + (1.5 * -2.25)) / 0.1) + (((float64(1.5) * float64(-2.25)) - 0.1) - (float64(((1e10 + ((0.0 + float32(0.0)) * 1.5)) / 3.14159)) - 0.0)))
                _ = fi17
                _ = v18
            }
            v19 := []int{128}
            if (!(!(((1e10 - -2.25) >= 3.14159) || (!((!v4) || (-1 >= 9223372036854775807)))))) {
                _ = len(v19)
                v10 = ((v10 + ("" + "\n\t")) + v10)
                v20 := GenS2{F0: -1, F1: v13, F2: 0}
                _ = v20
            }
            _ = v19
        }
    case "hello":
        if (true && (4294967295 != 18446744073709551615)) {
            switch ("foo bar" + v10) {
            case "":
                v21 := GenS1{F0: ""}
                _ = v21
            case "hello":
                var v22 int16 = -32768
                _ = v22
            default:
                v4 = v4
            }
        }
    }
    fmt.Println(v5, v10)
    var v23 int16 = 1000
    v4 = (false || (((v10 + (v10 + ("" + "\n\t"))) <= v10) && v4))
    var v24 float64 = -2.25
    v25 := GenS3{F0: 127, F1: v23, F2: 9223372036854775807}
    fmt.Println(v10, v23)
    v23 = (v23 * v23)
    v1[("\n\t" + "hello")] = v4
    if true {
        var v26 int8 = 127
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
