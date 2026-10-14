package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS4 struct { F0 int16; F1 int32; F2 uint }

func main() {
    var v1 uint64 = 4294967295
    var v2 int16 = 1
    var v3 float32 = 3.14159
    var v4 int8 = -16
    v3 += (float32((0.1 / ((0.1 + float64((0.0 * (((1.5 / ((float32(0.0) - 3.14159) - (1.5 + 3.14159))) - v3) - (1.5 - (0.0 / float32((((float64(float32(1.5)) - 0.1) - 1e10) / (3.14159 + 0.1))))))))) * (-2.25 + (1.5 + 3.14159))))) * v3)
    v4 -= v4
    if false {
        v5 := []int{65, 32768, 370, 80}
        v6 := GenS1{F0: int16(v4)}
        _ = v5
        _ = v6
    }
    if (561 > 65535) {
        if (-1 >= -9223372036854775808) {
            v7 := GenS2{F0: 65535, F1: 255, F2: 127}
            _ = v7
        } else {
            v1 = v1
            v8 := []int{int(v4), 127, 868}
            _ = v8
        }
        v2 -= v2
        var v9 uint32 = 0
        _ = v9
    } else {
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS3{F0: 65535, F1: 4294967295, F2: (v2 + (v2 * (v2 ^ (v2 % ((76) | 1)))))}
            for fi12 := 0; fi12 < 12; fi12++ {
                if (!false) {
                    switch "a" {
                    case "foo bar":
                        var v13 int = fi12
                        _ = v13
                    case "hello":
                        v14 := []int{}
                        _ = v14
                    case "a":
                        fmt.Println(fi10, fi12)
                    default:
                        v2 -= 42
                    }
                    var v15 int8 = v4
                    _ = v15
                } else {
                    if (4294967295 < 255) {
                        fmt.Println(v4, fi12)
                        v16 := GenS2{F0: 18446744073709551615, F1: 82, F2: 255}
                        fmt.Println(fi12)
                        _ = v16
                    }
                    v2 = int16(fi10)
                    var v17 int16 = v2
                    _ = v17
                }
                var v18 int64 = 248
                v19 := GenS2{F0: 838, F1: 23, F2: 255}
                _ = fi12
                _ = v18
                _ = v19
            }
            switch "\n\t" {
            case "foo bar":
                var v20 uint32 = 65535
                _ = v20
            case "":
                v21 := []int{fi10}
                _ = v21
            default:
                if (!(0 > 38)) {
                    var v22 int16 = -1
                    _ = v22
                } else {
                    var v23 int16 = v2
                    _ = v23
                }
            }
            _ = fi10
            _ = v11
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            switch (("hello" + (("\n\t" + "") + "\n\t")) + "") {
            case "foo bar":
                v1 -= v1
            case "a":
                for fi25 := 0; fi25 < 16; fi25++ {
                    v26 := []int{}
                    if ((true && false) || false) {
                        fmt.Println(v1)
                        fmt.Println(fi25, v4, v4)
                        fmt.Println(fi24, v1)
                    }
                    v27 := GenS4{F0: v2, F1: 1, F2: 2}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
            var v28 int16 = (-(v2))
            _ = fi24
            _ = v28
        }
    }
    switch (("a" + ("" + ("foo bar" + ""))) + ("\n\t" + "\n\t")) {
    case "":
        var v29 uint16 = 42
        _ = v29
    case "foo bar":
        fmt.Println(v3, v2, v4)
    case "hello":
        v30 := GenS2{F0: 4294967295, F1: 280, F2: 47}
        _ = v30
    }
    v31 := []int{-904, -807}
    var v32 int32 = 0
    var v33 float32 = float32(-2.25)
    _ = len(v31)
    var v34 bool = (true || ((v32 - -1) != v32))
    var v35 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
