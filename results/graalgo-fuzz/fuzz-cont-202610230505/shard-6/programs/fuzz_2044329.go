package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 string, p1 uint) int64 {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 1000}
        p1 += (p1 % ((p1) | 1))
        _ = fi5
        _ = v6
    }
    p0 = p0
    p1 = p1
    _ = p0
    _ = p1
    return -550
}

func main() {
    var v1 string = "a"
    v2 := []uint32{4294967295, 65535}
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    v2 = append(v2, 1)
    v2 = append(v2, 2147483648)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v7 := GenS2{F0: 65535, F1: 0, F2: 32767}
    v8 := GenS3{F0: 4294967295, F1: 100}
    v1 = ("" + ((v1 + v1) + v1))
    _ = len(v3)
    var v9 uint = 942
    if len(v2) > 0 {
        v2[0] = 975
    }
    switch 114 {
    case 32768:
        v10 := GenS2{F0: 65535, F1: 1, F2: 0}
        _ = v10
    }
    v11 := GenS4{F0: -116}
    var v12 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
