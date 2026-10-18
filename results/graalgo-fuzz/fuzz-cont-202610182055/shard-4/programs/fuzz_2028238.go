package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int8 }
func genFn0(p0 uint16, p1 bool, p2 int) int32 {
    var v7 uint16 = 255
    if (247 > 1) {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := []int{fi8, (p2 - (p2 - (fi8 - p2))), (-(fi8)), fi8}
            var v10 int8 = 76
            for fi11 := 0; fi11 < 16; fi11++ {
                if false {
                    p0 += (^(v7))
                }
                var v12 uint = 42
                _ = fi11
                _ = v12
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        v13 := []int{}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 uint8, p1 int8, p2 uint8) int8 {
    var v14 uint16 = 1
    var v15 uint64 = 42
    fmt.Println(v14)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return p1
}

func main() {
    var v1 uint8 = 55
    var v2 string = "foo bar"
    var v3 uint8 = 0
    v4 := []int{2147483647, -2147483648, -1, -2147483648}
    var v5 string = ("" + (v2 + "a"))
    v6 := make(map[string]uint32)
    fmt.Println(v5, v3)
    fmt.Println(v5, v2, v2)
    v16 := GenS1{F0: v1, F1: false, F2: 0}
    for fi17 := 0; fi17 < 8; fi17++ {
        fmt.Println(v3, v5, v5)
        v18 := GenS1{F0: 42, F1: true, F2: 65535}
        v5 = ((v2 + v2) + v5)
        _ = fi17
        _ = v18
    }
    v2 = v5
    if (v5 != v2) {
        v2 = v5
        for fi19 := 0; fi19 < 1; fi19++ {
            v20 := GenS2{F0: 931, F1: 2}
            _ = fi19
            _ = v20
        }
    } else {
        if false {
            if (0 >= 1) {
                _ = len(v4)
            }
        }
        v21 := GenS3{F0: uint32(v1), F1: 439, F2: 100}
        _ = v21
    }
    if (866 <= 1) {
        for fi22 := 0; fi22 < 1; fi22++ {
            fmt.Println(v3, fi22, v3)
            fmt.Println(v5, v2, v5)
            fmt.Println(v1)
            _ = fi22
        }
    } else {
        _ = len(v4)
        for fi23 := 0; fi23 < 18; fi23++ {
            for fi24 := 0; fi24 < 12; fi24++ {
                v4 = append(v4, 4294967295)
                _ = fi24
            }
            _ = fi23
        }
        v2 = v5
    }
    fmt.Println(v1, v5)
    for fi25 := 0; fi25 < 5; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            var v27 int8 = -1
            _ = fi26
            _ = v27
        }
        v1 = (v3 / (((^(v1))) | 1))
        _ = fi25
    }
    var v28 bool = (!false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
