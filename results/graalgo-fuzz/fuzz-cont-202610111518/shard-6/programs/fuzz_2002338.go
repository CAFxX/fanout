package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 bool; F2 int }
type GenS3 struct { F0 uint8; F1 bool }

func main() {
    var v1 int64 = 1
    v2 := []uint32{2, 42}
    var v3 int32 = -232
    var v4 int = -2147483648
    var v5 uint64 = 18446744073709551615
    fmt.Println(v3)
    var v6 float32 = 0.0
    v3 = (-(v3))
    var v7 float32 = float32(3.14159)
    if (((v4 % (((^(v4))) | 1)) != 128) && false) {
        switch "a" {
        case "\n\t":
            v4 = v4
        case "foo bar":
            var v8 float32 = v7
            _ = v8
        default:
            for fi9 := 0; fi9 < 13; fi9++ {
                if true {
                    v5 = (v5 + (v5 ^ (^(v5))))
                } else {
                    var v10 int8 = 0
                    var v11 uint8 = 255
                    var v12 string = "hello"
                    _ = v10
                    _ = v11
                    _ = v12
                }
                _ = fi9
            }
        }
        _ = len(v2)
    } else {
        v13 := GenS1{F0: ((4294967295 < (v4 + -1)) && false)}
        v14 := GenS1{F0: (true && (-19 != -77))}
        var v15 int8 = -117
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 string = ("" + "a")
    v17 := GenS1{F0: (!false)}
    fmt.Println(v3, v16)
    if false {
        v16 = v16
        fmt.Println(v3, v4)
    }
    if false {
        v7 = ((3.14159 + ((0.0 * 3.14159) - (1e10 - (float32(0.0) * v6)))) / v7)
    }
    var v18 string = (v16 + v16)
    if (!true) {
        var v19 int32 = 2147483647
        v7 += (-2.25 + v7)
        switch ((v18 + "foo bar") + (((v16 + "hello") + (v18 + ((v16 + v16) + "hello"))) + v18)) {
        case "\n\t":
            v4 += 32768
        case "":
            v6 = (v6 * v7)
        case "foo bar":
            v2 = append(v2, 394)
        }
        _ = v19
    } else {
        fmt.Println(v7)
        v20 := GenS2{F0: v3, F1: true, F2: (v4 * v4)}
        _ = v20
    }
    var v21 uint8 = 70
    fmt.Println(v1, v4)
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 15; fi23++ {
            switch ("foo bar" + v16) {
            case "hello":
                v24 := GenS3{F0: v21, F1: true}
                _ = v24
            }
            _ = fi23
        }
        v4 -= v4
        _ = fi22
    }
    if len(v2) > 0 {
        v2[0] = 149
    }
    if ((v4 == (^(v4))) && (false && ((^(v21)) > (v21 >> ((1) & 63))))) {
        for fi25 := 0; fi25 < 2; fi25++ {
            switch v18 {
            case "":
                v4 += 42
            case "a":
                v5 += uint64(v21)
            }
            v26 := GenS1{F0: ((v7 * v6) < float32((float64(float32(0.0)) - -2.25)))}
            v27 := GenS3{F0: (v21 & 193), F1: (4294967295 != 1)}
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
