package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int32; F1 int; F2 int8 }
func genFn0(p0 int, p1 int64, p2 int32) uint16 {
    v5 := GenS1{F0: 2, F1: 1}
    v6 := GenS1{F0: 42, F1: 204}
    var v7 int64 = (^(p1))
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    switch p0 {
    case 32767:
        v10 := []int{32767, (p0 - (p0 % ((p0) | 1)))}
        _ = v10
    case -1:
        v7 -= -1
    case 1:
        var v11 uint8 = 145
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 546
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 3.14159
    v3 := make(map[string]string)
    v4 := []int{-2147483648, 127, 32768}
    if ((401 < 32767) || true) {
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            var v12 uint64 = 963
            _ = v12
        case "":
            if (956 != -150) {
                var v13 string = (((("foo bar" + "a") + ("foo bar" + ("" + "foo bar"))) + ("" + (((("foo bar" + "") + "a") + ("a" + (("a" + ("hello" + "foo bar")) + "\n\t"))) + ""))) + "a")
                _ = v13
            } else {
                v2 -= float64((3.14159 * (1.5 + 0.1)))
            }
        }
        v14 := GenS1{F0: 127, F1: 0}
        _ = v14
    } else {
        fmt.Println(v2)
        v2 -= v2
    }
    fmt.Println(v2)
    v15 := v1[127]
    v16 := GenS1{F0: 176, F1: 255}
    var v17 int16 = 32767
    if false {
        var v18 uint8 = 255
        var v19 uint32 = 2
        var v20 uint32 = v19
        _ = v18
        _ = v19
        _ = v20
    }
    if len(v4) > 0 {
        v4[0] = v15
    }
    switch 32767 {
    case -294:
        v21 := GenS2{F0: 1, F1: (-(v15)), F2: 0}
        _ = v21
    }
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
