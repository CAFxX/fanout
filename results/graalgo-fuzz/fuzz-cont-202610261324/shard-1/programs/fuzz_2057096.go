package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
func genFn0(p0 uint16, p1 int32) string {
    var v4 int = 4294967295
    var v5 uint32 = 2147483648
    switch "foo bar" {
    case "hello":
        var v6 uint32 = v5
        _ = v6
    case "a":
        v7 := []int{}
        _ = v7
    }
    v5 += (v5 * 2)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return ("\n\t" + "")
}

func main() {
    var v1 int16 = 101
    v2 := make(map[string]int)
    v3 := []int{-9223372036854775808, -9223372036854775808, -315, 2147483647}
    v3 = append(v3, 32767)
    var v8 uint16 = 1
    switch -2147483648 {
    case 0:
        switch -9223372036854775808 {
        case 0:
            v1 = -307
        case 530:
            fmt.Println(v1, v1)
        case 42:
            var v9 int = 128
            _ = v9
        default:
            if true {
                for fi10 := 0; fi10 < 12; fi10++ {
                    for fi11 := 0; fi11 < 12; fi11++ {
                        v8 = v8
                        var v12 int32 = -952
                        _ = len(v3)
                        _ = fi11
                        _ = v12
                    }
                    _ = fi10
                }
                v13 := GenS1{F0: 247, F1: -1}
                _ = v13
            } else {
                var v14 uint = 4294967295
                _ = v14
            }
        }
    default:
        var v15 int64 = 809
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 int16 = int16(fi16)
        var v18 int = -1
        var v19 float64 = (((1e10 * 3.14159) * (float64(0.0) - (float64(3.14159) * (float64(1e10) + ((((((-2.25 / -2.25) / 1e10) / float64(1.5)) * 0.0) + 3.14159) + 3.14159))))) - 1.5)
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 int8 = 42
    v20 = v20
    v20 = (v20 / ((47) | 1))
    var v21 int16 = 42
    if (((v20 - (v20 | v20)) != 1) || true) {
        if (!false) {
            v22 := GenS1{F0: 4294967295, F1: 42}
            if (2147483647 == 401) {
                if false {
                    fmt.Println(v20)
                }
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 int32 = 0
                    _ = fi23
                    _ = v24
                }
            } else {
                if true {
                    fmt.Println(v1)
                    v25 := GenS1{F0: 18446744073709551615, F1: 1}
                    _ = v25
                }
            }
            var v26 string = "hello"
            _ = v22
            _ = v26
        }
        var v27 float32 = (3.14159 / float32(1.5))
        _ = v27
    }
    fmt.Println(v1, v21)
    if (v1 > (v1 - 42)) {
        _ = len(v2)
        fmt.Println(v1, v1)
        v28 := GenS2{F0: 128}
        _ = v28
    }
    v29 := GenS1{F0: 42, F1: 540}
    v20 = (v20 - (v20 ^ 1))
    var v30 int8 = v20
    var v31 int32 = -1
    var v32 uint16 = (v8 * 32767)
    v33 := GenS3{F0: 414}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v8
}
