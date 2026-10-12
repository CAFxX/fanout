package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int16 }
func genFn0(p0 int16, p1 uint) int {
    switch 128 {
    case 32767:
        fmt.Println(p1)
    case 880:
        v7 := GenS1{F0: true}
        _ = v7
    case 4294967295:
        v8 := GenS2{F0: -2147483648, F1: -1}
        _ = v8
    default:
        v9 := []int{int(p1), -655, 4294967295}
        _ = v9
    }
    v10 := []int{-747, 0, 2}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if true {
        p0 = 1
        var v11 int32 = -708
        _ = v11
    } else {
        p1 = p1
        switch 834 {
        case 32768:
            switch ((("hello" + "foo bar") + "") + ("a" + (("hello" + "") + "foo bar"))) {
            case "\n\t":
                var v12 string = ("\n\t" + ("" + ("hello" + ("\n\t" + "\n\t"))))
                _ = v12
            }
        default:
            fmt.Println(p0, p1)
        }
        if false {
            if true {
                var v13 uint16 = 65535
                var v14 bool = true
                if (("" + "") != (("\n\t" + "") + "foo bar")) {
                    v15 := GenS1{F0: (false || v14)}
                    fmt.Println(v14)
                    p1 = p1
                    _ = v15
                }
                _ = v13
                _ = v14
            }
            fmt.Println(p0, p1)
        } else {
            v16 := GenS1{F0: (!false)}
            v17 := GenS3{F0: 42}
            _ = v16
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = v10
    return 128
}
func genFn1() uint16 {
    var v18 int16 = 1
    v19 := []int{-730, genFn0(v18, 18446744073709551615), 542}
    v20 := GenS2{F0: -916, F1: 127}
    for fi21 := 0; fi21 < 3; fi21++ {
        if len(v19) > 0 {
            v19[0] = (fi21 % ((genFn0(v18, 0)) | 1))
        }
        fmt.Println(v18, v18)
        var v22 int = fi21
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: false}
    _ = v18
    _ = v19
    _ = v20
    _ = v23
    return 83
}

func main() {
    var v1 int64 = 42
    var v2 int32 = 42
    var v3 string = ("" + ("\n\t" + ("hello" + (("hello" + "foo bar") + "hello"))))
    v4 := []int{}
    var v5 uint8 = 1
    v6 := make(map[int]int)
    var v24 int32 = (v2 + v2)
    v24 = (v2 & v2)
    if true {
        v3 = v3
    }
    delete(v6, 32768)
    var v25 string = v3
    if len(v4) > 0 {
        v4[0] = 42
    }
    v26 := GenS2{F0: v24, F1: 4294967295}
    fmt.Println(v25, v25)
    v27 := GenS2{F0: v2, F1: 4294967295}
    var v28 float64 = ((3.14159 - 3.14159) - (0.0 * 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
