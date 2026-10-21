package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint32, p1 string) uint32 {
    var v4 bool = false
    for fi5 := 0; fi5 < 9; fi5++ {
        p0 = p0
        v6 := GenS1{F0: 2147483647, F1: 128}
        _ = fi5
        _ = v6
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    return (p0 << ((2) & 63))
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    var v7 uint16 = 742
    v1 = 0
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 float64 = (1.5 - 3.14159)
        _ = fi8
        _ = v9
    }
    v7 = v7
    if true {
        if false {
            fmt.Println(v1)
            if true {
                var v10 uint64 = 890
                _ = v10
            } else {
                if (65535 <= 18446744073709551615) {
                    if false {
                        var v11 int16 = v1
                        _ = v11
                    }
                    v1 = v1
                }
            }
        }
    } else {
        var v12 uint16 = v7
        v7 = v7
        v13 := []int{0, -592, -479, 127}
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 42, F1: 2}
    v15 := []int{4294967295, 2147483647, 32767}
    v7 = 2
    fmt.Println(v7, v1)
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    v16 := GenS1{F0: -846, F1: 1}
    var v17 uint = 65535
    if len(v15) > 0 {
        v15[0] = -841
    }
    var v18 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v7
}
