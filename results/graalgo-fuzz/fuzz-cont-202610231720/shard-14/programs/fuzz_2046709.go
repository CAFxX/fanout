package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int16; F2 int16 }
func genFn0() uint {
    var v5 int = 32767
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint32 = 840
        switch ("foo bar" + ("" + "a")) {
        case "":
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 uint8 = 83
                var v10 uint = 321
                _ = fi8
                _ = v10
                _ = v9
            }
        default:
            fmt.Println(fi6, fi6, v5)
        }
        for fi11 := 0; fi11 < 15; fi11++ {
            fmt.Println(fi11, fi6)
            v7 = 965
            v5 -= (fi6 >> ((1) & 63))
            v12 := []int{-2147483648, (fi11 * -373), -2147483648, -818}
            _ = fi11
            _ = v12
        }
        _ = fi6
        _ = v7
    }
    v13 := GenS1{F0: "foo bar", F1: 2147483647}
    _ = v13
    _ = v5
    return 18446744073709551615
}
func genFn1() uint {
    v14 := GenS2{F0: 1}
    if true {
        var v15 string = "a"
        var v16 uint64 = 471
        v17 := GenS3{F0: 100, F1: 32767, F2: 801}
        _ = v15
        _ = v16
        _ = v17
    }
    _ = v14
    return 255
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    var v3 uint16 = 804
    v4 := make(map[int]string)
    v18 := []int{32768, -453, 746, -440}
    v3 += (v3 + (v3 & 65535))
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    var v19 float32 = 0.1
    v19 = 1.5
    v19 += float32(-2.25)
    fmt.Println(v3, v19)
    var v20 int = 90
    var v21 uint8 = 42
    v3 += (^(v3))
    switch (-(v20)) {
    case 4294967295:
        switch "a" {
        case "\n\t":
            for fi22 := 0; fi22 < 5; fi22++ {
                delete(v1, 4294967295)
                _ = fi22
            }
        case "a":
            v21 = v21
        }
    case 267:
        var v23 uint16 = 901
        _ = v23
    case -60:
        switch "a" {
        case "a":
            delete(v1, 32767)
        default:
            switch ("\n\t" + "hello") {
            case "hello":
                v21 += 128
            case "foo bar":
                v24 := v2[v20]
                _ = v24
            default:
                v18 = append(v18, (v20 >> ((uint(v21)) & 63)))
            }
        }
    }
    var v25 float64 = float64((((v19 * (0.0 + float32((-2.25 * (1e10 / 1.0))))) + v19) - 1.5))
    fmt.Println(v3, v20)
    var v26 float32 = 0.1
    for fi27 := 0; fi27 < 2; fi27++ {
        v19 += (((v26 * v26) + 0.1) - (0.0 * 0.1))
        _ = fi27
    }
    var v28 uint16 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
}
