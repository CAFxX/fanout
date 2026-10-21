package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS6 struct { F0 int64 }
func genFn0() int32 {
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 int8 = -72
        _ = fi4
        _ = v5
    }
    var v6 uint8 = 42
    _ = v6
    return 1
}
func genFn1(p0 int8, p1 int, p2 uint16) bool {
    var v7 int8 = (p0 << ((18446744073709551615) & 63))
    for fi8 := 0; fi8 < 15; fi8++ {
        if true {
            if ((v7 ^ 0) < p0) {
                p0 = (p0 | 1)
                var v9 uint = 154
                v10 := GenS1{F0: 576}
                _ = v10
                _ = v9
            } else {
                v11 := []int{2147483647, fi8, (fi8 * -2147483648)}
                _ = v11
            }
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS1{F0: fi8}
                _ = fi12
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := []int{32768, (-(fi14))}
                p1 -= p1
                if false {
                    switch (fi8 >> ((255) & 63)) {
                    case -2147483648:
                        v16 := GenS1{F0: fi8}
                        _ = v16
                    }
                    var v17 int16 = 42
                    var v18 bool = true
                    _ = v17
                    _ = v18
                } else {
                    p0 -= (p0 / (((p0 >> ((445) & 63))) | 1))
                    var v19 uint32 = 42
                    if true {
                        p0 -= p0
                        var v20 int16 = int16(fi8)
                        _ = v20
                    }
                    _ = v19
                }
                _ = fi14
                _ = v15
            }
            v7 = 127
        }
        v21 := GenS2{F0: 65}
        _ = fi8
        _ = v21
    }
    if (441 != 1) {
        p2 += 668
        p0 = p0
    } else {
        p1 = (p1 + p1)
        fmt.Println(p1, v7)
        p1 -= (p1 / (((p1 / ((p1) | 1))) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return false
}

func main() {
    var v1 float32 = -2.25
    var v2 uint64 = 775
    v3 := make(map[int]int)
    v1 += v1
    v22 := GenS3{F0: 42, F1: 42}
    v23 := GenS4{F0: -9223372036854775808, F1: true}
    v2 -= 4294967295
    v24 := GenS2{F0: 1}
    for fi25 := 0; fi25 < 12; fi25++ {
        switch "a" {
        case "":
            v26 := []int{128, 4294967295}
            _ = v26
        }
        _ = fi25
    }
    if true {
        v27 := GenS5{F0: 653, F1: 2, F2: v2}
        v1 = float32((3.14159 - 3.14159))
        if false {
            if true {
                v28 := GenS6{F0: 0}
                _ = v28
            }
        }
        _ = v27
    }
    switch "a" {
    case "\n\t":
        v2 += (v2 << ((255) & 63))
    case "a":
        if false {
            fmt.Println(v2)
            var v29 uint64 = 521
            _ = v29
        }
    case "foo bar":
        var v30 string = "hello"
        _ = v30
    }
    v1 -= v1
    v31 := []int{487, 17, -1, -336}
    var v32 int64 = 0
    v33 := GenS4{F0: v32, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
