package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 string; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0() uint {
    var v7 int = 2147483647
    if ("" <= "\n\t") {
        var v8 int32 = 2147483647
        v8 += v8
        if (230 == 2) {
            if ((((0.0 * 3.14159) / ((-2.25 / 1e10) / 1.0)) + (3.14159 * -2.25)) < 0.0) {
                switch "\n\t" {
                case "hello":
                    v9 := []int{(v7 ^ v7), 659, (^(v7))}
                    _ = v9
                }
            }
            if true {
                switch 9223372036854775807 {
                case -965:
                    if false {
                        fmt.Println(v8, v7)
                        var v10 int32 = (v8 ^ 0)
                        _ = v10
                    }
                default:
                    fmt.Println(v7, v8)
                }
                v8 += (v8 * v8)
            }
        }
        _ = v8
    }
    _ = v7
    return 675
}

func main() {
    var v1 int8 = 42
    var v2 uint8 = 241
    var v3 uint32 = 0
    v4 := []int{-811, -934, -326}
    v5 := []int{128, -471, -849, 2147483647}
    var v6 int16 = 1000
    v3 -= 632
    v6 -= (v6 * v6)
    fmt.Println(v6, v6, v6)
    v11 := GenS1{F0: 42}
    v1 = -128
    if ("\n\t" != "") {
        if true {
            switch "" {
            case "hello":
                if true {
                    v12 := GenS1{F0: 784}
                    _ = v12
                }
            }
            var v13 int8 = (v1 - -96)
            _ = v13
        }
    }
    v5 = append(v5, 42)
    for fi14 := 0; fi14 < 3; fi14++ {
        if len(v4) > 0 {
            v4[0] = (fi14 % (((fi14 + fi14)) | 1))
        }
        var v15 int64 = 42
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 329}
    for fi17 := 0; fi17 < 9; fi17++ {
        fmt.Println(v3)
        _ = fi17
    }
    v18 := GenS2{F0: 2, F1: "a", F2: uint64(v2)}
    var v19 string = "hello"
    v20 := GenS3{F0: 1}
    var v21 int32 = 2147483647
    v22 := GenS1{F0: 0}
    if (true && (v6 >= int16(v1))) {
        for fi23 := 0; fi23 < 10; fi23++ {
            v2 = 0
            v5 = append(v5, fi23)
            _ = fi23
        }
        v6 += (v6 & (v6 | v6))
        v24 := GenS1{F0: 352}
        _ = v24
    }
    v25 := GenS4{F0: v3, F1: 18446744073709551615}
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
