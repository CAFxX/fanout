package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint; F1 int64 }
func genFn0(p0 int) string {
    if (!false) {
        p0 = p0
        p0 -= -1
        p0 = p0
    } else {
        var v6 int32 = -411
        _ = v6
    }
    v7 := []int{(p0 - (p0 - -721)), 32768, p0}
    p0 = (p0 - p0)
    p0 += (p0 | 81)
    _ = p0
    _ = v7
    return "a"
}

func main() {
    var v1 uint = 255
    var v2 int = 9223372036854775807
    var v3 float32 = float32(0.1)
    v4 := make(map[uint32]uint32)
    v5 := []int{-995}
    v8 := GenS1{F0: true, F1: false, F2: false}
    v3 -= v3
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := GenS2{F0: 127, F1: 1}
        if (-9223372036854775808 <= v2) {
            v4[184] = 2
            var v11 uint64 = 18446744073709551615
            switch "" {
            case "a":
                for fi12 := 0; fi12 < 3; fi12++ {
                    v11 -= (v11 + (v11 + (v11 - 351)))
                    _ = fi12
                }
            case "":
                for fi13 := 0; fi13 < 19; fi13++ {
                    fmt.Println(fi13)
                    v11 = (v11 >> ((v1) & 63))
                    _ = fi13
                }
            }
            _ = v11
        }
        var v14 uint = (v1 << (((v1 << (((^(v1))) & 63))) & 63))
        _ = fi9
        _ = v10
        _ = v14
    }
    if ("a" == "\n\t") {
        v3 = v3
        var v15 int16 = -515
        _ = v15
    }
    switch 127 {
    case 32767:
        for fi16 := 0; fi16 < 1; fi16++ {
            v2 = 320
            _ = fi16
        }
    case 128:
        v17 := GenS3{F0: 800, F1: 389}
        _ = v17
    case 2147483647:
        v2 -= (v2 % (((v2 / (((-(v2))) | 1))) | 1))
    default:
        var v18 float64 = 1e10
        _ = v18
    }
    if (!true) {
        v19 := GenS3{F0: (^(v1)), F1: -1}
        v2 = 32768
        var v20 uint16 = 42
        _ = v19
        _ = v20
    } else {
        v1 = (v1 | (v1 % ((v1) | 1)))
        v3 = ((v3 + (v3 + v3)) + 1e10)
    }
    switch ("hello" + "") {
    case "foo bar":
        delete(v4, 672)
    case "hello":
        switch "" {
        case "":
            fmt.Println(v1)
        case "foo bar":
            v2 -= (v2 - -697)
        case "hello":
            var v21 int8 = 1
            _ = v21
        }
    }
    v3 -= 0.0
    if (0.1 <= ((v3 - ((v3 + (-2.25 / (v3 + (float32((-2.25 / 3.14159)) - v3)))) * v3)) - 3.14159)) {
        v22 := GenS3{F0: (v1 * 2), F1: 0}
        v1 = v1
        var v23 int = v2
        _ = v22
        _ = v23
    } else {
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 uint64 = 724
            var v26 float32 = v3
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = len(v4)
    }
    v1 = 330
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
