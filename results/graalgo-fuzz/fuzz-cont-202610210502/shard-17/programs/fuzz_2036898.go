package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 4294967295
    var v3 uint8 = 2
    v4 := v1["foo bar"]
    fmt.Println(v3, v4)
    v4 = v4
    if false {
        v5 := []int{2147483647, -659, 109}
        v4 = (232 <= v3)
        _ = v5
    }
    switch "foo bar" {
    case "foo bar":
        v1["hello"] = (42 != -71)
    default:
        v3 = 19
    }
    v2 -= v2
    v1["hello"] = false
    _ = len(v1)
    switch "hello" {
    case "a":
        for fi6 := 0; fi6 < 19; fi6++ {
            if (v4 || ((393 > 924) && (!true))) {
                var v7 int16 = -32768
                if v4 {
                    v8 := GenS1{F0: 1000, F1: 65535}
                    _ = v8
                }
                if ((v4 && v4) || (v4 && v4)) {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        var v10 uint32 = 4294967295
                        _ = fi9
                        _ = v10
                    }
                    v11 := []int{fi6, 4294967295}
                    for fi12 := 0; fi12 < 2; fi12++ {
                        v13 := GenS2{F0: 2, F1: 862, F2: 1}
                        var v14 int = -99
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                }
                _ = v7
            } else {
                v1["foo bar"] = ((3.14159 <= ((1.5 + float64(-2.25)) - 3.14159)) && (1 < 42))
                var v15 int16 = 0
                _ = v15
            }
            _ = fi6
        }
    default:
        if false {
            _ = len(v1)
            v2 -= 522
            v16 := []int{1}
            _ = v16
        } else {
            for fi17 := 0; fi17 < 16; fi17++ {
                if v4 {
                    fmt.Println(v4)
                } else {
                    if v4 {
                        var v18 int8 = 100
                        _ = v18
                    }
                    v4 = v4
                }
                v19 := GenS3{F0: -1}
                for fi20 := 0; fi20 < 11; fi20++ {
                    v3 += (v3 << (((v2 & (^(v2)))) & 63))
                    var v21 int32 = 2147483647
                    v22 := []int{fi17, -837, fi17, 40}
                    if (226 > 65535) {
                        var v23 int16 = 1000
                        fmt.Println(v21, v3)
                        fmt.Println(v4, v3)
                        _ = v23
                    }
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = fi17
                _ = v19
            }
            v24 := GenS4{F0: 1, F1: 2147483647}
            v25 := GenS2{F0: 4294967295, F1: v2, F2: 9223372036854775807}
            _ = v24
            _ = v25
        }
    }
    v3 = (v3 | (v3 >> (((v2 + 18446744073709551615)) & 63)))
    v26 := GenS4{F0: 108, F1: 2147483647}
    for fi27 := 0; fi27 < 5; fi27++ {
        var v28 uint8 = (^(v3))
        var v29 float64 = 1.5
        v30 := GenS5{F0: -1}
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    var v31 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v4
}
