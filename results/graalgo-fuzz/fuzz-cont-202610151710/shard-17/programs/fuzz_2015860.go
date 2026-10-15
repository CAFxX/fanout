package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint }
func genFn0(p0 string, p1 int64, p2 int8) uint64 {
    var v6 int16 = -676
    var v7 string = (((p0 + p0) + p0) + p0)
    var v8 int16 = v6
    var v9 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 uint8, p1 uint16) int32 {
    v10 := []int{}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 bool = ("foo bar" == "")
    var v4 uint8 = 0
    var v5 float64 = (-2.25 * -2.25)
    var v11 uint8 = 1
    fmt.Println(v5, v3, v11)
    switch "hello" {
    case "\n\t":
        v12 := []int{32768, 0}
        _ = v12
    case "hello":
        v1["a"] = 2147483647
    case "":
        if (v3 || v3) {
            if (1 == 18446744073709551615) {
                switch "\n\t" {
                case "hello":
                    var v13 uint16 = 893
                    _ = v13
                case "foo bar":
                    v11 += v11
                case "":
                    if true {
                        v4 = 207
                    }
                }
                switch 0 {
                case 1:
                    v4 += v11
                case -400:
                    v5 = v5
                }
            }
            v3 = (v3 || ((128 != (v11 ^ v4)) && (v3 || ((127 >= -47) || true))))
            v14 := []int{32767, -116}
            _ = v14
        } else {
            if (!(-1 > 100)) {
                var v15 int64 = 0
                fmt.Println(v4, v15, v15)
                var v16 float32 = (3.14159 - 3.14159)
                _ = v15
                _ = v16
            } else {
                v5 = 1.5
                v17 := GenS1{F0: "", F1: 255}
                _ = v17
            }
            switch "hello" {
            case "hello":
                var v18 float32 = ((3.14159 / (1e10 + 0.1)) * float32(-2.25))
                _ = v18
            }
            if v3 {
                var v19 uint32 = 2147483647
                _ = v19
            }
        }
    default:
        var v20 bool = ((2147483648 >= 133) || v3)
        _ = v20
    }
    v5 = float64(0.1)
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS1{F0: ("\n\t" + "hello"), F1: 333}
        v4 += (^(v4))
        switch "a" {
        case "hello":
            switch fi21 {
            case -825:
                var v23 int = (fi21 - 127)
                _ = v23
            default:
                v24 := []int{(fi21 % ((-9223372036854775808) | 1)), 290, (fi21 & fi21)}
                _ = v24
            }
        default:
            if (true || (!(true && (!(v3 || false))))) {
                switch "a" {
                case "foo bar":
                    fmt.Println(v4)
                case "\n\t":
                    if (1e10 <= -2.25) {
                        var v25 uint = uint(v4)
                        var v26 uint = 255
                        _ = v25
                        _ = v26
                    }
                case "hello":
                    v5 = v5
                default:
                    v27 := GenS1{F0: "hello", F1: 370}
                    _ = v27
                }
            } else {
                v28 := []int{(-(fi21)), fi21, 42, fi21}
                v3 = true
                _ = v28
            }
        }
        _ = fi21
        _ = v22
    }
    switch 1 {
    case 9223372036854775807:
        v11 = (^(v11))
    case 0:
        for fi29 := 0; fi29 < 7; fi29++ {
            v30 := GenS2{F0: 9223372036854775807, F1: (v11 << ((65535) & 63))}
            var v31 int64 = 9223372036854775807
            v32 := GenS3{F0: -1, F1: 0}
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
    case -1:
        v3 = false
    }
    fmt.Println(v11)
    for fi33 := 0; fi33 < 20; fi33++ {
        for fi34 := 0; fi34 < 13; fi34++ {
            v35 := GenS2{F0: 42, F1: v11}
            v4 -= 0
            v36 := v2[(fi34 ^ fi33)]
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi33
    }
    v37 := []int{-679, 0}
    v11 += v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
