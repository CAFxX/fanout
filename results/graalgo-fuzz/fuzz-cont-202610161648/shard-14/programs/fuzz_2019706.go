package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int; F1 int; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint32 }
func genFn0(p0 bool, p1 int32, p2 uint) int {
    switch -152 {
    case 42:
        var v6 uint = p2
        _ = v6
    case 162:
        v7 := []int{32768}
        _ = v7
    }
    var v8 uint64 = 4294967295
    var v9 uint64 = v8
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(p1, p1)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 437
}

func main() {
    var v1 uint8 = 1
    var v2 float64 = 1e10
    var v3 int16 = -1
    var v4 string = ("foo bar" + ("hello" + ("" + "a")))
    v5 := make(map[int]int)
    v3 = v3
    v11 := []int{32767, -821, -1}
    for fi12 := 0; fi12 < 12; fi12++ {
        v1 += v1
        _ = fi12
    }
    switch 505 {
    case 32768:
        v5[-9223372036854775808] = 9223372036854775807
    case -519:
        var v13 int32 = -469
        _ = v13
    case 0:
        v2 = 0.0
    }
    v3 -= (^(v3))
    v14 := GenS1{F0: 0, F1: 9223372036854775807}
    switch v4 {
    case "":
        delete(v5, 9223372036854775807)
    case "a":
        if ((true && (v3 >= v3)) && (65535 < uint32(v1))) {
            _ = len(v11)
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 int32 = 2147483647
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 1; fi17++ {
                _ = len(v5)
                v5[127] = genFn0((-2.25 == (0.1 + (float32(v2) - (float32(v2) - float32((v2 / 1.0)))))), 0, 255)
                if true {
                    v3 = 32767
                    v18 := GenS1{F0: 1, F1: 1}
                    v19 := GenS1{F0: 249, F1: -610}
                    _ = v18
                    _ = v19
                } else {
                    v20 := GenS1{F0: 636, F1: -1}
                    _ = v20
                }
                _ = fi17
            }
        }
    case "foo bar":
        var v21 int64 = 1
        _ = v21
    default:
        v1 -= v1
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        if (!(0 > 2)) {
            fmt.Println(fi22)
            v4 = ((("foo bar" + "hello") + ("" + ("hello" + (("\n\t" + "hello") + v4)))) + "foo bar")
            switch fi22 {
            case 42:
                v4 = ("a" + v4)
            case 32768:
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            }
        } else {
            var v23 uint8 = (v1 << ((2) & 63))
            if true {
                var v24 float32 = 1e10
                v25 := GenS2{F0: (fi22 >> ((0) & 63)), F1: (-(fi22)), F2: 4294967295}
                v26 := GenS3{F0: 4294967295, F1: 0, F2: 1}
                _ = v24
                _ = v25
                _ = v26
            } else {
                v27 := v5[fi22]
                _ = v27
            }
            v28 := GenS2{F0: (fi22 / ((238) | 1)), F1: 456, F2: 367}
            _ = v23
            _ = v28
        }
        fmt.Println(v1, v3)
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
