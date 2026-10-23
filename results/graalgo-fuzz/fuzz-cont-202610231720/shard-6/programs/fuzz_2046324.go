package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int = fi5
        var v7 float64 = -2.25
        switch ("foo bar" + (("" + "a") + ("a" + ("hello" + ("hello" + "foo bar"))))) {
        case "a":
            if false {
                for fi8 := 0; fi8 < 14; fi8++ {
                    var v9 int8 = 42
                    _ = fi8
                    _ = v9
                }
                if (774 >= 42) {
                    var v10 int16 = -32768
                    var v11 int64 = 0
                    _ = v10
                    _ = v11
                }
            }
        default:
            v12 := []int{fi5, (v6 * 9223372036854775807), (v6 * (fi5 + v6))}
            _ = v12
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v13 := GenS1{F0: 1}
    switch "hello" {
    case "a":
        v14 := []int{4294967295, -304}
        _ = v14
    case "foo bar":
        var v15 float32 = -2.25
        _ = v15
    case "hello":
        for fi16 := 0; fi16 < 20; fi16++ {
            v17 := GenS1{F0: 65535}
            fmt.Println(fi16)
            v18 := []int{fi16, fi16, fi16, fi16}
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    _ = v13
    return 1
}
func genFn1(p0 int16, p1 uint32) uint8 {
    var v19 uint32 = (p1 >> ((42) & 63))
    v20 := GenS1{F0: 0}
    var v21 int = int(p0)
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v21
    return 128
}

func main() {
    var v1 uint32 = 702
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    v4 := make(map[uint32]string)
    v1 += v1
    if (1 <= 0) {
        fmt.Println(v1)
    } else {
        if true {
            fmt.Println(v1)
            var v22 bool = (0.1 > (-2.25 / 3.14159))
            _ = v22
        }
    }
    v3 = append(v3, (^(v1)))
    v1 = (v1 | 0)
    v1 = 2
    var v23 uint16 = genFn0()
    if (-357 != -1) {
        var v24 uint = 2
        fmt.Println(v23)
        _ = v24
    }
    var v25 int = 127
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
