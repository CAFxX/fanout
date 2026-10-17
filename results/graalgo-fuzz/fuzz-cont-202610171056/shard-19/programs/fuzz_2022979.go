package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 string; F2 int32 }
func genFn0() uint {
    if ((9223372036854775807 > -9223372036854775808) && false) {
        var v4 float32 = float32(1.5)
        _ = v4
    } else {
        switch 1 {
        case 1:
            v5 := GenS1{F0: -323, F1: -812, F2: 2}
            _ = v5
        default:
            var v6 string = ""
            _ = v6
        }
        v7 := []int{1, 911, -534}
        _ = v7
    }
    var v8 int64 = -435
    v8 += v8
    v8 = v8
    _ = v8
    return 1
}
func genFn1(p0 uint16) int8 {
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS2{F0: (!(!(2147483647 == 65535)))}
        var v11 string = "\n\t"
        p0 += p0
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            p0 = (p0 | (p0 % ((p0) | 1)))
            p0 -= p0
            if ((!(genFn0() == genFn0())) || false) {
                switch fi12 {
                case 0:
                    v14 := []int{(fi13 - fi12), fi12, (fi13 + fi13), fi12}
                    _ = v14
                case -4:
                    v15 := GenS3{F0: 451, F1: "\n\t", F2: 0}
                    _ = v15
                case 9223372036854775807:
                    switch "" {
                    case "":
                        v16 := []int{}
                        _ = v16
                    }
                default:
                    v17 := []int{1}
                    _ = v17
                }
            } else {
                var v18 int64 = 0
                v19 := []int{-2147483648}
                var v20 uint16 = (p0 % (((p0 / (((^(p0))) | 1))) | 1))
                _ = v18
                _ = v19
                _ = v20
            }
            _ = fi13
        }
        _ = fi12
    }
    p0 -= p0
    p0 -= 65535
    _ = p0
    return 40
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 float32 = float32((0.1 / 1e10))
    v21 := []int{-154, 382, 0, -718}
    if true {
        var v22 int8 = 42
        v23 := GenS3{F0: 4294967295, F1: (("" + ("a" + "\n\t")) + "hello"), F2: -54}
        _ = v22
        _ = v23
    } else {
        var v24 uint32 = 42
        _ = v24
    }
    v25 := GenS2{F0: (((true && true) && true) || true)}
    v26 := GenS2{F0: true}
    var v27 int16 = -32768
    v27 = (^(v27))
    if (v3 < v3) {
        var v28 string = "foo bar"
        var v29 string = v28
        v30 := GenS1{F0: (v27 % ((v27) | 1)), F1: 32767, F2: 77}
        _ = v28
        _ = v29
        _ = v30
    } else {
        v27 -= (-(v27))
    }
    var v31 float64 = ((float64((((((1e10 / v3) / 0.1) - -2.25) - v3) * (v3 - v3))) - ((float64(3.14159) - (0.0 + 3.14159)) - 0.0)) + (1e10 - 1e10))
    v1[1] = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
}
