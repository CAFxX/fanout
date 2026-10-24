package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 string; F1 int; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint32 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 17; fi4++ {
        switch -1 {
        case -1:
            v5 := []int{(^(fi4)), fi4}
            _ = v5
        case 32768:
            var v6 string = "hello"
            _ = v6
        default:
            v7 := GenS1{F0: (1e10 < ((0.0 + (3.14159 * (1.5 + 1e10))) + ((1e10 * 0.0) - (((3.14159 + 1.5) * 1.5) / 3.14159)))), F1: (fi4 - 823)}
            _ = v7
        }
        v8 := GenS2{F0: "foo bar", F1: 1, F2: 1}
        _ = fi4
        _ = v8
    }
    v9 := []int{-2147483648, -2147483648, 2147483647, 42}
    switch 0 {
    case 2147483647:
        if (((((((1e10 / 3.14159) / 1e10) * 0.1) / 1e10) + (1.5 - (((0.1 / 1e10) + (0.1 / (((float32(float64(-2.25)) * 0.1) / 1e10) * (1.5 / ((float32(0.1) + 1.5) + (0.1 / 1.0)))))) - 1e10))) * 1e10) < ((3.14159 / 1.5) * 1e10)) {
            if false {
                var v10 uint16 = 2
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = v10
            }
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 int64 = 42
                _ = fi11
                _ = v12
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    }
    var v13 int64 = -1
    _ = v13
    _ = v9
    return 2147483647
}
func genFn1(p0 uint) int16 {
    var v14 uint16 = 65535
    switch "a" {
    case "hello":
        v15 := GenS2{F0: "foo bar", F1: -2147483648, F2: 42}
        _ = v15
    default:
        var v16 string = (("\n\t" + ("a" + "")) + ("\n\t" + "\n\t"))
        _ = v16
    }
    var v17 float32 = 0.1
    v18 := GenS3{F0: (^(v14)), F1: 255, F2: -128}
    _ = p0
    _ = v14
    _ = v17
    _ = v18
    return 1000
}

func main() {
    var v1 int32 = -2147483648
    var v2 int8 = -54
    v3 := []int{1}
    _ = len(v3)
    if (("hello" + ("\n\t" + (("hello" + "") + "hello"))) >= "") {
        var v19 float64 = float64(0.0)
        var v20 uint8 = 0
        switch "a" {
        case "foo bar":
            if false {
                v21 := GenS1{F0: (!(0 > 429)), F1: 128}
                _ = v21
            }
        }
        _ = v19
        _ = v20
    } else {
        var v22 uint = 18446744073709551615
        if len(v3) > 0 {
            v3[0] = 195
        }
        _ = v22
    }
    v23 := GenS2{F0: ("" + "foo bar"), F1: 42, F2: 18446744073709551615}
    fmt.Println(v2)
    var v24 float32 = 3.14159
    switch 9223372036854775807 {
    case 470:
        _ = len(v3)
    case -918:
        v3 = append(v3, 32768)
    }
    v25 := GenS4{F0: 13, F1: 4294967295}
    fmt.Println(v1, v1)
    switch -700 {
    case 152:
        v3 = append(v3, 128)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
