package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 string; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 int64 }
func genFn0(p0 int64) int {
    var v7 uint16 = 649
    v8 := GenS1{F0: false}
    switch "\n\t" {
    case "":
        v9 := GenS1{F0: false}
        _ = v9
    case "\n\t":
        if true {
            if true {
                var v10 int8 = 100
                v11 := []int{-2147483648, 42, 4294967295}
                p0 -= 9223372036854775807
                _ = v10
                _ = v11
            }
            v12 := GenS1{F0: (false || true)}
            v13 := GenS2{F0: 1, F1: ("a" + "foo bar"), F2: 128}
            _ = v12
            _ = v13
        } else {
            if (42 >= 42) {
                v7 += (^(v7))
            } else {
                for fi14 := 0; fi14 < 10; fi14++ {
                    v15 := GenS3{F0: 4294967295}
                    v16 := GenS4{F0: 0, F1: (p0 & (p0 << ((uint(v7)) & 63)))}
                    var v17 int = (-(fi14))
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                v7 -= v7
            }
        }
    case "foo bar":
        v18 := []int{-829, -365, 42}
        _ = v18
    }
    fmt.Println(p0)
    var v19 uint = 10
    _ = p0
    _ = v19
    _ = v7
    _ = v8
    return 636
}
func genFn1(p0 bool, p1 int32, p2 uint16) int8 {
    switch "" {
    case "a":
        p0 = p0
    case "foo bar":
        v20 := []int{-9223372036854775808, -2147483648, 127}
        _ = v20
    }
    v21 := []int{899}
    var v22 int8 = 42
    v23 := GenS3{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    _ = v23
    return (^(v22))
}

func main() {
    var v1 string = "foo bar"
    var v2 uint16 = 363
    v3 := make(map[int]int)
    var v4 float32 = (1.5 - 0.1)
    v5 := make(map[uint32]int)
    var v6 uint32 = 2
    v24 := GenS1{F0: true}
    v6 += v6
    v25 := GenS1{F0: true}
    for fi26 := 0; fi26 < 17; fi26++ {
        v27 := []int{}
        _ = fi26
        _ = v27
    }
    v28 := []int{42, 4294967295, -9223372036854775808, -83}
    var v29 uint = uint(v6)
    fmt.Println(v29, v4, v2)
    v5[(^(v6))] = -9223372036854775808
    fmt.Println(v4, v1, v1)
    var v30 bool = (1 <= -9223372036854775808)
    v31 := GenS2{F0: 42, F1: ("" + v1), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
