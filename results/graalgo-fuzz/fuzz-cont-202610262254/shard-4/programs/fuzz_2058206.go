package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int16) uint32 {
    for fi5 := 0; fi5 < 14; fi5++ {
        p0 -= p0
        _ = fi5
    }
    var v6 uint8 = 49
    var v7 uint = 173
    _ = p0
    _ = v6
    _ = v7
    return 2
}
func genFn1() int32 {
    if true {
        v8 := []int{-758, 42, -286}
        for fi9 := 0; fi9 < 16; fi9++ {
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 int = fi9
                v12 := GenS1{F0: v11, F1: 32767}
                var v13 string = ""
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            _ = fi9
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := GenS2{F0: 65535}
            var v16 int16 = 1000
            v17 := GenS2{F0: 255}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = v8
    }
    var v18 float64 = 1e10
    v18 += v18
    var v19 uint = 1
    _ = v18
    _ = v19
    return int32(v19)
}

func main() {
    var v1 uint16 = 2
    var v2 int32 = 42
    v3 := make(map[string]string)
    v4 := []int{596}
    v2 = v2
    if (0.0 <= 0.1) {
        fmt.Println(v1, v1)
    } else {
        var v20 int8 = 1
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v20
    }
    var v21 uint16 = 398
    v22 := GenS1{F0: 32767, F1: 974}
    var v23 string = "a"
    if (true && true) {
        var v24 int64 = -361
        v25 := GenS2{F0: uint(v1)}
        v1 += v21
        _ = v24
        _ = v25
    } else {
        fmt.Println(v2, v1)
    }
    v4 = append(v4, 128)
    v1 = (v21 & v21)
    switch "a" {
    case "foo bar":
        switch 32768 {
        case 42:
            v21 -= 255
        }
    case "\n\t":
        v4 = append(v4, 2147483647)
    }
    delete(v3, v23)
    var v26 string = ""
    switch v26 {
    case "a":
        fmt.Println(v26, v21)
    }
    v27 := GenS3{F0: 128}
    fmt.Println(v21)
    var v28 int64 = 0
    switch -9223372036854775808 {
    case -827:
        v1 += (v21 | (v1 | v21))
    case 32767:
        delete(v3, v23)
    }
    v23 = v23
    if (2 < 958) {
        fmt.Println(v1, v26, v23)
        var v29 uint8 = 255
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
