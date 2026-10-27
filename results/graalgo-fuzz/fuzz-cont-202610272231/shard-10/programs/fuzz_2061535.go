package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }
type GenS2 struct { F0 int32; F1 int; F2 int32 }
func genFn0(p0 int64) int32 {
    var v7 uint32 = 4294967295
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(fi8, fi8, v7)
        for fi9 := 0; fi9 < 8; fi9++ {
            v7 -= v7
            if (-988 <= -1) {
                if true {
                    fmt.Println(p0, p0)
                    v7 = (v7 / ((234) | 1))
                    v7 -= v7
                    fmt.Println(p0, p0, fi8)
                } else {
                    fmt.Println(v7)
                    switch "a" {
                    case "a":
                        v10 := []int{(fi8 - fi9), fi9}
                        _ = v10
                    case "foo bar":
                        v7 -= v7
                    }
                }
            }
            _ = fi9
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return -397
}
func genFn1() uint16 {
    var v11 int = 32768
    switch ("hello" + ("\n\t" + ("a" + ("a" + (("hello" + (("foo bar" + "foo bar") + "a")) + "hello"))))) {
    case "\n\t":
        var v12 bool = (false || true)
        _ = v12
    default:
        var v13 uint32 = 2
        _ = v13
    }
    var v14 uint8 = 1
    _ = v11
    _ = v14
    return 704
}

func main() {
    var v1 bool = false
    var v2 int64 = 9223372036854775807
    v3 := []int{964, 1, -260}
    var v4 string = ("\n\t" + (("foo bar" + "\n\t") + "\n\t"))
    var v5 int = 840
    v6 := []string{v4, "", v4, v4}
    v15 := GenS1{F0: 190, F1: ((("" + v4) + "hello") + "hello"), F2: -30}
    switch v4 {
    case "a":
        var v16 float32 = float32(0.1)
        _ = v16
    case "foo bar":
        var v17 float64 = 1e10
        _ = v17
    case "\n\t":
        switch v5 {
        case 32768:
            if v1 {
                var v18 float64 = 1.5
                fmt.Println(v18)
                var v19 float32 = 0.1
                _ = v18
                _ = v19
            }
        default:
            v1 = ((!true) && (((255 >= 154) && (!false)) && ((!false) && ((v2 | (^(v2))) < 1))))
        }
    }
    if len(v6) > 0 {
        v6[0] = ""
    }
    v1 = true
    v1 = (true && v1)
    v6 = append(v6, "\n\t")
    v20 := GenS2{F0: -491, F1: 32768, F2: -1}
    v5 += v5
    var v21 int8 = -1
    var v22 int8 = v21
    v2 = -1
    var v23 bool = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
