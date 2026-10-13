package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 string }
type GenS2 struct { F0 bool; F1 int16 }
func genFn0(p0 int32, p1 string, p2 int) uint64 {
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: -1, F1: p0, F2: p1}
    if true {
        for fi9 := 0; fi9 < 7; fi9++ {
            if (v7 > (v7 | (v7 & (^(v7))))) {
                v10 := GenS1{F0: -1, F1: p0, F2: ""}
                var v11 int8 = 102
                _ = v10
                _ = v11
            } else {
                p0 = -2147483648
            }
            var v12 float32 = 1e10
            _ = fi9
            _ = v12
        }
        if true {
            var v13 float32 = float32(0.0)
            p2 += (p2 | p2)
            var v14 uint = 255
            _ = v13
            _ = v14
        }
    }
    fmt.Println(p1, p1, p1)
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint16 = 42
        fmt.Println(v7, v16)
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint64) int8 {
    v17 := GenS2{F0: (!(46 > 128)), F1: 0}
    p0 = 4294967295
    var v18 float32 = 0.1
    _ = p0
    _ = v17
    _ = v18
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    v3 := []string{"hello"}
    var v4 int = 2147483647
    var v5 int8 = -1
    var v6 int64 = 0
    fmt.Println(v4, v4)
    v19 := GenS2{F0: (false || (0 != 320)), F1: 1}
    v20 := GenS1{F0: (v5 >> ((4294967295) & 63)), F1: 0, F2: ("foo bar" + "\n\t")}
    fmt.Println(v5, v5)
    v5 -= genFn1(42)
    if false {
        fmt.Println(v5, v5, v4)
        v4 += v4
    } else {
        v3 = append(v3, ("" + ("hello" + ("\n\t" + "hello"))))
        var v21 int32 = -1
        _ = v21
    }
    var v22 int = v4
    var v23 int8 = (v5 >> ((1) & 63))
    for fi24 := 0; fi24 < 5; fi24++ {
        fmt.Println(v4, v6)
        _ = fi24
    }
    v6 += 0
    var v25 int8 = (^(v23))
    v23 = genFn1(347)
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v26 int = v4
    v25 = 1
    var v27 float32 = -2.25
    var v28 uint32 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
