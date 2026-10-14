package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := []uint32{42}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 uint64 = 0
    v4 += (v4 & v4)
    v4 = v4
    v4 = v4
    v4 -= (^(v4))
    switch "" {
    case "hello":
        var v5 bool = (!((!(200 == 255)) && (!(true && true))))
        _ = v5
    case "foo bar":
        var v6 int8 = 127
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int64 = 1
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(fi7)
            var v10 int32 = -290
            _ = fi9
            _ = v10
        }
        switch ("a" + "") {
        case "hello":
            v11 := GenS1{F0: 1, F1: 127}
            _ = v11
        case "\n\t":
            v4 -= v4
        default:
            var v12 uint64 = (v4 % (((v4 >> ((2) & 63))) | 1))
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    if (0.0 > (3.14159 / ((-2.25 / 0.1) + 1.5))) {
        switch 318 {
        case 1:
            v13 := GenS2{F0: -1}
            _ = v13
        }
        switch 128 {
        case -9223372036854775808:
            fmt.Println(v4)
        case -1:
            var v14 uint16 = 0
            _ = v14
        }
    } else {
        v15 := GenS3{F0: 2}
        v16 := GenS2{F0: -9223372036854775808}
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 int64 = -1
            fmt.Println(fi17)
            _ = fi17
            _ = v18
        }
        _ = v15
        _ = v16
    }
    var v19 float64 = 0.1
    var v20 float64 = v19
    var v21 uint = 65535
    v4 = v4
    v4 = v4
    var v22 float32 = 1e10
    fmt.Println(v22, v4, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
