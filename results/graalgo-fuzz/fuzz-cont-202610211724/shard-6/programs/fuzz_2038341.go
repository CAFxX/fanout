package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32; F1 uint; F2 int }
func genFn0() int8 {
    switch 267 {
    case -406:
        v7 := []int{}
        _ = v7
    case 9223372036854775807:
        var v8 uint64 = 4294967295
        _ = v8
    default:
        var v9 uint32 = 0
        _ = v9
    }
    var v10 uint = 1
    var v11 float64 = float64(0.1)
    _ = v10
    _ = v11
    return 42
}
func genFn1(p0 string, p1 uint16, p2 string) string {
    fmt.Println(p1, p1, p1)
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 42
    v4 := make(map[int]uint32)
    var v5 int16 = 0
    v6 := make(map[string]string)
    v12 := []int{}
    if true {
        switch (("a" + "a") + "\n\t") {
        case "\n\t":
            v3 = v3
        case "foo bar":
            v12 = append(v12, -130)
        case "":
            v13 := GenS1{F0: (^(v3)), F1: 42}
            _ = v13
        default:
            var v14 int8 = 1
            _ = v14
        }
        if true {
            v12 = append(v12, 9223372036854775807)
        }
        v15 := GenS1{F0: (^(v3)), F1: 272}
        _ = v15
    }
    v3 += (v3 / ((18) | 1))
    var v16 int32 = -394
    v17 := GenS1{F0: v3, F1: 243}
    v5 = (v5 << ((42) & 63))
    for fi18 := 0; fi18 < 12; fi18++ {
        switch "hello" {
        case "a":
            var v19 string = ""
            _ = v19
        }
        _ = fi18
    }
    if true {
        switch 32767 {
        case 32768:
            fmt.Println(v3, v5, v16)
        }
        _ = len(v12)
    } else {
        v20 := GenS2{F0: (float64((3.14159 * 3.14159)) != 1.5), F1: (v5 + (v5 ^ v5))}
        for fi21 := 0; fi21 < 19; fi21++ {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            fmt.Println(fi21)
            _ = fi21
        }
        _ = v20
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        v5 = 32767
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS2{F0: true, F1: v5}
            v25 := GenS3{F0: 0, F1: 255, F2: (fi22 << ((42) & 63))}
            _ = len(v12)
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    v5 += 32767
    _ = len(v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
