package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint64; F1 int; F2 int32 }
func genFn0() uint8 {
    v7 := GenS1{F0: 127, F1: 177}
    var v8 float32 = 0.1
    switch 4294967295 {
    case -2147483648:
        for fi9 := 0; fi9 < 20; fi9++ {
            if true {
                var v10 uint32 = 4294967295
                v10 += v10
                _ = v10
            }
            if (!((65535 <= 809) || (!true))) {
                v8 += 1e10
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 int = fi11
                    _ = fi11
                    _ = v12
                }
                v13 := []int{}
                _ = v13
            } else {
                var v14 string = ""
                for fi15 := 0; fi15 < 2; fi15++ {
                    var v16 uint = 684
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            _ = fi9
        }
    case 32767:
        v17 := GenS2{F0: 1, F1: 65535}
        _ = v17
    case 904:
        var v18 int = 42
        _ = v18
    }
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint32, p1 uint, p2 uint32) string {
    var v19 uint16 = 0
    var v20 bool = false
    var v21 uint = p1
    fmt.Println(p1, v21)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v21
    return ""
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -2147483648
    var v3 uint = 110
    v4 := make(map[string]bool)
    var v5 int16 = -1
    v6 := []int{989}
    v22 := v4[((((("foo bar" + "") + "a") + "\n\t") + ("a" + ("foo bar" + (("a" + "hello") + "hello")))) + ((((("hello" + "foo bar") + "a") + "\n\t") + "") + "hello"))]
    _ = len(v6)
    v23 := GenS1{F0: -105, F1: genFn0()}
    v24 := GenS3{F0: -9223372036854775808, F1: ((v3 << (((v3 / ((42) | 1))) & 63)) >= 42)}
    fmt.Println(v5)
    fmt.Println(v2, v22)
    var v25 int32 = (v2 << (((v3 >> ((v3) & 63))) & 63))
    var v26 int = 2147483647
    v6 = append(v6, 1)
    var v27 int32 = -778
    var v28 float64 = (0.0 - 0.1)
    v29 := GenS4{F0: 128, F1: uint64(v3)}
    if (("" != "foo bar") || (!((v25 / ((v25) | 1)) <= -414))) {
        switch v26 {
        case 42:
            v2 -= 563
        case 32768:
            var v30 uint32 = 2147483648
            _ = v30
        case -9223372036854775808:
            if (42 >= 255) {
                for fi31 := 0; fi31 < 12; fi31++ {
                    v32 := GenS5{F0: 42, F1: (v26 >> (((^(v3))) & 63)), F2: v2}
                    _ = fi31
                    _ = v32
                }
            }
        }
        v6 = append(v6, v26)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
