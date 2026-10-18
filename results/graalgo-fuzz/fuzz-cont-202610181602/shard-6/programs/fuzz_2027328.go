package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 bool }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int, p1 uint, p2 bool) int8 {
    if (p2 || (!p2)) {
        v4 := GenS1{F0: "a", F1: 0, F2: (221 == 128)}
        p2 = true
        var v5 uint16 = 610
        _ = v4
        _ = v5
    }
    var v6 uint32 = 0
    v7 := []int{p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 string) uint8 {
    v8 := GenS2{F0: 42}
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 126
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]uint32)
    v3 := []uint32{1, 4294967295, 2, 0}
    fmt.Println(v1)
    for fi9 := 0; fi9 < 8; fi9++ {
        switch (fi9 << ((548) & 63)) {
        case 102:
            var v10 uint = 42
            _ = v10
        case 9223372036854775807:
            v11 := GenS2{F0: 956}
            _ = v11
        }
        _ = fi9
    }
    _ = len(v2)
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            v1 -= (v1 << ((0) & 63))
            v14 := GenS1{F0: "a", F1: 2147483647, F2: true}
            v15 := GenS1{F0: "", F1: 42, F2: (false && (2 == 492))}
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(fi12, fi12)
        _ = fi12
    }
    var v16 uint64 = 0
    if (!false) {
        v1 -= v1
    } else {
        if ((v1 ^ v1) != 42) {
            var v17 string = "\n\t"
            v18 := GenS2{F0: 1}
            _ = v17
            _ = v18
        } else {
            v19 := GenS1{F0: "foo bar", F1: -68, F2: (false && true)}
            if true {
                for fi20 := 0; fi20 < 7; fi20++ {
                    var v21 float64 = ((-2.25 / (float64(3.14159) / 3.14159)) - 1e10)
                    if true {
                        for fi22 := 0; fi22 < 1; fi22++ {
                            delete(v2, 32768)
                            _ = fi22
                        }
                    }
                    var v23 bool = (v16 <= (v16 % ((v16) | 1)))
                    _ = fi20
                    _ = v21
                    _ = v23
                }
            }
            _ = v19
        }
        var v24 int16 = -1
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 string = "hello"
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    v1 = v1
    var v27 uint32 = 2147483648
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v28 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
