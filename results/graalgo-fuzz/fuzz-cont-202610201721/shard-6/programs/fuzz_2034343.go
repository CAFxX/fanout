package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint16) uint {
    switch "foo bar" {
    case "hello":
        var v4 uint64 = 1
        _ = v4
    case "":
        p0 = (p0 % ((42) | 1))
    }
    fmt.Println(p0)
    p0 = p0
    var v5 string = ("foo bar" + "a")
    p0 += p0
    _ = p0
    _ = v5
    return 255
}
func genFn1(p0 uint32) int {
    for fi6 := 0; fi6 < 12; fi6++ {
        p0 += (p0 + p0)
        switch 967 {
        case 805:
            for fi7 := 0; fi7 < 1; fi7++ {
                switch ("" + ("hello" + "foo bar")) {
                case "foo bar":
                    fmt.Println(fi7, fi6)
                case "hello":
                    var v8 uint32 = (^(p0))
                    _ = v8
                default:
                    var v9 float64 = (0.0 + (1.5 * 3.14159))
                    v10 := []int{(-(fi6)), 32768, fi6}
                    _ = v10
                    _ = v9
                }
                var v11 uint32 = 2
                var v12 int32 = -1
                _ = fi7
                _ = v11
                _ = v12
            }
        default:
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 float64 = 0.1
                p0 += (p0 / (((p0 | 539)) | 1))
                v15 := GenS1{F0: -32768}
                var v16 string = "a"
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        }
        _ = fi6
    }
    var v17 float64 = -2.25
    switch ("a" + ("" + "\n\t")) {
    case "":
        v18 := []int{-565, 191, -9223372036854775808, 32767}
        _ = v18
    case "a":
        var v19 float64 = 3.14159
        _ = v19
    case "hello":
        var v20 bool = false
        _ = v20
    }
    v21 := []int{-1, -2147483648}
    var v22 float64 = ((v17 + 1e10) - (v17 * v17))
    _ = p0
    _ = v17
    _ = v21
    _ = v22
    return 32768
}

func main() {
    var v1 bool = (("hello" + "") > ("" + ("hello" + "hello")))
    var v2 uint64 = 52
    var v3 uint32 = 2147483648
    if v1 {
        v23 := GenS2{F0: 0, F1: v1}
        _ = v23
    } else {
        for fi24 := 0; fi24 < 11; fi24++ {
            v25 := GenS2{F0: v3, F1: true}
            _ = fi24
            _ = v25
        }
    }
    v3 += v3
    v2 = (v2 & (v2 % ((v2) | 1)))
    v1 = (2 <= 1)
    v26 := GenS3{F0: 253, F1: 42}
    fmt.Println(v1, v3, v1)
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 int32 = 2147483647
        v29 := []int{(fi27 % ((9223372036854775807) | 1)), (^(fi27))}
        var v30 int32 = (v28 ^ int32(fi27))
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    var v31 string = "a"
    v3 -= (v3 * v3)
    fmt.Println(v1, v31)
    var v32 int = -837
    for fi33 := 0; fi33 < 4; fi33++ {
        fmt.Println(v3, fi33, v3)
        var v34 bool = true
        var v35 float64 = (0.0 * (((((float64(-2.25) - ((3.14159 * -2.25) + 3.14159)) + float64(3.14159)) - 1.5) * ((((1e10 / (((1e10 + -2.25) / 1.0) - 1e10)) / 3.14159) * ((-2.25 + 0.1) - -2.25)) / -2.25)) + 0.0))
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := []int{v32, v32}
    for fi37 := 0; fi37 < 15; fi37++ {
        if false {
            v38 := GenS3{F0: 92, F1: 230}
            _ = len(v36)
            _ = v38
        } else {
            v31 = (v31 + v31)
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v36
}
