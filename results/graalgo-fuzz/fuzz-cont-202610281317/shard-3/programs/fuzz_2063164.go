package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int64; F1 int32; F2 int8 }
type GenS3 struct { F0 int; F1 uint32 }
func genFn0(p0 uint64) uint64 {
    v5 := []int{128, 2147483647}
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint32 = 964
        if (true && false) {
            _ = len(v5)
            p0 = 42
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return p0
}
func genFn1(p0 uint, p1 string, p2 int64) uint32 {
    switch p1 {
    case "hello":
        fmt.Println(p1)
    }
    if true {
        fmt.Println(p0, p2, p2)
        v8 := GenS1{F0: (912 > 0), F1: 52}
        var v9 int16 = 0
        _ = v8
        _ = v9
    } else {
        v10 := []int{2147483647}
        _ = v10
    }
    var v11 uint8 = 8
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 2
}

func main() {
    v1 := []uint32{1, 2, 65535}
    v2 := []string{}
    var v3 int = -952
    var v4 bool = true
    v12 := GenS1{F0: (42 >= 1), F1: 127}
    v1 = append(v1, genFn1(42, "hello", 469))
    for fi13 := 0; fi13 < 5; fi13++ {
        fmt.Println(v3, v3, fi13)
        fmt.Println(v4, v3)
        _ = fi13
    }
    v4 = (366 <= 77)
    var v14 int16 = -411
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v14, v14)
    fmt.Println(v3)
    switch "a" {
    case "hello":
        fmt.Println(v14)
    case "":
        var v15 uint16 = 65535
        _ = v15
    }
    var v16 uint16 = 301
    v17 := GenS1{F0: v4, F1: 2}
    switch (v3 | (v3 | v3)) {
    case 127:
        var v18 uint16 = v16
        _ = v18
    case -145:
        for fi19 := 0; fi19 < 11; fi19++ {
            v20 := GenS2{F0: 0, F1: -2147483648, F2: -128}
            _ = fi19
            _ = v20
        }
    default:
        if v4 {
            v14 += (-(v14))
            var v21 int64 = 1
            _ = v21
        } else {
            var v22 string = "a"
            if true {
                var v23 bool = (983 <= 2147483647)
                v24 := GenS1{F0: v4, F1: 128}
                v23 = false
                _ = v23
                _ = v24
            } else {
                var v25 uint64 = 2
                for fi26 := 0; fi26 < 8; fi26++ {
                    switch (fi26 % ((fi26) | 1)) {
                    case -680:
                        fmt.Println(fi26, v14, fi26)
                    }
                    _ = fi26
                }
                _ = v25
            }
            _ = v22
        }
    }
    v3 = (v3 & (v3 >> ((42) & 63)))
    v14 = (v14 << ((255) & 63))
    v27 := GenS3{F0: v3, F1: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
