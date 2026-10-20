package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
func genFn0() uint64 {
    v4 := GenS1{F0: 127}
    var v5 int = -860
    var v6 int64 = 1
    var v7 uint8 = 0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return uint64(v7)
}
func genFn1(p0 int) int64 {
    if (!(p0 >= (p0 & p0))) {
        for fi8 := 0; fi8 < 12; fi8++ {
            p0 -= (-(p0))
            switch (p0 | 445) {
            case -880:
                var v9 float32 = float32(1e10)
                _ = v9
            case -9223372036854775808:
                fmt.Println(fi8, fi8)
            }
            _ = fi8
        }
        if (!false) {
            for fi10 := 0; fi10 < 10; fi10++ {
                if (2147483647 != 42) {
                    if (!(2147483647 > -2147483648)) {
                        if false {
                            if (2 > 42) {
                                fmt.Println(fi10, fi10)
                                fmt.Println(p0, fi10)
                            } else {
                                fmt.Println(fi10, fi10)
                                v11 := []int{p0, p0, 9223372036854775807}
                                v11 = append(v11, p0)
                                _ = v11
                            }
                            v12 := GenS1{F0: 1}
                            v13 := []int{42, (fi10 & (fi10 >> ((2) & 63))), -1}
                            _ = v12
                            _ = v13
                        }
                    } else {
                        var v14 uint8 = 232
                        v15 := GenS1{F0: 0}
                        p0 = p0
                        _ = v14
                        _ = v15
                    }
                    var v16 uint64 = 144
                    for fi17 := 0; fi17 < 8; fi17++ {
                        v16 = v16
                        _ = fi17
                    }
                    _ = v16
                }
                _ = fi10
            }
            switch "a" {
            case "hello":
                v18 := GenS2{F0: false}
                _ = v18
            case "":
                var v19 uint32 = 170
                _ = v19
            case "\n\t":
                v20 := []int{2147483647}
                _ = v20
            }
            v21 := GenS2{F0: false}
            _ = v21
        } else {
            v22 := []int{p0}
            _ = v22
        }
        p0 = p0
    } else {
        v23 := GenS2{F0: false}
        _ = v23
    }
    var v24 int64 = 253
    var v25 int8 = 78
    v26 := []int{int(v24), (p0 << ((0) & 63))}
    v27 := GenS3{F0: 255}
    _ = p0
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    return 1
}

func main() {
    var v1 float32 = 0.0
    var v2 uint64 = 2
    v3 := []int{1}
    for fi28 := 0; fi28 < 11; fi28++ {
        fmt.Println(fi28, v2, v1)
        for fi29 := 0; fi29 < 18; fi29++ {
            v1 += (v1 / v1)
            _ = fi29
        }
        _ = fi28
    }
    v1 = float32(0.1)
    v3 = append(v3, 786)
    v2 -= 636
    fmt.Println(v1)
    var v30 int64 = 0
    fmt.Println(v30, v30)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
}
