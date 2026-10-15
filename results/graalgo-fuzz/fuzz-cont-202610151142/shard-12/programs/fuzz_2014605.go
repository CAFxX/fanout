package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int16; F1 int8; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 uint32 }
func genFn0(p0 int64, p1 uint64, p2 bool) int {
    v7 := []int{9223372036854775807, 734}
    var v8 int64 = -740
    fmt.Println(p0, p0, p1)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if true {
        if (!p2) {
            var v9 bool = (p0 >= p0)
            if v9 {
                v10 := GenS1{F0: (-1 <= 1), F1: 32767, F2: 255}
                switch -1 {
                case -34:
                    v11 := GenS2{F0: 32767, F1: -128, F2: 1}
                    _ = v11
                default:
                    var v12 int8 = 36
                    _ = v12
                }
                _ = v10
            } else {
                p0 -= (v8 / ((-135) | 1))
                var v13 uint16 = 255
                fmt.Println(v13, v8, p1)
                _ = v13
            }
            _ = v9
        } else {
            var v14 string = "a"
            for fi15 := 0; fi15 < 9; fi15++ {
                v16 := GenS1{F0: (18446744073709551615 >= 1), F1: 65535, F2: 0}
                _ = fi15
                _ = v16
            }
            var v17 uint16 = 993
            _ = v14
            _ = v17
        }
        v8 -= v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return int(p1)
}
func genFn1(p0 uint8, p1 int32) int {
    var v18 int = 32767
    v19 := GenS3{F0: 1}
    v20 := GenS2{F0: -32768, F1: 36, F2: 315}
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v20
    return (v18 / (((v18 % (((v18 & (v18 * v18))) | 1))) | 1))
}

func main() {
    v1 := []uint32{1, 2, 4294967295, 42}
    v2 := make(map[string]uint32)
    var v3 uint64 = 864
    var v4 string = ((("" + "") + "\n\t") + "hello")
    var v5 string = "hello"
    var v6 string = (v5 + v4)
    fmt.Println(v4)
    if (false && true) {
        var v21 int16 = 32767
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 bool = (-1 >= 618)
            v6 = "a"
            v3 = 2
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        _ = len(v1)
        var v24 int = -2147483648
        _ = v24
    }
    if (42 < 59) {
        v4 = (v4 + ((v5 + v5) + v6))
    } else {
        var v25 uint16 = 323
        for fi26 := 0; fi26 < 20; fi26++ {
            fmt.Println(v6, v25)
            _ = fi26
        }
        var v27 float32 = (1.5 + (((0.0 * 0.0) / (((3.14159 * -2.25) / 1.0) - (3.14159 * float32(1e10)))) / 1.5))
        _ = v25
        _ = v27
    }
    var v28 uint16 = 689
    var v29 int32 = 1
    for fi30 := 0; fi30 < 2; fi30++ {
        v29 += v29
        _ = fi30
    }
    v31 := GenS4{F0: 603, F1: 866}
    switch v6 {
    case "a":
        var v32 int64 = 9223372036854775807
        _ = v32
    }
    v3 = 1
    if (((((!false) && (32767 < 4294967295)) && false) && false) || false) {
        var v33 int = 397
        fmt.Println(v3, v28)
        _ = v33
    }
    var v34 float32 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
