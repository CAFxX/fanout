package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int8; F2 string }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 uint32 }
func genFn0() uint32 {
    v4 := GenS1{F0: -237}
    var v5 uint32 = 42
    var v6 uint = uint(v5)
    fmt.Println(v5, v5)
    if (false || false) {
        if true {
            v7 := []int{-9223372036854775808, -9223372036854775808, 127, 2147483647}
            v8 := GenS1{F0: -32768}
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 1000}
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := []int{fi10, 227, -9223372036854775808, fi10}
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 int32 = -1
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS2{F0: uint64(v5), F1: -121, F2: "\n\t"}
                v6 += v6
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi14
                _ = v15
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    _ = v4
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 int16) uint32 {
    v16 := []int{}
    switch 0 {
    case -261:
        p0 += p0
    case 42:
        v17 := GenS3{F0: -370, F1: 255}
        _ = v17
    }
    p0 += p0
    _ = p0
    _ = v16
    return 977
}

func main() {
    var v1 uint8 = 64
    var v2 float32 = 3.14159
    v3 := []int{42, -520, 4294967295, 32768}
    for fi18 := 0; fi18 < 19; fi18++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi18
    }
    if (false || (!(-1 == 4294967295))) {
        if (2 > 1) {
            var v19 int = 32767
            var v20 float64 = (((float64(float32(1e10)) + 0.1) * 0.0) - (-2.25 / 1.5))
            _ = v19
            _ = v20
        } else {
            fmt.Println(v1, v2)
        }
        for fi21 := 0; fi21 < 7; fi21++ {
            switch fi21 {
            case -2147483648:
                var v22 uint64 = 0
                _ = v22
            case 32767:
                var v23 int16 = 42
                _ = v23
            case -297:
                _ = len(v3)
            default:
                if (((3.14159 / 1.0) != float64(v2)) && ((true && false) && ((v1 << ((2) & 63)) <= v1))) {
                    var v24 int = fi21
                    switch -2147483648 {
                    case 128:
                        v1 += (v1 << ((18446744073709551615) & 63))
                    default:
                        v2 -= (float32(0.0) + -2.25)
                    }
                    _ = v24
                }
            }
            v2 += v2
            var v25 uint = 316
            _ = fi21
            _ = v25
        }
    }
    v26 := GenS4{F0: 18446744073709551615}
    if len(v3) > 0 {
        v3[0] = -2147483648
    }
    fmt.Println(v2, v2)
    v27 := GenS2{F0: 885, F1: 0, F2: "a"}
    var v28 uint = 18446744073709551615
    var v29 bool = (v28 != 297)
    v30 := GenS3{F0: -281, F1: 65535}
    if (627 <= 1) {
        var v31 uint64 = 0
        _ = v31
    } else {
        if v29 {
            v32 := GenS5{F0: uint64(v28), F1: 2}
            _ = v32
        } else {
            _ = len(v3)
        }
    }
    if (v1 >= v1) {
        if ((("a" + ("foo bar" + "hello")) > "\n\t") || false) {
            var v33 uint64 = 0
            _ = v33
        }
    }
    v2 = (v2 + (float32((3.14159 / float64(v2))) * ((v2 * ((v2 / v2) / v2)) - (float32((0.1 + (1.5 - 1.5))) + (v2 * 1e10)))))
    var v34 uint16 = 65535
    fmt.Println(v34, v29, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v35 uint8 = v1
    v36 := GenS4{F0: 652}
    var v37 uint16 = v34
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
