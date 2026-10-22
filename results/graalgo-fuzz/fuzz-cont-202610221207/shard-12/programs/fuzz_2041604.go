package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint; F2 string }
func genFn0(p0 string, p1 bool) int32 {
    var v6 uint16 = 811
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{fi7, (fi7 << ((125) & 63))}
        _ = fi7
        _ = v8
    }
    v9 := []int{-2147483648, 0, 0, 127}
    var v10 float32 = (1e10 * float32(-2.25))
    if (p1 && true) {
        if false {
            v11 := GenS1{F0: uint8(v6)}
            _ = v11
        }
        var v12 int32 = 2147483647
        _ = len(v9)
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v9
    return 329
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = 1e10
    var v3 bool = (18446744073709551615 < 18446744073709551615)
    var v4 string = "\n\t"
    v5 := make(map[string]string)
    switch v4 {
    case "a":
        v13 := v5[(v4 + (v4 + "hello"))]
        _ = v13
    case "":
        var v14 int = 2147483647
        _ = v14
    default:
        v3 = v3
    }
    if v3 {
        v15 := GenS2{F0: "hello", F1: 42, F2: 0}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 12; fi16++ {
            if true {
                delete(v5, "hello")
                var v17 int16 = 42
                v18 := GenS1{F0: 127}
                _ = v17
                _ = v18
            } else {
                var v19 uint8 = 195
                _ = v19
            }
            fmt.Println(fi16, v4, v1)
            _ = fi16
        }
        v20 := GenS3{F0: 2, F1: uint(v1), F2: "foo bar"}
        if v3 {
            v3 = true
        } else {
            var v21 float64 = float64(1.5)
            _ = v21
        }
        _ = v20
    }
    fmt.Println(v3, v3)
    v22 := v5[(((v4 + (v4 + v4)) + v4) + (v4 + (v4 + v4)))]
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 uint32 = 0
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 int8 = -128
            var v27 uint = 255
            switch fi25 {
            case 2147483647:
                var v28 uint32 = v24
                _ = v28
            case 32767:
                if (v3 || (128 != 128)) {
                    var v29 int32 = genFn0("foo bar", (!((9223372036854775807 == 1) && ((v26 ^ (v26 % (((v26 * v26)) | 1))) >= (v26 % (((v26 << (((v27 + uint(v24))) & 63))) | 1))))))
                    _ = v29
                } else {
                    v30 := GenS1{F0: 0}
                    v31 := []int{fi25, fi25, (fi23 << ((v27) & 63))}
                    _ = v30
                    _ = v31
                }
            default:
                var v32 string = "hello"
                v2 -= v2
                var v33 uint = (v27 / ((v27) | 1))
                v3 = (v33 >= 205)
                if (928 == -1) {
                    v4 = "hello"
                    v34 := []int{}
                    _ = v34
                } else {
                    v2 += v2
                    v35 := []int{(^(fi25))}
                    _ = v35
                }
                _ = v32
                _ = v33
            }
            var v36 int = 446
            _ = fi25
            _ = v26
            _ = v27
            _ = v36
        }
        _ = fi23
        _ = v24
    }
    v5[v4] = v4
    var v37 int64 = 0
    v38 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
