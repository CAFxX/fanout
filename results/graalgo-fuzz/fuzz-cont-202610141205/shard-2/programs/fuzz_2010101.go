package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int }
func genFn0(p0 int64) int16 {
    v5 := GenS1{F0: (-(p0)), F1: 537}
    v6 := []int{}
    v7 := GenS2{F0: 42, F1: 4294967295}
    fmt.Println(p0)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -711
}

func main() {
    var v1 float32 = (1.5 - (1.5 / float32(((0.1 / (0.1 + 1.5)) + (1e10 + 0.1)))))
    v2 := []string{(((((((("foo bar" + (("hello" + "") + "hello")) + "a") + "hello") + "\n\t") + "\n\t") + "hello") + "") + "\n\t"), ("" + "hello")}
    var v3 uint32 = 42
    v4 := []uint32{(v3 ^ (^(v3)))}
    fmt.Println(v1, v3)
    switch 349 {
    case 127:
        if true {
            v4 = append(v4, (v3 / ((v3) | 1)))
            if len(v2) > 0 {
                v2[0] = "a"
            }
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 int16 = 32767
                _ = fi8
                _ = v9
            }
        }
    case -69:
        v10 := GenS2{F0: 1, F1: 884}
        _ = v10
    }
    var v11 uint64 = 18446744073709551615
    switch 128 {
    case 2147483647:
        v11 = 18446744073709551615
    case 4294967295:
        var v12 float32 = 1.5
        _ = v12
    }
    var v13 uint64 = uint64(v3)
    v14 := GenS3{F0: -876, F1: 8}
    switch "a" {
    case "a":
        if (160 > 1) {
            var v15 float64 = (3.14159 + ((((float64(float32(3.14159)) + 0.1) + float64(((1.5 - ((3.14159 + float32(0.0)) * 1e10)) * (float32(float64(0.0)) + v1)))) - 0.1) * 1.5))
            _ = v15
        } else {
            var v16 int16 = 1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v3 += 1
            _ = v16
        }
    case "":
        if (-1 == -88) {
            switch 4294967295 {
            case -9223372036854775808:
                v3 += (^(v3))
            case 128:
                if (0.1 < (1e10 * -2.25)) {
                    switch -37 {
                    case 2147483647:
                        v13 = v11
                    case -1:
                        var v17 int8 = 42
                        _ = v17
                    }
                    v18 := GenS4{F0: 1}
                    _ = v18
                } else {
                    v13 = uint64(v3)
                }
            }
        }
    case "foo bar":
        switch "hello" {
        case "\n\t":
            var v19 uint32 = v3
            _ = v19
        case "hello":
            v13 = (v13 % (((v13 & 541)) | 1))
        case "a":
            _ = len(v2)
        default:
            v20 := GenS5{F0: "a", F1: 42}
            _ = v20
        }
    }
    v21 := GenS2{F0: v11, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
