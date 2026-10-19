package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint16; F1 int8; F2 int64 }
type GenS3 struct { F0 int }

func main() {
    var v1 int32 = -1
    var v2 float32 = float32(float64(1e10))
    var v3 int32 = v1
    var v4 int = -2147483648
    var v5 float32 = -2.25
    var v6 uint8 = 2
    var v7 uint16 = uint16(v6)
    if true {
        fmt.Println(v4, v2)
    }
    v8 := GenS1{F0: 127, F1: ("\n\t" + "")}
    v9 := GenS2{F0: (v7 % ((42) | 1)), F1: -90, F2: -445}
    fmt.Println(v3, v3)
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int32 = 2147483647
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 string = "hello"
            if false {
                if false {
                    v11 -= (v1 * v11)
                    v14 := GenS3{F0: (fi10 ^ 42)}
                    _ = v14
                } else {
                    v15 := GenS1{F0: int8(v11), F1: "\n\t"}
                    switch (fi10 & -2147483648) {
                    case -9223372036854775808:
                        v13 = "a"
                    default:
                        if (int64(v7) > int64(fi12)) {
                            fmt.Println(v6)
                            var v16 string = "a"
                            _ = v16
                        }
                    }
                    _ = v15
                }
                v1 = (v1 ^ (v3 / ((v11) | 1)))
            }
            switch "a" {
            case "a":
                v17 := []int{(fi12 << ((18446744073709551615) & 63)), (^(fi10)), fi12, fi12}
                _ = v17
            case "foo bar":
                if false {
                    var v18 float64 = -2.25
                    v19 := []int{(fi12 * (fi12 >> ((316) & 63))), v4}
                    if ((255 == (v6 | v6)) && true) {
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        v5 = (-2.25 + 0.1)
                        v20 := GenS3{F0: fi12}
                        _ = v20
                    }
                    _ = v18
                    _ = v19
                }
            }
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v21 := []int{-2147483648}
    v6 = (v6 ^ v6)
    switch 32767 {
    case 1:
        var v22 int8 = 1
        _ = v22
    case -9223372036854775808:
        for fi23 := 0; fi23 < 10; fi23++ {
            switch "a" {
            case "":
                switch "a" {
                case "":
                    var v24 int16 = 34
                    _ = v24
                case "hello":
                    if (((v4 >= (fi23 | v4)) && true) && false) {
                        v6 = (v6 & (^(v6)))
                    } else {
                        switch fi23 {
                        case 42:
                            fmt.Println(v7, v4, v5)
                        default:
                            v1 = (v1 % (((v3 / ((2147483647) | 1))) | 1))
                        }
                        var v25 uint8 = (v6 * 127)
                        _ = v25
                    }
                }
            }
            _ = fi23
        }
    default:
        v1 = (v3 ^ 42)
    }
    v2 = 3.14159
    v1 -= 2147483647
    fmt.Println(v1, v2, v4)
    v4 -= v4
    _ = len(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
