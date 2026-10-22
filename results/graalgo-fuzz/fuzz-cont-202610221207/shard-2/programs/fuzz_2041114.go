package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint32; F2 int32 }
func genFn0(p0 uint8, p1 bool) int {
    var v5 int8 = 0
    v5 += v5
    v6 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    var v1 int32 = 0
    var v2 uint8 = 119
    v3 := make(map[uint32]uint32)
    var v4 string = "a"
    if true {
        v1 += v1
    } else {
        v7 := GenS1{F0: -49}
        _ = v7
    }
    switch 127 {
    case 999:
        v4 = v4
    case 32768:
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(fi8, fi8, fi8)
            _ = fi8
        }
    case -278:
        switch v4 {
        case "":
            var v9 int8 = 42
            _ = v9
        default:
            var v10 bool = (1e10 > (1e10 + 0.1))
            _ = v10
        }
    }
    var v11 uint8 = (v2 << ((139) & 63))
    v12 := GenS2{F0: 42}
    var v13 float32 = 3.14159
    v4 = (((v4 + "\n\t") + v4) + "\n\t")
    switch (v4 + (v4 + "hello")) {
    case "hello":
        if (((1 != 4294967295) && false) && false) {
            var v14 int8 = -100
            if false {
                if true {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        for fi16 := 0; fi16 < 15; fi16++ {
                            var v17 int32 = (v1 / ((v1) | 1))
                            _ = fi16
                            _ = v17
                        }
                        v18 := []int{}
                        switch (fi15 - (fi15 - (fi15 % ((fi15) | 1)))) {
                        case 9223372036854775807:
                            delete(v3, 886)
                        }
                        _ = fi15
                        _ = v18
                    }
                    if true {
                        v1 -= v1
                        if false {
                            fmt.Println(v11)
                            v19 := GenS1{F0: genFn0(127, (0 <= 0))}
                            fmt.Println(v2)
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 15; fi20++ {
                            var v21 int = (fi20 * -835)
                            _ = fi20
                            _ = v21
                        }
                    }
                    var v22 uint32 = 2
                    _ = v22
                }
                var v23 int8 = v14
                _ = v23
            } else {
                v11 -= v11
                var v24 uint16 = 1
                _ = v24
            }
            v25 := []int{-513, genFn0((v11 / ((v2) | 1)), true), 128}
            _ = v14
            _ = v25
        }
    default:
        if false {
            v26 := v3[2]
            var v27 float64 = 1.5
            _ = v26
            _ = v27
        }
    }
    v28 := v3[42]
    if (0 >= -2147483648) {
        fmt.Println(v2, v1, v11)
    }
    if true {
        switch ("a" + (v4 + v4)) {
        case "foo bar":
            v29 := GenS3{F0: 707, F1: 2, F2: (v1 << ((519) & 63))}
            _ = v29
        }
        if (9223372036854775807 != -9223372036854775808) {
            var v30 uint64 = 1
            _ = v30
        } else {
            v31 := []int{}
            _ = v31
        }
        fmt.Println(v28, v11)
    }
    v32 := []int{2147483647, 9223372036854775807, 642, 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v4
}
