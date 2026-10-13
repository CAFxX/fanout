package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int16 }
func genFn0(p0 int8) uint64 {
    var v5 float64 = (1e10 * (1e10 + 1.5))
    var v6 int = -1
    v7 := []int{}
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: 2147483647, F1: p0, F2: "hello"}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 236
    v2 := make(map[string]bool)
    var v3 string = ((("a" + "hello") + ("foo bar" + "a")) + "\n\t")
    v4 := make(map[string]bool)
    v10 := []int{-610, -329, 1, 100}
    var v11 int64 = -1
    if ((45 <= 0) && (502 < (v11 >> ((42) & 63)))) {
        var v12 float32 = 1.5
        _ = v12
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if len(v10) > 0 {
        v10[0] = 127
    }
    if (!(!false)) {
        fmt.Println(v3)
        v11 -= (^(v11))
        v3 = "\n\t"
    }
    v1 = v1
    switch "a" {
    case "\n\t":
        switch ("hello" + (("foo bar" + ((((("hello" + v3) + v3) + v3) + "a") + v3)) + v3)) {
        case "\n\t":
            v4[(v3 + "a")] = true
        case "hello":
            switch "\n\t" {
            case "a":
                fmt.Println(v3)
            }
        case "":
            var v13 uint16 = 944
            _ = v13
        default:
            if false {
                if (352 != genFn0(1)) {
                    fmt.Println(v3, v11)
                    var v14 int8 = 1
                    _ = v14
                }
                fmt.Println(v11, v1)
                fmt.Println(v3, v3)
            } else {
                v3 = (((v3 + (v3 + "foo bar")) + v3) + (v3 + v3))
                fmt.Println(v11, v11, v1)
            }
        }
    default:
        fmt.Println(v1, v3, v1)
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(v11, v1)
        _ = fi15
    }
    switch 903 {
    case -2147483648:
        var v16 float32 = (1.5 / ((0.1 - 3.14159) / float32(1e10)))
        _ = v16
    case 0:
        var v17 int32 = -770
        _ = v17
    }
    var v18 uint = 918
    delete(v2, v3)
    if (true && false) {
        var v19 float64 = (1e10 * 0.1)
        v20 := GenS2{F0: true}
        var v21 int16 = 32767
        _ = v19
        _ = v20
        _ = v21
    } else {
        v22 := GenS3{F0: 65535, F1: 32767}
        _ = v22
    }
    if (true && true) {
        fmt.Println(v1, v1)
        v18 = (v18 - 18446744073709551615)
    }
    var v23 int8 = 1
    v24 := GenS1{F0: 0, F1: v23, F2: v3}
    fmt.Println(v11, v11, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
