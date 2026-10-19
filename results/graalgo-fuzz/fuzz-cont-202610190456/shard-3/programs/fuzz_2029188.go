package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
func genFn0(p0 int, p1 int32) string {
    for fi7 := 0; fi7 < 20; fi7++ {
        p0 = p0
        if true {
            var v8 uint64 = 966
            p1 = int32(fi7)
            var v9 int64 = 1
            _ = v8
            _ = v9
        } else {
            fmt.Println(fi7)
            var v10 int64 = 42
            _ = v10
        }
        _ = fi7
    }
    if false {
        v11 := GenS1{F0: 255}
        var v12 float64 = 1.5
        _ = v11
        _ = v12
    } else {
        if (104 != 127) {
            var v13 uint16 = 255
            _ = v13
        } else {
            var v14 uint = 986
            _ = v14
        }
        if ((0.1 / 3.14159) <= 1e10) {
            if false {
                v15 := []int{-717, p0, p0}
                var v16 int16 = 32767
                _ = v15
                _ = v16
            }
            var v17 uint = 53
            v18 := GenS1{F0: 786}
            _ = v17
            _ = v18
        } else {
            var v19 float32 = 0.0
            _ = v19
        }
    }
    _ = p0
    _ = p1
    return ""
}

func main() {
    v1 := []string{"foo bar", "hello", (((("hello" + (("\n\t" + "a") + ("\n\t" + ("hello" + (("\n\t" + "\n\t") + "hello"))))) + ("\n\t" + "hello")) + "") + ("" + ("foo bar" + "hello"))), "foo bar"}
    var v2 uint32 = 1
    var v3 int16 = 814
    var v4 uint = 406
    v5 := []int{9223372036854775807, -816}
    v6 := []uint32{v2, v2, (v2 >> ((uint(v2)) & 63)), v2}
    for fi20 := 0; fi20 < 6; fi20++ {
        fmt.Println(v2, fi20, v4)
        fmt.Println(v3, fi20, v3)
        if true {
            v21 := GenS2{F0: "foo bar"}
            _ = v21
        }
        _ = fi20
    }
    v3 = v3
    var v22 uint64 = 42
    var v23 int32 = 259
    v3 -= 1
    for fi24 := 0; fi24 < 7; fi24++ {
        fmt.Println(v2)
        _ = fi24
    }
    var v25 int16 = (v3 ^ v3)
    var v26 uint32 = v2
    var v27 int64 = 1
    _ = len(v5)
    v2 += (^(v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
