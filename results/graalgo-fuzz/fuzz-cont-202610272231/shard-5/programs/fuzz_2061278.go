package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint64, p1 string) bool {
    switch ("hello" + "hello") {
    case "hello":
        p0 -= p0
    case "foo bar":
        var v4 float64 = 3.14159
        _ = v4
    case "\n\t":
        p1 = "\n\t"
    }
    v5 := GenS1{F0: 42, F1: (p1 + p1), F2: (p0 << ((uint(p0)) & 63))}
    p0 -= (p0 + p0)
    v6 := GenS1{F0: 0, F1: "hello", F2: (^(p0))}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return true
}
func genFn1() int8 {
    v7 := GenS1{F0: -2147483648, F1: ("a" + "foo bar"), F2: 1}
    if true {
        switch -873 {
        case 127:
            v8 := []int{-9223372036854775808, -2147483648, 555}
            _ = v8
        }
    }
    _ = v7
    return 127
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = 77
    var v3 uint = 2
    switch "a" {
    case "foo bar":
        v3 = 0
    default:
        switch -343 {
        case 128:
            for fi9 := 0; fi9 < 7; fi9++ {
                v3 -= 255
                _ = fi9
            }
        case 729:
            v10 := []int{164, 640, 42}
            _ = v10
        case -49:
            fmt.Println(v3, v3)
        default:
            v11 := []int{-303, 127}
            _ = v11
        }
    }
    var v12 uint8 = 2
    if true {
        v13 := []int{1, 1, int(v12)}
        fmt.Println(v1, v2)
        if (-1 >= -9223372036854775808) {
            _ = len(v13)
            var v14 int = 33
            for fi15 := 0; fi15 < 5; fi15++ {
                v16 := GenS2{F0: (^(v12))}
                _ = fi15
                _ = v16
            }
            _ = v14
        } else {
            if false {
                v17 := GenS3{F0: 1}
                var v18 uint8 = v12
                _ = v17
                _ = v18
            }
        }
        _ = v13
    }
    v19 := []int{-9223372036854775808, 32767, 1, 42}
    fmt.Println(v2, v12, v12)
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 uint16 = 32767
        if false {
            v22 := GenS2{F0: v12}
            switch fi20 {
            case 2147483647:
                var v23 bool = (!(v12 <= v12))
                _ = v23
            }
            _ = v22
        } else {
            v2 += v2
            v24 := GenS1{F0: 1, F1: (v1 + (v1 + "a")), F2: 522}
            _ = v24
        }
        _ = fi20
        _ = v21
    }
    fmt.Println(v12, v1, v1)
    var v25 string = "a"
    fmt.Println(v2, v3)
    fmt.Println(v12, v1)
    var v26 int32 = 1
    var v27 uint = (v3 - v3)
    var v28 int16 = (v2 << (((v3 + (v3 - v27))) & 63))
    v28 = (-(v28))
    v3 = v3
    var v29 int32 = v26
    var v30 int = -82
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
