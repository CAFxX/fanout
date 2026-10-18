package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint64) int16 {
    v5 := []int{-1, -283, 2147483647}
    v6 := GenS1{F0: true, F1: 88, F2: 32767}
    var v7 float32 = 1e10
    v7 -= 0.1
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 32767
}

func main() {
    v1 := []string{"\n\t", "foo bar", "foo bar"}
    v2 := []int{-437}
    v3 := make(map[uint32]int)
    var v4 int8 = 1
    fmt.Println(v4)
    switch 4294967295 {
    case 4294967295:
        v8 := GenS2{F0: (!(2147483647 < 366))}
        _ = v8
    default:
        if false {
            v4 += v4
            var v9 int16 = 42
            _ = v9
        }
    }
    v10 := GenS1{F0: true, F1: -1, F2: 32767}
    fmt.Println(v4)
    v4 = 1
    switch 4294967295 {
    case 443:
        v11 := GenS1{F0: (1 != 2147483647), F1: (v4 % ((v4) | 1)), F2: 682}
        _ = v11
    }
    var v12 float64 = float64(-2.25)
    var v13 uint32 = 740
    fmt.Println(v12, v12, v4)
    if true {
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS2{F0: false}
            v13 -= v13
            _ = fi14
            _ = v15
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v16 int16 = 757
    v13 = (v13 % ((v13) | 1))
    _ = len(v1)
    v4 = (^(v4))
    v12 = (float64(-2.25) * v12)
    var v17 uint8 = 2
    for fi18 := 0; fi18 < 5; fi18++ {
        switch fi18 {
        case 0:
            v17 = uint8(v13)
        case 32767:
            v19 := GenS1{F0: false, F1: v4, F2: 32767}
            _ = v19
        case -954:
            v20 := GenS2{F0: ((((("\n\t" + (("hello" + "a") + "foo bar")) + "") + "hello") + "\n\t") > "hello")}
            _ = v20
        }
        v16 = 32767
        _ = fi18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
