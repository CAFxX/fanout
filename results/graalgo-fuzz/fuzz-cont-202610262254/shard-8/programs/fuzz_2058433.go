package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 string; F1 uint16; F2 uint16 }
func genFn0(p0 uint64, p1 uint32, p2 uint) int16 {
    var v6 int64 = 264
    v7 := GenS1{F0: p2, F1: (2 <= 2), F2: p1}
    if ((42 != -1) && ((false && (((^(p0)) != p0) || (-455 <= (v6 >> ((p2) & 63))))) && false)) {
        p0 = p0
        fmt.Println(p0, p0)
        var v8 uint = p2
        _ = v8
    }
    v9 := GenS1{F0: 255, F1: false, F2: 656}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v9
    return -32768
}

func main() {
    v1 := []uint32{2}
    v2 := []int{-395, 653}
    var v3 int = 42
    var v4 bool = false
    var v5 bool = false
    for fi10 := 0; fi10 < 3; fi10++ {
        v5 = (-83 >= 9223372036854775807)
        v11 := GenS2{F0: "\n\t", F1: 42, F2: 0}
        v4 = false
        _ = fi10
        _ = v11
    }
    if ((!v4) || v5) {
        v12 := GenS1{F0: 447, F1: true, F2: 256}
        fmt.Println(v5)
        _ = v12
    } else {
        v13 := GenS1{F0: 1, F1: (true && v4), F2: 0}
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 int8 = 0
            v15 -= v15
            _ = fi14
            _ = v15
        }
        _ = len(v1)
        _ = v13
    }
    v4 = (255 == 42)
    v16 := GenS2{F0: "hello", F1: 1, F2: 2}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 = v5
    v4 = true
    v3 = 0
    switch "hello" {
    case "a":
        v4 = (v3 > v3)
    case "hello":
        var v17 uint = 235
        _ = v17
    case "":
        switch "hello" {
        case "foo bar":
            v4 = (4294967295 >= 265)
        }
    }
    var v18 int32 = 162
    var v19 int64 = 9223372036854775807
    var v20 float64 = float64(0.1)
    v19 = (v19 % ((1) | 1))
    v3 += (-(v3))
    v19 -= 0
    v5 = (128 < 8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
