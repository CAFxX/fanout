package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 bool = (!(0.1 == -2.25))
    var v2 uint = 406
    v3 := make(map[int]string)
    var v4 int64 = 0
    if false {
        for fi5 := 0; fi5 < 6; fi5++ {
            for fi6 := 0; fi6 < 20; fi6++ {
                v7 := GenS1{F0: v1, F1: 90}
                var v8 int64 = (v4 >> (((^(v2))) & 63))
                v2 -= (v2 - (v2 / ((419) | 1)))
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = fi5
        }
    }
    v4 += (-(v4))
    fmt.Println(v2, v4)
    v9 := []int{}
    v4 = v4
    v10 := GenS1{F0: v1, F1: 970}
    v11 := GenS2{F0: 0}
    var v12 float64 = -2.25
    for fi13 := 0; fi13 < 18; fi13++ {
        v4 = v4
        v2 += 18446744073709551615
        var v14 string = "\n\t"
        _ = fi13
        _ = v14
    }
    var v15 int = 9223372036854775807
    for fi16 := 0; fi16 < 14; fi16++ {
        fmt.Println(v2)
        _ = fi16
    }
    v2 -= 18446744073709551615
    if v1 {
        var v17 uint32 = 65535
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v17
    } else {
        var v18 uint16 = 814
        v19 := GenS2{F0: 1000}
        var v20 uint16 = v18
        _ = v18
        _ = v19
        _ = v20
    }
    v12 += ((v12 - float64(float32(float64((0.0 * 0.1))))) + 1.5)
    switch (v15 >> ((2) & 63)) {
    case -9223372036854775808:
        for fi21 := 0; fi21 < 1; fi21++ {
            if true {
                var v22 int32 = -92
                if len(v9) > 0 {
                    v9[0] = fi21
                }
                var v23 float32 = 3.14159
                _ = v22
                _ = v23
            } else {
                v2 = (^(v2))
            }
            fmt.Println(v15, v12)
            _ = fi21
        }
    case -2147483648:
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 uint8 = 154
            for fi26 := 0; fi26 < 6; fi26++ {
                v27 := GenS1{F0: (!(1.5 == 1.5)), F1: 4294967295}
                v15 = fi26
                var v28 string = ("a" + (("foo bar" + "a") + (("hello" + ("a" + ("foo bar" + ("hello" + "hello")))) + "hello")))
                if v1 {
                    var v29 string = ((v28 + v28) + v28)
                    v2 += 65535
                    if (121 > -1) {
                        v30 := GenS3{F0: 42}
                        _ = v30
                    }
                    _ = v29
                } else {
                    var v31 uint = v2
                    _ = v31
                }
                _ = fi26
                _ = v27
                _ = v28
            }
            fmt.Println(fi24, v2, v25)
            _ = fi24
            _ = v25
        }
    case -780:
        fmt.Println(v1)
    default:
        _ = len(v9)
    }
    v9 = append(v9, 642)
    var v32 float64 = (0.0 + (v12 / v12))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v9
}
