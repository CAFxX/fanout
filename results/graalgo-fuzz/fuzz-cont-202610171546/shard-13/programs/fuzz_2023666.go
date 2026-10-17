package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (1e10 + float32(((1e10 - 0.0) * (1.5 + -2.25))))
    var v3 uint8 = 128
    var v4 float64 = 1e10
    if (47 != 0) {
        delete(v1, 217)
        switch "a" {
        case "hello":
            var v5 int32 = 0
            _ = v5
        case "":
            switch "" {
            case "foo bar":
                var v6 int32 = -2147483648
                _ = v6
            default:
                v7 := GenS1{F0: true, F1: 65535, F2: -353}
                _ = v7
            }
        case "\n\t":
            fmt.Println(v2, v3)
        }
        fmt.Println(v2, v3)
    }
    fmt.Println(v3)
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: ((!false) && false), F1: 936, F2: -129}
        v10 := GenS1{F0: ((("a" + "\n\t") < ((("foo bar" + "hello") + "hello") + "foo bar")) && false), F1: 255, F2: 2147483647}
        v11 := []int{fi8, fi8, fi8}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := GenS2{F0: 718, F1: 314}
    v13 := GenS2{F0: -864, F1: 32767}
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 uint16 = 912
        _ = fi14
        _ = v15
    }
    if true {
        v16 := GenS3{F0: 100, F1: 255}
        v4 -= 3.14159
        fmt.Println(v2, v3)
        _ = v16
    } else {
        switch "foo bar" {
        case "foo bar":
            v17 := GenS2{F0: 550, F1: 962}
            _ = v17
        case "a":
            for fi18 := 0; fi18 < 17; fi18++ {
                if false {
                    v4 = 1.5
                }
                switch "a" {
                case "\n\t":
                    v4 += v4
                case "":
                    _ = len(v1)
                case "foo bar":
                    v19 := []int{-9223372036854775808, (-(fi18)), fi18, (^(fi18))}
                    _ = v19
                }
                _ = fi18
            }
        default:
            delete(v1, 42)
        }
        v1[0] = ("foo bar" + "a")
    }
    v4 += (v4 + ((v4 / float64(1e10)) - 1e10))
    v20 := []int{9223372036854775807}
    var v21 float32 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
