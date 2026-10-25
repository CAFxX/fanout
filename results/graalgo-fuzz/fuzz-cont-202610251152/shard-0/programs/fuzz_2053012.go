package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 uint16, p1 int16) int {
    if (1e10 <= (((1e10 * 0.0) * 0.1) * 1.5)) {
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int16 = 1000
            switch "" {
            case "\n\t":
                if false {
                    var v9 float64 = 0.1
                    _ = v9
                }
            default:
                p0 = 32767
            }
            p1 = 32767
            _ = fi7
            _ = v8
        }
    }
    if ((true && (("\n\t" + "") >= "\n\t")) && (2147483648 <= 2147483648)) {
        fmt.Println(p1, p0)
    }
    switch ("" + "foo bar") {
    case "\n\t":
        fmt.Println(p0, p0)
    case "hello":
        v10 := GenS1{F0: 2, F1: (p1 * (p1 * 935)), F2: 149}
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 14; fi12++ {
                fmt.Println(p0, fi11, fi11)
                var v13 float32 = float32(((float64(0.0) - -2.25) - float64((1e10 * 0.1))))
                fmt.Println(v13, fi12)
                var v14 float64 = ((-2.25 / 1e10) * 1e10)
                _ = fi12
                _ = v13
                _ = v14
            }
            switch ("a" + "hello") {
            case "a":
                for fi15 := 0; fi15 < 2; fi15++ {
                    var v16 uint = 0
                    v17 := GenS1{F0: 884, F1: -811, F2: 128}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            case "\n\t":
                var v18 uint16 = 42
                _ = v18
            case "hello":
                var v19 int16 = p1
                _ = v19
            }
            _ = fi11
        }
    default:
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    return 437
}

func main() {
    var v1 uint64 = 2
    var v2 int = 42
    var v3 float32 = (3.14159 / 3.14159)
    var v4 uint16 = 952
    var v5 string = ("\n\t" + "hello")
    v6 := make(map[uint32]string)
    v20 := []int{831, (v2 ^ (v2 - v2)), v2}
    v3 -= v3
    v3 = float32((((float64(3.14159) + 3.14159) - 0.0) * float64(v3)))
    var v21 int16 = -304
    switch genFn0(635, v21) {
    case 32767:
        if (!(v1 == 4294967295)) {
            switch 1 {
            case 127:
                v1 = (v1 % (((v1 * (v1 & v1))) | 1))
            default:
                v22 := v6[684]
                _ = v22
            }
        }
    case -339:
        if len(v20) > 0 {
            v20[0] = genFn0((v4 % ((v4) | 1)), -32768)
        }
    case -253:
        v23 := GenS2{F0: 127, F1: 1}
        _ = v23
    }
    var v24 uint32 = 2147483647
    v2 = (^(v2))
    v20 = append(v20, v2)
    v1 = v1
    v25 := GenS1{F0: v1, F1: -32768, F2: uint8(v4)}
    var v26 int32 = 0
    switch v5 {
    case "\n\t":
        v6[0] = v5
    case "hello":
        var v27 int = 528
        _ = v27
    case "":
        var v28 int64 = -10
        _ = v28
    default:
        switch v2 {
        case 127:
            if true {
                for fi29 := 0; fi29 < 3; fi29++ {
                    var v30 float32 = (v3 - (((v3 * 1e10) - v3) + v3))
                    _ = fi29
                    _ = v30
                }
            } else {
                fmt.Println(v2)
            }
        case 0:
            v3 -= (float32(0.1) + v3)
        default:
            for fi31 := 0; fi31 < 13; fi31++ {
                var v32 int16 = v21
                _ = fi31
                _ = v32
            }
        }
    }
    if true {
        if true {
            v26 += (v26 | (v26 * (v26 % (((v26 >> ((772) & 63))) | 1))))
        }
    }
    fmt.Println(v5)
    v33 := GenS2{F0: 143, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
