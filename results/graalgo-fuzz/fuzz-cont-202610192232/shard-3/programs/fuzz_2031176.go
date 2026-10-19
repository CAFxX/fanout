package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }
func genFn0() uint8 {
    switch -429 {
    case 32768:
        var v4 int64 = 1
        _ = v4
    case 567:
        var v5 uint64 = 4294967295
        _ = v5
    case 127:
        var v6 float32 = float32(float64(1e10))
        _ = v6
    }
    switch "" {
    case "a":
        for fi7 := 0; fi7 < 12; fi7++ {
            if false {
                if ((18446744073709551615 == 1) || (false && (15 <= 194))) {
                    var v8 int64 = 42
                    _ = v8
                }
            } else {
                for fi9 := 0; fi9 < 1; fi9++ {
                    v10 := GenS1{F0: 2, F1: -2147483648, F2: 999}
                    var v11 uint64 = 227
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                fmt.Println(fi7)
                var v12 int64 = -1
                _ = v12
            }
            var v13 float32 = (((1e10 - 0.0) + ((float32(3.14159) - float32(float64(((0.0 / (3.14159 + 0.0)) - -2.25)))) / 1e10)) - 0.1)
            _ = fi7
            _ = v13
        }
    default:
        v14 := GenS1{F0: 255, F1: 82, F2: 255}
        _ = v14
    }
    var v15 bool = false
    _ = v15
    return 134
}

func main() {
    var v1 float32 = 1e10
    var v2 uint32 = 373
    var v3 int = -754
    v1 += 1.5
    v16 := []int{v3, (v3 % (((v3 % (((v3 << ((398) & 63))) | 1))) | 1))}
    v17 := GenS1{F0: 65535, F1: 2147483647, F2: 337}
    var v18 uint = 65535
    var v19 int64 = 42
    if false {
        var v20 uint = v18
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        v1 = (((v1 / float32(0.1)) / (((-2.25 + v1) - float32(3.14159)) * v1)) / v1)
        var v22 int16 = 1000
        _ = len(v16)
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: (^(v18)), F1: 1, F2: 42}
    var v24 int32 = 42
    if (-59 < 73) {
        if len(v16) > 0 {
            v16[0] = v3
        }
    } else {
        var v25 uint = v18
        if true {
            switch "hello" {
            case "a":
                for fi26 := 0; fi26 < 18; fi26++ {
                    fmt.Println(fi26, v24)
                    var v27 float64 = 1.5
                    _ = fi26
                    _ = v27
                }
            }
            var v28 int8 = -128
            v2 = 42
            _ = v28
        } else {
            for fi29 := 0; fi29 < 12; fi29++ {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                switch v3 {
                case 32768:
                    _ = len(v16)
                case 1:
                    if len(v16) > 0 {
                        v16[0] = (v3 - v3)
                    }
                }
                var v30 float64 = -2.25
                _ = fi29
                _ = v30
            }
            var v31 int8 = 100
            var v32 int64 = v19
            _ = v31
            _ = v32
        }
        _ = v25
    }
    var v33 int32 = 2147483647
    _ = len(v16)
    var v34 int32 = -2147483648
    if len(v16) > 0 {
        v16[0] = (^(v3))
    }
    v24 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v34
}
