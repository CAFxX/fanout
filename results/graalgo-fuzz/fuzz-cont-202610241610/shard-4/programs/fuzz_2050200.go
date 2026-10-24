package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint64 }
func genFn0(p0 int, p1 bool) int64 {
    fmt.Println(p1, p1)
    fmt.Println(p0, p1)
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: (65 >= 65535), F1: 0, F2: 18446744073709551615}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    var v3 int16 = -777
    var v4 bool = true
    v5 := make(map[int]bool)
    var v8 int32 = -1
    switch 134 {
    case 190:
        var v9 int8 = 1
        _ = v9
    case 2147483647:
        v8 += (v8 * v8)
    case 435:
        for fi10 := 0; fi10 < 12; fi10++ {
            delete(v5, (fi10 % (((^(fi10))) | 1)))
            var v11 uint16 = 255
            _ = fi10
            _ = v11
        }
    }
    v4 = (!true)
    _ = len(v1)
    v4 = (4294967295 == 0)
    v4 = v4
    var v12 float32 = 0.0
    var v13 uint32 = 2147483647
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 int16 = 42
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 bool = v4
        v18 := GenS1{F0: (true && (-1 == (v8 ^ (-(v8))))), F1: 42, F2: 4294967295}
        v3 = (v3 ^ v3)
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
