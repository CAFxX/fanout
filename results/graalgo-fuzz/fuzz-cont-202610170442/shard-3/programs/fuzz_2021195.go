package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int32 }
func genFn0(p0 string) int32 {
    switch (p0 + (((p0 + p0) + p0) + p0)) {
    case "\n\t":
        switch p0 {
        case "":
            var v5 uint64 = 0
            _ = v5
        }
    case "foo bar":
        for fi6 := 0; fi6 < 18; fi6++ {
            switch fi6 {
            case 1:
                p0 = p0
            }
            v7 := []int{fi6, fi6, 42, (fi6 | -572)}
            var v8 uint8 = 25
            _ = fi6
            _ = v7
            _ = v8
        }
    case "":
        v9 := []int{-837, 42, 42}
        _ = v9
    }
    fmt.Println(p0)
    v10 := []int{32768, 834, 2147483647}
    _ = p0
    _ = v10
    return -126
}
func genFn1(p0 int8, p1 int16, p2 int16) uint32 {
    p2 = (p2 * (p2 % ((568) | 1)))
    p1 = (p2 | p2)
    var v11 int16 = p1
    var v12 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 2
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    v1 = 0
    fmt.Println(v1)
    v1 = (v1 / ((127) | 1))
    var v13 uint64 = 188
    var v14 uint32 = 2
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 string = "\n\t"
            var v18 int64 = 0
            var v19 int8 = -128
            v20 := GenS1{F0: uint32(v13), F1: 416, F2: -348}
            var v21 int32 = 1
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
            _ = v20
            _ = v21
        }
        v14 += 42
        _ = fi15
    }
    v1 += v1
    v14 = 2147483648
    for fi22 := 0; fi22 < 7; fi22++ {
        v1 -= v1
        v1 = v1
        delete(v4, (v14 >> ((188) & 63)))
        _ = fi22
    }
    v23 := GenS1{F0: (v14 >> ((55) & 63)), F1: 65535, F2: genFn0("a")}
    v24 := []int{-197, 179, 128}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
