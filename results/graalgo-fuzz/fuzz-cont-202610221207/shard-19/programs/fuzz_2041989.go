package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int8; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint64 }
func genFn0() uint16 {
    var v6 float32 = 0.0
    var v7 string = ("\n\t" + "\n\t")
    var v8 bool = (((v6 / (float32(3.14159) + ((float32(3.14159) * 3.14159) / float32(1.5)))) < float32(0.1)) || ((531 < 0) && false))
    var v9 int = 4294967295
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int32 = 1
        if v8 {
            v11 -= 42
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := []int{(fi12 & -949)}
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 int32 = v11
                var v16 int8 = 42
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        _ = fi10
        _ = v11
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 640
}

func main() {
    var v1 uint = 858
    var v2 float32 = float32(3.14159)
    var v3 int16 = -961
    var v4 float64 = (float64(v2) + 1e10)
    var v5 uint64 = uint64(v1)
    var v17 int32 = -2147483648
    v17 -= (v17 >> ((0) & 63))
    v18 := []int{-1, 32768, -1, -603}
    if (v2 == (3.14159 + 1.5)) {
        var v19 float32 = v2
        fmt.Println(v1, v2, v1)
        if ((-5 > 1) && (!true)) {
            v4 += v4
        } else {
            fmt.Println(v1, v3, v3)
        }
        _ = v19
    } else {
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        var v20 float64 = 0.0
        _ = v20
    }
    fmt.Println(v4, v1, v17)
    if false {
        v4 = v4
        v18 = append(v18, 32768)
        for fi21 := 0; fi21 < 13; fi21++ {
            v3 -= v3
            fmt.Println(v2, v4)
            v22 := GenS1{F0: ("hello" + (("a" + "\n\t") + "foo bar")), F1: (fi21 ^ -818)}
            _ = fi21
            _ = v22
        }
    } else {
        v18 = append(v18, -2147483648)
        for fi23 := 0; fi23 < 5; fi23++ {
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 uint16 = 42
                v1 -= (^(v1))
                _ = fi24
                _ = v25
            }
            v26 := GenS1{F0: "hello", F1: (fi23 & fi23)}
            for fi27 := 0; fi27 < 8; fi27++ {
                switch "\n\t" {
                case "a":
                    var v28 int64 = -939
                    _ = v28
                }
                var v29 float32 = 3.14159
                var v30 uint = v1
                var v31 uint32 = 271
                v32 := GenS1{F0: "foo bar", F1: 1}
                _ = fi27
                _ = v29
                _ = v30
                _ = v31
                _ = v32
            }
            _ = fi23
            _ = v26
        }
    }
    v18 = append(v18, -977)
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    var v33 uint8 = 146
    if ((!((v1 & v1) > v1)) || (float32(((v4 + (((v4 * v4) + v4) * (((v4 / ((float64((v2 - v2)) - (float64((v2 * (v2 + 0.0))) + -2.25)) + 1.5)) + (float64(1.5) / v4)) - v4))) * (float64((float32((v4 * (((3.14159 / (v4 / v4)) - float64((v2 + (1e10 - (float32(v4) * ((v2 + v2) + float32(v4))))))) + ((v4 + v4) / float64(v2))))) * ((float32((float64((-2.25 / (float32(v4) - v2))) - 0.0)) * float32(v4)) / float32(3.14159)))) / (float64((v2 / v2)) * (v4 * v4))))) != (1e10 - (v2 + v2)))) {
        if true {
            v3 = (v3 ^ int16(v1))
            v34 := GenS2{F0: 1}
            _ = v34
        }
    }
    v35 := GenS3{F0: v3, F1: 127, F2: 0}
    v36 := GenS4{F0: 1, F1: (v5 * 4294967295), F2: (^(v5))}
    v33 = v33
    switch ("" + ("" + "hello")) {
    case "foo bar":
        for fi37 := 0; fi37 < 1; fi37++ {
            fmt.Println(v3, v3, v1)
            fmt.Println(v2)
            v3 = (v3 << (((v1 & (v1 + (v1 - 0)))) & 63))
            _ = fi37
        }
    default:
        v5 = 42
    }
    fmt.Println(v4, v33, v33)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
