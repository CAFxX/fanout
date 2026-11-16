package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 bool }
func genFn0(p0 uint64, p1 bool) uint32 {
    var v7 uint64 = (p0 >> ((42) & 63))
    v8 := GenS1{F0: 1, F1: 9223372036854775807}
    v7 -= (^(v7))
    p0 = v7
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 string, p1 int32, p2 uint32) int32 {
    v9 := []int{}
    var v10 float32 = ((1.5 - 1e10) - (1e10 * (0.1 / -2.25)))
    v11 := GenS1{F0: -128, F1: -2147483648}
    if true {
        if len(v9) > 0 {
            v9[0] = 32768
        }
    }
    fmt.Println(v10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return (p1 << ((42) & 63))
}

func main() {
    var v1 string = ""
    v2 := []uint32{2147483648}
    var v3 uint8 = 128
    var v4 bool = true
    var v5 float64 = (-2.25 / 1.5)
    var v6 string = v1
    v12 := GenS1{F0: 0, F1: 42}
    v13 := GenS1{F0: 127, F1: 42}
    switch "" {
    case "\n\t":
        v5 -= v5
    case "":
        var v14 int = 9223372036854775807
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS1{F0: 51, F1: fi15}
        _ = fi15
        _ = v16
    }
    var v17 string = v6
    switch ("" + v1) {
    case "":
        if v4 {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v1, v6)
            fmt.Println(v17)
        }
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 string = v17
        if len(v2) > 0 {
            v2[0] = 42
        }
        _ = fi18
        _ = v19
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v20 uint32 = 0
    v21 := GenS2{F0: -550}
    for fi22 := 0; fi22 < 7; fi22++ {
        v23 := GenS3{F0: (((((1e10 / 1.5) * 0.0) * 0.0) * 1.5) > (((3.14159 / 1.0) - float32(v5)) / (1.5 + 1e10))), F1: (-9223372036854775808 == 0)}
        for fi24 := 0; fi24 < 13; fi24++ {
            fmt.Println(fi24)
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    if (-1 != 127) {
        v5 = v5
        for fi25 := 0; fi25 < 5; fi25++ {
            for fi26 := 0; fi26 < 6; fi26++ {
                v1 = ""
                switch v17 {
                case "":
                    v3 = 128
                case "hello":
                    var v27 uint = 730
                    _ = v27
                case "a":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    var v28 uint64 = 207
                    _ = v28
                }
                if v4 {
                    v6 = ("a" + v17)
                    var v29 float64 = float64(-2.25)
                    _ = v29
                } else {
                    v5 += v5
                    v5 += ((v5 - ((-2.25 - (v5 - float64(((((1.5 + float32(0.1)) + ((3.14159 / (1e10 + 3.14159)) * (float32(v5) / 1.0))) - ((0.1 / (-2.25 + ((0.0 + 3.14159) * (1.5 / 1e10)))) * (((0.1 - -2.25) + (3.14159 / float32(float64(float32(v5))))) * (1e10 / (float32(float64(3.14159)) + -2.25))))) + -2.25)))) + ((float64(0.1) + 1e10) + float64(1e10)))) * v5)
                }
                _ = fi26
            }
            for fi30 := 0; fi30 < 1; fi30++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi30
            }
            _ = fi25
        }
    } else {
        v1 = v1
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        switch (fi31 - int(v3)) {
        case 127:
            v20 -= 272
        }
        _ = fi31
    }
    v3 = 1
    var v32 int8 = -8
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
