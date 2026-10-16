package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint) uint16 {
    if false {
        v6 := []int{2147483647, 1}
        if false {
            var v7 uint8 = 14
            v8 := GenS1{F0: -128}
            _ = len(v6)
            _ = v7
            _ = v8
        } else {
            p0 -= (p0 | p0)
            var v9 int64 = 1
            fmt.Println(v9, v9)
            _ = v9
        }
        v10 := GenS1{F0: -128}
        _ = v10
        _ = v6
    }
    if (2147483647 >= 4294967295) {
        if (213 == 0) {
            p0 = 426
        }
        v11 := []int{4294967295, 127, int(p0), 128}
        _ = len(v11)
        _ = v11
    }
    _ = p0
    return 443
}
func genFn1(p0 uint64, p1 uint8) int64 {
    v12 := []int{1}
    var v13 bool = ((-2147483648 >= 175) && true)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return -685
}

func main() {
    v1 := []int{-314, 2147483647, -2147483648}
    var v2 int = 32767
    var v3 bool = false
    var v4 string = ""
    var v5 float64 = float64(3.14159)
    v14 := GenS1{F0: 63}
    fmt.Println(v4, v5)
    v5 = v5
    var v15 float64 = 3.14159
    var v16 int32 = -262
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 string = (((v4 + "") + v4) + v4)
        var v19 uint64 = 18446744073709551615
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int16 = 1000
    fmt.Println(v20, v16)
    v20 += (v20 >> ((65535) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v21 int16 = (-(v20))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
