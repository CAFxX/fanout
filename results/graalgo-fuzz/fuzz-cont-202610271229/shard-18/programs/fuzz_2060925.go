package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint64; F1 int32; F2 int16 }
type GenS3 struct { F0 int16; F1 int16 }
func genFn0(p0 string, p1 uint) int {
    if false {
        p1 -= p1
        switch (p0 + "hello") {
        case "\n\t":
            v5 := []int{-1}
            _ = v5
        case "":
            fmt.Println(p0)
        case "hello":
            for fi6 := 0; fi6 < 19; fi6++ {
                v7 := GenS1{F0: true, F1: 4294967295, F2: 1000}
                v8 := []int{-2147483648, fi6}
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 int64 = 42
                    _ = fi9
                    _ = v10
                }
                _ = fi6
                _ = v7
                _ = v8
            }
        default:
            var v11 int64 = 0
            _ = v11
        }
        var v12 int32 = 2147483647
        _ = v12
    }
    var v13 bool = false
    fmt.Println(p1, v13)
    v14 := GenS1{F0: (v13 && (!(1000 <= 1))), F1: 42, F2: -32768}
    var v15 int8 = -1
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return 1
}
func genFn1() bool {
    v16 := GenS1{F0: (-2.25 >= -2.25), F1: 1, F2: 707}
    var v17 string = "hello"
    _ = v16
    _ = v17
    return true
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    v3 := []int{}
    v4 := []uint32{2147483648, 65535, 30, 359}
    switch genFn0("foo bar", 4294967295) {
    case 2147483647:
        v4 = append(v4, 260)
    case 128:
        v18 := GenS2{F0: 1, F1: 42, F2: 42}
        _ = v18
    }
    if true {
        switch genFn0("\n\t", 65535) {
        case 2147483647:
            var v19 float32 = 0.1
            _ = v19
        case -2147483648:
            v20 := GenS2{F0: 4294967295, F1: 0, F2: 1000}
            _ = v20
        case 949:
            var v21 float32 = 3.14159
            _ = v21
        }
    }
    switch "foo bar" {
    case "a":
        var v22 int = -2147483648
        _ = v22
    case "\n\t":
        if (!false) {
            if false {
                v23 := GenS2{F0: 454, F1: -177, F2: -32768}
                var v24 uint32 = 264
                var v25 uint16 = 377
                _ = v23
                _ = v24
                _ = v25
            } else {
                var v26 int8 = 42
                v27 := GenS2{F0: 4294967295, F1: -500, F2: 696}
                _ = v26
                _ = v27
            }
            var v28 uint32 = 2147483647
            v29 := GenS3{F0: -162, F1: -1}
            _ = v28
            _ = v29
        }
    case "hello":
        if true {
            var v30 float64 = 0.1
            fmt.Println(v30)
            fmt.Println(v30)
            _ = v30
        }
    }
    var v31 int = 730
    for fi32 := 0; fi32 < 2; fi32++ {
        switch -26 {
        case 42:
            var v33 int = genFn0((("foo bar" + "") + "hello"), 654)
            _ = v33
        case 9223372036854775807:
            _ = len(v3)
        default:
            v31 = genFn0(("foo bar" + "\n\t"), 4294967295)
        }
        _ = fi32
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        var v35 uint = 42
        _ = fi34
        _ = v35
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v31 = 186
    fmt.Println(v31)
    var v36 uint = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v4
}
