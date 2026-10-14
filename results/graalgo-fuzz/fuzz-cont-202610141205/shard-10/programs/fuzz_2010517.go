package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int64) uint64 {
    fmt.Println(p0)
    if true {
        var v7 uint64 = 1
        _ = v7
    }
    _ = p0
    return 42
}

func main() {
    var v1 uint16 = 971
    v2 := []uint32{693, uint32(v1), 0, 853}
    var v3 bool = (!(0 <= -32768))
    v4 := []int{4294967295}
    var v5 float32 = (3.14159 - 0.1)
    v6 := []string{"hello", "hello"}
    v8 := GenS1{F0: ((!true) || v3)}
    switch -9223372036854775808 {
    case 2147483647:
        var v9 int64 = 42
        _ = v9
    case 582:
        v5 = 1e10
    }
    fmt.Println(v5, v1)
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 float64 = 0.0
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 uint32 = uint32(v1)
            if true {
                fmt.Println(v11)
            }
            _ = fi12
            _ = v13
        }
        var v14 uint64 = 0
        _ = fi10
        _ = v11
        _ = v14
    }
    v15 := GenS1{F0: (1 < 1000)}
    fmt.Println(v3, v3, v3)
    v1 = (v1 | (^(v1)))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
