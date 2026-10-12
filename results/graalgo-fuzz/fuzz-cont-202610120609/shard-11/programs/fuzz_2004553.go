package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
type GenS2 struct { F0 uint }
func genFn0(p0 bool, p1 uint8) uint64 {
    v6 := []int{1, 873}
    p0 = p0
    p1 += p1
    var v7 bool = (1e10 <= 0.1)
    var v8 uint = 42
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 int32, p1 string) int {
    if (2 == 127) {
        p0 = p0
    } else {
        v9 := []int{-9223372036854775808, 1, 32767, 952}
        if true {
            switch p1 {
            case "a":
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 int16 = -32768
                    var v12 int32 = (p0 ^ (p0 & -423))
                    for fi13 := 0; fi13 < 20; fi13++ {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        fmt.Println(v11, fi13)
                        fmt.Println(p0, fi10, v11)
                        _ = fi13
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            default:
                var v14 uint64 = 0
                _ = v14
            }
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 int8 = -1
                var v17 float32 = ((0.0 * (0.1 / (1.5 + 0.0))) * (float32((0.1 / 1.0)) * 1.5))
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(p1, p1)
            _ = fi18
        }
        _ = v9
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 int = fi19
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    return 796
}

func main() {
    v1 := []int{308}
    var v2 uint8 = 127
    v3 := make(map[string]uint32)
    v4 := []string{("" + "a"), ("hello" + "hello"), "foo bar", "a"}
    var v5 int8 = 42
    var v21 int8 = v5
    v22 := GenS1{F0: 1000, F1: -190, F2: (^(v5))}
    switch "hello" {
    case "foo bar":
        if ((genFn1(-802, "") > 596) || true) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi23 := 0; fi23 < 4; fi23++ {
                v24 := GenS2{F0: 1}
                v2 -= (^(v2))
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 20; fi25++ {
                var v26 float32 = 0.0
                var v27 uint32 = 4294967295
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi25
                _ = v26
                _ = v27
            }
        } else {
            if false {
                var v28 int = -1
                v28 -= v28
                v21 = -1
                _ = v28
            }
        }
    }
    var v29 string = "hello"
    var v30 uint8 = (v2 + v2)
    fmt.Println(v30)
    var v31 int = -1
    v31 = (v31 * v31)
    v29 = v29
    var v32 uint = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
