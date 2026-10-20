package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint }
func genFn0(p0 int16, p1 int64, p2 uint64) bool {
    var v4 string = ("a" + "foo bar")
    var v5 uint32 = 0
    v6 := []int{}
    p0 += p0
    v7 := GenS1{F0: "\n\t", F1: ((v5 >> ((2) & 63)) >= 2147483648), F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (1 >= 32767)
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]int)
    var v3 uint8 = 42
    fmt.Println(v1, v3)
    var v8 float32 = v1
    if false {
        for fi9 := 0; fi9 < 5; fi9++ {
            v8 -= v8
            _ = fi9
        }
        v10 := []int{}
        _ = v10
    } else {
        fmt.Println(v1, v1, v3)
    }
    var v11 uint64 = 2
    var v12 uint64 = (v11 + 0)
    var v13 float64 = -2.25
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 string = ("" + "hello")
        _ = fi14
        _ = v15
    }
    var v16 int32 = 513
    v13 = (((float64((v1 * (v1 / ((0.0 / (float32((v13 - 0.0)) + (((-2.25 / (v8 * ((v8 - v1) * v1))) * (v8 * 3.14159)) * (v1 - ((((v8 * 1e10) - (v8 * v8)) - (1e10 + v1)) * (float32(v13) * v8)))))) / v8)))) * float64(0.1)) - ((((((v13 / (3.14159 / ((1e10 - float64(3.14159)) - 1.5))) + (float64(0.0) + v13)) + ((v13 + v13) * (0.1 * v13))) - 0.0) / (v13 + float64(v8))) + 1e10)) * (float64(3.14159) * (v13 + v13)))
    var v17 int32 = v16
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 uint32 = 64
        v20 := GenS1{F0: "hello", F1: true, F2: 255}
        v17 += (v16 / (((v16 | v17)) | 1))
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := []int{-668, 42, 42}
    if true {
        v12 -= v12
    } else {
        var v22 float32 = ((3.14159 / (v1 - v1)) + (v8 - 0.0))
        if true {
            var v23 string = "foo bar"
            v24 := GenS1{F0: v23, F1: true, F2: 615}
            if (("foo bar" + "hello") != ("\n\t" + v23)) {
                switch 128 {
                case 480:
                    v25 := GenS1{F0: v23, F1: (v23 == (("" + v23) + (v23 + v23))), F2: 1}
                    _ = v25
                default:
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                }
                var v26 string = ("a" + "hello")
                v12 += 4294967295
                _ = v26
            } else {
                v27 := GenS1{F0: v23, F1: (false || (0 > 2147483647)), F2: 192}
                _ = v27
            }
            _ = v23
            _ = v24
        } else {
            fmt.Println(v16)
            v11 -= v12
            var v28 int16 = 32767
            _ = v28
        }
        v22 += (v8 * (v1 / v22))
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v8
}
