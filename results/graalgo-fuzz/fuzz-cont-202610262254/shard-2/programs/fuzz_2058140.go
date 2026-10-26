package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int8 }
type GenS2 struct { F0 int; F1 int }
func genFn0(p0 uint64, p1 uint16) uint {
    var v5 string = "a"
    for fi6 := 0; fi6 < 2; fi6++ {
        v5 = ""
        v7 := []int{(fi6 + (fi6 >> ((538) & 63))), -1}
        _ = fi6
        _ = v7
    }
    var v8 int32 = int32(p0)
    if (false || (true || true)) {
        p1 = p1
        v5 = v5
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            if (0 < 42) {
                switch fi9 {
                case 9223372036854775807:
                    if (("hello" + "a") != v5) {
                        v11 := GenS1{F0: 4294967295, F1: 2, F2: 100}
                        p1 -= 255
                        v12 := GenS2{F0: fi9, F1: fi9}
                        _ = v11
                        _ = v12
                    }
                }
                v5 = v5
                p0 -= (p0 << ((255) & 63))
                switch "foo bar" {
                case "foo bar":
                    var v13 uint = 42
                    _ = v13
                case "":
                    var v14 float64 = (1e10 - 1e10)
                    _ = v14
                case "a":
                    var v15 uint = 18446744073709551615
                    _ = v15
                default:
                    p1 = p1
                }
                v16 := GenS1{F0: 1, F1: 1, F2: 0}
                v17 := GenS2{F0: (fi9 % ((fi9) | 1)), F1: fi9}
                _ = v16
                _ = v17
            }
            var v18 int = fi10
            _ = fi10
            _ = v18
        }
        v5 = v5
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 552
    v3 := make(map[int]int)
    var v4 uint = 4294967295
    var v19 uint8 = 11
    switch -313 {
    case 127:
        var v20 uint32 = 0
        _ = v20
    case -1:
        for fi21 := 0; fi21 < 10; fi21++ {
            for fi22 := 0; fi22 < 6; fi22++ {
                v23 := GenS1{F0: 478, F1: 0, F2: 23}
                var v24 int8 = 1
                switch (-(fi21)) {
                case -2147483648:
                    var v25 float64 = ((1.5 + (1e10 / (-2.25 - ((0.1 - ((((1.5 + ((0.0 / 0.1) * 0.0)) * ((0.0 + -2.25) - 1e10)) + 1e10) - (-2.25 * 3.14159))) * (1e10 + (0.0 * ((0.1 - 3.14159) - 1e10))))))) - 0.0)
                    _ = v25
                case 2147483647:
                    var v26 bool = ((v24 ^ (v24 << (((v4 << (((v4 - (^(v4)))) & 63))) & 63))) <= -1)
                    _ = v26
                case 42:
                    fmt.Println(v19)
                default:
                    v27 := GenS2{F0: (^(fi21)), F1: (fi21 - (fi21 >> ((v4) & 63)))}
                    _ = v27
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
    default:
        v2 = (v2 + v2)
    }
    var v28 uint32 = uint32(v2)
    if ((true && (!true)) && ((^(v4)) == v4)) {
        v2 += (v2 | v2)
        var v29 int = 128
        _ = v29
    }
    if (577 >= (^(v2))) {
        if (true || (v28 > (v28 % ((v28) | 1)))) {
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 bool = ((fi30 / (((fi30 / ((fi30) | 1))) | 1)) == (-(fi30)))
                _ = fi30
                _ = v31
            }
            v32 := GenS1{F0: (v4 | v4), F1: (v28 >> (((^(v4))) & 63)), F2: -92}
            v33 := GenS1{F0: v4, F1: 1, F2: 42}
            _ = v32
            _ = v33
        } else {
            switch "a" {
            case "foo bar":
                fmt.Println(v4)
            }
            v4 = (v4 + (v4 >> ((v4) & 63)))
            var v34 int64 = 42
            _ = v34
        }
        fmt.Println(v4, v19)
    } else {
        _ = len(v1)
    }
    delete(v3, 2147483647)
    switch "" {
    case "a":
        v2 -= v2
    case "":
        var v35 float32 = 1.5
        _ = v35
    case "\n\t":
        v28 = v28
    }
    v36 := GenS1{F0: v4, F1: uint32(v4), F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v4
}
