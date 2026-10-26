package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 bool; F1 int8 }
func genFn0(p0 string) uint64 {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 42, F1: 297, F2: 2}
        p0 = p0
        _ = fi7
        _ = v8
    }
    v9 := []int{127, 460}
    var v10 float32 = (3.14159 - 3.14159)
    _ = p0
    _ = v10
    _ = v9
    return 934
}

func main() {
    v1 := []int{-778}
    v2 := make(map[uint32]uint32)
    var v3 int16 = -130
    v4 := []string{("a" + ("foo bar" + "foo bar")), "a", ("a" + "hello")}
    var v5 int16 = v3
    var v6 string = ""
    fmt.Println(v3, v3, v5)
    var v11 int64 = 1
    v12 := GenS1{F0: 42, F1: (v5 - 32767), F2: 65535}
    if true {
        var v13 uint = 855
        _ = v13
    } else {
        v5 -= 292
    }
    if false {
        if false {
            _ = len(v1)
        }
        switch 568 {
        case 128:
            v2[2] = 0
        case 32767:
            var v14 uint = 989
            _ = v14
        }
        v15 := GenS2{F0: 42, F1: 1}
        _ = v15
    } else {
        fmt.Println(v6)
    }
    var v16 uint8 = 127
    var v17 uint8 = 128
    switch v6 {
    case "foo bar":
        v18 := GenS3{F0: true, F1: 42}
        _ = v18
    case "a":
        v16 -= (v16 - v17)
    }
    v19 := v2[0]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
