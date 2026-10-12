package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int; F1 uint8; F2 int16 }
func genFn0() uint {
    var v7 int = 127
    v8 := GenS1{F0: false}
    v9 := GenS2{F0: ("foo bar" + ("" + "\n\t")), F1: 255}
    if ((-9223372036854775808 < 0) && true) {
        v10 := GenS2{F0: ("hello" + "foo bar"), F1: 371}
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := []int{-559, fi11, (v7 ^ fi11), v7}
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        v13 := GenS2{F0: "hello", F1: 0}
        _ = v13
    }
    _ = v7
    _ = v8
    _ = v9
    return 796
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{4294967295}
    v3 := make(map[uint32]string)
    var v4 float32 = 1e10
    v5 := make(map[string]string)
    var v6 uint8 = 0
    for fi14 := 0; fi14 < 16; fi14++ {
        fmt.Println(fi14, v4)
        _ = fi14
    }
    var v15 int16 = -1
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 float32 = (-2.25 * float32(float64(v4)))
        _ = fi16
        _ = v17
    }
    var v18 string = "hello"
    var v19 string = ""
    for fi20 := 0; fi20 < 4; fi20++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi20
    }
    v6 = v6
    v21 := GenS3{F0: 0, F1: v6, F2: (v15 - v15)}
    _ = len(v3)
    _ = len(v1)
    var v22 bool = (127 >= v6)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
