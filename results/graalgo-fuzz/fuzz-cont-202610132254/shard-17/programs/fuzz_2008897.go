package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint; F1 int; F2 int32 }
type GenS3 struct { F0 bool }
func genFn0() int {
    var v7 bool = (-854 != 1)
    v8 := GenS1{F0: -32768, F1: 128, F2: 0}
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(v7)
        _ = fi9
    }
    fmt.Println(v7)
    _ = v7
    _ = v8
    return 984
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{(("a" + ("\n\t" + "\n\t")) + "\n\t"), ("\n\t" + "foo bar")}
    v3 := make(map[string]string)
    v4 := []string{("hello" + ""), ("foo bar" + (("" + "") + "a")), "\n\t"}
    v5 := make(map[uint32]bool)
    var v6 int32 = 0
    v6 -= -948
    v10 := v3[""]
    v10 = (v10 + "")
    v10 = (v10 + ((v10 + v10) + v10))
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(v6)
        fmt.Println(v6, fi11)
        _ = fi11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        fmt.Println(v10, fi12, v10)
        var v13 uint16 = 254
        var v14 float64 = 0.1
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 bool = true
    if v15 {
        v15 = false
        var v16 int32 = (v6 + (v6 | v6))
        _ = v16
    } else {
        v17 := GenS2{F0: 960, F1: 934, F2: (v6 ^ 452)}
        if (-128 != -25) {
            for fi18 := 0; fi18 < 11; fi18++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if true {
                    if (!v15) {
                        v19 := GenS1{F0: 931, F1: 237, F2: 545}
                        v15 = (1 <= -468)
                        _ = len(v1)
                        _ = v19
                    }
                    var v20 string = ("hello" + (v10 + v10))
                    _ = v20
                } else {
                    v21 := GenS1{F0: 1000, F1: 128, F2: -1}
                    _ = v21
                }
                _ = fi18
            }
        }
        if (!v15) {
            var v22 uint32 = 65535
            v6 -= (v6 << ((2) & 63))
            _ = v22
        }
        _ = v17
    }
    v23 := GenS2{F0: 327, F1: 42, F2: (v6 & v6)}
    switch 265 {
    case 878:
        var v24 int8 = -1
        _ = v24
    case 127:
        delete(v5, 42)
    default:
        var v25 bool = (65535 >= 1)
        _ = v25
    }
    var v26 uint8 = 127
    v10 = (("a" + v10) + v10)
    _ = len(v2)
    v27 := GenS3{F0: v15}
    v28 := GenS2{F0: 223, F1: -1, F2: (v6 % ((v6) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
