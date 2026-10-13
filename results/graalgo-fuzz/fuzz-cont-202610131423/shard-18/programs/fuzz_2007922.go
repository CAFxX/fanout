package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 string; F2 bool }
func genFn0(p0 bool, p1 bool, p2 int64) int16 {
    switch 127 {
    case 276:
        v7 := []int{85, 9223372036854775807, 4294967295, 128}
        _ = v7
    default:
        v8 := []int{-270, -612, 32767, 9223372036854775807}
        _ = v8
    }
    v9 := []int{32767, -1, 1, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -33
}
func genFn1(p0 uint64, p1 int, p2 int8) uint64 {
    var v10 int64 = 9223372036854775807
    v11 := []int{}
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return p0
}

func main() {
    var v1 int = -116
    var v2 float64 = 0.0
    v3 := []uint32{}
    v4 := make(map[string]int)
    var v5 uint8 = 255
    v6 := make(map[uint32]string)
    var v12 bool = false
    v5 -= (v5 << ((874) & 63))
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(v5)
        switch (v1 | (fi13 + 4294967295)) {
        case 127:
            delete(v4, "a")
        default:
            if (-128 != 42) {
                var v14 int16 = 990
                var v15 uint32 = 2
                var v16 uint8 = v5
                _ = v14
                _ = v15
                _ = v16
            }
        }
        var v17 uint8 = v5
        _ = fi13
        _ = v17
    }
    if false {
        v18 := GenS1{F0: 0}
        for fi19 := 0; fi19 < 1; fi19++ {
            for fi20 := 0; fi20 < 7; fi20++ {
                var v21 string = "hello"
                v2 = v2
                _ = fi20
                _ = v21
            }
            _ = len(v6)
            switch v1 {
            case 4294967295:
                v22 := GenS1{F0: 42}
                _ = v22
            default:
                if v12 {
                    v23 := GenS1{F0: 2}
                    var v24 int = 2147483647
                    var v25 string = (("foo bar" + "\n\t") + "")
                    _ = v23
                    _ = v24
                    _ = v25
                }
            }
            _ = fi19
        }
        _ = v18
    }
    var v26 string = "\n\t"
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := GenS1{F0: uint32(v5)}
        _ = fi27
        _ = v28
    }
    v29 := GenS2{F0: (v1 >> ((587) & 63)), F1: "a", F2: false}
    var v30 float32 = 1e10
    v12 = false
    fmt.Println(v12, v2)
    var v31 int = v1
    var v32 string = ("hello" + (v26 + v26))
    var v33 int64 = 42
    fmt.Println(v2)
    var v34 int64 = v33
    if len(v3) > 0 {
        v3[0] = 42
    }
    v26 = ""
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
