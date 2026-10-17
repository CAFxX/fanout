package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int32; F2 uint8 }
type GenS5 struct { F0 int32; F1 uint; F2 int8 }
type GenS6 struct { F0 string; F1 uint16 }
func genFn0(p0 uint32, p1 uint64, p2 int32) string {
    v7 := []int{-1}
    p2 = p2
    v8 := GenS1{F0: 42, F1: 42, F2: -391}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    var v1 uint64 = 2
    var v2 string = ""
    var v3 uint64 = 956
    var v4 int8 = -61
    var v5 float64 = ((1e10 * ((0.1 * 0.0) + float64(0.1))) + -2.25)
    v6 := make(map[uint32]string)
    var v9 uint = 4294967295
    if (32767 >= -844) {
        v10 := GenS1{F0: 42, F1: 469, F2: 127}
        _ = v10
    } else {
        switch "" {
        case "\n\t":
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 int32 = 1
                v13 := GenS2{F0: v4}
                v14 := GenS2{F0: v4}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        case "foo bar":
            var v15 uint8 = 128
            _ = v15
        case "a":
            v1 -= 777
        }
    }
    if false {
        v2 = (("\n\t" + v2) + v2)
    }
    v3 = v1
    var v16 int16 = -37
    if true {
        switch -759 {
        case 1:
            v17 := []int{}
            _ = v17
        }
    }
    var v18 float64 = v5
    v9 += 2
    v19 := GenS3{F0: 65535}
    var v20 int16 = v16
    var v21 uint16 = 2
    if false {
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 float32 = (float32(-2.25) - 3.14159)
            v24 := v6[1]
            if ((false && false) || ((^(v21)) > v21)) {
                v25 := []int{(fi22 - fi22), (-(fi22)), (fi22 >> ((0) & 63)), fi22}
                v26 := GenS4{F0: v20, F1: -719, F2: 2}
                v27 := GenS5{F0: 42, F1: (v9 & (v9 % (((v9 ^ (v9 & v9))) | 1))), F2: (v4 & -1)}
                _ = v25
                _ = v26
                _ = v27
            }
            _ = fi22
            _ = v23
            _ = v24
        }
        v2 = (v2 + v2)
    }
    v6[113] = ((v2 + "foo bar") + ("" + ("a" + v2)))
    if (!(int64(v9) == -962)) {
        var v28 bool = (!(-1 < -40))
        _ = v28
    }
    v29 := GenS6{F0: (v2 + (v2 + ("hello" + ((v2 + ((v2 + "hello") + (("" + v2) + "foo bar"))) + v2)))), F1: (v21 & v21)}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
