package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint }
type GenS2 struct { F0 int; F1 uint }
func genFn0(p0 uint32) uint {
    fmt.Println(p0)
    v4 := GenS1{F0: true, F1: -31, F2: 65535}
    _ = p0
    _ = v4
    return 65535
}
func genFn1() uint8 {
    if false {
        var v5 uint16 = 0
        v6 := []int{32768}
        var v7 uint64 = 965
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 uint64 = 2
    _ = v8
    return 128
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = -1
    var v3 uint16 = 1
    fmt.Println(v2, v3)
    var v9 bool = (!(!true))
    if (901 > genFn0(2147483647)) {
        v10 := []int{(v2 & (v2 / ((42) | 1))), -2147483648}
        v9 = true
        var v11 string = "foo bar"
        _ = v10
        _ = v11
    }
    fmt.Println(v9, v3)
    v12 := []int{v2, (v2 << ((287) & 63)), v2}
    if (977 <= 2) {
        if len(v12) > 0 {
            v12[0] = 738
        }
    } else {
        v13 := GenS2{F0: (v2 << ((65535) & 63)), F1: genFn0(4294967295)}
        delete(v1, 2)
        _ = v13
    }
    v9 = false
    v2 += v2
    var v14 float64 = float64(((1.5 / 1e10) * (-2.25 / 1.0)))
    var v15 string = ((((("foo bar" + "hello") + "hello") + "\n\t") + ("\n\t" + "\n\t")) + ("" + "hello"))
    fmt.Println(v14, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
