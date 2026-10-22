package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
type GenS3 struct { F0 bool; F1 string }
func genFn0(p0 int32) int16 {
    fmt.Println(p0)
    if (100 >= 127) {
        var v7 int16 = -32768
        _ = v7
    }
    _ = p0
    return 106
}
func genFn1(p0 int16, p1 int64) int16 {
    var v8 uint32 = 42
    v9 := GenS1{F0: 4294967295}
    fmt.Println(v8)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (p0 * 0)
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]int)
    var v3 int = -1
    v4 := []string{(((("\n\t" + "") + ("hello" + "hello")) + "a") + "a"), "a"}
    var v5 bool = true
    var v6 bool = v5
    v10 := GenS1{F0: 42}
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            _ = len(v2)
            var v13 bool = true
            var v14 uint = 42
            _ = fi12
            _ = v13
            _ = v14
        }
        fmt.Println(v3)
        v5 = false
        _ = fi11
    }
    v6 = v5
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 float32 = 1.5
        var v17 int16 = -32768
        v18 := GenS1{F0: 91}
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 0}
    _ = len(v2)
    v5 = v5
    v1 = v1
    v5 = ((0 <= 656) && false)
    if (((^(v1)) >= -827) && (921 == 597)) {
        v6 = v6
        var v20 uint8 = 0
        _ = v20
    }
    v1 = (-(v1))
    _ = len(v4)
    fmt.Println(v1, v6)
    v21 := GenS2{F0: 763, F1: -128, F2: 193}
    if (v6 || (!(v6 && true))) {
        switch (v3 | v3) {
        case 1:
            var v22 int8 = -1
            _ = v22
        case 127:
            var v23 int8 = -87
            _ = v23
        case 32768:
            v6 = (42 >= 65535)
        }
        if v6 {
            fmt.Println(v5, v5)
        } else {
            var v24 uint32 = 967
            v25 := GenS3{F0: v6, F1: "\n\t"}
            _ = v24
            _ = v25
        }
        var v26 int8 = -128
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
