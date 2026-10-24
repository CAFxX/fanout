package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int8; F1 int64; F2 int16 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int8; F1 int16 }
func genFn0(p0 uint8, p1 bool) uint8 {
    if false {
        var v4 int64 = 42
        var v5 bool = p1
        _ = v4
        _ = v5
    } else {
        v6 := []int{4294967295, -2147483648, -943, 0}
        _ = v6
    }
    p1 = p1
    v7 := GenS1{F0: -11, F1: 4294967295}
    fmt.Println(p1, p0)
    var v8 int8 = 127
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = true
    var v3 uint8 = 112
    for fi9 := 0; fi9 < 4; fi9++ {
        delete(v1, fi9)
        for fi10 := 0; fi10 < 7; fi10++ {
            switch 2147483647 {
            case -780:
                _ = len(v1)
            case 128:
                var v11 int64 = -9223372036854775808
                _ = v11
            case -958:
                v12 := v1[384]
                _ = v12
            default:
                if false {
                    v3 = (v3 % ((genFn0(v3, (0 > 42))) | 1))
                } else {
                    v13 := v1[(fi9 << ((65535) & 63))]
                    _ = v13
                }
            }
            var v14 int = fi9
            _ = fi10
            _ = v14
        }
        v3 = genFn0(genFn0(v3, (("hello" + (("hello" + ("hello" + "foo bar")) + "foo bar")) > "\n\t")), v2)
        _ = fi9
    }
    v15 := GenS1{F0: 1, F1: 18446744073709551615}
    if false {
        v16 := GenS1{F0: 42, F1: 311}
        _ = v16
    } else {
        v1[-9223372036854775808] = 65535
    }
    var v17 int8 = -128
    switch 1 {
    case 1:
        switch "\n\t" {
        case "hello":
            var v18 int8 = 0
            _ = v18
        case "":
            var v19 uint64 = 42
            _ = v19
        case "\n\t":
            v20 := GenS1{F0: -32768, F1: 0}
            _ = v20
        default:
            v21 := GenS2{F0: 42, F1: 35, F2: 0}
            _ = v21
        }
    case 128:
        fmt.Println(v2, v2, v2)
    default:
        if true {
            v3 = (v3 * v3)
            delete(v1, 1)
        }
    }
    switch ("a" + "a") {
    case "\n\t":
        fmt.Println(v3, v3)
    case "a":
        v17 = v17
    default:
        if (4294967295 != 0) {
            if v2 {
                var v22 uint8 = (v3 | (^(v3)))
                var v23 bool = ((v2 || v2) && ((0.0 / 1.5) <= 3.14159))
                _ = v22
                _ = v23
            }
        }
    }
    v24 := GenS2{F0: (v17 << ((984) & 63)), F1: 0, F2: 42}
    v2 = v2
    fmt.Println(v3)
    switch "foo bar" {
    case "\n\t":
        var v25 bool = (!(v2 && v2))
        _ = v25
    case "a":
        v26 := []int{}
        _ = v26
    }
    v27 := v1[408]
    for fi28 := 0; fi28 < 20; fi28++ {
        if v2 {
            fmt.Println(v27, fi28)
            fmt.Println(v27, v3, v2)
        } else {
            var v29 string = "hello"
            if (v2 || true) {
                fmt.Println(v27, v29)
                if true {
                    var v30 uint32 = 32
                    _ = v30
                } else {
                    v17 = v17
                }
                fmt.Println(v17, v27)
            }
            var v31 int64 = -1
            _ = v29
            _ = v31
        }
        for fi32 := 0; fi32 < 15; fi32++ {
            if false {
                var v33 int = fi28
                v34 := GenS3{F0: v3, F1: (false && (42 != -1))}
                _ = v33
                _ = v34
            }
            var v35 uint64 = 18446744073709551615
            var v36 int = (-(fi32))
            v37 := GenS4{F0: v17, F1: 0}
            fmt.Println(v36)
            _ = fi32
            _ = v35
            _ = v36
            _ = v37
        }
        _ = fi28
    }
    fmt.Println(v3)
    if true {
        for fi38 := 0; fi38 < 17; fi38++ {
            var v39 float32 = 1e10
            v17 = -1
            var v40 uint8 = genFn0(42, (!(!true)))
            _ = fi38
            _ = v39
            _ = v40
        }
    } else {
        var v41 uint = 498
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v24
    _ = v27
    _ = v3
}
