package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
func genFn0(p0 uint16, p1 int, p2 uint64) uint32 {
    if ((p2 == (^(p2))) || true) {
        v5 := []int{(p1 & (-(p1)))}
        _ = v5
    } else {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int = (p1 ^ 128)
            fmt.Println(fi6, fi6, p2)
            fmt.Println(p0, fi6)
            _ = fi6
            _ = v7
        }
    }
    var v8 int = (p1 ^ 32768)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 1
    var v3 string = ("foo bar" + "\n\t")
    v4 := []int{-149, 9223372036854775807, -2147483648}
    var v9 int8 = 42
    var v10 int64 = -419
    v11 := GenS1{F0: (-(v9)), F1: -598}
    var v12 string = ((((("a" + (v3 + (((v3 + "a") + "hello") + v3))) + v3) + ((v3 + v3) + v3)) + ((("hello" + v3) + v3) + v3)) + v3)
    v13 := GenS1{F0: v9, F1: 758}
    v12 = (v3 + v12)
    var v14 float32 = (-2.25 + 1.5)
    v10 -= (v10 * (v10 - v10))
    if false {
        var v15 uint16 = 65535
        if (false || (true && (912 >= 1))) {
            v16 := GenS1{F0: v9, F1: 2147483647}
            var v17 uint32 = 690
            _ = v16
            _ = v17
        }
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
