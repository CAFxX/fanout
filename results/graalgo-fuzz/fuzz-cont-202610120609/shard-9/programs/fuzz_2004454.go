package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 uint32, p1 int64) bool {
    var v5 int64 = 9223372036854775807
    var v6 uint16 = 2
    var v7 string = (("a" + (((("foo bar" + ("hello" + "a")) + "foo bar") + ("a" + "hello")) + "")) + "\n\t")
    v7 = ("hello" + v7)
    var v8 uint64 = 132
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (!true)
}
func genFn1() uint {
    if false {
        var v9 int32 = -1
        v10 := []int{42, 32768, 32767, -9223372036854775808}
        _ = v10
        _ = v9
    }
    v11 := []int{}
    var v12 int32 = 2147483647
    v12 -= v12
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -630
    var v3 uint8 = 127
    var v4 uint32 = 2147483647
    v4 -= (v4 + (v4 - (v4 | (v4 + 182))))
    v2 += v2
    if true {
        switch v2 {
        case 128:
            _ = len(v1)
        case -1:
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := []int{}
                var v15 string = "a"
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        fmt.Println(v2, v3, v3)
    } else {
        v16 := GenS1{F0: 2147483647}
        _ = v16
    }
    v17 := GenS1{F0: (v4 >> ((76) & 63))}
    var v18 float64 = float64(-2.25)
    v4 = v4
    var v19 bool = (32767 <= int16(v3))
    var v20 int = v2
    for fi21 := 0; fi21 < 18; fi21++ {
        switch -442 {
        case -2147483648:
            v22 := v1[(v4 ^ 455)]
            _ = v22
        }
        _ = fi21
    }
    v23 := GenS2{F0: -128, F1: 42}
    var v24 int64 = 465
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
