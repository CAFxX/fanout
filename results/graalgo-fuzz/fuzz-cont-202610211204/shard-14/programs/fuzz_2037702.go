package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 uint16 = 0
    var v2 bool = (!true)
    v3 := make(map[int]string)
    var v4 bool = v2
    v5 := []uint32{}
    switch 32768 {
    case 733:
        for fi6 := 0; fi6 < 14; fi6++ {
            v4 = (42 >= 794)
            var v7 int = (fi6 ^ (-(fi6)))
            switch v7 {
            case -302:
                delete(v3, (fi6 - (v7 + v7)))
            case 128:
                for fi8 := 0; fi8 < 18; fi8++ {
                    switch (-(fi8)) {
                    case 128:
                        fmt.Println(v2, v4, v4)
                    default:
                        var v9 float64 = (0.0 + -2.25)
                        var v10 uint16 = (^(v1))
                        fmt.Println(fi8, v9)
                        _ = v10
                        _ = v9
                    }
                    v2 = (!(true || (1 <= 540)))
                    v7 = (fi8 | (fi6 >> ((7) & 63)))
                    _ = fi8
                }
            }
            _ = fi6
            _ = v7
        }
    default:
        v11 := GenS1{F0: "foo bar", F1: 2147483648}
        _ = v11
    }
    switch -32 {
    case 1:
        if (float64((float32((float64(3.14159) - 0.0)) + 0.1)) > 1e10) {
            switch 0 {
            case 128:
                for fi12 := 0; fi12 < 16; fi12++ {
                    switch fi12 {
                    case 4294967295:
                        v2 = v2
                    case 32767:
                        v13 := GenS1{F0: "", F1: 65535}
                        _ = v13
                    default:
                        var v14 uint = 42
                        _ = v14
                    }
                    switch fi12 {
                    case 32768:
                        var v15 int64 = -9223372036854775808
                        _ = v15
                    case -131:
                        var v16 float64 = -2.25
                        _ = v16
                    }
                    _ = fi12
                }
            case 32767:
                var v17 int32 = 587
                _ = v17
            case -636:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                v18 := GenS1{F0: ("a" + "foo bar"), F1: 2147483648}
                _ = v18
            }
        }
    case -855:
        v2 = (-479 < 32768)
    case 615:
        v5 = append(v5, 0)
    }
    if (-2.25 != ((-2.25 - 1.5) - ((0.0 * ((((3.14159 / 1.5) / 1e10) * 0.1) / 0.1)) / ((1.5 * (-2.25 + (0.1 / float64(3.14159)))) - 0.0)))) {
        v2 = (!v2)
        for fi19 := 0; fi19 < 16; fi19++ {
            fmt.Println(v2, fi19, v4)
            switch fi19 {
            case 539:
                v20 := v3[(fi19 ^ (fi19 & (fi19 - (fi19 & (fi19 * (-(fi19)))))))]
                _ = v20
            }
            v21 := GenS1{F0: "a", F1: 1}
            _ = fi19
            _ = v21
        }
    } else {
        switch 0 {
        case -1:
            for fi22 := 0; fi22 < 20; fi22++ {
                v2 = (!v2)
                _ = fi22
            }
        case 1:
            v23 := GenS2{F0: (211 > -719), F1: 1, F2: -9223372036854775808}
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        v3[(fi24 | fi24)] = "foo bar"
        v4 = v2
        _ = fi24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := GenS3{F0: 42}
        _ = fi25
        _ = v26
    }
    v27 := GenS4{F0: 42, F1: 65535}
    if ((float32(3.14159) / 0.1) != (float32(0.0) + float32((1e10 * 0.0)))) {
        switch -548 {
        case 32767:
            var v28 bool = false
            _ = v28
        default:
            var v29 int16 = -448
            _ = v29
        }
        var v30 uint8 = 0
        _ = v30
    }
    v5 = append(v5, 2147483648)
    v2 = false
    fmt.Println(v2, v1, v4)
    v31 := GenS1{F0: "", F1: 2147483647}
    fmt.Println(v4, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
