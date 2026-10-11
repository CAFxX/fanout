package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int8 }
type GenS2 struct { F0 uint; F1 bool; F2 int }
type GenS3 struct { F0 string; F1 int8; F2 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    v4 := []uint32{2147483647, 42}
    v5 := []string{("hello" + ("" + ("hello" + ""))), "hello", "\n\t", "\n\t"}
    v6 := make(map[uint32]int)
    var v7 int64 = 20
    var v8 uint = 734
    v9 := GenS1{F0: (("foo bar" + "hello") + "hello"), F1: 0, F2: -1}
    v10 := GenS2{F0: v8, F1: ((!true) && (!(-1 < -912))), F2: 2147483647}
    if (true || ((42 < -1) || (1 <= 170))) {
        _ = len(v1)
    } else {
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := v6[358]
            switch 32767 {
            case 24:
                v13 := GenS2{F0: 792, F1: (0.1 > 3.14159), F2: (v12 | -906)}
                _ = v13
            default:
                fmt.Println(v12, v8, v12)
            }
            _ = fi11
            _ = v12
        }
    }
    var v14 uint = v8
    var v15 int16 = -1
    v16 := GenS3{F0: ("hello" + "a"), F1: 1, F2: 65535}
    v17 := GenS1{F0: "foo bar", F1: 42, F2: 100}
    var v18 int = 951
    _ = len(v1)
    v19 := GenS4{F0: -1}
    fmt.Println(v8, v14)
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 uint = v14
        var v22 uint64 = 42
        var v23 int = v18
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := GenS5{F0: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
