package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 string }
func genFn0(p0 uint, p1 uint) uint32 {
    v6 := []int{-9223372036854775808, -98, -2147483648}
    var v7 uint = p0
    for fi8 := 0; fi8 < 7; fi8++ {
        p0 = 198
        for fi9 := 0; fi9 < 3; fi9++ {
            fmt.Println(v7, fi9, fi8)
            fmt.Println(p0, fi8, p0)
            if true {
                p0 += 4294967295
            } else {
                var v10 uint = 122
                var v11 float32 = (float32(0.0) / -2.25)
                _ = v10
                _ = v11
            }
            _ = fi9
        }
        v6 = append(v6, (fi8 / (((fi8 + fi8)) | 1)))
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 431
}

func main() {
    var v1 string = ("\n\t" + "")
    v2 := []uint32{151}
    v3 := make(map[int]int)
    var v4 uint32 = 1
    var v5 float64 = (0.1 + 3.14159)
    v3[-9223372036854775808] = -2147483648
    v4 += v4
    _ = len(v3)
    fmt.Println(v4)
    var v12 string = v1
    v12 = "foo bar"
    if true {
        var v13 uint16 = 42
        var v14 int64 = -291
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        switch "" {
        case "hello":
            v4 -= 2147483648
        case "a":
            v1 = "foo bar"
        case "":
            v1 = "a"
        }
        switch fi15 {
        case -9223372036854775808:
            if (false || false) {
                var v16 uint = 4294967295
                _ = v16
            }
        case 4294967295:
            var v17 float64 = float64(0.0)
            _ = v17
        case 2147483647:
            v18 := GenS1{F0: 127, F1: v12, F2: ((v1 + v12) + v1)}
            _ = v18
        }
        _ = fi15
    }
    fmt.Println(v5, v4)
    v4 -= 90
    switch 9223372036854775807 {
    case 454:
        var v19 int16 = -1
        _ = v19
    default:
        var v20 int32 = 14
        _ = v20
    }
    switch ("foo bar" + v12) {
    case "\n\t":
        v12 = "\n\t"
    case "foo bar":
        if false {
            v2 = append(v2, v4)
            v21 := GenS2{F0: 255}
            v22 := GenS3{F0: 115, F1: (v1 + (((v12 + ((v1 + ("\n\t" + ("a" + ((((v12 + "\n\t") + v1) + "hello") + "foo bar")))) + v1)) + "a") + "foo bar"))}
            _ = v21
            _ = v22
        }
    default:
        v4 += (v4 % ((42) | 1))
    }
    v23 := GenS2{F0: 583}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
