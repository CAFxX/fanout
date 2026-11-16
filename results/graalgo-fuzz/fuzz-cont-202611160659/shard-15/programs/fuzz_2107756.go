package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
func genFn0(p0 string) int64 {
    p0 = "hello"
    var v6 float64 = 0.0
    v7 := []int{}
    var v8 int = 128
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 string, p1 int) int16 {
    v9 := []int{}
    if ((p1 >> ((1) & 63)) == 0) {
        if false {
            var v10 uint16 = 42
            p1 -= (p1 ^ 128)
            _ = v10
        }
    } else {
        p0 = "a"
        var v11 float32 = 1e10
        _ = v11
    }
    var v12 float32 = 1.5
    var v13 string = "\n\t"
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v9
    return 32767
}

func main() {
    v1 := []uint32{2147483647, 42, 42}
    var v2 string = ("hello" + "")
    var v3 int = 32767
    var v4 int32 = 1
    v5 := []string{("foo bar" + v2), ((v2 + v2) + v2), v2}
    v2 = v2
    if len(v1) > 0 {
        v1[0] = 65535
    }
    if len(v1) > 0 {
        v1[0] = 65535
    }
    var v14 int8 = 42
    if (128 >= 156) {
        var v15 int64 = -24
        fmt.Println(v2)
        v15 = (v15 >> ((65535) & 63))
        _ = v15
    } else {
        var v16 uint64 = 0
        _ = v16
    }
    var v17 bool = (!false)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v18 uint16 = 255
    var v19 float32 = ((1e10 + -2.25) * 1e10)
    v3 += -2147483648
    var v20 uint64 = 0
    var v21 string = v2
    v22 := GenS1{F0: -309, F1: v20}
    var v23 bool = (!(0 > 0))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
