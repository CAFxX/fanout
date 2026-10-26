package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int64; F2 int32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int64, p1 bool) int32 {
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int8 = -128
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 1, F1: 191}
    var v9 int = 42
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1(p0 int16, p1 bool, p2 uint8) uint16 {
    v10 := GenS2{F0: (^(p2)), F1: -9223372036854775808, F2: 42}
    var v11 int = 0
    v12 := []int{(v11 * v11), v11, (^(v11))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 255
}

func main() {
    var v1 int16 = 782
    v2 := make(map[int]string)
    var v3 uint16 = 171
    v4 := []uint32{2147483648, 707}
    var v5 int32 = int32(v1)
    v5 = 0
    var v13 bool = ("" == ((("\n\t" + (("foo bar" + ("\n\t" + (("" + (((((("\n\t" + "") + "hello") + "hello") + ("foo bar" + (("" + "a") + "foo bar"))) + "\n\t") + ("foo bar" + "\n\t"))) + ""))) + "a")) + "\n\t") + "a"))
    var v14 int16 = 290
    v3 = 42
    if ((!("foo bar" < "hello")) && v13) {
        if v13 {
            for fi15 := 0; fi15 < 7; fi15++ {
                fmt.Println(v3)
                fmt.Println(v14)
                v1 = v14
                _ = fi15
            }
            var v16 int = 32767
            var v17 int = -310
            _ = v16
            _ = v17
        } else {
            var v18 float64 = 1e10
            v18 = ((1e10 * (1.5 * (v18 - float64(1.5)))) + (v18 * v18))
            fmt.Println(v5, v13)
            _ = v18
        }
        v13 = ((v3 % ((genFn1(1, (!((("" + ("hello" + "\n\t")) + "a") == "a")), 1)) | 1)) == (^(v3)))
    }
    v3 = 862
    v13 = ((v3 / (((v3 >> ((42) & 63))) | 1)) >= genFn1((v14 & -558), v13, 2))
    for fi19 := 0; fi19 < 19; fi19++ {
        v3 -= (v3 * 65535)
        v14 = -488
        _ = fi19
    }
    if v13 {
        v14 -= 0
    }
    var v20 uint8 = 163
    var v21 float32 = (0.0 - 1e10)
    v22 := v2[128]
    var v23 uint64 = 0
    fmt.Println(v22, v3)
    if v13 {
        v24 := GenS3{F0: v23}
        _ = len(v2)
        _ = len(v4)
        _ = v24
    } else {
        for fi25 := 0; fi25 < 16; fi25++ {
            v22 = v22
            v2[(^(fi25))] = v22
            _ = fi25
        }
        v1 = (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
