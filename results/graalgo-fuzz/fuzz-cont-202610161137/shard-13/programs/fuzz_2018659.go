package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int8; F1 uint8 }
func genFn0(p0 bool) int32 {
    var v4 bool = (!(true && (2 < 781)))
    v4 = v4
    var v5 float32 = (float32((3.14159 / 1.0)) / 1.0)
    fmt.Println(v5)
    _ = p0
    _ = v4
    _ = v5
    return 303
}

func main() {
    var v1 uint = 64
    var v2 int = 127
    v3 := make(map[string]string)
    if ((^(v2)) != -493) {
        fmt.Println(v2, v2)
    } else {
        var v6 int16 = 0
        _ = v6
    }
    if ((v2 & -611) < 127) {
        if true {
            fmt.Println(v2, v2)
            var v7 bool = true
            _ = v7
        } else {
            var v8 int16 = 42
            _ = v8
        }
        var v9 int32 = -2147483648
        v10 := GenS1{F0: -1}
        _ = v10
        _ = v9
    }
    var v11 int32 = -1
    v12 := []int{}
    v13 := GenS2{F0: (v2 + v2), F1: 283}
    switch "\n\t" {
    case "":
        var v14 int = 1
        _ = v14
    }
    switch "hello" {
    case "hello":
        delete(v3, "foo bar")
    }
    v15 := v3[("hello" + "\n\t")]
    fmt.Println(v11, v11)
    var v16 float32 = 1.5
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 uint32 = 2
        if (42 >= 127) {
            var v19 float64 = 0.1
            if true {
                var v20 int32 = (v11 % ((2147483647) | 1))
                _ = v20
            }
            v21 := GenS3{F0: (v1 ^ v1)}
            _ = v19
            _ = v21
        }
        switch fi17 {
        case 1:
            if true {
                fmt.Println(fi17)
            } else {
                var v22 int32 = genFn0((true && false))
                if (!true) {
                    if len(v12) > 0 {
                        v12[0] = 0
                    }
                    var v23 int64 = -401
                    _ = v23
                }
                v24 := GenS1{F0: 0}
                _ = v22
                _ = v24
            }
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v2, v11)
    v16 -= 1.5
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    if (-1 < -2147483648) {
        v25 := GenS4{F0: true, F1: 0}
        var v26 uint32 = 366
        _ = v25
        _ = v26
    }
    switch ("a" + v15) {
    case "hello":
        v27 := GenS5{F0: 25, F1: 1}
        _ = v27
    case "":
        _ = len(v3)
    case "foo bar":
        switch ("" + (v15 + v15)) {
        case "a":
            v1 -= v1
        case "hello":
            v15 = v15
        }
    }
    if false {
        var v28 uint32 = 968
        if (732 > 1) {
            var v29 int32 = genFn0((v15 != v15))
            fmt.Println(v11, v15)
            v1 = 42
            _ = v29
        }
        _ = len(v3)
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
