package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint16 }
func genFn0() int {
    var v7 float32 = 3.14159
    v8 := GenS1{F0: 2, F1: 2147483647, F2: 107}
    switch "hello" {
    case "":
        v9 := GenS1{F0: 1, F1: 2, F2: 0}
        _ = v9
    default:
        for fi10 := 0; fi10 < 9; fi10++ {
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 int8 = -100
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    _ = v7
    _ = v8
    return 436
}
func genFn1(p0 int8) int {
    var v13 int64 = 474
    switch ("a" + "foo bar") {
    case "\n\t":
        v14 := GenS1{F0: 255, F1: 65535, F2: 741}
        _ = v14
    case "a":
        var v15 int32 = 1
        _ = v15
    case "foo bar":
        fmt.Println(v13)
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        if false {
            v17 := []int{genFn0(), (fi16 % ((-9223372036854775808) | 1))}
            var v18 int64 = (-(v13))
            _ = v17
            _ = v18
        }
        _ = fi16
    }
    _ = p0
    _ = v13
    return 32767
}

func main() {
    var v1 uint64 = 647
    var v2 string = "a"
    v3 := make(map[uint32]uint32)
    var v4 int64 = 9223372036854775807
    var v5 int64 = v4
    var v6 uint8 = 187
    var v19 uint16 = 0
    if false {
        for fi20 := 0; fi20 < 18; fi20++ {
            fmt.Println(v19)
            var v21 uint = 0
            v6 = 255
            _ = fi20
            _ = v21
        }
        var v22 uint32 = 65535
        _ = v22
    }
    v1 = v1
    v23 := []int{-9223372036854775808, -153}
    switch 128 {
    case 128:
        switch int(v4) {
        case 32767:
            v6 += v6
        case 42:
            var v24 string = v2
            _ = v24
        case 128:
            if ("a" >= (v2 + v2)) {
                v2 = "foo bar"
                var v25 float64 = ((0.0 * 1.5) * 0.1)
                _ = v25
            } else {
                var v26 int32 = 0
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 uint64 = (v1 % (((^(v1))) | 1))
                    _ = fi27
                    _ = v28
                }
                _ = v26
            }
        }
    case 9223372036854775807:
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 int64 = v5
            _ = fi29
            _ = v30
        }
    default:
        v2 = "\n\t"
    }
    delete(v3, 2)
    fmt.Println(v4, v1)
    if (!(true && false)) {
        v31 := GenS1{F0: 0, F1: uint32(v1), F2: v19}
        _ = v31
    }
    var v32 int16 = 42
    v5 -= (v5 ^ (v4 << ((122) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
