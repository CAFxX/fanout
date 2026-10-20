package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 bool; F1 string }
func genFn0(p0 int32, p1 uint16, p2 uint32) uint32 {
    if (2147483647 >= -783) {
        v6 := []int{-2147483648}
        _ = v6
    }
    fmt.Println(p2, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    return p2
}
func genFn1(p0 uint64) int16 {
    fmt.Println(p0)
    if ((false || (128 < 9223372036854775807)) || true) {
        var v7 int32 = 42
        _ = v7
    }
    _ = p0
    return 624
}

func main() {
    var v1 int64 = 42
    var v2 bool = (true && (-2147483648 == 957))
    var v3 uint16 = 0
    var v4 uint8 = 127
    var v5 int = -782
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: 21, F1: 32768, F2: 0}
            switch "foo bar" {
            case "":
                if (91 > 38) {
                    v11 := GenS1{F0: 0, F1: -9223372036854775808, F2: (v1 | (v1 & v1))}
                    _ = v11
                }
            case "hello":
                if v2 {
                    v12 := []int{fi8, (fi9 + (-(v5))), (fi8 * fi8)}
                    _ = v12
                } else {
                    fmt.Println(fi8, v2)
                    v13 := GenS1{F0: 975, F1: fi9, F2: (v1 * -1)}
                    _ = v13
                }
                var v14 bool = (v1 >= v1)
                _ = v14
            case "\n\t":
                var v15 uint16 = 2
                _ = v15
            }
            v4 += 128
            v16 := GenS2{F0: -44, F1: v5}
            _ = fi9
            _ = v10
            _ = v16
        }
        var v17 float32 = 1.5
        switch ("hello" + "foo bar") {
        case "a":
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := []int{fi18, (-(fi8)), 1, (fi18 % ((fi8) | 1))}
                _ = fi18
                _ = v19
            }
        case "":
            if false {
                var v20 uint64 = 18446744073709551615
                _ = v20
            }
        }
        _ = fi8
        _ = v17
    }
    v21 := []int{(v5 | v5), v5, v5, v5}
    var v22 int = (v5 >> ((2) & 63))
    for fi23 := 0; fi23 < 7; fi23++ {
        v22 = (fi23 % (((-(fi23))) | 1))
        fmt.Println(v22, fi23, v22)
        var v24 int64 = (v1 % (((v1 << ((272) & 63))) | 1))
        _ = fi23
        _ = v24
    }
    if len(v21) > 0 {
        v21[0] = -825
    }
    _ = len(v21)
    switch "" {
    case "\n\t":
        switch v5 {
        case -9223372036854775808:
            v5 = v5
        }
    }
    v25 := GenS1{F0: 305, F1: v5, F2: (v1 << ((65535) & 63))}
    if ((1 >= -175) || ("\n\t" >= "hello")) {
        fmt.Println(v4, v2, v3)
    }
    var v26 int8 = -88
    fmt.Println(v3, v26)
    for fi27 := 0; fi27 < 6; fi27++ {
        if v2 {
            for fi28 := 0; fi28 < 3; fi28++ {
                switch int(v4) {
                case 0:
                    v29 := GenS3{F0: false, F1: ""}
                    _ = v29
                default:
                    if v2 {
                        v1 = v1
                        var v30 bool = (!v2)
                        _ = v30
                    }
                }
                _ = fi28
            }
        } else {
            fmt.Println(v4, v5, v5)
        }
        _ = fi27
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        v21 = append(v21, (fi31 - (fi31 % (((v5 - -977)) | 1))))
        if len(v21) > 0 {
            v21[0] = (v5 | v5)
        }
        _ = fi31
    }
    if len(v21) > 0 {
        v21[0] = (-(v22))
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        var v33 float32 = 3.14159
        var v34 uint64 = 1
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
