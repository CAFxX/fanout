package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{761}
    var v3 string = "a"
    var v4 uint16 = 367
    if (42 >= 0) {
        v5 := GenS1{F0: true}
        fmt.Println(v3, v4)
        _ = v5
    } else {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int32 = 2147483647
            var v8 int = fi6
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 float64 = 3.14159
            fmt.Println(fi9, v4, fi9)
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS2{F0: (!(!true)), F1: 127}
    for fi12 := 0; fi12 < 3; fi12++ {
        if false {
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 int32 = 1
                switch ("" + v3) {
                case "foo bar":
                    if true {
                        v3 = "a"
                    } else {
                        v3 = (v3 + v3)
                        var v15 int8 = 55
                        fmt.Println(v15)
                        _ = v15
                    }
                case "a":
                    v3 = (v3 + "a")
                }
                fmt.Println(fi12)
                _ = fi13
                _ = v14
            }
            v4 += (v4 + (v4 & (v4 * (v4 | v4))))
            var v16 uint16 = v4
            _ = v16
        }
        v3 = (((v3 + (v3 + ((((("" + v3) + (v3 + v3)) + ("hello" + v3)) + v3) + ("a" + v3)))) + ((v3 + (v3 + (v3 + v3))) + (("hello" + "foo bar") + v3))) + v3)
        for fi17 := 0; fi17 < 10; fi17++ {
            v4 += 32767
            fmt.Println(fi12, v3, fi12)
            fmt.Println(v4, fi12, fi12)
            _ = fi17
        }
        _ = fi12
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        fmt.Println(fi18, v4, v4)
        var v19 int8 = 42
        _ = fi18
        _ = v19
    }
    _ = len(v2)
    v3 = (("a" + "\n\t") + ("a" + v3))
    var v20 float64 = -2.25
    switch v3 {
    case "a":
        v4 -= (v4 ^ (v4 << ((4294967295) & 63)))
    case "hello":
        var v21 int8 = -128
        _ = v21
    case "foo bar":
        _ = len(v2)
    default:
        var v22 uint16 = 42
        _ = v22
    }
    switch ((v3 + (v3 + "a")) + v3) {
    case "foo bar":
        v23 := GenS3{F0: -1}
        _ = v23
    case "":
        var v24 uint = uint(v4)
        _ = v24
    case "a":
        if (18446744073709551615 <= 42) {
            var v25 float32 = (0.1 - 1.5)
            _ = v25
        } else {
            var v26 float32 = 1.5
            _ = v26
        }
    default:
        v27 := GenS4{F0: 2147483648, F1: -1}
        _ = v27
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
