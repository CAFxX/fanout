package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint64 = 1
    var v2 uint16 = 42
    var v3 int = -443
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int64 = 245
        _ = fi4
        _ = v5
    }
    if false {
        v6 := GenS1{F0: (v1 % ((v1) | 1))}
        var v7 int8 = 100
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 1}
        fmt.Println(v3)
        _ = v8
    }
    switch v3 {
    case -267:
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := []int{(^(v3)), (v3 << ((uint(v2)) & 63)), (fi9 + -476)}
            for fi11 := 0; fi11 < 19; fi11++ {
                if true {
                    if false {
                        fmt.Println(v1, v1, fi9)
                    } else {
                        fmt.Println(fi9)
                        v2 = (v2 - v2)
                        v2 = 542
                        fmt.Println(v3, v1)
                    }
                    if (false && false) {
                        v3 = fi11
                    }
                } else {
                    switch v3 {
                    case 4294967295:
                        var v12 int = fi11
                        _ = v12
                    case -212:
                        v13 := GenS1{F0: v1}
                        _ = v13
                    default:
                        v10 = append(v10, fi9)
                    }
                }
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
    case 94:
        v14 := []int{v3}
        _ = v14
    default:
        fmt.Println(v1, v2)
    }
    if false {
        if ((-2.25 - float32(0.1)) != (((-2.25 + ((1e10 * (0.1 * 3.14159)) + 0.0)) - (-2.25 + 0.1)) - 3.14159)) {
            switch "\n\t" {
            case "\n\t":
                v3 = v3
            case "foo bar":
                if ((true || (true && false)) && true) {
                    v15 := []int{v3, -191}
                    v15 = append(v15, (v3 >> ((1) & 63)))
                    if true {
                        var v16 int8 = 0
                        var v17 float64 = (0.0 + (1e10 / 1e10))
                        fmt.Println(v3)
                        _ = v16
                        _ = v17
                    }
                    _ = v15
                }
            case "":
                var v18 int64 = 560
                _ = v18
            default:
                var v19 float32 = (0.1 - (float32(-2.25) - 0.1))
                _ = v19
            }
            v20 := GenS1{F0: 4294967295}
            _ = v20
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 float64 = 3.14159
            v1 -= (v1 * (^(v1)))
            _ = fi21
            _ = v22
        }
    } else {
        v23 := GenS2{F0: ("\n\t" + ("a" + "hello"))}
        _ = v23
    }
    v24 := GenS2{F0: "a"}
    if false {
        fmt.Println(v2)
        fmt.Println(v3, v2)
    } else {
        v3 -= (v3 - 32767)
    }
    if (1 != (v2 << ((1) & 63))) {
        var v25 bool = false
        _ = v25
    }
    if (18446744073709551615 >= 65535) {
        for fi26 := 0; fi26 < 5; fi26++ {
            for fi27 := 0; fi27 < 20; fi27++ {
                v28 := []int{}
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
        v29 := []int{-851, 9223372036854775807, (^(v3))}
        for fi30 := 0; fi30 < 1; fi30++ {
            fmt.Println(fi30, v2)
            switch v3 {
            case -892:
                var v31 uint8 = 121
                _ = v31
            case -1:
                v2 = v2
            default:
                v2 -= (v2 / ((2) | 1))
            }
            var v32 int = (v3 % ((32768) | 1))
            _ = fi30
            _ = v32
        }
        _ = v29
    } else {
        if (float64(0.0) != 0.0) {
            var v33 bool = (2 > 42)
            _ = v33
        }
        v2 += v2
    }
    v3 += v3
    v34 := []int{(v3 ^ v3)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
}
