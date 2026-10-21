package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() uint16 {
    var v7 int16 = 0
    v8 := GenS1{F0: 466}
    var v9 string = ("" + ("" + (("" + "a") + "")))
    var v10 float64 = 1e10
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint16, p1 uint32) int {
    var v11 float64 = 1e10
    var v12 int = -807
    switch (v12 & (v12 ^ v12)) {
    case -963:
        var v13 int8 = 1
        _ = v13
    case 194:
        v14 := []int{}
        _ = v14
    case 43:
        switch (v12 / ((v12) | 1)) {
        case 2147483647:
            var v15 string = "hello"
            _ = v15
        case -2147483648:
            var v16 int64 = -1
            _ = v16
        }
    }
    v17 := []int{(-(v12)), -609}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v17
    return -1
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint8 = 202
    var v3 uint64 = uint64(v2)
    v4 := []int{1}
    v5 := []uint32{99, 2147483647, 2147483648}
    var v6 int = 42
    fmt.Println(v6)
    var v18 int64 = -9223372036854775808
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v19 string = ""
    var v20 int64 = 42
    var v21 int16 = -32768
    if false {
        if len(v5) > 0 {
            v5[0] = 1
        }
    }
    var v22 uint16 = 0
    v21 = v21
    v20 = v20
    if len(v5) > 0 {
        v5[0] = 1
    }
    var v23 int32 = -1
    if len(v5) > 0 {
        v5[0] = 908
    }
    var v24 uint16 = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
