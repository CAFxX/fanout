package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint32; F1 uint32; F2 string }
func genFn0(p0 uint8) bool {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: ("\n\t" + ("hello" + "foo bar")), F1: 1, F2: 0}
        _ = fi4
        _ = v5
    }
    if false {
        v6 := []int{2147483647, -9223372036854775808, 9223372036854775807}
        if len(v6) > 0 {
            v6[0] = 128
        }
        _ = v6
    }
    if ((0 != 42) || true) {
        var v7 uint8 = p0
        _ = v7
    }
    _ = p0
    return false
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 255
    v3 := make(map[int]int)
    var v8 uint64 = 624
    v8 = 1
    v9 := GenS1{F0: "foo bar", F1: 644, F2: 42}
    switch "a" {
    case "hello":
        v10 := []int{128, -692}
        _ = v10
    }
    if (42 > v8) {
        _ = len(v1)
    } else {
        var v11 int32 = 314
        _ = v11
    }
    var v12 uint16 = 42
    var v13 uint16 = 32767
    v12 -= 1
    if true {
        v14 := GenS2{F0: 2147483647}
        _ = v14
    }
    switch -1 {
    case 0:
        v13 = 32767
    case -508:
        v15 := GenS3{F0: (v13 % ((v13) | 1)), F1: 4294967295, F2: 53}
        _ = v15
    }
    v8 = 4294967295
    if true {
        switch "" {
        case "hello":
            v13 -= 1
        }
    }
    var v16 float32 = -2.25
    var v17 string = "a"
    var v18 string = (v17 + (v17 + v17))
    v19 := GenS4{F0: uint32(v12), F1: 2, F2: v17}
    for fi20 := 0; fi20 < 6; fi20++ {
        if false {
            if true {
                v12 = v13
            } else {
                fmt.Println(v8)
            }
            v21 := v1[(-(fi20))]
            _ = v21
        }
        v22 := []int{}
        _ = fi20
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
