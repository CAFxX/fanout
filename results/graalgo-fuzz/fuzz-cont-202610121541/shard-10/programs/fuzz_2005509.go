package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int16; F2 int8 }
type GenS5 struct { F0 int16; F1 uint; F2 uint32 }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint64) int16 {
    p0 = (p0 % ((p0) | 1))
    var v4 int = -2147483648
    if false {
        var v5 bool = true
        _ = v5
    }
    _ = p0
    _ = v4
    return 1000
}

func main() {
    var v1 int64 = 42
    var v2 uint64 = 2
    var v3 float64 = (float64(-2.25) - 3.14159)
    if true {
        for fi6 := 0; fi6 < 15; fi6++ {
            if (1e10 != (float32((-2.25 + 3.14159)) + float32(0.0))) {
                fmt.Println(v2, v1)
            }
            if (!true) {
                for fi7 := 0; fi7 < 9; fi7++ {
                    switch "a" {
                    case "foo bar":
                        v1 = 0
                    case "a":
                        var v8 uint32 = 2147483647
                        _ = v8
                    default:
                        v9 := GenS1{F0: 18446744073709551615, F1: (v1 << ((213) & 63)), F2: (fi7 / (((fi7 >> ((660) & 63))) | 1))}
                        _ = v9
                    }
                    var v10 uint16 = 32767
                    _ = fi7
                    _ = v10
                }
                v11 := []int{(fi6 - fi6), (fi6 ^ fi6)}
                v3 += (v3 + v3)
                _ = v11
            }
            _ = fi6
        }
        switch 64 {
        case 0:
            if ((128 < 214) && true) {
                var v12 float64 = 0.1
                for fi13 := 0; fi13 < 10; fi13++ {
                    v12 = -2.25
                    v2 -= (v2 & 42)
                    _ = fi13
                }
                _ = v12
            }
        case 32768:
            v14 := []int{-9223372036854775808, 42}
            _ = v14
        case 42:
            if (((((0.1 - 0.0) / (0.0 - -2.25)) * float32(float64(-2.25))) + (-2.25 + 0.0)) >= 1.5) {
                switch -1 {
                case -9223372036854775808:
                    v15 := GenS1{F0: 0, F1: (v1 & -1), F2: 32767}
                    _ = v15
                case 0:
                    if (1 < (v1 | (v1 ^ v1))) {
                        fmt.Println(v1, v2, v2)
                        v16 := []int{}
                        v17 := GenS1{F0: 1, F1: v1, F2: -2147483648}
                        _ = v16
                        _ = v17
                    } else {
                        v3 = 1e10
                    }
                case -597:
                    var v18 uint16 = 744
                    _ = v18
                default:
                    v19 := GenS1{F0: 0, F1: v1, F2: 42}
                    _ = v19
                }
            } else {
                v20 := GenS1{F0: 0, F1: (v1 % ((v1) | 1)), F2: 4294967295}
                v21 := []int{-989, -294, -410}
                _ = v20
                _ = v21
            }
        default:
            for fi22 := 0; fi22 < 12; fi22++ {
                v23 := []int{(fi22 << ((4294967295) & 63)), 9223372036854775807}
                _ = fi22
                _ = v23
            }
        }
        var v24 int = 4294967295
        _ = v24
    } else {
        if false {
            v25 := GenS2{F0: int32(v2), F1: int8(v2)}
            v2 = (v2 << ((121) & 63))
            _ = v25
        } else {
            var v26 uint16 = 42
            v27 := GenS3{F0: 197}
            _ = v26
            _ = v27
        }
    }
    if (("foo bar" + (("a" + "foo bar") + "foo bar")) < ("hello" + "\n\t")) {
        var v28 int16 = 659
        _ = v28
    }
    if true {
        var v29 string = ("\n\t" + ("a" + "hello"))
        _ = v29
    }
    var v30 float32 = (-2.25 / 0.1)
    switch "\n\t" {
    case "\n\t":
        v31 := GenS4{F0: -112, F1: 0, F2: 1}
        _ = v31
    case "a":
        v32 := []int{}
        _ = v32
    }
    v33 := GenS5{F0: -772, F1: 42, F2: 514}
    switch -446 {
    case -870:
        switch "" {
        case "foo bar":
            v3 += float64(((1e10 / 0.1) / 0.1))
        default:
            v34 := GenS1{F0: 561, F1: v1, F2: 9223372036854775807}
            _ = v34
        }
    case -1:
        v35 := []int{}
        _ = v35
    }
    switch ("\n\t" + "hello") {
    case "":
        v36 := []int{-1, 42}
        _ = v36
    case "a":
        v2 -= (^(v2))
    case "hello":
        v37 := GenS6{F0: (v1 >> ((762) & 63))}
        _ = v37
    default:
        v38 := []int{-9223372036854775808, -9223372036854775808, 32768}
        _ = v38
    }
    var v39 uint8 = uint8(v2)
    var v40 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v39
    _ = v40
}
