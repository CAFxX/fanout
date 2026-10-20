package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int64 }
type GenS2 struct { F0 int64 }
func genFn0() uint16 {
    if true {
        var v4 bool = false
        if v4 {
            var v5 string = ("\n\t" + "foo bar")
            _ = v5
        } else {
            var v6 uint16 = 0
            var v7 uint32 = 42
            _ = v6
            _ = v7
        }
        _ = v4
    } else {
        var v8 uint = 42
        v9 := []int{-355, 32768}
        _ = v8
        _ = v9
    }
    if (511 <= 351) {
        var v10 int32 = 0
        v11 := []int{-1, 4294967295, 42, 9223372036854775807}
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 2147483647, F1: 1, F2: 42}
    _ = v12
    return 32767
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{}
    v3 := []string{"\n\t", "hello"}
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 int16 = 32767
        v14 -= v14
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 int32 = 118
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v3) > 0 {
            v3[0] = ""
        }
        _ = fi15
        _ = v16
    }
    switch "foo bar" {
    case "":
        var v17 uint32 = 1
        _ = v17
    default:
        var v18 int = 928
        _ = v18
    }
    v19 := v1[42]
    for fi20 := 0; fi20 < 7; fi20++ {
        fmt.Println(fi20, v19)
        _ = fi20
    }
    v19 = v19
    for fi21 := 0; fi21 < 17; fi21++ {
        if (918 > 510) {
            _ = len(v3)
            if true {
                for fi22 := 0; fi22 < 17; fi22++ {
                    v23 := GenS2{F0: 0}
                    delete(v1, 128)
                    _ = fi22
                    _ = v23
                }
                var v24 uint32 = (v19 >> ((2) & 63))
                _ = v24
            }
        }
        var v25 uint = 0
        var v26 string = "\n\t"
        _ = fi21
        _ = v25
        _ = v26
    }
    v19 += v19
    v19 = (v19 << ((384) & 63))
    fmt.Println(v19)
    v1[0] = v19
    var v27 int = -226
    v27 = v27
    for fi28 := 0; fi28 < 8; fi28++ {
        for fi29 := 0; fi29 < 8; fi29++ {
            switch (("foo bar" + "") + "\n\t") {
            case "hello":
                var v30 uint = 0
                _ = v30
            }
            var v31 float32 = (1e10 - 0.0)
            _ = fi29
            _ = v31
        }
        _ = fi28
    }
    var v32 bool = (119 > -60)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v32
}
