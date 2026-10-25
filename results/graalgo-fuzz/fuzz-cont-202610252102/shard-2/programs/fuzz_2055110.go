package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 bool }
type GenS2 struct { F0 uint16 }
func genFn0() string {
    v7 := GenS1{F0: (((("hello" + ("\n\t" + "hello")) + "a") + "a") + (("" + ("hello" + "")) + "foo bar")), F1: 0, F2: (!false)}
    var v8 int = 628
    if ((v8 << ((244) & 63)) < (-(v8))) {
        fmt.Println(v8)
    } else {
        v9 := []int{(v8 >> ((255) & 63)), 486}
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS2{F0: 0}
        v8 = (-(fi10))
        _ = fi10
        _ = v11
    }
    _ = v7
    _ = v8
    return ("" + "foo bar")
}
func genFn1(p0 uint16, p1 int8, p2 uint64) uint16 {
    fmt.Println(p2)
    var v12 int = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 2
}

func main() {
    var v1 uint16 = 283
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    v4 := make(map[uint32]int)
    var v5 string = "foo bar"
    var v6 float32 = (((1.5 - (0.1 + 0.1)) - 3.14159) * 1e10)
    var v13 float32 = v6
    v1 -= genFn1((v1 + genFn1(998, 51, 2)), 42, 0)
    v13 -= v6
    fmt.Println(v1, v13, v6)
    var v14 uint = 18446744073709551615
    if true {
        v15 := []int{}
        var v16 uint64 = 60
        _ = v15
        _ = v16
    }
    var v17 uint8 = 127
    v1 = v1
    v18 := []int{-499, 128}
    v13 = 1.5
    var v19 uint16 = (v1 & (v1 * 255))
    fmt.Println(v13, v13)
    for fi20 := 0; fi20 < 19; fi20++ {
        if ((32767 >= v1) || true) {
            v1 = v19
        } else {
            if (!(!(genFn1(32767, 93, 42) <= 384))) {
                v5 = ("a" + v5)
                v14 -= (v14 - 822)
            }
            switch v5 {
            case "":
                var v21 uint = v14
                _ = v21
            case "a":
                v13 = 3.14159
            default:
                fmt.Println(v14, fi20, v5)
            }
        }
        _ = fi20
    }
    var v22 uint64 = 2
    var v23 uint32 = 2
    v24 := GenS1{F0: v5, F1: -123, F2: (!(!(v6 != float32((float64(float32(1.5)) - (0.0 - -2.25))))))}
    if false {
        fmt.Println(v17)
        _ = len(v18)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
