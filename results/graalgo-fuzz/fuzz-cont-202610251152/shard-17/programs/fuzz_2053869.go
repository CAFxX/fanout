package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 string; F1 int32; F2 uint64 }
func genFn0(p0 uint, p1 int) int {
    p1 += (p1 ^ (p1 / (((^(p1))) | 1)))
    var v7 uint = p0
    v8 := []int{(p1 << ((v7) & 63)), p1, -2147483648, (p1 | p1)}
    if true {
        var v9 uint64 = 798
        var v10 int32 = 1
        if (0 != 1) {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int16 = -255
                var v13 bool = false
                var v14 int = fi11
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            v15 := GenS1{F0: -9223372036854775808, F1: ("hello" + "a")}
            v10 = 1
            _ = v15
        } else {
            for fi16 := 0; fi16 < 1; fi16++ {
                v7 = v7
                switch "foo bar" {
                case "":
                    var v17 uint32 = 4294967295
                    _ = v17
                case "hello":
                    var v18 uint8 = 127
                    _ = v18
                }
                var v19 float64 = (-2.25 + 0.1)
                _ = fi16
                _ = v19
            }
        }
        _ = v10
        _ = v9
    } else {
        for fi20 := 0; fi20 < 1; fi20++ {
            _ = len(v8)
            _ = fi20
        }
    }
    if true {
        p1 -= -514
        if len(v8) > 0 {
            v8[0] = (-(p1))
        }
        var v21 bool = (1e10 >= (3.14159 - float32(float64(1e10))))
        _ = v21
    } else {
        if true {
            var v22 string = ""
            v7 = (^(v7))
            _ = v22
        }
        var v23 int32 = -1
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (p1 & 2147483647)
}

func main() {
    var v1 uint32 = 0
    var v2 uint64 = 104
    var v3 int16 = 1
    v4 := make(map[uint32]string)
    v5 := make(map[string]bool)
    var v6 uint64 = v2
    v24 := GenS1{F0: -1, F1: "a"}
    for fi25 := 0; fi25 < 17; fi25++ {
        v26 := GenS1{F0: 530, F1: ("hello" + "hello")}
        var v27 float64 = (-2.25 * 1.5)
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := GenS2{F0: 42, F1: v2}
            v30 := GenS3{F0: 100, F1: v1, F2: 2}
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    var v31 uint = 739
    fmt.Println(v1, v31)
    v6 -= 2
    v32 := GenS4{F0: -1, F1: true}
    v1 += v1
    var v33 int8 = -122
    v34 := v4[v1]
    v35 := GenS5{F0: ("hello" + v34), F1: -1, F2: v6}
    var v36 uint = uint(v1)
    v37 := GenS4{F0: 1, F1: ((((3.14159 / 1.5) - 1.5) / 3.14159) >= (1e10 / 1.0))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
