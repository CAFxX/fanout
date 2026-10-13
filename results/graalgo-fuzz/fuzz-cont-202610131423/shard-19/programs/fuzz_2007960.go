package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int8; F2 int16 }
type GenS5 struct { F0 uint32; F1 string; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint64; F2 int32 }
func genFn0(p0 uint32) int32 {
    v4 := GenS1{F0: 383, F1: -2147483648}
    if (251 > -1) {
        fmt.Println(p0)
    } else {
        v5 := GenS1{F0: 32767, F1: -1}
        v6 := []int{2147483647, 32768, 928}
        _ = v5
        _ = v6
    }
    v7 := []int{32768, 9223372036854775807, 4294967295, 32767}
    if true {
        switch "" {
        case "":
            if (0 > 4294967295) {
                var v8 float32 = ((1.5 / 1.5) / ((0.0 - 0.1) / (1e10 / 1e10)))
                var v9 uint16 = 65535
                _ = v8
                _ = v9
            }
        case "a":
            v7 = append(v7, -2147483648)
        case "\n\t":
            v10 := GenS1{F0: -32768, F1: 42}
            _ = v10
        default:
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 uint32 = (p0 >> ((255) & 63))
                var v13 int16 = int16(p0)
                v14 := GenS1{F0: 32767, F1: 0}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        }
        p0 += 1
        var v15 float32 = ((-2.25 * -2.25) - (float32(0.0) - 0.1))
        _ = v15
    }
    _ = p0
    _ = v4
    _ = v7
    return 1
}
func genFn1(p0 uint8, p1 uint64) uint8 {
    var v16 uint64 = 42
    fmt.Println(p0, p0)
    v17 := GenS1{F0: -966, F1: 42}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return p0
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 127
    var v3 float64 = 0.0
    v18 := GenS1{F0: 0, F1: 452}
    v19 := GenS2{F0: 65535, F1: 0, F2: -807}
    v3 += 3.14159
    var v20 uint64 = 966
    v21 := []int{}
    v1 = "a"
    var v22 uint32 = 2147483647
    if ((0.0 / (((0.0 + 3.14159) + (((1.5 * -2.25) - 0.0) / float32(((v3 * -2.25) + ((0.0 / (((v3 - 0.0) + v3) + 0.1)) + v3))))) - 3.14159)) < 3.14159) {
        v21 = append(v21, 42)
        v22 = 1
        switch v1 {
        case "a":
            fmt.Println(v22)
        case "hello":
            v23 := GenS2{F0: 1, F1: 255, F2: 1}
            _ = v23
        default:
            _ = len(v21)
        }
    } else {
        v22 = (v22 - (^(v22)))
    }
    switch ("foo bar" + (((((v1 + v1) + "") + ("a" + v1)) + v1) + ("\n\t" + ""))) {
    case "":
        if len(v21) > 0 {
            v21[0] = 42
        }
    case "\n\t":
        if (-1 <= 9223372036854775807) {
            if ((^(v2)) < genFn1(v2, 42)) {
                v24 := GenS3{F0: 778}
                v25 := GenS2{F0: 1, F1: 2, F2: int(v2)}
                _ = v24
                _ = v25
            } else {
                v1 = ("hello" + ("hello" + "\n\t"))
                switch 2147483647 {
                case -9223372036854775808:
                    var v26 int8 = 42
                    _ = v26
                }
            }
            if false {
                v20 = v20
            } else {
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
                v27 := GenS4{F0: (!false), F1: -3, F2: 1}
                var v28 uint8 = v2
                _ = v27
                _ = v28
            }
            v29 := GenS2{F0: 2, F1: 556, F2: -2147483648}
            _ = v29
        }
    }
    v2 = 128
    for fi30 := 0; fi30 < 7; fi30++ {
        v3 += (((float64(1.5) - 0.1) * (v3 * v3)) + 0.1)
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 uint16 = 32767
            switch "hello" {
            case "":
                v33 := GenS5{F0: (v22 / (((^(v22))) | 1)), F1: "", F2: 9223372036854775807}
                _ = v33
            case "hello":
                v34 := GenS3{F0: 347}
                _ = v34
            }
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    switch 983 {
    case -367:
        for fi35 := 0; fi35 < 5; fi35++ {
            if false {
                var v36 bool = true
                v37 := GenS6{F0: (v2 + (v2 | (^(v2)))), F1: 709, F2: 2147483647}
                _ = v36
                _ = v37
            }
            _ = fi35
        }
    case -2147483648:
        fmt.Println(v2, v20)
    case -1:
        v38 := GenS4{F0: (-2147483648 < int(v22)), F1: -128, F2: 32767}
        _ = v38
    default:
        v39 := GenS1{F0: 0, F1: -629}
        _ = v39
    }
    var v40 int8 = 127
    fmt.Println(v2, v22, v40)
    v3 -= v3
    var v41 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v40
    _ = v41
}
