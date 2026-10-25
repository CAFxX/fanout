package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 string; F1 uint16; F2 int16 }
type GenS3 struct { F0 int64; F1 string; F2 uint16 }
func genFn0() int64 {
    v4 := []int{-351, 2147483647, 215, -2147483648}
    var v5 string = "foo bar"
    v5 = v5
    for fi6 := 0; fi6 < 5; fi6++ {
        switch fi6 {
        case -1:
            v5 = v5
        case 32768:
            for fi7 := 0; fi7 < 20; fi7++ {
                if (!(0 < 0)) {
                    v8 := GenS1{F0: (fi7 * fi6), F1: "foo bar"}
                    v5 = v5
                    switch v5 {
                    case "hello":
                        v9 := GenS2{F0: "\n\t", F1: 65535, F2: -857}
                        _ = v9
                    case "foo bar":
                        var v10 uint32 = 65535
                        _ = v10
                    default:
                        var v11 uint32 = 4294967295
                        _ = v11
                    }
                    _ = v8
                }
                var v12 int8 = -71
                _ = fi7
                _ = v12
            }
        case 260:
            var v13 float32 = (1.5 / (((0.1 * ((1e10 / 3.14159) - 1e10)) - 0.0) - (float32((float64((3.14159 * 1e10)) + 3.14159)) + float32((3.14159 / 1e10)))))
            _ = v13
        default:
            v4 = append(v4, fi6)
        }
        _ = fi6
    }
    v5 = (v5 + v5)
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    var v14 float32 = 1e10
    v1 = (44 >= 42)
    var v15 int = 32768
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 string = ("hello" + "foo bar")
        _ = fi16
        _ = v17
    }
    fmt.Println(v1, v1)
    if (65535 >= 0) {
        v18 := []int{(v15 / (((v15 << ((54) & 63))) | 1)), v15, v15, v15}
        v19 := GenS2{F0: "\n\t", F1: 0, F2: -613}
        _ = v18
        _ = v19
    }
    var v20 bool = v1
    var v21 string = "hello"
    v22 := GenS2{F0: v21, F1: 32767, F2: 1000}
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint = 4294967295
        _ = fi23
        _ = v24
    }
    var v25 float64 = 1.5
    if (!(2147483647 != 0)) {
        var v26 int8 = 45
        _ = v26
    }
    var v27 int = 32768
    v20 = v1
    switch v21 {
    case "":
        v28 := GenS3{F0: -9223372036854775808, F1: (v21 + v21), F2: 1}
        _ = v28
    case "\n\t":
        v29 := []int{v27}
        _ = v29
    default:
        v15 = 32768
    }
    v30 := GenS1{F0: (v15 & (v15 | (v27 + (v27 ^ v15)))), F1: ("hello" + "\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v27
    _ = v3
    _ = v30
}
