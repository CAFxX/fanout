package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint; F1 uint; F2 bool }
func genFn0(p0 int32) uint8 {
    if (true && (!((1.5 != 1e10) || (!false)))) {
        v5 := []int{32768, -2147483648}
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: 944, F1: 1, F2: 0}
            _ = fi6
            _ = v7
        }
        var v8 float64 = ((((3.14159 + 0.0) + (1.5 + (1e10 * float64(3.14159)))) * float64((0.1 - -2.25))) + 3.14159)
        _ = v5
        _ = v8
    }
    var v9 float64 = -2.25
    var v10 string = ("" + ("" + (("a" + (("a" + "\n\t") + ("hello" + (("foo bar" + "hello") + "hello")))) + ("a" + ""))))
    v11 := []int{-628}
    var v12 int32 = (p0 * p0)
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 41
}
func genFn1() uint32 {
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := []int{fi13, fi13}
        var v15 int = fi13
        _ = fi13
        _ = v14
        _ = v15
    }
    switch 1 {
    case 0:
        var v16 int64 = 0
        _ = v16
    case 758:
        var v17 uint64 = 1
        _ = v17
    default:
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := []int{}
            _ = fi18
            _ = v19
        }
    }
    return 2147483647
}

func main() {
    v1 := []int{1, -2147483648}
    var v2 int = 1
    var v3 uint8 = 2
    var v4 uint64 = 0
    var v20 uint8 = (^(v3))
    v4 += (^(v4))
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 float32 = (0.1 - -2.25)
        v22 += (((0.1 * (float32(((((3.14159 * ((1e10 / float64((v22 / -2.25))) / (((1.5 * float64(v22)) - 0.0) - 1.5))) - 3.14159) * 0.0) - 0.0)) / 0.1)) / (v22 * v22)) * float32(0.1))
        if true {
            var v23 int64 = 1
            fmt.Println(v23, v22)
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    v4 += 0
    v2 -= 30
    if false {
        v24 := GenS2{F0: 302, F1: 42, F2: ((v20 & (v20 % ((127) | 1))) != (v3 % ((v20) | 1)))}
        var v25 uint32 = genFn1()
        _ = v24
        _ = v25
    } else {
        v26 := GenS1{F0: v4, F1: 127, F2: 219}
        _ = v26
    }
    v4 = (v4 >> ((425) & 63))
    fmt.Println(v3)
    switch (-(v2)) {
    case 0:
        var v27 bool = false
        _ = v27
    }
    v1 = append(v1, v2)
    if len(v1) > 0 {
        v1[0] = (v2 % ((-350) | 1))
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        if true {
            _ = len(v1)
            for fi29 := 0; fi29 < 18; fi29++ {
                v4 += v4
                _ = fi29
            }
        }
        if false {
            v20 -= v20
            v30 := GenS2{F0: 65535, F1: 744, F2: (v3 > v20)}
            _ = v30
        }
        var v31 string = "foo bar"
        _ = fi28
        _ = v31
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        fmt.Println(v3, v4, v2)
        _ = fi32
    }
    var v33 int32 = 582
    fmt.Println(v20, v3)
    v1 = append(v1, 709)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v4
}
