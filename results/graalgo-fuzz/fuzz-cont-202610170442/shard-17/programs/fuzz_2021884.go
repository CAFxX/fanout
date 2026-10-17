package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8; F1 bool; F2 uint32 }
type GenS3 struct { F0 string }
func genFn0(p0 int, p1 bool) string {
    for fi4 := 0; fi4 < 7; fi4++ {
        p1 = true
        _ = fi4
    }
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 1, F1: 1}
        v7 := GenS1{F0: 18446744073709551615, F1: 1}
        _ = fi5
        _ = v6
        _ = v7
    }
    if (!true) {
        if p1 {
            switch (p0 & 9223372036854775807) {
            case -432:
                p1 = ((p0 ^ p0) <= -1)
            }
            v8 := []int{-2147483648, p0, p0}
            _ = v8
        }
        var v9 int = 948
        _ = v9
    }
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    var v1 string = "hello"
    var v2 int64 = 1
    v3 := make(map[int]bool)
    v10 := v3[42]
    v11 := []int{32767, -544, 32768}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v2 = (-(v2))
    fmt.Println(v1)
    for fi12 := 0; fi12 < 12; fi12++ {
        v2 = v2
        var v13 int = (fi12 / ((2147483647) | 1))
        var v14 uint = 221
        _ = fi12
        _ = v13
        _ = v14
    }
    v2 = -226
    switch v1 {
    case "a":
        if v10 {
            if v10 {
                fmt.Println(v1, v2)
                v15 := GenS2{F0: 127, F1: (!(4294967295 <= 110)), F2: 2}
                var v16 int64 = -1
                _ = v15
                _ = v16
            } else {
                var v17 int32 = 42
                _ = v17
            }
            v18 := GenS1{F0: 1, F1: 42}
            _ = v18
        }
    case "hello":
        for fi19 := 0; fi19 < 13; fi19++ {
            for fi20 := 0; fi20 < 11; fi20++ {
                fmt.Println(v1, v2, fi19)
                _ = len(v3)
                v3[fi20] = (!(42 < 255))
                _ = fi20
            }
            var v21 uint32 = 521
            switch (v1 + v1) {
            case "foo bar":
                fmt.Println(v21)
            case "\n\t":
                var v22 uint = 431
                _ = v22
            default:
                switch "\n\t" {
                case "a":
                    v10 = (-2.25 > (-2.25 * ((1.5 - 1.5) * 1e10)))
                case "hello":
                    fmt.Println(v2, v21)
                case "foo bar":
                    v1 = ("foo bar" + v1)
                }
            }
            _ = fi19
            _ = v21
        }
    }
    v23 := GenS3{F0: v1}
    _ = len(v11)
    v1 = ((v1 + v1) + v1)
    v24 := v3[-2147483648]
    var v25 float32 = float32(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
