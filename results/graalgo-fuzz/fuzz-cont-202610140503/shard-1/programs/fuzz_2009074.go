package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint32 = 2147483648
    var v2 float32 = float32((-2.25 - -2.25))
    var v3 uint8 = 255
    var v4 int32 = -865
    var v5 int32 = (v4 * v4)
    var v6 uint32 = (v1 / (((v1 / (((v1 ^ (v1 ^ v1))) | 1))) | 1))
    fmt.Println(v4, v2)
    v1 = 42
    fmt.Println(v4)
    var v7 int8 = 100
    v7 -= (v7 * -128)
    var v8 string = ("hello" + "")
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int = fi9
        fmt.Println(v7, v6, v10)
        if false {
            v3 = uint8(v6)
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 int64 = 9223372036854775807
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 uint8 = 2
                v6 += 2
                if false {
                    var v15 bool = (918 < 337)
                    switch ("" + (("a" + v8) + v8)) {
                    case "foo bar":
                        v10 += -1
                    case "a":
                        var v16 int = fi13
                        _ = v16
                    case "\n\t":
                        v8 = v8
                    default:
                        v17 := []int{v10, (fi13 ^ -137)}
                        _ = v17
                    }
                    var v18 int64 = -9223372036854775808
                    _ = v15
                    _ = v18
                } else {
                    var v19 uint = 18446744073709551615
                    v7 -= -55
                    _ = v19
                }
                _ = fi13
                _ = v14
            }
        } else {
            switch (((v8 + "hello") + v8) + "hello") {
            case "\n\t":
                var v20 int = v10
                _ = v20
            default:
                v21 := []int{128}
                _ = v21
            }
            v22 := []int{1}
            _ = v22
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v2, v7)
    var v23 uint8 = 128
    v5 -= v4
    if (true && (32767 == -840)) {
        v24 := []int{0, 768}
        if (65535 < 255) {
            v8 = "a"
            v25 := GenS1{F0: "hello", F1: v5}
            _ = v25
        }
        _ = v24
    }
    if false {
        var v26 float32 = (((1.5 / v2) + 0.1) + v2)
        if false {
            switch v8 {
            case "hello":
                if (!false) {
                    var v27 uint16 = 367
                    for fi28 := 0; fi28 < 11; fi28++ {
                        v29 := []int{fi28, (fi28 | (-(fi28)))}
                        fmt.Println(v23, v5, v2)
                        v6 -= v1
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                }
            }
        } else {
            v30 := []int{}
            for fi31 := 0; fi31 < 5; fi31++ {
                fmt.Println(v1, v4)
                _ = len(v30)
                _ = fi31
            }
            _ = v30
        }
        v2 += ((0.0 + (v26 - 1.5)) * 1e10)
        _ = v26
    }
    v32 := []int{}
    v33 := GenS2{F0: 761}
    var v34 int64 = 9223372036854775807
    for fi35 := 0; fi35 < 16; fi35++ {
        v36 := GenS1{F0: ("" + v8), F1: v4}
        _ = len(v32)
        fmt.Println(fi35, v8, v7)
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
