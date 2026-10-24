package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 string }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }
func genFn0(p0 bool) int16 {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{(fi6 % ((fi6) | 1))}
        for fi8 := 0; fi8 < 20; fi8++ {
            p0 = (-1 > 42)
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    var v9 uint = 4294967295
    _ = p0
    _ = v9
    return 1
}

func main() {
    var v1 int8 = 100
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]int)
    fmt.Println(v1)
    switch 42 {
    case 460:
        if true {
            fmt.Println(v1)
        }
    case 32767:
        if (false && true) {
            for fi10 := 0; fi10 < 7; fi10++ {
                fmt.Println(fi10, v1)
                _ = fi10
            }
        }
    case -76:
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint64 = 18446744073709551615
            switch ("hello" + "\n\t") {
            case "foo bar":
                v13 := GenS1{F0: 368, F1: 0, F2: "\n\t"}
                _ = v13
            default:
                v1 += (v1 | 100)
            }
            var v14 uint32 = 42
            _ = fi11
            _ = v12
            _ = v14
        }
    default:
        v15 := []int{646, 42, 127}
        _ = v15
    }
    var v16 uint = 708
    var v17 float64 = ((3.14159 / -2.25) / 1.0)
    fmt.Println(v17, v16, v16)
    if true {
        v18 := GenS1{F0: -199, F1: 0, F2: ("a" + "foo bar")}
        var v19 float64 = v17
        _ = v18
        _ = v19
    }
    var v20 uint16 = 42
    switch (("" + "") + "foo bar") {
    case "\n\t":
        var v21 bool = (-9223372036854775808 >= 0)
        _ = v21
    case "a":
        v16 = v16
    case "":
        if true {
            switch 32768 {
            case 128:
                fmt.Println(v20, v20, v16)
            default:
                if true {
                    v22 := GenS1{F0: 32768, F1: 0, F2: "foo bar"}
                    _ = v22
                }
            }
            switch 0 {
            case 32767:
                fmt.Println(v17, v1, v1)
            case -1:
                v23 := GenS2{F0: 4294967295, F1: 0, F2: v20}
                _ = v23
            }
        }
    default:
        var v24 float64 = v17
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
