package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 string; F2 int32 }
func genFn0() uint8 {
    if false {
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 uint = 513
            for fi7 := 0; fi7 < 13; fi7++ {
                v6 = v6
                _ = fi7
            }
            _ = fi5
            _ = v6
        }
        if false {
            if (479 != 659) {
                v8 := []int{}
                var v9 float32 = ((1.5 + float32(0.1)) / -2.25)
                _ = v8
                _ = v9
            } else {
                if ((-1 == 0) && true) {
                    switch "foo bar" {
                    case "hello":
                        v10 := []int{-14, 127}
                        _ = v10
                    case "foo bar":
                        v11 := []int{-1, 127, 128}
                        _ = v11
                    default:
                        var v12 uint16 = 2
                        _ = v12
                    }
                    var v13 int64 = 1
                    _ = v13
                }
                var v14 uint8 = 1
                _ = v14
            }
            var v15 bool = (!(("\n\t" > "") && true))
            v16 := []int{-1, 0}
            _ = v15
            _ = v16
        }
    }
    var v17 uint8 = 187
    v18 := []int{-2147483648, -9223372036854775808, 1, 311}
    v19 := GenS1{F0: 32767}
    _ = v17
    _ = v18
    _ = v19
    return (v17 - 0)
}

func main() {
    var v1 uint8 = 0
    var v2 uint8 = v1
    v3 := make(map[string]bool)
    var v4 uint8 = 186
    switch -69 {
    case 4294967295:
        var v20 uint = 947
        _ = v20
    case 0:
        var v21 uint32 = 484
        _ = v21
    case 32767:
        v4 += (^(v4))
    default:
        var v22 float32 = 3.14159
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        delete(v3, "a")
        v2 = (^(v4))
        _ = fi23
    }
    fmt.Println(v2, v2)
    v24 := GenS1{F0: 42}
    if ((42 >= -1) || (0.0 < 0.0)) {
        v1 = 128
    }
    if (!true) {
        fmt.Println(v4)
        var v25 float64 = 1e10
        _ = len(v3)
        _ = v25
    } else {
        var v26 uint32 = 4294967295
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 int8 = 127
            v29 := []int{-2147483648}
            _ = fi27
            _ = v28
            _ = v29
        }
        v4 += (v1 / (((v4 - (^(v4)))) | 1))
        _ = v26
    }
    v30 := GenS2{F0: 535}
    var v31 int8 = -128
    v2 += (v1 - (v2 / (((v4 | 127)) | 1)))
    v32 := []int{-9223372036854775808}
    v33 := GenS3{F0: uint32(v2), F1: (("\n\t" + "foo bar") + ""), F2: -435}
    var v34 float32 = (((1e10 + (0.0 + (1e10 - (1e10 + float32(-2.25))))) / 1.5) * -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
