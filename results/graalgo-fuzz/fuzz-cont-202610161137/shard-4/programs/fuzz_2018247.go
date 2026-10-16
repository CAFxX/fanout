package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int; F2 bool }
type GenS4 struct { F0 uint8; F1 uint; F2 uint32 }
func genFn0(p0 uint16, p1 uint16, p2 uint16) uint {
    fmt.Println(p2)
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 0, F1: 550}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 string = "foo bar"
    var v2 uint = 2
    v3 := make(map[uint32]int)
    var v4 uint32 = 101
    var v5 int32 = 42
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(v2)
        for fi9 := 0; fi9 < 6; fi9++ {
            if (!false) {
                if ("" <= ("foo bar" + "foo bar")) {
                    v10 := []int{(fi9 << ((2) & 63)), 32767, fi9}
                    fmt.Println(fi9, fi8, v1)
                    _ = v10
                } else {
                    var v11 float32 = (float32(0.0) + 3.14159)
                    v12 := v3[v4]
                    var v13 bool = true
                    _ = v11
                    _ = v12
                    _ = v13
                }
                if false {
                    delete(v3, v4)
                }
            }
            var v14 uint32 = (v4 >> (((v2 >> ((v2) & 63))) & 63))
            v1 = (v1 + v1)
            _ = fi9
            _ = v14
        }
        _ = fi8
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint64 = 0
        v17 := GenS2{F0: v4}
        _ = fi15
        _ = v16
        _ = v17
    }
    switch ("\n\t" + v1) {
    case "":
        v18 := GenS2{F0: (^(v4))}
        _ = v18
    case "foo bar":
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 float32 = 1.5
            if (255 > 255) {
                v20 -= v20
            } else {
                var v21 uint32 = 940
                if true {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        fmt.Println(fi22, v2, v1)
                        _ = fi22
                    }
                    var v23 uint16 = 134
                    v4 = v4
                    _ = v23
                } else {
                    fmt.Println(v5)
                    v21 = 2
                    fmt.Println(v2, v5)
                }
                _ = v21
            }
            fmt.Println(v1, v1, v2)
            _ = fi19
            _ = v20
        }
    case "a":
        fmt.Println(v5, v5, v2)
    default:
        fmt.Println(v4, v4)
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        for fi25 := 0; fi25 < 7; fi25++ {
            v26 := GenS3{F0: 644, F1: (fi25 - 0), F2: ("foo bar" <= v1)}
            v1 = "foo bar"
            _ = fi25
            _ = v26
        }
        var v27 uint8 = 0
        v2 = 975
        _ = fi24
        _ = v27
    }
    var v28 uint32 = 872
    v29 := GenS2{F0: v28}
    v4 = 2147483647
    v30 := GenS4{F0: 76, F1: v2, F2: 65535}
    var v31 int16 = 32767
    var v32 uint = v2
    fmt.Println(v4, v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
