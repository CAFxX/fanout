package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int64 }
func genFn0() int32 {
    if true {
        var v5 float32 = ((0.0 / 0.1) / 3.14159)
        _ = v5
    } else {
        if (0.0 <= 0.0) {
            var v6 int = -653
            _ = v6
        } else {
            for fi7 := 0; fi7 < 12; fi7++ {
                v8 := GenS1{F0: 18446744073709551615, F1: ("a" + "hello")}
                var v9 uint64 = 1
                var v10 float32 = (3.14159 / ((1e10 * ((3.14159 + 0.1) - 1.5)) - 1.5))
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            var v11 float32 = (0.0 * (0.1 + (float32(0.0) * ((float32(1e10) * -2.25) / float32(1.5)))))
            _ = v11
        }
        v12 := GenS1{F0: 4294967295, F1: ""}
        v13 := []int{-657, 128}
        _ = v12
        _ = v13
    }
    var v14 string = "a"
    _ = v14
    return -2147483648
}
func genFn1() uint8 {
    if (0 == 42) {
        var v15 uint = 130
        _ = v15
    }
    var v16 uint8 = 1
    _ = v16
    return v16
}

func main() {
    var v1 float64 = (1e10 + 3.14159)
    var v2 int64 = 1
    v3 := make(map[string]int)
    var v4 int64 = (v2 ^ -9223372036854775808)
    if false {
        var v17 int32 = 42
        switch "a" {
        case "\n\t":
            v18 := []int{4294967295}
            _ = v18
        case "foo bar":
            var v19 uint = 4294967295
            _ = v19
        case "":
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 int32 = 1
                fmt.Println(v21, v4, v4)
                _ = fi20
                _ = v21
            }
        default:
            v1 -= (v1 + ((float64(0.0) * (1.5 / v1)) + ((v1 - ((0.1 + ((float64((0.0 + (0.1 / float32(0.1)))) + v1) / 3.14159)) - 0.0)) / float64(0.1))))
        }
        var v22 uint8 = genFn1()
        _ = v17
        _ = v22
    } else {
        v23 := GenS1{F0: 292, F1: "a"}
        fmt.Println(v4, v2, v1)
        _ = v23
    }
    v24 := GenS1{F0: 0, F1: "a"}
    v25 := GenS2{F0: -9223372036854775808}
    if false {
        fmt.Println(v4, v1)
        v26 := []int{2147483647, 2147483647}
        _ = v26
    } else {
        var v27 int8 = 52
        v27 = v27
        _ = v27
    }
    if (-2.25 <= (3.14159 + (float32((3.14159 / (v1 / 1.0))) + (0.1 - -2.25)))) {
        v4 += 0
        v2 = -228
    }
    v28 := []int{996, 32768}
    if false {
        switch ("a" + ("" + (("a" + "a") + ("foo bar" + "a")))) {
        case "a":
            for fi29 := 0; fi29 < 8; fi29++ {
                v4 = -760
                _ = fi29
            }
        case "":
            v30 := GenS1{F0: 18446744073709551615, F1: (("\n\t" + ("\n\t" + ("" + "hello"))) + "a")}
            _ = v30
        default:
            v1 = v1
        }
        v4 += (v2 % (((v2 & v2)) | 1))
    }
    v2 = v2
    var v31 uint16 = 745
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
