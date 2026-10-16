package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint32 }
func genFn0(p0 int16) int {
    v4 := []int{}
    var v5 int8 = 0
    var v6 int32 = -2147483648
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -812
}

func main() {
    var v1 int = 735
    var v2 string = "a"
    var v3 uint64 = 0
    var v7 int32 = -173
    for fi8 := 0; fi8 < 14; fi8++ {
        if true {
            v9 := []int{(fi8 >> ((5) & 63)), v1, genFn0(int16(v3))}
            v7 = v7
            _ = v9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            v7 = (v7 & 1)
            if (false || false) {
                v11 := GenS1{F0: v2, F1: genFn0(-32768), F2: 42}
                _ = v11
            }
            var v12 uint16 = 32767
            _ = fi10
            _ = v12
        }
        var v13 int32 = (v7 & v7)
        _ = fi8
        _ = v13
    }
    switch v1 {
    case 1:
        if (!true) {
            v3 += (v3 | 2)
            if true {
                v14 := []int{v1, genFn0(32767), genFn0(-826)}
                _ = v14
            }
            fmt.Println(v2, v3)
        } else {
            v7 -= (-(v7))
            v15 := []int{}
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = v15
        }
    case 127:
        fmt.Println(v2, v7)
    }
    fmt.Println(v3, v2)
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS1{F0: ("" + ""), F1: fi16, F2: 2147483648}
        var v18 string = ("" + (("foo bar" + ((v2 + (v2 + "a")) + "\n\t")) + "\n\t"))
        v19 := []int{v1}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := []int{42}
    v20 = append(v20, genFn0(-620))
    v21 := GenS1{F0: v2, F1: v1, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v7
}
