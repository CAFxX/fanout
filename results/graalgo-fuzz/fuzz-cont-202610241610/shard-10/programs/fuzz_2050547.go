package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint }
func genFn0(p0 string) int32 {
    switch ("foo bar" + p0) {
    case "foo bar":
        v6 := GenS1{F0: 42, F1: -1}
        _ = v6
    }
    switch ((p0 + (p0 + ((p0 + p0) + p0))) + p0) {
    case "":
        if (944 < 2147483647) {
            for fi7 := 0; fi7 < 4; fi7++ {
                fmt.Println(p0, fi7)
                fmt.Println(fi7, p0)
                _ = fi7
            }
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 bool = false
                for fi10 := 0; fi10 < 16; fi10++ {
                    switch p0 {
                    case "foo bar":
                        fmt.Println(fi8, fi10, p0)
                        v9 = v9
                    case "hello":
                        fmt.Println(fi10, v9)
                    case "a":
                        var v11 int16 = -1
                        var v12 uint64 = 0
                        _ = v11
                        _ = v12
                    default:
                        v9 = true
                    }
                    _ = fi10
                }
                fmt.Println(v9, fi8, fi8)
                _ = fi8
                _ = v9
            }
            if ((24 <= 29) && true) {
                var v13 int8 = 127
                fmt.Println(v13, p0)
                _ = v13
            } else {
                v14 := []int{459, -2147483648, -888}
                _ = v14
            }
        }
    default:
        p0 = "hello"
    }
    p0 = (("" + p0) + p0)
    _ = p0
    return -839
}

func main() {
    var v1 int64 = -1
    v2 := []int{0, 992, 128, 719}
    var v3 uint16 = 255
    var v4 uint16 = (v3 >> ((142) & 63))
    var v5 uint8 = 255
    if true {
        if false {
            for fi15 := 0; fi15 < 15; fi15++ {
                fmt.Println(v5, v3)
                v16 := GenS1{F0: 455, F1: fi15}
                for fi17 := 0; fi17 < 14; fi17++ {
                    v4 = 956
                    v3 = (v4 / ((670) | 1))
                    _ = fi17
                }
                _ = fi15
                _ = v16
            }
            v18 := GenS2{F0: uint32(v5), F1: -32768}
            _ = v18
        }
        if len(v2) > 0 {
            v2[0] = -298
        }
    }
    var v19 string = ("hello" + ("\n\t" + ("" + ((("a" + "") + (("a" + ("" + ("hello" + "a"))) + "hello")) + "a"))))
    var v20 float32 = 3.14159
    switch (v19 + v19) {
    case "\n\t":
        v21 := GenS1{F0: 0, F1: 128}
        _ = v21
    default:
        var v22 uint16 = v3
        _ = v22
    }
    if false {
        v23 := GenS3{F0: -9223372036854775808, F1: (^(v4))}
        v24 := GenS4{F0: v1, F1: v4}
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: 226, F1: 42}
    for fi26 := 0; fi26 < 17; fi26++ {
        if false {
            var v27 float32 = (3.14159 / 1.0)
            var v28 string = ((v19 + "foo bar") + (("hello" + v19) + v19))
            _ = v27
            _ = v28
        } else {
            v29 := GenS5{F0: 973}
            _ = v29
        }
        _ = fi26
    }
    v3 -= v3
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
