package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }
type GenS3 struct { F0 uint8; F1 int }
func genFn0() int16 {
    var v6 string = "hello"
    v7 := []int{}
    fmt.Println(v6)
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int32 = 1
        v10 := GenS1{F0: 64}
        var v11 uint = 2
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v6 = v6
    _ = v6
    _ = v7
    return -32768
}
func genFn1(p0 uint16, p1 int) uint32 {
    fmt.Println(p0, p1)
    p0 = (p0 * 255)
    var v12 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v12
    return 881
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 128
    var v3 int32 = 256
    var v4 int16 = -32768
    var v5 uint64 = 988
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int32 = (v3 + int32(fi13))
        var v15 int = 1
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 float32 = 1e10
    v2 -= v2
    var v17 float64 = 1e10
    if ((!(("a" <= ("\n\t" + "foo bar")) || true)) && false) {
        v18 := []int{v2, 1, 869}
        v19 := GenS1{F0: v5}
        fmt.Println(v5, v16)
        _ = v18
        _ = v19
    }
    switch v2 {
    case -9223372036854775808:
        v3 += 42
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        if false {
            v21 := GenS1{F0: 42}
            _ = v21
        }
        v22 := []int{}
        _ = fi20
        _ = v22
    }
    if (!true) {
        fmt.Println(v2, v2)
        v23 := GenS2{F0: 2147483648, F1: 2, F2: ""}
        fmt.Println(v3)
        _ = v23
    } else {
        v24 := GenS3{F0: 0, F1: (v2 / ((v2) | 1))}
        _ = v24
    }
    if false {
        v25 := []int{(v2 + (v2 & v2)), -2147483648, v2, 0}
        if len(v25) > 0 {
            v25[0] = (v2 >> ((2) & 63))
        }
        _ = v25
    } else {
        fmt.Println(v3)
        v26 := []int{}
        _ = v26
    }
    var v27 uint = 0
    v28 := []int{(v2 & (^(v2))), (v2 % ((-274) | 1)), (^(v2)), int(v3)}
    if true {
        v27 = v27
    }
    v29 := GenS2{F0: 2147483648, F1: 2, F2: ("a" + (("" + "\n\t") + "a"))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
