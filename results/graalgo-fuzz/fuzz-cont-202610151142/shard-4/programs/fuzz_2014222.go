package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int16; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint; F1 int8; F2 uint8 }
func genFn0(p0 string, p1 uint8) int16 {
    var v5 int32 = 2147483647
    v6 := GenS1{F0: 42, F1: ((2147483647 >= 32768) && (true || true))}
    for fi7 := 0; fi7 < 16; fi7++ {
        if (true || false) {
            var v8 uint32 = 469
            switch (p0 + p0) {
            case "a":
                fmt.Println(v5, v5)
            }
            _ = v8
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 905
}

func main() {
    v1 := []int{}
    var v2 uint16 = 32767
    v3 := make(map[uint32]int)
    var v4 uint16 = (v2 >> ((4294967295) & 63))
    v9 := v3[2]
    v4 = 2
    delete(v3, 205)
    v10 := v3[65535]
    _ = len(v3)
    switch "hello" {
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "foo bar":
        v4 = 65535
    case "\n\t":
        switch "\n\t" {
        case "":
            v1 = append(v1, (-(v10)))
        case "hello":
            if false {
                _ = len(v1)
            }
        default:
            var v11 int16 = 0
            _ = v11
        }
    }
    switch "a" {
    case "":
        v12 := GenS2{F0: 42, F1: v9}
        _ = v12
    }
    if (false || false) {
        switch "" {
        case "":
            for fi13 := 0; fi13 < 9; fi13++ {
                fmt.Println(v4)
                v14 := GenS3{F0: -1, F1: false, F2: 1}
                switch "\n\t" {
                case "\n\t":
                    if ((^(fi13)) == v10) {
                        for fi15 := 0; fi15 < 9; fi15++ {
                            v16 := GenS4{F0: 2, F1: -1, F2: 128}
                            _ = fi15
                            _ = v16
                        }
                        v2 += (^(v2))
                        if (false || false) {
                            v2 = v4
                            var v17 float32 = 1.5
                            v2 += v4
                            _ = v17
                        }
                    }
                }
                _ = fi13
                _ = v14
            }
        }
        v9 += v9
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v18 int32 = 1
    v4 -= (v4 ^ v4)
    var v19 int = (^(v9))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
