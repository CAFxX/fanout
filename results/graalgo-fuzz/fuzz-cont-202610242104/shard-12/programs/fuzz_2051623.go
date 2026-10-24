package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int8 }
func genFn0() int {
    switch -419 {
    case 32768:
        if false {
            var v4 int32 = 0
            switch -475 {
            case 227:
                if (-128 <= -58) {
                    var v5 int8 = 100
                    var v6 uint = 977
                    _ = v5
                    _ = v6
                }
            case 4294967295:
                v7 := []int{-9223372036854775808, 977}
                _ = v7
            }
            v8 := GenS1{F0: 2, F1: 128}
            _ = v4
            _ = v8
        } else {
            var v9 uint8 = 247
            v10 := []int{0, 127, -190, 632}
            var v11 uint32 = 2147483648
            _ = v10
            _ = v11
            _ = v9
        }
    case 4294967295:
        v12 := GenS2{F0: 2147483647, F1: 69}
        _ = v12
    }
    v13 := []int{}
    if (!((((1 == 42) && (!(0.0 == (-2.25 * 3.14159)))) && (((-2147483648 != -1) && (3.14159 < 1.5)) && (!false))) && (false || (4294967295 != 127)))) {
        if len(v13) > 0 {
            v13[0] = 9223372036854775807
        }
    }
    _ = v13
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 bool = (18446744073709551615 != 18446744073709551615)
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            v3 = v3
            fmt.Println(fi15)
            _ = fi15
        }
        _ = fi14
    }
    if true {
        switch -14 {
        case 32767:
            var v16 uint = 1
            _ = v16
        case 127:
            v1[(("a" + "") + "foo bar")] = "a"
        case 128:
            switch 513 {
            case 130:
                if (!((0.1 / 1e10) > 1e10)) {
                    var v17 float32 = (float32(0.1) - 3.14159)
                    _ = v17
                } else {
                    v3 = v3
                    v3 = v3
                    v3 = (!(489 >= 9223372036854775807))
                }
            }
        default:
            if false {
                var v18 uint = 639
                var v19 uint16 = 0
                if (!v3) {
                    v20 := GenS1{F0: 2147483648, F1: 0}
                    _ = v20
                }
                _ = v18
                _ = v19
            }
        }
        switch "a" {
        case "a":
            v1[""] = "a"
        }
    }
    v3 = v3
    var v21 int8 = -128
    fmt.Println(v3, v3)
    var v22 string = ""
    v3 = false
    var v23 int8 = v21
    v24 := GenS3{F0: 0}
    fmt.Println(v23, v3)
    v25 := GenS1{F0: 2147483648, F1: 162}
    var v26 uint = 42
    fmt.Println(v3)
    for fi27 := 0; fi27 < 1; fi27++ {
        if v3 {
            var v28 int32 = 2147483647
            _ = v28
        }
        _ = fi27
    }
    v26 += (v26 * (v26 >> (((v26 * v26)) & 63)))
    _ = len(v1)
    switch 42 {
    case -1:
        var v29 int64 = 1
        _ = v29
    case -652:
        for fi30 := 0; fi30 < 20; fi30++ {
            switch "" {
            case "a":
                fmt.Println(fi30)
            case "":
                fmt.Println(v3, v26)
            case "\n\t":
                v31 := v2[""]
                _ = v31
            default:
                _ = len(v1)
            }
            v21 = (v23 & v23)
            _ = fi30
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
