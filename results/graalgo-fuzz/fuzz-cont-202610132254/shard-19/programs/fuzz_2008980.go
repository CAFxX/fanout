package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint }
func genFn0(p0 int8, p1 int) int8 {
    fmt.Println(p0)
    var v6 int8 = -1
    v7 := GenS1{F0: ((((float64(0.0) * 0.1) + float64((1.5 + 3.14159))) - ((-2.25 + ((1e10 + 1e10) - ((1e10 + (float64((((-2.25 / 1.0) * 0.0) + 1e10)) * 0.0)) - float64(0.0)))) + (float64(0.1) / 1e10))) > 1e10), F1: 381, F2: 325}
    v8 := []int{(^(p1)), (p1 * (p1 & 4294967295)), p1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 100
}
func genFn1(p0 int64) int64 {
    for fi9 := 0; fi9 < 9; fi9++ {
        p0 = (p0 | (p0 ^ p0))
        switch "a" {
        case "hello":
            if false {
                for fi10 := 0; fi10 < 20; fi10++ {
                    if (0 <= 1) {
                        v11 := GenS1{F0: (1000 < 0), F1: 4294967295, F2: 2}
                        v12 := GenS1{F0: true, F1: 503, F2: 0}
                        _ = v11
                        _ = v12
                    }
                    _ = fi10
                }
                fmt.Println(p0)
            }
        case "":
            var v13 int32 = 1
            _ = v13
        default:
            for fi14 := 0; fi14 < 17; fi14++ {
                fmt.Println(p0, fi9)
                v15 := []int{(fi14 ^ 128), fi9, (fi9 ^ (^(fi14)))}
                v16 := GenS1{F0: (65535 > 65535), F1: 1, F2: 1}
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        switch ("" + "\n\t") {
        case "":
            fmt.Println(fi9)
        default:
            if false {
                var v17 string = "\n\t"
                v18 := GenS1{F0: (true || ((-1 == 0) || (true || false))), F1: 2, F2: 2}
                var v19 bool = false
                _ = v17
                _ = v18
                _ = v19
            }
        }
        _ = fi9
    }
    var v20 uint = 65535
    _ = p0
    _ = v20
    return 9223372036854775807
}

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    v3 := []string{"hello", ""}
    var v4 uint64 = 42
    var v5 float32 = float32(3.14159)
    var v21 float32 = (0.0 / float32((1.5 / (3.14159 * float64((0.1 + (v5 + 1e10)))))))
    var v22 float32 = (v5 + (1.5 - 1.5))
    _ = len(v2)
    v23 := v2["\n\t"]
    if len(v3) > 0 {
        v3[0] = v23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        fmt.Println(v1, v1)
        fmt.Println(v4, v23)
        _ = fi24
    }
    var v25 uint64 = 2
    v23 = (("\n\t" + (((((v23 + "a") + (v23 + v23)) + "hello") + v23) + (v23 + "hello"))) + v23)
    if (807 == 9223372036854775807) {
        v26 := GenS1{F0: (1e10 != 0.1), F1: 4294967295, F2: 65535}
        _ = v26
    } else {
        v4 += 0
        fmt.Println(v21, v4)
        fmt.Println(v23, v22)
    }
    fmt.Println(v1, v5, v22)
    switch 4294967295 {
    case 42:
        var v27 int64 = -27
        _ = v27
    case -438:
        var v28 int32 = 42
        _ = v28
    }
    fmt.Println(v21)
    for fi29 := 0; fi29 < 11; fi29++ {
        v1 = v1
        var v30 int16 = 1000
        for fi31 := 0; fi31 < 12; fi31++ {
            var v32 int = (fi29 / (((fi31 / (((^(fi31))) | 1))) | 1))
            v4 -= v25
            fmt.Println(v21, v30, v1)
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v1, v4)
    var v33 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
