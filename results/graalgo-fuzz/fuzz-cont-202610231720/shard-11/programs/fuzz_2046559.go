package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 string }
func genFn0(p0 uint16, p1 int16, p2 bool) int16 {
    v5 := []int{}
    var v6 uint8 = 255
    if p2 {
        switch "hello" {
        case "hello":
            if ((v6 >> ((255) & 63)) <= 1) {
                p0 -= 42
            } else {
                v7 := GenS1{F0: 452, F1: 42, F2: 9223372036854775807}
                for fi8 := 0; fi8 < 8; fi8++ {
                    v9 := GenS2{F0: 18446744073709551615, F1: p2}
                    _ = fi8
                    _ = v9
                }
                _ = v7
            }
        case "\n\t":
            var v10 int = 32767
            _ = v10
        case "":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v5 = append(v5, -149)
        v11 := GenS2{F0: 18446744073709551615, F1: p2}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p1
}
func genFn1() uint {
    v12 := GenS1{F0: -9, F1: 618, F2: -460}
    v13 := GenS2{F0: 735, F1: (!true)}
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint64 = 168
        for fi16 := 0; fi16 < 18; fi16++ {
            fmt.Println(v15)
            _ = fi16
        }
        fmt.Println(v15, fi14)
        _ = fi14
        _ = v15
    }
    var v17 bool = true
    _ = v12
    _ = v13
    _ = v17
    return 862
}

func main() {
    var v1 int64 = 96
    v2 := make(map[string]string)
    var v3 bool = (1 != 4294967295)
    var v4 string = ""
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := GenS3{F0: v4}
        v2[v4] = v4
        _ = fi18
        _ = v19
    }
    delete(v2, "")
    fmt.Println(v3)
    var v20 uint = 18446744073709551615
    var v21 int16 = -32768
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := []int{(fi22 << (((v20 ^ 2)) & 63)), (fi22 << ((567) & 63)), (fi22 << ((0) & 63)), 9223372036854775807}
        _ = fi22
        _ = v23
    }
    switch v4 {
    case "\n\t":
        fmt.Println(v21, v20, v3)
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
