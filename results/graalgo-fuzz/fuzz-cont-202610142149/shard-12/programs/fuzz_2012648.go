package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint; F1 int16; F2 int }
func genFn0(p0 uint32) uint16 {
    v6 := []int{}
    fmt.Println(p0)
    if true {
        var v7 int8 = -1
        for fi8 := 0; fi8 < 9; fi8++ {
            fmt.Println(p0)
            _ = fi8
        }
        _ = v7
    } else {
        if (true && false) {
            v9 := GenS1{F0: 32767, F1: "foo bar"}
            _ = v9
        }
    }
    var v10 int64 = -1
    _ = p0
    _ = v10
    _ = v6
    return 65535
}

func main() {
    var v1 uint16 = 32767
    var v2 int32 = 0
    var v3 int32 = 122
    v4 := []string{}
    v5 := make(map[int]uint32)
    var v11 int64 = 1
    switch "" {
    case "":
        v12 := v5[186]
        _ = v12
    case "foo bar":
        if false {
            v3 -= v2
            v2 = (v2 + (-(v2)))
            var v13 uint = 798
            _ = v13
        }
    }
    var v14 float32 = (((float32(((0.0 + float64(float32((0.1 / (3.14159 - 0.0))))) * float64((-2.25 - -2.25)))) / float32(1.5)) - (3.14159 / (3.14159 * -2.25))) * float32(0.1))
    v2 = 386
    v2 += 2147483647
    v15 := GenS1{F0: 42, F1: ("foo bar" + ("foo bar" + "foo bar"))}
    for fi16 := 0; fi16 < 18; fi16++ {
        v1 = (v1 * v1)
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 float64 = 1e10
            v2 -= v2
            var v19 uint = 65535
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    fmt.Println(v14, v3, v11)
    v14 = v14
    var v20 float64 = float64((((1e10 / v14) / float32((1e10 * -2.25))) - (v14 + v14)))
    v3 += (-(v2))
    var v21 bool = ((136 == 204) && (2 != 2))
    if v21 {
        v20 = (v20 * 1.5)
        for fi22 := 0; fi22 < 10; fi22++ {
            if len(v4) > 0 {
                v4[0] = ("hello" + ("hello" + "a"))
            }
            _ = fi22
        }
        v23 := v5[-293]
        _ = v23
    } else {
        v24 := GenS2{F0: 1, F1: 0, F2: 128}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v25 uint = 2
        _ = v24
        _ = v25
    }
    if false {
        var v26 int8 = -122
        fmt.Println(v21)
        _ = v26
    }
    v27 := GenS1{F0: 1, F1: ""}
    v14 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
