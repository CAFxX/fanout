package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int16; F1 int32; F2 int8 }

func main() {
    var v1 uint = 255
    var v2 int = 2147483647
    var v3 string = (((("\n\t" + "foo bar") + "a") + "hello") + "")
    if ((true || (false || false)) || true) {
        for fi4 := 0; fi4 < 4; fi4++ {
            v1 = v1
            for fi5 := 0; fi5 < 9; fi5++ {
                var v6 float32 = (float32(1.5) / (float32(0.1) * ((1e10 / (0.1 * (1.5 * 1e10))) + float32((0.1 / 1.0)))))
                var v7 uint64 = 1
                var v8 uint8 = uint8(v1)
                _ = fi5
                _ = v6
                _ = v7
                _ = v8
            }
            _ = fi4
        }
        for fi9 := 0; fi9 < 9; fi9++ {
            v3 = v3
            v2 = 42
            _ = fi9
        }
    }
    v10 := []int{}
    _ = len(v10)
    switch (v2 / (((-(v2))) | 1)) {
    case -949:
        v11 := GenS1{F0: -1, F1: 42, F2: 127}
        _ = v11
    case 1:
        v2 = v2
    case 0:
        if ((42 >= 42) || true) {
            for fi12 := 0; fi12 < 7; fi12++ {
                v13 := GenS1{F0: 42, F1: 1, F2: -128}
                _ = fi12
                _ = v13
            }
        } else {
            fmt.Println(v3, v2, v2)
            v10 = append(v10, v2)
            v10 = append(v10, (v2 | -9223372036854775808))
        }
    default:
        v2 = v2
    }
    var v14 float32 = 0.0
    _ = len(v10)
    for fi15 := 0; fi15 < 18; fi15++ {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi15
    }
    v1 += (v1 % ((v1) | 1))
    v14 -= 3.14159
    var v16 uint64 = 18446744073709551615
    for fi17 := 0; fi17 < 11; fi17++ {
        switch (v2 + (fi17 >> ((v1) & 63))) {
        case -915:
            var v18 bool = (9223372036854775807 == -1)
            _ = v18
        case 0:
            v19 := GenS2{F0: -1, F1: true}
            _ = v19
        case -529:
            if true {
                v14 = 0.0
            } else {
                v20 := GenS3{F0: 177, F1: 255, F2: 2147483648}
                var v21 int = 32768
                v14 = v14
                _ = v20
                _ = v21
            }
        }
        switch v3 {
        case "\n\t":
            _ = len(v10)
        }
        _ = fi17
    }
    v3 = "foo bar"
    v22 := GenS4{F0: -1, F1: 0, F2: -89}
    var v23 int64 = 0
    var v24 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
