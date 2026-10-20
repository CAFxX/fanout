package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int8; F2 string }
type GenS4 struct { F0 uint16 }
func genFn0() int32 {
    var v7 uint32 = 33
    switch 32767 {
    case -85:
        v7 = (v7 & v7)
    default:
        switch 42 {
        case 128:
            if false {
                var v8 float32 = (float32(0.1) - 0.0)
                _ = v8
            } else {
                var v9 uint64 = 2
                var v10 int16 = 0
                _ = v10
                _ = v9
            }
        default:
            v7 = (v7 + (v7 * (v7 | (v7 % ((v7) | 1)))))
        }
    }
    v11 := []int{128, 756, 458, -2147483648}
    var v12 uint32 = (v7 >> ((0) & 63))
    _ = v11
    _ = v12
    _ = v7
    return 1
}

func main() {
    v1 := []string{("" + ((("a" + ((("foo bar" + "hello") + "a") + "hello")) + ("a" + "hello")) + "foo bar")), (("" + ("\n\t" + "\n\t")) + (("hello" + "a") + "foo bar"))}
    var v2 uint = 4294967295
    v3 := []int{249, 127}
    v4 := []uint32{}
    var v5 uint32 = 4294967295
    var v6 uint = (v2 ^ uint(v5))
    v13 := GenS1{F0: 2, F1: 0, F2: v6}
    v14 := GenS2{F0: (v5 + v5)}
    v15 := GenS1{F0: 18446744073709551615, F1: 1000, F2: 2}
    switch "foo bar" {
    case "a":
        var v16 int32 = genFn0()
        _ = v16
    }
    var v17 float32 = -2.25
    v2 -= v2
    fmt.Println(v2, v17)
    if (!(!(!true))) {
        fmt.Println(v6)
        var v18 int16 = -816
        fmt.Println(v2, v6)
        _ = v18
    }
    v19 := GenS1{F0: 42, F1: 0, F2: (v2 % ((1) | 1))}
    for fi20 := 0; fi20 < 13; fi20++ {
        fmt.Println(fi20)
        fmt.Println(v2, v6, v17)
        _ = fi20
    }
    v21 := GenS1{F0: 4294967295, F1: 42, F2: (^(v2))}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        if ((-20 <= genFn0()) || false) {
            var v22 float64 = ((float64(float32((-2.25 * 0.1))) * 0.0) * ((1e10 / 1e10) + 1.5))
            v17 = v17
            _ = v22
        }
        v23 := GenS3{F0: 32768, F1: 42, F2: "hello"}
        _ = v23
    }
    v24 := GenS2{F0: v5}
    var v25 float64 = (3.14159 / 0.1)
    _ = len(v4)
    v26 := GenS4{F0: 432}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v19
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
