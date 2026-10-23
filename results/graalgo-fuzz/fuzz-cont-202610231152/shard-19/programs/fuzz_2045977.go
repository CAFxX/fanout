package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint32) uint32 {
    var v7 int8 = -1
    v8 := []int{-495}
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 int8 = -1
    var v2 int16 = -1
    v3 := make(map[string]bool)
    var v4 int = -2147483648
    v5 := make(map[string]string)
    v6 := make(map[string]uint32)
    v9 := []int{}
    v10 := GenS1{F0: 237}
    v11 := v5[("\n\t" + "a")]
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(v2, fi12)
        v9 = append(v9, 127)
        var v13 int8 = (v1 & (v1 + (v1 ^ v1)))
        _ = fi12
        _ = v13
    }
    if (v11 != ("" + ("a" + "a"))) {
        if false {
            _ = len(v9)
        }
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int32 = -1
            _ = fi14
            _ = v15
        }
        v16 := GenS1{F0: 127}
        _ = v16
    }
    var v17 bool = false
    if v17 {
        var v18 int8 = v1
        for fi19 := 0; fi19 < 17; fi19++ {
            if v17 {
                fmt.Println(v17, v18, v4)
            }
            for fi20 := 0; fi20 < 14; fi20++ {
                fmt.Println(fi20, v18)
                fmt.Println(v2)
                v4 += (fi20 | 4294967295)
                _ = fi20
            }
            v18 = 42
            _ = fi19
        }
        var v21 int64 = 1
        _ = v18
        _ = v21
    }
    v22 := GenS1{F0: 128}
    fmt.Println(v11, v11, v1)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
