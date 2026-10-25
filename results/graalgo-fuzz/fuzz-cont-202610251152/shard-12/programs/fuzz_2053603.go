package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 string }
type GenS2 struct { F0 int16; F1 int8; F2 int64 }
func genFn0() int {
    switch 0 {
    case 2147483647:
        var v4 uint32 = 2147483647
        _ = v4
    case -455:
        var v5 float32 = 3.14159
        _ = v5
    }
    var v6 bool = (true && true)
    var v7 int64 = 0
    v8 := []int{2147483647}
    _ = v6
    _ = v7
    _ = v8
    return 128
}
func genFn1(p0 int8, p1 int32) uint8 {
    v9 := GenS1{F0: (-32768 == -661), F1: -2147483648, F2: ((("hello" + "hello") + ("a" + "foo bar")) + "hello")}
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: ((3.14159 + 1e10) != 1e10), F1: 425, F2: "a"}
        _ = fi10
        _ = v11
    }
    var v12 uint = 1
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := GenS1{F0: false, F1: (fi13 % ((fi14) | 1)), F2: "a"}
            _ = fi14
            _ = v15
        }
        v16 := GenS1{F0: false, F1: 459, F2: "\n\t"}
        _ = fi13
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 int64 = 42
        _ = fi17
        _ = v18
    }
    switch 42 {
    case 32767:
        var v19 int8 = -1
        _ = v19
    case -2147483648:
        v20 := []int{-1, -86, 127, -186}
        _ = v20
    case 376:
        var v21 int64 = 469
        _ = v21
    }
    if true {
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := v2[(("" + "a") + "a")]
            _ = fi22
            _ = v23
        }
        v24 := v2[("hello" + "\n\t")]
        _ = v24
    } else {
        var v25 int8 = 1
        var v26 uint64 = 2
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: true, F1: genFn0(), F2: "\n\t"}
    var v28 uint8 = 1
    v28 += 127
    v29 := GenS2{F0: 328, F1: 127, F2: 9223372036854775807}
    var v30 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
