package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int32; F1 int; F2 uint8 }
func genFn0(p0 uint8, p1 int32) uint8 {
    var v4 int16 = -32768
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := []int{fi6}
            _ = len(v7)
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 11; fi8++ {
            switch 42 {
            case -1:
                v4 = (v4 >> ((51) & 63))
                if (1 > 0) {
                    fmt.Println(p1, p1, p0)
                } else {
                    if (!(false && true)) {
                        v9 := []int{(^(fi5))}
                        v10 := GenS1{F0: -317}
                        _ = v10
                        _ = v9
                    }
                }
                v4 = (v4 << ((1) & 63))
                if (42 >= -98) {
                    p1 += -1
                    v4 -= v4
                } else {
                    fmt.Println(fi5, fi8, v4)
                    p0 = 96
                }
            case 128:
                var v11 uint8 = (p0 >> ((0) & 63))
                _ = v11
            case 127:
                v12 := []int{(fi8 >> ((18446744073709551615) & 63)), 9223372036854775807, fi5}
                _ = v12
            }
            p0 = p0
            var v13 int = (fi8 ^ (^(fi5)))
            _ = fi8
            _ = v13
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 int64) int64 {
    fmt.Println(p0)
    v14 := []int{-15}
    _ = p0
    _ = v14
    return (p0 & (p0 & 970))
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 0
    var v3 uint32 = 1
    switch 9223372036854775807 {
    case 1:
        v3 = (v3 * v3)
    case 2147483647:
        if false {
            if false {
                v15 := []int{9223372036854775807, 0, 9223372036854775807, 9223372036854775807}
                _ = v15
            }
            var v16 int = 128
            _ = v16
        }
    case 127:
        v17 := GenS1{F0: -819}
        _ = v17
    }
    v2 -= v2
    if false {
        v2 = v2
        v2 -= uint16(v3)
    }
    var v18 uint64 = 227
    v1["\n\t"] = ("\n\t" + "a")
    fmt.Println(v18, v2)
    v19 := GenS2{F0: true, F1: 490}
    if (!false) {
        if (true || (v18 == (v18 ^ (^(v18))))) {
            var v20 int32 = 1
            v21 := GenS3{F0: -779, F1: (v3 | v3)}
            var v22 uint32 = (v3 * uint32(v2))
            _ = v20
            _ = v21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        v3 += (v3 & uint32(v18))
        v24 := GenS3{F0: -1, F1: (v3 | v3)}
        _ = fi23
        _ = v24
    }
    v25 := GenS4{F0: 733, F1: 32767}
    v26 := GenS5{F0: 747, F1: -1, F2: 90}
    v27 := GenS1{F0: 483}
    switch ("a" + "hello") {
    case "foo bar":
        v2 += (v2 * v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
