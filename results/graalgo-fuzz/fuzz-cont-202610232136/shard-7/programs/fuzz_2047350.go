package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 string; F1 bool; F2 string }
type GenS3 struct { F0 uint; F1 uint16; F2 string }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 int8 = 0
    v4 := make(map[string]uint32)
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]int)
    var v7 int64 = -932
    v3 -= -128
    var v8 int32 = -1
    v9 := []int{}
    var v10 int8 = (v3 - v3)
    var v11 int64 = -222
    switch ("a" + "") {
    case "a":
        fmt.Println(v11, v7, v8)
    case "hello":
        v8 -= 2147483647
    case "foo bar":
        for fi12 := 0; fi12 < 10; fi12++ {
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := GenS1{F0: 292, F1: false}
                _ = fi13
                _ = v14
            }
            if false {
                var v15 bool = false
                for fi16 := 0; fi16 < 15; fi16++ {
                    v15 = v15
                    fmt.Println(v11)
                    var v17 int = (fi16 - (^(fi16)))
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            _ = fi12
        }
    }
    if false {
        v8 = 875
        switch "foo bar" {
        case "hello":
            v18 := v6[0]
            _ = v18
        case "a":
            if false {
                v8 = v8
                v10 = (v10 >> ((1) & 63))
                var v19 float64 = 1.5
                _ = v19
            } else {
                v11 = v11
            }
        default:
            var v20 float64 = ((0.0 + float64(0.1)) / -2.25)
            _ = v20
        }
        switch 2147483647 {
        case 77:
            var v21 uint8 = 2
            _ = v21
        case -44:
            fmt.Println(v7)
        case -640:
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        default:
            var v22 float32 = 3.14159
            _ = v22
        }
    }
    v8 -= v8
    fmt.Println(v11, v8, v10)
    if true {
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS2{F0: "\n\t", F1: false, F2: "a"}
            _ = fi23
            _ = v24
        }
        if (127 < 32767) {
            fmt.Println(v7)
        } else {
            var v25 int64 = -176
            if false {
                var v26 int32 = v8
                v7 -= 42
                fmt.Println(v8)
                _ = v26
            }
            _ = v25
        }
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 3; fi28++ {
            v29 := GenS3{F0: 717, F1: 1, F2: "foo bar"}
            var v30 int16 = -763
            delete(v5, (fi28 | fi28))
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi27
    }
    var v31 uint8 = 255
    if ((127 >= 889) && ((825 > -1) || (-20 == v10))) {
        delete(v4, ("hello" + "\n\t"))
        v32 := GenS4{F0: 262}
        v10 += v3
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
