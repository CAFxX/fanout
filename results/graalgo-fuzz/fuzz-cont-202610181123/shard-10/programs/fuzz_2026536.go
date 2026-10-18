package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 int32) int16 {
    var v5 uint32 = 65535
    switch 127 {
    case -9223372036854775808:
        fmt.Println(v5, v5)
    }
    if true {
        v6 := GenS1{F0: p0, F1: (p0 % (((p0 + (p0 >> ((1) & 63)))) | 1))}
        for fi7 := 0; fi7 < 4; fi7++ {
            if false {
                p0 += p0
            } else {
                var v8 bool = true
                switch "hello" {
                case "a":
                    p0 = p0
                case "\n\t":
                    switch "a" {
                    case "a":
                        v9 := GenS2{F0: fi7}
                        _ = v9
                    default:
                        fmt.Println(fi7, v8, v8)
                    }
                }
                _ = v8
            }
            _ = fi7
        }
        _ = v6
    } else {
        v5 -= (v5 - v5)
    }
    var v10 float32 = (3.14159 * 3.14159)
    _ = p0
    _ = v10
    _ = v5
    return 462
}
func genFn1() int16 {
    if (255 <= 18446744073709551615) {
        if (1 > 2147483648) {
            var v11 uint8 = 0
            v12 := []int{32767, 114}
            _ = v11
            _ = v12
        }
    }
    var v13 bool = false
    if (0 == 2147483647) {
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS2{F0: fi14}
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 uint64 = 489
                var v18 uint8 = 128
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi14
            _ = v15
        }
        fmt.Println(v13)
        for fi19 := 0; fi19 < 17; fi19++ {
            switch "foo bar" {
            case "":
                fmt.Println(fi19)
            }
            var v20 uint16 = 317
            fmt.Println(fi19, v13)
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 float64 = -2.25
        _ = fi21
        _ = v22
    }
    v23 := []int{-94}
    _ = v13
    _ = v23
    return 0
}

func main() {
    v1 := []uint32{597}
    var v2 uint = 690
    v3 := make(map[string]string)
    var v4 uint32 = 65535
    v24 := GenS3{F0: 42, F1: true}
    v25 := GenS1{F0: -151, F1: -1}
    v26 := GenS2{F0: -9223372036854775808}
    var v27 uint32 = (v4 + 1)
    v28 := GenS2{F0: -9223372036854775808}
    v29 := GenS1{F0: 42, F1: 42}
    fmt.Println(v2, v27, v27)
    fmt.Println(v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v30 := GenS2{F0: -148}
    _ = len(v1)
    var v31 uint16 = 2
    v27 = 42
    if ((("hello" + "hello") >= "a") && (0.1 > (0.1 / (float32(3.14159) + 1.5)))) {
        v1 = append(v1, (v4 << (((v2 / ((v2) | 1))) & 63)))
        var v32 int = -881
        _ = v32
    }
    v2 = (v2 % (((^(v2))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
