package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int16; F1 int }

func main() {
    var v1 int32 = -1
    var v2 float64 = (1.5 - (((float64(-2.25) * (3.14159 / (float64((0.0 + 3.14159)) / ((1.5 - 1e10) / float64(1.5))))) - (1.5 * 0.0)) + -2.25))
    var v3 int8 = 1
    var v4 uint64 = 4294967295
    v5 := []int{-1, -9223372036854775808, 2147483647}
    for fi6 := 0; fi6 < 5; fi6++ {
        v5 = append(v5, fi6)
        if true {
            v2 = v2
            v7 := GenS1{F0: (255 > 224)}
            _ = v7
        }
        _ = fi6
    }
    if (((!(2 != 0)) && (0 >= int(v3))) || false) {
        if true {
            var v8 string = ""
            v2 += ((float64(3.14159) / (v2 - v2)) + v2)
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 float32 = 3.14159
                _ = len(v5)
                var v11 string = v8
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v8
        } else {
            var v12 int16 = 1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v13 int32 = -930
            _ = v12
            _ = v13
        }
    }
    v3 = (^(v3))
    for fi14 := 0; fi14 < 10; fi14++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            v2 += float64(float32(float64(-2.25)))
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v5 = append(v5, fi15)
            _ = fi15
        }
        var v16 string = "a"
        _ = fi14
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v3 -= -1
        _ = fi17
    }
    var v18 uint64 = (v4 / (((^(v4))) | 1))
    v5 = append(v5, 128)
    for fi19 := 0; fi19 < 12; fi19++ {
        v20 := GenS1{F0: (0 >= 795)}
        v18 = 560
        _ = fi19
        _ = v20
    }
    if true {
        v21 := GenS1{F0: ((0.1 >= (1.5 * -2.25)) || false)}
        v4 += 2
        v2 = v2
        _ = v21
    }
    v22 := GenS2{F0: ""}
    for fi23 := 0; fi23 < 9; fi23++ {
        v5 = append(v5, (fi23 | fi23))
        fmt.Println(v1, fi23, v2)
        _ = fi23
    }
    switch "\n\t" {
    case "hello":
        if true {
            var v24 int32 = -2147483648
            var v25 bool = (!((!false) && (-1 >= v3)))
            var v26 uint = 996
            _ = v24
            _ = v25
            _ = v26
        }
    default:
        v27 := GenS3{F0: 2, F1: -558}
        _ = v27
    }
    if (!false) {
        v5 = append(v5, -1)
        var v28 uint16 = 65535
        if true {
            v3 = v3
            v1 = 277
        }
        _ = v28
    } else {
        v3 = 100
        var v29 uint32 = 834
        if len(v5) > 0 {
            v5[0] = 39
        }
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        if (("hello" + (("a" + "hello") + "a")) >= "hello") {
            v31 := GenS4{F0: (v3 & v3), F1: 2}
            _ = v31
        }
        var v32 int8 = -128
        if (128 != 128) {
            fmt.Println(v32, v32, fi30)
            v32 += (^(v32))
        }
        _ = fi30
        _ = v32
    }
    fmt.Println(v2, v18)
    v33 := GenS5{F0: 42, F1: 9223372036854775807}
    for fi34 := 0; fi34 < 1; fi34++ {
        var v35 uint8 = 2
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
