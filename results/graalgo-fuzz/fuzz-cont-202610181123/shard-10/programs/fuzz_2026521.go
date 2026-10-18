package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int16; F2 uint64 }
func genFn0(p0 int32) string {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 bool = (p0 >= (p0 * (p0 / (((^(p0))) | 1))))
        fmt.Println(v6)
        v7 := GenS1{F0: -233}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: -1}
    var v9 uint = 65535
    var v10 int8 = 0
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return ((("foo bar" + ("foo bar" + ("foo bar" + (("hello" + ("foo bar" + ("\n\t" + "hello"))) + "")))) + "foo bar") + "\n\t")
}

func main() {
    var v1 uint = 65535
    v2 := make(map[int]uint32)
    v3 := []string{"\n\t", "", ((("" + (("foo bar" + ("foo bar" + ((((("hello" + "\n\t") + "") + "a") + ("foo bar" + (((("hello" + "hello") + ("\n\t" + "")) + "foo bar") + "\n\t"))) + (("hello" + "foo bar") + "")))) + "hello")) + "a") + ((("\n\t" + "\n\t") + "\n\t") + "a")), "hello"}
    var v4 bool = (!true)
    if true {
        v11 := GenS3{F0: 0}
        var v12 uint = v1
        v13 := GenS1{F0: 42}
        _ = v11
        _ = v12
        _ = v13
    } else {
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    }
    v14 := GenS4{F0: 42}
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 uint64 = 4294967295
        v17 := GenS5{F0: 2, F1: 1000, F2: v16}
        var v18 int8 = -1
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := GenS5{F0: v1, F1: -1, F2: 707}
    var v20 int8 = -18
    var v21 int32 = 2147483647
    var v22 uint = (v1 % ((v1) | 1))
    v23 := GenS1{F0: 708}
    fmt.Println(v20, v1, v21)
    fmt.Println(v1)
    var v24 float32 = (0.0 + 0.0)
    var v25 uint16 = 0
    _ = len(v2)
    if (v4 && false) {
        v26 := GenS1{F0: 760}
        if (v20 >= v20) {
            var v27 float64 = 3.14159
            var v28 int32 = -2147483648
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v27
            _ = v28
        }
        fmt.Println(v25, v20, v22)
        _ = v26
    }
    var v29 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
