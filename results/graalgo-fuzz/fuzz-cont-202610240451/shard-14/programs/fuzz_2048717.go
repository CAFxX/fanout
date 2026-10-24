package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int16 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int16; F1 int32; F2 uint }
func genFn0(p0 uint) int16 {
    switch 128 {
    case 0:
        fmt.Println(p0)
    case 4294967295:
        var v5 int16 = 1
        _ = v5
    case -134:
        fmt.Println(p0)
    }
    v6 := []int{837, 704}
    var v7 bool = (p0 == p0)
    var v8 int64 = int64(p0)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint8) uint {
    v9 := []int{2147483647, 4294967295, 32767, 0}
    switch (("a" + ((("\n\t" + "\n\t") + "a") + "\n\t")) + "foo bar") {
    case "a":
        v9 = append(v9, -2147483648)
    case "\n\t":
        var v10 float32 = ((3.14159 * -2.25) * -2.25)
        _ = v10
    }
    v11 := GenS1{F0: (("hello" + (("hello" + "a") + "\n\t")) + (("foo bar" + "hello") + "foo bar")), F1: 127, F2: genFn0(18446744073709551615)}
    p0 = (^(p0))
    for fi12 := 0; fi12 < 3; fi12++ {
        if len(v9) > 0 {
            v9[0] = (fi12 ^ (fi12 * (fi12 - fi12)))
        }
        p0 = p0
        var v13 uint = 268
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v11
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 0
    var v3 int64 = 42
    var v4 int = -516
    v14 := GenS2{F0: ((^(v2)) != v2), F1: 127}
    v15 := GenS2{F0: (false && (!false)), F1: -5}
    for fi16 := 0; fi16 < 7; fi16++ {
        v2 = v2
        _ = fi16
    }
    var v17 uint = 2
    var v18 uint32 = 4294967295
    v19 := GenS3{F0: 32767, F1: 1, F2: 2}
    for fi20 := 0; fi20 < 15; fi20++ {
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 int8 = -77
            _ = fi21
            _ = v22
        }
        v23 := v1[v18]
        _ = fi20
        _ = v23
    }
    v4 -= (^(v4))
    v17 = v17
    v24 := []int{}
    v25 := GenS2{F0: false, F1: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
