package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int64; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint }
func genFn0() int16 {
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint16 = 32767
        _ = fi7
        _ = v8
    }
    if false {
        var v9 string = "a"
        var v10 int64 = 1
        var v11 uint = 0
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 bool = false
    _ = v12
    return -1
}
func genFn1(p0 int16) uint64 {
    fmt.Println(p0)
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS1{F0: (fi13 & fi13)}
        _ = fi13
        _ = v14
    }
    p0 = genFn0()
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    var v1 int64 = 0
    var v2 uint = 255
    var v3 uint8 = 128
    var v4 uint64 = 4294967295
    var v5 uint = v2
    var v6 string = "\n\t"
    fmt.Println(v6, v1)
    v15 := []int{-9223372036854775808}
    var v16 string = "\n\t"
    if ((!(-266 < -1)) || (!(!(((!false) || (v1 <= (v1 * -385))) && (true && (!(v6 >= "foo bar"))))))) {
        if false {
            v4 -= v4
        }
    } else {
        var v17 float32 = (-2.25 * (float32((1e10 / 1e10)) / -2.25))
        for fi18 := 0; fi18 < 7; fi18++ {
            v19 := GenS1{F0: fi18}
            v6 = ("hello" + "foo bar")
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 16; fi20++ {
            for fi21 := 0; fi21 < 20; fi21++ {
                fmt.Println(v2)
                v22 := GenS1{F0: fi20}
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
        _ = v17
    }
    switch (("a" + v6) + (v6 + "")) {
    case "":
        v5 = (v5 ^ v2)
    case "\n\t":
        _ = len(v15)
    default:
        var v23 float32 = (float32(0.1) - 0.0)
        _ = v23
    }
    if ((0.1 * 1e10) == (1e10 / 1e10)) {
        v15 = append(v15, 2147483647)
        v24 := GenS2{F0: v2, F1: int64(v2), F2: (v4 << ((4294967295) & 63))}
        _ = v24
    }
    if true {
        var v25 int64 = (^(v1))
        v26 := GenS3{F0: true, F1: uint(v3)}
        var v27 bool = false
        _ = v25
        _ = v26
        _ = v27
    }
    var v28 uint = 18446744073709551615
    _ = len(v15)
    var v29 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
