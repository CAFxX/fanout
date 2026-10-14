package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 string; F1 int8 }
func genFn0(p0 uint64) bool {
    for fi5 := 0; fi5 < 10; fi5++ {
        p0 = p0
        _ = fi5
    }
    var v6 float32 = 1.5
    var v7 uint = 796
    if (-41 == 100) {
        p0 = (^(p0))
    } else {
        v6 = v6
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 float64 = 0.0
            v6 += 0.0
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return true
}
func genFn1() uint8 {
    switch "hello" {
    case "":
        v10 := GenS1{F0: "foo bar"}
        _ = v10
    case "hello":
        var v11 uint64 = 4294967295
        _ = v11
    case "a":
        var v12 int16 = 0
        _ = v12
    default:
        v13 := GenS2{F0: 64, F1: (2147483647 != 689)}
        _ = v13
    }
    v14 := GenS3{F0: "", F1: 119}
    var v15 bool = ((32767 > 203) && false)
    _ = v14
    _ = v15
    return 0
}

func main() {
    v1 := []uint32{42, 36, 2147483647, 65535}
    var v2 uint16 = 266
    var v3 uint = 673
    v4 := make(map[uint32]uint32)
    v16 := GenS2{F0: 1, F1: false}
    var v17 float64 = (3.14159 - -2.25)
    _ = len(v4)
    v2 += (^(v2))
    var v18 uint16 = 65535
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(v2, v17, v18)
        _ = fi19
    }
    v20 := GenS3{F0: ("hello" + "a"), F1: 42}
    v2 += v18
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint16 = 2
        var v23 string = "a"
        fmt.Println(v22, v2)
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v3, v17)
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
