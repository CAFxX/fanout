package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint32) int32 {
    if false {
        if false {
            v4 := GenS1{F0: ("hello" + "hello"), F1: 883}
            _ = v4
        }
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 uint64 = 759
            v7 := []int{fi5}
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    v8 := []int{}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v9 uint32 = p0
    v10 := GenS1{F0: "", F1: 740}
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 uint64 = 0
    v3 := make(map[uint32]uint32)
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 bool = (fi11 <= -1)
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: "", F1: -9223372036854775808}
    var v14 int8 = -128
    var v15 int16 = 0
    v3[458] = 190
    switch (("" + "") + "hello") {
    case "":
        var v16 bool = (!false)
        _ = v16
    }
    fmt.Println(v2, v15)
    if true {
        var v17 int8 = v14
        fmt.Println(v15, v15, v15)
        var v18 int = 42
        _ = v17
        _ = v18
    }
    var v19 uint = 0
    _ = len(v3)
    v20 := GenS2{F0: 165, F1: 127}
    v21 := GenS2{F0: 189, F1: 190}
    if false {
        fmt.Println(v15, v19)
        v2 = (v2 - (v2 - v2))
    }
    if ((v2 * (v2 >> ((542) & 63))) != v2) {
        v14 = (v14 * v14)
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        if (128 < 232) {
            v15 = 32767
            var v23 uint = 42
            _ = v23
        } else {
            _ = len(v1)
        }
        _ = fi22
    }
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
