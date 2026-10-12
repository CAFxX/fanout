package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int16; F2 int16 }
type GenS3 struct { F0 uint; F1 uint; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint; F2 uint }
type GenS5 struct { F0 int16 }
func genFn0() uint {
    var v4 uint = 505
    if true {
        var v5 uint64 = 18446744073709551615
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 uint = (v4 - v4)
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    v4 = (v4 >> ((4294967295) & 63))
    v8 := []int{-990, 803, -2147483648, 32768}
    v8 = append(v8, 4294967295)
    _ = v4
    _ = v8
    return v4
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 614
    var v3 uint32 = v2
    v3 += v2
    if (((((-2.25 - -2.25) + 1.5) - -2.25) <= 1.5) && ((806 != 1) && ((121 > 93) || true))) {
        v2 -= 0
    } else {
        v9 := GenS1{F0: -73, F1: 29}
        if false {
            v3 += (v2 + (v3 % ((42) | 1)))
        }
        fmt.Println(v3, v3)
        _ = v9
    }
    delete(v1, "hello")
    v10 := GenS2{F0: 1, F1: 734, F2: 32767}
    switch 4294967295 {
    case 32767:
        v11 := GenS2{F0: 32767, F1: 477, F2: 264}
        _ = v11
    case 605:
        v12 := []int{-126}
        _ = v12
    case 0:
        switch 128 {
        case 127:
            v2 = 2147483647
        case -1:
            v13 := v1[("hello" + ("" + ("\n\t" + "hello")))]
            _ = v13
        case -9223372036854775808:
            switch 32767 {
            case 0:
                v14 := GenS2{F0: 32767, F1: 42, F2: 684}
                _ = v14
            case 128:
                v15 := GenS3{F0: 255, F1: genFn0(), F2: 127}
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 1; fi16++ {
                v17 := GenS2{F0: 182, F1: -553, F2: -237}
                v2 = v3
                v18 := GenS4{F0: 1, F1: 4294967295, F2: 0}
                _ = fi16
                _ = v17
                _ = v18
            }
        }
    }
    v2 += (v3 + 2147483647)
    var v19 int64 = 909
    switch "" {
    case "":
        if true {
            v20 := []int{32768, 42}
            var v21 uint32 = v3
            _ = v20
            _ = v21
        }
    case "\n\t":
        v19 = (^(v19))
    case "foo bar":
        switch "a" {
        case "\n\t":
            v2 = (^(v3))
        default:
            switch "a" {
            case "":
                v19 = 1
            case "a":
                if (!true) {
                    v22 := []int{-251, 42, 2147483647}
                    switch -254 {
                    case 106:
                        var v23 bool = false
                        _ = v23
                    case 0:
                        fmt.Println(v3)
                    }
                    _ = v22
                } else {
                    v2 = 1
                }
            default:
                var v24 int64 = (v19 / (((v19 | (v19 | 2))) | 1))
                _ = v24
            }
        }
    }
    var v25 string = "\n\t"
    v26 := GenS3{F0: 321, F1: 18446744073709551615, F2: 81}
    var v27 int8 = 0
    v28 := GenS5{F0: -32768}
    v29 := v1[("hello" + "a")]
    v29 -= 2147483647
    for fi30 := 0; fi30 < 12; fi30++ {
        v31 := []int{-2147483648, (v29 - (v29 * (v29 >> ((855) & 63)))), (v29 | -1)}
        v2 += (^(v3))
        _ = fi30
        _ = v31
    }
    var v32 int = 42
    fmt.Println(v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
}
