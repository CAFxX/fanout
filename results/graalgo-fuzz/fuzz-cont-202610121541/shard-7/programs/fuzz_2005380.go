package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 string }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint8 }
func genFn0() int8 {
    var v6 uint = 255
    v7 := []int{-1, 942, 32768, 54}
    var v8 float64 = 0.1
    _ = v6
    _ = v7
    _ = v8
    return -128
}

func main() {
    v1 := []string{(((("hello" + "\n\t") + "\n\t") + "") + "foo bar"), ("hello" + "hello"), "a", "foo bar"}
    var v2 float64 = -2.25
    var v3 int16 = 757
    var v4 int32 = -909
    v5 := make(map[string]int)
    switch -633 {
    case 9223372036854775807:
        _ = len(v5)
    case 127:
        v2 = ((v2 - v2) - 1.5)
    }
    v2 -= v2
    fmt.Println(v3)
    v9 := GenS1{F0: 0, F1: 118, F2: 2}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v4 -= v4
        _ = fi10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        delete(v5, "a")
        _ = fi11
    }
    _ = len(v1)
    var v12 uint64 = 42
    if (338 == -1) {
        v13 := GenS1{F0: -58, F1: 128, F2: 1}
        var v14 float32 = 1.5
        _ = v13
        _ = v14
    }
    var v15 int16 = (v3 - (v3 / ((181) | 1)))
    for fi16 := 0; fi16 < 7; fi16++ {
        v4 -= v4
        v2 -= v2
        _ = fi16
    }
    fmt.Println(v15)
    if false {
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := GenS2{F0: v12, F1: ("\n\t" + "hello"), F2: ("" + "")}
            _ = fi17
            _ = v18
        }
    } else {
        v19 := GenS3{F0: v15, F1: (v12 * 1)}
        var v20 int64 = -1
        _ = v19
        _ = v20
    }
    if true {
        if false {
            v21 := v5["\n\t"]
            _ = v21
        } else {
            var v22 uint32 = 0
            delete(v5, "a")
            _ = v22
        }
        v23 := GenS4{F0: v15, F1: (v12 * 1), F2: 74}
        _ = v23
    }
    v12 += 64
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
