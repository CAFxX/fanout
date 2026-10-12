package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int; F1 int; F2 int16 }

func main() {
    var v1 int = -200
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 int8 = 60
        _ = fi5
        _ = v6
    }
    switch "a" {
    case "\n\t":
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := []int{(-(v1)), (v1 % ((127) | 1))}
            _ = fi7
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int = (v1 / (((v1 | v1)) | 1))
            v11 := GenS1{F0: -1}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 float32 = 0.1
    var v13 int8 = -1
    var v14 int8 = 100
    switch "\n\t" {
    case "foo bar":
        var v15 uint16 = 927
        _ = v15
    default:
        v16 := GenS1{F0: 9223372036854775807}
        _ = v16
    }
    v17 := v3[636]
    v18 := GenS1{F0: 0}
    for fi19 := 0; fi19 < 20; fi19++ {
        switch ((("foo bar" + "hello") + v17) + ("foo bar" + (v17 + ("" + "foo bar")))) {
        case "":
            if (!true) {
                v1 -= (fi19 ^ (v1 >> ((18446744073709551615) & 63)))
                v2[(fi19 / ((132) | 1))] = (fi19 + (v1 % (((-(v1))) | 1)))
            } else {
                v17 = v17
                var v20 int32 = 0
                delete(v3, 869)
                _ = v20
            }
        case "\n\t":
            var v21 int64 = 42
            _ = v21
        default:
            switch v17 {
            case "foo bar":
                v22 := []int{2147483647}
                _ = v22
            case "a":
                v23 := GenS2{F0: 128, F1: 1000}
                _ = v23
            case "\n\t":
                var v24 bool = ((255 <= 0) && ((!(v14 > 42)) && (42 >= 1)))
                _ = v24
            }
        }
        _ = fi19
    }
    switch "\n\t" {
    case "hello":
        var v25 int = 32768
        _ = v25
    case "":
        v26 := GenS3{F0: -885, F1: (v1 * (v1 | (v1 * v1))), F2: 255}
        _ = v26
    case "\n\t":
        _ = len(v3)
    }
    v27 := GenS1{F0: 42}
    var v28 int16 = 568
    v12 -= ((1e10 - 0.0) / v12)
    v13 += -116
    v14 += (v14 + (-(v13)))
    for fi29 := 0; fi29 < 17; fi29++ {
        for fi30 := 0; fi30 < 13; fi30++ {
            switch fi29 {
            case -2147483648:
                var v31 int32 = int32(fi29)
                _ = v31
            case 352:
                v14 = v14
            }
            var v32 float32 = -2.25
            v33 := v2[(v1 ^ fi29)]
            _ = fi30
            _ = v32
            _ = v33
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
