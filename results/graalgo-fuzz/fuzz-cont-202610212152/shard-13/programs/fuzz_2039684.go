package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int; F1 string; F2 uint64 }

func main() {
    var v1 int16 = 1
    var v2 float64 = -2.25
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    var v5 uint32 = 2147483647
    v1 = v1
    var v6 string = "hello"
    v7 := []int{}
    if (-128 > -1) {
        var v8 float32 = float32((((v2 - ((v2 - (0.1 + (((((v2 + (v2 - (float64(3.14159) + (3.14159 + v2)))) + ((float64(0.0) + v2) / (1e10 - v2))) * ((-2.25 / v2) * ((((float64(0.1) - ((((v2 + (v2 / v2)) + v2) + 3.14159) - v2)) + ((((0.1 * (1e10 - (v2 * float64(1.5)))) / v2) + (3.14159 * ((1e10 + 3.14159) + v2))) * (0.0 * v2))) / ((float64(float32(1.5)) + (float64((((1e10 * ((float32(v2) * 3.14159) * float32(v2))) * float32(-2.25)) + float32(3.14159))) * v2)) - (3.14159 / 3.14159))) - float64((((-2.25 * (1.5 / 1.5)) / -2.25) / (1.5 + (0.0 * float32(-2.25)))))))) + 0.0) - (1.5 / 3.14159)))) / v2)) / float64(3.14159)) - 0.0))
        _ = len(v4)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 10; fi9++ {
            _ = len(v3)
            _ = fi9
        }
        fmt.Println(v6, v6, v5)
        switch v6 {
        case "\n\t":
            fmt.Println(v5, v2, v6)
        case "hello":
            delete(v4, 32767)
        }
    }
    _ = len(v7)
    if (2 >= 128) {
        for fi10 := 0; fi10 < 7; fi10++ {
            switch 1 {
            case 32768:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            switch (^(fi10)) {
            case -9223372036854775808:
                v5 = 2147483647
            default:
                fmt.Println(v1)
            }
            var v11 uint16 = 2
            _ = fi10
            _ = v11
        }
        var v12 int8 = 118
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint64 = 18446744073709551615
            if false {
                switch v6 {
                case "":
                    for fi15 := 0; fi15 < 17; fi15++ {
                        fmt.Println(v1)
                        _ = fi15
                    }
                }
                var v16 float64 = v2
                if true {
                    fmt.Println(v16, v16)
                    if len(v7) > 0 {
                        v7[0] = (-(fi13))
                    }
                    fmt.Println(v2, v6, v6)
                }
                _ = v16
            } else {
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 uint16 = 739
                    _ = fi17
                    _ = v18
                }
                if (0 != 42) {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        v6 = v6
                        _ = fi19
                    }
                    switch (v6 + ("hello" + v6)) {
                    case "foo bar":
                        _ = len(v7)
                    case "\n\t":
                        fmt.Println(v1, fi13)
                    default:
                        delete(v4, -283)
                    }
                    v20 := GenS1{F0: (-(fi13))}
                    _ = v20
                }
            }
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    v21 := GenS2{F0: 2147483647, F1: 0}
    var v22 uint16 = 42
    var v23 int32 = -1
    for fi24 := 0; fi24 < 15; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            if ((v23 & (v23 & -190)) > (-(v23))) {
                var v26 int8 = -1
                var v27 int8 = (v26 - int8(v1))
                v28 := v4[(fi24 / ((int(v1)) | 1))]
                _ = v26
                _ = v27
                _ = v28
            }
            v1 -= v1
            fmt.Println(v22)
            _ = fi25
        }
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 int32 = (v23 >> ((0) & 63))
            _ = fi29
            _ = v30
        }
        v31 := GenS3{F0: fi24, F1: "a", F2: 82}
        _ = fi24
        _ = v31
    }
    fmt.Println(v1, v5, v6)
    for fi32 := 0; fi32 < 1; fi32++ {
        var v33 float64 = v2
        v23 -= 0
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
