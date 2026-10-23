package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 string }
type GenS2 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    var v3 float32 = float32(1.5)
    var v4 uint8 = 0
    var v5 int = 2147483647
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 string = "\n\t"
        for fi8 := 0; fi8 < 5; fi8++ {
            v3 = 0.0
            _ = fi8
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS1{F0: -128, F1: v7, F2: (v7 + "hello")}
            var v11 int8 = -1
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    v3 = 0.0
    var v12 int = v5
    for fi13 := 0; fi13 < 19; fi13++ {
        if (!true) {
            v5 = 2147483647
            v4 += (v4 + v4)
            v14 := GenS1{F0: -114, F1: "foo bar", F2: "foo bar"}
            _ = v14
        }
        switch (v5 + (v5 >> ((4294967295) & 63))) {
        case 0:
            var v15 int64 = 0
            _ = v15
        case -1:
            if true {
                var v16 uint16 = 908
                _ = v16
            } else {
                if true {
                    fmt.Println(v5, v12)
                } else {
                    v17 := GenS1{F0: -128, F1: "\n\t", F2: "hello"}
                    v5 = 4294967295
                    _ = v17
                }
                fmt.Println(v4, v12, v3)
                v18 := []int{(fi13 & 333), v5}
                _ = v18
            }
        case 1:
            for fi19 := 0; fi19 < 4; fi19++ {
                if (284 == -32768) {
                    v5 = fi19
                    v3 = float32((1e10 / 1.0))
                    v20 := []int{v5, v5}
                    _ = v20
                }
                fmt.Println(v12, fi19, v3)
                _ = fi19
            }
        }
        _ = fi13
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 int16 = 1
        v23 := []int{fi21, 32767}
        if len(v23) > 0 {
            v23[0] = 279
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    v2[v5] = 2147483648
    v24 := GenS1{F0: 127, F1: ((("a" + "hello") + "\n\t") + "\n\t"), F2: ("hello" + "hello")}
    switch v12 {
    case -9223372036854775808:
        v25 := GenS1{F0: 0, F1: "a", F2: ("hello" + ("hello" + "foo bar"))}
        _ = v25
    case 127:
        var v26 uint64 = 18446744073709551615
        _ = v26
    }
    v12 -= 9223372036854775807
    var v27 uint8 = v4
    v12 -= -160
    for fi28 := 0; fi28 < 16; fi28++ {
        if true {
            v29 := GenS2{F0: 32767, F1: v5}
            _ = v29
        } else {
            fmt.Println(v5, v3)
        }
        fmt.Println(fi28, v3, v3)
        v12 -= (v5 - v5)
        _ = fi28
    }
    if ((v3 != -2.25) || true) {
        var v30 int8 = 1
        _ = v30
    }
    v31 := []int{v12, 9223372036854775807}
    v32 := v2[32767]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
