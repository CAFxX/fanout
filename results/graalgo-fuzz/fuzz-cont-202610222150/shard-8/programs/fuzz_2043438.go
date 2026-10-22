package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint16; F2 int }

func main() {
    v1 := []uint32{2147483648, 2147483647, 58}
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    v4 := []int{-1}
    var v5 uint8 = 0
    if true {
        if (true || (!(((("foo bar" + "a") + ("foo bar" + ((("hello" + ("hello" + "\n\t")) + ("hello" + (("a" + "hello") + (("" + "a") + (("foo bar" + "\n\t") + "\n\t"))))) + ("\n\t" + "a")))) + "hello") == ("a" + "hello")))) {
            v6 := GenS1{F0: (false || true), F1: 0}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 float64 = (-2.25 + -2.25)
                var v9 string = "hello"
                _ = fi7
                _ = v8
                _ = v9
            }
            v5 = (v5 >> ((4294967295) & 63))
            if len(v1) > 0 {
                v1[0] = 93
            }
        }
        var v10 int32 = -846
        if false {
            if len(v4) > 0 {
                v4[0] = 32768
            }
        }
        _ = v10
    }
    var v11 uint8 = (v5 << ((0) & 63))
    v12 := GenS2{F0: uint(v5)}
    if true {
        v11 += (v11 >> ((uint(v5)) & 63))
    } else {
        _ = len(v4)
    }
    v13 := GenS3{F0: true, F1: 959, F2: 32767}
    v14 := GenS3{F0: (false && false), F1: 1, F2: -1}
    v15 := GenS3{F0: true, F1: 1, F2: 128}
    v16 := GenS1{F0: (65535 > 0), F1: 571}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
