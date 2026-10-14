package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int; F2 uint }

func main() {
    var v1 int16 = 42
    v2 := make(map[int]string)
    var v3 uint = 4294967295
    var v4 uint = (v3 + (v3 & (^(v3))))
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 float64 = 0.1
            _ = fi6
            _ = v7
        }
        switch ("foo bar" + ("\n\t" + "hello")) {
        case "\n\t":
            v4 += (v4 >> ((v4) & 63))
        default:
            switch int(v1) {
            case 766:
                var v8 float32 = 3.14159
                _ = v8
            case 127:
                for fi9 := 0; fi9 < 11; fi9++ {
                    v4 = 125
                    _ = fi9
                }
            }
        }
        var v10 float64 = float64(1e10)
        _ = fi5
        _ = v10
    }
    v1 = 42
    fmt.Println(v1, v1, v4)
    var v11 int16 = v1
    for fi12 := 0; fi12 < 19; fi12++ {
        v11 += (v1 / ((v1) | 1))
        _ = fi12
    }
    fmt.Println(v3, v3, v1)
    if false {
        fmt.Println(v11, v1)
        v4 = 42
    } else {
        switch -9223372036854775808 {
        case -2147483648:
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS1{F0: false, F1: false}
                v15 := GenS2{F0: 650}
                for fi16 := 0; fi16 < 2; fi16++ {
                    fmt.Println(v4, v11)
                    _ = fi16
                }
                _ = fi13
                _ = v14
                _ = v15
            }
        case 42:
            v11 -= v1
        default:
            v2[0] = (((("a" + "hello") + "\n\t") + "hello") + "")
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            switch fi17 {
            case 42:
                v18 := GenS3{F0: 417, F1: (fi17 % ((fi17) | 1)), F2: (v4 * 42)}
                _ = v18
            case 4294967295:
                for fi19 := 0; fi19 < 6; fi19++ {
                    v3 = (v3 & v3)
                    v20 := []int{fi19, fi17, 32768, (fi17 << (((v3 * 18446744073709551615)) & 63))}
                    _ = fi19
                    _ = v20
                }
            case 9223372036854775807:
                var v21 float32 = (1e10 - -2.25)
                _ = v21
            }
            v22 := v2[fi17]
            v1 = (v11 ^ v1)
            _ = fi17
            _ = v22
        }
        if false {
            v11 += (-(v11))
        }
    }
    if true {
        v1 = v11
        if false {
            var v23 float64 = (-2.25 / 3.14159)
            _ = v23
        } else {
            delete(v2, 643)
        }
    }
    v2[0] = "hello"
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 uint = 18446744073709551615
        _ = fi24
        _ = v25
    }
    var v26 bool = false
    var v27 bool = (65535 == 0)
    v28 := GenS1{F0: true, F1: v26}
    v27 = false
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
