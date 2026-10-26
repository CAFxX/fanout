package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0(p0 int8) uint8 {
    var v5 uint = 0
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{128}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi6
        _ = v7
    }
    v8 := []int{153, 4294967295, 32768}
    _ = p0
    _ = v5
    _ = v8
    return 127
}
func genFn1(p0 uint, p1 int64) int64 {
    var v9 string = ""
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(p1, p1, fi10)
        if true {
            v9 = (("hello" + (v9 + ((v9 + ("a" + "")) + ("foo bar" + "\n\t")))) + v9)
            p1 += (p1 >> (((p0 / (((p0 << ((p0) & 63))) | 1))) & 63))
            fmt.Println(p1)
        } else {
            if true {
                var v11 int64 = (p1 << (((p0 % (((p0 - p0)) | 1))) & 63))
                _ = v11
            }
            switch fi10 {
            case 32767:
                fmt.Println(p0, v9)
            case 2147483647:
                v12 := []int{-401}
                _ = v12
            default:
                var v13 bool = true
                _ = v13
            }
            v14 := []int{}
            _ = v14
        }
        p1 = 42
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    var v4 int = 2147483647
    var v15 uint64 = 951
    if (0 < 1) {
        switch (v4 ^ (v4 | (v4 + v4))) {
        case 9223372036854775807:
            v16 := GenS1{F0: -97, F1: ("\n\t" + "")}
            _ = v16
        default:
            var v17 int64 = 655
            _ = v17
        }
    }
    v18 := GenS2{F0: 0, F1: 737}
    if true {
        for fi19 := 0; fi19 < 14; fi19++ {
            v15 += v15
            v4 = (v4 % (((fi19 - fi19)) | 1))
            var v20 int8 = -1
            _ = fi19
            _ = v20
        }
    }
    delete(v2, (v4 - 128))
    var v21 int64 = genFn1(65535, 438)
    var v22 uint = 4294967295
    for fi23 := 0; fi23 < 9; fi23++ {
        _ = len(v3)
        v24 := GenS2{F0: (v4 ^ v4), F1: 42}
        fmt.Println(v15)
        _ = fi23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
