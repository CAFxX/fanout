package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int8; F1 int8; F2 int8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int64; F1 int64; F2 int }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[uint32]uint32)
    var v3 float64 = -2.25
    v4 := []string{"hello"}
    var v5 uint32 = 580
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: -3, F1: -747}
        _ = fi6
        _ = v7
    }
    switch ("\n\t" + ("a" + (("foo bar" + "") + "\n\t"))) {
    case "":
        v2[v5] = (v5 % (((^(v5))) | 1))
    case "foo bar":
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 uint = 166
                v11 := v2[(v5 - 42)]
                v3 = (3.14159 + float64(1e10))
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
    case "hello":
        v1 -= v1
    }
    var v12 int = 9223372036854775807
    var v13 float32 = float32((v3 / v3))
    var v14 uint64 = 4294967295
    switch v12 {
    case 4294967295:
        switch v12 {
        case -299:
            var v15 uint32 = 455
            _ = v15
        case -2147483648:
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 int64 = 42
                fmt.Println(v17, v5)
                switch (v12 | (v12 << ((65535) & 63))) {
                case -1:
                    v18 := GenS2{F0: (v1 ^ v1), F1: 2}
                    _ = v18
                case 869:
                    v19 := v2[42]
                    _ = v19
                case 1:
                    var v20 int = -1
                    _ = v20
                }
                _ = fi16
                _ = v17
            }
        }
    case -845:
        var v21 int = 127
        _ = v21
    default:
        fmt.Println(v3, v5, v3)
    }
    v13 = 0.1
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 int8 = -1
        var v24 int32 = -567
        if true {
            v25 := GenS2{F0: (v1 * (v1 ^ (v1 | v1))), F1: 42}
            var v26 int = fi22
            v14 = 4294967295
            _ = v25
            _ = v26
        } else {
            if false {
                var v27 int16 = -964
                switch v12 {
                case 4294967295:
                    var v28 float32 = 0.1
                    _ = v28
                case 1:
                    var v29 int32 = v24
                    _ = v29
                }
                for fi30 := 0; fi30 < 11; fi30++ {
                    if (true || true) {
                        _ = len(v2)
                        switch (fi30 + (fi22 % ((127) | 1))) {
                        case -2147483648:
                            v31 := GenS3{F0: v23, F1: (v23 << ((1) & 63)), F2: (v23 - (v23 << ((2) & 63)))}
                            _ = v31
                        case 540:
                            fmt.Println(v12, fi22)
                        case 128:
                            fmt.Println(v3, fi30, v12)
                        }
                    } else {
                        var v32 int32 = (v24 - (v24 * -1))
                        v24 += (v24 + 678)
                        v33 := GenS4{F0: v32, F1: "hello"}
                        _ = v32
                        _ = v33
                    }
                    _ = fi30
                }
                _ = v27
            }
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    switch (v12 - (v12 | 1)) {
    case 128:
        for fi34 := 0; fi34 < 3; fi34++ {
            v5 = (v5 >> ((65535) & 63))
            _ = fi34
        }
    case -347:
        v14 = v14
    case 3:
        v35 := GenS5{F0: -1, F1: 9223372036854775807, F2: 1}
        _ = v35
    default:
        if len(v4) > 0 {
            v4[0] = ("a" + "")
        }
    }
    var v36 int64 = 42
    v2[v5] = (v5 | (v5 + v5))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
