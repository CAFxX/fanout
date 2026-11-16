package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int64; F1 bool; F2 int64 }
type GenS3 struct { F0 int64; F1 string }

func main() {
    var v1 float64 = 0.1
    var v2 float64 = 1.5
    var v3 float32 = ((((float32(float64(((-2.25 + (((-2.25 + (1.5 - ((1e10 * -2.25) + 0.0))) * ((0.1 - (((0.1 / ((1.5 * (-2.25 - -2.25)) - 3.14159)) + 0.0) / (float32((float64(1.5) * (v2 - (1.5 + 1.5)))) / (3.14159 - (1e10 / 1.5))))) + (0.1 / 1.0))) + 1e10)) - 1e10))) - (0.0 + float32(v1))) * 0.1) * (0.1 / (1e10 * float32((3.14159 - -2.25))))) + ((0.0 * -2.25) / 3.14159))
    var v4 uint8 = 42
    var v5 int64 = -1
    var v6 uint64 = 0
    if true {
        v7 := GenS1{F0: 4294967295, F1: 632}
        _ = v7
    }
    var v8 uint32 = 4294967295
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(v3, v4, fi9)
            var v11 uint16 = 1
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    if false {
        v12 := GenS1{F0: 42, F1: (^(v6))}
        v13 := []int{-9223372036854775808, -86, -9223372036854775808}
        _ = v12
        _ = v13
    }
    switch 513 {
    case -900:
        v14 := []int{2147483647, 127}
        _ = v14
    default:
        var v15 uint8 = (^(v4))
        _ = v15
    }
    switch (("\n\t" + "") + "foo bar") {
    case "hello":
        if true {
            fmt.Println(v5)
            if (913 == v6) {
                var v16 string = (("foo bar" + "hello") + (("\n\t" + ("\n\t" + "hello")) + "\n\t"))
                var v17 int = 492
                _ = v16
                _ = v17
            } else {
                var v18 uint16 = 230
                var v19 int8 = -128
                fmt.Println(v1, v2)
                _ = v18
                _ = v19
            }
            var v20 int32 = 42
            _ = v20
        } else {
            v8 = (v8 ^ (v8 % ((v8) | 1)))
            if (!((v5 >> ((4294967295) & 63)) != (v5 * (v5 & v5)))) {
                v21 := GenS1{F0: 331, F1: 0}
                v4 -= 0
                _ = v21
            }
            v22 := []int{}
            _ = v22
        }
    case "foo bar":
        var v23 int8 = -1
        _ = v23
    }
    v8 = (v8 + (v8 / ((uint32(v6)) | 1)))
    var v24 uint16 = 2
    switch "hello" {
    case "foo bar":
        for fi25 := 0; fi25 < 19; fi25++ {
            v26 := []int{}
            v1 = v2
            _ = fi25
            _ = v26
        }
    case "\n\t":
        switch ("a" + "a") {
        case "\n\t":
            v27 := []int{}
            _ = v27
        case "":
            var v28 uint32 = uint32(v24)
            _ = v28
        }
    default:
        for fi29 := 0; fi29 < 16; fi29++ {
            v30 := GenS2{F0: (v5 / ((int64(v8)) | 1)), F1: false, F2: (v5 ^ -389)}
            _ = fi29
            _ = v30
        }
    }
    var v31 int16 = 0
    if false {
        var v32 uint = 0
        _ = v32
    } else {
        v33 := GenS3{F0: (-(v5)), F1: "foo bar"}
        v4 = v4
        v34 := []int{4294967295, -87}
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
