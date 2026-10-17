package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint32; F2 int }
func genFn0(p0 int) uint8 {
    p0 -= p0
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 0
}
func genFn1(p0 uint32, p1 uint32, p2 int64) int8 {
    var v4 int8 = -128
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (v4 - v4)
}

func main() {
    var v1 uint = 0
    v2 := []int{9223372036854775807, -9223372036854775808, -9223372036854775808}
    v3 := make(map[int]int)
    v1 -= 713
    v3[-57] = 2147483647
    var v5 bool = true
    v6 := GenS1{F0: 833, F1: 70}
    v5 = v5
    var v7 int8 = 127
    if v5 {
        v7 = genFn1(2147483647, 1, 1)
        v1 -= v1
        v7 += v7
    }
    fmt.Println(v5, v5, v1)
    v5 = (42 <= -2147483648)
    v8 := GenS1{F0: 65535, F1: 65535}
    v9 := v3[1]
    var v10 uint64 = 420
    var v11 int = v9
    v12 := GenS2{F0: ("a" + "\n\t"), F1: 2, F2: int(v1)}
    if v5 {
        v13 := GenS2{F0: (("\n\t" + ("a" + ("\n\t" + ("\n\t" + "a")))) + "\n\t"), F1: 2147483648, F2: v9}
        if (!v5) {
            v9 -= -624
            var v14 uint16 = 217
            _ = v14
        } else {
            var v15 uint64 = 612
            var v16 int = (v9 ^ -1)
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 float32 = 3.14159
            var v19 int32 = -656
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v13
    } else {
        v10 = (v10 - v10)
    }
    var v20 int = v9
    if (18446744073709551615 < 241) {
        v11 += -2
        var v21 uint8 = 128
        _ = v21
    } else {
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 int8 = int8(v1)
            _ = fi22
            _ = v23
        }
        delete(v3, -230)
        if len(v2) > 0 {
            v2[0] = (^(v20))
        }
    }
    if (!v5) {
        fmt.Println(v9, v5, v5)
        if (14 > genFn1(1, 0, -882)) {
            fmt.Println(v10, v20, v7)
        }
        v24 := GenS1{F0: 469, F1: 2}
        _ = v24
    } else {
        if (v1 <= v1) {
            if true {
                var v25 string = ("foo bar" + "foo bar")
                v26 := GenS1{F0: 1, F1: 2147483648}
                _ = v25
                _ = v26
            }
        }
        fmt.Println(v20)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
