package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 int16) int16 {
    v6 := GenS1{F0: (p0 / ((p0) | 1))}
    p0 = p0
    if false {
        v7 := GenS1{F0: 0}
        p0 = p0
        _ = v7
    } else {
        p0 -= p0
        var v8 int16 = p0
        for fi9 := 0; fi9 < 2; fi9++ {
            v8 -= (p0 - p0)
            _ = fi9
        }
        _ = v8
    }
    v10 := []int{0, 32767, 400, 32767}
    fmt.Println(p0)
    _ = p0
    _ = v10
    _ = v6
    return p0
}
func genFn1() int8 {
    var v11 int16 = -1
    v12 := []int{4294967295}
    var v13 uint16 = 1
    _ = v11
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 int8 = 42
    var v2 uint16 = 255
    v3 := []int{34, -1, 2147483647}
    v4 := []string{}
    v5 := make(map[int]string)
    if ("" != "") {
        _ = len(v3)
    }
    switch "" {
    case "":
        v1 = v1
    case "foo bar":
        v14 := GenS2{F0: 1, F1: 420}
        _ = v14
    }
    v2 -= 68
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS1{F0: 32767}
        v17 := GenS1{F0: 32767}
        fmt.Println(v2)
        _ = fi15
        _ = v16
        _ = v17
    }
    v2 = (^(v2))
    fmt.Println(v2, v2)
    var v18 uint16 = (v2 % (((v2 ^ 581)) | 1))
    var v19 bool = (v18 > (v2 / ((65535) | 1)))
    var v20 uint32 = 65535
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 uint = 546
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        fmt.Println(v19)
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 int = fi24
            var v26 int16 = 931
            if true {
                fmt.Println(v25, v20)
                var v27 uint = 42
                _ = v27
            }
            v2 = 0
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi23
    }
    _ = len(v5)
    var v28 int32 = 335
    v18 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
