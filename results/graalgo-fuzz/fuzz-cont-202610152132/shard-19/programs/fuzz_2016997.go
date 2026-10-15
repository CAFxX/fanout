package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int8 }
type GenS2 struct { F0 uint32; F1 string; F2 int16 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint32 = 511
    var v2 uint8 = 1
    var v3 uint = 42
    v4 := make(map[int]string)
    v5 := make(map[uint32]int)
    fmt.Println(v2, v3)
    for fi6 := 0; fi6 < 1; fi6++ {
        delete(v5, (^(v1)))
        fmt.Println(fi6)
        v1 = v1
        _ = fi6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        v2 = (^(v2))
        fmt.Println(v2)
        _ = fi7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int8 = -29
        var v10 float64 = 1.5
        fmt.Println(v2)
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = len(v4)
    if ((((false || false) || false) && (uint8(v3) <= v2)) || false) {
        delete(v5, v1)
        v11 := []int{}
        v2 += uint8(v3)
        _ = v11
    } else {
        v12 := []int{-1, -9223372036854775808, 607}
        _ = v12
    }
    switch 128 {
    case -2147483648:
        v2 = (v2 / (((v2 - v2)) | 1))
    case 42:
        v1 -= v1
    case 4294967295:
        v13 := v5[(v1 * (^(v1)))]
        _ = v13
    }
    v14 := GenS1{F0: (57 >= 2), F1: 2147483647, F2: 28}
    v15 := GenS2{F0: (v1 / ((v1) | 1)), F1: ("" + ("hello" + (("\n\t" + (("foo bar" + "foo bar") + "hello")) + "hello"))), F2: 929}
    switch "" {
    case "hello":
        v1 = v1
    case "foo bar":
        if ((^(v2)) != (^(v2))) {
            if false {
                v2 += (v2 - (v2 / (((v2 >> ((v3) & 63))) | 1)))
                switch -2147483648 {
                case -538:
                    fmt.Println(v1)
                default:
                    v16 := GenS3{F0: 0}
                    _ = v16
                }
                delete(v4, -9223372036854775808)
            } else {
                delete(v5, v1)
                v1 -= (v1 - 823)
                fmt.Println(v3, v1)
            }
        } else {
            if false {
                v2 += (v2 | 127)
            }
            var v17 string = ("\n\t" + "\n\t")
            _ = v17
        }
    }
    if (((0.0 * -2.25) / (3.14159 - (((-2.25 / 3.14159) * (((3.14159 * (0.0 / (float32(3.14159) + 1.5))) + 0.1) + 1e10)) + (((float32(((((-2.25 - 1e10) * 1.5) / (float64((0.0 - 0.0)) + 3.14159)) + 1.5)) * 3.14159) + (float32(-2.25) * 3.14159)) * (1e10 - ((1e10 + 1e10) * ((-2.25 / 1.0) * 0.0))))))) != 1e10) {
        v18 := GenS3{F0: 913}
        v1 += 1
        v2 -= (v2 - v2)
        _ = v18
    }
    fmt.Println(v1)
    v19 := GenS2{F0: v1, F1: ((("foo bar" + (("" + (("hello" + "a") + "a")) + ("" + "foo bar"))) + ("foo bar" + "a")) + "hello"), F2: -389}
    var v20 uint = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
