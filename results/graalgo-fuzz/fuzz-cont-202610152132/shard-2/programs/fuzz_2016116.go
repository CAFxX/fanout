package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 bool }
type GenS3 struct { F0 uint; F1 int; F2 uint16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint64, p1 uint, p2 int32) int64 {
    fmt.Println(p1)
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 uint = p1
        fmt.Println(v5, p1, fi4)
        fmt.Println(fi4, p2)
        _ = fi4
        _ = v5
    }
    var v6 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -205
}
func genFn1(p0 int32) uint64 {
    var v7 int32 = p0
    for fi8 := 0; fi8 < 13; fi8++ {
        switch "foo bar" {
        case "a":
            var v9 int8 = -22
            _ = v9
        case "foo bar":
            v10 := GenS1{F0: ("" != ("\n\t" + "\n\t")), F1: 179, F2: 0}
            _ = v10
        default:
            switch "a" {
            case "\n\t":
                v11 := GenS1{F0: false, F1: 0, F2: -32768}
                _ = v11
            default:
                if (213 != 1) {
                    var v12 uint64 = 0
                    _ = v12
                }
            }
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            v7 = p0
            if (0.1 < 0.1) {
                var v14 int8 = -1
                p0 -= p0
                v15 := []int{32768, fi8, (fi13 + (fi13 % ((fi8) | 1)))}
                _ = v14
                _ = v15
            } else {
                var v16 uint16 = 65535
                v7 += (^(v7))
                if false {
                    var v17 int32 = p0
                    _ = v17
                } else {
                    v7 = (v7 | v7)
                    var v18 int = fi13
                    if (0 != 42) {
                        var v19 string = ""
                        _ = v19
                    } else {
                        v7 += (v7 + (v7 - v7))
                        v20 := GenS1{F0: true, F1: 42, F2: 1000}
                        var v21 uint32 = 849
                        _ = v20
                        _ = v21
                    }
                    _ = v18
                }
                _ = v16
            }
            p0 -= -750
            v7 += (v7 << ((142) & 63))
            _ = fi13
        }
        v22 := []int{fi8, (-(fi8)), fi8}
        _ = fi8
        _ = v22
    }
    var v23 float32 = ((0.1 + ((0.1 * float32((3.14159 * -2.25))) / (1e10 - (0.0 + (3.14159 - (1.5 / 3.14159)))))) / 0.1)
    v7 += (p0 * v7)
    fmt.Println(v7, v23)
    _ = p0
    _ = v23
    _ = v7
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 65535
    var v3 int8 = 127
    var v24 int64 = genFn0(genFn1(1), 255, 2147483647)
    v25 := []int{782, 32768, 2147483647, -9223372036854775808}
    var v26 uint8 = 0
    v3 = -1
    if false {
        if true {
            var v27 int = -432
            switch -215 {
            case -2147483648:
                if (!true) {
                    var v28 int = (v27 / (((v27 * (v27 + v27))) | 1))
                    _ = v28
                }
            case 9223372036854775807:
                v26 = v26
            case -9223372036854775808:
                var v29 float32 = (0.0 / 1.0)
                _ = v29
            }
            v25 = append(v25, (^(v27)))
            _ = v27
        }
        var v30 int = 698
        _ = v30
    } else {
        if len(v25) > 0 {
            v25[0] = 889
        }
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
        _ = fi31
    }
    v32 := GenS2{F0: 0, F1: v26, F2: true}
    if (0.0 < (0.0 - ((0.1 * 3.14159) / 1.0))) {
        v33 := GenS3{F0: 18446744073709551615, F1: -16, F2: 32767}
        _ = v33
    }
    var v34 float32 = float32(3.14159)
    v35 := GenS4{F0: (v2 | 65535)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v34
    _ = v35
}
