package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint8) string {
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(fi8, p0)
            switch (fi8 - 4294967295) {
            case 246:
                v9 := []int{-9223372036854775808, fi8, 9223372036854775807}
                _ = v9
            case -1:
                if true {
                    var v10 int32 = -775
                    fmt.Println(fi7)
                    v11 := []int{fi7, 1, (fi8 << ((651) & 63)), 9223372036854775807}
                    _ = v10
                    _ = v11
                } else {
                    var v12 uint64 = 47
                    p0 += p0
                    var v13 int32 = 118
                    _ = v12
                    _ = v13
                }
            case 9223372036854775807:
                var v14 uint64 = 615
                _ = v14
            default:
                v15 := GenS1{F0: 2}
                _ = v15
            }
            _ = fi8
        }
        switch "hello" {
        case "hello":
            p0 = p0
        case "foo bar":
            fmt.Println(p0, fi7)
        }
        _ = fi7
    }
    p0 = p0
    _ = p0
    return ("" + "")
}
func genFn1(p0 bool, p1 int8) int64 {
    var v16 int = 1
    var v17 uint32 = 875
    v18 := []int{(^(v16)), v16}
    p1 = p1
    fmt.Println(v16)
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    return -30
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]uint32)
    var v3 int16 = 1000
    var v4 int32 = 42
    var v5 uint = 42
    var v6 uint64 = 1
    fmt.Println(v3)
    fmt.Println(v4, v5, v5)
    var v19 uint16 = 65535
    var v20 string = "foo bar"
    v1 = (v1 % (((v1 % ((2147483648) | 1))) | 1))
    var v21 int64 = -459
    var v22 uint16 = (v19 & 918)
    switch 128 {
    case 1:
        v23 := GenS2{F0: 42}
        _ = v23
    case -2147483648:
        if true {
            switch -2147483648 {
            case -82:
                v1 += v1
            case -2147483648:
                v5 += (^(v5))
            case 32768:
                v24 := []int{}
                _ = v24
            default:
                var v25 int = 863
                _ = v25
            }
        } else {
            switch v20 {
            case "foo bar":
                var v26 uint64 = 1
                _ = v26
            }
        }
    case 9223372036854775807:
        var v27 int32 = v4
        _ = v27
    }
    var v28 uint16 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
