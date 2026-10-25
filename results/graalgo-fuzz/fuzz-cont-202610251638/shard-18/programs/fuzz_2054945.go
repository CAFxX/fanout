package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 int; F1 bool; F2 int64 }
func genFn0(p0 uint64) string {
    v6 := GenS1{F0: true}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint16 = 255
        _ = fi7
        _ = v8
    }
    if (!true) {
        var v9 int = -25
        v10 := GenS2{F0: 4294967295, F1: 0, F2: -32}
        if ((-2.25 / (1e10 - (1e10 * 3.14159))) < ((-2.25 - -2.25) + 1.5)) {
            switch -504 {
            case -2147483648:
                v11 := GenS2{F0: 495, F1: 184, F2: 9223372036854775807}
                _ = v11
            case -636:
                for fi12 := 0; fi12 < 15; fi12++ {
                    v13 := GenS2{F0: (^(p0)), F1: 0, F2: 9223372036854775807}
                    _ = fi12
                    _ = v13
                }
            case 128:
                switch v9 {
                case 42:
                    if false {
                        fmt.Println(v9, p0)
                    }
                default:
                    v14 := GenS1{F0: (true && true)}
                    _ = v14
                }
            }
            v15 := GenS1{F0: true}
            v16 := []int{v9, 32768, (-(v9)), v9}
            _ = v15
            _ = v16
        }
        _ = v10
        _ = v9
    } else {
        p0 = p0
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int8 = -127
        for fi19 := 0; fi19 < 18; fi19++ {
            p0 -= p0
            if ((p0 | (p0 / ((1) | 1))) > p0) {
                fmt.Println(fi17)
                var v20 bool = false
                _ = v20
            }
            if true {
                var v21 int = fi17
                v22 := GenS3{F0: v18}
                _ = v21
                _ = v22
            }
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = v6
    return ("a" + ("a" + "\n\t"))
}
func genFn1(p0 int32, p1 string, p2 int) uint16 {
    v23 := []int{p2, p2, p2, (p2 * p2)}
    p2 += (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 65535
}

func main() {
    var v1 uint8 = 35
    v2 := []uint32{0, 0, 2147483648, 0}
    var v3 float64 = ((-2.25 * (0.1 - ((1.5 * (-2.25 / 1e10)) - 1e10))) * (1.5 * 0.1))
    v4 := []uint32{65535}
    v5 := make(map[int]bool)
    v24 := GenS1{F0: (!(976 != 42))}
    v25 := GenS1{F0: (!(255 == 18446744073709551615))}
    var v26 float64 = 1e10
    var v27 int32 = 0
    v27 += -430
    if true {
        var v28 float64 = (v26 + v26)
        var v29 uint64 = 886
        fmt.Println(v28, v26, v1)
        _ = v28
        _ = v29
    } else {
        fmt.Println(v26, v27)
        for fi30 := 0; fi30 < 16; fi30++ {
            v26 += (v26 + 0.1)
            var v31 uint32 = 805
            _ = fi30
            _ = v31
        }
    }
    var v32 float32 = -2.25
    for fi33 := 0; fi33 < 10; fi33++ {
        v34 := GenS1{F0: true}
        var v35 bool = true
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println(v27, v27)
    v36 := GenS4{F0: 1000, F1: -1}
    v37 := GenS5{F0: 237, F1: (2 == 42), F2: 490}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
