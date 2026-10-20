package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint32; F2 int64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 string, p1 int32) uint8 {
    p0 = p0
    v7 := GenS1{F0: 65535}
    if (true && false) {
        for fi8 := 0; fi8 < 10; fi8++ {
            p0 = p0
            _ = fi8
        }
    } else {
        p1 -= (p1 ^ p1)
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
    }
    switch "\n\t" {
    case "a":
        var v11 string = "hello"
        _ = v11
    default:
        if true {
            v12 := GenS1{F0: 2147483647}
            for fi13 := 0; fi13 < 9; fi13++ {
                fmt.Println(p1, p1)
                if (42 <= 813) {
                    var v14 int32 = p1
                    v15 := GenS2{F0: true}
                    var v16 uint64 = 4294967295
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = fi13
            }
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 2
}

func main() {
    var v1 uint8 = 0
    var v2 float32 = (3.14159 + 1e10)
    v3 := make(map[uint32]string)
    var v4 int8 = 0
    v5 := []string{}
    var v6 bool = true
    var v17 float64 = 1.5
    if len(v5) > 0 {
        v5[0] = ""
    }
    _ = len(v5)
    switch 32767 {
    case 128:
        v6 = v6
    }
    var v18 uint16 = 0
    var v19 uint64 = 758
    v6 = false
    v20 := GenS3{F0: int8(v18), F1: 4294967295, F2: -1}
    var v21 float32 = float32((v17 * (float64(v2) * 3.14159)))
    fmt.Println(v21, v17)
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 uint16 = 786
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 int32 = 42
            v3[0] = "a"
            _ = fi24
            _ = v25
        }
        fmt.Println(v17)
        _ = fi22
        _ = v23
    }
    var v26 float32 = (0.0 - v2)
    v21 = (v26 - float32(v17))
    var v27 string = ("a" + "")
    v28 := GenS2{F0: (float64(v21) >= v17)}
    switch v27 {
    case "\n\t":
        var v29 int32 = 1
        _ = v29
    }
    if (true && ((false && v6) && v6)) {
        v2 = (v21 / v2)
    } else {
        var v30 float32 = (v21 + (float32((v17 / (v17 + v17))) + -2.25))
        v27 = ""
        var v31 uint32 = uint32(v1)
        _ = v30
        _ = v31
    }
    switch v27 {
    case "a":
        v32 := GenS4{F0: -516}
        _ = v32
    case "foo bar":
        var v33 int = int(v1)
        _ = v33
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
