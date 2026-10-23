package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int }
func genFn0() int32 {
    v5 := []int{-9223372036854775808, 42}
    var v6 int64 = 850
    _ = v5
    _ = v6
    return -883
}
func genFn1(p0 int16, p1 string) uint {
    var v7 bool = (4294967295 > 548)
    var v8 bool = v7
    if (1e10 < (0.0 + 1.5)) {
        fmt.Println(v7)
        v8 = (-1 == 42)
    } else {
        v9 := []int{-831, -9223372036854775808, 128, -2147483648}
        _ = v9
    }
    var v10 int = -453
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    return 651
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 42
    var v3 float64 = 0.0
    v4 := make(map[string]int)
    var v11 float32 = ((1.5 / (-2.25 * ((3.14159 * (1.5 + (0.1 + 1.5))) * 0.1))) * 3.14159)
    if (0 <= 18446744073709551615) {
        var v12 float32 = (((v11 - v11) * ((float32(v3) - v11) - v11)) - v11)
        v13 := GenS1{F0: 2147483647, F1: 4294967295}
        var v14 int64 = 9223372036854775807
        _ = v12
        _ = v13
        _ = v14
    } else {
        fmt.Println(v3, v2)
        fmt.Println(v2)
    }
    var v15 uint = 255
    switch ("" + "a") {
    case "foo bar":
        v15 += 1
    case "hello":
        var v16 uint16 = 65535
        _ = v16
    case "":
        for fi17 := 0; fi17 < 4; fi17++ {
            if (-49 <= 100) {
                v15 = (v15 ^ 582)
                fmt.Println(fi17)
            } else {
                var v18 string = ((("foo bar" + "hello") + "foo bar") + "foo bar")
                fmt.Println(v2, v3)
                _ = v18
            }
            if false {
                v19 := GenS1{F0: fi17, F1: 42}
                fmt.Println(v3, fi17, fi17)
                v2 += v2
                _ = v19
            } else {
                fmt.Println(fi17, v2, v3)
            }
            _ = fi17
        }
    }
    v11 = v11
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int = fi20
        v22 := GenS2{F0: (fi20 | -661)}
        v23 := []int{(fi20 / (((v21 | 1)) | 1))}
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    v3 = v3
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 int = -854
        var v26 uint16 = 42
        _ = fi24
        _ = v25
        _ = v26
    }
    v3 = v3
    v1[""] = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
