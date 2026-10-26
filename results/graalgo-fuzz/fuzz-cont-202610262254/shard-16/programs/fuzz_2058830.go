package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int64 }
func genFn0() uint32 {
    var v7 string = ("foo bar" + ("hello" + ("foo bar" + ("foo bar" + "hello"))))
    v8 := GenS1{F0: -1}
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 bool) string {
    switch 425 {
    case -13:
        p0 = p0
    default:
        switch "\n\t" {
        case "":
            for fi9 := 0; fi9 < 14; fi9++ {
                v10 := GenS1{F0: -986}
                v11 := GenS1{F0: 8}
                _ = fi9
                _ = v10
                _ = v11
            }
        case "foo bar":
            var v12 uint32 = 2147483648
            _ = v12
        case "a":
            var v13 int16 = 32767
            _ = v13
        default:
            switch "foo bar" {
            case "foo bar":
                var v14 float32 = 0.1
                _ = v14
            }
        }
    }
    fmt.Println(p0)
    v15 := GenS1{F0: 2147483647}
    var v16 string = "a"
    _ = p0
    _ = v15
    _ = v16
    return v16
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = ((1.5 * 1.5) / float32(3.14159))
    var v3 float64 = float64(v2)
    var v4 uint64 = 494
    var v5 uint32 = 356
    v6 := make(map[string]string)
    v17 := GenS2{F0: 525, F1: false}
    v3 -= 0.0
    var v18 uint16 = 960
    _ = len(v1)
    var v19 uint32 = 2147483648
    fmt.Println(v4, v19)
    for fi20 := 0; fi20 < 10; fi20++ {
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 uint16 = (^(v18))
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v23 := []int{int(v18), -9223372036854775808, 127}
    var v24 uint = 4294967295
    for fi25 := 0; fi25 < 10; fi25++ {
        fmt.Println(fi25, v4)
        v23 = append(v23, (fi25 % ((fi25) | 1)))
        if false {
            var v26 int8 = -108
            if (!false) {
                var v27 int8 = v26
                var v28 bool = (82 < 42)
                _ = v27
                _ = v28
            } else {
                var v29 bool = false
                var v30 uint32 = 1
                _ = v29
                _ = v30
            }
            var v31 int32 = 2147483647
            _ = v26
            _ = v31
        }
        _ = fi25
    }
    if (42 < 42) {
        if false {
            for fi32 := 0; fi32 < 8; fi32++ {
                switch "\n\t" {
                case "a":
                    v18 = (^(v18))
                }
                _ = fi32
            }
        } else {
            var v33 int32 = 0
            v34 := GenS3{F0: 394}
            switch ("" + "") {
            case "\n\t":
                v18 = v18
            }
            _ = v33
            _ = v34
        }
        var v35 int64 = -1
        _ = v35
    }
    var v36 int16 = -752
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
