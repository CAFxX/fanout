package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 string, p1 bool, p2 bool) bool {
    if (p2 || false) {
        p2 = (0.1 == 3.14159)
    } else {
        var v7 float64 = -2.25
        p0 = (p0 + p0)
        _ = v7
    }
    var v8 float32 = 3.14159
    if (411 != 488) {
        var v9 int64 = 42
        _ = v9
    }
    if (2 == 255) {
        v10 := []int{187, 32768, -83}
        v11 := GenS1{F0: 0}
        _ = v10
        _ = v11
    }
    if p2 {
        p2 = (0 > -62)
        v12 := GenS1{F0: 132}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return true
}
func genFn1(p0 uint, p1 uint16, p2 string) int {
    var v13 int64 = -1
    p1 += p1
    var v14 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 128
}

func main() {
    v1 := make(map[int]int)
    var v2 string = ((("foo bar" + "foo bar") + "a") + "foo bar")
    var v3 int64 = -863
    var v4 int8 = -39
    var v5 uint32 = 1
    var v6 int16 = -1
    var v15 int64 = (^(v3))
    var v16 uint32 = (v5 & v5)
    v4 -= (^(v4))
    v17 := GenS2{F0: 42}
    var v18 float64 = 0.0
    var v19 bool = false
    switch ((("foo bar" + (((v2 + (("hello" + v2) + v2)) + ("" + "\n\t")) + v2)) + "hello") + v2) {
    case "":
        var v20 string = "foo bar"
        _ = v20
    case "hello":
        var v21 int8 = 69
        _ = v21
    case "a":
        var v22 uint16 = 510
        _ = v22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            if (2 >= (v5 + (v16 - 716))) {
                delete(v1, fi23)
                _ = len(v1)
            } else {
                var v25 string = "hello"
                _ = v25
            }
            v26 := []int{}
            fmt.Println(fi23, fi23, v18)
            _ = len(v1)
            _ = fi24
            _ = v26
        }
        v27 := []int{238, -1, (^(fi23))}
        v16 -= (v16 / ((0) | 1))
        _ = fi23
        _ = v27
    }
    v4 -= v4
    var v28 int32 = 42
    v15 -= v15
    if (true || false) {
        _ = len(v1)
        v29 := v1[-218]
        _ = v29
    }
    v30 := []int{-1, -285}
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 float64 = (float64(1e10) * 1.5)
        var v33 uint64 = 0
        v1[(^(fi31))] = (fi31 | int(v6))
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
