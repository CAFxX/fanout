package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 int16, p1 uint8, p2 uint16) uint8 {
    p0 = 313
    if ((968 <= 42) && false) {
        p0 = p0
        v5 := GenS1{F0: (false || (((0.1 / (((float32(1.5) / 1e10) * (3.14159 / -2.25)) * 0.1)) / 1.0) >= 1e10)), F1: -128, F2: -614}
        _ = v5
    } else {
        var v6 float64 = float64((0.1 / 1.5))
        v7 := []int{128}
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 string = (("foo bar" + (("a" + (("hello" + ("\n\t" + "hello")) + "\n\t")) + ("a" + ("hello" + ("" + ("\n\t" + (("\n\t" + ((("hello" + "foo bar") + "\n\t") + "hello")) + (((("\n\t" + ("\n\t" + "foo bar")) + "hello") + "foo bar") + "foo bar")))))))) + (("hello" + "\n\t") + "hello"))
    var v2 string = (v1 + (v1 + v1))
    v3 := make(map[string]bool)
    v4 := make(map[string]string)
    for fi8 := 0; fi8 < 18; fi8++ {
        v1 = ("a" + v1)
        v2 = ("\n\t" + ("foo bar" + (v2 + v2)))
        _ = fi8
    }
    if false {
        v2 = "a"
    }
    if true {
        var v9 string = v2
        if (2 <= 18446744073709551615) {
            v1 = v2
            v10 := GenS2{F0: "", F1: 336}
            var v11 float32 = 1e10
            _ = v10
            _ = v11
        } else {
            var v12 int8 = 127
            v1 = v2
            fmt.Println(v2, v12)
            _ = v12
        }
        fmt.Println(v2, v1, v2)
        _ = v9
    } else {
        var v13 uint8 = 255
        _ = v13
    }
    delete(v4, (("foo bar" + v2) + ("" + (v2 + v2))))
    fmt.Println(v2, v1)
    v14 := GenS2{F0: "a", F1: 9223372036854775807}
    var v15 int = -673
    var v16 uint32 = 0
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int = (v15 + 42)
        var v19 uint32 = v16
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 float64 = (1e10 + float64((1e10 + float32(((3.14159 + (0.0 + 0.0)) + (float64(0.1) - 0.0))))))
    var v21 int8 = 127
    var v22 int64 = 9223372036854775807
    var v23 uint = 877
    var v24 uint32 = 2147483647
    var v25 float64 = (-2.25 - ((1.5 + 1.5) * 0.1))
    var v26 uint = uint(v16)
    var v27 uint64 = 914
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
