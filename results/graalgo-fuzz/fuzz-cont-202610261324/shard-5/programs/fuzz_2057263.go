package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int32 }
func genFn0() int {
    var v4 int = 698
    v4 += 9223372036854775807
    v5 := []int{-902, 32767, (v4 & (-(v4))), (^(v4))}
    _ = v4
    _ = v5
    return v4
}
func genFn1(p0 int64, p1 uint32, p2 uint32) uint8 {
    fmt.Println(p1)
    switch "\n\t" {
    case "":
        var v6 float32 = (1e10 * 1.5)
        _ = v6
    }
    v7 := GenS1{F0: -398, F1: 327, F2: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 155
}

func main() {
    var v1 uint64 = 816
    v2 := make(map[int]bool)
    var v3 float32 = 1e10
    v8 := []int{32768, -971, 127}
    var v9 uint = 688
    v10 := GenS2{F0: 1, F1: 1}
    var v11 float64 = float64(v3)
    v1 = (v1 | v1)
    fmt.Println(v11, v11)
    _ = len(v2)
    v2[-495] = false
    v12 := GenS3{F0: 2147483647}
    switch int(v1) {
    case 32767:
        if true {
            for fi13 := 0; fi13 < 3; fi13++ {
                v1 += (^(v1))
                if false {
                    v14 := GenS1{F0: 812, F1: v9, F2: 0}
                    if (!(-2147483648 != 2147483647)) {
                        fmt.Println(v11)
                    }
                    v2[-13] = true
                    _ = v14
                } else {
                    v8 = append(v8, (fi13 % (((fi13 / ((-2147483648) | 1))) | 1)))
                    var v15 bool = (((211 != 127) || true) || true)
                    _ = v15
                }
                _ = fi13
            }
            var v16 int32 = 397
            _ = v16
        } else {
            v1 = (v1 % (((v1 / (((v1 >> (((v9 + v9)) & 63))) | 1))) | 1))
        }
    case 665:
        if true {
            v1 -= v1
            switch -619 {
            case 9223372036854775807:
                fmt.Println(v1)
            case 127:
                v17 := GenS2{F0: 379, F1: 193}
                _ = v17
            }
        } else {
            v9 += v9
        }
    }
    if (!(true && ((v9 <= v9) || true))) {
        var v18 uint8 = 164
        _ = v18
    }
    v19 := GenS1{F0: -32768, F1: v9, F2: 9223372036854775807}
    var v20 string = ((("a" + ("a" + "\n\t")) + "") + "\n\t")
    for fi21 := 0; fi21 < 3; fi21++ {
        switch (^(fi21)) {
        case 9223372036854775807:
            v2[fi21] = false
        case -1:
            for fi22 := 0; fi22 < 6; fi22++ {
                v3 -= v3
                _ = fi22
            }
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
