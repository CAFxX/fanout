package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int; F2 int16 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int }
func genFn0() uint {
    v4 := GenS1{F0: 0}
    var v5 uint64 = 2
    _ = v4
    _ = v5
    return 255
}
func genFn1(p0 int64, p1 int16) int32 {
    if true {
        v6 := GenS2{F0: 38, F1: -731, F2: p1}
        v7 := GenS3{F0: 100, F1: -406}
        _ = v6
        _ = v7
    } else {
        fmt.Println(p0, p1)
        p1 = 42
        switch "hello" {
        case "hello":
            var v8 string = ("hello" + ("a" + ("foo bar" + ("foo bar" + "hello"))))
            _ = v8
        }
    }
    var v9 int32 = 42
    v10 := []int{int(v9)}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (v9 >> ((1) & 63))
}

func main() {
    var v1 bool = (213 == 42)
    v2 := make(map[uint32]string)
    var v3 uint32 = 2147483648
    var v11 int32 = -1
    if (4294967295 <= v3) {
        fmt.Println(v3, v1)
        v12 := []int{-2147483648, 32768, 1}
        _ = v12
    } else {
        v13 := GenS3{F0: -105, F1: v11}
        var v14 uint64 = 0
        _ = v13
        _ = v14
    }
    var v15 uint = genFn0()
    if true {
        for fi16 := 0; fi16 < 5; fi16++ {
            if true {
                var v17 uint16 = 518
                v11 = v11
                fmt.Println(v1, v1, fi16)
                _ = v17
            }
            v3 = (v3 - (v3 | v3))
            _ = fi16
        }
    }
    v11 -= v11
    v18 := []int{-2147483648}
    var v19 string = "hello"
    v20 := GenS3{F0: -51, F1: -428}
    switch "" {
    case "\n\t":
        if (v1 || (!(!false))) {
            v21 := GenS4{F0: -2147483648}
            _ = v21
        } else {
            switch "" {
            case "foo bar":
                fmt.Println(v19)
            default:
                var v22 int16 = 32767
                _ = v22
            }
            v19 = v19
        }
    case "foo bar":
        v23 := v2[(v3 ^ (^(v3)))]
        _ = v23
    case "":
        switch ((v19 + ((("a" + v19) + v19) + (v19 + "foo bar"))) + "a") {
        case "foo bar":
            _ = len(v18)
        case "":
            if v1 {
                var v24 float64 = 1e10
                var v25 int64 = -824
                for fi26 := 0; fi26 < 12; fi26++ {
                    v24 = float64((0.0 - 0.0))
                    _ = fi26
                }
                _ = v24
                _ = v25
            }
        case "hello":
            var v27 uint16 = 956
            _ = v27
        }
    default:
        fmt.Println(v15, v11, v11)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
