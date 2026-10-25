package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int }
func genFn0(p0 int32) bool {
    v6 := GenS1{F0: "", F1: 42, F2: 42}
    switch int(p0) {
    case -9223372036854775808:
        fmt.Println(p0)
    }
    v7 := GenS1{F0: "hello", F1: -2147483648, F2: 144}
    switch 448 {
    case 1:
        for fi8 := 0; fi8 < 2; fi8++ {
            p0 += 42
            fmt.Println(p0, p0)
            _ = fi8
        }
    default:
        for fi9 := 0; fi9 < 2; fi9++ {
            p0 -= 42
            if true {
                p0 += 2147483647
                p0 = (p0 % (((p0 % ((1) | 1))) | 1))
            }
            v10 := GenS2{F0: 1}
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 int32, p1 int, p2 uint) uint8 {
    var v11 float64 = 0.1
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 255
}

func main() {
    var v1 uint8 = 189
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    var v4 string = "foo bar"
    v5 := []int{-386, 32767}
    var v12 bool = false
    v1 += genFn1(-677, -76, 255)
    v13 := GenS3{F0: false}
    v12 = (1e10 == (1.5 * ((3.14159 + float32((((1.5 - -2.25) - ((1e10 + 0.1) / 1.5)) / 1e10))) * (0.0 + ((0.1 - 1.5) + 1e10)))))
    if (1 < -261) {
        switch v4 {
        case "":
            var v14 bool = (352 != uint16(v1))
            _ = v14
        }
        fmt.Println(v12)
        delete(v2, -2147483648)
    } else {
        var v15 int32 = -1
        _ = v15
    }
    var v16 uint8 = 2
    switch ((v4 + "foo bar") + (v4 + (v4 + "foo bar"))) {
    case "a":
        for fi17 := 0; fi17 < 16; fi17++ {
            if false {
                v12 = false
                var v18 int = (fi17 - (fi17 / ((272) | 1)))
                _ = v18
            }
            v19 := GenS4{F0: ("" + "foo bar"), F1: 42}
            var v20 string = v4
            _ = fi17
            _ = v19
            _ = v20
        }
    case "hello":
        if (217 <= (v16 ^ (^(v1)))) {
            v4 = (v4 + (v4 + "a"))
        } else {
            var v21 string = "a"
            v21 = v21
            _ = v21
        }
    default:
        switch 391 {
        case 127:
            switch 2147483647 {
            case -453:
                v22 := GenS1{F0: v4, F1: 528, F2: (v16 + (v1 + genFn1(-646, -388, 765)))}
                _ = v22
            default:
                v16 -= genFn1(912, 128, 1)
            }
        }
    }
    v23 := GenS5{F0: int(v1)}
    v1 = v16
    v16 += (^(v16))
    fmt.Println(v1)
    var v24 float32 = (float32((0.0 / -2.25)) + (1.5 / 1e10))
    var v25 uint = 255
    v2[2147483647] = 4294967295
    for fi26 := 0; fi26 < 19; fi26++ {
        if ((v12 && (!(false || (236 < 0)))) && v12) {
            var v27 uint = v25
            _ = v27
        }
        _ = fi26
    }
    var v28 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
