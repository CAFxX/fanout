package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int }
type GenS2 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint, p1 uint32) uint {
    p1 = 904
    v5 := []int{42}
    _ = p0
    _ = p1
    _ = v5
    return p0
}

func main() {
    var v1 uint16 = 255
    var v2 float64 = 0.0
    v3 := []uint32{4294967295, 65535, 1}
    var v4 int16 = 32767
    v1 = 65535
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int64 = -1
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v1, v1, fi9)
        var v10 int32 = -2147483648
        switch fi9 {
        case 42:
            v11 := GenS1{F0: false, F1: (("a" + "a") + "foo bar"), F2: fi9}
            _ = v11
        case 2147483647:
            v10 += (v10 | 42)
        case 32768:
            v10 = (-(v10))
        default:
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS1{F0: false, F1: "a", F2: (-(fi9))}
                _ = fi12
                _ = v13
            }
        }
        _ = fi9
        _ = v10
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v14 uint8 = 42
    fmt.Println(v4)
    if (false && (false && true)) {
        if false {
            if ((!false) || true) {
                switch "hello" {
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "\n\t":
                    _ = len(v3)
                case "":
                    var v15 int32 = -2147483648
                    _ = v15
                }
                v16 := GenS1{F0: ((true && true) || ((-1 < -1) || false)), F1: (("hello" + "a") + "foo bar"), F2: -2147483648}
                _ = v16
            }
        } else {
            switch 1 {
            case 32768:
                if false {
                    var v17 float64 = float64(0.1)
                    v4 += (v4 / ((v4) | 1))
                    _ = v17
                } else {
                    var v18 int16 = v4
                    var v19 float64 = (0.1 / (((float64(1.5) + 1e10) - 3.14159) * v2))
                    _ = v18
                    _ = v19
                }
            case -652:
                v20 := GenS1{F0: false, F1: "foo bar", F2: 4294967295}
                _ = v20
            case -9223372036854775808:
                v21 := GenS1{F0: (false || (v2 <= (v2 + float64(float32(1e10))))), F1: "foo bar", F2: -1}
                _ = v21
            }
            var v22 uint8 = 1
            v4 = 1000
            _ = v22
        }
        v14 = v14
        var v23 string = "a"
        _ = v23
    }
    switch (("hello" + "") + "\n\t") {
    case "a":
        v1 = v1
    default:
        v1 -= v1
    }
    v4 -= 0
    v24 := GenS1{F0: true, F1: "foo bar", F2: 128}
    v25 := GenS2{F0: 66, F1: 707}
    v2 -= 3.14159
    fmt.Println(v14, v14, v14)
    var v26 int64 = -453
    if ((1e10 == float32(3.14159)) || true) {
        var v27 float32 = 1e10
        _ = v27
    } else {
        var v28 bool = ((^(v26)) != (v26 / ((0) | 1)))
        _ = v28
    }
    v2 -= v2
    if (false || ((false && true) && false)) {
        var v29 uint64 = 1
        var v30 uint16 = (v1 / ((v1) | 1))
        _ = v29
        _ = v30
    }
    var v31 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v4
}
