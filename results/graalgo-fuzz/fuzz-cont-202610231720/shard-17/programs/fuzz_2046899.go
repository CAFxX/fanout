package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0(p0 int64, p1 uint8, p2 int16) uint8 {
    v4 := GenS1{F0: 42}
    if true {
        v5 := []int{}
        if (true && false) {
            switch 128 {
            case -660:
                if (!(32767 != 32767)) {
                    p0 = 42
                    if true {
                        var v6 string = "a"
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v6
                    }
                    if true {
                        _ = len(v5)
                    } else {
                        var v7 uint32 = 42
                        _ = v7
                    }
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(p2, p0, p0)
                    }
                    fmt.Println(p1, p2, p1)
                }
            case 0:
                p1 = 252
            default:
                v8 := GenS2{F0: 342, F1: 0}
                _ = v8
            }
        } else {
            var v9 int16 = int16(p1)
            fmt.Println(v9)
            v10 := GenS2{F0: 484, F1: -128}
            _ = v10
            _ = v9
        }
        _ = len(v5)
        _ = v5
    }
    v11 := []int{2147483647}
    p1 -= (p1 - (p1 / ((p1) | 1)))
    var v12 uint8 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v4
    return 1
}
func genFn1(p0 uint16, p1 uint64) int16 {
    var v13 string = ("\n\t" + "a")
    v14 := GenS1{F0: 1}
    p1 = p1
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return -32768
}

func main() {
    var v1 int64 = 163
    var v2 string = "hello"
    var v3 float32 = 3.14159
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 uint64 = 548
        var v17 string = (v2 + v2)
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := []int{-766, 9223372036854775807, 0, -921}
    fmt.Println(v1, v2)
    var v19 string = "foo bar"
    var v20 int16 = 32767
    for fi21 := 0; fi21 < 5; fi21++ {
        v20 = v20
        _ = fi21
    }
    if (((!true) && (92 > 127)) || true) {
        var v22 int64 = (v1 | (^(v1)))
        v19 = (v2 + v19)
        _ = v22
    } else {
        if true {
            fmt.Println(v1)
        }
        var v23 uint8 = 128
        var v24 uint16 = 65535
        _ = v23
        _ = v24
    }
    var v25 bool = false
    v3 = v3
    if (!(v3 == (-2.25 + (v3 - float32((3.14159 * 0.1)))))) {
        _ = len(v18)
    } else {
        _ = len(v18)
    }
    fmt.Println(v19, v20, v25)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
}
