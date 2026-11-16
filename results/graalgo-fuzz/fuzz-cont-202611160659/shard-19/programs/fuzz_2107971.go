package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int8 }
type GenS2 struct { F0 int; F1 uint64; F2 uint64 }
func genFn0() int8 {
    var v5 uint8 = 128
    for fi6 := 0; fi6 < 3; fi6++ {
        if false {
            v5 += v5
            if false {
                if (642 <= 255) {
                    var v7 uint = 325
                    v7 = (^(v7))
                    _ = v7
                }
                v8 := GenS1{F0: 0, F1: 2, F2: 0}
                _ = v8
            }
        }
        v5 = 14
        _ = fi6
    }
    var v9 uint8 = (^(v5))
    v10 := []int{}
    v11 := GenS1{F0: 127, F1: 42, F2: -1}
    _ = v10
    _ = v11
    _ = v5
    _ = v9
    return -57
}
func genFn1(p0 uint8, p1 int64) uint {
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            p1 = 655
            fmt.Println(p1, p0)
            var v14 uint = 42
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    fmt.Println(p1, p0)
    var v15 float64 = -2.25
    v15 += v15
    var v16 uint16 = 255
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 42
}

func main() {
    v1 := []string{}
    var v2 uint64 = 4294967295
    var v3 uint16 = uint16(v2)
    var v4 int = -1
    switch (v4 / ((0) | 1)) {
    case 537:
        fmt.Println(v4, v3)
    }
    v17 := GenS1{F0: genFn0(), F1: v2, F2: 57}
    switch (v4 / ((v4) | 1)) {
    case 128:
        for fi18 := 0; fi18 < 14; fi18++ {
            if true {
                v19 := GenS2{F0: 1, F1: v2, F2: (v2 - (v2 * (v2 * (^(v2)))))}
                _ = v19
            }
            _ = fi18
        }
    }
    switch "foo bar" {
    case "hello":
        if (int8(v4) == 1) {
            var v20 uint = genFn1(255, 0)
            fmt.Println(v4, v2, v20)
            if (false || ((true && true) && (false || ((!((^(v3)) > (^(v3)))) || true)))) {
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 float64 = ((float64(3.14159) / 1.0) - (0.1 - float64(((-2.25 + -2.25) + 1e10))))
                    _ = fi21
                    _ = v22
                }
                var v23 uint8 = 176
                _ = v23
            }
            _ = v20
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v24 uint32 = 922
    fmt.Println(v3, v4)
    v4 += (v4 % ((-673) | 1))
    v24 = v24
    v4 -= v4
    var v25 int32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
