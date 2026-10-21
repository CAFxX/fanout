package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 bool; F2 string }
type GenS5 struct { F0 uint8; F1 string; F2 uint8 }

func main() {
    v1 := []string{"\n\t", ("\n\t" + "foo bar"), "a", "\n\t"}
    var v2 uint8 = 42
    v3 := make(map[int]bool)
    var v4 int = -9223372036854775808
    var v5 uint32 = 589
    var v6 float64 = -2.25
    var v7 uint32 = v5
    if (-2.25 < 1e10) {
        v1 = append(v1, "foo bar")
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(fi8, v4, v4)
            var v9 string = "\n\t"
            v10 := GenS1{F0: (v4 * fi8), F1: fi8}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v11 := GenS2{F0: v4}
    var v12 int32 = 42
    if true {
        if (128 > v2) {
            var v13 string = "a"
            v2 = (v2 % (((v2 / ((v2) | 1))) | 1))
            v14 := GenS2{F0: v4}
            _ = v13
            _ = v14
        }
        v5 = 2147483648
    } else {
        v2 = v2
        v3[4294967295] = (true || (((((v6 / (v6 + v6)) - v6) * 0.0) / 1.0) <= (1e10 - float64(1.5))))
    }
    v3[(v4 - v4)] = (false && false)
    switch (v4 >> ((255) & 63)) {
    case 978:
        delete(v3, (v4 % (((v4 * v4)) | 1)))
    case -2147483648:
        if true {
            v6 -= 3.14159
            v15 := GenS3{F0: 32767}
            _ = v15
        }
    case -931:
        switch "\n\t" {
        case "foo bar":
            var v16 uint16 = 2
            _ = v16
        case "":
            var v17 uint32 = v5
            _ = v17
        case "hello":
            fmt.Println(v12, v7)
        }
    }
    v18 := GenS4{F0: 255, F1: false, F2: "hello"}
    var v19 uint16 = 16
    v20 := GenS5{F0: 0, F1: (("a" + "a") + ""), F2: v2}
    switch "a" {
    case "\n\t":
        v21 := GenS1{F0: v4, F1: (v4 % (((v4 >> ((255) & 63))) | 1))}
        _ = v21
    case "a":
        var v22 string = "hello"
        _ = v22
    default:
        _ = len(v1)
    }
    fmt.Println(v2, v12)
    var v23 uint8 = (v2 % (((v2 - v2)) | 1))
    var v24 float64 = 1.5
    fmt.Println(v6, v23, v6)
    var v25 string = ("" + ("" + ("a" + "a")))
    v4 += v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
