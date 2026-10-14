package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0(p0 bool, p1 int8, p2 int) int32 {
    var v4 int = (^(p2))
    if p0 {
        switch ((("foo bar" + ("a" + "")) + "foo bar") + "foo bar") {
        case "\n\t":
            p0 = (!((-167 < 803) && p0))
        case "":
            switch "" {
            case "":
                v5 := GenS1{F0: (v4 + v4), F1: (p1 * (-(p1)))}
                _ = v5
            case "foo bar":
                p2 = p2
            case "\n\t":
                p1 -= p1
            }
        case "a":
            v6 := GenS1{F0: p2, F1: (p1 + (p1 & 88))}
            _ = v6
        default:
            v7 := GenS1{F0: p2, F1: p1}
            _ = v7
        }
        v8 := []int{}
        _ = v8
    }
    v9 := GenS2{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return -2147483648
}

func main() {
    var v1 int = 128
    v2 := []int{32768, -206}
    var v3 float32 = (((1e10 + 0.0) - (1.5 * 3.14159)) + ((1.5 * 1e10) + ((0.1 / 3.14159) / float32(float64(1.5)))))
    var v10 uint = 255
    v1 = (v1 % ((v1) | 1))
    var v11 float32 = (v3 - v3)
    fmt.Println(v1)
    if false {
        var v12 uint8 = uint8(v10)
        v3 = (3.14159 + float32(1e10))
        _ = v12
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 uint32 = 2147483647
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            switch ("foo bar" + "a") {
            case "a":
                var v16 int32 = genFn0(true, 82, (v1 - (fi15 | 9223372036854775807)))
                _ = v16
            case "hello":
                var v17 uint = v10
                _ = v17
            }
            if true {
                v18 := GenS1{F0: 926, F1: -73}
                _ = v18
            }
            _ = fi15
        }
    }
    v10 += v10
    v19 := GenS3{F0: 0, F1: 53}
    v3 = float32(0.1)
    v3 = float32(1.5)
    v3 += (v3 * ((v11 + (3.14159 / (1.5 + ((v3 - (1.5 * -2.25)) * v11)))) / -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
}
