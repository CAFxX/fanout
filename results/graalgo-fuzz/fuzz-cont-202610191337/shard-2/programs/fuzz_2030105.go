package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint16 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]int)
    var v3 uint16 = 0
    fmt.Println(v1, v1)
    var v4 int32 = -2147483648
    v5 := GenS1{F0: 42, F1: 232}
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 uint16 = v3
        v8 := v2[546]
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 string = "hello"
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 float64 = 1.5
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        if true {
            var v14 int64 = 244
            v2[2147483648] = fi13
            fmt.Println(fi13, v1, v9)
            _ = v14
        } else {
            fmt.Println(fi13)
            var v15 int16 = 1
            _ = v15
        }
        fmt.Println(v9, v3, v3)
        v16 := GenS1{F0: 8, F1: 138}
        _ = fi13
        _ = v16
    }
    v17 := GenS1{F0: 127, F1: 2}
    var v18 float32 = v1
    v19 := GenS1{F0: -1, F1: 128}
    fmt.Println(v9, v4, v4)
    if true {
        v20 := GenS2{F0: 2147483648}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 7; fi21++ {
            v9 = ("hello" + ("a" + v9))
            delete(v2, 76)
            _ = fi21
        }
        v3 = 410
    }
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        fmt.Println(v9, v4, v3)
    case "":
        if false {
            v22 := []int{32767, -2147483648}
            _ = v22
        } else {
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 uint32 = 65535
                v2[(v24 >> ((1) & 63))] = fi23
                v25 := GenS1{F0: -128, F1: 127}
                _ = fi23
                _ = v24
                _ = v25
            }
        }
    }
    if ("foo bar" >= (v9 + ("a" + v9))) {
        if true {
            var v26 int32 = 0
            _ = v26
        } else {
            v27 := GenS1{F0: -128, F1: 228}
            v28 := GenS3{F0: -49, F1: (v3 * v3)}
            var v29 int8 = -1
            _ = v27
            _ = v28
            _ = v29
        }
        v18 = 0.1
        if true {
            for fi30 := 0; fi30 < 3; fi30++ {
                fmt.Println(v1)
                _ = fi30
            }
            v4 = 0
            var v31 int16 = 42
            _ = v31
        }
    } else {
        v32 := GenS1{F0: 1, F1: 1}
        v33 := []int{32767}
        _ = v32
        _ = v33
    }
    fmt.Println(v18)
    var v34 int64 = -336
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v9
}
