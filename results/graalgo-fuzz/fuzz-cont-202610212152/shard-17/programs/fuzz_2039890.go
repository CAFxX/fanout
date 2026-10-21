package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint, p1 uint16, p2 uint8) uint {
    for fi7 := 0; fi7 < 15; fi7++ {
        if true {
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 string = "hello"
                v9 = "hello"
                if (!true) {
                    v10 := []int{(fi8 * fi7), fi7}
                    fmt.Println(p1)
                    _ = v10
                } else {
                    p2 -= (p2 - 127)
                }
                _ = fi8
                _ = v9
            }
            p2 = p2
            var v11 int32 = 2147483647
            _ = v11
        }
        v12 := []int{}
        _ = fi7
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 int32 = 1
        v15 := GenS1{F0: 918, F1: 553}
        switch "foo bar" {
        case "":
            switch (("" + "foo bar") + ("hello" + "foo bar")) {
            case "foo bar":
                p2 = (p2 << (((p0 - (p0 | (p0 / ((uint(p2)) | 1))))) & 63))
            }
        case "hello":
            var v16 int64 = -1
            _ = v16
        case "\n\t":
            var v17 string = "hello"
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return 145
}
func genFn1(p0 int, p1 int32, p2 uint32) uint32 {
    var v18 string = "\n\t"
    fmt.Println(p0)
    var v19 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return (^(p2))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = -89
    v3 := []string{}
    v4 := []int{42, -1, -495, -1}
    v5 := make(map[int]bool)
    var v6 uint64 = 886
    if true {
        _ = len(v4)
    }
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v20 float32 = 0.0
        v20 = v20
        _ = v20
    }
    var v21 uint8 = uint8(v6)
    v2 += 1
    var v22 uint32 = 65535
    v23 := GenS2{F0: v22, F1: 127}
    v22 = 788
    v22 += v22
    _ = len(v4)
    switch -237 {
    case 127:
        v2 -= (v2 - (v2 ^ 1))
    case -656:
        v24 := GenS3{F0: 142}
        _ = v24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 uint8 = (v21 / (((v21 & (v21 >> ((genFn0(4294967295, 1, v21)) & 63)))) | 1))
        _ = fi25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
