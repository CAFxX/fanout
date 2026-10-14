package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint8 }
func genFn0(p0 int8, p1 int32, p2 uint32) int16 {
    v7 := GenS1{F0: 0, F1: 1}
    for fi8 := 0; fi8 < 9; fi8++ {
        if true {
            if (false || (!((!(false && (!true))) && false))) {
                v9 := []int{-831}
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = v9
            } else {
                var v10 int64 = 9223372036854775807
                _ = v10
            }
            v11 := GenS2{F0: 0}
            v12 := []int{fi8, 32767}
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 string = "foo bar"
                if (!((!((^(p2)) >= (p2 & (p2 - p2)))) || ((65535 <= 4294967295) || true))) {
                    p2 += 4294967295
                    v15 := GenS2{F0: 1}
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
            p1 -= (p1 >> ((850) & 63))
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}
func genFn1() int32 {
    switch "hello" {
    case "\n\t":
        v16 := GenS2{F0: 18446744073709551615}
        _ = v16
    }
    v17 := []int{}
    v18 := GenS2{F0: 959}
    var v19 uint64 = 17
    _ = v17
    _ = v18
    _ = v19
    return -2147483648
}

func main() {
    var v1 float64 = 1e10
    var v2 int = 32768
    v3 := make(map[int]int)
    v4 := []uint32{}
    v5 := []uint32{0, 0, 0, 310}
    var v6 uint64 = 146
    v1 = 0.0
    var v20 string = "hello"
    var v21 string = v20
    var v22 uint16 = 984
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 int = 42
        var v25 uint16 = 65535
        var v26 int64 = -9223372036854775808
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    v27 := GenS3{F0: -1, F1: 109}
    var v28 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
