package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 int16; F1 bool; F2 int8 }
func genFn0(p0 int32) int {
    p0 = p0
    var v7 int64 = 0
    if (!(((0 < 128) && true) || (0.0 != 1e10))) {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int8 = -128
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: (65535 > 4294967295), F1: 4294967295, F2: -1}
        _ = v10
    }
    _ = p0
    _ = v7
    return 128
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 65535
    v3 := make(map[string]bool)
    var v4 uint = 4294967295
    var v5 float32 = 1e10
    v6 := make(map[int]string)
    var v11 float64 = (0.0 + 1.5)
    v12 := GenS1{F0: true, F1: (v4 / ((249) | 1)), F2: 723}
    var v13 float64 = float64(1e10)
    v5 = (v5 - 0.0)
    if ((("" + "hello") + ("a" + "")) >= "a") {
        if false {
            v14 := GenS2{F0: -32768, F1: false, F2: 0}
            v15 := []int{4294967295, genFn0(1), 32768, -2147483648}
            var v16 int = genFn0(-1)
            _ = v14
            _ = v15
            _ = v16
        }
        fmt.Println(v4)
    } else {
        var v17 uint64 = uint64(v2)
        _ = v17
    }
    fmt.Println(v11)
    var v18 float32 = 0.1
    for fi19 := 0; fi19 < 5; fi19++ {
        v11 = 3.14159
        v4 = v4
        for fi20 := 0; fi20 < 11; fi20++ {
            v13 = 1e10
            _ = fi20
        }
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
