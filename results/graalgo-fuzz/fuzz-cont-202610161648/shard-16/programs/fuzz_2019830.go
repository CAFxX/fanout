package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint16 }
func genFn0(p0 int8, p1 string) int8 {
    if false {
        p1 = ((p1 + (("\n\t" + ("hello" + p1)) + "")) + "\n\t")
    }
    p1 = (("" + p1) + p1)
    p1 = "\n\t"
    p1 = ((p1 + (p1 + p1)) + "\n\t")
    p0 = (p0 & (p0 ^ (p0 & p0)))
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 string = ("foo bar" + "hello")
    v2 := make(map[int]string)
    var v3 int8 = -128
    var v4 float32 = 1e10
    var v5 string = ((v1 + v1) + v1)
    v6 := GenS1{F0: 9223372036854775807, F1: (v1 + v5), F2: 42}
    v4 = v4
    var v7 float32 = 0.1
    var v8 int8 = v3
    if true {
        fmt.Println(v5, v5)
        v8 = (v8 + (v8 ^ (v8 - v3)))
    } else {
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 float64 = float64(float32(((1.5 + 0.1) - float64((3.14159 * v7)))))
            v11 := []int{(fi9 >> ((4294967295) & 63)), fi9}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS1{F0: 1, F1: v1, F2: 32767}
        v14 := []int{(fi12 % (((-(fi12))) | 1)), 127, fi12, (fi12 << ((912) & 63))}
        v4 += v4
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint8 = 9
    v3 = (v8 & v3)
    v16 := GenS1{F0: 9223372036854775807, F1: v5, F2: 1}
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 uint = 42
        var v19 uint32 = uint32(v15)
        _ = fi17
        _ = v18
        _ = v19
    }
    v4 = 3.14159
    v7 = float32(-2.25)
    var v20 int32 = 42
    v1 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
