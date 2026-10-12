package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 bool) int8 {
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            fmt.Println(fi5, p0, p0)
            var v7 int64 = 42
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    p0 = p0
    v8 := []int{-9223372036854775808}
    _ = p0
    _ = v8
    return 25
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int8 = 100
    v3 := make(map[string]bool)
    var v4 float64 = 1.5
    delete(v3, "foo bar")
    _ = len(v3)
    v1 += v1
    switch 4294967295 {
    case -9223372036854775808:
        v9 := GenS1{F0: 99}
        _ = v9
    case 32768:
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 uint32 = v1
            v12 := []int{(^(fi10)), (fi10 | (fi10 << ((716) & 63))), 4294967295}
            _ = fi10
            _ = v11
            _ = v12
        }
    default:
        v13 := []int{-188, -472, 666, 2147483647}
        _ = v13
    }
    fmt.Println(v4)
    v1 = (^(v1))
    var v14 uint = 1
    v15 := GenS1{F0: 11}
    for fi16 := 0; fi16 < 16; fi16++ {
        v2 += v2
        v17 := GenS1{F0: 1}
        v18 := GenS1{F0: 236}
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint16 = 1
    var v20 int16 = 998
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
