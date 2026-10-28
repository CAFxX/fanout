package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 int }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{42, 530}
    var v3 float64 = ((0.1 * (((3.14159 + -2.25) / (-2.25 / 3.14159)) / float64(0.1))) + -2.25)
    v4 := GenS1{F0: 791, F1: -2147483648}
    v3 = float64(1.5)
    var v5 uint8 = 255
    v3 = 3.14159
    v6 := GenS2{F0: 32767, F1: "foo bar"}
    var v7 uint8 = v5
    switch "" {
    case "\n\t":
        v1[127] = 42
    default:
        var v8 uint8 = v5
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            v7 = v5
            fmt.Println(fi10, fi10)
            _ = len(v1)
            var v11 uint = 4294967295
            fmt.Println(fi10)
            _ = fi10
            _ = v11
        }
        v1[fi9] = 2147483647
        for fi12 := 0; fi12 < 3; fi12++ {
            if true {
                v13 := GenS2{F0: -1, F1: ("foo bar" + (("hello" + "\n\t") + ("a" + "foo bar")))}
                v14 := GenS3{F0: 489, F1: (false && (!(1 < 4294967295))), F2: 84}
                _ = v13
                _ = v14
            } else {
                if true {
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(fi9, v5, v5)
                    case "":
                        if true {
                            v15 := GenS3{F0: 65535, F1: false, F2: 0}
                            _ = v15
                        }
                    }
                } else {
                    if true {
                        v5 += v5
                    } else {
                        _ = len(v1)
                        v3 += (1.5 + v3)
                        v16 := GenS1{F0: 42, F1: 2147483647}
                        _ = v16
                    }
                    v17 := GenS4{F0: 65535}
                    _ = v17
                }
                v18 := GenS5{F0: ("hello" + ("\n\t" + ("" + "hello"))), F1: (^(fi12))}
                v5 -= v7
                _ = v18
            }
            _ = fi12
        }
        _ = fi9
    }
    var v19 uint32 = 972
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
