package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int32; F2 string }
func genFn0() uint16 {
    var v6 int8 = 1
    v6 += v6
    _ = v6
    return 255
}
func genFn1(p0 int8) uint8 {
    var v7 bool = false
    v8 := []int{-2147483648}
    v9 := GenS1{F0: 162}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]bool)
    var v3 int8 = 42
    var v4 uint8 = 42
    v5 := make(map[uint32]bool)
    var v10 bool = (("\n\t" >= "\n\t") && false)
    v11 := GenS1{F0: 606}
    for fi12 := 0; fi12 < 12; fi12++ {
        v1 = (("\n\t" + (v1 + v1)) + ("\n\t" + v1))
        v10 = (fi12 < (fi12 * (^(fi12))))
        _ = fi12
    }
    var v13 int16 = 32767
    v1 = v1
    v14 := []int{128}
    switch 854 {
    case 484:
        v4 = genFn1(42)
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        v3 = (v3 ^ (^(v3)))
        fmt.Println(v3)
        _ = fi15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        delete(v5, 849)
        delete(v5, 0)
        if v10 {
            var v17 uint8 = v4
            _ = v17
        }
        _ = fi16
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        fmt.Println(fi18)
        v1 = ("hello" + (("\n\t" + ((v1 + ("hello" + (v1 + (((v1 + ((v1 + ((((v1 + ((v1 + v1) + v1)) + v1) + v1) + v1)) + v1)) + v1) + v1)))) + "foo bar")) + (v1 + v1)))
        var v19 uint16 = uint16(v4)
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        if (!(-9223372036854775808 >= 23)) {
            fmt.Println(v4)
        }
        v21 := GenS2{F0: 914}
        _ = fi20
        _ = v21
    }
    v10 = v10
    var v22 float64 = (3.14159 - 3.14159)
    v23 := GenS3{F0: true, F1: -1, F2: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
