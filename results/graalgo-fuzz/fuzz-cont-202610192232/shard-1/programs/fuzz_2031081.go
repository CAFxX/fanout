package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
func genFn0(p0 bool, p1 bool) uint16 {
    var v6 uint = 0
    if p0 {
        var v7 uint16 = 993
        var v8 uint32 = 4294967295
        var v9 float64 = 1.5
        _ = v7
        _ = v8
        _ = v9
    } else {
        p0 = ("foo bar" <= "hello")
        var v10 bool = ((true || p0) || false)
        v10 = ("foo bar" != ("hello" + ("a" + "\n\t")))
        _ = v10
    }
    var v11 uint = v6
    var v12 int8 = -1
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v6
    return 0
}

func main() {
    var v1 uint = 373
    v2 := make(map[uint32]int)
    var v3 float32 = (0.1 * 1.5)
    var v4 uint = (v1 * (v1 * (v1 << (((^(v1))) & 63))))
    var v5 int32 = -2147483648
    v1 += (v1 % ((693) | 1))
    var v13 int16 = 808
    v14 := GenS1{F0: (v5 / ((-50) | 1))}
    v2[4294967295] = 9223372036854775807
    switch "a" {
    case "\n\t":
        for fi15 := 0; fi15 < 17; fi15++ {
            fmt.Println(v13)
            switch "\n\t" {
            case "":
                v5 += v5
            case "a":
                for fi16 := 0; fi16 < 19; fi16++ {
                    if (!(255 < 112)) {
                        v17 := GenS1{F0: (v5 * v5)}
                        _ = v17
                    }
                    if (false || (65535 != 303)) {
                        v18 := []int{fi15, fi15, 0}
                        _ = v18
                    } else {
                        v19 := []int{(-(fi15)), (fi15 / (((fi15 + (fi16 / (((^(fi16))) | 1)))) | 1)), (fi16 - (fi16 << ((0) & 63))), -2147483648}
                        var v20 int = (fi15 | (fi15 + fi15))
                        fmt.Println(v3)
                        _ = v19
                        _ = v20
                    }
                    v5 += (v5 - v5)
                    _ = fi16
                }
            }
            fmt.Println(v1)
            _ = fi15
        }
    case "foo bar":
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int32 = (v5 >> ((160) & 63))
            for fi23 := 0; fi23 < 18; fi23++ {
                fmt.Println(fi21, v3)
                v5 = -1
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    case "":
        if true {
            var v24 uint64 = 4294967295
            v25 := GenS1{F0: (v5 * (v5 & -550))}
            _ = v24
            _ = v25
        } else {
            v26 := GenS2{F0: 1}
            v5 = v5
            var v27 uint64 = 18446744073709551615
            _ = v26
            _ = v27
        }
    default:
        v28 := []int{}
        _ = v28
    }
    switch -111 {
    case 32768:
        v4 += (v1 % ((v4) | 1))
    case -135:
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 int16 = (v13 % ((913) | 1))
            v3 += float32(1.5)
            _ = fi29
            _ = v30
        }
    }
    var v31 float32 = v3
    v32 := GenS3{F0: 4294967295}
    fmt.Println(v4, v3, v3)
    for fi33 := 0; fi33 < 12; fi33++ {
        v34 := GenS2{F0: 670}
        v2[1] = -148
        _ = len(v2)
        _ = fi33
        _ = v34
    }
    v13 -= (v13 / (((v13 & v13)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
