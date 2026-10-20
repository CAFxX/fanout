package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint; F1 uint32; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 bool }

func main() {
    var v1 float32 = 0.1
    v2 := []int{-531, 2147483647}
    v3 := []uint32{4294967295, 1, 0}
    for fi4 := 0; fi4 < 9; fi4++ {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        v3 = append(v3, 1)
        _ = fi4
    }
    _ = len(v3)
    fmt.Println(v1)
    if (("foo bar" >= "") && ((!(-1 <= -1)) || ((0.0 / v1) <= 0.0))) {
        v5 := GenS1{F0: 2, F1: (((("foo bar" + (("hello" + ("foo bar" + "\n\t")) + "")) + ("\n\t" + "a")) + ("a" + "")) + "a")}
        v1 = -2.25
        _ = v5
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(v1, fi6)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi6
    }
    fmt.Println(v1)
    var v7 float64 = -2.25
    var v8 int = -633
    if (!(!(!(-1 != 0)))) {
        fmt.Println(v1, v1, v8)
        switch (v8 ^ v8) {
        case -670:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -1:
            if true {
                v9 := GenS2{F0: 1, F1: 111, F2: 941}
                if false {
                    var v10 uint8 = 1
                    _ = v10
                }
                _ = v9
            }
        case 496:
            switch -9223372036854775808 {
            case 128:
                if (v8 > (v8 / ((32768) | 1))) {
                    v8 += v8
                    var v11 int64 = 0
                    _ = v11
                }
            default:
                v3 = append(v3, 2147483648)
            }
        }
    }
    if (42 == 2147483647) {
        fmt.Println(v1)
    }
    v12 := GenS1{F0: 65535, F1: "foo bar"}
    v13 := GenS3{F0: -32768}
    v14 := GenS3{F0: -137}
    v15 := GenS4{F0: 506, F1: true}
    var v16 int = v8
    var v17 uint8 = 42
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
