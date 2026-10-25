package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint32 }
type GenS2 struct { F0 bool; F1 int; F2 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 0
    var v3 uint8 = 42
    v4 := make(map[int]int)
    var v5 float64 = (3.14159 - float64(3.14159))
    if (0 >= 32767) {
        if (128 < -178) {
            switch -2147483648 {
            case 4294967295:
                fmt.Println(v3, v3)
            case 32767:
                var v6 int64 = -9223372036854775808
                _ = v6
            case -354:
                if ((!false) || (!false)) {
                    delete(v1, -9223372036854775808)
                    switch "a" {
                    case "foo bar":
                        _ = len(v4)
                    }
                    var v7 uint8 = 255
                    _ = v7
                }
            }
            switch 9223372036854775807 {
            case 4294967295:
                var v8 float32 = (-2.25 * ((1.5 - ((-2.25 - 1.5) - 3.14159)) + 1.5))
                _ = v8
            default:
                v9 := []int{5, -1, 9223372036854775807}
                _ = v9
            }
            v10 := GenS1{F0: false, F1: -128, F2: 2147483648}
            _ = v10
        } else {
            if true {
                var v11 uint64 = 18446744073709551615
                _ = v11
            }
        }
        var v12 int32 = (v2 / (((v2 | 0)) | 1))
        _ = v12
    } else {
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 uint16 = 65535
            _ = len(v4)
            _ = fi13
            _ = v14
        }
        switch ("" + "a") {
        case "hello":
            v15 := GenS1{F0: ((("a" + ("" + "a")) < "hello") && false), F1: 0, F2: 42}
            _ = v15
        case "a":
            v2 = (^(v2))
        case "foo bar":
            var v16 float32 = float32(float64(((0.0 / -2.25) - ((((1e10 + 1e10) + float32(((float64((0.1 * 1e10)) / 1e10) + (v5 + v5)))) - float32(v5)) / -2.25))))
            _ = v16
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            v18 := []int{fi17, (fi17 & -80), fi17}
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            _ = fi17
            _ = v18
        }
    }
    if false {
        var v19 int32 = 972
        _ = v19
    }
    v20 := GenS2{F0: false, F1: 228, F2: 1}
    delete(v4, 9223372036854775807)
    if (!true) {
        fmt.Println(v2, v2, v3)
        if false {
            if ("" != "foo bar") {
                if false {
                    v5 -= -2.25
                    v3 = v3
                } else {
                    if (0 >= 42) {
                        v21 := GenS1{F0: true, F1: 1, F2: 2147483648}
                        _ = v21
                    }
                    var v22 float64 = 1.5
                    _ = v22
                }
            }
            if (true || (!true)) {
                switch "\n\t" {
                case "":
                    v23 := v1[42]
                    _ = v23
                default:
                    v24 := []int{-34, 42, -9223372036854775808, 2147483647}
                    _ = v24
                }
                v25 := GenS1{F0: false, F1: -1, F2: 2147483647}
                _ = v25
            } else {
                v5 += 3.14159
                v26 := GenS3{F0: 2}
                _ = v26
            }
            fmt.Println(v3, v5)
        }
    }
    var v27 int16 = -837
    if true {
        var v28 float32 = float32((float64(float32(v5)) + 1e10))
        var v29 int8 = -128
        _ = v28
        _ = v29
    } else {
        v5 += -2.25
        v3 = 127
    }
    var v30 int16 = (v27 | (v27 | 683))
    var v31 int64 = 295
    v32 := GenS4{F0: ((("foo bar" + ("" + ("foo bar" + ("" + "foo bar")))) + ("\n\t" + ((("hello" + ("\n\t" + "a")) + "foo bar") + "\n\t"))) + ("" + "")), F1: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
