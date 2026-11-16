package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint16; F2 int16 }
func genFn0(p0 uint, p1 uint16, p2 int) uint {
    switch (p2 + 977) {
    case 0:
        var v7 uint = (p0 % ((p0) | 1))
        _ = v7
    case -653:
        var v8 uint8 = 127
        _ = v8
    case 1:
        if (!false) {
            v9 := GenS1{F0: -9223372036854775808, F1: 1000}
            _ = v9
        } else {
            v10 := []int{}
            p2 += p2
            switch "hello" {
            case "hello":
                if true {
                    if false {
                        var v11 uint32 = 65535
                        p1 -= (^(p1))
                        var v12 int = p2
                        _ = v11
                        _ = v12
                    } else {
                        if false {
                            var v13 bool = false
                            var v14 int32 = -1
                            var v15 uint16 = (p1 - p1)
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                        v16 := GenS2{F0: 2}
                        _ = v16
                    }
                    v10 = append(v10, p2)
                    var v17 int8 = -113
                    _ = v17
                }
            }
            _ = v10
        }
    default:
        v18 := []int{(p2 << ((0) & 63)), p2, p2, p2}
        _ = v18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 uint16 = 65535
        v20 = 2
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 int = -77
    v4 := make(map[uint32]bool)
    var v5 uint16 = 0
    v6 := []int{v3, (v3 - v3), v3, (v3 & v3)}
    switch ("\n\t" + "foo bar") {
    case "a":
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 float64 = (float64(3.14159) - 3.14159)
            for fi23 := 0; fi23 < 15; fi23++ {
                fmt.Println(fi23)
                v3 -= int(v5)
                v3 += 127
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    }
    v3 += (v3 + v3)
    fmt.Println(v5, v3)
    fmt.Println(v5, v3)
    v3 = 127
    v5 = (^(v5))
    fmt.Println(v3)
    v3 = v3
    v5 = 2
    v24 := GenS3{F0: -1, F1: (v5 + 41), F2: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
