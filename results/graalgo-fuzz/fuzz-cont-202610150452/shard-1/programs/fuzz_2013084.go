package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int16; F1 int32 }
func genFn0() uint64 {
    if false {
        var v7 uint = 4294967295
        v8 := []int{830, 127, -2, -9223372036854775808}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    var v10 uint16 = 2
    _ = v10
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 float64 = -2.25
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    var v5 uint32 = 4294967295
    v6 := []uint32{}
    fmt.Println(v5)
    v11 := GenS1{F0: 255, F1: -1}
    var v12 int = 128
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(v2, fi13)
        var v14 int16 = 0
        _ = fi13
        _ = v14
    }
    if false {
        switch (v12 & v12) {
        case 127:
            if (949 == 487) {
                var v15 int64 = 1
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 float32 = 1e10
                v18 := GenS1{F0: 2, F1: 36}
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        v12 = (v12 & v12)
        if (!true) {
            if (true && true) {
                v2 = v2
                v2 = 1e10
            } else {
                fmt.Println(v2, v2, v2)
                if true {
                    var v19 float32 = 0.0
                    v19 = (float32((float64(v19) + (1.5 - 1e10))) * v19)
                    if ((!true) && true) {
                        v5 -= v5
                        v19 = v19
                        for fi20 := 0; fi20 < 9; fi20++ {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = fi20
                        }
                    }
                    _ = v19
                }
            }
            var v21 int8 = 42
            _ = v21
        }
    } else {
        var v22 bool = (1 != 411)
        _ = v22
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        v12 += v12
        switch "foo bar" {
        case "a":
            for fi24 := 0; fi24 < 20; fi24++ {
                v25 := GenS1{F0: 0, F1: 1}
                v26 := GenS2{F0: 839, F1: 0}
                _ = fi24
                _ = v25
                _ = v26
            }
        case "":
            v27 := v4["\n\t"]
            _ = v27
        }
        _ = fi23
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        fmt.Println(fi28, fi28)
        v2 -= 0.1
        v2 += (0.1 * 1.5)
        _ = fi28
    }
    v12 = v12
    v29 := GenS2{F0: 638, F1: 42}
    if len(v1) > 0 {
        v1[0] = v5
    }
    fmt.Println(v5, v5, v2)
    var v30 uint16 = 32767
    if len(v6) > 0 {
        v6[0] = 167
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
