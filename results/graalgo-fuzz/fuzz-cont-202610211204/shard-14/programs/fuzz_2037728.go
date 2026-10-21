package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
func genFn0(p0 string, p1 uint16) uint64 {
    if true {
        v6 := []int{32767}
        _ = v6
    } else {
        switch ((p0 + "hello") + p0) {
        case "foo bar":
            var v7 uint8 = 128
            _ = v7
        case "hello":
            v8 := []int{32768}
            _ = v8
        case "":
            for fi9 := 0; fi9 < 10; fi9++ {
                fmt.Println(fi9, p1)
                switch fi9 {
                case -359:
                    for fi10 := 0; fi10 < 18; fi10++ {
                        var v11 uint16 = p1
                        var v12 int64 = 9223372036854775807
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
                _ = fi9
            }
        default:
            fmt.Println(p0)
        }
    }
    p0 = ""
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    return 668
}

func main() {
    var v1 string = ""
    var v2 string = v1
    v3 := make(map[uint32]int)
    var v4 uint32 = 658
    v5 := make(map[int]string)
    var v13 float32 = (((3.14159 + -2.25) * 0.0) / 1.0)
    v14 := GenS1{F0: 65535, F1: -1, F2: -1}
    v1 = (v1 + v1)
    v15 := []int{int(v4), 0, 2147483647}
    v13 += float32(1.5)
    switch 32767 {
    case 42:
        v16 := GenS1{F0: 0, F1: 1, F2: 0}
        _ = v16
    case 128:
        switch v1 {
        case "":
            var v17 int8 = 1
            _ = v17
        default:
            v2 = "\n\t"
        }
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS1{F0: 42, F1: 51, F2: -275}
        _ = fi18
        _ = v19
    }
    fmt.Println(v2, v4)
    v20 := GenS1{F0: 2, F1: 42, F2: -32768}
    v21 := GenS1{F0: 1, F1: 2147483647, F2: -32768}
    fmt.Println(v1, v13, v4)
    if ((2 > v4) || ((("foo bar" + ("foo bar" + v1)) + v1) != v1)) {
        var v22 int16 = -1
        var v23 uint = 255
        _ = v22
        _ = v23
    } else {
        v24 := v3[v4]
        _ = v24
    }
    v25 := v3[v4]
    switch (v2 + v2) {
    case "foo bar":
        _ = len(v3)
    case "hello":
        var v26 int32 = 1
        _ = v26
    }
    v13 -= v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
