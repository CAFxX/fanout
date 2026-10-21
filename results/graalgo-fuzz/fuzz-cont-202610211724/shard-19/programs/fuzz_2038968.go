package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 string) int16 {
    p0 = "\n\t"
    var v5 float64 = 1e10
    if true {
        var v6 uint8 = 0
        var v7 uint = 42
        v7 = (v7 << (((v7 / (((v7 | v7)) | 1))) & 63))
        _ = v6
        _ = v7
    } else {
        p0 = "foo bar"
        v8 := []int{}
        _ = v8
    }
    var v9 uint8 = 156
    switch -219 {
    case 0:
        p0 = "hello"
    case 1:
        switch 32767 {
        case 2147483647:
            v5 += v5
        default:
            v10 := GenS1{F0: -713, F1: 42}
            _ = v10
        }
    }
    _ = p0
    _ = v5
    _ = v9
    return -214
}

func main() {
    var v1 int16 = 42
    var v2 int = 9223372036854775807
    var v3 float32 = 1.5
    var v4 uint = 4294967295
    var v11 int16 = v1
    v4 -= 4294967295
    v12 := []int{(v2 % ((v2) | 1))}
    _ = len(v12)
    switch ("foo bar" + "a") {
    case "foo bar":
        if (0 > 681) {
            v1 = v11
            var v13 uint32 = 944
            _ = v13
        } else {
            fmt.Println(v4)
            if true {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                v4 = v4
            }
        }
    case "\n\t":
        var v14 bool = (!("hello" != ""))
        _ = v14
    case "hello":
        for fi15 := 0; fi15 < 19; fi15++ {
            if (-51 > -128) {
                if (((v11 | genFn0("foo bar")) == v1) || (float32(0.0) < v3)) {
                    v16 := GenS2{F0: 100}
                    _ = v16
                } else {
                    var v17 string = ((("a" + "") + ("a" + "\n\t")) + "foo bar")
                    _ = v17
                }
                fmt.Println(v3, v3, v1)
            } else {
                var v18 int32 = -31
                v19 := GenS1{F0: (-(v18)), F1: 305}
                var v20 bool = false
                _ = v18
                _ = v19
                _ = v20
            }
            _ = fi15
        }
    }
    v1 = (v1 | int16(v2))
    var v21 int32 = -1
    var v22 float32 = (1e10 / 0.1)
    fmt.Println(v3, v21, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
