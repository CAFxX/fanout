package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint; F2 int64 }
func genFn0(p0 uint) int32 {
    v4 := []int{}
    var v5 uint = 18446744073709551615
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int16 = -1
        var v8 int32 = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    v5 = v5
    var v9 uint16 = 166
    _ = p0
    _ = v4
    _ = v5
    _ = v9
    return -2147483648
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    var v3 uint = 961
    var v10 int64 = -357
    fmt.Println(v3, v3)
    if false {
        fmt.Println(v3, v10)
    }
    switch (("" + "a") + ("a" + "")) {
    case "a":
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint16 = 32767
            _ = fi11
            _ = v12
        }
    case "hello":
        var v13 int8 = 100
        _ = v13
    case "foo bar":
        switch 128 {
        case 42:
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 int16 = -32768
                v10 += 0
                var v16 uint64 = 942
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            if true {
                v17 := GenS1{F0: 127}
                _ = v17
            } else {
                v3 = v3
                for fi18 := 0; fi18 < 9; fi18++ {
                    var v19 string = ""
                    v20 := GenS1{F0: 42}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
        }
    default:
        v2["a"] = false
    }
    v21 := []int{}
    var v22 int32 = 996
    var v23 uint16 = 131
    switch 9223372036854775807 {
    case -981:
        var v24 uint32 = 505
        _ = v24
    default:
        if len(v21) > 0 {
            v21[0] = 128
        }
    }
    _ = len(v21)
    v25 := GenS2{F0: 18446744073709551615, F1: uint(v23), F2: (v10 ^ (v10 % (((v10 - (v10 % ((v10) | 1)))) | 1)))}
    v10 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
}
