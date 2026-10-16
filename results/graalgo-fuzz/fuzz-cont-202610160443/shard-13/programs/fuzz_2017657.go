package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint16 }
func genFn0(p0 uint8, p1 uint, p2 uint) uint {
    if (!(true && ((^(p0)) == p0))) {
        v7 := []int{-952, 32768, -543, -9223372036854775808}
        fmt.Println(p0, p0, p1)
        for fi8 := 0; fi8 < 3; fi8++ {
            p0 += 89
            _ = len(v7)
            _ = fi8
        }
        _ = v7
    }
    v9 := GenS1{F0: 32767, F1: false, F2: 918}
    v10 := []int{}
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 float64 = -2.25
        var v13 uint64 = 42
        v10 = append(v10, fi11)
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint16 = 668
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v14
    _ = v9
    return p1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    v3 := []string{"hello", (("a" + "foo bar") + ((("hello" + "\n\t") + "") + "")), "\n\t", "a"}
    var v4 int8 = -30
    var v5 float64 = 0.1
    var v6 int = -47
    if true {
        var v15 int64 = 42
        fmt.Println(v6, v4)
        _ = v15
    }
    var v16 int8 = (v4 % ((-1) | 1))
    var v17 float32 = 1.5
    for fi18 := 0; fi18 < 17; fi18++ {
        _ = len(v3)
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := GenS1{F0: 800, F1: (65535 < 643), F2: 42}
            v3 = append(v3, "foo bar")
            v21 := GenS1{F0: 255, F1: false, F2: 255}
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 int = fi18
        _ = fi18
        _ = v22
    }
    fmt.Println(v5, v5, v4)
    var v23 float32 = 1e10
    if (!true) {
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 int32 = 0
            v26 := GenS1{F0: 1, F1: true, F2: 32767}
            v27 := GenS1{F0: 935, F1: (-39 < v4), F2: 536}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
    }
    var v28 uint64 = 18446744073709551615
    var v29 uint = 877
    v28 += v28
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
